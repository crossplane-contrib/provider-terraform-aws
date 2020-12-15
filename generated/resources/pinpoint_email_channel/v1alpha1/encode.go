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
	r, ok := mr.(*PinpointEmailChannel)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a PinpointEmailChannel.")
	}
	return EncodePinpointEmailChannel(*r), nil
}

func EncodePinpointEmailChannel(r PinpointEmailChannel) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodePinpointEmailChannel_ApplicationId(r.Spec.ForProvider, ctyVal)
	EncodePinpointEmailChannel_Enabled(r.Spec.ForProvider, ctyVal)
	EncodePinpointEmailChannel_FromAddress(r.Spec.ForProvider, ctyVal)
	EncodePinpointEmailChannel_Identity(r.Spec.ForProvider, ctyVal)
	EncodePinpointEmailChannel_RoleArn(r.Spec.ForProvider, ctyVal)
	EncodePinpointEmailChannel_MessagesPerSecond(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodePinpointEmailChannel_ApplicationId(p PinpointEmailChannelParameters, vals map[string]cty.Value) {
	vals["application_id"] = cty.StringVal(p.ApplicationId)
}

func EncodePinpointEmailChannel_Enabled(p PinpointEmailChannelParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodePinpointEmailChannel_FromAddress(p PinpointEmailChannelParameters, vals map[string]cty.Value) {
	vals["from_address"] = cty.StringVal(p.FromAddress)
}

func EncodePinpointEmailChannel_Identity(p PinpointEmailChannelParameters, vals map[string]cty.Value) {
	vals["identity"] = cty.StringVal(p.Identity)
}

func EncodePinpointEmailChannel_RoleArn(p PinpointEmailChannelParameters, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodePinpointEmailChannel_MessagesPerSecond(p PinpointEmailChannelObservation, vals map[string]cty.Value) {
	vals["messages_per_second"] = cty.NumberIntVal(p.MessagesPerSecond)
}