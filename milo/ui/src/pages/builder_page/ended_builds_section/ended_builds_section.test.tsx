// Copyright 2023 The LUCI Authors.
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

import { render, screen, fireEvent, act } from '@testing-library/react';
import { DateTime } from 'luxon';

import {
  Build,
  BuildStatus,
  BuildsService,
  SearchBuildsResponse,
} from '@/common/services/buildbucket';
import { FakeContextProvider } from '@/testing_tools/fakes/fake_context_provider';

import { EndedBuildsSection } from './ended_builds_section';
import { EndedBuildsTable } from './ended_builds_table';

jest.mock('./ended_builds_table', () => ({
  EndedBuildsTable: jest.fn(() => <></>),
}));

const builderId = {
  bucket: 'buck',
  builder: 'builder',
  project: 'proj',
};

function createMockBuild(id: string): Build {
  return {
    id,
    builder: builderId,
    status: BuildStatus.Success,
    createTime: '2020-01-01',
  };
}

const builds = [
  createMockBuild('1234'),
  createMockBuild('2345'),
  createMockBuild('3456'),
  createMockBuild('4567'),
  createMockBuild('5678'),
];

const pages: {
  [timestamp: string]: { [pageToken: string]: SearchBuildsResponse };
} = {
  '': {
    '': {
      builds: builds.slice(0, 2),
      nextPageToken: 'page2',
    },
    page2: {
      builds: builds.slice(2, 4),
      nextPageToken: 'page3',
    },
    page3: {
      builds: builds.slice(4, 5),
    },
  },
  '2020-02-02T02:02:02.000+00:00': {
    '': {
      builds: builds.slice(1, 3),
      nextPageToken: 'page2',
    },
    page2: {
      builds: builds.slice(3, 5),
    },
  },
};

describe('EndedBuildsSection', () => {
  let endedBuildsTableMock: jest.MockedFunctionDeep<typeof EndedBuildsTable>;

  beforeEach(() => {
    jest.useFakeTimers();
    jest
      .spyOn(BuildsService.prototype, 'searchBuilds')
      .mockImplementation(
        async ({ pageToken, predicate }) =>
          pages[predicate.createTime?.endTime || ''][pageToken || '']
      );
    endedBuildsTableMock = jest.mocked(EndedBuildsTable);
  });

  afterEach(() => {
    jest.useRealTimers();
  });

  test('should clear page tokens after date filter is reset', async () => {
    render(
      <FakeContextProvider>
        <EndedBuildsSection builderId={builderId} />
      </FakeContextProvider>
    );
    await act(() => jest.runAllTimersAsync());

    expect(endedBuildsTableMock).toHaveBeenCalledWith(
      {
        endedBuilds: builds.slice(0, 2),
        isLoading: false,
      },
      expect.anything()
    );
    endedBuildsTableMock.mockClear();

    fireEvent.click(screen.getByText('Next Page'));
    await act(() => jest.runAllTimersAsync());
    expect(screen.getByText('Previous Page')).toBeEnabled();
    expect(endedBuildsTableMock).toHaveBeenCalledWith(
      {
        endedBuilds: builds.slice(2, 4),
        isLoading: false,
      },
      expect.anything()
    );
    endedBuildsTableMock.mockClear();

    jest.setSystemTime(
      DateTime.fromISO('2020-02-02T02:02:02.000+00:00').toMillis()
    );
    fireEvent.click(screen.getByTestId('CalendarIcon'));
    fireEvent.click(screen.getByText('Today'));

    await act(() => jest.runAllTimersAsync());
    // Prev page tokens are purged.
    expect(screen.getByText('Previous Page')).toBeDisabled();
    expect(endedBuildsTableMock).toHaveBeenCalledWith(
      {
        endedBuilds: builds.slice(1, 3),
        isLoading: false,
      },
      expect.anything()
    );
  });
});