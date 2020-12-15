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
	r, ok := mr.(*IamInstanceProfile)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a IamInstanceProfile.")
	}
	return EncodeIamInstanceProfile(*r), nil
}

func EncodeIamInstanceProfile(r IamInstanceProfile) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeIamInstanceProfile_Name(r.Spec.ForProvider, ctyVal)
	EncodeIamInstanceProfile_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeIamInstanceProfile_Path(r.Spec.ForProvider, ctyVal)
	EncodeIamInstanceProfile_Role(r.Spec.ForProvider, ctyVal)
	EncodeIamInstanceProfile_UniqueId(r.Status.AtProvider, ctyVal)
	EncodeIamInstanceProfile_Arn(r.Status.AtProvider, ctyVal)
	EncodeIamInstanceProfile_CreateDate(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeIamInstanceProfile_Name(p IamInstanceProfileParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeIamInstanceProfile_NamePrefix(p IamInstanceProfileParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeIamInstanceProfile_Path(p IamInstanceProfileParameters, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeIamInstanceProfile_Role(p IamInstanceProfileParameters, vals map[string]cty.Value) {
	vals["role"] = cty.StringVal(p.Role)
}

func EncodeIamInstanceProfile_UniqueId(p IamInstanceProfileObservation, vals map[string]cty.Value) {
	vals["unique_id"] = cty.StringVal(p.UniqueId)
}

func EncodeIamInstanceProfile_Arn(p IamInstanceProfileObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeIamInstanceProfile_CreateDate(p IamInstanceProfileObservation, vals map[string]cty.Value) {
	vals["create_date"] = cty.StringVal(p.CreateDate)
}