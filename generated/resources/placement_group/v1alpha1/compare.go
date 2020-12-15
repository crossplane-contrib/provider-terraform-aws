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
	k := kube.(*PlacementGroup)
	p := prov.(*PlacementGroup)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergePlacementGroup_Strategy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergePlacementGroup_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergePlacementGroup_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergePlacementGroup_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergePlacementGroup_PlacementGroupId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergePlacementGroup_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergePlacementGroup_Strategy(k *PlacementGroupParameters, p *PlacementGroupParameters, md *plugin.MergeDescription) bool {
	if k.Strategy != p.Strategy {
		p.Strategy = k.Strategy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergePlacementGroup_Tags(k *PlacementGroupParameters, p *PlacementGroupParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(p.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergePlacementGroup_Id(k *PlacementGroupParameters, p *PlacementGroupParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergePlacementGroup_Name(k *PlacementGroupParameters, p *PlacementGroupParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergePlacementGroup_PlacementGroupId(k *PlacementGroupObservation, p *PlacementGroupObservation, md *plugin.MergeDescription) bool {
	if k.PlacementGroupId != p.PlacementGroupId {
		k.PlacementGroupId = p.PlacementGroupId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergePlacementGroup_Arn(k *PlacementGroupObservation, p *PlacementGroupObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}