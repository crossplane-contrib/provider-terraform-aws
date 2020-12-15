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
	k := kube.(*SecurityGroupRule)
	p := prov.(*SecurityGroupRule)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeSecurityGroupRule_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSecurityGroupRule_PrefixListIds(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSecurityGroupRule_SecurityGroupId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSecurityGroupRule_Self(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSecurityGroupRule_SourceSecurityGroupId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSecurityGroupRule_ToPort(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSecurityGroupRule_CidrBlocks(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSecurityGroupRule_FromPort(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSecurityGroupRule_Ipv6CidrBlocks(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSecurityGroupRule_Protocol(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSecurityGroupRule_Type(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeSecurityGroupRule_Description(k *SecurityGroupRuleParameters, p *SecurityGroupRuleParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeSecurityGroupRule_PrefixListIds(k *SecurityGroupRuleParameters, p *SecurityGroupRuleParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.PrefixListIds, p.PrefixListIds) {
		p.PrefixListIds = k.PrefixListIds
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSecurityGroupRule_SecurityGroupId(k *SecurityGroupRuleParameters, p *SecurityGroupRuleParameters, md *plugin.MergeDescription) bool {
	if k.SecurityGroupId != p.SecurityGroupId {
		p.SecurityGroupId = k.SecurityGroupId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSecurityGroupRule_Self(k *SecurityGroupRuleParameters, p *SecurityGroupRuleParameters, md *plugin.MergeDescription) bool {
	if k.Self != p.Self {
		p.Self = k.Self
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSecurityGroupRule_SourceSecurityGroupId(k *SecurityGroupRuleParameters, p *SecurityGroupRuleParameters, md *plugin.MergeDescription) bool {
	if k.SourceSecurityGroupId != p.SourceSecurityGroupId {
		p.SourceSecurityGroupId = k.SourceSecurityGroupId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSecurityGroupRule_ToPort(k *SecurityGroupRuleParameters, p *SecurityGroupRuleParameters, md *plugin.MergeDescription) bool {
	if k.ToPort != p.ToPort {
		p.ToPort = k.ToPort
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeSecurityGroupRule_CidrBlocks(k *SecurityGroupRuleParameters, p *SecurityGroupRuleParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.CidrBlocks, p.CidrBlocks) {
		p.CidrBlocks = k.CidrBlocks
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSecurityGroupRule_FromPort(k *SecurityGroupRuleParameters, p *SecurityGroupRuleParameters, md *plugin.MergeDescription) bool {
	if k.FromPort != p.FromPort {
		p.FromPort = k.FromPort
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeSecurityGroupRule_Ipv6CidrBlocks(k *SecurityGroupRuleParameters, p *SecurityGroupRuleParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.Ipv6CidrBlocks, p.Ipv6CidrBlocks) {
		p.Ipv6CidrBlocks = k.Ipv6CidrBlocks
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSecurityGroupRule_Protocol(k *SecurityGroupRuleParameters, p *SecurityGroupRuleParameters, md *plugin.MergeDescription) bool {
	if k.Protocol != p.Protocol {
		p.Protocol = k.Protocol
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSecurityGroupRule_Type(k *SecurityGroupRuleParameters, p *SecurityGroupRuleParameters, md *plugin.MergeDescription) bool {
	if k.Type != p.Type {
		p.Type = k.Type
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}