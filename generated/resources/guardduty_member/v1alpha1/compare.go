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
	k := kube.(*GuarddutyMember)
	p := prov.(*GuarddutyMember)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeGuarddutyMember_DisableEmailNotification(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeGuarddutyMember_Email(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeGuarddutyMember_InvitationMessage(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeGuarddutyMember_Invite(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeGuarddutyMember_AccountId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeGuarddutyMember_DetectorId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeGuarddutyMember_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeGuarddutyMember_RelationshipStatus(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeGuarddutyMember_DisableEmailNotification(k *GuarddutyMemberParameters, p *GuarddutyMemberParameters, md *plugin.MergeDescription) bool {
	if k.DisableEmailNotification != p.DisableEmailNotification {
		p.DisableEmailNotification = k.DisableEmailNotification
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeGuarddutyMember_Email(k *GuarddutyMemberParameters, p *GuarddutyMemberParameters, md *plugin.MergeDescription) bool {
	if k.Email != p.Email {
		p.Email = k.Email
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeGuarddutyMember_InvitationMessage(k *GuarddutyMemberParameters, p *GuarddutyMemberParameters, md *plugin.MergeDescription) bool {
	if k.InvitationMessage != p.InvitationMessage {
		p.InvitationMessage = k.InvitationMessage
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeGuarddutyMember_Invite(k *GuarddutyMemberParameters, p *GuarddutyMemberParameters, md *plugin.MergeDescription) bool {
	if k.Invite != p.Invite {
		p.Invite = k.Invite
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeGuarddutyMember_AccountId(k *GuarddutyMemberParameters, p *GuarddutyMemberParameters, md *plugin.MergeDescription) bool {
	if k.AccountId != p.AccountId {
		p.AccountId = k.AccountId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeGuarddutyMember_DetectorId(k *GuarddutyMemberParameters, p *GuarddutyMemberParameters, md *plugin.MergeDescription) bool {
	if k.DetectorId != p.DetectorId {
		p.DetectorId = k.DetectorId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeGuarddutyMember_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeGuarddutyMember_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeGuarddutyMember_Timeouts_Update(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeGuarddutyMember_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeGuarddutyMember_Timeouts_Update(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Update != p.Update {
		p.Update = k.Update
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeGuarddutyMember_RelationshipStatus(k *GuarddutyMemberObservation, p *GuarddutyMemberObservation, md *plugin.MergeDescription) bool {
	if k.RelationshipStatus != p.RelationshipStatus {
		k.RelationshipStatus = p.RelationshipStatus
		md.StatusUpdated = true
		return true
	}
	return false
}