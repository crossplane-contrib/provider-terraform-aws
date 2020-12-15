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
	k := kube.(*OrganizationsAccount)
	p := prov.(*OrganizationsAccount)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeOrganizationsAccount_ParentId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeOrganizationsAccount_Email(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeOrganizationsAccount_IamUserAccessToBilling(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeOrganizationsAccount_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeOrganizationsAccount_RoleName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeOrganizationsAccount_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeOrganizationsAccount_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeOrganizationsAccount_JoinedMethod(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeOrganizationsAccount_JoinedTimestamp(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeOrganizationsAccount_Status(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeOrganizationsAccount_ParentId(k *OrganizationsAccountParameters, p *OrganizationsAccountParameters, md *plugin.MergeDescription) bool {
	if k.ParentId != p.ParentId {
		p.ParentId = k.ParentId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeOrganizationsAccount_Email(k *OrganizationsAccountParameters, p *OrganizationsAccountParameters, md *plugin.MergeDescription) bool {
	if k.Email != p.Email {
		p.Email = k.Email
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeOrganizationsAccount_IamUserAccessToBilling(k *OrganizationsAccountParameters, p *OrganizationsAccountParameters, md *plugin.MergeDescription) bool {
	if k.IamUserAccessToBilling != p.IamUserAccessToBilling {
		p.IamUserAccessToBilling = k.IamUserAccessToBilling
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeOrganizationsAccount_Name(k *OrganizationsAccountParameters, p *OrganizationsAccountParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeOrganizationsAccount_RoleName(k *OrganizationsAccountParameters, p *OrganizationsAccountParameters, md *plugin.MergeDescription) bool {
	if k.RoleName != p.RoleName {
		p.RoleName = k.RoleName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeOrganizationsAccount_Tags(k *OrganizationsAccountParameters, p *OrganizationsAccountParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeOrganizationsAccount_Arn(k *OrganizationsAccountObservation, p *OrganizationsAccountObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeOrganizationsAccount_JoinedMethod(k *OrganizationsAccountObservation, p *OrganizationsAccountObservation, md *plugin.MergeDescription) bool {
	if k.JoinedMethod != p.JoinedMethod {
		k.JoinedMethod = p.JoinedMethod
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeOrganizationsAccount_JoinedTimestamp(k *OrganizationsAccountObservation, p *OrganizationsAccountObservation, md *plugin.MergeDescription) bool {
	if k.JoinedTimestamp != p.JoinedTimestamp {
		k.JoinedTimestamp = p.JoinedTimestamp
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeOrganizationsAccount_Status(k *OrganizationsAccountObservation, p *OrganizationsAccountObservation, md *plugin.MergeDescription) bool {
	if k.Status != p.Status {
		k.Status = p.Status
		md.StatusUpdated = true
		return true
	}
	return false
}