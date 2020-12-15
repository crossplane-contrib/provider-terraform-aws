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
	r, ok := mr.(*EipAssociation)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a EipAssociation.")
	}
	return EncodeEipAssociation(*r), nil
}

func EncodeEipAssociation(r EipAssociation) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEipAssociation_NetworkInterfaceId(r.Spec.ForProvider, ctyVal)
	EncodeEipAssociation_PrivateIpAddress(r.Spec.ForProvider, ctyVal)
	EncodeEipAssociation_PublicIp(r.Spec.ForProvider, ctyVal)
	EncodeEipAssociation_AllocationId(r.Spec.ForProvider, ctyVal)
	EncodeEipAssociation_AllowReassociation(r.Spec.ForProvider, ctyVal)
	EncodeEipAssociation_InstanceId(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeEipAssociation_NetworkInterfaceId(p EipAssociationParameters, vals map[string]cty.Value) {
	vals["network_interface_id"] = cty.StringVal(p.NetworkInterfaceId)
}

func EncodeEipAssociation_PrivateIpAddress(p EipAssociationParameters, vals map[string]cty.Value) {
	vals["private_ip_address"] = cty.StringVal(p.PrivateIpAddress)
}

func EncodeEipAssociation_PublicIp(p EipAssociationParameters, vals map[string]cty.Value) {
	vals["public_ip"] = cty.StringVal(p.PublicIp)
}

func EncodeEipAssociation_AllocationId(p EipAssociationParameters, vals map[string]cty.Value) {
	vals["allocation_id"] = cty.StringVal(p.AllocationId)
}

func EncodeEipAssociation_AllowReassociation(p EipAssociationParameters, vals map[string]cty.Value) {
	vals["allow_reassociation"] = cty.BoolVal(p.AllowReassociation)
}

func EncodeEipAssociation_InstanceId(p EipAssociationParameters, vals map[string]cty.Value) {
	vals["instance_id"] = cty.StringVal(p.InstanceId)
}