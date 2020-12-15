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
	r, ok := mr.(*DxGateway)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDxGateway(r, ctyValue)
}

func DecodeDxGateway(prev *DxGateway, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDxGateway_AmazonSideAsn(&new.Spec.ForProvider, valMap)
	DecodeDxGateway_Name(&new.Spec.ForProvider, valMap)
	DecodeDxGateway_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeDxGateway_OwnerAccountId(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeDxGateway_AmazonSideAsn(p *DxGatewayParameters, vals map[string]cty.Value) {
	p.AmazonSideAsn = ctwhy.ValueAsString(vals["amazon_side_asn"])
}

//primitiveTypeDecodeTemplate
func DecodeDxGateway_Name(p *DxGatewayParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//containerTypeDecodeTemplate
func DecodeDxGateway_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeDxGateway_Timeouts_Delete(p, valMap)
	DecodeDxGateway_Timeouts_Create(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeDxGateway_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

//primitiveTypeDecodeTemplate
func DecodeDxGateway_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveTypeDecodeTemplate
func DecodeDxGateway_OwnerAccountId(p *DxGatewayObservation, vals map[string]cty.Value) {
	p.OwnerAccountId = ctwhy.ValueAsString(vals["owner_account_id"])
}