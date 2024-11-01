// Copyright 2024 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { useContext } from 'react';

import { DepthCtx, OffsetsCtx, SizeRecorderCtx } from './context';

export function useSizeRecorder() {
  const ctx = useContext(SizeRecorderCtx);
  if (ctx === undefined) {
    throw new Error(
      'useSizeRecorder can only be used in a QueuedStickyContextProvider',
    );
  }

  return ctx;
}

export function useOffsets() {
  const ctx = useContext(OffsetsCtx);
  if (ctx === undefined) {
    throw new Error(
      'useOffsets can only be used in a QueuedStickyContextProvider',
    );
  }

  return ctx;
}

export function useDepth() {
  const ctx = useContext(DepthCtx);
  if (ctx === undefined) {
    throw new Error(
      'useDepth can only be used in a QueuedStickyContextProvider',
    );
  }
  return ctx;
}
