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

package v1alpha1func EncodeNetworkInterface(r NetworkInterface) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeNetworkInterface_PrivateIp(r.Spec.ForProvider, ctyVal)
	EncodeNetworkInterface_PrivateIps(r.Spec.ForProvider, ctyVal)
	EncodeNetworkInterface_SecurityGroups(r.Spec.ForProvider, ctyVal)
	EncodeNetworkInterface_SourceDestCheck(r.Spec.ForProvider, ctyVal)
	EncodeNetworkInterface_SubnetId(r.Spec.ForProvider, ctyVal)
	EncodeNetworkInterface_Tags(r.Spec.ForProvider, ctyVal)
	EncodeNetworkInterface_Id(r.Spec.ForProvider, ctyVal)
	EncodeNetworkInterface_PrivateIpsCount(r.Spec.ForProvider, ctyVal)
	EncodeNetworkInterface_Description(r.Spec.ForProvider, ctyVal)
	EncodeNetworkInterface_Attachment(r.Spec.ForProvider.Attachment, ctyVal)
	EncodeNetworkInterface_PrivateDnsName(r.Status.AtProvider, ctyVal)
	EncodeNetworkInterface_MacAddress(r.Status.AtProvider, ctyVal)
	EncodeNetworkInterface_OutpostArn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeNetworkInterface_PrivateIp(p *NetworkInterfaceParameters, vals map[string]cty.Value) {
	vals["private_ip"] = cty.StringVal(p.PrivateIp)
}

func EncodeNetworkInterface_PrivateIps(p *NetworkInterfaceParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.PrivateIps {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["private_ips"] = cty.SetVal(colVals)
}

func EncodeNetworkInterface_SecurityGroups(p *NetworkInterfaceParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroups {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_groups"] = cty.SetVal(colVals)
}

func EncodeNetworkInterface_SourceDestCheck(p *NetworkInterfaceParameters, vals map[string]cty.Value) {
	vals["source_dest_check"] = cty.BoolVal(p.SourceDestCheck)
}

func EncodeNetworkInterface_SubnetId(p *NetworkInterfaceParameters, vals map[string]cty.Value) {
	vals["subnet_id"] = cty.StringVal(p.SubnetId)
}

func EncodeNetworkInterface_Tags(p *NetworkInterfaceParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeNetworkInterface_Id(p *NetworkInterfaceParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeNetworkInterface_PrivateIpsCount(p *NetworkInterfaceParameters, vals map[string]cty.Value) {
	vals["private_ips_count"] = cty.IntVal(p.PrivateIpsCount)
}

func EncodeNetworkInterface_Description(p *NetworkInterfaceParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeNetworkInterface_Attachment(p *Attachment, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Attachment {
		ctyVal = make(map[string]cty.Value)
		EncodeNetworkInterface_Attachment_AttachmentId(v, ctyVal)
		EncodeNetworkInterface_Attachment_DeviceIndex(v, ctyVal)
		EncodeNetworkInterface_Attachment_Instance(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["attachment"] = cty.SetVal(valsForCollection)
}

func EncodeNetworkInterface_Attachment_AttachmentId(p *Attachment, vals map[string]cty.Value) {
	vals["attachment_id"] = cty.StringVal(p.AttachmentId)
}

func EncodeNetworkInterface_Attachment_DeviceIndex(p *Attachment, vals map[string]cty.Value) {
	vals["device_index"] = cty.IntVal(p.DeviceIndex)
}

func EncodeNetworkInterface_Attachment_Instance(p *Attachment, vals map[string]cty.Value) {
	vals["instance"] = cty.StringVal(p.Instance)
}

func EncodeNetworkInterface_PrivateDnsName(p *NetworkInterfaceObservation, vals map[string]cty.Value) {
	vals["private_dns_name"] = cty.StringVal(p.PrivateDnsName)
}

func EncodeNetworkInterface_MacAddress(p *NetworkInterfaceObservation, vals map[string]cty.Value) {
	vals["mac_address"] = cty.StringVal(p.MacAddress)
}

func EncodeNetworkInterface_OutpostArn(p *NetworkInterfaceObservation, vals map[string]cty.Value) {
	vals["outpost_arn"] = cty.StringVal(p.OutpostArn)
}