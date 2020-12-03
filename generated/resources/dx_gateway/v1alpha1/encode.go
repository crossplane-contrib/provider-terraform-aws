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

package v1alpha1func EncodeDxGateway(r DxGateway) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeDxGateway_Id(r.Spec.ForProvider, ctyVal)
	EncodeDxGateway_Name(r.Spec.ForProvider, ctyVal)
	EncodeDxGateway_AmazonSideAsn(r.Spec.ForProvider, ctyVal)
	EncodeDxGateway_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDxGateway_OwnerAccountId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeDxGateway_Id(p *DxGatewayParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDxGateway_Name(p *DxGatewayParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDxGateway_AmazonSideAsn(p *DxGatewayParameters, vals map[string]cty.Value) {
	vals["amazon_side_asn"] = cty.StringVal(p.AmazonSideAsn)
}

func EncodeDxGateway_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	ctyVal = make(map[string]cty.Value)
	EncodeDxGateway_Timeouts_Create(p, ctyVal)
	EncodeDxGateway_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDxGateway_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDxGateway_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeDxGateway_OwnerAccountId(p *DxGatewayObservation, vals map[string]cty.Value) {
	vals["owner_account_id"] = cty.StringVal(p.OwnerAccountId)
}