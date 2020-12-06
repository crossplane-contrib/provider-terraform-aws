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

func EncodeEip(r Eip) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEip_PublicIpv4Pool(r.Spec.ForProvider, ctyVal)
	EncodeEip_Vpc(r.Spec.ForProvider, ctyVal)
	EncodeEip_Id(r.Spec.ForProvider, ctyVal)
	EncodeEip_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEip_AssociateWithPrivateIp(r.Spec.ForProvider, ctyVal)
	EncodeEip_CustomerOwnedIpv4Pool(r.Spec.ForProvider, ctyVal)
	EncodeEip_Instance(r.Spec.ForProvider, ctyVal)
	EncodeEip_NetworkInterface(r.Spec.ForProvider, ctyVal)
	EncodeEip_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeEip_CustomerOwnedIp(r.Status.AtProvider, ctyVal)
	EncodeEip_Domain(r.Status.AtProvider, ctyVal)
	EncodeEip_PublicDns(r.Status.AtProvider, ctyVal)
	EncodeEip_AllocationId(r.Status.AtProvider, ctyVal)
	EncodeEip_PublicIp(r.Status.AtProvider, ctyVal)
	EncodeEip_PrivateIp(r.Status.AtProvider, ctyVal)
	EncodeEip_AssociationId(r.Status.AtProvider, ctyVal)
	EncodeEip_PrivateDns(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeEip_PublicIpv4Pool(p EipParameters, vals map[string]cty.Value) {
	vals["public_ipv4_pool"] = cty.StringVal(p.PublicIpv4Pool)
}

func EncodeEip_Vpc(p EipParameters, vals map[string]cty.Value) {
	vals["vpc"] = cty.BoolVal(p.Vpc)
}

func EncodeEip_Id(p EipParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEip_Tags(p EipParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeEip_AssociateWithPrivateIp(p EipParameters, vals map[string]cty.Value) {
	vals["associate_with_private_ip"] = cty.StringVal(p.AssociateWithPrivateIp)
}

func EncodeEip_CustomerOwnedIpv4Pool(p EipParameters, vals map[string]cty.Value) {
	vals["customer_owned_ipv4_pool"] = cty.StringVal(p.CustomerOwnedIpv4Pool)
}

func EncodeEip_Instance(p EipParameters, vals map[string]cty.Value) {
	vals["instance"] = cty.StringVal(p.Instance)
}

func EncodeEip_NetworkInterface(p EipParameters, vals map[string]cty.Value) {
	vals["network_interface"] = cty.StringVal(p.NetworkInterface)
}

func EncodeEip_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeEip_Timeouts_Delete(p, ctyVal)
	EncodeEip_Timeouts_Read(p, ctyVal)
	EncodeEip_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeEip_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeEip_Timeouts_Read(p Timeouts, vals map[string]cty.Value) {
	vals["read"] = cty.StringVal(p.Read)
}

func EncodeEip_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeEip_CustomerOwnedIp(p EipObservation, vals map[string]cty.Value) {
	vals["customer_owned_ip"] = cty.StringVal(p.CustomerOwnedIp)
}

func EncodeEip_Domain(p EipObservation, vals map[string]cty.Value) {
	vals["domain"] = cty.StringVal(p.Domain)
}

func EncodeEip_PublicDns(p EipObservation, vals map[string]cty.Value) {
	vals["public_dns"] = cty.StringVal(p.PublicDns)
}

func EncodeEip_AllocationId(p EipObservation, vals map[string]cty.Value) {
	vals["allocation_id"] = cty.StringVal(p.AllocationId)
}

func EncodeEip_PublicIp(p EipObservation, vals map[string]cty.Value) {
	vals["public_ip"] = cty.StringVal(p.PublicIp)
}

func EncodeEip_PrivateIp(p EipObservation, vals map[string]cty.Value) {
	vals["private_ip"] = cty.StringVal(p.PrivateIp)
}

func EncodeEip_AssociationId(p EipObservation, vals map[string]cty.Value) {
	vals["association_id"] = cty.StringVal(p.AssociationId)
}

func EncodeEip_PrivateDns(p EipObservation, vals map[string]cty.Value) {
	vals["private_dns"] = cty.StringVal(p.PrivateDns)
}