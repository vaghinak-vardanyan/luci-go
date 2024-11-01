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

import { useEffect } from 'react';
import { useOutletContext } from 'react-router-dom';

import bassFavicon from '@/common/assets/favicons/bass-32.png';
import { RecoverableErrorBoundary } from '@/common/components/error_handling';
import { PageMeta } from '@/common/components/page_meta';
import { GenFeedbackUrlArgs } from '@/common/tools/utils';
import { TrackLeafRoutePageView } from '@/generic_libs/components/google_analytics';

export const DeviceListPage = () => {
  return (
    <>
      <PageMeta title="Streamlined Fleet UI" favicon={bassFavicon} />
      Hello world
    </>
  );
};

export function Component() {
  const { setFeedbackUrlArgs } = useOutletContext() as {
    feedbackUrlArgs: GenFeedbackUrlArgs;
    setFeedbackUrlArgs: React.Dispatch<
      React.SetStateAction<GenFeedbackUrlArgs>
    >;
  };

  useEffect(() => {
    setFeedbackUrlArgs({ bugComponent: '1664178' });
    return () => setFeedbackUrlArgs(undefined);
  }, [setFeedbackUrlArgs]);

  return (
    <TrackLeafRoutePageView contentGroup="fleet-console-device-list">
      <RecoverableErrorBoundary
        // See the documentation for `<LoginPage />` for why we handle error
        // this way.
        key="fleet-device-list-page"
      >
        <DeviceListPage />
      </RecoverableErrorBoundary>
    </TrackLeafRoutePageView>
  );
}
