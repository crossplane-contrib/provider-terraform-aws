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
	k := kube.(*CodeartifactDomainPermissionsPolicy)
	p := prov.(*CodeartifactDomainPermissionsPolicy)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeCodeartifactDomainPermissionsPolicy_Domain(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCodeartifactDomainPermissionsPolicy_DomainOwner(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCodeartifactDomainPermissionsPolicy_PolicyDocument(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCodeartifactDomainPermissionsPolicy_PolicyRevision(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCodeartifactDomainPermissionsPolicy_ResourceArn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeCodeartifactDomainPermissionsPolicy_Domain(k *CodeartifactDomainPermissionsPolicyParameters, p *CodeartifactDomainPermissionsPolicyParameters, md *plugin.MergeDescription) bool {
	if k.Domain != p.Domain {
		p.Domain = k.Domain
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCodeartifactDomainPermissionsPolicy_DomainOwner(k *CodeartifactDomainPermissionsPolicyParameters, p *CodeartifactDomainPermissionsPolicyParameters, md *plugin.MergeDescription) bool {
	if k.DomainOwner != p.DomainOwner {
		p.DomainOwner = k.DomainOwner
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCodeartifactDomainPermissionsPolicy_PolicyDocument(k *CodeartifactDomainPermissionsPolicyParameters, p *CodeartifactDomainPermissionsPolicyParameters, md *plugin.MergeDescription) bool {
	if k.PolicyDocument != p.PolicyDocument {
		p.PolicyDocument = k.PolicyDocument
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCodeartifactDomainPermissionsPolicy_PolicyRevision(k *CodeartifactDomainPermissionsPolicyParameters, p *CodeartifactDomainPermissionsPolicyParameters, md *plugin.MergeDescription) bool {
	if k.PolicyRevision != p.PolicyRevision {
		p.PolicyRevision = k.PolicyRevision
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCodeartifactDomainPermissionsPolicy_ResourceArn(k *CodeartifactDomainPermissionsPolicyObservation, p *CodeartifactDomainPermissionsPolicyObservation, md *plugin.MergeDescription) bool {
	if k.ResourceArn != p.ResourceArn {
		k.ResourceArn = p.ResourceArn
		md.StatusUpdated = true
		return true
	}
	return false
}