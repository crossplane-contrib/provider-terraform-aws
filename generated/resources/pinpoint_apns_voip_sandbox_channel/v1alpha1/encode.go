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

package v1alpha1func EncodePinpointApnsVoipSandboxChannel(r PinpointApnsVoipSandboxChannel) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodePinpointApnsVoipSandboxChannel_TokenKey(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsVoipSandboxChannel_TokenKeyId(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsVoipSandboxChannel_ApplicationId(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsVoipSandboxChannel_Certificate(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsVoipSandboxChannel_DefaultAuthenticationMethod(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsVoipSandboxChannel_Id(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsVoipSandboxChannel_PrivateKey(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsVoipSandboxChannel_BundleId(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsVoipSandboxChannel_Enabled(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsVoipSandboxChannel_TeamId(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodePinpointApnsVoipSandboxChannel_TokenKey(p *PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	vals["token_key"] = cty.StringVal(p.TokenKey)
}

func EncodePinpointApnsVoipSandboxChannel_TokenKeyId(p *PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	vals["token_key_id"] = cty.StringVal(p.TokenKeyId)
}

func EncodePinpointApnsVoipSandboxChannel_ApplicationId(p *PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	vals["application_id"] = cty.StringVal(p.ApplicationId)
}

func EncodePinpointApnsVoipSandboxChannel_Certificate(p *PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	vals["certificate"] = cty.StringVal(p.Certificate)
}

func EncodePinpointApnsVoipSandboxChannel_DefaultAuthenticationMethod(p *PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	vals["default_authentication_method"] = cty.StringVal(p.DefaultAuthenticationMethod)
}

func EncodePinpointApnsVoipSandboxChannel_Id(p *PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodePinpointApnsVoipSandboxChannel_PrivateKey(p *PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	vals["private_key"] = cty.StringVal(p.PrivateKey)
}

func EncodePinpointApnsVoipSandboxChannel_BundleId(p *PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	vals["bundle_id"] = cty.StringVal(p.BundleId)
}

func EncodePinpointApnsVoipSandboxChannel_Enabled(p *PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodePinpointApnsVoipSandboxChannel_TeamId(p *PinpointApnsVoipSandboxChannelParameters, vals map[string]cty.Value) {
	vals["team_id"] = cty.StringVal(p.TeamId)
}