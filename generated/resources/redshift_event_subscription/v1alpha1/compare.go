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
	k := kube.(*RedshiftEventSubscription)
	p := prov.(*RedshiftEventSubscription)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeRedshiftEventSubscription_SourceIds(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRedshiftEventSubscription_SourceType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRedshiftEventSubscription_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRedshiftEventSubscription_Enabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRedshiftEventSubscription_EventCategories(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRedshiftEventSubscription_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRedshiftEventSubscription_Severity(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRedshiftEventSubscription_SnsTopicArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRedshiftEventSubscription_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRedshiftEventSubscription_CustomerAwsId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRedshiftEventSubscription_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRedshiftEventSubscription_Status(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeRedshiftEventSubscription_SourceIds(k *RedshiftEventSubscriptionParameters, p *RedshiftEventSubscriptionParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.SourceIds, p.SourceIds) {
		p.SourceIds = k.SourceIds
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRedshiftEventSubscription_SourceType(k *RedshiftEventSubscriptionParameters, p *RedshiftEventSubscriptionParameters, md *plugin.MergeDescription) bool {
	if k.SourceType != p.SourceType {
		p.SourceType = k.SourceType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeRedshiftEventSubscription_Tags(k *RedshiftEventSubscriptionParameters, p *RedshiftEventSubscriptionParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRedshiftEventSubscription_Enabled(k *RedshiftEventSubscriptionParameters, p *RedshiftEventSubscriptionParameters, md *plugin.MergeDescription) bool {
	if k.Enabled != p.Enabled {
		p.Enabled = k.Enabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeRedshiftEventSubscription_EventCategories(k *RedshiftEventSubscriptionParameters, p *RedshiftEventSubscriptionParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.EventCategories, p.EventCategories) {
		p.EventCategories = k.EventCategories
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRedshiftEventSubscription_Name(k *RedshiftEventSubscriptionParameters, p *RedshiftEventSubscriptionParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRedshiftEventSubscription_Severity(k *RedshiftEventSubscriptionParameters, p *RedshiftEventSubscriptionParameters, md *plugin.MergeDescription) bool {
	if k.Severity != p.Severity {
		p.Severity = k.Severity
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRedshiftEventSubscription_SnsTopicArn(k *RedshiftEventSubscriptionParameters, p *RedshiftEventSubscriptionParameters, md *plugin.MergeDescription) bool {
	if k.SnsTopicArn != p.SnsTopicArn {
		p.SnsTopicArn = k.SnsTopicArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeRedshiftEventSubscription_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeRedshiftEventSubscription_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRedshiftEventSubscription_Timeouts_Delete(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRedshiftEventSubscription_Timeouts_Update(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeRedshiftEventSubscription_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRedshiftEventSubscription_Timeouts_Delete(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Delete != p.Delete {
		p.Delete = k.Delete
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRedshiftEventSubscription_Timeouts_Update(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Update != p.Update {
		p.Update = k.Update
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeRedshiftEventSubscription_CustomerAwsId(k *RedshiftEventSubscriptionObservation, p *RedshiftEventSubscriptionObservation, md *plugin.MergeDescription) bool {
	if k.CustomerAwsId != p.CustomerAwsId {
		k.CustomerAwsId = p.CustomerAwsId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeRedshiftEventSubscription_Arn(k *RedshiftEventSubscriptionObservation, p *RedshiftEventSubscriptionObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeRedshiftEventSubscription_Status(k *RedshiftEventSubscriptionObservation, p *RedshiftEventSubscriptionObservation, md *plugin.MergeDescription) bool {
	if k.Status != p.Status {
		k.Status = p.Status
		md.StatusUpdated = true
		return true
	}
	return false
}