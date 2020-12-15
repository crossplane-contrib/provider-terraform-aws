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
	k := kube.(*NeptuneEventSubscription)
	p := prov.(*NeptuneEventSubscription)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeNeptuneEventSubscription_NamePrefix(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneEventSubscription_SourceType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneEventSubscription_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneEventSubscription_EventCategories(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneEventSubscription_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneEventSubscription_SnsTopicArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneEventSubscription_SourceIds(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneEventSubscription_Enabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneEventSubscription_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneEventSubscription_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneEventSubscription_CustomerAwsId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneEventSubscription_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeNeptuneEventSubscription_NamePrefix(k *NeptuneEventSubscriptionParameters, p *NeptuneEventSubscriptionParameters, md *plugin.MergeDescription) bool {
	if k.NamePrefix != p.NamePrefix {
		p.NamePrefix = k.NamePrefix
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneEventSubscription_SourceType(k *NeptuneEventSubscriptionParameters, p *NeptuneEventSubscriptionParameters, md *plugin.MergeDescription) bool {
	if k.SourceType != p.SourceType {
		p.SourceType = k.SourceType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeNeptuneEventSubscription_Tags(k *NeptuneEventSubscriptionParameters, p *NeptuneEventSubscriptionParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(p.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeNeptuneEventSubscription_EventCategories(k *NeptuneEventSubscriptionParameters, p *NeptuneEventSubscriptionParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(p.EventCategories, p.EventCategories) {
		p.EventCategories = k.EventCategories
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneEventSubscription_Id(k *NeptuneEventSubscriptionParameters, p *NeptuneEventSubscriptionParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneEventSubscription_SnsTopicArn(k *NeptuneEventSubscriptionParameters, p *NeptuneEventSubscriptionParameters, md *plugin.MergeDescription) bool {
	if k.SnsTopicArn != p.SnsTopicArn {
		p.SnsTopicArn = k.SnsTopicArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeNeptuneEventSubscription_SourceIds(k *NeptuneEventSubscriptionParameters, p *NeptuneEventSubscriptionParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(p.SourceIds, p.SourceIds) {
		p.SourceIds = k.SourceIds
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneEventSubscription_Enabled(k *NeptuneEventSubscriptionParameters, p *NeptuneEventSubscriptionParameters, md *plugin.MergeDescription) bool {
	if k.Enabled != p.Enabled {
		p.Enabled = k.Enabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneEventSubscription_Name(k *NeptuneEventSubscriptionParameters, p *NeptuneEventSubscriptionParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeNeptuneEventSubscription_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeNeptuneEventSubscription_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneEventSubscription_Timeouts_Delete(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNeptuneEventSubscription_Timeouts_Update(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeNeptuneEventSubscription_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneEventSubscription_Timeouts_Delete(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Delete != p.Delete {
		p.Delete = k.Delete
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNeptuneEventSubscription_Timeouts_Update(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Update != p.Update {
		p.Update = k.Update
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNeptuneEventSubscription_CustomerAwsId(k *NeptuneEventSubscriptionObservation, p *NeptuneEventSubscriptionObservation, md *plugin.MergeDescription) bool {
	if k.CustomerAwsId != p.CustomerAwsId {
		k.CustomerAwsId = p.CustomerAwsId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNeptuneEventSubscription_Arn(k *NeptuneEventSubscriptionObservation, p *NeptuneEventSubscriptionObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}