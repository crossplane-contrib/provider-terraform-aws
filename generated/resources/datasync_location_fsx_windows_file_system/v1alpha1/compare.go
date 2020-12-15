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
	k := kube.(*DatasyncLocationFsxWindowsFileSystem)
	p := prov.(*DatasyncLocationFsxWindowsFileSystem)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDatasyncLocationFsxWindowsFileSystem_User(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatasyncLocationFsxWindowsFileSystem_Domain(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatasyncLocationFsxWindowsFileSystem_FsxFilesystemArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatasyncLocationFsxWindowsFileSystem_Password(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatasyncLocationFsxWindowsFileSystem_SecurityGroupArns(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatasyncLocationFsxWindowsFileSystem_Subdirectory(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatasyncLocationFsxWindowsFileSystem_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatasyncLocationFsxWindowsFileSystem_CreationTime(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatasyncLocationFsxWindowsFileSystem_Uri(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatasyncLocationFsxWindowsFileSystem_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeDatasyncLocationFsxWindowsFileSystem_User(k *DatasyncLocationFsxWindowsFileSystemParameters, p *DatasyncLocationFsxWindowsFileSystemParameters, md *plugin.MergeDescription) bool {
	if k.User != p.User {
		p.User = k.User
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatasyncLocationFsxWindowsFileSystem_Domain(k *DatasyncLocationFsxWindowsFileSystemParameters, p *DatasyncLocationFsxWindowsFileSystemParameters, md *plugin.MergeDescription) bool {
	if k.Domain != p.Domain {
		p.Domain = k.Domain
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatasyncLocationFsxWindowsFileSystem_FsxFilesystemArn(k *DatasyncLocationFsxWindowsFileSystemParameters, p *DatasyncLocationFsxWindowsFileSystemParameters, md *plugin.MergeDescription) bool {
	if k.FsxFilesystemArn != p.FsxFilesystemArn {
		p.FsxFilesystemArn = k.FsxFilesystemArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatasyncLocationFsxWindowsFileSystem_Password(k *DatasyncLocationFsxWindowsFileSystemParameters, p *DatasyncLocationFsxWindowsFileSystemParameters, md *plugin.MergeDescription) bool {
	if k.Password != p.Password {
		p.Password = k.Password
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDatasyncLocationFsxWindowsFileSystem_SecurityGroupArns(k *DatasyncLocationFsxWindowsFileSystemParameters, p *DatasyncLocationFsxWindowsFileSystemParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.SecurityGroupArns, p.SecurityGroupArns) {
		p.SecurityGroupArns = k.SecurityGroupArns
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatasyncLocationFsxWindowsFileSystem_Subdirectory(k *DatasyncLocationFsxWindowsFileSystemParameters, p *DatasyncLocationFsxWindowsFileSystemParameters, md *plugin.MergeDescription) bool {
	if k.Subdirectory != p.Subdirectory {
		p.Subdirectory = k.Subdirectory
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDatasyncLocationFsxWindowsFileSystem_Tags(k *DatasyncLocationFsxWindowsFileSystemParameters, p *DatasyncLocationFsxWindowsFileSystemParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDatasyncLocationFsxWindowsFileSystem_CreationTime(k *DatasyncLocationFsxWindowsFileSystemObservation, p *DatasyncLocationFsxWindowsFileSystemObservation, md *plugin.MergeDescription) bool {
	if k.CreationTime != p.CreationTime {
		k.CreationTime = p.CreationTime
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDatasyncLocationFsxWindowsFileSystem_Uri(k *DatasyncLocationFsxWindowsFileSystemObservation, p *DatasyncLocationFsxWindowsFileSystemObservation, md *plugin.MergeDescription) bool {
	if k.Uri != p.Uri {
		k.Uri = p.Uri
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDatasyncLocationFsxWindowsFileSystem_Arn(k *DatasyncLocationFsxWindowsFileSystemObservation, p *DatasyncLocationFsxWindowsFileSystemObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}