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
	r, ok := mr.(*Ec2TransitGateway)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Ec2TransitGateway.")
	}
	return EncodeEc2TransitGateway(*r), nil
}

func EncodeEc2TransitGateway(r Ec2TransitGateway) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEc2TransitGateway_AutoAcceptSharedAttachments(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGateway_DefaultRouteTableAssociation(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGateway_Description(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGateway_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGateway_AmazonSideAsn(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGateway_DefaultRouteTablePropagation(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGateway_DnsSupport(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGateway_VpnEcmpSupport(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGateway_OwnerId(r.Status.AtProvider, ctyVal)
	EncodeEc2TransitGateway_PropagationDefaultRouteTableId(r.Status.AtProvider, ctyVal)
	EncodeEc2TransitGateway_AssociationDefaultRouteTableId(r.Status.AtProvider, ctyVal)
	EncodeEc2TransitGateway_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeEc2TransitGateway_AutoAcceptSharedAttachments(p Ec2TransitGatewayParameters, vals map[string]cty.Value) {
	vals["auto_accept_shared_attachments"] = cty.StringVal(p.AutoAcceptSharedAttachments)
}

func EncodeEc2TransitGateway_DefaultRouteTableAssociation(p Ec2TransitGatewayParameters, vals map[string]cty.Value) {
	vals["default_route_table_association"] = cty.StringVal(p.DefaultRouteTableAssociation)
}

func EncodeEc2TransitGateway_Description(p Ec2TransitGatewayParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeEc2TransitGateway_Tags(p Ec2TransitGatewayParameters, vals map[string]cty.Value) {
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

func EncodeEc2TransitGateway_AmazonSideAsn(p Ec2TransitGatewayParameters, vals map[string]cty.Value) {
	vals["amazon_side_asn"] = cty.NumberIntVal(p.AmazonSideAsn)
}

func EncodeEc2TransitGateway_DefaultRouteTablePropagation(p Ec2TransitGatewayParameters, vals map[string]cty.Value) {
	vals["default_route_table_propagation"] = cty.StringVal(p.DefaultRouteTablePropagation)
}

func EncodeEc2TransitGateway_DnsSupport(p Ec2TransitGatewayParameters, vals map[string]cty.Value) {
	vals["dns_support"] = cty.StringVal(p.DnsSupport)
}

func EncodeEc2TransitGateway_VpnEcmpSupport(p Ec2TransitGatewayParameters, vals map[string]cty.Value) {
	vals["vpn_ecmp_support"] = cty.StringVal(p.VpnEcmpSupport)
}

func EncodeEc2TransitGateway_OwnerId(p Ec2TransitGatewayObservation, vals map[string]cty.Value) {
	vals["owner_id"] = cty.StringVal(p.OwnerId)
}

func EncodeEc2TransitGateway_PropagationDefaultRouteTableId(p Ec2TransitGatewayObservation, vals map[string]cty.Value) {
	vals["propagation_default_route_table_id"] = cty.StringVal(p.PropagationDefaultRouteTableId)
}

func EncodeEc2TransitGateway_AssociationDefaultRouteTableId(p Ec2TransitGatewayObservation, vals map[string]cty.Value) {
	vals["association_default_route_table_id"] = cty.StringVal(p.AssociationDefaultRouteTableId)
}

func EncodeEc2TransitGateway_Arn(p Ec2TransitGatewayObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}