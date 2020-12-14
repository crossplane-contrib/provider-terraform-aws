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
	r, ok := mr.(*CloudformationStack)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeCloudformationStack(r, ctyValue)
}

func DecodeCloudformationStack(prev *CloudformationStack, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeCloudformationStack_Id(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStack_Capabilities(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStack_DisableRollback(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStack_IamRoleArn(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStack_NotificationArns(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStack_PolicyBody(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStack_PolicyUrl(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStack_Tags(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStack_TimeoutInMinutes(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStack_Name(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStack_OnFailure(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStack_Parameters(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStack_TemplateBody(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStack_TemplateUrl(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStack_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeCloudformationStack_Outputs(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeCloudformationStack_Id(p *CloudformationStackParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeCloudformationStack_Capabilities(p *CloudformationStackParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["capabilities"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.Capabilities = goVals
}

func DecodeCloudformationStack_DisableRollback(p *CloudformationStackParameters, vals map[string]cty.Value) {
	p.DisableRollback = ctwhy.ValueAsBool(vals["disable_rollback"])
}

func DecodeCloudformationStack_IamRoleArn(p *CloudformationStackParameters, vals map[string]cty.Value) {
	p.IamRoleArn = ctwhy.ValueAsString(vals["iam_role_arn"])
}

func DecodeCloudformationStack_NotificationArns(p *CloudformationStackParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["notification_arns"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.NotificationArns = goVals
}

func DecodeCloudformationStack_PolicyBody(p *CloudformationStackParameters, vals map[string]cty.Value) {
	p.PolicyBody = ctwhy.ValueAsString(vals["policy_body"])
}

func DecodeCloudformationStack_PolicyUrl(p *CloudformationStackParameters, vals map[string]cty.Value) {
	p.PolicyUrl = ctwhy.ValueAsString(vals["policy_url"])
}

func DecodeCloudformationStack_Tags(p *CloudformationStackParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeCloudformationStack_TimeoutInMinutes(p *CloudformationStackParameters, vals map[string]cty.Value) {
	p.TimeoutInMinutes = ctwhy.ValueAsInt64(vals["timeout_in_minutes"])
}

func DecodeCloudformationStack_Name(p *CloudformationStackParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

func DecodeCloudformationStack_OnFailure(p *CloudformationStackParameters, vals map[string]cty.Value) {
	p.OnFailure = ctwhy.ValueAsString(vals["on_failure"])
}

func DecodeCloudformationStack_Parameters(p *CloudformationStackParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["parameters"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Parameters = vMap
}

func DecodeCloudformationStack_TemplateBody(p *CloudformationStackParameters, vals map[string]cty.Value) {
	p.TemplateBody = ctwhy.ValueAsString(vals["template_body"])
}

func DecodeCloudformationStack_TemplateUrl(p *CloudformationStackParameters, vals map[string]cty.Value) {
	p.TemplateUrl = ctwhy.ValueAsString(vals["template_url"])
}

func DecodeCloudformationStack_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeCloudformationStack_Timeouts_Update(p, valMap)
	DecodeCloudformationStack_Timeouts_Create(p, valMap)
	DecodeCloudformationStack_Timeouts_Delete(p, valMap)
}

func DecodeCloudformationStack_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}

func DecodeCloudformationStack_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

func DecodeCloudformationStack_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

func DecodeCloudformationStack_Outputs(p *CloudformationStackObservation, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["outputs"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Outputs = vMap
}