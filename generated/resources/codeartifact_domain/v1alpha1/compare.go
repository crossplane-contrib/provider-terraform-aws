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
	k := kube.(*CodeartifactDomain)
	p := prov.(*CodeartifactDomain)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeCodeartifactDomain_Domain(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCodeartifactDomain_EncryptionKey(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCodeartifactDomain_RepositoryCount(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCodeartifactDomain_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCodeartifactDomain_AssetSizeBytes(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCodeartifactDomain_CreatedTime(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCodeartifactDomain_Owner(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeCodeartifactDomain_Domain(k *CodeartifactDomainParameters, p *CodeartifactDomainParameters, md *plugin.MergeDescription) bool {
	if k.Domain != p.Domain {
		p.Domain = k.Domain
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCodeartifactDomain_EncryptionKey(k *CodeartifactDomainParameters, p *CodeartifactDomainParameters, md *plugin.MergeDescription) bool {
	if k.EncryptionKey != p.EncryptionKey {
		p.EncryptionKey = k.EncryptionKey
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCodeartifactDomain_RepositoryCount(k *CodeartifactDomainObservation, p *CodeartifactDomainObservation, md *plugin.MergeDescription) bool {
	if k.RepositoryCount != p.RepositoryCount {
		k.RepositoryCount = p.RepositoryCount
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCodeartifactDomain_Arn(k *CodeartifactDomainObservation, p *CodeartifactDomainObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCodeartifactDomain_AssetSizeBytes(k *CodeartifactDomainObservation, p *CodeartifactDomainObservation, md *plugin.MergeDescription) bool {
	if k.AssetSizeBytes != p.AssetSizeBytes {
		k.AssetSizeBytes = p.AssetSizeBytes
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCodeartifactDomain_CreatedTime(k *CodeartifactDomainObservation, p *CodeartifactDomainObservation, md *plugin.MergeDescription) bool {
	if k.CreatedTime != p.CreatedTime {
		k.CreatedTime = p.CreatedTime
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCodeartifactDomain_Owner(k *CodeartifactDomainObservation, p *CodeartifactDomainObservation, md *plugin.MergeDescription) bool {
	if k.Owner != p.Owner {
		k.Owner = p.Owner
		md.StatusUpdated = true
		return true
	}
	return false
}