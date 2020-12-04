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

func EncodeLoadBalancerBackendServerPolicy(r LoadBalancerBackendServerPolicy) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeLoadBalancerBackendServerPolicy_LoadBalancerName(r.Spec.ForProvider, ctyVal)
	EncodeLoadBalancerBackendServerPolicy_PolicyNames(r.Spec.ForProvider, ctyVal)
	EncodeLoadBalancerBackendServerPolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeLoadBalancerBackendServerPolicy_InstancePort(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeLoadBalancerBackendServerPolicy_LoadBalancerName(p LoadBalancerBackendServerPolicyParameters, vals map[string]cty.Value) {
	vals["load_balancer_name"] = cty.StringVal(p.LoadBalancerName)
}

func EncodeLoadBalancerBackendServerPolicy_PolicyNames(p LoadBalancerBackendServerPolicyParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.PolicyNames {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["policy_names"] = cty.SetVal(colVals)
}

func EncodeLoadBalancerBackendServerPolicy_Id(p LoadBalancerBackendServerPolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLoadBalancerBackendServerPolicy_InstancePort(p LoadBalancerBackendServerPolicyParameters, vals map[string]cty.Value) {
	vals["instance_port"] = cty.NumberIntVal(p.InstancePort)
}