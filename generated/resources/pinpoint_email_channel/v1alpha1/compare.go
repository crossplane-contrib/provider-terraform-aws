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
	k := kube.(*PinpointEmailChannel)
	p := prov.(*PinpointEmailChannel)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergePinpointEmailChannel_ApplicationId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergePinpointEmailChannel_Enabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergePinpointEmailChannel_FromAddress(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergePinpointEmailChannel_Identity(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergePinpointEmailChannel_RoleArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergePinpointEmailChannel_MessagesPerSecond(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergePinpointEmailChannel_ApplicationId(k *PinpointEmailChannelParameters, p *PinpointEmailChannelParameters, md *plugin.MergeDescription) bool {
	if k.ApplicationId != p.ApplicationId {
		p.ApplicationId = k.ApplicationId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergePinpointEmailChannel_Enabled(k *PinpointEmailChannelParameters, p *PinpointEmailChannelParameters, md *plugin.MergeDescription) bool {
	if k.Enabled != p.Enabled {
		p.Enabled = k.Enabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergePinpointEmailChannel_FromAddress(k *PinpointEmailChannelParameters, p *PinpointEmailChannelParameters, md *plugin.MergeDescription) bool {
	if k.FromAddress != p.FromAddress {
		p.FromAddress = k.FromAddress
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergePinpointEmailChannel_Identity(k *PinpointEmailChannelParameters, p *PinpointEmailChannelParameters, md *plugin.MergeDescription) bool {
	if k.Identity != p.Identity {
		p.Identity = k.Identity
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergePinpointEmailChannel_RoleArn(k *PinpointEmailChannelParameters, p *PinpointEmailChannelParameters, md *plugin.MergeDescription) bool {
	if k.RoleArn != p.RoleArn {
		p.RoleArn = k.RoleArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergePinpointEmailChannel_MessagesPerSecond(k *PinpointEmailChannelObservation, p *PinpointEmailChannelObservation, md *plugin.MergeDescription) bool {
	if k.MessagesPerSecond != p.MessagesPerSecond {
		k.MessagesPerSecond = p.MessagesPerSecond
		md.StatusUpdated = true
		return true
	}
	return false
}