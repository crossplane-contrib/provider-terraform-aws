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
	r, ok := mr.(*IamRole)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeIamRole(r, ctyValue)
}

func DecodeIamRole(prev *IamRole, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeIamRole_ForceDetachPolicies(&new.Spec.ForProvider, valMap)
	DecodeIamRole_MaxSessionDuration(&new.Spec.ForProvider, valMap)
	DecodeIamRole_Name(&new.Spec.ForProvider, valMap)
	DecodeIamRole_NamePrefix(&new.Spec.ForProvider, valMap)
	DecodeIamRole_Path(&new.Spec.ForProvider, valMap)
	DecodeIamRole_AssumeRolePolicy(&new.Spec.ForProvider, valMap)
	DecodeIamRole_Description(&new.Spec.ForProvider, valMap)
	DecodeIamRole_PermissionsBoundary(&new.Spec.ForProvider, valMap)
	DecodeIamRole_Tags(&new.Spec.ForProvider, valMap)
	DecodeIamRole_Arn(&new.Status.AtProvider, valMap)
	DecodeIamRole_CreateDate(&new.Status.AtProvider, valMap)
	DecodeIamRole_UniqueId(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeIamRole_ForceDetachPolicies(p *IamRoleParameters, vals map[string]cty.Value) {
	p.ForceDetachPolicies = ctwhy.ValueAsBool(vals["force_detach_policies"])
}

//primitiveTypeDecodeTemplate
func DecodeIamRole_MaxSessionDuration(p *IamRoleParameters, vals map[string]cty.Value) {
	p.MaxSessionDuration = ctwhy.ValueAsInt64(vals["max_session_duration"])
}

//primitiveTypeDecodeTemplate
func DecodeIamRole_Name(p *IamRoleParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeIamRole_NamePrefix(p *IamRoleParameters, vals map[string]cty.Value) {
	p.NamePrefix = ctwhy.ValueAsString(vals["name_prefix"])
}

//primitiveTypeDecodeTemplate
func DecodeIamRole_Path(p *IamRoleParameters, vals map[string]cty.Value) {
	p.Path = ctwhy.ValueAsString(vals["path"])
}

//primitiveTypeDecodeTemplate
func DecodeIamRole_AssumeRolePolicy(p *IamRoleParameters, vals map[string]cty.Value) {
	p.AssumeRolePolicy = ctwhy.ValueAsString(vals["assume_role_policy"])
}

//primitiveTypeDecodeTemplate
func DecodeIamRole_Description(p *IamRoleParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeIamRole_PermissionsBoundary(p *IamRoleParameters, vals map[string]cty.Value) {
	p.PermissionsBoundary = ctwhy.ValueAsString(vals["permissions_boundary"])
}

//primitiveMapTypeDecodeTemplate
func DecodeIamRole_Tags(p *IamRoleParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeIamRole_Arn(p *IamRoleObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeIamRole_CreateDate(p *IamRoleObservation, vals map[string]cty.Value) {
	p.CreateDate = ctwhy.ValueAsString(vals["create_date"])
}

//primitiveTypeDecodeTemplate
func DecodeIamRole_UniqueId(p *IamRoleObservation, vals map[string]cty.Value) {
	p.UniqueId = ctwhy.ValueAsString(vals["unique_id"])
}