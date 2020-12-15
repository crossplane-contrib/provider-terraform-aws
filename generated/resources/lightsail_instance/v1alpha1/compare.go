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
	k := kube.(*LightsailInstance)
	p := prov.(*LightsailInstance)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeLightsailInstance_AvailabilityZone(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLightsailInstance_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLightsailInstance_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLightsailInstance_BlueprintId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLightsailInstance_BundleId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLightsailInstance_KeyPairName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLightsailInstance_UserData(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLightsailInstance_Ipv6Address(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLightsailInstance_Username(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLightsailInstance_CpuCount(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLightsailInstance_CreatedAt(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLightsailInstance_IsStaticIp(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLightsailInstance_PrivateIpAddress(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLightsailInstance_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLightsailInstance_PublicIpAddress(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLightsailInstance_RamSize(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeLightsailInstance_AvailabilityZone(k *LightsailInstanceParameters, p *LightsailInstanceParameters, md *plugin.MergeDescription) bool {
	if k.AvailabilityZone != p.AvailabilityZone {
		p.AvailabilityZone = k.AvailabilityZone
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLightsailInstance_Name(k *LightsailInstanceParameters, p *LightsailInstanceParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeLightsailInstance_Tags(k *LightsailInstanceParameters, p *LightsailInstanceParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLightsailInstance_BlueprintId(k *LightsailInstanceParameters, p *LightsailInstanceParameters, md *plugin.MergeDescription) bool {
	if k.BlueprintId != p.BlueprintId {
		p.BlueprintId = k.BlueprintId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLightsailInstance_BundleId(k *LightsailInstanceParameters, p *LightsailInstanceParameters, md *plugin.MergeDescription) bool {
	if k.BundleId != p.BundleId {
		p.BundleId = k.BundleId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLightsailInstance_KeyPairName(k *LightsailInstanceParameters, p *LightsailInstanceParameters, md *plugin.MergeDescription) bool {
	if k.KeyPairName != p.KeyPairName {
		p.KeyPairName = k.KeyPairName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLightsailInstance_UserData(k *LightsailInstanceParameters, p *LightsailInstanceParameters, md *plugin.MergeDescription) bool {
	if k.UserData != p.UserData {
		p.UserData = k.UserData
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeLightsailInstance_Ipv6Address(k *LightsailInstanceObservation, p *LightsailInstanceObservation, md *plugin.MergeDescription) bool {
	if k.Ipv6Address != p.Ipv6Address {
		k.Ipv6Address = p.Ipv6Address
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeLightsailInstance_Username(k *LightsailInstanceObservation, p *LightsailInstanceObservation, md *plugin.MergeDescription) bool {
	if k.Username != p.Username {
		k.Username = p.Username
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeLightsailInstance_CpuCount(k *LightsailInstanceObservation, p *LightsailInstanceObservation, md *plugin.MergeDescription) bool {
	if k.CpuCount != p.CpuCount {
		k.CpuCount = p.CpuCount
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeLightsailInstance_CreatedAt(k *LightsailInstanceObservation, p *LightsailInstanceObservation, md *plugin.MergeDescription) bool {
	if k.CreatedAt != p.CreatedAt {
		k.CreatedAt = p.CreatedAt
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeLightsailInstance_IsStaticIp(k *LightsailInstanceObservation, p *LightsailInstanceObservation, md *plugin.MergeDescription) bool {
	if k.IsStaticIp != p.IsStaticIp {
		k.IsStaticIp = p.IsStaticIp
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeLightsailInstance_PrivateIpAddress(k *LightsailInstanceObservation, p *LightsailInstanceObservation, md *plugin.MergeDescription) bool {
	if k.PrivateIpAddress != p.PrivateIpAddress {
		k.PrivateIpAddress = p.PrivateIpAddress
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeLightsailInstance_Arn(k *LightsailInstanceObservation, p *LightsailInstanceObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeLightsailInstance_PublicIpAddress(k *LightsailInstanceObservation, p *LightsailInstanceObservation, md *plugin.MergeDescription) bool {
	if k.PublicIpAddress != p.PublicIpAddress {
		k.PublicIpAddress = p.PublicIpAddress
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeLightsailInstance_RamSize(k *LightsailInstanceObservation, p *LightsailInstanceObservation, md *plugin.MergeDescription) bool {
	if k.RamSize != p.RamSize {
		k.RamSize = p.RamSize
		md.StatusUpdated = true
		return true
	}
	return false
}