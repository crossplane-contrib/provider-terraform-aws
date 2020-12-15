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
	k := kube.(*NetworkAclRule)
	p := prov.(*NetworkAclRule)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeNetworkAclRule_IcmpCode(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNetworkAclRule_Ipv6CidrBlock(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNetworkAclRule_NetworkAclId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNetworkAclRule_Protocol(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNetworkAclRule_ToPort(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNetworkAclRule_Egress(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNetworkAclRule_FromPort(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNetworkAclRule_IcmpType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNetworkAclRule_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNetworkAclRule_RuleAction(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNetworkAclRule_RuleNumber(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNetworkAclRule_CidrBlock(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeNetworkAclRule_IcmpCode(k *NetworkAclRuleParameters, p *NetworkAclRuleParameters, md *plugin.MergeDescription) bool {
	if k.IcmpCode != p.IcmpCode {
		p.IcmpCode = k.IcmpCode
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNetworkAclRule_Ipv6CidrBlock(k *NetworkAclRuleParameters, p *NetworkAclRuleParameters, md *plugin.MergeDescription) bool {
	if k.Ipv6CidrBlock != p.Ipv6CidrBlock {
		p.Ipv6CidrBlock = k.Ipv6CidrBlock
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNetworkAclRule_NetworkAclId(k *NetworkAclRuleParameters, p *NetworkAclRuleParameters, md *plugin.MergeDescription) bool {
	if k.NetworkAclId != p.NetworkAclId {
		p.NetworkAclId = k.NetworkAclId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNetworkAclRule_Protocol(k *NetworkAclRuleParameters, p *NetworkAclRuleParameters, md *plugin.MergeDescription) bool {
	if k.Protocol != p.Protocol {
		p.Protocol = k.Protocol
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNetworkAclRule_ToPort(k *NetworkAclRuleParameters, p *NetworkAclRuleParameters, md *plugin.MergeDescription) bool {
	if k.ToPort != p.ToPort {
		p.ToPort = k.ToPort
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNetworkAclRule_Egress(k *NetworkAclRuleParameters, p *NetworkAclRuleParameters, md *plugin.MergeDescription) bool {
	if k.Egress != p.Egress {
		p.Egress = k.Egress
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNetworkAclRule_FromPort(k *NetworkAclRuleParameters, p *NetworkAclRuleParameters, md *plugin.MergeDescription) bool {
	if k.FromPort != p.FromPort {
		p.FromPort = k.FromPort
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNetworkAclRule_IcmpType(k *NetworkAclRuleParameters, p *NetworkAclRuleParameters, md *plugin.MergeDescription) bool {
	if k.IcmpType != p.IcmpType {
		p.IcmpType = k.IcmpType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNetworkAclRule_Id(k *NetworkAclRuleParameters, p *NetworkAclRuleParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNetworkAclRule_RuleAction(k *NetworkAclRuleParameters, p *NetworkAclRuleParameters, md *plugin.MergeDescription) bool {
	if k.RuleAction != p.RuleAction {
		p.RuleAction = k.RuleAction
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNetworkAclRule_RuleNumber(k *NetworkAclRuleParameters, p *NetworkAclRuleParameters, md *plugin.MergeDescription) bool {
	if k.RuleNumber != p.RuleNumber {
		p.RuleNumber = k.RuleNumber
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNetworkAclRule_CidrBlock(k *NetworkAclRuleParameters, p *NetworkAclRuleParameters, md *plugin.MergeDescription) bool {
	if k.CidrBlock != p.CidrBlock {
		p.CidrBlock = k.CidrBlock
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}