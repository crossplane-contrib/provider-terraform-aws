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
	k := kube.(*ServicequotasServiceQuota)
	p := prov.(*ServicequotasServiceQuota)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeServicequotasServiceQuota_Value(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeServicequotasServiceQuota_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeServicequotasServiceQuota_QuotaCode(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeServicequotasServiceQuota_ServiceCode(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeServicequotasServiceQuota_RequestId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeServicequotasServiceQuota_Adjustable(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeServicequotasServiceQuota_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeServicequotasServiceQuota_DefaultValue(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeServicequotasServiceQuota_ServiceName(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeServicequotasServiceQuota_QuotaName(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeServicequotasServiceQuota_RequestStatus(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeServicequotasServiceQuota_Value(k *ServicequotasServiceQuotaParameters, p *ServicequotasServiceQuotaParameters, md *plugin.MergeDescription) bool {
	if k.Value != p.Value {
		p.Value = k.Value
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeServicequotasServiceQuota_Id(k *ServicequotasServiceQuotaParameters, p *ServicequotasServiceQuotaParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeServicequotasServiceQuota_QuotaCode(k *ServicequotasServiceQuotaParameters, p *ServicequotasServiceQuotaParameters, md *plugin.MergeDescription) bool {
	if k.QuotaCode != p.QuotaCode {
		p.QuotaCode = k.QuotaCode
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeServicequotasServiceQuota_ServiceCode(k *ServicequotasServiceQuotaParameters, p *ServicequotasServiceQuotaParameters, md *plugin.MergeDescription) bool {
	if k.ServiceCode != p.ServiceCode {
		p.ServiceCode = k.ServiceCode
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeServicequotasServiceQuota_RequestId(k *ServicequotasServiceQuotaObservation, p *ServicequotasServiceQuotaObservation, md *plugin.MergeDescription) bool {
	if k.RequestId != p.RequestId {
		k.RequestId = p.RequestId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeServicequotasServiceQuota_Adjustable(k *ServicequotasServiceQuotaObservation, p *ServicequotasServiceQuotaObservation, md *plugin.MergeDescription) bool {
	if k.Adjustable != p.Adjustable {
		k.Adjustable = p.Adjustable
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeServicequotasServiceQuota_Arn(k *ServicequotasServiceQuotaObservation, p *ServicequotasServiceQuotaObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeServicequotasServiceQuota_DefaultValue(k *ServicequotasServiceQuotaObservation, p *ServicequotasServiceQuotaObservation, md *plugin.MergeDescription) bool {
	if k.DefaultValue != p.DefaultValue {
		k.DefaultValue = p.DefaultValue
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeServicequotasServiceQuota_ServiceName(k *ServicequotasServiceQuotaObservation, p *ServicequotasServiceQuotaObservation, md *plugin.MergeDescription) bool {
	if k.ServiceName != p.ServiceName {
		k.ServiceName = p.ServiceName
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeServicequotasServiceQuota_QuotaName(k *ServicequotasServiceQuotaObservation, p *ServicequotasServiceQuotaObservation, md *plugin.MergeDescription) bool {
	if k.QuotaName != p.QuotaName {
		k.QuotaName = p.QuotaName
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeServicequotasServiceQuota_RequestStatus(k *ServicequotasServiceQuotaObservation, p *ServicequotasServiceQuotaObservation, md *plugin.MergeDescription) bool {
	if k.RequestStatus != p.RequestStatus {
		k.RequestStatus = p.RequestStatus
		md.StatusUpdated = true
		return true
	}
	return false
}