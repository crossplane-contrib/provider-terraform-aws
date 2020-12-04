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

func EncodeSnsTopicSubscription(r SnsTopicSubscription) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSnsTopicSubscription_DeliveryPolicy(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopicSubscription_Endpoint(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopicSubscription_EndpointAutoConfirms(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopicSubscription_Id(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopicSubscription_Protocol(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopicSubscription_RawMessageDelivery(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopicSubscription_ConfirmationTimeoutInMinutes(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopicSubscription_FilterPolicy(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopicSubscription_TopicArn(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopicSubscription_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeSnsTopicSubscription_DeliveryPolicy(p SnsTopicSubscriptionParameters, vals map[string]cty.Value) {
	vals["delivery_policy"] = cty.StringVal(p.DeliveryPolicy)
}

func EncodeSnsTopicSubscription_Endpoint(p SnsTopicSubscriptionParameters, vals map[string]cty.Value) {
	vals["endpoint"] = cty.StringVal(p.Endpoint)
}

func EncodeSnsTopicSubscription_EndpointAutoConfirms(p SnsTopicSubscriptionParameters, vals map[string]cty.Value) {
	vals["endpoint_auto_confirms"] = cty.BoolVal(p.EndpointAutoConfirms)
}

func EncodeSnsTopicSubscription_Id(p SnsTopicSubscriptionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSnsTopicSubscription_Protocol(p SnsTopicSubscriptionParameters, vals map[string]cty.Value) {
	vals["protocol"] = cty.StringVal(p.Protocol)
}

func EncodeSnsTopicSubscription_RawMessageDelivery(p SnsTopicSubscriptionParameters, vals map[string]cty.Value) {
	vals["raw_message_delivery"] = cty.BoolVal(p.RawMessageDelivery)
}

func EncodeSnsTopicSubscription_ConfirmationTimeoutInMinutes(p SnsTopicSubscriptionParameters, vals map[string]cty.Value) {
	vals["confirmation_timeout_in_minutes"] = cty.NumberIntVal(p.ConfirmationTimeoutInMinutes)
}

func EncodeSnsTopicSubscription_FilterPolicy(p SnsTopicSubscriptionParameters, vals map[string]cty.Value) {
	vals["filter_policy"] = cty.StringVal(p.FilterPolicy)
}

func EncodeSnsTopicSubscription_TopicArn(p SnsTopicSubscriptionParameters, vals map[string]cty.Value) {
	vals["topic_arn"] = cty.StringVal(p.TopicArn)
}

func EncodeSnsTopicSubscription_Arn(p SnsTopicSubscriptionObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}