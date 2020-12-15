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
	k := kube.(*Ec2CapacityReservation)
	p := prov.(*Ec2CapacityReservation)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeEc2CapacityReservation_Tenancy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2CapacityReservation_InstanceType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2CapacityReservation_EndDateType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2CapacityReservation_EphemeralStorage(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2CapacityReservation_InstanceCount(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2CapacityReservation_InstanceMatchCriteria(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2CapacityReservation_InstancePlatform(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2CapacityReservation_AvailabilityZone(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2CapacityReservation_EbsOptimized(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2CapacityReservation_EndDate(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2CapacityReservation_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEc2CapacityReservation_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeEc2CapacityReservation_Tenancy(k *Ec2CapacityReservationParameters, p *Ec2CapacityReservationParameters, md *plugin.MergeDescription) bool {
	if k.Tenancy != p.Tenancy {
		p.Tenancy = k.Tenancy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2CapacityReservation_InstanceType(k *Ec2CapacityReservationParameters, p *Ec2CapacityReservationParameters, md *plugin.MergeDescription) bool {
	if k.InstanceType != p.InstanceType {
		p.InstanceType = k.InstanceType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2CapacityReservation_EndDateType(k *Ec2CapacityReservationParameters, p *Ec2CapacityReservationParameters, md *plugin.MergeDescription) bool {
	if k.EndDateType != p.EndDateType {
		p.EndDateType = k.EndDateType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2CapacityReservation_EphemeralStorage(k *Ec2CapacityReservationParameters, p *Ec2CapacityReservationParameters, md *plugin.MergeDescription) bool {
	if k.EphemeralStorage != p.EphemeralStorage {
		p.EphemeralStorage = k.EphemeralStorage
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2CapacityReservation_InstanceCount(k *Ec2CapacityReservationParameters, p *Ec2CapacityReservationParameters, md *plugin.MergeDescription) bool {
	if k.InstanceCount != p.InstanceCount {
		p.InstanceCount = k.InstanceCount
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2CapacityReservation_InstanceMatchCriteria(k *Ec2CapacityReservationParameters, p *Ec2CapacityReservationParameters, md *plugin.MergeDescription) bool {
	if k.InstanceMatchCriteria != p.InstanceMatchCriteria {
		p.InstanceMatchCriteria = k.InstanceMatchCriteria
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2CapacityReservation_InstancePlatform(k *Ec2CapacityReservationParameters, p *Ec2CapacityReservationParameters, md *plugin.MergeDescription) bool {
	if k.InstancePlatform != p.InstancePlatform {
		p.InstancePlatform = k.InstancePlatform
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2CapacityReservation_AvailabilityZone(k *Ec2CapacityReservationParameters, p *Ec2CapacityReservationParameters, md *plugin.MergeDescription) bool {
	if k.AvailabilityZone != p.AvailabilityZone {
		p.AvailabilityZone = k.AvailabilityZone
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2CapacityReservation_EbsOptimized(k *Ec2CapacityReservationParameters, p *Ec2CapacityReservationParameters, md *plugin.MergeDescription) bool {
	if k.EbsOptimized != p.EbsOptimized {
		p.EbsOptimized = k.EbsOptimized
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEc2CapacityReservation_EndDate(k *Ec2CapacityReservationParameters, p *Ec2CapacityReservationParameters, md *plugin.MergeDescription) bool {
	if k.EndDate != p.EndDate {
		p.EndDate = k.EndDate
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeEc2CapacityReservation_Tags(k *Ec2CapacityReservationParameters, p *Ec2CapacityReservationParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEc2CapacityReservation_Arn(k *Ec2CapacityReservationObservation, p *Ec2CapacityReservationObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}