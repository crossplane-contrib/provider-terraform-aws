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

package v1alpha1func EncodePinpointApnsChannel(r PinpointApnsChannel) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodePinpointApnsChannel_ApplicationId(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsChannel_Enabled(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsChannel_TokenKey(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsChannel_TokenKeyId(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsChannel_BundleId(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsChannel_Certificate(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsChannel_DefaultAuthenticationMethod(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsChannel_Id(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsChannel_PrivateKey(r.Spec.ForProvider, ctyVal)
	EncodePinpointApnsChannel_TeamId(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodePinpointApnsChannel_ApplicationId(p *PinpointApnsChannelParameters, vals map[string]cty.Value) {
	vals["application_id"] = cty.StringVal(p.ApplicationId)
}

func EncodePinpointApnsChannel_Enabled(p *PinpointApnsChannelParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodePinpointApnsChannel_TokenKey(p *PinpointApnsChannelParameters, vals map[string]cty.Value) {
	vals["token_key"] = cty.StringVal(p.TokenKey)
}

func EncodePinpointApnsChannel_TokenKeyId(p *PinpointApnsChannelParameters, vals map[string]cty.Value) {
	vals["token_key_id"] = cty.StringVal(p.TokenKeyId)
}

func EncodePinpointApnsChannel_BundleId(p *PinpointApnsChannelParameters, vals map[string]cty.Value) {
	vals["bundle_id"] = cty.StringVal(p.BundleId)
}

func EncodePinpointApnsChannel_Certificate(p *PinpointApnsChannelParameters, vals map[string]cty.Value) {
	vals["certificate"] = cty.StringVal(p.Certificate)
}

func EncodePinpointApnsChannel_DefaultAuthenticationMethod(p *PinpointApnsChannelParameters, vals map[string]cty.Value) {
	vals["default_authentication_method"] = cty.StringVal(p.DefaultAuthenticationMethod)
}

func EncodePinpointApnsChannel_Id(p *PinpointApnsChannelParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodePinpointApnsChannel_PrivateKey(p *PinpointApnsChannelParameters, vals map[string]cty.Value) {
	vals["private_key"] = cty.StringVal(p.PrivateKey)
}

func EncodePinpointApnsChannel_TeamId(p *PinpointApnsChannelParameters, vals map[string]cty.Value) {
	vals["team_id"] = cty.StringVal(p.TeamId)
}