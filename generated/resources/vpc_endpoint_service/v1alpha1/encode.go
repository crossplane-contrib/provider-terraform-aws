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

package v1alpha1func EncodeVpcEndpointService(r VpcEndpointService) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeVpcEndpointService_AllowedPrincipals(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpointService_Id(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpointService_NetworkLoadBalancerArns(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpointService_Tags(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpointService_AcceptanceRequired(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpointService_Arn(r.Status.AtProvider, ctyVal)
	EncodeVpcEndpointService_AvailabilityZones(r.Status.AtProvider, ctyVal)
	EncodeVpcEndpointService_BaseEndpointDnsNames(r.Status.AtProvider, ctyVal)
	EncodeVpcEndpointService_ServiceName(r.Status.AtProvider, ctyVal)
	EncodeVpcEndpointService_State(r.Status.AtProvider, ctyVal)
	EncodeVpcEndpointService_ManagesVpcEndpoints(r.Status.AtProvider, ctyVal)
	EncodeVpcEndpointService_PrivateDnsName(r.Status.AtProvider, ctyVal)
	EncodeVpcEndpointService_ServiceType(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeVpcEndpointService_AllowedPrincipals(p *VpcEndpointServiceParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AllowedPrincipals {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["allowed_principals"] = cty.SetVal(colVals)
}

func EncodeVpcEndpointService_Id(p *VpcEndpointServiceParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeVpcEndpointService_NetworkLoadBalancerArns(p *VpcEndpointServiceParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.NetworkLoadBalancerArns {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["network_load_balancer_arns"] = cty.SetVal(colVals)
}

func EncodeVpcEndpointService_Tags(p *VpcEndpointServiceParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeVpcEndpointService_AcceptanceRequired(p *VpcEndpointServiceParameters, vals map[string]cty.Value) {
	vals["acceptance_required"] = cty.BoolVal(p.AcceptanceRequired)
}

func EncodeVpcEndpointService_Arn(p *VpcEndpointServiceObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeVpcEndpointService_AvailabilityZones(p *VpcEndpointServiceObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AvailabilityZones {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["availability_zones"] = cty.SetVal(colVals)
}

func EncodeVpcEndpointService_BaseEndpointDnsNames(p *VpcEndpointServiceObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.BaseEndpointDnsNames {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["base_endpoint_dns_names"] = cty.SetVal(colVals)
}

func EncodeVpcEndpointService_ServiceName(p *VpcEndpointServiceObservation, vals map[string]cty.Value) {
	vals["service_name"] = cty.StringVal(p.ServiceName)
}

func EncodeVpcEndpointService_State(p *VpcEndpointServiceObservation, vals map[string]cty.Value) {
	vals["state"] = cty.StringVal(p.State)
}

func EncodeVpcEndpointService_ManagesVpcEndpoints(p *VpcEndpointServiceObservation, vals map[string]cty.Value) {
	vals["manages_vpc_endpoints"] = cty.BoolVal(p.ManagesVpcEndpoints)
}

func EncodeVpcEndpointService_PrivateDnsName(p *VpcEndpointServiceObservation, vals map[string]cty.Value) {
	vals["private_dns_name"] = cty.StringVal(p.PrivateDnsName)
}

func EncodeVpcEndpointService_ServiceType(p *VpcEndpointServiceObservation, vals map[string]cty.Value) {
	vals["service_type"] = cty.StringVal(p.ServiceType)
}