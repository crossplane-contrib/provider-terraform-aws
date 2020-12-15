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
	r, ok := mr.(*PinpointSmsChannel)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a PinpointSmsChannel.")
	}
	return EncodePinpointSmsChannel(*r), nil
}

func EncodePinpointSmsChannel(r PinpointSmsChannel) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodePinpointSmsChannel_ApplicationId(r.Spec.ForProvider, ctyVal)
	EncodePinpointSmsChannel_Enabled(r.Spec.ForProvider, ctyVal)
	EncodePinpointSmsChannel_Id(r.Spec.ForProvider, ctyVal)
	EncodePinpointSmsChannel_SenderId(r.Spec.ForProvider, ctyVal)
	EncodePinpointSmsChannel_ShortCode(r.Spec.ForProvider, ctyVal)
	EncodePinpointSmsChannel_PromotionalMessagesPerSecond(r.Status.AtProvider, ctyVal)
	EncodePinpointSmsChannel_TransactionalMessagesPerSecond(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodePinpointSmsChannel_ApplicationId(p PinpointSmsChannelParameters, vals map[string]cty.Value) {
	vals["application_id"] = cty.StringVal(p.ApplicationId)
}

func EncodePinpointSmsChannel_Enabled(p PinpointSmsChannelParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodePinpointSmsChannel_Id(p PinpointSmsChannelParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodePinpointSmsChannel_SenderId(p PinpointSmsChannelParameters, vals map[string]cty.Value) {
	vals["sender_id"] = cty.StringVal(p.SenderId)
}

func EncodePinpointSmsChannel_ShortCode(p PinpointSmsChannelParameters, vals map[string]cty.Value) {
	vals["short_code"] = cty.StringVal(p.ShortCode)
}

func EncodePinpointSmsChannel_PromotionalMessagesPerSecond(p PinpointSmsChannelObservation, vals map[string]cty.Value) {
	vals["promotional_messages_per_second"] = cty.NumberIntVal(p.PromotionalMessagesPerSecond)
}

func EncodePinpointSmsChannel_TransactionalMessagesPerSecond(p PinpointSmsChannelObservation, vals map[string]cty.Value) {
	vals["transactional_messages_per_second"] = cty.NumberIntVal(p.TransactionalMessagesPerSecond)
}