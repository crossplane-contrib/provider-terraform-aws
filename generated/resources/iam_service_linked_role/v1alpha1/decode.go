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
	r, ok := mr.(*IamServiceLinkedRole)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeIamServiceLinkedRole(r, ctyValue)
}

func DecodeIamServiceLinkedRole(prev *IamServiceLinkedRole, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeIamServiceLinkedRole_CustomSuffix(&new.Spec.ForProvider, valMap)
	DecodeIamServiceLinkedRole_AwsServiceName(&new.Spec.ForProvider, valMap)
	DecodeIamServiceLinkedRole_Description(&new.Spec.ForProvider, valMap)
	DecodeIamServiceLinkedRole_Arn(&new.Status.AtProvider, valMap)
	DecodeIamServiceLinkedRole_CreateDate(&new.Status.AtProvider, valMap)
	DecodeIamServiceLinkedRole_Name(&new.Status.AtProvider, valMap)
	DecodeIamServiceLinkedRole_Path(&new.Status.AtProvider, valMap)
	DecodeIamServiceLinkedRole_UniqueId(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeIamServiceLinkedRole_CustomSuffix(p *IamServiceLinkedRoleParameters, vals map[string]cty.Value) {
	p.CustomSuffix = ctwhy.ValueAsString(vals["custom_suffix"])
}

//primitiveTypeDecodeTemplate
func DecodeIamServiceLinkedRole_AwsServiceName(p *IamServiceLinkedRoleParameters, vals map[string]cty.Value) {
	p.AwsServiceName = ctwhy.ValueAsString(vals["aws_service_name"])
}

//primitiveTypeDecodeTemplate
func DecodeIamServiceLinkedRole_Description(p *IamServiceLinkedRoleParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeIamServiceLinkedRole_Arn(p *IamServiceLinkedRoleObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeIamServiceLinkedRole_CreateDate(p *IamServiceLinkedRoleObservation, vals map[string]cty.Value) {
	p.CreateDate = ctwhy.ValueAsString(vals["create_date"])
}

//primitiveTypeDecodeTemplate
func DecodeIamServiceLinkedRole_Name(p *IamServiceLinkedRoleObservation, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeIamServiceLinkedRole_Path(p *IamServiceLinkedRoleObservation, vals map[string]cty.Value) {
	p.Path = ctwhy.ValueAsString(vals["path"])
}

//primitiveTypeDecodeTemplate
func DecodeIamServiceLinkedRole_UniqueId(p *IamServiceLinkedRoleObservation, vals map[string]cty.Value) {
	p.UniqueId = ctwhy.ValueAsString(vals["unique_id"])
}