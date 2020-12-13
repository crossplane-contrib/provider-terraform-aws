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
	"fmt"

	"github.com/zclconf/go-cty/cty"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*VpcEndpointService)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a VpcEndpointService.")
	}
	return EncodeVpcEndpointService(*r), nil
}

func EncodeVpcEndpointService(r VpcEndpointService) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeVpcEndpointService_NetworkLoadBalancerArns(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpointService_AllowedPrincipals(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpointService_Id(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpointService_Tags(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpointService_AcceptanceRequired(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpointService_BaseEndpointDnsNames(r.Status.AtProvider, ctyVal)
	EncodeVpcEndpointService_PrivateDnsName(r.Status.AtProvider, ctyVal)
	EncodeVpcEndpointService_ServiceType(r.Status.AtProvider, ctyVal)
	EncodeVpcEndpointService_State(r.Status.AtProvider, ctyVal)
	EncodeVpcEndpointService_Arn(r.Status.AtProvider, ctyVal)
	EncodeVpcEndpointService_AvailabilityZones(r.Status.AtProvider, ctyVal)
	EncodeVpcEndpointService_ManagesVpcEndpoints(r.Status.AtProvider, ctyVal)
	EncodeVpcEndpointService_ServiceName(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeVpcEndpointService_NetworkLoadBalancerArns(p VpcEndpointServiceParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.NetworkLoadBalancerArns {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["network_load_balancer_arns"] = cty.SetVal(colVals)
}

func EncodeVpcEndpointService_AllowedPrincipals(p VpcEndpointServiceParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AllowedPrincipals {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["allowed_principals"] = cty.SetVal(colVals)
}

func EncodeVpcEndpointService_Id(p VpcEndpointServiceParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeVpcEndpointService_Tags(p VpcEndpointServiceParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeVpcEndpointService_AcceptanceRequired(p VpcEndpointServiceParameters, vals map[string]cty.Value) {
	vals["acceptance_required"] = cty.BoolVal(p.AcceptanceRequired)
}

func EncodeVpcEndpointService_BaseEndpointDnsNames(p VpcEndpointServiceObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.BaseEndpointDnsNames {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["base_endpoint_dns_names"] = cty.SetVal(colVals)
}

func EncodeVpcEndpointService_PrivateDnsName(p VpcEndpointServiceObservation, vals map[string]cty.Value) {
	vals["private_dns_name"] = cty.StringVal(p.PrivateDnsName)
}

func EncodeVpcEndpointService_ServiceType(p VpcEndpointServiceObservation, vals map[string]cty.Value) {
	vals["service_type"] = cty.StringVal(p.ServiceType)
}

func EncodeVpcEndpointService_State(p VpcEndpointServiceObservation, vals map[string]cty.Value) {
	vals["state"] = cty.StringVal(p.State)
}

func EncodeVpcEndpointService_Arn(p VpcEndpointServiceObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeVpcEndpointService_AvailabilityZones(p VpcEndpointServiceObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AvailabilityZones {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["availability_zones"] = cty.SetVal(colVals)
}

func EncodeVpcEndpointService_ManagesVpcEndpoints(p VpcEndpointServiceObservation, vals map[string]cty.Value) {
	vals["manages_vpc_endpoints"] = cty.BoolVal(p.ManagesVpcEndpoints)
}

func EncodeVpcEndpointService_ServiceName(p VpcEndpointServiceObservation, vals map[string]cty.Value) {
	vals["service_name"] = cty.StringVal(p.ServiceName)
}