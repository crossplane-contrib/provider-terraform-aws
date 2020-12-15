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
	k := kube.(*VpcEndpointService)
	p := prov.(*VpcEndpointService)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeVpcEndpointService_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcEndpointService_AllowedPrincipals(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcEndpointService_NetworkLoadBalancerArns(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcEndpointService_AcceptanceRequired(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcEndpointService_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcEndpointService_AvailabilityZones(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcEndpointService_ManagesVpcEndpoints(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcEndpointService_PrivateDnsName(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcEndpointService_State(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcEndpointService_BaseEndpointDnsNames(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcEndpointService_ServiceName(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcEndpointService_ServiceType(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeVpcEndpointService_Tags(k *VpcEndpointServiceParameters, p *VpcEndpointServiceParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVpcEndpointService_AllowedPrincipals(k *VpcEndpointServiceParameters, p *VpcEndpointServiceParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.AllowedPrincipals, p.AllowedPrincipals) {
		p.AllowedPrincipals = k.AllowedPrincipals
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVpcEndpointService_NetworkLoadBalancerArns(k *VpcEndpointServiceParameters, p *VpcEndpointServiceParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.NetworkLoadBalancerArns, p.NetworkLoadBalancerArns) {
		p.NetworkLoadBalancerArns = k.NetworkLoadBalancerArns
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVpcEndpointService_AcceptanceRequired(k *VpcEndpointServiceParameters, p *VpcEndpointServiceParameters, md *plugin.MergeDescription) bool {
	if k.AcceptanceRequired != p.AcceptanceRequired {
		p.AcceptanceRequired = k.AcceptanceRequired
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVpcEndpointService_Arn(k *VpcEndpointServiceObservation, p *VpcEndpointServiceObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateStatus
func MergeVpcEndpointService_AvailabilityZones(k *VpcEndpointServiceObservation, p *VpcEndpointServiceObservation, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.AvailabilityZones, p.AvailabilityZones) {
		k.AvailabilityZones = p.AvailabilityZones
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVpcEndpointService_ManagesVpcEndpoints(k *VpcEndpointServiceObservation, p *VpcEndpointServiceObservation, md *plugin.MergeDescription) bool {
	if k.ManagesVpcEndpoints != p.ManagesVpcEndpoints {
		k.ManagesVpcEndpoints = p.ManagesVpcEndpoints
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVpcEndpointService_PrivateDnsName(k *VpcEndpointServiceObservation, p *VpcEndpointServiceObservation, md *plugin.MergeDescription) bool {
	if k.PrivateDnsName != p.PrivateDnsName {
		k.PrivateDnsName = p.PrivateDnsName
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVpcEndpointService_State(k *VpcEndpointServiceObservation, p *VpcEndpointServiceObservation, md *plugin.MergeDescription) bool {
	if k.State != p.State {
		k.State = p.State
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateStatus
func MergeVpcEndpointService_BaseEndpointDnsNames(k *VpcEndpointServiceObservation, p *VpcEndpointServiceObservation, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.BaseEndpointDnsNames, p.BaseEndpointDnsNames) {
		k.BaseEndpointDnsNames = p.BaseEndpointDnsNames
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVpcEndpointService_ServiceName(k *VpcEndpointServiceObservation, p *VpcEndpointServiceObservation, md *plugin.MergeDescription) bool {
	if k.ServiceName != p.ServiceName {
		k.ServiceName = p.ServiceName
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVpcEndpointService_ServiceType(k *VpcEndpointServiceObservation, p *VpcEndpointServiceObservation, md *plugin.MergeDescription) bool {
	if k.ServiceType != p.ServiceType {
		k.ServiceType = p.ServiceType
		md.StatusUpdated = true
		return true
	}
	return false
}