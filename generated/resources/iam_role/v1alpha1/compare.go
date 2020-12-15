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
	k := kube.(*IamRole)
	p := prov.(*IamRole)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeIamRole_ForceDetachPolicies(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamRole_MaxSessionDuration(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamRole_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamRole_NamePrefix(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamRole_Path(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamRole_AssumeRolePolicy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamRole_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamRole_PermissionsBoundary(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamRole_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamRole_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamRole_CreateDate(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamRole_UniqueId(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeIamRole_ForceDetachPolicies(k *IamRoleParameters, p *IamRoleParameters, md *plugin.MergeDescription) bool {
	if k.ForceDetachPolicies != p.ForceDetachPolicies {
		p.ForceDetachPolicies = k.ForceDetachPolicies
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamRole_MaxSessionDuration(k *IamRoleParameters, p *IamRoleParameters, md *plugin.MergeDescription) bool {
	if k.MaxSessionDuration != p.MaxSessionDuration {
		p.MaxSessionDuration = k.MaxSessionDuration
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamRole_Name(k *IamRoleParameters, p *IamRoleParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamRole_NamePrefix(k *IamRoleParameters, p *IamRoleParameters, md *plugin.MergeDescription) bool {
	if k.NamePrefix != p.NamePrefix {
		p.NamePrefix = k.NamePrefix
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamRole_Path(k *IamRoleParameters, p *IamRoleParameters, md *plugin.MergeDescription) bool {
	if k.Path != p.Path {
		p.Path = k.Path
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamRole_AssumeRolePolicy(k *IamRoleParameters, p *IamRoleParameters, md *plugin.MergeDescription) bool {
	if k.AssumeRolePolicy != p.AssumeRolePolicy {
		p.AssumeRolePolicy = k.AssumeRolePolicy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamRole_Description(k *IamRoleParameters, p *IamRoleParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamRole_PermissionsBoundary(k *IamRoleParameters, p *IamRoleParameters, md *plugin.MergeDescription) bool {
	if k.PermissionsBoundary != p.PermissionsBoundary {
		p.PermissionsBoundary = k.PermissionsBoundary
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeIamRole_Tags(k *IamRoleParameters, p *IamRoleParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeIamRole_Arn(k *IamRoleObservation, p *IamRoleObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeIamRole_CreateDate(k *IamRoleObservation, p *IamRoleObservation, md *plugin.MergeDescription) bool {
	if k.CreateDate != p.CreateDate {
		k.CreateDate = p.CreateDate
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeIamRole_UniqueId(k *IamRoleObservation, p *IamRoleObservation, md *plugin.MergeDescription) bool {
	if k.UniqueId != p.UniqueId {
		k.UniqueId = p.UniqueId
		md.StatusUpdated = true
		return true
	}
	return false
}