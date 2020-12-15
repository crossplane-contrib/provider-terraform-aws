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
	r, ok := mr.(*PinpointAdmChannel)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a PinpointAdmChannel.")
	}
	return EncodePinpointAdmChannel(*r), nil
}

func EncodePinpointAdmChannel(r PinpointAdmChannel) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodePinpointAdmChannel_ClientSecret(r.Spec.ForProvider, ctyVal)
	EncodePinpointAdmChannel_Enabled(r.Spec.ForProvider, ctyVal)
	EncodePinpointAdmChannel_ApplicationId(r.Spec.ForProvider, ctyVal)
	EncodePinpointAdmChannel_ClientId(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodePinpointAdmChannel_ClientSecret(p PinpointAdmChannelParameters, vals map[string]cty.Value) {
	vals["client_secret"] = cty.StringVal(p.ClientSecret)
}

func EncodePinpointAdmChannel_Enabled(p PinpointAdmChannelParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodePinpointAdmChannel_ApplicationId(p PinpointAdmChannelParameters, vals map[string]cty.Value) {
	vals["application_id"] = cty.StringVal(p.ApplicationId)
}

func EncodePinpointAdmChannel_ClientId(p PinpointAdmChannelParameters, vals map[string]cty.Value) {
	vals["client_id"] = cty.StringVal(p.ClientId)
}