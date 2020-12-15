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

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*VpcEndpointService)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeVpcEndpointService(r, ctyValue)
}

func DecodeVpcEndpointService(prev *VpcEndpointService, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeVpcEndpointService_AcceptanceRequired(&new.Spec.ForProvider, valMap)
	DecodeVpcEndpointService_AllowedPrincipals(&new.Spec.ForProvider, valMap)
	DecodeVpcEndpointService_NetworkLoadBalancerArns(&new.Spec.ForProvider, valMap)
	DecodeVpcEndpointService_Tags(&new.Spec.ForProvider, valMap)
	DecodeVpcEndpointService_State(&new.Status.AtProvider, valMap)
	DecodeVpcEndpointService_Arn(&new.Status.AtProvider, valMap)
	DecodeVpcEndpointService_AvailabilityZones(&new.Status.AtProvider, valMap)
	DecodeVpcEndpointService_PrivateDnsName(&new.Status.AtProvider, valMap)
	DecodeVpcEndpointService_ServiceName(&new.Status.AtProvider, valMap)
	DecodeVpcEndpointService_BaseEndpointDnsNames(&new.Status.AtProvider, valMap)
	DecodeVpcEndpointService_ManagesVpcEndpoints(&new.Status.AtProvider, valMap)
	DecodeVpcEndpointService_ServiceType(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeVpcEndpointService_AcceptanceRequired(p *VpcEndpointServiceParameters, vals map[string]cty.Value) {
	p.AcceptanceRequired = ctwhy.ValueAsBool(vals["acceptance_required"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeVpcEndpointService_AllowedPrincipals(p *VpcEndpointServiceParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["allowed_principals"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.AllowedPrincipals = goVals
}

//primitiveCollectionTypeDecodeTemplate
func DecodeVpcEndpointService_NetworkLoadBalancerArns(p *VpcEndpointServiceParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["network_load_balancer_arns"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.NetworkLoadBalancerArns = goVals
}

//primitiveMapTypeDecodeTemplate
func DecodeVpcEndpointService_Tags(p *VpcEndpointServiceParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeVpcEndpointService_State(p *VpcEndpointServiceObservation, vals map[string]cty.Value) {
	p.State = ctwhy.ValueAsString(vals["state"])
}

//primitiveTypeDecodeTemplate
func DecodeVpcEndpointService_Arn(p *VpcEndpointServiceObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeVpcEndpointService_AvailabilityZones(p *VpcEndpointServiceObservation, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["availability_zones"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.AvailabilityZones = goVals
}

//primitiveTypeDecodeTemplate
func DecodeVpcEndpointService_PrivateDnsName(p *VpcEndpointServiceObservation, vals map[string]cty.Value) {
	p.PrivateDnsName = ctwhy.ValueAsString(vals["private_dns_name"])
}

//primitiveTypeDecodeTemplate
func DecodeVpcEndpointService_ServiceName(p *VpcEndpointServiceObservation, vals map[string]cty.Value) {
	p.ServiceName = ctwhy.ValueAsString(vals["service_name"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeVpcEndpointService_BaseEndpointDnsNames(p *VpcEndpointServiceObservation, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["base_endpoint_dns_names"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.BaseEndpointDnsNames = goVals
}

//primitiveTypeDecodeTemplate
func DecodeVpcEndpointService_ManagesVpcEndpoints(p *VpcEndpointServiceObservation, vals map[string]cty.Value) {
	p.ManagesVpcEndpoints = ctwhy.ValueAsBool(vals["manages_vpc_endpoints"])
}

//primitiveTypeDecodeTemplate
func DecodeVpcEndpointService_ServiceType(p *VpcEndpointServiceObservation, vals map[string]cty.Value) {
	p.ServiceType = ctwhy.ValueAsString(vals["service_type"])
}