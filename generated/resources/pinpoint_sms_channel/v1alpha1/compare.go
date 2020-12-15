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
	k := kube.(*PinpointSmsChannel)
	p := prov.(*PinpointSmsChannel)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergePinpointSmsChannel_ApplicationId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergePinpointSmsChannel_Enabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergePinpointSmsChannel_SenderId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergePinpointSmsChannel_ShortCode(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergePinpointSmsChannel_PromotionalMessagesPerSecond(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergePinpointSmsChannel_TransactionalMessagesPerSecond(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergePinpointSmsChannel_ApplicationId(k *PinpointSmsChannelParameters, p *PinpointSmsChannelParameters, md *plugin.MergeDescription) bool {
	if k.ApplicationId != p.ApplicationId {
		p.ApplicationId = k.ApplicationId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergePinpointSmsChannel_Enabled(k *PinpointSmsChannelParameters, p *PinpointSmsChannelParameters, md *plugin.MergeDescription) bool {
	if k.Enabled != p.Enabled {
		p.Enabled = k.Enabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergePinpointSmsChannel_SenderId(k *PinpointSmsChannelParameters, p *PinpointSmsChannelParameters, md *plugin.MergeDescription) bool {
	if k.SenderId != p.SenderId {
		p.SenderId = k.SenderId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergePinpointSmsChannel_ShortCode(k *PinpointSmsChannelParameters, p *PinpointSmsChannelParameters, md *plugin.MergeDescription) bool {
	if k.ShortCode != p.ShortCode {
		p.ShortCode = k.ShortCode
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergePinpointSmsChannel_PromotionalMessagesPerSecond(k *PinpointSmsChannelObservation, p *PinpointSmsChannelObservation, md *plugin.MergeDescription) bool {
	if k.PromotionalMessagesPerSecond != p.PromotionalMessagesPerSecond {
		k.PromotionalMessagesPerSecond = p.PromotionalMessagesPerSecond
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergePinpointSmsChannel_TransactionalMessagesPerSecond(k *PinpointSmsChannelObservation, p *PinpointSmsChannelObservation, md *plugin.MergeDescription) bool {
	if k.TransactionalMessagesPerSecond != p.TransactionalMessagesPerSecond {
		k.TransactionalMessagesPerSecond = p.TransactionalMessagesPerSecond
		md.StatusUpdated = true
		return true
	}
	return false
}