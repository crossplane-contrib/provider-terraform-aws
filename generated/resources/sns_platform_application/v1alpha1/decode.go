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

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*SnsPlatformApplication)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeSnsPlatformApplication(r, ctyValue)
}

func DecodeSnsPlatformApplication(prev *SnsPlatformApplication, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeSnsPlatformApplication_EventDeliveryFailureTopicArn(&new.Spec.ForProvider, valMap)
	DecodeSnsPlatformApplication_EventEndpointDeletedTopicArn(&new.Spec.ForProvider, valMap)
	DecodeSnsPlatformApplication_EventEndpointUpdatedTopicArn(&new.Spec.ForProvider, valMap)
	DecodeSnsPlatformApplication_FailureFeedbackRoleArn(&new.Spec.ForProvider, valMap)
	DecodeSnsPlatformApplication_EventEndpointCreatedTopicArn(&new.Spec.ForProvider, valMap)
	DecodeSnsPlatformApplication_Id(&new.Spec.ForProvider, valMap)
	DecodeSnsPlatformApplication_Name(&new.Spec.ForProvider, valMap)
	DecodeSnsPlatformApplication_Platform(&new.Spec.ForProvider, valMap)
	DecodeSnsPlatformApplication_PlatformCredential(&new.Spec.ForProvider, valMap)
	DecodeSnsPlatformApplication_PlatformPrincipal(&new.Spec.ForProvider, valMap)
	DecodeSnsPlatformApplication_SuccessFeedbackRoleArn(&new.Spec.ForProvider, valMap)
	DecodeSnsPlatformApplication_SuccessFeedbackSampleRate(&new.Spec.ForProvider, valMap)
	DecodeSnsPlatformApplication_Arn(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeSnsPlatformApplication_EventDeliveryFailureTopicArn(p *SnsPlatformApplicationParameters, vals map[string]cty.Value) {
	p.EventDeliveryFailureTopicArn = ctwhy.ValueAsString(vals["event_delivery_failure_topic_arn"])
}

func DecodeSnsPlatformApplication_EventEndpointDeletedTopicArn(p *SnsPlatformApplicationParameters, vals map[string]cty.Value) {
	p.EventEndpointDeletedTopicArn = ctwhy.ValueAsString(vals["event_endpoint_deleted_topic_arn"])
}

func DecodeSnsPlatformApplication_EventEndpointUpdatedTopicArn(p *SnsPlatformApplicationParameters, vals map[string]cty.Value) {
	p.EventEndpointUpdatedTopicArn = ctwhy.ValueAsString(vals["event_endpoint_updated_topic_arn"])
}

func DecodeSnsPlatformApplication_FailureFeedbackRoleArn(p *SnsPlatformApplicationParameters, vals map[string]cty.Value) {
	p.FailureFeedbackRoleArn = ctwhy.ValueAsString(vals["failure_feedback_role_arn"])
}

func DecodeSnsPlatformApplication_EventEndpointCreatedTopicArn(p *SnsPlatformApplicationParameters, vals map[string]cty.Value) {
	p.EventEndpointCreatedTopicArn = ctwhy.ValueAsString(vals["event_endpoint_created_topic_arn"])
}

func DecodeSnsPlatformApplication_Id(p *SnsPlatformApplicationParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeSnsPlatformApplication_Name(p *SnsPlatformApplicationParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

func DecodeSnsPlatformApplication_Platform(p *SnsPlatformApplicationParameters, vals map[string]cty.Value) {
	p.Platform = ctwhy.ValueAsString(vals["platform"])
}

func DecodeSnsPlatformApplication_PlatformCredential(p *SnsPlatformApplicationParameters, vals map[string]cty.Value) {
	p.PlatformCredential = ctwhy.ValueAsString(vals["platform_credential"])
}

func DecodeSnsPlatformApplication_PlatformPrincipal(p *SnsPlatformApplicationParameters, vals map[string]cty.Value) {
	p.PlatformPrincipal = ctwhy.ValueAsString(vals["platform_principal"])
}

func DecodeSnsPlatformApplication_SuccessFeedbackRoleArn(p *SnsPlatformApplicationParameters, vals map[string]cty.Value) {
	p.SuccessFeedbackRoleArn = ctwhy.ValueAsString(vals["success_feedback_role_arn"])
}

func DecodeSnsPlatformApplication_SuccessFeedbackSampleRate(p *SnsPlatformApplicationParameters, vals map[string]cty.Value) {
	p.SuccessFeedbackSampleRate = ctwhy.ValueAsString(vals["success_feedback_sample_rate"])
}

func DecodeSnsPlatformApplication_Arn(p *SnsPlatformApplicationObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}