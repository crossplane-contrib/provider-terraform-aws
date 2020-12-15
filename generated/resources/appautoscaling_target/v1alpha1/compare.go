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
	k := kube.(*AppautoscalingTarget)
	p := prov.(*AppautoscalingTarget)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeAppautoscalingTarget_MaxCapacity(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAppautoscalingTarget_MinCapacity(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAppautoscalingTarget_ResourceId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAppautoscalingTarget_RoleArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAppautoscalingTarget_ScalableDimension(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAppautoscalingTarget_ServiceNamespace(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeAppautoscalingTarget_MaxCapacity(k *AppautoscalingTargetParameters, p *AppautoscalingTargetParameters, md *plugin.MergeDescription) bool {
	if k.MaxCapacity != p.MaxCapacity {
		p.MaxCapacity = k.MaxCapacity
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAppautoscalingTarget_MinCapacity(k *AppautoscalingTargetParameters, p *AppautoscalingTargetParameters, md *plugin.MergeDescription) bool {
	if k.MinCapacity != p.MinCapacity {
		p.MinCapacity = k.MinCapacity
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAppautoscalingTarget_ResourceId(k *AppautoscalingTargetParameters, p *AppautoscalingTargetParameters, md *plugin.MergeDescription) bool {
	if k.ResourceId != p.ResourceId {
		p.ResourceId = k.ResourceId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAppautoscalingTarget_RoleArn(k *AppautoscalingTargetParameters, p *AppautoscalingTargetParameters, md *plugin.MergeDescription) bool {
	if k.RoleArn != p.RoleArn {
		p.RoleArn = k.RoleArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAppautoscalingTarget_ScalableDimension(k *AppautoscalingTargetParameters, p *AppautoscalingTargetParameters, md *plugin.MergeDescription) bool {
	if k.ScalableDimension != p.ScalableDimension {
		p.ScalableDimension = k.ScalableDimension
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAppautoscalingTarget_ServiceNamespace(k *AppautoscalingTargetParameters, p *AppautoscalingTargetParameters, md *plugin.MergeDescription) bool {
	if k.ServiceNamespace != p.ServiceNamespace {
		p.ServiceNamespace = k.ServiceNamespace
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}