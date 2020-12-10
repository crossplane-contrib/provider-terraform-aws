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
	r, ok := mr.(*CloudformationStack)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a CloudformationStack.")
	}
	return EncodeCloudformationStack(*r), nil
}

func EncodeCloudformationStack(r CloudformationStack) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCloudformationStack_TemplateUrl(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStack_Capabilities(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStack_DisableRollback(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStack_Id(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStack_PolicyBody(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStack_TemplateBody(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStack_NotificationArns(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStack_Parameters(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStack_PolicyUrl(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStack_OnFailure(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStack_Tags(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStack_IamRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStack_Name(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStack_TimeoutInMinutes(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStack_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeCloudformationStack_Outputs(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeCloudformationStack_TemplateUrl(p CloudformationStackParameters, vals map[string]cty.Value) {
	vals["template_url"] = cty.StringVal(p.TemplateUrl)
}

func EncodeCloudformationStack_Capabilities(p CloudformationStackParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Capabilities {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["capabilities"] = cty.SetVal(colVals)
}

func EncodeCloudformationStack_DisableRollback(p CloudformationStackParameters, vals map[string]cty.Value) {
	vals["disable_rollback"] = cty.BoolVal(p.DisableRollback)
}

func EncodeCloudformationStack_Id(p CloudformationStackParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCloudformationStack_PolicyBody(p CloudformationStackParameters, vals map[string]cty.Value) {
	vals["policy_body"] = cty.StringVal(p.PolicyBody)
}

func EncodeCloudformationStack_TemplateBody(p CloudformationStackParameters, vals map[string]cty.Value) {
	vals["template_body"] = cty.StringVal(p.TemplateBody)
}

func EncodeCloudformationStack_NotificationArns(p CloudformationStackParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.NotificationArns {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["notification_arns"] = cty.SetVal(colVals)
}

func EncodeCloudformationStack_Parameters(p CloudformationStackParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Parameters {
		mVals[key] = cty.StringVal(value)
	}
	vals["parameters"] = cty.MapVal(mVals)
}

func EncodeCloudformationStack_PolicyUrl(p CloudformationStackParameters, vals map[string]cty.Value) {
	vals["policy_url"] = cty.StringVal(p.PolicyUrl)
}

func EncodeCloudformationStack_OnFailure(p CloudformationStackParameters, vals map[string]cty.Value) {
	vals["on_failure"] = cty.StringVal(p.OnFailure)
}

func EncodeCloudformationStack_Tags(p CloudformationStackParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeCloudformationStack_IamRoleArn(p CloudformationStackParameters, vals map[string]cty.Value) {
	vals["iam_role_arn"] = cty.StringVal(p.IamRoleArn)
}

func EncodeCloudformationStack_Name(p CloudformationStackParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCloudformationStack_TimeoutInMinutes(p CloudformationStackParameters, vals map[string]cty.Value) {
	vals["timeout_in_minutes"] = cty.NumberIntVal(p.TimeoutInMinutes)
}

func EncodeCloudformationStack_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeCloudformationStack_Timeouts_Create(p, ctyVal)
	EncodeCloudformationStack_Timeouts_Delete(p, ctyVal)
	EncodeCloudformationStack_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeCloudformationStack_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeCloudformationStack_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeCloudformationStack_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeCloudformationStack_Outputs(p CloudformationStackObservation, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Outputs {
		mVals[key] = cty.StringVal(value)
	}
	vals["outputs"] = cty.MapVal(mVals)
}