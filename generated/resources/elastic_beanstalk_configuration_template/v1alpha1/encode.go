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
	r, ok := mr.(*ElasticBeanstalkConfigurationTemplate)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ElasticBeanstalkConfigurationTemplate.")
	}
	return EncodeElasticBeanstalkConfigurationTemplate(*r), nil
}

func EncodeElasticBeanstalkConfigurationTemplate(r ElasticBeanstalkConfigurationTemplate) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeElasticBeanstalkConfigurationTemplate_Name(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkConfigurationTemplate_SolutionStackName(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkConfigurationTemplate_Application(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkConfigurationTemplate_Description(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkConfigurationTemplate_EnvironmentId(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkConfigurationTemplate_Id(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkConfigurationTemplate_Setting(r.Spec.ForProvider.Setting, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeElasticBeanstalkConfigurationTemplate_Name(p ElasticBeanstalkConfigurationTemplateParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeElasticBeanstalkConfigurationTemplate_SolutionStackName(p ElasticBeanstalkConfigurationTemplateParameters, vals map[string]cty.Value) {
	vals["solution_stack_name"] = cty.StringVal(p.SolutionStackName)
}

func EncodeElasticBeanstalkConfigurationTemplate_Application(p ElasticBeanstalkConfigurationTemplateParameters, vals map[string]cty.Value) {
	vals["application"] = cty.StringVal(p.Application)
}

func EncodeElasticBeanstalkConfigurationTemplate_Description(p ElasticBeanstalkConfigurationTemplateParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeElasticBeanstalkConfigurationTemplate_EnvironmentId(p ElasticBeanstalkConfigurationTemplateParameters, vals map[string]cty.Value) {
	vals["environment_id"] = cty.StringVal(p.EnvironmentId)
}

func EncodeElasticBeanstalkConfigurationTemplate_Id(p ElasticBeanstalkConfigurationTemplateParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeElasticBeanstalkConfigurationTemplate_Setting(p Setting, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeElasticBeanstalkConfigurationTemplate_Setting_Value(p, ctyVal)
	EncodeElasticBeanstalkConfigurationTemplate_Setting_Name(p, ctyVal)
	EncodeElasticBeanstalkConfigurationTemplate_Setting_Namespace(p, ctyVal)
	EncodeElasticBeanstalkConfigurationTemplate_Setting_Resource(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["setting"] = cty.SetVal(valsForCollection)
}

func EncodeElasticBeanstalkConfigurationTemplate_Setting_Value(p Setting, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeElasticBeanstalkConfigurationTemplate_Setting_Name(p Setting, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeElasticBeanstalkConfigurationTemplate_Setting_Namespace(p Setting, vals map[string]cty.Value) {
	vals["namespace"] = cty.StringVal(p.Namespace)
}

func EncodeElasticBeanstalkConfigurationTemplate_Setting_Resource(p Setting, vals map[string]cty.Value) {
	vals["resource"] = cty.StringVal(p.Resource)
}