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
	r, ok := mr.(*OpsworksRdsDbInstance)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a OpsworksRdsDbInstance.")
	}
	return EncodeOpsworksRdsDbInstance(*r), nil
}

func EncodeOpsworksRdsDbInstance(r OpsworksRdsDbInstance) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksRdsDbInstance_DbPassword(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRdsDbInstance_DbUser(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRdsDbInstance_Id(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRdsDbInstance_RdsDbInstanceArn(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRdsDbInstance_StackId(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
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