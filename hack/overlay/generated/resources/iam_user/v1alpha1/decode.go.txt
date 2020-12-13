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

// hand-written overlay for iam decoding

package v1alpha1

import (
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
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
	DecodeIamUser_Id(&new.Spec.ForProvider, valMap)
	DecodeIamUser_Name(&new.Spec.ForProvider, valMap)
	DecodeIamUser_Path(&new.Spec.ForProvider, valMap)
	DecodeIamUser_PermissionsBoundary(&new.Spec.ForProvider, valMap)
	DecodeIamUser_Tags(&new.Spec.ForProvider, valMap)
	DecodeIamUser_Arn(&new.Status.AtProvider, valMap)
	DecodeIamUser_UniqueId(&new.Status.AtProvider, valMap)
	return new, nil
}

func DecodeIamUser_ForceDestroy(p *IamUserParameters, vals map[string]cty.Value) {
	p.ForceDestroy = vals["force_destroy"].True()
}

func DecodeIamUser_Id(p *IamUserParameters, vals map[string]cty.Value) {
	p.Id = vals["id"].AsString()
}

func DecodeIamUser_Name(p *IamUserParameters, vals map[string]cty.Value) {
	p.Name = vals["name"].AsString()
}

func DecodeIamUser_Path(p *IamUserParameters, vals map[string]cty.Value) {
	p.Path = vals["path"].AsString()
}

func DecodeIamUser_PermissionsBoundary(p *IamUserParameters, vals map[string]cty.Value) {
	p.PermissionsBoundary = vals["permissions_boundary"].AsString()
}

func DecodeIamUser_Tags(p *IamUserParameters, vals map[string]cty.Value) {
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = value.AsString()
	}
	p.Tags = vMap
}

func DecodeIamUser_Arn(p *IamUserObservation, vals map[string]cty.Value) {
	p.Arn = vals["arn"].AsString()
}

func DecodeIamUser_UniqueId(p *IamUserObservation, vals map[string]cty.Value) {
	p.UniqueId = vals["unique_id"].AsString()
}
