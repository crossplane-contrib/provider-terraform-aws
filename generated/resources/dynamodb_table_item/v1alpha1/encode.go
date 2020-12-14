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
	r, ok := mr.(*DynamodbTableItem)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DynamodbTableItem.")
	}
	return EncodeDynamodbTableItem(*r), nil
}

func EncodeDynamodbTableItem(r DynamodbTableItem) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDynamodbTableItem_RangeKey(r.Spec.ForProvider, ctyVal)
	EncodeDynamodbTableItem_TableName(r.Spec.ForProvider, ctyVal)
	EncodeDynamodbTableItem_HashKey(r.Spec.ForProvider, ctyVal)
	EncodeDynamodbTableItem_Id(r.Spec.ForProvider, ctyVal)
	EncodeDynamodbTableItem_Item(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeDynamodbTableItem_RangeKey(p DynamodbTableItemParameters, vals map[string]cty.Value) {
	vals["range_key"] = cty.StringVal(p.RangeKey)
}

func EncodeDynamodbTableItem_TableName(p DynamodbTableItemParameters, vals map[string]cty.Value) {
	vals["table_name"] = cty.StringVal(p.TableName)
}

func EncodeDynamodbTableItem_HashKey(p DynamodbTableItemParameters, vals map[string]cty.Value) {
	vals["hash_key"] = cty.StringVal(p.HashKey)
}

func EncodeDynamodbTableItem_Id(p DynamodbTableItemParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDynamodbTableItem_Item(p DynamodbTableItemParameters, vals map[string]cty.Value) {
	vals["item"] = cty.StringVal(p.Item)
}