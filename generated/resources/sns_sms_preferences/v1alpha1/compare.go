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
	k := kube.(*SnsSmsPreferences)
	p := prov.(*SnsSmsPreferences)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeSnsSmsPreferences_MonthlySpendLimit(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsSmsPreferences_UsageReportS3Bucket(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsSmsPreferences_DefaultSenderId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsSmsPreferences_DefaultSmsType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsSmsPreferences_DeliveryStatusIamRoleArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsSmsPreferences_DeliveryStatusSuccessSamplingRate(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeSnsSmsPreferences_MonthlySpendLimit(k *SnsSmsPreferencesParameters, p *SnsSmsPreferencesParameters, md *plugin.MergeDescription) bool {
	if k.MonthlySpendLimit != p.MonthlySpendLimit {
		p.MonthlySpendLimit = k.MonthlySpendLimit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsSmsPreferences_UsageReportS3Bucket(k *SnsSmsPreferencesParameters, p *SnsSmsPreferencesParameters, md *plugin.MergeDescription) bool {
	if k.UsageReportS3Bucket != p.UsageReportS3Bucket {
		p.UsageReportS3Bucket = k.UsageReportS3Bucket
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsSmsPreferences_DefaultSenderId(k *SnsSmsPreferencesParameters, p *SnsSmsPreferencesParameters, md *plugin.MergeDescription) bool {
	if k.DefaultSenderId != p.DefaultSenderId {
		p.DefaultSenderId = k.DefaultSenderId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsSmsPreferences_DefaultSmsType(k *SnsSmsPreferencesParameters, p *SnsSmsPreferencesParameters, md *plugin.MergeDescription) bool {
	if k.DefaultSmsType != p.DefaultSmsType {
		p.DefaultSmsType = k.DefaultSmsType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsSmsPreferences_DeliveryStatusIamRoleArn(k *SnsSmsPreferencesParameters, p *SnsSmsPreferencesParameters, md *plugin.MergeDescription) bool {
	if k.DeliveryStatusIamRoleArn != p.DeliveryStatusIamRoleArn {
		p.DeliveryStatusIamRoleArn = k.DeliveryStatusIamRoleArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsSmsPreferences_DeliveryStatusSuccessSamplingRate(k *SnsSmsPreferencesParameters, p *SnsSmsPreferencesParameters, md *plugin.MergeDescription) bool {
	if k.DeliveryStatusSuccessSamplingRate != p.DeliveryStatusSuccessSamplingRate {
		p.DeliveryStatusSuccessSamplingRate = k.DeliveryStatusSuccessSamplingRate
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}