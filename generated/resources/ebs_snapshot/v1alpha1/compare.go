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
	k := kube.(*EbsSnapshot)
	p := prov.(*EbsSnapshot)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeEbsSnapshot_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEbsSnapshot_VolumeId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEbsSnapshot_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEbsSnapshot_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEbsSnapshot_VolumeSize(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEbsSnapshot_DataEncryptionKeyId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEbsSnapshot_KmsKeyId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEbsSnapshot_OwnerAlias(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEbsSnapshot_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEbsSnapshot_Encrypted(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEbsSnapshot_OwnerId(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeEbsSnapshot_Tags(k *EbsSnapshotParameters, p *EbsSnapshotParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEbsSnapshot_VolumeId(k *EbsSnapshotParameters, p *EbsSnapshotParameters, md *plugin.MergeDescription) bool {
	if k.VolumeId != p.VolumeId {
		p.VolumeId = k.VolumeId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEbsSnapshot_Description(k *EbsSnapshotParameters, p *EbsSnapshotParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeEbsSnapshot_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeEbsSnapshot_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEbsSnapshot_Timeouts_Delete(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeEbsSnapshot_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEbsSnapshot_Timeouts_Delete(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Delete != p.Delete {
		p.Delete = k.Delete
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEbsSnapshot_VolumeSize(k *EbsSnapshotObservation, p *EbsSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.VolumeSize != p.VolumeSize {
		k.VolumeSize = p.VolumeSize
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEbsSnapshot_DataEncryptionKeyId(k *EbsSnapshotObservation, p *EbsSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.DataEncryptionKeyId != p.DataEncryptionKeyId {
		k.DataEncryptionKeyId = p.DataEncryptionKeyId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEbsSnapshot_KmsKeyId(k *EbsSnapshotObservation, p *EbsSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.KmsKeyId != p.KmsKeyId {
		k.KmsKeyId = p.KmsKeyId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEbsSnapshot_OwnerAlias(k *EbsSnapshotObservation, p *EbsSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.OwnerAlias != p.OwnerAlias {
		k.OwnerAlias = p.OwnerAlias
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEbsSnapshot_Arn(k *EbsSnapshotObservation, p *EbsSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEbsSnapshot_Encrypted(k *EbsSnapshotObservation, p *EbsSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.Encrypted != p.Encrypted {
		k.Encrypted = p.Encrypted
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEbsSnapshot_OwnerId(k *EbsSnapshotObservation, p *EbsSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.OwnerId != p.OwnerId {
		k.OwnerId = p.OwnerId
		md.StatusUpdated = true
		return true
	}
	return false
}