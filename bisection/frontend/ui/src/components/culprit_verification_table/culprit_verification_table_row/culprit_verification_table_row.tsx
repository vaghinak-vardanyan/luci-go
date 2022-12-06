// Copyright 2022 The LUCI Authors.
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

import './culprit_verification_table_row.css';

import Link from '@mui/material/Link';
import TableCell from '@mui/material/TableCell';
import TableRow from '@mui/material/TableRow';

import { getCommitShortHash } from '../../../tools/commit_formatters';
import { HeuristicSuspect } from '../../../services/luci_bisection';
import { VerificationDetailsTable } from '../../culprits_table/culprits_table';

interface Props {
  suspect: HeuristicSuspect;
}

export const CulpritVerificationTableRow = ({ suspect }: Props) => {
  const {
    gitilesCommit,
    reviewUrl,
    reviewTitle,
    verificationDetails
  } = suspect;

  let suspectDescription = getCommitShortHash(gitilesCommit.id);
  if (reviewTitle) {
    suspectDescription += `: ${reviewTitle}`;
  }

  return (
    <>
      <TableRow data-testid='culprit_verification_table_row'>
        <TableCell className='overview-cell'>
          <Link
            href={reviewUrl}
            target='_blank'
            rel='noreferrer'
            underline='always'
          >
            {suspectDescription}
          </Link>
        </TableCell>
        <TableCell className='overview-cell'>
          Heuristic
        </TableCell>
        <TableCell className='overview-cell'>
          {verificationDetails.status}
        </TableCell>
        <TableCell className='overview-cell'>
        <VerificationDetailsTable details={verificationDetails} />
        </TableCell>
      </TableRow>
    </>
  );
};