// Copyright 2023 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package realmsinternals

import (
	"errors"
	"sort"

	"go.chromium.org/luci/common/data/sortby"
	"go.chromium.org/luci/common/data/stringset"
	realmsconf "go.chromium.org/luci/common/proto/realms"
	"go.chromium.org/luci/server/auth/service/protocol"

	"google.golang.org/protobuf/proto"
)

var (
	// ErrFinalized is used when the ConditionsSet has already been finalized
	// and further modifications are attempted.
	ErrFinalized = errors.New("conditions set has already been finalized")
)

//	ConditionsSet normalizes and dedups conditions, maps them to integers.
//	Assumes all incoming realmsconf.Condition are immutable and dedups
//	them by pointer, as well as by normalized values.
//	Also assumes the set of all possible *objects* ever passed to indexes(...) was
//	also passed to addCond(...) first (so it could build id => index map).
//
// This makes hot indexes(...) function fast by allowing to lookup ids instead
// of (potentially huge) protobuf message values.
type ConditionsSet struct {
	// normalized is a mapping from a serialized normalized protocol.Condition
	// to a pair (normalized *protocol.Condition, its unique index)
	normalized map[string]*conditionMapTuple

	// indexMapping from serialized realms_config to its index.
	indexMapping map[*realmsconf.Condition]uint32

	// finalized is true if finalize() was called, see finalize for more info.
	finalized bool
}

// conditionMapTuple is to represent the entries of normalized, reflects
// what index a Condition is tied to.
type conditionMapTuple struct {
	cond *protocol.Condition
	idx  uint32
}

// addCond adds a *Condition from realms.cfg definition to the set if it's
// not already there.
//
// Returns ErrFinalized -- if set has already been finalized
func (cs *ConditionsSet) addCond(cond *realmsconf.Condition) error {
	if cs.finalized {
		return ErrFinalized
	}
	if _, ok := cs.indexMapping[cond]; ok {
		return nil
	}

	norm := &protocol.Condition{}
	var attr string

	if cond.GetRestrict() != nil {
		condValues := make([]string, len(cond.GetRestrict().GetValues()))
		copy(condValues, cond.GetRestrict().GetValues())
		condSet := stringset.NewFromSlice(condValues...)
		attr = cond.GetRestrict().GetAttribute()
		norm.Op = &protocol.Condition_Restrict{
			Restrict: &protocol.Condition_AttributeRestriction{
				Attribute: attr,
				Values:    condSet.ToSortedSlice(),
			},
		}
	}

	idx := uint32(len(cs.normalized))
	if condTup, ok := cs.normalized[conditionKey(norm)]; ok {
		idx = condTup.idx
	}
	cs.normalized[conditionKey(norm)] = &conditionMapTuple{norm, idx}
	cs.indexMapping[cond] = idx
	return nil
}

// conditionKey generates a key by serializing a protocol.Condition.
func conditionKey(cond *protocol.Condition) string {
	key, err := proto.Marshal(cond)
	if err != nil {
		return ""
	}
	return string(key)
}

// sortConditions sorts a given conditions slice by attribute first
// then by values.
func sortConditions(conds []*protocol.Condition) {
	sort.Slice(conds, sortby.Chain{
		func(i, j int) bool {
			return conds[i].GetRestrict().GetAttribute() < conds[j].GetRestrict().GetAttribute()
		},
		func(i, j int) bool {
			iValsLen, jValsLen := len(conds[i].GetRestrict().GetValues()), len(conds[j].GetRestrict().GetValues())
			iVals, jVals := conds[i].GetRestrict().GetValues(), conds[j].GetRestrict().GetValues()
			if iValsLen == jValsLen {
				for idx, iVal := range iVals {
					if iVal == jVals[idx] {
						continue
					}
					return iVal < jVals[idx]
				}
			}
			return iValsLen < jValsLen
		},
	}.Use)
}

// finalize finalizes the set by preventing any future addCond calls.
//
// Sorts the list of stored conditions by attribute first then by values.
// returns the final sorted list of protocol.Condition. Returns nil if
// ConditionSet is not finalized.
//
// Indexes returned by indexes() will refer to the indexes in this list.
func (cs *ConditionsSet) finalize() []*protocol.Condition {
	if cs.finalized {
		return nil
	}
	cs.finalized = true

	conds := []*protocol.Condition{}
	for _, curr := range cs.normalized {
		conds = append(conds, curr.cond)
	}

	sortConditions(conds)

	oldToNew := map[uint32]uint32{}

	for idx, cond := range conds {
		old := cs.normalized[conditionKey(cond)]
		oldToNew[old.idx] = uint32(idx)
	}

	for key, old := range cs.indexMapping {
		cs.indexMapping[key] = oldToNew[old]
	}

	return conds
}

// indexes returns a sorted slice of indexes
//
// Can be called only after finalize(). All given conditions must have previously
// been put into the set via addCond(). The returned indexes can have fewer
// elements if some conditions in conds are equivalent.
//
// The returned indexes is essentially a compact encoding of the overall AND
// condition expression in a binding.
func (cs *ConditionsSet) indexes(conds []*realmsconf.Condition) []uint32 {
	if !cs.finalized {
		return nil
	}
	if conds == nil {
		return nil
	}
	if len(conds) == 1 {
		if idx, ok := cs.indexMapping[conds[0]]; ok {
			return []uint32{idx}
		}
		return nil
	}

	indexesSet := emptyIndexSet()

	for _, cond := range conds {
		v, ok := cs.indexMapping[cond]
		if !ok {
			return nil
		}
		indexesSet.add(v)
	}

	return indexesSet.toSortedSlice()
}

// indexSet is a set data structure for managing indexes when expanding realms and permissions.
type indexSet struct {
	set map[uint32]struct{}
}

// add adds a given uint32 to the index set.
func (is *indexSet) add(v uint32) {
	is.set[v] = struct{}{}
}

// IndexSetFromSlice converts a given slice of indexes and returns an IndexSet from them.
func IndexSetFromSlice(src []uint32) *indexSet {
	res := emptyIndexSet()
	for _, val := range src {
		res.set[val] = struct{}{}
	}
	return res
}

// emptyIndexSet initializes and returns an empty IndexSet.
func emptyIndexSet() *indexSet {
	return &indexSet{make(map[uint32]struct{})}
}

// toSortedSlice converts and IndexSet to a slice and then sorts the indexes, returning the
// result.
func (is *indexSet) toSortedSlice() []uint32 {
	res := make([]uint32, 0, len(is.set))
	for k := range is.set {
		res = append(res, k)
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	return res
}