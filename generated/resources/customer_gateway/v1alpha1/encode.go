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

package v1alpha1func EncodeCustomerGateway(r CustomerGateway) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeCustomerGateway_Id(r.Spec.ForProvider, ctyVal)
	EncodeCustomerGateway_IpAddress(r.Spec.ForProvider, ctyVal)
	EncodeCustomerGateway_Tags(r.Spec.ForProvider, ctyVal)
	EncodeCustomerGateway_Type(r.Spec.ForProvider, ctyVal)
	EncodeCustomerGateway_BgpAsn(r.Spec.ForProvider, ctyVal)
	EncodeCustomerGateway_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeCustomerGateway_Id(p *CustomerGatewayParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCustomerGateway_IpAddress(p *CustomerGatewayParameters, vals map[string]cty.Value) {
	vals["ip_address"] = cty.StringVal(p.IpAddress)
}

func EncodeCustomerGateway_Tags(p *CustomerGatewayParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeCustomerGateway_Type(p *CustomerGatewayParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeCustomerGateway_BgpAsn(p *CustomerGatewayParameters, vals map[string]cty.Value) {
	vals["bgp_asn"] = cty.StringVal(p.BgpAsn)
}

func EncodeCustomerGateway_Arn(p *CustomerGatewayObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}