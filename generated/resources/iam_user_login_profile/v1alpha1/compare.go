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
	k := kube.(*IamUserLoginProfile)
	p := prov.(*IamUserLoginProfile)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeIamUserLoginProfile_PasswordLength(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamUserLoginProfile_PasswordResetRequired(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamUserLoginProfile_PgpKey(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamUserLoginProfile_User(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamUserLoginProfile_EncryptedPassword(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamUserLoginProfile_KeyFingerprint(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeIamUserLoginProfile_PasswordLength(k *IamUserLoginProfileParameters, p *IamUserLoginProfileParameters, md *plugin.MergeDescription) bool {
	if k.PasswordLength != p.PasswordLength {
		p.PasswordLength = k.PasswordLength
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamUserLoginProfile_PasswordResetRequired(k *IamUserLoginProfileParameters, p *IamUserLoginProfileParameters, md *plugin.MergeDescription) bool {
	if k.PasswordResetRequired != p.PasswordResetRequired {
		p.PasswordResetRequired = k.PasswordResetRequired
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamUserLoginProfile_PgpKey(k *IamUserLoginProfileParameters, p *IamUserLoginProfileParameters, md *plugin.MergeDescription) bool {
	if k.PgpKey != p.PgpKey {
		p.PgpKey = k.PgpKey
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamUserLoginProfile_User(k *IamUserLoginProfileParameters, p *IamUserLoginProfileParameters, md *plugin.MergeDescription) bool {
	if k.User != p.User {
		p.User = k.User
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeIamUserLoginProfile_EncryptedPassword(k *IamUserLoginProfileObservation, p *IamUserLoginProfileObservation, md *plugin.MergeDescription) bool {
	if k.EncryptedPassword != p.EncryptedPassword {
		k.EncryptedPassword = p.EncryptedPassword
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeIamUserLoginProfile_KeyFingerprint(k *IamUserLoginProfileObservation, p *IamUserLoginProfileObservation, md *plugin.MergeDescription) bool {
	if k.KeyFingerprint != p.KeyFingerprint {
		k.KeyFingerprint = p.KeyFingerprint
		md.StatusUpdated = true
		return true
	}
	return false
}