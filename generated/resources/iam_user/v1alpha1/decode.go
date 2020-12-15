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
	r, ok := mr.(*IamUser)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeIamUser(r, ctyValue)
}

func DecodeIamUser(prev *IamUser, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeIamUser_ForceDestroy(&new.Spec.ForProvider, valMap)
	DecodeIamUser_Name(&new.Spec.ForProvider, valMap)
	DecodeIamUser_Path(&new.Spec.ForProvider, valMap)
	DecodeIamUser_PermissionsBoundary(&new.Spec.ForProvider, valMap)
	DecodeIamUser_Tags(&new.Spec.ForProvider, valMap)
	DecodeIamUser_Arn(&new.Status.AtProvider, valMap)
	DecodeIamUser_UniqueId(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeIamUser_ForceDestroy(p *IamUserParameters, vals map[string]cty.Value) {
	p.ForceDestroy = ctwhy.ValueAsBool(vals["force_destroy"])
}

//primitiveTypeDecodeTemplate
func DecodeIamUser_Name(p *IamUserParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeIamUser_Path(p *IamUserParameters, vals map[string]cty.Value) {
	p.Path = ctwhy.ValueAsString(vals["path"])
}

//primitiveTypeDecodeTemplate
func DecodeIamUser_PermissionsBoundary(p *IamUserParameters, vals map[string]cty.Value) {
	p.PermissionsBoundary = ctwhy.ValueAsString(vals["permissions_boundary"])
}

//primitiveMapTypeDecodeTemplate
func DecodeIamUser_Tags(p *IamUserParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeIamUser_Arn(p *IamUserObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeIamUser_UniqueId(p *IamUserObservation, vals map[string]cty.Value) {
	p.UniqueId = ctwhy.ValueAsString(vals["unique_id"])
}