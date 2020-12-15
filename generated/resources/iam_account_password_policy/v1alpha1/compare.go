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
	k := kube.(*IamAccountPasswordPolicy)
	p := prov.(*IamAccountPasswordPolicy)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeIamAccountPasswordPolicy_RequireUppercaseCharacters(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamAccountPasswordPolicy_AllowUsersToChangePassword(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamAccountPasswordPolicy_PasswordReusePrevention(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamAccountPasswordPolicy_RequireNumbers(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamAccountPasswordPolicy_HardExpiry(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamAccountPasswordPolicy_MaxPasswordAge(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamAccountPasswordPolicy_MinimumPasswordLength(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamAccountPasswordPolicy_RequireLowercaseCharacters(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamAccountPasswordPolicy_RequireSymbols(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamAccountPasswordPolicy_ExpirePasswords(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeIamAccountPasswordPolicy_RequireUppercaseCharacters(k *IamAccountPasswordPolicyParameters, p *IamAccountPasswordPolicyParameters, md *plugin.MergeDescription) bool {
	if k.RequireUppercaseCharacters != p.RequireUppercaseCharacters {
		p.RequireUppercaseCharacters = k.RequireUppercaseCharacters
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamAccountPasswordPolicy_AllowUsersToChangePassword(k *IamAccountPasswordPolicyParameters, p *IamAccountPasswordPolicyParameters, md *plugin.MergeDescription) bool {
	if k.AllowUsersToChangePassword != p.AllowUsersToChangePassword {
		p.AllowUsersToChangePassword = k.AllowUsersToChangePassword
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamAccountPasswordPolicy_PasswordReusePrevention(k *IamAccountPasswordPolicyParameters, p *IamAccountPasswordPolicyParameters, md *plugin.MergeDescription) bool {
	if k.PasswordReusePrevention != p.PasswordReusePrevention {
		p.PasswordReusePrevention = k.PasswordReusePrevention
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamAccountPasswordPolicy_RequireNumbers(k *IamAccountPasswordPolicyParameters, p *IamAccountPasswordPolicyParameters, md *plugin.MergeDescription) bool {
	if k.RequireNumbers != p.RequireNumbers {
		p.RequireNumbers = k.RequireNumbers
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamAccountPasswordPolicy_HardExpiry(k *IamAccountPasswordPolicyParameters, p *IamAccountPasswordPolicyParameters, md *plugin.MergeDescription) bool {
	if k.HardExpiry != p.HardExpiry {
		p.HardExpiry = k.HardExpiry
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamAccountPasswordPolicy_MaxPasswordAge(k *IamAccountPasswordPolicyParameters, p *IamAccountPasswordPolicyParameters, md *plugin.MergeDescription) bool {
	if k.MaxPasswordAge != p.MaxPasswordAge {
		p.MaxPasswordAge = k.MaxPasswordAge
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamAccountPasswordPolicy_MinimumPasswordLength(k *IamAccountPasswordPolicyParameters, p *IamAccountPasswordPolicyParameters, md *plugin.MergeDescription) bool {
	if k.MinimumPasswordLength != p.MinimumPasswordLength {
		p.MinimumPasswordLength = k.MinimumPasswordLength
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamAccountPasswordPolicy_RequireLowercaseCharacters(k *IamAccountPasswordPolicyParameters, p *IamAccountPasswordPolicyParameters, md *plugin.MergeDescription) bool {
	if k.RequireLowercaseCharacters != p.RequireLowercaseCharacters {
		p.RequireLowercaseCharacters = k.RequireLowercaseCharacters
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamAccountPasswordPolicy_RequireSymbols(k *IamAccountPasswordPolicyParameters, p *IamAccountPasswordPolicyParameters, md *plugin.MergeDescription) bool {
	if k.RequireSymbols != p.RequireSymbols {
		p.RequireSymbols = k.RequireSymbols
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeIamAccountPasswordPolicy_ExpirePasswords(k *IamAccountPasswordPolicyObservation, p *IamAccountPasswordPolicyObservation, md *plugin.MergeDescription) bool {
	if k.ExpirePasswords != p.ExpirePasswords {
		k.ExpirePasswords = p.ExpirePasswords
		md.StatusUpdated = true
		return true
	}
	return false
}