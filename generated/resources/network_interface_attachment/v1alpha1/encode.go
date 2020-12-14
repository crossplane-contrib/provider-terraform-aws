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
	r, ok := mr.(*NetworkInterfaceAttachment)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a NetworkInterfaceAttachment.")
	}
	return EncodeNetworkInterfaceAttachment(*r), nil
}

func EncodeNetworkInterfaceAttachment(r NetworkInterfaceAttachment) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeNetworkInterfaceAttachment_Id(r.Spec.ForProvider, ctyVal)
	EncodeNetworkInterfaceAttachment_InstanceId(r.Spec.ForProvider, ctyVal)
	EncodeNetworkInterfaceAttachment_NetworkInterfaceId(r.Spec.ForProvider, ctyVal)
	EncodeNetworkInterfaceAttachment_DeviceIndex(r.Spec.ForProvider, ctyVal)
	EncodeNetworkInterfaceAttachment_Status(r.Status.AtProvider, ctyVal)
	EncodeNetworkInterfaceAttachment_AttachmentId(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeNetworkInterfaceAttachment_Id(p NetworkInterfaceAttachmentParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeNetworkInterfaceAttachment_InstanceId(p NetworkInterfaceAttachmentParameters, vals map[string]cty.Value) {
	vals["instance_id"] = cty.StringVal(p.InstanceId)
}

func EncodeNetworkInterfaceAttachment_NetworkInterfaceId(p NetworkInterfaceAttachmentParameters, vals map[string]cty.Value) {
	vals["network_interface_id"] = cty.StringVal(p.NetworkInterfaceId)
}

func EncodeNetworkInterfaceAttachment_DeviceIndex(p NetworkInterfaceAttachmentParameters, vals map[string]cty.Value) {
	vals["device_index"] = cty.NumberIntVal(p.DeviceIndex)
}

func EncodeNetworkInterfaceAttachment_Status(p NetworkInterfaceAttachmentObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeNetworkInterfaceAttachment_AttachmentId(p NetworkInterfaceAttachmentObservation, vals map[string]cty.Value) {
	vals["attachment_id"] = cty.StringVal(p.AttachmentId)
}