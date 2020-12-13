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
	r, ok := mr.(*Route53ResolverEndpoint)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Route53ResolverEndpoint.")
	}
	return EncodeRoute53ResolverEndpoint(*r), nil
}

func EncodeRoute53ResolverEndpoint(r Route53ResolverEndpoint) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRoute53ResolverEndpoint_Id(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverEndpoint_Name(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverEndpoint_SecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverEndpoint_Tags(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverEndpoint_Direction(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverEndpoint_IpAddress(r.Spec.ForProvider.IpAddress, ctyVal)
	EncodeRoute53ResolverEndpoint_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeRoute53ResolverEndpoint_HostVpcId(r.Status.AtProvider, ctyVal)
	EncodeRoute53ResolverEndpoint_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeRoute53ResolverEndpoint_Id(p Route53ResolverEndpointParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRoute53ResolverEndpoint_Name(p Route53ResolverEndpointParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeRoute53ResolverEndpoint_SecurityGroupIds(p Route53ResolverEndpointParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_group_ids"] = cty.SetVal(colVals)
}

func EncodeRoute53ResolverEndpoint_Tags(p Route53ResolverEndpointParameters, vals map[string]cty.Value) {
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

func EncodeRoute53ResolverEndpoint_Direction(p Route53ResolverEndpointParameters, vals map[string]cty.Value) {
	vals["direction"] = cty.StringVal(p.Direction)
}

func EncodeRoute53ResolverEndpoint_IpAddress(p []IpAddress, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeRoute53ResolverEndpoint_IpAddress_IpId(v, ctyVal)
		EncodeRoute53ResolverEndpoint_IpAddress_SubnetId(v, ctyVal)
		EncodeRoute53ResolverEndpoint_IpAddress_Ip(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_address"] = cty.SetVal(valsForCollection)
}

func EncodeRoute53ResolverEndpoint_IpAddress_IpId(p IpAddress, vals map[string]cty.Value) {
	vals["ip_id"] = cty.StringVal(p.IpId)
}

func EncodeRoute53ResolverEndpoint_IpAddress_SubnetId(p IpAddress, vals map[string]cty.Value) {
	vals["subnet_id"] = cty.StringVal(p.SubnetId)
}

func EncodeRoute53ResolverEndpoint_IpAddress_Ip(p IpAddress, vals map[string]cty.Value) {
	vals["ip"] = cty.StringVal(p.Ip)
}

func EncodeRoute53ResolverEndpoint_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeRoute53ResolverEndpoint_Timeouts_Create(p, ctyVal)
	EncodeRoute53ResolverEndpoint_Timeouts_Delete(p, ctyVal)
	EncodeRoute53ResolverEndpoint_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeRoute53ResolverEndpoint_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeRoute53ResolverEndpoint_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeRoute53ResolverEndpoint_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeRoute53ResolverEndpoint_HostVpcId(p Route53ResolverEndpointObservation, vals map[string]cty.Value) {
	vals["host_vpc_id"] = cty.StringVal(p.HostVpcId)
}

func EncodeRoute53ResolverEndpoint_Arn(p Route53ResolverEndpointObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}