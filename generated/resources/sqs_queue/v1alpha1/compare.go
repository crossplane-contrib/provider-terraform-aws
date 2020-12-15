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
	k := kube.(*SqsQueue)
	p := prov.(*SqsQueue)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeSqsQueue_DelaySeconds(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSqsQueue_MessageRetentionSeconds(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSqsQueue_ContentBasedDeduplication(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSqsQueue_MaxMessageSize(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSqsQueue_ReceiveWaitTimeSeconds(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSqsQueue_FifoQueue(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSqsQueue_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSqsQueue_VisibilityTimeoutSeconds(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSqsQueue_KmsDataKeyReusePeriodSeconds(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSqsQueue_KmsMasterKeyId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSqsQueue_NamePrefix(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSqsQueue_Policy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSqsQueue_RedrivePolicy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSqsQueue_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSqsQueue_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeSqsQueue_DelaySeconds(k *SqsQueueParameters, p *SqsQueueParameters, md *plugin.MergeDescription) bool {
	if k.DelaySeconds != p.DelaySeconds {
		p.DelaySeconds = k.DelaySeconds
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSqsQueue_MessageRetentionSeconds(k *SqsQueueParameters, p *SqsQueueParameters, md *plugin.MergeDescription) bool {
	if k.MessageRetentionSeconds != p.MessageRetentionSeconds {
		p.MessageRetentionSeconds = k.MessageRetentionSeconds
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSqsQueue_ContentBasedDeduplication(k *SqsQueueParameters, p *SqsQueueParameters, md *plugin.MergeDescription) bool {
	if k.ContentBasedDeduplication != p.ContentBasedDeduplication {
		p.ContentBasedDeduplication = k.ContentBasedDeduplication
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSqsQueue_MaxMessageSize(k *SqsQueueParameters, p *SqsQueueParameters, md *plugin.MergeDescription) bool {
	if k.MaxMessageSize != p.MaxMessageSize {
		p.MaxMessageSize = k.MaxMessageSize
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSqsQueue_ReceiveWaitTimeSeconds(k *SqsQueueParameters, p *SqsQueueParameters, md *plugin.MergeDescription) bool {
	if k.ReceiveWaitTimeSeconds != p.ReceiveWaitTimeSeconds {
		p.ReceiveWaitTimeSeconds = k.ReceiveWaitTimeSeconds
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSqsQueue_FifoQueue(k *SqsQueueParameters, p *SqsQueueParameters, md *plugin.MergeDescription) bool {
	if k.FifoQueue != p.FifoQueue {
		p.FifoQueue = k.FifoQueue
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSqsQueue_Name(k *SqsQueueParameters, p *SqsQueueParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSqsQueue_VisibilityTimeoutSeconds(k *SqsQueueParameters, p *SqsQueueParameters, md *plugin.MergeDescription) bool {
	if k.VisibilityTimeoutSeconds != p.VisibilityTimeoutSeconds {
		p.VisibilityTimeoutSeconds = k.VisibilityTimeoutSeconds
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSqsQueue_KmsDataKeyReusePeriodSeconds(k *SqsQueueParameters, p *SqsQueueParameters, md *plugin.MergeDescription) bool {
	if k.KmsDataKeyReusePeriodSeconds != p.KmsDataKeyReusePeriodSeconds {
		p.KmsDataKeyReusePeriodSeconds = k.KmsDataKeyReusePeriodSeconds
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSqsQueue_KmsMasterKeyId(k *SqsQueueParameters, p *SqsQueueParameters, md *plugin.MergeDescription) bool {
	if k.KmsMasterKeyId != p.KmsMasterKeyId {
		p.KmsMasterKeyId = k.KmsMasterKeyId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSqsQueue_NamePrefix(k *SqsQueueParameters, p *SqsQueueParameters, md *plugin.MergeDescription) bool {
	if k.NamePrefix != p.NamePrefix {
		p.NamePrefix = k.NamePrefix
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSqsQueue_Policy(k *SqsQueueParameters, p *SqsQueueParameters, md *plugin.MergeDescription) bool {
	if k.Policy != p.Policy {
		p.Policy = k.Policy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSqsQueue_RedrivePolicy(k *SqsQueueParameters, p *SqsQueueParameters, md *plugin.MergeDescription) bool {
	if k.RedrivePolicy != p.RedrivePolicy {
		p.RedrivePolicy = k.RedrivePolicy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeSqsQueue_Tags(k *SqsQueueParameters, p *SqsQueueParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeSqsQueue_Arn(k *SqsQueueObservation, p *SqsQueueObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}