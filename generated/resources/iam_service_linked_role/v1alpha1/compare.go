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
	k := kube.(*IamServiceLinkedRole)
	p := prov.(*IamServiceLinkedRole)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeIamServiceLinkedRole_AwsServiceName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamServiceLinkedRole_CustomSuffix(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamServiceLinkedRole_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamServiceLinkedRole_Path(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamServiceLinkedRole_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamServiceLinkedRole_Name(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamServiceLinkedRole_UniqueId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamServiceLinkedRole_CreateDate(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeIamServiceLinkedRole_AwsServiceName(k *IamServiceLinkedRoleParameters, p *IamServiceLinkedRoleParameters, md *plugin.MergeDescription) bool {
	if k.AwsServiceName != p.AwsServiceName {
		p.AwsServiceName = k.AwsServiceName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamServiceLinkedRole_CustomSuffix(k *IamServiceLinkedRoleParameters, p *IamServiceLinkedRoleParameters, md *plugin.MergeDescription) bool {
	if k.CustomSuffix != p.CustomSuffix {
		p.CustomSuffix = k.CustomSuffix
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamServiceLinkedRole_Description(k *IamServiceLinkedRoleParameters, p *IamServiceLinkedRoleParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeIamServiceLinkedRole_Path(k *IamServiceLinkedRoleObservation, p *IamServiceLinkedRoleObservation, md *plugin.MergeDescription) bool {
	if k.Path != p.Path {
		k.Path = p.Path
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeIamServiceLinkedRole_Arn(k *IamServiceLinkedRoleObservation, p *IamServiceLinkedRoleObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeIamServiceLinkedRole_Name(k *IamServiceLinkedRoleObservation, p *IamServiceLinkedRoleObservation, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		k.Name = p.Name
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeIamServiceLinkedRole_UniqueId(k *IamServiceLinkedRoleObservation, p *IamServiceLinkedRoleObservation, md *plugin.MergeDescription) bool {
	if k.UniqueId != p.UniqueId {
		k.UniqueId = p.UniqueId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeIamServiceLinkedRole_CreateDate(k *IamServiceLinkedRoleObservation, p *IamServiceLinkedRoleObservation, md *plugin.MergeDescription) bool {
	if k.CreateDate != p.CreateDate {
		k.CreateDate = p.CreateDate
		md.StatusUpdated = true
		return true
	}
	return false
}