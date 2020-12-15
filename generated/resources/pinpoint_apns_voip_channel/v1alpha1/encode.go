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
	r, ok := mr.(*PinpointApnsVoipChannel)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a PinpointApnsVoipChannel.")
	}
	return EncodePinpointApnsVoipChannel(*r), nil
}

func EncodePinpointApnsVoipChannel(r PinpointApnsVoipChannel) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodePinpointApnsVoipChannel_TokenKeyId(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsVoipChannel_ApplicationId(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsVoipChannel_BundleId(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsVoipChannel_Certificate(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsVoipChannel_DefaultAuthenticationMethod(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsVoipChannel_Enabled(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsVoipChannel_PrivateKey(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsVoipChannel_TeamId(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsVoipChannel_Id(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsVoipChannel_TokenKey(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodePinpointApnsVoipChannel_TokenKeyId(p PinpointApnsVoipChannelParameters, vals map[string]cty.Value) {
	vals["token_key_id"] = cty.StringVal(p.TokenKeyId)
}

func EncodePinpointApnsVoipChannel_ApplicationId(p PinpointApnsVoipChannelParameters, vals map[string]cty.Value) {
	vals["application_id"] = cty.StringVal(p.ApplicationId)
}

func EncodePinpointApnsVoipChannel_BundleId(p PinpointApnsVoipChannelParameters, vals map[string]cty.Value) {
	vals["bundle_id"] = cty.StringVal(p.BundleId)
}

func EncodePinpointApnsVoipChannel_Certificate(p PinpointApnsVoipChannelParameters, vals map[string]cty.Value) {
	vals["certificate"] = cty.StringVal(p.Certificate)
}

func EncodePinpointApnsVoipChannel_DefaultAuthenticationMethod(p PinpointApnsVoipChannelParameters, vals map[string]cty.Value) {
	vals["default_authentication_method"] = cty.StringVal(p.DefaultAuthenticationMethod)
}

func EncodePinpointApnsVoipChannel_Enabled(p PinpointApnsVoipChannelParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodePinpointApnsVoipChannel_PrivateKey(p PinpointApnsVoipChannelParameters, vals map[string]cty.Value) {
	vals["private_key"] = cty.StringVal(p.PrivateKey)
}

func EncodePinpointApnsVoipChannel_TeamId(p PinpointApnsVoipChannelParameters, vals map[string]cty.Value) {
	vals["team_id"] = cty.StringVal(p.TeamId)
}

func EncodePinpointApnsVoipChannel_Id(p PinpointApnsVoipChannelParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodePinpointApnsVoipChannel_TokenKey(p PinpointApnsVoipChannelParameters, vals map[string]cty.Value) {
	vals["token_key"] = cty.StringVal(p.TokenKey)
}