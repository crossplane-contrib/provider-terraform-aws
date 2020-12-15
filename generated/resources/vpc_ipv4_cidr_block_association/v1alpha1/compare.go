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
	k := kube.(*VpcIpv4CidrBlockAssociation)
	p := prov.(*VpcIpv4CidrBlockAssociation)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeVpcIpv4CidrBlockAssociation_VpcId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcIpv4CidrBlockAssociation_CidrBlock(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcIpv4CidrBlockAssociation_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
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
func MergeVpcIpv4CidrBlockAssociation_VpcId(k *VpcIpv4CidrBlockAssociationParameters, p *VpcIpv4CidrBlockAssociationParameters, md *plugin.MergeDescription) bool {
	if k.VpcId != p.VpcId {
		p.VpcId = k.VpcId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVpcIpv4CidrBlockAssociation_CidrBlock(k *VpcIpv4CidrBlockAssociationParameters, p *VpcIpv4CidrBlockAssociationParameters, md *plugin.MergeDescription) bool {
	if k.CidrBlock != p.CidrBlock {
		p.CidrBlock = k.CidrBlock
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeVpcIpv4CidrBlockAssociation_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeVpcIpv4CidrBlockAssociation_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcIpv4CidrBlockAssociation_Timeouts_Delete(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeVpcIpv4CidrBlockAssociation_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVpcIpv4CidrBlockAssociation_Timeouts_Delete(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Delete != p.Delete {
		p.Delete = k.Delete
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}