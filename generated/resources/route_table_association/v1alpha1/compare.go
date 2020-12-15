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
	k := kube.(*RouteTableAssociation)
	p := prov.(*RouteTableAssociation)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeRouteTableAssociation_GatewayId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRouteTableAssociation_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRouteTableAssociation_RouteTableId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRouteTableAssociation_SubnetId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeRouteTableAssociation_GatewayId(k *RouteTableAssociationParameters, p *RouteTableAssociationParameters, md *plugin.MergeDescription) bool {
	if k.GatewayId != p.GatewayId {
		p.GatewayId = k.GatewayId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRouteTableAssociation_Id(k *RouteTableAssociationParameters, p *RouteTableAssociationParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRouteTableAssociation_RouteTableId(k *RouteTableAssociationParameters, p *RouteTableAssociationParameters, md *plugin.MergeDescription) bool {
	if k.RouteTableId != p.RouteTableId {
		p.RouteTableId = k.RouteTableId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRouteTableAssociation_SubnetId(k *RouteTableAssociationParameters, p *RouteTableAssociationParameters, md *plugin.MergeDescription) bool {
	if k.SubnetId != p.SubnetId {
		p.SubnetId = k.SubnetId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}