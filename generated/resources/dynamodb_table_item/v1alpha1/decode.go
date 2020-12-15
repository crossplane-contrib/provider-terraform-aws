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
	r, ok := mr.(*DynamodbTableItem)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDynamodbTableItem(r, ctyValue)
}

func DecodeDynamodbTableItem(prev *DynamodbTableItem, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDynamodbTableItem_HashKey(&new.Spec.ForProvider, valMap)
	DecodeDynamodbTableItem_Item(&new.Spec.ForProvider, valMap)
	DecodeDynamodbTableItem_RangeKey(&new.Spec.ForProvider, valMap)
	DecodeDynamodbTableItem_TableName(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeDynamodbTableItem_HashKey(p *DynamodbTableItemParameters, vals map[string]cty.Value) {
	p.HashKey = ctwhy.ValueAsString(vals["hash_key"])
}

//primitiveTypeDecodeTemplate
func DecodeDynamodbTableItem_Item(p *DynamodbTableItemParameters, vals map[string]cty.Value) {
	p.Item = ctwhy.ValueAsString(vals["item"])
}

//primitiveTypeDecodeTemplate
func DecodeDynamodbTableItem_RangeKey(p *DynamodbTableItemParameters, vals map[string]cty.Value) {
	p.RangeKey = ctwhy.ValueAsString(vals["range_key"])
}

//primitiveTypeDecodeTemplate
func DecodeDynamodbTableItem_TableName(p *DynamodbTableItemParameters, vals map[string]cty.Value) {
	p.TableName = ctwhy.ValueAsString(vals["table_name"])
}