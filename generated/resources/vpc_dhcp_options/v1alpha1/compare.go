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
	k := kube.(*VpcDhcpOptions)
	p := prov.(*VpcDhcpOptions)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeVpcDhcpOptions_DomainName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcDhcpOptions_NetbiosNodeType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcDhcpOptions_NtpServers(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcDhcpOptions_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcDhcpOptions_DomainNameServers(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcDhcpOptions_NetbiosNameServers(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcDhcpOptions_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcDhcpOptions_OwnerId(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeVpcDhcpOptions_DomainName(k *VpcDhcpOptionsParameters, p *VpcDhcpOptionsParameters, md *plugin.MergeDescription) bool {
	if k.DomainName != p.DomainName {
		p.DomainName = k.DomainName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVpcDhcpOptions_NetbiosNodeType(k *VpcDhcpOptionsParameters, p *VpcDhcpOptionsParameters, md *plugin.MergeDescription) bool {
	if k.NetbiosNodeType != p.NetbiosNodeType {
		p.NetbiosNodeType = k.NetbiosNodeType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVpcDhcpOptions_NtpServers(k *VpcDhcpOptionsParameters, p *VpcDhcpOptionsParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.NtpServers, p.NtpServers) {
		p.NtpServers = k.NtpServers
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVpcDhcpOptions_Tags(k *VpcDhcpOptionsParameters, p *VpcDhcpOptionsParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVpcDhcpOptions_DomainNameServers(k *VpcDhcpOptionsParameters, p *VpcDhcpOptionsParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.DomainNameServers, p.DomainNameServers) {
		p.DomainNameServers = k.DomainNameServers
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVpcDhcpOptions_NetbiosNameServers(k *VpcDhcpOptionsParameters, p *VpcDhcpOptionsParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.NetbiosNameServers, p.NetbiosNameServers) {
		p.NetbiosNameServers = k.NetbiosNameServers
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVpcDhcpOptions_Arn(k *VpcDhcpOptionsObservation, p *VpcDhcpOptionsObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVpcDhcpOptions_OwnerId(k *VpcDhcpOptionsObservation, p *VpcDhcpOptionsObservation, md *plugin.MergeDescription) bool {
	if k.OwnerId != p.OwnerId {
		k.OwnerId = p.OwnerId
		md.StatusUpdated = true
		return true
	}
	return false
}