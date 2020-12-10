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
	r, ok := mr.(*IamRole)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a IamRole.")
	}
	return EncodeIamRole(*r), nil
}

func EncodeIamRole(r IamRole) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeIamRole_MaxSessionDuration(r.Spec.ForProvider, ctyVal)
	EncodeIamRole_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeIamRole_PermissionsBoundary(r.Spec.ForProvider, ctyVal)
	EncodeIamRole_Tags(r.Spec.ForProvider, ctyVal)
	EncodeIamRole_ForceDetachPolicies(r.Spec.ForProvider, ctyVal)
	EncodeIamRole_Id(r.Spec.ForProvider, ctyVal)
	EncodeIamRole_AssumeRolePolicy(r.Spec.ForProvider, ctyVal)
	EncodeIamRole_Description(r.Spec.ForProvider, ctyVal)
	EncodeIamRole_Name(r.Spec.ForProvider, ctyVal)
	EncodeIamRole_Path(r.Spec.ForProvider, ctyVal)
	EncodeIamRole_Arn(r.Status.AtProvider, ctyVal)
	EncodeIamRole_CreateDate(r.Status.AtProvider, ctyVal)
	EncodeIamRole_UniqueId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeIamRole_MaxSessionDuration(p IamRoleParameters, vals map[string]cty.Value) {
	vals["max_session_duration"] = cty.NumberIntVal(p.MaxSessionDuration)
}

func EncodeIamRole_NamePrefix(p IamRoleParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeIamRole_PermissionsBoundary(p IamRoleParameters, vals map[string]cty.Value) {
	vals["permissions_boundary"] = cty.StringVal(p.PermissionsBoundary)
}

func EncodeIamRole_Tags(p IamRoleParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeIamRole_ForceDetachPolicies(p IamRoleParameters, vals map[string]cty.Value) {
	vals["force_detach_policies"] = cty.BoolVal(p.ForceDetachPolicies)
}

func EncodeIamRole_Id(p IamRoleParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeIamRole_AssumeRolePolicy(p IamRoleParameters, vals map[string]cty.Value) {
	vals["assume_role_policy"] = cty.StringVal(p.AssumeRolePolicy)
}

func EncodeIamRole_Description(p IamRoleParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeIamRole_Name(p IamRoleParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeIamRole_Path(p IamRoleParameters, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeIamRole_Arn(p IamRoleObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeIamRole_CreateDate(p IamRoleObservation, vals map[string]cty.Value) {
	vals["create_date"] = cty.StringVal(p.CreateDate)
}

func EncodeIamRole_UniqueId(p IamRoleObservation, vals map[string]cty.Value) {
	vals["unique_id"] = cty.StringVal(p.UniqueId)
}