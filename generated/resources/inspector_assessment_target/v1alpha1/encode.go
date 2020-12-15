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
	r, ok := mr.(*InspectorAssessmentTarget)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a InspectorAssessmentTarget.")
	}
	return EncodeInspectorAssessmentTarget(*r), nil
}

func EncodeInspectorAssessmentTarget(r InspectorAssessmentTarget) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeInspectorAssessmentTarget_Name(r.Spec.ForProvider, ctyVal)
	EncodeInspectorAssessmentTarget_ResourceGroupArn(r.Spec.ForProvider, ctyVal)
	EncodeInspectorAssessmentTarget_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeInspectorAssessmentTarget_Name(p InspectorAssessmentTargetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeInspectorAssessmentTarget_ResourceGroupArn(p InspectorAssessmentTargetParameters, vals map[string]cty.Value) {
	vals["resource_group_arn"] = cty.StringVal(p.ResourceGroupArn)
}

func EncodeInspectorAssessmentTarget_Arn(p InspectorAssessmentTargetObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}