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
	k := kube.(*Route53HealthCheck)
	p := prov.(*Route53HealthCheck)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeRoute53HealthCheck_ChildHealthchecks(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute53HealthCheck_CloudwatchAlarmName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute53HealthCheck_FailureThreshold(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute53HealthCheck_Port(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute53HealthCheck_ChildHealthThreshold(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute53HealthCheck_ResourcePath(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute53HealthCheck_EnableSni(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute53HealthCheck_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute53HealthCheck_Regions(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute53HealthCheck_RequestInterval(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute53HealthCheck_SearchString(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute53HealthCheck_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute53HealthCheck_Type(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute53HealthCheck_CloudwatchAlarmRegion(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute53HealthCheck_Fqdn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute53HealthCheck_InsufficientDataHealthStatus(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute53HealthCheck_InvertHealthcheck(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute53HealthCheck_IpAddress(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute53HealthCheck_MeasureLatency(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute53HealthCheck_ReferenceName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRoute53HealthCheck_Disabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeRoute53HealthCheck_ChildHealthchecks(k *Route53HealthCheckParameters, p *Route53HealthCheckParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(p.ChildHealthchecks, p.ChildHealthchecks) {
		p.ChildHealthchecks = k.ChildHealthchecks
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute53HealthCheck_CloudwatchAlarmName(k *Route53HealthCheckParameters, p *Route53HealthCheckParameters, md *plugin.MergeDescription) bool {
	if k.CloudwatchAlarmName != p.CloudwatchAlarmName {
		p.CloudwatchAlarmName = k.CloudwatchAlarmName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute53HealthCheck_FailureThreshold(k *Route53HealthCheckParameters, p *Route53HealthCheckParameters, md *plugin.MergeDescription) bool {
	if k.FailureThreshold != p.FailureThreshold {
		p.FailureThreshold = k.FailureThreshold
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute53HealthCheck_Port(k *Route53HealthCheckParameters, p *Route53HealthCheckParameters, md *plugin.MergeDescription) bool {
	if k.Port != p.Port {
		p.Port = k.Port
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute53HealthCheck_ChildHealthThreshold(k *Route53HealthCheckParameters, p *Route53HealthCheckParameters, md *plugin.MergeDescription) bool {
	if k.ChildHealthThreshold != p.ChildHealthThreshold {
		p.ChildHealthThreshold = k.ChildHealthThreshold
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute53HealthCheck_ResourcePath(k *Route53HealthCheckParameters, p *Route53HealthCheckParameters, md *plugin.MergeDescription) bool {
	if k.ResourcePath != p.ResourcePath {
		p.ResourcePath = k.ResourcePath
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute53HealthCheck_EnableSni(k *Route53HealthCheckParameters, p *Route53HealthCheckParameters, md *plugin.MergeDescription) bool {
	if k.EnableSni != p.EnableSni {
		p.EnableSni = k.EnableSni
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute53HealthCheck_Id(k *Route53HealthCheckParameters, p *Route53HealthCheckParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeRoute53HealthCheck_Regions(k *Route53HealthCheckParameters, p *Route53HealthCheckParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(p.Regions, p.Regions) {
		p.Regions = k.Regions
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute53HealthCheck_RequestInterval(k *Route53HealthCheckParameters, p *Route53HealthCheckParameters, md *plugin.MergeDescription) bool {
	if k.RequestInterval != p.RequestInterval {
		p.RequestInterval = k.RequestInterval
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute53HealthCheck_SearchString(k *Route53HealthCheckParameters, p *Route53HealthCheckParameters, md *plugin.MergeDescription) bool {
	if k.SearchString != p.SearchString {
		p.SearchString = k.SearchString
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeRoute53HealthCheck_Tags(k *Route53HealthCheckParameters, p *Route53HealthCheckParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(p.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute53HealthCheck_Type(k *Route53HealthCheckParameters, p *Route53HealthCheckParameters, md *plugin.MergeDescription) bool {
	if k.Type != p.Type {
		p.Type = k.Type
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute53HealthCheck_CloudwatchAlarmRegion(k *Route53HealthCheckParameters, p *Route53HealthCheckParameters, md *plugin.MergeDescription) bool {
	if k.CloudwatchAlarmRegion != p.CloudwatchAlarmRegion {
		p.CloudwatchAlarmRegion = k.CloudwatchAlarmRegion
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute53HealthCheck_Fqdn(k *Route53HealthCheckParameters, p *Route53HealthCheckParameters, md *plugin.MergeDescription) bool {
	if k.Fqdn != p.Fqdn {
		p.Fqdn = k.Fqdn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute53HealthCheck_InsufficientDataHealthStatus(k *Route53HealthCheckParameters, p *Route53HealthCheckParameters, md *plugin.MergeDescription) bool {
	if k.InsufficientDataHealthStatus != p.InsufficientDataHealthStatus {
		p.InsufficientDataHealthStatus = k.InsufficientDataHealthStatus
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute53HealthCheck_InvertHealthcheck(k *Route53HealthCheckParameters, p *Route53HealthCheckParameters, md *plugin.MergeDescription) bool {
	if k.InvertHealthcheck != p.InvertHealthcheck {
		p.InvertHealthcheck = k.InvertHealthcheck
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute53HealthCheck_IpAddress(k *Route53HealthCheckParameters, p *Route53HealthCheckParameters, md *plugin.MergeDescription) bool {
	if k.IpAddress != p.IpAddress {
		p.IpAddress = k.IpAddress
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute53HealthCheck_MeasureLatency(k *Route53HealthCheckParameters, p *Route53HealthCheckParameters, md *plugin.MergeDescription) bool {
	if k.MeasureLatency != p.MeasureLatency {
		p.MeasureLatency = k.MeasureLatency
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute53HealthCheck_ReferenceName(k *Route53HealthCheckParameters, p *Route53HealthCheckParameters, md *plugin.MergeDescription) bool {
	if k.ReferenceName != p.ReferenceName {
		p.ReferenceName = k.ReferenceName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRoute53HealthCheck_Disabled(k *Route53HealthCheckParameters, p *Route53HealthCheckParameters, md *plugin.MergeDescription) bool {
	if k.Disabled != p.Disabled {
		p.Disabled = k.Disabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}