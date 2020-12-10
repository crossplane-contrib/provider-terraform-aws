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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*LightsailStaticIpAttachment)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a LightsailStaticIpAttachment.")
	}
	return EncodeLightsailStaticIpAttachment(*r), nil
}

func EncodeLightsailStaticIpAttachment(r LightsailStaticIpAttachment) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeLightsailStaticIpAttachment_Id(r.Spec.ForProvider, ctyVal)
	EncodeLightsailStaticIpAttachment_InstanceName(r.Spec.ForProvider, ctyVal)
	EncodeLightsailStaticIpAttachment_StaticIpName(r.Spec.ForProvider, ctyVal)
	EncodeLightsailStaticIpAttachment_IpAddress(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeLightsailStaticIpAttachment_Id(p LightsailStaticIpAttachmentParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLightsailStaticIpAttachment_InstanceName(p LightsailStaticIpAttachmentParameters, vals map[string]cty.Value) {
	vals["instance_name"] = cty.StringVal(p.InstanceName)
}

func EncodeLightsailStaticIpAttachment_StaticIpName(p LightsailStaticIpAttachmentParameters, vals map[string]cty.Value) {
	vals["static_ip_name"] = cty.StringVal(p.StaticIpName)
}

func EncodeLightsailStaticIpAttachment_IpAddress(p LightsailStaticIpAttachmentObservation, vals map[string]cty.Value) {
	vals["ip_address"] = cty.StringVal(p.IpAddress)
}