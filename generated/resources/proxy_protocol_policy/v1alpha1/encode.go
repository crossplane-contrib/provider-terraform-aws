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

package v1alpha1func EncodeProxyProtocolPolicy(r ProxyProtocolPolicy) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeProxyProtocolPolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeProxyProtocolPolicy_InstancePorts(r.Spec.ForProvider, ctyVal)
	EncodeProxyProtocolPolicy_LoadBalancer(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeProxyProtocolPolicy_Id(p *ProxyProtocolPolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeProxyProtocolPolicy_InstancePorts(p *ProxyProtocolPolicyParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.InstancePorts {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["instance_ports"] = cty.SetVal(colVals)
}

func EncodeProxyProtocolPolicy_LoadBalancer(p *ProxyProtocolPolicyParameters, vals map[string]cty.Value) {
	vals["load_balancer"] = cty.StringVal(p.LoadBalancer)
}