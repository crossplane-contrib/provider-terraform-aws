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
	k := kube.(*RamResourceShareAccepter)
	p := prov.(*RamResourceShareAccepter)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeRamResourceShareAccepter_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRamResourceShareAccepter_ShareArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRamResourceShareAccepter_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRamResourceShareAccepter_Status(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRamResourceShareAccepter_Resources(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRamResourceShareAccepter_ShareId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRamResourceShareAccepter_ShareName(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRamResourceShareAccepter_InvitationArn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRamResourceShareAccepter_ReceiverAccountId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRamResourceShareAccepter_SenderAccountId(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeRamResourceShareAccepter_Id(k *RamResourceShareAccepterParameters, p *RamResourceShareAccepterParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRamResourceShareAccepter_ShareArn(k *RamResourceShareAccepterParameters, p *RamResourceShareAccepterParameters, md *plugin.MergeDescription) bool {
	if k.ShareArn != p.ShareArn {
		p.ShareArn = k.ShareArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeRamResourceShareAccepter_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeRamResourceShareAccepter_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRamResourceShareAccepter_Timeouts_Delete(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeRamResourceShareAccepter_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRamResourceShareAccepter_Timeouts_Delete(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Delete != p.Delete {
		p.Delete = k.Delete
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeRamResourceShareAccepter_Status(k *RamResourceShareAccepterObservation, p *RamResourceShareAccepterObservation, md *plugin.MergeDescription) bool {
	if k.Status != p.Status {
		k.Status = p.Status
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateStatus
func MergeRamResourceShareAccepter_Resources(k *RamResourceShareAccepterObservation, p *RamResourceShareAccepterObservation, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(p.Resources, p.Resources) {
		k.Resources = p.Resources
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeRamResourceShareAccepter_ShareId(k *RamResourceShareAccepterObservation, p *RamResourceShareAccepterObservation, md *plugin.MergeDescription) bool {
	if k.ShareId != p.ShareId {
		k.ShareId = p.ShareId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeRamResourceShareAccepter_ShareName(k *RamResourceShareAccepterObservation, p *RamResourceShareAccepterObservation, md *plugin.MergeDescription) bool {
	if k.ShareName != p.ShareName {
		k.ShareName = p.ShareName
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeRamResourceShareAccepter_InvitationArn(k *RamResourceShareAccepterObservation, p *RamResourceShareAccepterObservation, md *plugin.MergeDescription) bool {
	if k.InvitationArn != p.InvitationArn {
		k.InvitationArn = p.InvitationArn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeRamResourceShareAccepter_ReceiverAccountId(k *RamResourceShareAccepterObservation, p *RamResourceShareAccepterObservation, md *plugin.MergeDescription) bool {
	if k.ReceiverAccountId != p.ReceiverAccountId {
		k.ReceiverAccountId = p.ReceiverAccountId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeRamResourceShareAccepter_SenderAccountId(k *RamResourceShareAccepterObservation, p *RamResourceShareAccepterObservation, md *plugin.MergeDescription) bool {
	if k.SenderAccountId != p.SenderAccountId {
		k.SenderAccountId = p.SenderAccountId
		md.StatusUpdated = true
		return true
	}
	return false
}