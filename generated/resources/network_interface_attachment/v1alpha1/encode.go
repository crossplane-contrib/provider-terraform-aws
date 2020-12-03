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

package v1alpha1func EncodeNetworkInterfaceAttachment(r NetworkInterfaceAttachment) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeNetworkInterfaceAttachment_InstanceId(r.Spec.ForProvider, ctyVal)
	EncodeNetworkInterfaceAttachment_NetworkInterfaceId(r.Spec.ForProvider, ctyVal)
	EncodeNetworkInterfaceAttachment_DeviceIndex(r.Spec.ForProvider, ctyVal)
	EncodeNetworkInterfaceAttachment_Id(r.Spec.ForProvider, ctyVal)
	EncodeNetworkInterfaceAttachment_Status(r.Status.AtProvider, ctyVal)
	EncodeNetworkInterfaceAttachment_AttachmentId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeNetworkInterfaceAttachment_InstanceId(p *NetworkInterfaceAttachmentParameters, vals map[string]cty.Value) {
	vals["instance_id"] = cty.StringVal(p.InstanceId)
}

func EncodeNetworkInterfaceAttachment_NetworkInterfaceId(p *NetworkInterfaceAttachmentParameters, vals map[string]cty.Value) {
	vals["network_interface_id"] = cty.StringVal(p.NetworkInterfaceId)
}

func EncodeNetworkInterfaceAttachment_DeviceIndex(p *NetworkInterfaceAttachmentParameters, vals map[string]cty.Value) {
	vals["device_index"] = cty.IntVal(p.DeviceIndex)
}

func EncodeNetworkInterfaceAttachment_Id(p *NetworkInterfaceAttachmentParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeNetworkInterfaceAttachment_Status(p *NetworkInterfaceAttachmentObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeNetworkInterfaceAttachment_AttachmentId(p *NetworkInterfaceAttachmentObservation, vals map[string]cty.Value) {
	vals["attachment_id"] = cty.StringVal(p.AttachmentId)
}