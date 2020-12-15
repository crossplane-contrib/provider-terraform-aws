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
	k := kube.(*QuicksightUser)
	p := prov.(*QuicksightUser)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeQuicksightUser_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeQuicksightUser_IdentityType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeQuicksightUser_Namespace(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeQuicksightUser_SessionName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeQuicksightUser_UserName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeQuicksightUser_UserRole(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeQuicksightUser_AwsAccountId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeQuicksightUser_IamArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeQuicksightUser_Email(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeQuicksightUser_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeQuicksightUser_Id(k *QuicksightUserParameters, p *QuicksightUserParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeQuicksightUser_IdentityType(k *QuicksightUserParameters, p *QuicksightUserParameters, md *plugin.MergeDescription) bool {
	if k.IdentityType != p.IdentityType {
		p.IdentityType = k.IdentityType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeQuicksightUser_Namespace(k *QuicksightUserParameters, p *QuicksightUserParameters, md *plugin.MergeDescription) bool {
	if k.Namespace != p.Namespace {
		p.Namespace = k.Namespace
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeQuicksightUser_SessionName(k *QuicksightUserParameters, p *QuicksightUserParameters, md *plugin.MergeDescription) bool {
	if k.SessionName != p.SessionName {
		p.SessionName = k.SessionName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeQuicksightUser_UserName(k *QuicksightUserParameters, p *QuicksightUserParameters, md *plugin.MergeDescription) bool {
	if k.UserName != p.UserName {
		p.UserName = k.UserName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeQuicksightUser_UserRole(k *QuicksightUserParameters, p *QuicksightUserParameters, md *plugin.MergeDescription) bool {
	if k.UserRole != p.UserRole {
		p.UserRole = k.UserRole
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeQuicksightUser_AwsAccountId(k *QuicksightUserParameters, p *QuicksightUserParameters, md *plugin.MergeDescription) bool {
	if k.AwsAccountId != p.AwsAccountId {
		p.AwsAccountId = k.AwsAccountId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeQuicksightUser_IamArn(k *QuicksightUserParameters, p *QuicksightUserParameters, md *plugin.MergeDescription) bool {
	if k.IamArn != p.IamArn {
		p.IamArn = k.IamArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeQuicksightUser_Email(k *QuicksightUserParameters, p *QuicksightUserParameters, md *plugin.MergeDescription) bool {
	if k.Email != p.Email {
		p.Email = k.Email
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeQuicksightUser_Arn(k *QuicksightUserObservation, p *QuicksightUserObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}