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

func EncodeStoragegatewayCachedIscsiVolume(r StoragegatewayCachedIscsiVolume) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeStoragegatewayCachedIscsiVolume_GatewayArn(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayCachedIscsiVolume_KmsEncrypted(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayCachedIscsiVolume_Id(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayCachedIscsiVolume_NetworkInterfaceId(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayCachedIscsiVolume_SnapshotId(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayCachedIscsiVolume_Tags(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayCachedIscsiVolume_TargetName(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayCachedIscsiVolume_KmsKey(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayCachedIscsiVolume_SourceVolumeArn(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayCachedIscsiVolume_VolumeSizeInBytes(r.Spec.ForProvider, ctyVal)
	EncodeStoragegatewayCachedIscsiVolume_Arn(r.Status.AtProvider, ctyVal)
	EncodeStoragegatewayCachedIscsiVolume_NetworkInterfacePort(r.Status.AtProvider, ctyVal)
	EncodeStoragegatewayCachedIscsiVolume_VolumeId(r.Status.AtProvider, ctyVal)
	EncodeStoragegatewayCachedIscsiVolume_ChapEnabled(r.Status.AtProvider, ctyVal)
	EncodeStoragegatewayCachedIscsiVolume_LunNumber(r.Status.AtProvider, ctyVal)
	EncodeStoragegatewayCachedIscsiVolume_TargetArn(r.Status.AtProvider, ctyVal)
	EncodeStoragegatewayCachedIscsiVolume_VolumeArn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeStoragegatewayCachedIscsiVolume_GatewayArn(p StoragegatewayCachedIscsiVolumeParameters, vals map[string]cty.Value) {
	vals["gateway_arn"] = cty.StringVal(p.GatewayArn)
}

func EncodeStoragegatewayCachedIscsiVolume_KmsEncrypted(p StoragegatewayCachedIscsiVolumeParameters, vals map[string]cty.Value) {
	vals["kms_encrypted"] = cty.BoolVal(p.KmsEncrypted)
}

func EncodeStoragegatewayCachedIscsiVolume_Id(p StoragegatewayCachedIscsiVolumeParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeStoragegatewayCachedIscsiVolume_NetworkInterfaceId(p StoragegatewayCachedIscsiVolumeParameters, vals map[string]cty.Value) {
	vals["network_interface_id"] = cty.StringVal(p.NetworkInterfaceId)
}

func EncodeStoragegatewayCachedIscsiVolume_SnapshotId(p StoragegatewayCachedIscsiVolumeParameters, vals map[string]cty.Value) {
	vals["snapshot_id"] = cty.StringVal(p.SnapshotId)
}

func EncodeStoragegatewayCachedIscsiVolume_Tags(p StoragegatewayCachedIscsiVolumeParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeStoragegatewayCachedIscsiVolume_TargetName(p StoragegatewayCachedIscsiVolumeParameters, vals map[string]cty.Value) {
	vals["target_name"] = cty.StringVal(p.TargetName)
}

func EncodeStoragegatewayCachedIscsiVolume_KmsKey(p StoragegatewayCachedIscsiVolumeParameters, vals map[string]cty.Value) {
	vals["kms_key"] = cty.StringVal(p.KmsKey)
}

func EncodeStoragegatewayCachedIscsiVolume_SourceVolumeArn(p StoragegatewayCachedIscsiVolumeParameters, vals map[string]cty.Value) {
	vals["source_volume_arn"] = cty.StringVal(p.SourceVolumeArn)
}

func EncodeStoragegatewayCachedIscsiVolume_VolumeSizeInBytes(p StoragegatewayCachedIscsiVolumeParameters, vals map[string]cty.Value) {
	vals["volume_size_in_bytes"] = cty.NumberIntVal(p.VolumeSizeInBytes)
}

func EncodeStoragegatewayCachedIscsiVolume_Arn(p StoragegatewayCachedIscsiVolumeObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeStoragegatewayCachedIscsiVolume_NetworkInterfacePort(p StoragegatewayCachedIscsiVolumeObservation, vals map[string]cty.Value) {
	vals["network_interface_port"] = cty.NumberIntVal(p.NetworkInterfacePort)
}

func EncodeStoragegatewayCachedIscsiVolume_VolumeId(p StoragegatewayCachedIscsiVolumeObservation, vals map[string]cty.Value) {
	vals["volume_id"] = cty.StringVal(p.VolumeId)
}

func EncodeStoragegatewayCachedIscsiVolume_ChapEnabled(p StoragegatewayCachedIscsiVolumeObservation, vals map[string]cty.Value) {
	vals["chap_enabled"] = cty.BoolVal(p.ChapEnabled)
}

func EncodeStoragegatewayCachedIscsiVolume_LunNumber(p StoragegatewayCachedIscsiVolumeObservation, vals map[string]cty.Value) {
	vals["lun_number"] = cty.NumberIntVal(p.LunNumber)
}

func EncodeStoragegatewayCachedIscsiVolume_TargetArn(p StoragegatewayCachedIscsiVolumeObservation, vals map[string]cty.Value) {
	vals["target_arn"] = cty.StringVal(p.TargetArn)
}

func EncodeStoragegatewayCachedIscsiVolume_VolumeArn(p StoragegatewayCachedIscsiVolumeObservation, vals map[string]cty.Value) {
	vals["volume_arn"] = cty.StringVal(p.VolumeArn)
}