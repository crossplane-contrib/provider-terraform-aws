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
	k := kube.(*ServiceDiscoveryPrivateDnsNamespace)
	p := prov.(*ServiceDiscoveryPrivateDnsNamespace)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeServiceDiscoveryPrivateDnsNamespace_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeServiceDiscoveryPrivateDnsNamespace_Vpc(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeServiceDiscoveryPrivateDnsNamespace_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeServiceDiscoveryPrivateDnsNamespace_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeServiceDiscoveryPrivateDnsNamespace_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeServiceDiscoveryPrivateDnsNamespace_HostedZone(&k.Status.AtProvider, &p.Status.AtProvider, md)
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

//mergePrimitiveContainerTemplateSpec
func MergeServiceDiscoveryPrivateDnsNamespace_Tags(k *ServiceDiscoveryPrivateDnsNamespaceParameters, p *ServiceDiscoveryPrivateDnsNamespaceParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeServiceDiscoveryPrivateDnsNamespace_Vpc(k *ServiceDiscoveryPrivateDnsNamespaceParameters, p *ServiceDiscoveryPrivateDnsNamespaceParameters, md *plugin.MergeDescription) bool {
	if k.Vpc != p.Vpc {
		p.Vpc = k.Vpc
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeServiceDiscoveryPrivateDnsNamespace_Description(k *ServiceDiscoveryPrivateDnsNamespaceParameters, p *ServiceDiscoveryPrivateDnsNamespaceParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeServiceDiscoveryPrivateDnsNamespace_Name(k *ServiceDiscoveryPrivateDnsNamespaceParameters, p *ServiceDiscoveryPrivateDnsNamespaceParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeServiceDiscoveryPrivateDnsNamespace_Arn(k *ServiceDiscoveryPrivateDnsNamespaceObservation, p *ServiceDiscoveryPrivateDnsNamespaceObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeServiceDiscoveryPrivateDnsNamespace_HostedZone(k *ServiceDiscoveryPrivateDnsNamespaceObservation, p *ServiceDiscoveryPrivateDnsNamespaceObservation, md *plugin.MergeDescription) bool {
	if k.HostedZone != p.HostedZone {
		k.HostedZone = p.HostedZone
		md.StatusUpdated = true
		return true
	}
	return false
}