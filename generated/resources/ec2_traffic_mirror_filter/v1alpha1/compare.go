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
	k := kube.(*Ec2TrafficMirrorFilter)
	p := prov.(*Ec2TrafficMirrorFilter)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeEc2TrafficMirrorFilter_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TrafficMirrorFilter_NetworkServices(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2TrafficMirrorFilter_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeEc2TrafficMirrorFilter_Description(k *Ec2TrafficMirrorFilterParameters, p *Ec2TrafficMirrorFilterParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeEc2TrafficMirrorFilter_NetworkServices(k *Ec2TrafficMirrorFilterParameters, p *Ec2TrafficMirrorFilterParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.NetworkServices, p.NetworkServices) {
		p.NetworkServices = k.NetworkServices
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeEc2TrafficMirrorFilter_Tags(k *Ec2TrafficMirrorFilterParameters, p *Ec2TrafficMirrorFilterParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}