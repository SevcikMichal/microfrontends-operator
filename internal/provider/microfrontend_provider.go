/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

// TODO: Improve the implementation of MicroFrontendProvider currently it is not very sofiticated

import (
	"sync"

	"github.com/SevcikMichal/microfrontends-controller/contract"
	"github.com/SevcikMichal/microfrontends-controller/internal/model"
	"k8s.io/apimachinery/pkg/types"
)

const (
	appTransfersKey     = "app-transfers"
	preloadTrnasfersKey = "preload-transfers"
	contextTransfersKey = "context-transfers"
)

type MicroFrontendProvider struct {
	MicroFrontendTransferStorage *sync.Map
}

func (r *MicroFrontendProvider) SetMicroFrontendConfig(key types.UID, microFrontendConfig *model.MicroFrontendConfig) {
	r.updateWebAppTransfers(key, microFrontendConfig)
	r.updatePreloadTransfers(key, microFrontendConfig)
	r.updateContextTransfers(key, microFrontendConfig)
}

func (r *MicroFrontendProvider) DeleteMicroFrontendConfig(key types.UID) {
	r.deleteWebAppTransfers(key)
	r.deletePreloadTransfers(key)
	r.deleteContextTransfers(key)
}

func (r *MicroFrontendProvider) GetMicroFrontendConfigTransfer() *contract.MicroFrontendConfigurationTransfer {
	frontendConfigTransfer := &contract.MicroFrontendConfigurationTransfer{
		Apps:     r.getWebAppTransfers(),
		Preload:  r.getPreloadTransfers(),
		Contexts: r.getContextTransfers(),
		// Anonymous: ,
		// User: ,
	}

	return frontendConfigTransfer
}
