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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane-contrib/terraform-runtime/pkg/plugin"
)

//mergeManagedResourceEntrypointTemplate
type resourceMerger struct{}

func (r *resourceMerger) MergeResources(kube resource.Managed, prov resource.Managed) plugin.MergeDescription {
	k := kube.(*StoragegatewayCachedIscsiVolume)
	p := prov.(*StoragegatewayCachedIscsiVolume)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeStoragegatewayCachedIscsiVolume_TargetName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeStoragegatewayCachedIscsiVolume_NetworkInterfaceId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeStoragegatewayCachedIscsiVolume_KmsEncrypted(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeStoragegatewayCachedIscsiVolume_KmsKey(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeStoragegatewayCachedIscsiVolume_VolumeSizeInBytes(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeStoragegatewayCachedIscsiVolume_GatewayArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeStoragegatewayCachedIscsiVolume_SnapshotId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeStoragegatewayCachedIscsiVolume_SourceVolumeArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeStoragegatewayCachedIscsiVolume_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeStoragegatewayCachedIscsiVolume_LunNumber(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeStoragegatewayCachedIscsiVolume_NetworkInterfacePort(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeStoragegatewayCachedIscsiVolume_VolumeId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeStoragegatewayCachedIscsiVolume_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeStoragegatewayCachedIscsiVolume_VolumeArn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeStoragegatewayCachedIscsiVolume_ChapEnabled(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeStoragegatewayCachedIscsiVolume_TargetArn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	for key, v := range p.Annotations {
		if k.Annotations[key] != v {
			k.Annotations[key] = v
			md.AnnotationsUpdated = true
		}
	}
	md.AnyFieldUpdated = anyChildUpdated
	return *md
}

//mergePrimitiveTemplateSpec
func MergeStoragegatewayCachedIscsiVolume_TargetName(k *StoragegatewayCachedIscsiVolumeParameters, p *StoragegatewayCachedIscsiVolumeParameters, md *plugin.MergeDescription) bool {
	if k.TargetName != p.TargetName {
		p.TargetName = k.TargetName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeStoragegatewayCachedIscsiVolume_NetworkInterfaceId(k *StoragegatewayCachedIscsiVolumeParameters, p *StoragegatewayCachedIscsiVolumeParameters, md *plugin.MergeDescription) bool {
	if k.NetworkInterfaceId != p.NetworkInterfaceId {
		p.NetworkInterfaceId = k.NetworkInterfaceId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeStoragegatewayCachedIscsiVolume_KmsEncrypted(k *StoragegatewayCachedIscsiVolumeParameters, p *StoragegatewayCachedIscsiVolumeParameters, md *plugin.MergeDescription) bool {
	if k.KmsEncrypted != p.KmsEncrypted {
		p.KmsEncrypted = k.KmsEncrypted
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeStoragegatewayCachedIscsiVolume_KmsKey(k *StoragegatewayCachedIscsiVolumeParameters, p *StoragegatewayCachedIscsiVolumeParameters, md *plugin.MergeDescription) bool {
	if k.KmsKey != p.KmsKey {
		p.KmsKey = k.KmsKey
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeStoragegatewayCachedIscsiVolume_VolumeSizeInBytes(k *StoragegatewayCachedIscsiVolumeParameters, p *StoragegatewayCachedIscsiVolumeParameters, md *plugin.MergeDescription) bool {
	if k.VolumeSizeInBytes != p.VolumeSizeInBytes {
		p.VolumeSizeInBytes = k.VolumeSizeInBytes
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeStoragegatewayCachedIscsiVolume_GatewayArn(k *StoragegatewayCachedIscsiVolumeParameters, p *StoragegatewayCachedIscsiVolumeParameters, md *plugin.MergeDescription) bool {
	if k.GatewayArn != p.GatewayArn {
		p.GatewayArn = k.GatewayArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeStoragegatewayCachedIscsiVolume_SnapshotId(k *StoragegatewayCachedIscsiVolumeParameters, p *StoragegatewayCachedIscsiVolumeParameters, md *plugin.MergeDescription) bool {
	if k.SnapshotId != p.SnapshotId {
		p.SnapshotId = k.SnapshotId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeStoragegatewayCachedIscsiVolume_SourceVolumeArn(k *StoragegatewayCachedIscsiVolumeParameters, p *StoragegatewayCachedIscsiVolumeParameters, md *plugin.MergeDescription) bool {
	if k.SourceVolumeArn != p.SourceVolumeArn {
		p.SourceVolumeArn = k.SourceVolumeArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeStoragegatewayCachedIscsiVolume_Tags(k *StoragegatewayCachedIscsiVolumeParameters, p *StoragegatewayCachedIscsiVolumeParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeStoragegatewayCachedIscsiVolume_LunNumber(k *StoragegatewayCachedIscsiVolumeObservation, p *StoragegatewayCachedIscsiVolumeObservation, md *plugin.MergeDescription) bool {
	if k.LunNumber != p.LunNumber {
		k.LunNumber = p.LunNumber
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeStoragegatewayCachedIscsiVolume_NetworkInterfacePort(k *StoragegatewayCachedIscsiVolumeObservation, p *StoragegatewayCachedIscsiVolumeObservation, md *plugin.MergeDescription) bool {
	if k.NetworkInterfacePort != p.NetworkInterfacePort {
		k.NetworkInterfacePort = p.NetworkInterfacePort
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeStoragegatewayCachedIscsiVolume_VolumeId(k *StoragegatewayCachedIscsiVolumeObservation, p *StoragegatewayCachedIscsiVolumeObservation, md *plugin.MergeDescription) bool {
	if k.VolumeId != p.VolumeId {
		k.VolumeId = p.VolumeId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeStoragegatewayCachedIscsiVolume_Arn(k *StoragegatewayCachedIscsiVolumeObservation, p *StoragegatewayCachedIscsiVolumeObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeStoragegatewayCachedIscsiVolume_VolumeArn(k *StoragegatewayCachedIscsiVolumeObservation, p *StoragegatewayCachedIscsiVolumeObservation, md *plugin.MergeDescription) bool {
	if k.VolumeArn != p.VolumeArn {
		k.VolumeArn = p.VolumeArn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeStoragegatewayCachedIscsiVolume_ChapEnabled(k *StoragegatewayCachedIscsiVolumeObservation, p *StoragegatewayCachedIscsiVolumeObservation, md *plugin.MergeDescription) bool {
	if k.ChapEnabled != p.ChapEnabled {
		k.ChapEnabled = p.ChapEnabled
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeStoragegatewayCachedIscsiVolume_TargetArn(k *StoragegatewayCachedIscsiVolumeObservation, p *StoragegatewayCachedIscsiVolumeObservation, md *plugin.MergeDescription) bool {
	if k.TargetArn != p.TargetArn {
		k.TargetArn = p.TargetArn
		md.StatusUpdated = true
		return true
	}
	return false
}