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
	k := kube.(*Eip)
	p := prov.(*Eip)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeEip_NetworkInterface(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEip_Vpc(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEip_Instance(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEip_PublicIpv4Pool(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEip_AssociateWithPrivateIp(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEip_CustomerOwnedIpv4Pool(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEip_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEip_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEip_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEip_PrivateIp(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEip_PublicIp(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEip_PublicDns(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEip_PrivateDns(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEip_AllocationId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEip_CustomerOwnedIp(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEip_Domain(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEip_AssociationId(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeEip_NetworkInterface(k *EipParameters, p *EipParameters, md *plugin.MergeDescription) bool {
	if k.NetworkInterface != p.NetworkInterface {
		p.NetworkInterface = k.NetworkInterface
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEip_Vpc(k *EipParameters, p *EipParameters, md *plugin.MergeDescription) bool {
	if k.Vpc != p.Vpc {
		p.Vpc = k.Vpc
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEip_Instance(k *EipParameters, p *EipParameters, md *plugin.MergeDescription) bool {
	if k.Instance != p.Instance {
		p.Instance = k.Instance
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEip_PublicIpv4Pool(k *EipParameters, p *EipParameters, md *plugin.MergeDescription) bool {
	if k.PublicIpv4Pool != p.PublicIpv4Pool {
		p.PublicIpv4Pool = k.PublicIpv4Pool
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEip_AssociateWithPrivateIp(k *EipParameters, p *EipParameters, md *plugin.MergeDescription) bool {
	if k.AssociateWithPrivateIp != p.AssociateWithPrivateIp {
		p.AssociateWithPrivateIp = k.AssociateWithPrivateIp
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEip_CustomerOwnedIpv4Pool(k *EipParameters, p *EipParameters, md *plugin.MergeDescription) bool {
	if k.CustomerOwnedIpv4Pool != p.CustomerOwnedIpv4Pool {
		p.CustomerOwnedIpv4Pool = k.CustomerOwnedIpv4Pool
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEip_Id(k *EipParameters, p *EipParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeEip_Tags(k *EipParameters, p *EipParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(p.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeEip_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeEip_Timeouts_Delete(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEip_Timeouts_Read(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEip_Timeouts_Update(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeEip_Timeouts_Delete(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Delete != p.Delete {
		p.Delete = k.Delete
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEip_Timeouts_Read(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Read != p.Read {
		p.Read = k.Read
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEip_Timeouts_Update(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Update != p.Update {
		p.Update = k.Update
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEip_PrivateIp(k *EipObservation, p *EipObservation, md *plugin.MergeDescription) bool {
	if k.PrivateIp != p.PrivateIp {
		k.PrivateIp = p.PrivateIp
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEip_PublicIp(k *EipObservation, p *EipObservation, md *plugin.MergeDescription) bool {
	if k.PublicIp != p.PublicIp {
		k.PublicIp = p.PublicIp
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEip_PublicDns(k *EipObservation, p *EipObservation, md *plugin.MergeDescription) bool {
	if k.PublicDns != p.PublicDns {
		k.PublicDns = p.PublicDns
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEip_PrivateDns(k *EipObservation, p *EipObservation, md *plugin.MergeDescription) bool {
	if k.PrivateDns != p.PrivateDns {
		k.PrivateDns = p.PrivateDns
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEip_AllocationId(k *EipObservation, p *EipObservation, md *plugin.MergeDescription) bool {
	if k.AllocationId != p.AllocationId {
		k.AllocationId = p.AllocationId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEip_CustomerOwnedIp(k *EipObservation, p *EipObservation, md *plugin.MergeDescription) bool {
	if k.CustomerOwnedIp != p.CustomerOwnedIp {
		k.CustomerOwnedIp = p.CustomerOwnedIp
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEip_Domain(k *EipObservation, p *EipObservation, md *plugin.MergeDescription) bool {
	if k.Domain != p.Domain {
		k.Domain = p.Domain
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeEip_AssociationId(k *EipObservation, p *EipObservation, md *plugin.MergeDescription) bool {
	if k.AssociationId != p.AssociationId {
		k.AssociationId = p.AssociationId
		md.StatusUpdated = true
		return true
	}
	return false
}