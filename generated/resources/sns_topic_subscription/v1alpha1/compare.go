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
	k := kube.(*SnsTopicSubscription)
	p := prov.(*SnsTopicSubscription)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeSnsTopicSubscription_Endpoint(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopicSubscription_EndpointAutoConfirms(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopicSubscription_FilterPolicy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopicSubscription_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopicSubscription_Protocol(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopicSubscription_RawMessageDelivery(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopicSubscription_ConfirmationTimeoutInMinutes(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopicSubscription_TopicArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopicSubscription_DeliveryPolicy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopicSubscription_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeSnsTopicSubscription_Endpoint(k *SnsTopicSubscriptionParameters, p *SnsTopicSubscriptionParameters, md *plugin.MergeDescription) bool {
	if k.Endpoint != p.Endpoint {
		p.Endpoint = k.Endpoint
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsTopicSubscription_EndpointAutoConfirms(k *SnsTopicSubscriptionParameters, p *SnsTopicSubscriptionParameters, md *plugin.MergeDescription) bool {
	if k.EndpointAutoConfirms != p.EndpointAutoConfirms {
		p.EndpointAutoConfirms = k.EndpointAutoConfirms
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsTopicSubscription_FilterPolicy(k *SnsTopicSubscriptionParameters, p *SnsTopicSubscriptionParameters, md *plugin.MergeDescription) bool {
	if k.FilterPolicy != p.FilterPolicy {
		p.FilterPolicy = k.FilterPolicy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsTopicSubscription_Id(k *SnsTopicSubscriptionParameters, p *SnsTopicSubscriptionParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsTopicSubscription_Protocol(k *SnsTopicSubscriptionParameters, p *SnsTopicSubscriptionParameters, md *plugin.MergeDescription) bool {
	if k.Protocol != p.Protocol {
		p.Protocol = k.Protocol
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsTopicSubscription_RawMessageDelivery(k *SnsTopicSubscriptionParameters, p *SnsTopicSubscriptionParameters, md *plugin.MergeDescription) bool {
	if k.RawMessageDelivery != p.RawMessageDelivery {
		p.RawMessageDelivery = k.RawMessageDelivery
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsTopicSubscription_ConfirmationTimeoutInMinutes(k *SnsTopicSubscriptionParameters, p *SnsTopicSubscriptionParameters, md *plugin.MergeDescription) bool {
	if k.ConfirmationTimeoutInMinutes != p.ConfirmationTimeoutInMinutes {
		p.ConfirmationTimeoutInMinutes = k.ConfirmationTimeoutInMinutes
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsTopicSubscription_TopicArn(k *SnsTopicSubscriptionParameters, p *SnsTopicSubscriptionParameters, md *plugin.MergeDescription) bool {
	if k.TopicArn != p.TopicArn {
		p.TopicArn = k.TopicArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsTopicSubscription_DeliveryPolicy(k *SnsTopicSubscriptionParameters, p *SnsTopicSubscriptionParameters, md *plugin.MergeDescription) bool {
	if k.DeliveryPolicy != p.DeliveryPolicy {
		p.DeliveryPolicy = k.DeliveryPolicy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeSnsTopicSubscription_Arn(k *SnsTopicSubscriptionObservation, p *SnsTopicSubscriptionObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}