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
	k := kube.(*IamAccessKey)
	p := prov.(*IamAccessKey)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeIamAccessKey_PgpKey(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamAccessKey_Status(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamAccessKey_User(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamAccessKey_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamAccessKey_KeyFingerprint(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamAccessKey_Secret(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamAccessKey_SesSmtpPasswordV4(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamAccessKey_EncryptedSecret(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeIamAccessKey_PgpKey(k *IamAccessKeyParameters, p *IamAccessKeyParameters, md *plugin.MergeDescription) bool {
	if k.PgpKey != p.PgpKey {
		p.PgpKey = k.PgpKey
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamAccessKey_Status(k *IamAccessKeyParameters, p *IamAccessKeyParameters, md *plugin.MergeDescription) bool {
	if k.Status != p.Status {
		p.Status = k.Status
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamAccessKey_User(k *IamAccessKeyParameters, p *IamAccessKeyParameters, md *plugin.MergeDescription) bool {
	if k.User != p.User {
		p.User = k.User
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamAccessKey_Id(k *IamAccessKeyParameters, p *IamAccessKeyParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeIamAccessKey_KeyFingerprint(k *IamAccessKeyObservation, p *IamAccessKeyObservation, md *plugin.MergeDescription) bool {
	if k.KeyFingerprint != p.KeyFingerprint {
		k.KeyFingerprint = p.KeyFingerprint
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeIamAccessKey_Secret(k *IamAccessKeyObservation, p *IamAccessKeyObservation, md *plugin.MergeDescription) bool {
	if k.Secret != p.Secret {
		k.Secret = p.Secret
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeIamAccessKey_SesSmtpPasswordV4(k *IamAccessKeyObservation, p *IamAccessKeyObservation, md *plugin.MergeDescription) bool {
	if k.SesSmtpPasswordV4 != p.SesSmtpPasswordV4 {
		k.SesSmtpPasswordV4 = p.SesSmtpPasswordV4
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeIamAccessKey_EncryptedSecret(k *IamAccessKeyObservation, p *IamAccessKeyObservation, md *plugin.MergeDescription) bool {
	if k.EncryptedSecret != p.EncryptedSecret {
		k.EncryptedSecret = p.EncryptedSecret
		md.StatusUpdated = true
		return true
	}
	return false
}