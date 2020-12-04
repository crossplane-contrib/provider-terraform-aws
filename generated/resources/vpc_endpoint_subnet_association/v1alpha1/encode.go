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

func EncodeVpcEndpointSubnetAssociation(r VpcEndpointSubnetAssociation) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeVpcEndpointSubnetAssociation_VpcEndpointId(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpointSubnetAssociation_Id(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpointSubnetAssociation_SubnetId(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpointSubnetAssociation_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeVpcEndpointSubnetAssociation_VpcEndpointId(p VpcEndpointSubnetAssociationParameters, vals map[string]cty.Value) {
	vals["vpc_endpoint_id"] = cty.StringVal(p.VpcEndpointId)
}

func EncodeVpcEndpointSubnetAssociation_Id(p VpcEndpointSubnetAssociationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeVpcEndpointSubnetAssociation_SubnetId(p VpcEndpointSubnetAssociationParameters, vals map[string]cty.Value) {
	vals["subnet_id"] = cty.StringVal(p.SubnetId)
}

func EncodeVpcEndpointSubnetAssociation_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeVpcEndpointSubnetAssociation_Timeouts_Create(p, ctyVal)
	EncodeVpcEndpointSubnetAssociation_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeVpcEndpointSubnetAssociation_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeVpcEndpointSubnetAssociation_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}