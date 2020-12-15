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
	r, ok := mr.(*IamInstanceProfile)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeIamInstanceProfile(r, ctyValue)
}

func DecodeIamInstanceProfile(prev *IamInstanceProfile, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeIamInstanceProfile_Path(&new.Spec.ForProvider, valMap)
	DecodeIamInstanceProfile_Role(&new.Spec.ForProvider, valMap)
	DecodeIamInstanceProfile_Id(&new.Spec.ForProvider, valMap)
	DecodeIamInstanceProfile_Name(&new.Spec.ForProvider, valMap)
	DecodeIamInstanceProfile_NamePrefix(&new.Spec.ForProvider, valMap)
	DecodeIamInstanceProfile_UniqueId(&new.Status.AtProvider, valMap)
	DecodeIamInstanceProfile_Arn(&new.Status.AtProvider, valMap)
	DecodeIamInstanceProfile_CreateDate(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeIamInstanceProfile_Path(p *IamInstanceProfileParameters, vals map[string]cty.Value) {
	p.Path = ctwhy.ValueAsString(vals["path"])
}

//primitiveTypeDecodeTemplate
func DecodeIamInstanceProfile_Role(p *IamInstanceProfileParameters, vals map[string]cty.Value) {
	p.Role = ctwhy.ValueAsString(vals["role"])
}

//primitiveTypeDecodeTemplate
func DecodeIamInstanceProfile_Id(p *IamInstanceProfileParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeIamInstanceProfile_Name(p *IamInstanceProfileParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeIamInstanceProfile_NamePrefix(p *IamInstanceProfileParameters, vals map[string]cty.Value) {
	p.NamePrefix = ctwhy.ValueAsString(vals["name_prefix"])
}

//primitiveTypeDecodeTemplate
func DecodeIamInstanceProfile_UniqueId(p *IamInstanceProfileObservation, vals map[string]cty.Value) {
	p.UniqueId = ctwhy.ValueAsString(vals["unique_id"])
}

//primitiveTypeDecodeTemplate
func DecodeIamInstanceProfile_Arn(p *IamInstanceProfileObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeIamInstanceProfile_CreateDate(p *IamInstanceProfileObservation, vals map[string]cty.Value) {
	p.CreateDate = ctwhy.ValueAsString(vals["create_date"])
}