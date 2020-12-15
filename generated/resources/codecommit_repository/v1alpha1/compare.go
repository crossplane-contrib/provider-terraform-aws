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
	k := kube.(*CodecommitRepository)
	p := prov.(*CodecommitRepository)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeCodecommitRepository_DefaultBranch(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCodecommitRepository_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCodecommitRepository_RepositoryName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCodecommitRepository_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCodecommitRepository_CloneUrlHttp(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCodecommitRepository_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCodecommitRepository_CloneUrlSsh(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCodecommitRepository_RepositoryId(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeCodecommitRepository_DefaultBranch(k *CodecommitRepositoryParameters, p *CodecommitRepositoryParameters, md *plugin.MergeDescription) bool {
	if k.DefaultBranch != p.DefaultBranch {
		p.DefaultBranch = k.DefaultBranch
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCodecommitRepository_Description(k *CodecommitRepositoryParameters, p *CodecommitRepositoryParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCodecommitRepository_RepositoryName(k *CodecommitRepositoryParameters, p *CodecommitRepositoryParameters, md *plugin.MergeDescription) bool {
	if k.RepositoryName != p.RepositoryName {
		p.RepositoryName = k.RepositoryName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeCodecommitRepository_Tags(k *CodecommitRepositoryParameters, p *CodecommitRepositoryParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCodecommitRepository_CloneUrlHttp(k *CodecommitRepositoryObservation, p *CodecommitRepositoryObservation, md *plugin.MergeDescription) bool {
	if k.CloneUrlHttp != p.CloneUrlHttp {
		k.CloneUrlHttp = p.CloneUrlHttp
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCodecommitRepository_Arn(k *CodecommitRepositoryObservation, p *CodecommitRepositoryObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCodecommitRepository_CloneUrlSsh(k *CodecommitRepositoryObservation, p *CodecommitRepositoryObservation, md *plugin.MergeDescription) bool {
	if k.CloneUrlSsh != p.CloneUrlSsh {
		k.CloneUrlSsh = p.CloneUrlSsh
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCodecommitRepository_RepositoryId(k *CodecommitRepositoryObservation, p *CodecommitRepositoryObservation, md *plugin.MergeDescription) bool {
	if k.RepositoryId != p.RepositoryId {
		k.RepositoryId = p.RepositoryId
		md.StatusUpdated = true
		return true
	}
	return false
}