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
	"github.com/zclconf/go-cty/cty"
)

func EncodeDbInstanceRoleAssociation(r DbInstanceRoleAssociation) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDbInstanceRoleAssociation_FeatureName(r.Spec.ForProvider, ctyVal)
	EncodeDbInstanceRoleAssociation_Id(r.Spec.ForProvider, ctyVal)
	EncodeDbInstanceRoleAssociation_RoleArn(r.Spec.ForProvider, ctyVal)
	EncodeDbInstanceRoleAssociation_DbInstanceIdentifier(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeDbInstanceRoleAssociation_FeatureName(p DbInstanceRoleAssociationParameters, vals map[string]cty.Value) {
	vals["feature_name"] = cty.StringVal(p.FeatureName)
}

func EncodeDbInstanceRoleAssociation_Id(p DbInstanceRoleAssociationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDbInstanceRoleAssociation_RoleArn(p DbInstanceRoleAssociationParameters, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeDbInstanceRoleAssociation_DbInstanceIdentifier(p DbInstanceRoleAssociationParameters, vals map[string]cty.Value) {
	vals["db_instance_identifier"] = cty.StringVal(p.DbInstanceIdentifier)
}