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
	k := kube.(*EbsSnapshotCopy)
	p := prov.(*EbsSnapshotCopy)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeEbsSnapshotCopy_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEbsSnapshotCopy_Encrypted(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEbsSnapshotCopy_SourceRegion(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEbsSnapshotCopy_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEbsSnapshotCopy_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEbsSnapshotCopy_KmsKeyId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEbsSnapshotCopy_SourceSnapshotId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEbsSnapshotCopy_VolumeId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEbsSnapshotCopy_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEbsSnapshotCopy_OwnerAlias(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEbsSnapshotCopy_OwnerId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEbsSnapshotCopy_VolumeSize(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEbsSnapshotCopy_DataEncryptionKeyId(&k.Status.AtProvider, &p.Status.AtProvider, md)
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

//mergePrimitiveContainerTemplateSpec
func MergeEbsSnapshotCopy_Tags(k *EbsSnapshotCopyParameters, p *EbsSnapshotCopyParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(p.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEbsSnapshotCopy_Encrypted(k *EbsSnapshotCopyParameters, p *EbsSnapshotCopyParameters, md *plugin.MergeDescription) bool {
	if k.Encrypted != p.Encrypted {
		p.Encrypted = k.Encrypted
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEbsSnapshotCopy_SourceRegion(k *EbsSnapshotCopyParameters, p *EbsSnapshotCopyParameters, md *plugin.MergeDescription) bool {
	if k.SourceRegion != p.SourceRegion {
		p.SourceRegion = k.SourceRegion
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEbsSnapshotCopy_Description(k *EbsSnapshotCopyParameters, p *EbsSnapshotCopyParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEbsSnapshotCopy_Id(k *EbsSnapshotCopyParameters, p *EbsSnapshotCopyParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEbsSnapshotCopy_KmsKeyId(k *EbsSnapshotCopyParameters, p *EbsSnapshotCopyParameters, md *plugin.MergeDescription) bool {
	if k.KmsKeyId != p.KmsKeyId {
		p.KmsKeyId = k.KmsKeyId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEbsSnapshotCopy_SourceSnapshotId(k *EbsSnapshotCopyParameters, p *EbsSnapshotCopyParameters, md *plugin.MergeDescription) bool {
	if k.SourceSnapshotId != p.SourceSnapshotId {
		p.SourceSnapshotId = k.SourceSnapshotId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEbsSnapshotCopy_VolumeId(k *EbsSnapshotCopyObservation, p *EbsSnapshotCopyObservation, md *plugin.MergeDescription) bool {
	if k.VolumeId != p.VolumeId {
		k.VolumeId = p.VolumeId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEbsSnapshotCopy_Arn(k *EbsSnapshotCopyObservation, p *EbsSnapshotCopyObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEbsSnapshotCopy_OwnerAlias(k *EbsSnapshotCopyObservation, p *EbsSnapshotCopyObservation, md *plugin.MergeDescription) bool {
	if k.OwnerAlias != p.OwnerAlias {
		k.OwnerAlias = p.OwnerAlias
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEbsSnapshotCopy_OwnerId(k *EbsSnapshotCopyObservation, p *EbsSnapshotCopyObservation, md *plugin.MergeDescription) bool {
	if k.OwnerId != p.OwnerId {
		k.OwnerId = p.OwnerId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEbsSnapshotCopy_VolumeSize(k *EbsSnapshotCopyObservation, p *EbsSnapshotCopyObservation, md *plugin.MergeDescription) bool {
	if k.VolumeSize != p.VolumeSize {
		k.VolumeSize = p.VolumeSize
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEbsSnapshotCopy_DataEncryptionKeyId(k *EbsSnapshotCopyObservation, p *EbsSnapshotCopyObservation, md *plugin.MergeDescription) bool {
	if k.DataEncryptionKeyId != p.DataEncryptionKeyId {
		k.DataEncryptionKeyId = p.DataEncryptionKeyId
		md.StatusUpdated = true
		return true
	}
	return false
}