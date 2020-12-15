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
	k := kube.(*SsmActivation)
	p := prov.(*SsmActivation)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeSsmActivation_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSsmActivation_ExpirationDate(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSsmActivation_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSsmActivation_RegistrationLimit(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSsmActivation_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSsmActivation_IamRole(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSsmActivation_RegistrationCount(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSsmActivation_ActivationCode(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSsmActivation_Expired(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeSsmActivation_Tags(k *SsmActivationParameters, p *SsmActivationParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSsmActivation_ExpirationDate(k *SsmActivationParameters, p *SsmActivationParameters, md *plugin.MergeDescription) bool {
	if k.ExpirationDate != p.ExpirationDate {
		p.ExpirationDate = k.ExpirationDate
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSsmActivation_Name(k *SsmActivationParameters, p *SsmActivationParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSsmActivation_RegistrationLimit(k *SsmActivationParameters, p *SsmActivationParameters, md *plugin.MergeDescription) bool {
	if k.RegistrationLimit != p.RegistrationLimit {
		p.RegistrationLimit = k.RegistrationLimit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSsmActivation_Description(k *SsmActivationParameters, p *SsmActivationParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSsmActivation_IamRole(k *SsmActivationParameters, p *SsmActivationParameters, md *plugin.MergeDescription) bool {
	if k.IamRole != p.IamRole {
		p.IamRole = k.IamRole
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeSsmActivation_RegistrationCount(k *SsmActivationObservation, p *SsmActivationObservation, md *plugin.MergeDescription) bool {
	if k.RegistrationCount != p.RegistrationCount {
		k.RegistrationCount = p.RegistrationCount
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeSsmActivation_ActivationCode(k *SsmActivationObservation, p *SsmActivationObservation, md *plugin.MergeDescription) bool {
	if k.ActivationCode != p.ActivationCode {
		k.ActivationCode = p.ActivationCode
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeSsmActivation_Expired(k *SsmActivationObservation, p *SsmActivationObservation, md *plugin.MergeDescription) bool {
	if k.Expired != p.Expired {
		k.Expired = p.Expired
		md.StatusUpdated = true
		return true
	}
	return false
}