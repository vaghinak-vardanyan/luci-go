// Copyright 2020 The LUCI Authors.
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

import { MobxLitElement } from '@adobe/lit-mobx';
import { css, html } from 'lit';
import { customElement } from 'lit/decorators.js';
import { computed, makeObservable, observable } from 'mobx';
import { fromPromise } from 'mobx-utils';
import { useSearchParams } from 'react-router-dom';

import '../../components/image_diff_viewer';
import '../../components/status_bar';
import '../../components/dot_spinner';
import { consumer } from '../../libs/context';
import { reportRenderError } from '../../libs/error_handler';
import { unwrapObservable } from '../../libs/milo_mobx_utils';
import { ArtifactIdentifier, constructArtifactName } from '../../services/resultdb';
import { consumeStore, StoreInstance } from '../../store';
import commonStyle from '../../styles/common_style.css';
import { consumeArtifactIdent } from './artifact_page_layout';

/**
 * Renders an image diff artifact set, including expected image, actual image
 * and image diff.
 */
// TODO(weiweilin): improve error handling.
@customElement('milo-image-diff-artifact-page')
@consumer
export class ImageDiffArtifactPageElement extends MobxLitElement {
  static get properties() {
    return {
      expectedArtifactId: {
        type: String,
      },
      actualArtifactId: {
        type: String,
      },
    };
  }

  @observable.ref
  @consumeStore()
  store!: StoreInstance;

  @observable.ref
  @consumeArtifactIdent()
  artifactIdent!: ArtifactIdentifier;

  @observable.ref expectedArtifactId!: string;
  @observable.ref actualArtifactId!: string;

  @computed private get diffArtifactName() {
    return constructArtifactName({ ...this.artifactIdent });
  }
  @computed private get expectedArtifactName() {
    return constructArtifactName({ ...this.artifactIdent, artifactId: this.expectedArtifactId });
  }
  @computed private get actualArtifactName() {
    return constructArtifactName({ ...this.artifactIdent, artifactId: this.actualArtifactId });
  }

  @computed
  private get diffArtifact$() {
    if (!this.store.services.resultDb) {
      return fromPromise(Promise.race([]));
    }
    return fromPromise(this.store.services.resultDb.getArtifact({ name: this.diffArtifactName }));
  }
  @computed private get diffArtifact() {
    return unwrapObservable(this.diffArtifact$, null);
  }

  @computed
  private get expectedArtifact$() {
    if (!this.store.services.resultDb) {
      return fromPromise(Promise.race([]));
    }
    return fromPromise(this.store.services.resultDb.getArtifact({ name: this.expectedArtifactName }));
  }
  @computed private get expectedArtifact() {
    return unwrapObservable(this.expectedArtifact$, null);
  }

  @computed
  private get actualArtifact$() {
    if (!this.store.services.resultDb) {
      return fromPromise(Promise.race([]));
    }
    return fromPromise(this.store.services.resultDb.getArtifact({ name: this.actualArtifactName }));
  }
  @computed private get actualArtifact() {
    return unwrapObservable(this.actualArtifact$, null);
  }

  @computed get isLoading() {
    return !this.expectedArtifact || !this.actualArtifact || !this.diffArtifact;
  }

  constructor() {
    super();
    makeObservable(this);
  }

  protected render = reportRenderError(this, () => {
    if (this.isLoading) {
      return html`<div id="loading-spinner" class="active-text">Loading <milo-dot-spinner></milo-dot-spinner></div>`;
    }

    return html`
      <milo-image-diff-viewer
        .expected=${this.expectedArtifact}
        .actual=${this.actualArtifact}
        .diff=${this.diffArtifact}
      >
      </milo-image-diff-viewer>
    `;
  });

  static styles = [
    commonStyle,
    css`
      :host {
        display: block;
      }

      #loading-spinner {
        margin: 20px;
      }
    `,
  ];
}

declare global {
  // eslint-disable-next-line @typescript-eslint/no-namespace
  namespace JSX {
    interface IntrinsicElements {
      'milo-image-diff-artifact-page': {
        expectedArtifactId: string;
        actualArtifactId: string;
      };
    }
  }
}

export function ImageDiffArtifactPage() {
  const [search] = useSearchParams();
  const expectedArtifactId = search.get('expectedArtifactId')!;
  const actualArtifactId = search.get('actualArtifactId')!;
  return (
    <milo-image-diff-artifact-page
      expectedArtifactId={expectedArtifactId}
      actualArtifactId={actualArtifactId}
    ></milo-image-diff-artifact-page>
  );
}