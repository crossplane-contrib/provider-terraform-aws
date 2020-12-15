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
	k := kube.(*VpcEndpointRouteTableAssociation)
	p := prov.(*VpcEndpointRouteTableAssociation)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeVpcEndpointRouteTableAssociation_RouteTableId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcEndpointRouteTableAssociation_VpcEndpointId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeVpcEndpointRouteTableAssociation_RouteTableId(k *VpcEndpointRouteTableAssociationParameters, p *VpcEndpointRouteTableAssociationParameters, md *plugin.MergeDescription) bool {
	if k.RouteTableId != p.RouteTableId {
		p.RouteTableId = k.RouteTableId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVpcEndpointRouteTableAssociation_VpcEndpointId(k *VpcEndpointRouteTableAssociationParameters, p *VpcEndpointRouteTableAssociationParameters, md *plugin.MergeDescription) bool {
	if k.VpcEndpointId != p.VpcEndpointId {
		p.VpcEndpointId = k.VpcEndpointId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}