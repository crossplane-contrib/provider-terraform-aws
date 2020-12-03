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

package v1alpha1func EncodeLbSslNegotiationPolicy(r LbSslNegotiationPolicy) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeLbSslNegotiationPolicy_Name(r.Spec.ForProvider, ctyVal)
	EncodeLbSslNegotiationPolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeLbSslNegotiationPolicy_LbPort(r.Spec.ForProvider, ctyVal)
	EncodeLbSslNegotiationPolicy_LoadBalancer(r.Spec.ForProvider, ctyVal)
	EncodeLbSslNegotiationPolicy_Attribute(r.Spec.ForProvider.Attribute, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeLbSslNegotiationPolicy_Name(p *LbSslNegotiationPolicyParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeLbSslNegotiationPolicy_Id(p *LbSslNegotiationPolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLbSslNegotiationPolicy_LbPort(p *LbSslNegotiationPolicyParameters, vals map[string]cty.Value) {
	vals["lb_port"] = cty.IntVal(p.LbPort)
}

func EncodeLbSslNegotiationPolicy_LoadBalancer(p *LbSslNegotiationPolicyParameters, vals map[string]cty.Value) {
	vals["load_balancer"] = cty.StringVal(p.LoadBalancer)
}

func EncodeLbSslNegotiationPolicy_Attribute(p *Attribute, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Attribute {
		ctyVal = make(map[string]cty.Value)
		EncodeLbSslNegotiationPolicy_Attribute_Name(v, ctyVal)
		EncodeLbSslNegotiationPolicy_Attribute_Value(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["attribute"] = cty.SetVal(valsForCollection)
}

func EncodeLbSslNegotiationPolicy_Attribute_Name(p *Attribute, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeLbSslNegotiationPolicy_Attribute_Value(p *Attribute, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}