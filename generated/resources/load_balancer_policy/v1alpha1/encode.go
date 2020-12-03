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

package v1alpha1func EncodeLoadBalancerPolicy(r LoadBalancerPolicy) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeLoadBalancerPolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeLoadBalancerPolicy_LoadBalancerName(r.Spec.ForProvider, ctyVal)
	EncodeLoadBalancerPolicy_PolicyName(r.Spec.ForProvider, ctyVal)
	EncodeLoadBalancerPolicy_PolicyTypeName(r.Spec.ForProvider, ctyVal)
	EncodeLoadBalancerPolicy_PolicyAttribute(r.Spec.ForProvider.PolicyAttribute, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeLoadBalancerPolicy_Id(p *LoadBalancerPolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLoadBalancerPolicy_LoadBalancerName(p *LoadBalancerPolicyParameters, vals map[string]cty.Value) {
	vals["load_balancer_name"] = cty.StringVal(p.LoadBalancerName)
}

func EncodeLoadBalancerPolicy_PolicyName(p *LoadBalancerPolicyParameters, vals map[string]cty.Value) {
	vals["policy_name"] = cty.StringVal(p.PolicyName)
}

func EncodeLoadBalancerPolicy_PolicyTypeName(p *LoadBalancerPolicyParameters, vals map[string]cty.Value) {
	vals["policy_type_name"] = cty.StringVal(p.PolicyTypeName)
}

func EncodeLoadBalancerPolicy_PolicyAttribute(p *PolicyAttribute, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.PolicyAttribute {
		ctyVal = make(map[string]cty.Value)
		EncodeLoadBalancerPolicy_PolicyAttribute_Name(v, ctyVal)
		EncodeLoadBalancerPolicy_PolicyAttribute_Value(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["policy_attribute"] = cty.SetVal(valsForCollection)
}

func EncodeLoadBalancerPolicy_PolicyAttribute_Name(p *PolicyAttribute, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeLoadBalancerPolicy_PolicyAttribute_Value(p *PolicyAttribute, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}