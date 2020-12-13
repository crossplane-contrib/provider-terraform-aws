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
	r, ok := mr.(*DynamodbGlobalTable)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DynamodbGlobalTable.")
	}
	return EncodeDynamodbGlobalTable(*r), nil
}

func EncodeDynamodbGlobalTable(r DynamodbGlobalTable) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDynamodbGlobalTable_Id(r.Spec.ForProvider, ctyVal)
	EncodeDynamodbGlobalTable_Name(r.Spec.ForProvider, ctyVal)
	EncodeDynamodbGlobalTable_Replica(r.Spec.ForProvider.Replica, ctyVal)
	EncodeDynamodbGlobalTable_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDynamodbGlobalTable_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeDynamodbGlobalTable_Id(p DynamodbGlobalTableParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDynamodbGlobalTable_Name(p DynamodbGlobalTableParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDynamodbGlobalTable_Replica(p []Replica, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeDynamodbGlobalTable_Replica_RegionName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["replica"] = cty.SetVal(valsForCollection)
}

func EncodeDynamodbGlobalTable_Replica_RegionName(p Replica, vals map[string]cty.Value) {
	vals["region_name"] = cty.StringVal(p.RegionName)
}

func EncodeDynamodbGlobalTable_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeDynamodbGlobalTable_Timeouts_Create(p, ctyVal)
	EncodeDynamodbGlobalTable_Timeouts_Delete(p, ctyVal)
	EncodeDynamodbGlobalTable_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDynamodbGlobalTable_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDynamodbGlobalTable_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeDynamodbGlobalTable_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeDynamodbGlobalTable_Arn(p DynamodbGlobalTableObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}