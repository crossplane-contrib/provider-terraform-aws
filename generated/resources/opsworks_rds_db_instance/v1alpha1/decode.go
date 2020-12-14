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
	r, ok := mr.(*OpsworksRdsDbInstance)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeOpsworksRdsDbInstance(r, ctyValue)
}

func DecodeOpsworksRdsDbInstance(prev *OpsworksRdsDbInstance, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeOpsworksRdsDbInstance_Id(&new.Spec.ForProvider, valMap)
	DecodeOpsworksRdsDbInstance_RdsDbInstanceArn(&new.Spec.ForProvider, valMap)
	DecodeOpsworksRdsDbInstance_StackId(&new.Spec.ForProvider, valMap)
	DecodeOpsworksRdsDbInstance_DbPassword(&new.Spec.ForProvider, valMap)
	DecodeOpsworksRdsDbInstance_DbUser(&new.Spec.ForProvider, valMap)

	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeOpsworksRdsDbInstance_Id(p *OpsworksRdsDbInstanceParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeOpsworksRdsDbInstance_RdsDbInstanceArn(p *OpsworksRdsDbInstanceParameters, vals map[string]cty.Value) {
	p.RdsDbInstanceArn = ctwhy.ValueAsString(vals["rds_db_instance_arn"])
}

func DecodeOpsworksRdsDbInstance_StackId(p *OpsworksRdsDbInstanceParameters, vals map[string]cty.Value) {
	p.StackId = ctwhy.ValueAsString(vals["stack_id"])
}

func DecodeOpsworksRdsDbInstance_DbPassword(p *OpsworksRdsDbInstanceParameters, vals map[string]cty.Value) {
	p.DbPassword = ctwhy.ValueAsString(vals["db_password"])
}

func DecodeOpsworksRdsDbInstance_DbUser(p *OpsworksRdsDbInstanceParameters, vals map[string]cty.Value) {
	p.DbUser = ctwhy.ValueAsString(vals["db_user"])
}