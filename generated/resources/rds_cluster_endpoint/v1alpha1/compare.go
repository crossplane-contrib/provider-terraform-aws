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
	k := kube.(*RdsClusterEndpoint)
	p := prov.(*RdsClusterEndpoint)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeRdsClusterEndpoint_CustomEndpointType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterEndpoint_ExcludedMembers(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterEndpoint_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterEndpoint_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterEndpoint_ClusterIdentifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterEndpoint_StaticMembers(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterEndpoint_ClusterEndpointIdentifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterEndpoint_Endpoint(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRdsClusterEndpoint_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeRdsClusterEndpoint_CustomEndpointType(k *RdsClusterEndpointParameters, p *RdsClusterEndpointParameters, md *plugin.MergeDescription) bool {
	if k.CustomEndpointType != p.CustomEndpointType {
		p.CustomEndpointType = k.CustomEndpointType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeRdsClusterEndpoint_ExcludedMembers(k *RdsClusterEndpointParameters, p *RdsClusterEndpointParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(p.ExcludedMembers, p.ExcludedMembers) {
		p.ExcludedMembers = k.ExcludedMembers
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRdsClusterEndpoint_Id(k *RdsClusterEndpointParameters, p *RdsClusterEndpointParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeRdsClusterEndpoint_Tags(k *RdsClusterEndpointParameters, p *RdsClusterEndpointParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(p.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRdsClusterEndpoint_ClusterIdentifier(k *RdsClusterEndpointParameters, p *RdsClusterEndpointParameters, md *plugin.MergeDescription) bool {
	if k.ClusterIdentifier != p.ClusterIdentifier {
		p.ClusterIdentifier = k.ClusterIdentifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeRdsClusterEndpoint_StaticMembers(k *RdsClusterEndpointParameters, p *RdsClusterEndpointParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(p.StaticMembers, p.StaticMembers) {
		p.StaticMembers = k.StaticMembers
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRdsClusterEndpoint_ClusterEndpointIdentifier(k *RdsClusterEndpointParameters, p *RdsClusterEndpointParameters, md *plugin.MergeDescription) bool {
	if k.ClusterEndpointIdentifier != p.ClusterEndpointIdentifier {
		p.ClusterEndpointIdentifier = k.ClusterEndpointIdentifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeRdsClusterEndpoint_Endpoint(k *RdsClusterEndpointObservation, p *RdsClusterEndpointObservation, md *plugin.MergeDescription) bool {
	if k.Endpoint != p.Endpoint {
		k.Endpoint = p.Endpoint
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeRdsClusterEndpoint_Arn(k *RdsClusterEndpointObservation, p *RdsClusterEndpointObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}