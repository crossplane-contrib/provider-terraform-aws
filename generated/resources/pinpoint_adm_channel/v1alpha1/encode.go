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
	"github.com/zclconf/go-cty/cty"
)

func EncodePinpointAdmChannel(r PinpointAdmChannel) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodePinpointAdmChannel_ApplicationId(r.Spec.ForProvider, ctyVal)
	EncodePinpointAdmChannel_ClientId(r.Spec.ForProvider, ctyVal)
	EncodePinpointAdmChannel_ClientSecret(r.Spec.ForProvider, ctyVal)
	EncodePinpointAdmChannel_Enabled(r.Spec.ForProvider, ctyVal)
	EncodePinpointAdmChannel_Id(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodePinpointAdmChannel_ApplicationId(p PinpointAdmChannelParameters, vals map[string]cty.Value) {
	vals["application_id"] = cty.StringVal(p.ApplicationId)
}

func EncodePinpointAdmChannel_ClientId(p PinpointAdmChannelParameters, vals map[string]cty.Value) {
	vals["client_id"] = cty.StringVal(p.ClientId)
}

func EncodePinpointAdmChannel_ClientSecret(p PinpointAdmChannelParameters, vals map[string]cty.Value) {
	vals["client_secret"] = cty.StringVal(p.ClientSecret)
}

func EncodePinpointAdmChannel_Enabled(p PinpointAdmChannelParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodePinpointAdmChannel_Id(p PinpointAdmChannelParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}