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
	k := kube.(*EfsMountTarget)
	p := prov.(*EfsMountTarget)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeEfsMountTarget_SubnetId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEfsMountTarget_IpAddress(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEfsMountTarget_SecurityGroups(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEfsMountTarget_FileSystemId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEfsMountTarget_FileSystemArn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEfsMountTarget_AvailabilityZoneId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEfsMountTarget_DnsName(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEfsMountTarget_MountTargetDnsName(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEfsMountTarget_NetworkInterfaceId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEfsMountTarget_OwnerId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEfsMountTarget_AvailabilityZoneName(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeEfsMountTarget_SubnetId(k *EfsMountTargetParameters, p *EfsMountTargetParameters, md *plugin.MergeDescription) bool {
	if k.SubnetId != p.SubnetId {
		p.SubnetId = k.SubnetId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEfsMountTarget_IpAddress(k *EfsMountTargetParameters, p *EfsMountTargetParameters, md *plugin.MergeDescription) bool {
	if k.IpAddress != p.IpAddress {
		p.IpAddress = k.IpAddress
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeEfsMountTarget_SecurityGroups(k *EfsMountTargetParameters, p *EfsMountTargetParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.SecurityGroups, p.SecurityGroups) {
		p.SecurityGroups = k.SecurityGroups
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEfsMountTarget_FileSystemId(k *EfsMountTargetParameters, p *EfsMountTargetParameters, md *plugin.MergeDescription) bool {
	if k.FileSystemId != p.FileSystemId {
		p.FileSystemId = k.FileSystemId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEfsMountTarget_FileSystemArn(k *EfsMountTargetObservation, p *EfsMountTargetObservation, md *plugin.MergeDescription) bool {
	if k.FileSystemArn != p.FileSystemArn {
		k.FileSystemArn = p.FileSystemArn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEfsMountTarget_AvailabilityZoneId(k *EfsMountTargetObservation, p *EfsMountTargetObservation, md *plugin.MergeDescription) bool {
	if k.AvailabilityZoneId != p.AvailabilityZoneId {
		k.AvailabilityZoneId = p.AvailabilityZoneId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEfsMountTarget_DnsName(k *EfsMountTargetObservation, p *EfsMountTargetObservation, md *plugin.MergeDescription) bool {
	if k.DnsName != p.DnsName {
		k.DnsName = p.DnsName
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEfsMountTarget_MountTargetDnsName(k *EfsMountTargetObservation, p *EfsMountTargetObservation, md *plugin.MergeDescription) bool {
	if k.MountTargetDnsName != p.MountTargetDnsName {
		k.MountTargetDnsName = p.MountTargetDnsName
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEfsMountTarget_NetworkInterfaceId(k *EfsMountTargetObservation, p *EfsMountTargetObservation, md *plugin.MergeDescription) bool {
	if k.NetworkInterfaceId != p.NetworkInterfaceId {
		k.NetworkInterfaceId = p.NetworkInterfaceId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEfsMountTarget_OwnerId(k *EfsMountTargetObservation, p *EfsMountTargetObservation, md *plugin.MergeDescription) bool {
	if k.OwnerId != p.OwnerId {
		k.OwnerId = p.OwnerId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEfsMountTarget_AvailabilityZoneName(k *EfsMountTargetObservation, p *EfsMountTargetObservation, md *plugin.MergeDescription) bool {
	if k.AvailabilityZoneName != p.AvailabilityZoneName {
		k.AvailabilityZoneName = p.AvailabilityZoneName
		md.StatusUpdated = true
		return true
	}
	return false
}