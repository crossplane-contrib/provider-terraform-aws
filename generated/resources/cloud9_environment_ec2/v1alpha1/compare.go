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
	k := kube.(*Cloud9EnvironmentEc2)
	p := prov.(*Cloud9EnvironmentEc2)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeCloud9EnvironmentEc2_AutomaticStopTimeMinutes(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloud9EnvironmentEc2_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloud9EnvironmentEc2_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloud9EnvironmentEc2_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloud9EnvironmentEc2_InstanceType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloud9EnvironmentEc2_OwnerArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloud9EnvironmentEc2_SubnetId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloud9EnvironmentEc2_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloud9EnvironmentEc2_Type(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeCloud9EnvironmentEc2_AutomaticStopTimeMinutes(k *Cloud9EnvironmentEc2Parameters, p *Cloud9EnvironmentEc2Parameters, md *plugin.MergeDescription) bool {
	if k.AutomaticStopTimeMinutes != p.AutomaticStopTimeMinutes {
		p.AutomaticStopTimeMinutes = k.AutomaticStopTimeMinutes
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloud9EnvironmentEc2_Name(k *Cloud9EnvironmentEc2Parameters, p *Cloud9EnvironmentEc2Parameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeCloud9EnvironmentEc2_Tags(k *Cloud9EnvironmentEc2Parameters, p *Cloud9EnvironmentEc2Parameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloud9EnvironmentEc2_Description(k *Cloud9EnvironmentEc2Parameters, p *Cloud9EnvironmentEc2Parameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloud9EnvironmentEc2_InstanceType(k *Cloud9EnvironmentEc2Parameters, p *Cloud9EnvironmentEc2Parameters, md *plugin.MergeDescription) bool {
	if k.InstanceType != p.InstanceType {
		p.InstanceType = k.InstanceType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloud9EnvironmentEc2_OwnerArn(k *Cloud9EnvironmentEc2Parameters, p *Cloud9EnvironmentEc2Parameters, md *plugin.MergeDescription) bool {
	if k.OwnerArn != p.OwnerArn {
		p.OwnerArn = k.OwnerArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloud9EnvironmentEc2_SubnetId(k *Cloud9EnvironmentEc2Parameters, p *Cloud9EnvironmentEc2Parameters, md *plugin.MergeDescription) bool {
	if k.SubnetId != p.SubnetId {
		p.SubnetId = k.SubnetId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCloud9EnvironmentEc2_Arn(k *Cloud9EnvironmentEc2Observation, p *Cloud9EnvironmentEc2Observation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCloud9EnvironmentEc2_Type(k *Cloud9EnvironmentEc2Observation, p *Cloud9EnvironmentEc2Observation, md *plugin.MergeDescription) bool {
	if k.Type != p.Type {
		k.Type = p.Type
		md.StatusUpdated = true
		return true
	}
	return false
}