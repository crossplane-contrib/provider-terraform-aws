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
	"github.com/zclconf/go-cty/cty"
)

func EncodeLoadBalancerListenerPolicy(r LoadBalancerListenerPolicy) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeLoadBalancerListenerPolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeLoadBalancerListenerPolicy_LoadBalancerName(r.Spec.ForProvider, ctyVal)
	EncodeLoadBalancerListenerPolicy_LoadBalancerPort(r.Spec.ForProvider, ctyVal)
	EncodeLoadBalancerListenerPolicy_PolicyNames(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeLoadBalancerListenerPolicy_Id(p LoadBalancerListenerPolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLoadBalancerListenerPolicy_LoadBalancerName(p LoadBalancerListenerPolicyParameters, vals map[string]cty.Value) {
	vals["load_balancer_name"] = cty.StringVal(p.LoadBalancerName)
}

func EncodeLoadBalancerListenerPolicy_LoadBalancerPort(p LoadBalancerListenerPolicyParameters, vals map[string]cty.Value) {
	vals["load_balancer_port"] = cty.NumberIntVal(p.LoadBalancerPort)
}

func EncodeLoadBalancerListenerPolicy_PolicyNames(p LoadBalancerListenerPolicyParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.PolicyNames {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["policy_names"] = cty.SetVal(colVals)
}