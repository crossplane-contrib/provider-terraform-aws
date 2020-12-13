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
	r, ok := mr.(*CloudformationStackSet)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a CloudformationStackSet.")
	}
	return EncodeCloudformationStackSet(*r), nil
}

func EncodeCloudformationStackSet(r CloudformationStackSet) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCloudformationStackSet_Description(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStackSet_ExecutionRoleName(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStackSet_Id(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStackSet_Tags(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStackSet_AdministrationRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStackSet_Capabilities(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStackSet_TemplateUrl(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStackSet_Name(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStackSet_Parameters(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStackSet_TemplateBody(r.Spec.ForProvider, ctyVal)
	EncodeCloudformationStackSet_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeCloudformationStackSet_StackSetId(r.Status.AtProvider, ctyVal)
	EncodeCloudformationStackSet_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeCloudformationStackSet_Description(p CloudformationStackSetParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeCloudformationStackSet_ExecutionRoleName(p CloudformationStackSetParameters, vals map[string]cty.Value) {
	vals["execution_role_name"] = cty.StringVal(p.ExecutionRoleName)
}

func EncodeCloudformationStackSet_Id(p CloudformationStackSetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCloudformationStackSet_Tags(p CloudformationStackSetParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeCloudformationStackSet_AdministrationRoleArn(p CloudformationStackSetParameters, vals map[string]cty.Value) {
	vals["administration_role_arn"] = cty.StringVal(p.AdministrationRoleArn)
}

func EncodeCloudformationStackSet_Capabilities(p CloudformationStackSetParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Capabilities {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["capabilities"] = cty.SetVal(colVals)
}

func EncodeCloudformationStackSet_TemplateUrl(p CloudformationStackSetParameters, vals map[string]cty.Value) {
	vals["template_url"] = cty.StringVal(p.TemplateUrl)
}

func EncodeCloudformationStackSet_Name(p CloudformationStackSetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCloudformationStackSet_Parameters(p CloudformationStackSetParameters, vals map[string]cty.Value) {
	if len(p.Parameters) == 0 {
		vals["parameters"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Parameters {
		mVals[key] = cty.StringVal(value)
	}
	vals["parameters"] = cty.MapVal(mVals)
}

func EncodeCloudformationStackSet_TemplateBody(p CloudformationStackSetParameters, vals map[string]cty.Value) {
	vals["template_body"] = cty.StringVal(p.TemplateBody)
}

func EncodeCloudformationStackSet_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeCloudformationStackSet_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeCloudformationStackSet_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeCloudformationStackSet_StackSetId(p CloudformationStackSetObservation, vals map[string]cty.Value) {
	vals["stack_set_id"] = cty.StringVal(p.StackSetId)
}

func EncodeCloudformationStackSet_Arn(p CloudformationStackSetObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}