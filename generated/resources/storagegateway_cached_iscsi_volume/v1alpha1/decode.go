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

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*StoragegatewayCachedIscsiVolume)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeStoragegatewayCachedIscsiVolume(r, ctyValue)
}

func DecodeStoragegatewayCachedIscsiVolume(prev *StoragegatewayCachedIscsiVolume, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeStoragegatewayCachedIscsiVolume_SourceVolumeArn(&new.Spec.ForProvider, valMap)
	DecodeStoragegatewayCachedIscsiVolume_VolumeSizeInBytes(&new.Spec.ForProvider, valMap)
	DecodeStoragegatewayCachedIscsiVolume_GatewayArn(&new.Spec.ForProvider, valMap)
	DecodeStoragegatewayCachedIscsiVolume_KmsEncrypted(&new.Spec.ForProvider, valMap)
	DecodeStoragegatewayCachedIscsiVolume_TargetName(&new.Spec.ForProvider, valMap)
	DecodeStoragegatewayCachedIscsiVolume_Id(&new.Spec.ForProvider, valMap)
	DecodeStoragegatewayCachedIscsiVolume_NetworkInterfaceId(&new.Spec.ForProvider, valMap)
	DecodeStoragegatewayCachedIscsiVolume_Tags(&new.Spec.ForProvider, valMap)
	DecodeStoragegatewayCachedIscsiVolume_KmsKey(&new.Spec.ForProvider, valMap)
	DecodeStoragegatewayCachedIscsiVolume_SnapshotId(&new.Spec.ForProvider, valMap)
	DecodeStoragegatewayCachedIscsiVolume_Arn(&new.Status.AtProvider, valMap)
	DecodeStoragegatewayCachedIscsiVolume_ChapEnabled(&new.Status.AtProvider, valMap)
	DecodeStoragegatewayCachedIscsiVolume_LunNumber(&new.Status.AtProvider, valMap)
	DecodeStoragegatewayCachedIscsiVolume_VolumeId(&new.Status.AtProvider, valMap)
	DecodeStoragegatewayCachedIscsiVolume_NetworkInterfacePort(&new.Status.AtProvider, valMap)
	DecodeStoragegatewayCachedIscsiVolume_VolumeArn(&new.Status.AtProvider, valMap)
	DecodeStoragegatewayCachedIscsiVolume_TargetArn(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeStoragegatewayCachedIscsiVolume_SourceVolumeArn(p *StoragegatewayCachedIscsiVolumeParameters, vals map[string]cty.Value) {
	p.SourceVolumeArn = ctwhy.ValueAsString(vals["source_volume_arn"])
}

func DecodeStoragegatewayCachedIscsiVolume_VolumeSizeInBytes(p *StoragegatewayCachedIscsiVolumeParameters, vals map[string]cty.Value) {
	p.VolumeSizeInBytes = ctwhy.ValueAsInt64(vals["volume_size_in_bytes"])
}

func DecodeStoragegatewayCachedIscsiVolume_GatewayArn(p *StoragegatewayCachedIscsiVolumeParameters, vals map[string]cty.Value) {
	p.GatewayArn = ctwhy.ValueAsString(vals["gateway_arn"])
}

func DecodeStoragegatewayCachedIscsiVolume_KmsEncrypted(p *StoragegatewayCachedIscsiVolumeParameters, vals map[string]cty.Value) {
	p.KmsEncrypted = ctwhy.ValueAsBool(vals["kms_encrypted"])
}

func DecodeStoragegatewayCachedIscsiVolume_TargetName(p *StoragegatewayCachedIscsiVolumeParameters, vals map[string]cty.Value) {
	p.TargetName = ctwhy.ValueAsString(vals["target_name"])
}

func DecodeStoragegatewayCachedIscsiVolume_Id(p *StoragegatewayCachedIscsiVolumeParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeStoragegatewayCachedIscsiVolume_NetworkInterfaceId(p *StoragegatewayCachedIscsiVolumeParameters, vals map[string]cty.Value) {
	p.NetworkInterfaceId = ctwhy.ValueAsString(vals["network_interface_id"])
}

func DecodeStoragegatewayCachedIscsiVolume_Tags(p *StoragegatewayCachedIscsiVolumeParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeStoragegatewayCachedIscsiVolume_KmsKey(p *StoragegatewayCachedIscsiVolumeParameters, vals map[string]cty.Value) {
	p.KmsKey = ctwhy.ValueAsString(vals["kms_key"])
}

func DecodeStoragegatewayCachedIscsiVolume_SnapshotId(p *StoragegatewayCachedIscsiVolumeParameters, vals map[string]cty.Value) {
	p.SnapshotId = ctwhy.ValueAsString(vals["snapshot_id"])
}

func DecodeStoragegatewayCachedIscsiVolume_Arn(p *StoragegatewayCachedIscsiVolumeObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

func DecodeStoragegatewayCachedIscsiVolume_ChapEnabled(p *StoragegatewayCachedIscsiVolumeObservation, vals map[string]cty.Value) {
	p.ChapEnabled = ctwhy.ValueAsBool(vals["chap_enabled"])
}

func DecodeStoragegatewayCachedIscsiVolume_LunNumber(p *StoragegatewayCachedIscsiVolumeObservation, vals map[string]cty.Value) {
	p.LunNumber = ctwhy.ValueAsInt64(vals["lun_number"])
}

func DecodeStoragegatewayCachedIscsiVolume_VolumeId(p *StoragegatewayCachedIscsiVolumeObservation, vals map[string]cty.Value) {
	p.VolumeId = ctwhy.ValueAsString(vals["volume_id"])
}

func DecodeStoragegatewayCachedIscsiVolume_NetworkInterfacePort(p *StoragegatewayCachedIscsiVolumeObservation, vals map[string]cty.Value) {
	p.NetworkInterfacePort = ctwhy.ValueAsInt64(vals["network_interface_port"])
}

func DecodeStoragegatewayCachedIscsiVolume_VolumeArn(p *StoragegatewayCachedIscsiVolumeObservation, vals map[string]cty.Value) {
	p.VolumeArn = ctwhy.ValueAsString(vals["volume_arn"])
}

func DecodeStoragegatewayCachedIscsiVolume_TargetArn(p *StoragegatewayCachedIscsiVolumeObservation, vals map[string]cty.Value) {
	p.TargetArn = ctwhy.ValueAsString(vals["target_arn"])
}