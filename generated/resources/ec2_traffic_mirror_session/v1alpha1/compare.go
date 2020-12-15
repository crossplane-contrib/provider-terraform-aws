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
	k := kube.(*Ec2TrafficMirrorSession)
	p := prov.(*Ec2TrafficMirrorSession)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeEc2TrafficMirrorSession_NetworkInterfaceId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TrafficMirrorSession_SessionNumber(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TrafficMirrorSession_TrafficMirrorFilterId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TrafficMirrorSession_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TrafficMirrorSession_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TrafficMirrorSession_PacketLength(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TrafficMirrorSession_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TrafficMirrorSession_TrafficMirrorTargetId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TrafficMirrorSession_VirtualNetworkId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TrafficMirrorSession_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeEc2TrafficMirrorSession_NetworkInterfaceId(k *Ec2TrafficMirrorSessionParameters, p *Ec2TrafficMirrorSessionParameters, md *plugin.MergeDescription) bool {
	if k.NetworkInterfaceId != p.NetworkInterfaceId {
		p.NetworkInterfaceId = k.NetworkInterfaceId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2TrafficMirrorSession_SessionNumber(k *Ec2TrafficMirrorSessionParameters, p *Ec2TrafficMirrorSessionParameters, md *plugin.MergeDescription) bool {
	if k.SessionNumber != p.SessionNumber {
		p.SessionNumber = k.SessionNumber
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2TrafficMirrorSession_TrafficMirrorFilterId(k *Ec2TrafficMirrorSessionParameters, p *Ec2TrafficMirrorSessionParameters, md *plugin.MergeDescription) bool {
	if k.TrafficMirrorFilterId != p.TrafficMirrorFilterId {
		p.TrafficMirrorFilterId = k.TrafficMirrorFilterId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2TrafficMirrorSession_Id(k *Ec2TrafficMirrorSessionParameters, p *Ec2TrafficMirrorSessionParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2TrafficMirrorSession_Description(k *Ec2TrafficMirrorSessionParameters, p *Ec2TrafficMirrorSessionParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2TrafficMirrorSession_PacketLength(k *Ec2TrafficMirrorSessionParameters, p *Ec2TrafficMirrorSessionParameters, md *plugin.MergeDescription) bool {
	if k.PacketLength != p.PacketLength {
		p.PacketLength = k.PacketLength
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeEc2TrafficMirrorSession_Tags(k *Ec2TrafficMirrorSessionParameters, p *Ec2TrafficMirrorSessionParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(p.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2TrafficMirrorSession_TrafficMirrorTargetId(k *Ec2TrafficMirrorSessionParameters, p *Ec2TrafficMirrorSessionParameters, md *plugin.MergeDescription) bool {
	if k.TrafficMirrorTargetId != p.TrafficMirrorTargetId {
		p.TrafficMirrorTargetId = k.TrafficMirrorTargetId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2TrafficMirrorSession_VirtualNetworkId(k *Ec2TrafficMirrorSessionParameters, p *Ec2TrafficMirrorSessionParameters, md *plugin.MergeDescription) bool {
	if k.VirtualNetworkId != p.VirtualNetworkId {
		p.VirtualNetworkId = k.VirtualNetworkId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEc2TrafficMirrorSession_Arn(k *Ec2TrafficMirrorSessionObservation, p *Ec2TrafficMirrorSessionObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}