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

func EncodeAmiFromInstance(r AmiFromInstance) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeAmiFromInstance_SnapshotWithoutReboot(r.Spec.ForProvider, ctyVal)
	EncodeAmiFromInstance_Tags(r.Spec.ForProvider, ctyVal)
	EncodeAmiFromInstance_Id(r.Spec.ForProvider, ctyVal)
	EncodeAmiFromInstance_Name(r.Spec.ForProvider, ctyVal)
	EncodeAmiFromInstance_SourceInstanceId(r.Spec.ForProvider, ctyVal)
	EncodeAmiFromInstance_Description(r.Spec.ForProvider, ctyVal)
	EncodeAmiFromInstance_EbsBlockDevice(r.Spec.ForProvider.EbsBlockDevice, ctyVal)
	EncodeAmiFromInstance_EphemeralBlockDevice(r.Spec.ForProvider.EphemeralBlockDevice, ctyVal)
	EncodeAmiFromInstance_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeAmiFromInstance_Arn(r.Status.AtProvider, ctyVal)
	EncodeAmiFromInstance_ImageLocation(r.Status.AtProvider, ctyVal)
	EncodeAmiFromInstance_ManageEbsSnapshots(r.Status.AtProvider, ctyVal)
	EncodeAmiFromInstance_RamdiskId(r.Status.AtProvider, ctyVal)
	EncodeAmiFromInstance_RootSnapshotId(r.Status.AtProvider, ctyVal)
	EncodeAmiFromInstance_KernelId(r.Status.AtProvider, ctyVal)
	EncodeAmiFromInstance_Architecture(r.Status.AtProvider, ctyVal)
	EncodeAmiFromInstance_EnaSupport(r.Status.AtProvider, ctyVal)
	EncodeAmiFromInstance_RootDeviceName(r.Status.AtProvider, ctyVal)
	EncodeAmiFromInstance_SriovNetSupport(r.Status.AtProvider, ctyVal)
	EncodeAmiFromInstance_VirtualizationType(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeAmiFromInstance_SnapshotWithoutReboot(p AmiFromInstanceParameters, vals map[string]cty.Value) {
	vals["snapshot_without_reboot"] = cty.BoolVal(p.SnapshotWithoutReboot)
}

func EncodeAmiFromInstance_Tags(p AmiFromInstanceParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeAmiFromInstance_Id(p AmiFromInstanceParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAmiFromInstance_Name(p AmiFromInstanceParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAmiFromInstance_SourceInstanceId(p AmiFromInstanceParameters, vals map[string]cty.Value) {
	vals["source_instance_id"] = cty.StringVal(p.SourceInstanceId)
}

func EncodeAmiFromInstance_Description(p AmiFromInstanceParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeAmiFromInstance_EbsBlockDevice(p EbsBlockDevice, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAmiFromInstance_EbsBlockDevice_DeleteOnTermination(p, ctyVal)
	EncodeAmiFromInstance_EbsBlockDevice_DeviceName(p, ctyVal)
	EncodeAmiFromInstance_EbsBlockDevice_Encrypted(p, ctyVal)
	EncodeAmiFromInstance_EbsBlockDevice_Iops(p, ctyVal)
	EncodeAmiFromInstance_EbsBlockDevice_SnapshotId(p, ctyVal)
	EncodeAmiFromInstance_EbsBlockDevice_VolumeSize(p, ctyVal)
	EncodeAmiFromInstance_EbsBlockDevice_VolumeType(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ebs_block_device"] = cty.SetVal(valsForCollection)
}

func EncodeAmiFromInstance_EbsBlockDevice_DeleteOnTermination(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["delete_on_termination"] = cty.BoolVal(p.DeleteOnTermination)
}

func EncodeAmiFromInstance_EbsBlockDevice_DeviceName(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["device_name"] = cty.StringVal(p.DeviceName)
}

func EncodeAmiFromInstance_EbsBlockDevice_Encrypted(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeAmiFromInstance_EbsBlockDevice_Iops(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["iops"] = cty.NumberIntVal(p.Iops)
}

func EncodeAmiFromInstance_EbsBlockDevice_SnapshotId(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["snapshot_id"] = cty.StringVal(p.SnapshotId)
}

func EncodeAmiFromInstance_EbsBlockDevice_VolumeSize(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["volume_size"] = cty.NumberIntVal(p.VolumeSize)
}

func EncodeAmiFromInstance_EbsBlockDevice_VolumeType(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["volume_type"] = cty.StringVal(p.VolumeType)
}

func EncodeAmiFromInstance_EphemeralBlockDevice(p EphemeralBlockDevice, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAmiFromInstance_EphemeralBlockDevice_DeviceName(p, ctyVal)
	EncodeAmiFromInstance_EphemeralBlockDevice_VirtualName(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ephemeral_block_device"] = cty.SetVal(valsForCollection)
}

func EncodeAmiFromInstance_EphemeralBlockDevice_DeviceName(p EphemeralBlockDevice, vals map[string]cty.Value) {
	vals["device_name"] = cty.StringVal(p.DeviceName)
}

func EncodeAmiFromInstance_EphemeralBlockDevice_VirtualName(p EphemeralBlockDevice, vals map[string]cty.Value) {
	vals["virtual_name"] = cty.StringVal(p.VirtualName)
}

func EncodeAmiFromInstance_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeAmiFromInstance_Timeouts_Create(p, ctyVal)
	EncodeAmiFromInstance_Timeouts_Delete(p, ctyVal)
	EncodeAmiFromInstance_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeAmiFromInstance_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeAmiFromInstance_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeAmiFromInstance_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeAmiFromInstance_Arn(p AmiFromInstanceObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeAmiFromInstance_ImageLocation(p AmiFromInstanceObservation, vals map[string]cty.Value) {
	vals["image_location"] = cty.StringVal(p.ImageLocation)
}

func EncodeAmiFromInstance_ManageEbsSnapshots(p AmiFromInstanceObservation, vals map[string]cty.Value) {
	vals["manage_ebs_snapshots"] = cty.BoolVal(p.ManageEbsSnapshots)
}

func EncodeAmiFromInstance_RamdiskId(p AmiFromInstanceObservation, vals map[string]cty.Value) {
	vals["ramdisk_id"] = cty.StringVal(p.RamdiskId)
}

func EncodeAmiFromInstance_RootSnapshotId(p AmiFromInstanceObservation, vals map[string]cty.Value) {
	vals["root_snapshot_id"] = cty.StringVal(p.RootSnapshotId)
}

func EncodeAmiFromInstance_KernelId(p AmiFromInstanceObservation, vals map[string]cty.Value) {
	vals["kernel_id"] = cty.StringVal(p.KernelId)
}

func EncodeAmiFromInstance_Architecture(p AmiFromInstanceObservation, vals map[string]cty.Value) {
	vals["architecture"] = cty.StringVal(p.Architecture)
}

func EncodeAmiFromInstance_EnaSupport(p AmiFromInstanceObservation, vals map[string]cty.Value) {
	vals["ena_support"] = cty.BoolVal(p.EnaSupport)
}

func EncodeAmiFromInstance_RootDeviceName(p AmiFromInstanceObservation, vals map[string]cty.Value) {
	vals["root_device_name"] = cty.StringVal(p.RootDeviceName)
}

func EncodeAmiFromInstance_SriovNetSupport(p AmiFromInstanceObservation, vals map[string]cty.Value) {
	vals["sriov_net_support"] = cty.StringVal(p.SriovNetSupport)
}

func EncodeAmiFromInstance_VirtualizationType(p AmiFromInstanceObservation, vals map[string]cty.Value) {
	vals["virtualization_type"] = cty.StringVal(p.VirtualizationType)
}