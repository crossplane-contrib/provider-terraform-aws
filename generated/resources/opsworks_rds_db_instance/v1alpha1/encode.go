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

func EncodeOpsworksRdsDbInstance(r OpsworksRdsDbInstance) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksRdsDbInstance_DbPassword(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRdsDbInstance_DbUser(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRdsDbInstance_Id(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRdsDbInstance_RdsDbInstanceArn(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRdsDbInstance_StackId(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeOpsworksRdsDbInstance_DbPassword(p OpsworksRdsDbInstanceParameters, vals map[string]cty.Value) {
	vals["db_password"] = cty.StringVal(p.DbPassword)
}

func EncodeOpsworksRdsDbInstance_DbUser(p OpsworksRdsDbInstanceParameters, vals map[string]cty.Value) {
	vals["db_user"] = cty.StringVal(p.DbUser)
}

func EncodeOpsworksRdsDbInstance_Id(p OpsworksRdsDbInstanceParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOpsworksRdsDbInstance_RdsDbInstanceArn(p OpsworksRdsDbInstanceParameters, vals map[string]cty.Value) {
	vals["rds_db_instance_arn"] = cty.StringVal(p.RdsDbInstanceArn)
}

func EncodeOpsworksRdsDbInstance_StackId(p OpsworksRdsDbInstanceParameters, vals map[string]cty.Value) {
	vals["stack_id"] = cty.StringVal(p.StackId)
}