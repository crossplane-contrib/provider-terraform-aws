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

package v1alpha1func EncodeNatGateway(r NatGateway) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeNatGateway_AllocationId(r.Spec.ForProvider, ctyVal)
	EncodeNatGateway_Id(r.Spec.ForProvider, ctyVal)
	EncodeNatGateway_SubnetId(r.Spec.ForProvider, ctyVal)
	EncodeNatGateway_Tags(r.Spec.ForProvider, ctyVal)
	EncodeNatGateway_NetworkInterfaceId(r.Status.AtProvider, ctyVal)
	EncodeNatGateway_PrivateIp(r.Status.AtProvider, ctyVal)
	EncodeNatGateway_PublicIp(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeNatGateway_AllocationId(p *NatGatewayParameters, vals map[string]cty.Value) {
	vals["allocation_id"] = cty.StringVal(p.AllocationId)
}

func EncodeNatGateway_Id(p *NatGatewayParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeNatGateway_SubnetId(p *NatGatewayParameters, vals map[string]cty.Value) {
	vals["subnet_id"] = cty.StringVal(p.SubnetId)
}

func EncodeNatGateway_Tags(p *NatGatewayParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeNatGateway_NetworkInterfaceId(p *NatGatewayObservation, vals map[string]cty.Value) {
	vals["network_interface_id"] = cty.StringVal(p.NetworkInterfaceId)
}

func EncodeNatGateway_PrivateIp(p *NatGatewayObservation, vals map[string]cty.Value) {
	vals["private_ip"] = cty.StringVal(p.PrivateIp)
}

func EncodeNatGateway_PublicIp(p *NatGatewayObservation, vals map[string]cty.Value) {
	vals["public_ip"] = cty.StringVal(p.PublicIp)
}