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

package v1alpha1func EncodeAmi(r Ami) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeAmi_RootDeviceName(r.Spec.ForProvider, ctyVal)
	EncodeAmi_Tags(r.Spec.ForProvider, ctyVal)
	EncodeAmi_KernelId(r.Spec.ForProvider, ctyVal)
	EncodeAmi_EnaSupport(r.Spec.ForProvider, ctyVal)
	EncodeAmi_Id(r.Spec.ForProvider, ctyVal)
	EncodeAmi_Architecture(r.Spec.ForProvider, ctyVal)
	EncodeAmi_ImageLocation(r.Spec.ForProvider, ctyVal)
	EncodeAmi_Name(r.Spec.ForProvider, ctyVal)
	EncodeAmi_RamdiskId(r.Spec.ForProvider, ctyVal)
	EncodeAmi_SriovNetSupport(r.Spec.ForProvider, ctyVal)
	EncodeAmi_VirtualizationType(r.Spec.ForProvider, ctyVal)
	EncodeAmi_Description(r.Spec.ForProvider, ctyVal)
	EncodeAmi_EbsBlockDevice(r.Spec.ForProvider.EbsBlockDevice, ctyVal)
	EncodeAmi_EphemeralBlockDevice(r.Spec.ForProvider.EphemeralBlockDevice, ctyVal)
	EncodeAmi_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeAmi_ManageEbsSnapshots(r.Status.AtProvider, ctyVal)
	EncodeAmi_RootSnapshotId(r.Status.AtProvider, ctyVal)
	EncodeAmi_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeAmi_RootDeviceName(p *AmiParameters, vals map[string]cty.Value) {
	vals["root_device_name"] = cty.StringVal(p.RootDeviceName)
}

func EncodeAmi_Tags(p *AmiParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeAmi_KernelId(p *AmiParameters, vals map[string]cty.Value) {
	vals["kernel_id"] = cty.StringVal(p.KernelId)
}

func EncodeAmi_EnaSupport(p *AmiParameters, vals map[string]cty.Value) {
	vals["ena_support"] = cty.BoolVal(p.EnaSupport)
}

func EncodeAmi_Id(p *AmiParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAmi_Architecture(p *AmiParameters, vals map[string]cty.Value) {
	vals["architecture"] = cty.StringVal(p.Architecture)
}

func EncodeAmi_ImageLocation(p *AmiParameters, vals map[string]cty.Value) {
	vals["image_location"] = cty.StringVal(p.ImageLocation)
}

func EncodeAmi_Name(p *AmiParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAmi_RamdiskId(p *AmiParameters, vals map[string]cty.Value) {
	vals["ramdisk_id"] = cty.StringVal(p.RamdiskId)
}

func EncodeAmi_SriovNetSupport(p *AmiParameters, vals map[string]cty.Value) {
	vals["sriov_net_support"] = cty.StringVal(p.SriovNetSupport)
}

func EncodeAmi_VirtualizationType(p *AmiParameters, vals map[string]cty.Value) {
	vals["virtualization_type"] = cty.StringVal(p.VirtualizationType)
}

func EncodeAmi_Description(p *AmiParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeAmi_EbsBlockDevice(p *EbsBlockDevice, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.EbsBlockDevice {
		ctyVal = make(map[string]cty.Value)
		EncodeAmi_EbsBlockDevice_DeleteOnTermination(v, ctyVal)
		EncodeAmi_EbsBlockDevice_DeviceName(v, ctyVal)
		EncodeAmi_EbsBlockDevice_Encrypted(v, ctyVal)
		EncodeAmi_EbsBlockDevice_Iops(v, ctyVal)
		EncodeAmi_EbsBlockDevice_SnapshotId(v, ctyVal)
		EncodeAmi_EbsBlockDevice_VolumeSize(v, ctyVal)
		EncodeAmi_EbsBlockDevice_VolumeType(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ebs_block_device"] = cty.SetVal(valsForCollection)
}

func EncodeAmi_EbsBlockDevice_DeleteOnTermination(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["delete_on_termination"] = cty.BoolVal(p.DeleteOnTermination)
}

func EncodeAmi_EbsBlockDevice_DeviceName(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["device_name"] = cty.StringVal(p.DeviceName)
}

func EncodeAmi_EbsBlockDevice_Encrypted(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeAmi_EbsBlockDevice_Iops(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["iops"] = cty.IntVal(p.Iops)
}

func EncodeAmi_EbsBlockDevice_SnapshotId(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["snapshot_id"] = cty.StringVal(p.SnapshotId)
}

func EncodeAmi_EbsBlockDevice_VolumeSize(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["volume_size"] = cty.IntVal(p.VolumeSize)
}

func EncodeAmi_EbsBlockDevice_VolumeType(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["volume_type"] = cty.StringVal(p.VolumeType)
}

func EncodeAmi_EphemeralBlockDevice(p *EphemeralBlockDevice, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.EphemeralBlockDevice {
		ctyVal = make(map[string]cty.Value)
		EncodeAmi_EphemeralBlockDevice_DeviceName(v, ctyVal)
		EncodeAmi_EphemeralBlockDevice_VirtualName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ephemeral_block_device"] = cty.SetVal(valsForCollection)
}

func EncodeAmi_EphemeralBlockDevice_DeviceName(p *EphemeralBlockDevice, vals map[string]cty.Value) {
	vals["device_name"] = cty.StringVal(p.DeviceName)
}

func EncodeAmi_EphemeralBlockDevice_VirtualName(p *EphemeralBlockDevice, vals map[string]cty.Value) {
	vals["virtual_name"] = cty.StringVal(p.VirtualName)
}

func EncodeAmi_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	ctyVal = make(map[string]cty.Value)
	EncodeAmi_Timeouts_Create(p, ctyVal)
	EncodeAmi_Timeouts_Delete(p, ctyVal)
	EncodeAmi_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeAmi_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeAmi_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeAmi_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeAmi_ManageEbsSnapshots(p *AmiObservation, vals map[string]cty.Value) {
	vals["manage_ebs_snapshots"] = cty.BoolVal(p.ManageEbsSnapshots)
}

func EncodeAmi_RootSnapshotId(p *AmiObservation, vals map[string]cty.Value) {
	vals["root_snapshot_id"] = cty.StringVal(p.RootSnapshotId)
}

func EncodeAmi_Arn(p *AmiObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}