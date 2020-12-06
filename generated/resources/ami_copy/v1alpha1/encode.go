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

func EncodeAmiCopy(r AmiCopy) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeAmiCopy_Description(r.Spec.ForProvider, ctyVal)
	EncodeAmiCopy_SourceAmiId(r.Spec.ForProvider, ctyVal)
	EncodeAmiCopy_SourceAmiRegion(r.Spec.ForProvider, ctyVal)
	EncodeAmiCopy_KmsKeyId(r.Spec.ForProvider, ctyVal)
	EncodeAmiCopy_Tags(r.Spec.ForProvider, ctyVal)
	EncodeAmiCopy_Id(r.Spec.ForProvider, ctyVal)
	EncodeAmiCopy_Name(r.Spec.ForProvider, ctyVal)
	EncodeAmiCopy_Encrypted(r.Spec.ForProvider, ctyVal)
	EncodeAmiCopy_EbsBlockDevice(r.Spec.ForProvider.EbsBlockDevice, ctyVal)
	EncodeAmiCopy_EphemeralBlockDevice(r.Spec.ForProvider.EphemeralBlockDevice, ctyVal)
	EncodeAmiCopy_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeAmiCopy_VirtualizationType(r.Status.AtProvider, ctyVal)
	EncodeAmiCopy_EnaSupport(r.Status.AtProvider, ctyVal)
	EncodeAmiCopy_KernelId(r.Status.AtProvider, ctyVal)
	EncodeAmiCopy_RootDeviceName(r.Status.AtProvider, ctyVal)
	EncodeAmiCopy_ManageEbsSnapshots(r.Status.AtProvider, ctyVal)
	EncodeAmiCopy_Architecture(r.Status.AtProvider, ctyVal)
	EncodeAmiCopy_ImageLocation(r.Status.AtProvider, ctyVal)
	EncodeAmiCopy_RootSnapshotId(r.Status.AtProvider, ctyVal)
	EncodeAmiCopy_Arn(r.Status.AtProvider, ctyVal)
	EncodeAmiCopy_RamdiskId(r.Status.AtProvider, ctyVal)
	EncodeAmiCopy_SriovNetSupport(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeAmiCopy_Description(p AmiCopyParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeAmiCopy_SourceAmiId(p AmiCopyParameters, vals map[string]cty.Value) {
	vals["source_ami_id"] = cty.StringVal(p.SourceAmiId)
}

func EncodeAmiCopy_SourceAmiRegion(p AmiCopyParameters, vals map[string]cty.Value) {
	vals["source_ami_region"] = cty.StringVal(p.SourceAmiRegion)
}

func EncodeAmiCopy_KmsKeyId(p AmiCopyParameters, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeAmiCopy_Tags(p AmiCopyParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeAmiCopy_Id(p AmiCopyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAmiCopy_Name(p AmiCopyParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAmiCopy_Encrypted(p AmiCopyParameters, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeAmiCopy_EbsBlockDevice(p EbsBlockDevice, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAmiCopy_EbsBlockDevice_SnapshotId(p, ctyVal)
	EncodeAmiCopy_EbsBlockDevice_VolumeSize(p, ctyVal)
	EncodeAmiCopy_EbsBlockDevice_VolumeType(p, ctyVal)
	EncodeAmiCopy_EbsBlockDevice_DeleteOnTermination(p, ctyVal)
	EncodeAmiCopy_EbsBlockDevice_DeviceName(p, ctyVal)
	EncodeAmiCopy_EbsBlockDevice_Encrypted(p, ctyVal)
	EncodeAmiCopy_EbsBlockDevice_Iops(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ebs_block_device"] = cty.SetVal(valsForCollection)
}

func EncodeAmiCopy_EbsBlockDevice_SnapshotId(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["snapshot_id"] = cty.StringVal(p.SnapshotId)
}

func EncodeAmiCopy_EbsBlockDevice_VolumeSize(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["volume_size"] = cty.NumberIntVal(p.VolumeSize)
}

func EncodeAmiCopy_EbsBlockDevice_VolumeType(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["volume_type"] = cty.StringVal(p.VolumeType)
}

func EncodeAmiCopy_EbsBlockDevice_DeleteOnTermination(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["delete_on_termination"] = cty.BoolVal(p.DeleteOnTermination)
}

func EncodeAmiCopy_EbsBlockDevice_DeviceName(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["device_name"] = cty.StringVal(p.DeviceName)
}

func EncodeAmiCopy_EbsBlockDevice_Encrypted(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeAmiCopy_EbsBlockDevice_Iops(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["iops"] = cty.NumberIntVal(p.Iops)
}

func EncodeAmiCopy_EphemeralBlockDevice(p EphemeralBlockDevice, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAmiCopy_EphemeralBlockDevice_DeviceName(p, ctyVal)
	EncodeAmiCopy_EphemeralBlockDevice_VirtualName(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ephemeral_block_device"] = cty.SetVal(valsForCollection)
}

func EncodeAmiCopy_EphemeralBlockDevice_DeviceName(p EphemeralBlockDevice, vals map[string]cty.Value) {
	vals["device_name"] = cty.StringVal(p.DeviceName)
}

func EncodeAmiCopy_EphemeralBlockDevice_VirtualName(p EphemeralBlockDevice, vals map[string]cty.Value) {
	vals["virtual_name"] = cty.StringVal(p.VirtualName)
}

func EncodeAmiCopy_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeAmiCopy_Timeouts_Delete(p, ctyVal)
	EncodeAmiCopy_Timeouts_Update(p, ctyVal)
	EncodeAmiCopy_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeAmiCopy_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeAmiCopy_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeAmiCopy_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeAmiCopy_VirtualizationType(p AmiCopyObservation, vals map[string]cty.Value) {
	vals["virtualization_type"] = cty.StringVal(p.VirtualizationType)
}

func EncodeAmiCopy_EnaSupport(p AmiCopyObservation, vals map[string]cty.Value) {
	vals["ena_support"] = cty.BoolVal(p.EnaSupport)
}

func EncodeAmiCopy_KernelId(p AmiCopyObservation, vals map[string]cty.Value) {
	vals["kernel_id"] = cty.StringVal(p.KernelId)
}

func EncodeAmiCopy_RootDeviceName(p AmiCopyObservation, vals map[string]cty.Value) {
	vals["root_device_name"] = cty.StringVal(p.RootDeviceName)
}

func EncodeAmiCopy_ManageEbsSnapshots(p AmiCopyObservation, vals map[string]cty.Value) {
	vals["manage_ebs_snapshots"] = cty.BoolVal(p.ManageEbsSnapshots)
}

func EncodeAmiCopy_Architecture(p AmiCopyObservation, vals map[string]cty.Value) {
	vals["architecture"] = cty.StringVal(p.Architecture)
}

func EncodeAmiCopy_ImageLocation(p AmiCopyObservation, vals map[string]cty.Value) {
	vals["image_location"] = cty.StringVal(p.ImageLocation)
}

func EncodeAmiCopy_RootSnapshotId(p AmiCopyObservation, vals map[string]cty.Value) {
	vals["root_snapshot_id"] = cty.StringVal(p.RootSnapshotId)
}

func EncodeAmiCopy_Arn(p AmiCopyObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeAmiCopy_RamdiskId(p AmiCopyObservation, vals map[string]cty.Value) {
	vals["ramdisk_id"] = cty.StringVal(p.RamdiskId)
}

func EncodeAmiCopy_SriovNetSupport(p AmiCopyObservation, vals map[string]cty.Value) {
	vals["sriov_net_support"] = cty.StringVal(p.SriovNetSupport)
}