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
	k := kube.(*IamGroupPolicy)
	p := prov.(*IamGroupPolicy)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeIamGroupPolicy_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamGroupPolicy_NamePrefix(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamGroupPolicy_Policy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamGroupPolicy_Group(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeIamGroupPolicy_Name(k *IamGroupPolicyParameters, p *IamGroupPolicyParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamGroupPolicy_NamePrefix(k *IamGroupPolicyParameters, p *IamGroupPolicyParameters, md *plugin.MergeDescription) bool {
	if k.NamePrefix != p.NamePrefix {
		p.NamePrefix = k.NamePrefix
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamGroupPolicy_Policy(k *IamGroupPolicyParameters, p *IamGroupPolicyParameters, md *plugin.MergeDescription) bool {
	if k.Policy != p.Policy {
		p.Policy = k.Policy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamGroupPolicy_Group(k *IamGroupPolicyParameters, p *IamGroupPolicyParameters, md *plugin.MergeDescription) bool {
	if k.Group != p.Group {
		p.Group = k.Group
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}