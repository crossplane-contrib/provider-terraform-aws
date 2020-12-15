/*
	Copyright 2019 The Crossplane Authors.

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

package v1alpha1

import (
	"fmt"

	"github.com/zclconf/go-cty/cty"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*PinpointApnsVoipSandboxChannel)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a PinpointApnsVoipSandboxChannel.")
	}
	return EncodePinpointApnsVoipSandboxChannel(*r), nil
}

func EncodePinpointApnsVoipSandboxChannel(r PinpointApnsVoipSandboxChannel) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodePinpointApnsVoipSandboxChannel_ApplicationId(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsVoipSandboxChannel_DefaultAuthenticationMethod(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsVoipSandboxChannel_Enabled(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsVoipSandboxChannel_TokenKeyId(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsVoipSandboxChannel_BundleId(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsVoipSandboxChannel_Certificate(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsVoipSandboxChannel_PrivateKey(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsVoipSandboxChannel_TeamId(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsVoipSandboxChannel_TokenKey(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodePinpointApnsVoipSandboxChannel_ApplicationId(p PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	vals["application_id"] = cty.StringVal(p.ApplicationId)
}

func EncodePinpointApnsVoipSandboxChannel_DefaultAuthenticationMethod(p PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	vals["default_authentication_method"] = cty.StringVal(p.DefaultAuthenticationMethod)
}

func EncodePinpointApnsVoipSandboxChannel_Enabled(p PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodePinpointApnsVoipSandboxChannel_TokenKeyId(p PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	vals["token_key_id"] = cty.StringVal(p.TokenKeyId)
}

func EncodePinpointApnsVoipSandboxChannel_BundleId(p PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	vals["bundle_id"] = cty.StringVal(p.BundleId)
}

func EncodePinpointApnsVoipSandboxChannel_Certificate(p PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	vals["certificate"] = cty.StringVal(p.Certificate)
}

func EncodePinpointApnsVoipSandboxChannel_PrivateKey(p PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	vals["private_key"] = cty.StringVal(p.PrivateKey)
}

func EncodePinpointApnsVoipSandboxChannel_TeamId(p PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	vals["team_id"] = cty.StringVal(p.TeamId)
}

func EncodePinpointApnsVoipSandboxChannel_TokenKey(p PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	vals["token_key"] = cty.StringVal(p.TokenKey)
}