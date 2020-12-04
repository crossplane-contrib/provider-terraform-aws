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

func EncodePinpointSmsChannel(r PinpointSmsChannel) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodePinpointSmsChannel_ShortCode(r.Spec.ForProvider, ctyVal)
	EncodePinpointSmsChannel_ApplicationId(r.Spec.ForProvider, ctyVal)
	EncodePinpointSmsChannel_Enabled(r.Spec.ForProvider, ctyVal)
	EncodePinpointSmsChannel_Id(r.Spec.ForProvider, ctyVal)
	EncodePinpointSmsChannel_SenderId(r.Spec.ForProvider, ctyVal)
	EncodePinpointSmsChannel_TransactionalMessagesPerSecond(r.Status.AtProvider, ctyVal)
	EncodePinpointSmsChannel_PromotionalMessagesPerSecond(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodePinpointSmsChannel_ShortCode(p PinpointSmsChannelParameters, vals map[string]cty.Value) {
	vals["short_code"] = cty.StringVal(p.ShortCode)
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

func EncodePinpointSmsChannel_TransactionalMessagesPerSecond(p PinpointSmsChannelObservation, vals map[string]cty.Value) {
	vals["transactional_messages_per_second"] = cty.NumberIntVal(p.TransactionalMessagesPerSecond)
}

func EncodePinpointSmsChannel_PromotionalMessagesPerSecond(p PinpointSmsChannelObservation, vals map[string]cty.Value) {
	vals["promotional_messages_per_second"] = cty.NumberIntVal(p.PromotionalMessagesPerSecond)
}