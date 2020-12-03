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

package v1alpha1func EncodeServiceDiscoveryPrivateDnsNamespace(r ServiceDiscoveryPrivateDnsNamespace) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeServiceDiscoveryPrivateDnsNamespace_Id(r.Spec.ForProvider, ctyVal)
	EncodeServiceDiscoveryPrivateDnsNamespace_Name(r.Spec.ForProvider, ctyVal)
	EncodeServiceDiscoveryPrivateDnsNamespace_Tags(r.Spec.ForProvider, ctyVal)
	EncodeServiceDiscoveryPrivateDnsNamespace_Vpc(r.Spec.ForProvider, ctyVal)
	EncodeServiceDiscoveryPrivateDnsNamespace_Description(r.Spec.ForProvider, ctyVal)
	EncodeServiceDiscoveryPrivateDnsNamespace_HostedZone(r.Status.AtProvider, ctyVal)
	EncodeServiceDiscoveryPrivateDnsNamespace_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeServiceDiscoveryPrivateDnsNamespace_Id(p *ServiceDiscoveryPrivateDnsNamespaceParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeServiceDiscoveryPrivateDnsNamespace_Name(p *ServiceDiscoveryPrivateDnsNamespaceParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeServiceDiscoveryPrivateDnsNamespace_Tags(p *ServiceDiscoveryPrivateDnsNamespaceParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeServiceDiscoveryPrivateDnsNamespace_Vpc(p *ServiceDiscoveryPrivateDnsNamespaceParameters, vals map[string]cty.Value) {
	vals["vpc"] = cty.StringVal(p.Vpc)
}

func EncodeServiceDiscoveryPrivateDnsNamespace_Description(p *ServiceDiscoveryPrivateDnsNamespaceParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeServiceDiscoveryPrivateDnsNamespace_HostedZone(p *ServiceDiscoveryPrivateDnsNamespaceObservation, vals map[string]cty.Value) {
	vals["hosted_zone"] = cty.StringVal(p.HostedZone)
}

func EncodeServiceDiscoveryPrivateDnsNamespace_Arn(p *ServiceDiscoveryPrivateDnsNamespaceObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}