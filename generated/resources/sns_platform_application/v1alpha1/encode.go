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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*SnsPlatformApplication)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a SnsPlatformApplication.")
	}
	return EncodeSnsPlatformApplication(*r), nil
}

func EncodeSnsPlatformApplication(r SnsPlatformApplication) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSnsPlatformApplication_Name(r.Spec.ForProvider, ctyVal)
	EncodeSnsPlatformApplication_Platform(r.Spec.ForProvider, ctyVal)
	EncodeSnsPlatformApplication_PlatformCredential(r.Spec.ForProvider, ctyVal)
	EncodeSnsPlatformApplication_PlatformPrincipal(r.Spec.ForProvider, ctyVal)
	EncodeSnsPlatformApplication_EventDeliveryFailureTopicArn(r.Spec.ForProvider, ctyVal)
	EncodeSnsPlatformApplication_EventEndpointCreatedTopicArn(r.Spec.ForProvider, ctyVal)
	EncodeSnsPlatformApplication_EventEndpointUpdatedTopicArn(r.Spec.ForProvider, ctyVal)
	EncodeSnsPlatformApplication_SuccessFeedbackRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeSnsPlatformApplication_SuccessFeedbackSampleRate(r.Spec.ForProvider, ctyVal)
	EncodeSnsPlatformApplication_EventEndpointDeletedTopicArn(r.Spec.ForProvider, ctyVal)
	EncodeSnsPlatformApplication_FailureFeedbackRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeSnsPlatformApplication_Id(r.Spec.ForProvider, ctyVal)
	EncodeSnsPlatformApplication_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeSnsPlatformApplication_Name(p SnsPlatformApplicationParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSnsPlatformApplication_Platform(p SnsPlatformApplicationParameters, vals map[string]cty.Value) {
	vals["platform"] = cty.StringVal(p.Platform)
}

func EncodeSnsPlatformApplication_PlatformCredential(p SnsPlatformApplicationParameters, vals map[string]cty.Value) {
	vals["platform_credential"] = cty.StringVal(p.PlatformCredential)
}

func EncodeSnsPlatformApplication_PlatformPrincipal(p SnsPlatformApplicationParameters, vals map[string]cty.Value) {
	vals["platform_principal"] = cty.StringVal(p.PlatformPrincipal)
}

func EncodeSnsPlatformApplication_EventDeliveryFailureTopicArn(p SnsPlatformApplicationParameters, vals map[string]cty.Value) {
	vals["event_delivery_failure_topic_arn"] = cty.StringVal(p.EventDeliveryFailureTopicArn)
}

func EncodeSnsPlatformApplication_EventEndpointCreatedTopicArn(p SnsPlatformApplicationParameters, vals map[string]cty.Value) {
	vals["event_endpoint_created_topic_arn"] = cty.StringVal(p.EventEndpointCreatedTopicArn)
}

func EncodeSnsPlatformApplication_EventEndpointUpdatedTopicArn(p SnsPlatformApplicationParameters, vals map[string]cty.Value) {
	vals["event_endpoint_updated_topic_arn"] = cty.StringVal(p.EventEndpointUpdatedTopicArn)
}

func EncodeSnsPlatformApplication_SuccessFeedbackRoleArn(p SnsPlatformApplicationParameters, vals map[string]cty.Value) {
	vals["success_feedback_role_arn"] = cty.StringVal(p.SuccessFeedbackRoleArn)
}

func EncodeSnsPlatformApplication_SuccessFeedbackSampleRate(p SnsPlatformApplicationParameters, vals map[string]cty.Value) {
	vals["success_feedback_sample_rate"] = cty.StringVal(p.SuccessFeedbackSampleRate)
}

func EncodeSnsPlatformApplication_EventEndpointDeletedTopicArn(p SnsPlatformApplicationParameters, vals map[string]cty.Value) {
	vals["event_endpoint_deleted_topic_arn"] = cty.StringVal(p.EventEndpointDeletedTopicArn)
}

func EncodeSnsPlatformApplication_FailureFeedbackRoleArn(p SnsPlatformApplicationParameters, vals map[string]cty.Value) {
	vals["failure_feedback_role_arn"] = cty.StringVal(p.FailureFeedbackRoleArn)
}

func EncodeSnsPlatformApplication_Id(p SnsPlatformApplicationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSnsPlatformApplication_Arn(p SnsPlatformApplicationObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}