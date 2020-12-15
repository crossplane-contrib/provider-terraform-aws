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
	k := kube.(*SnsTopic)
	p := prov.(*SnsTopic)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeSnsTopic_LambdaFailureFeedbackRoleArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopic_LambdaSuccessFeedbackRoleArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopic_SqsSuccessFeedbackRoleArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopic_SqsSuccessFeedbackSampleRate(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopic_DisplayName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopic_KmsMasterKeyId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopic_LambdaSuccessFeedbackSampleRate(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopic_NamePrefix(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopic_Policy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopic_ApplicationSuccessFeedbackSampleRate(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopic_HttpSuccessFeedbackRoleArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopic_HttpSuccessFeedbackSampleRate(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopic_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopic_ApplicationSuccessFeedbackRoleArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopic_DeliveryPolicy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopic_HttpFailureFeedbackRoleArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopic_SqsFailureFeedbackRoleArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopic_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopic_ApplicationFailureFeedbackRoleArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsTopic_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeSnsTopic_LambdaFailureFeedbackRoleArn(k *SnsTopicParameters, p *SnsTopicParameters, md *plugin.MergeDescription) bool {
	if k.LambdaFailureFeedbackRoleArn != p.LambdaFailureFeedbackRoleArn {
		p.LambdaFailureFeedbackRoleArn = k.LambdaFailureFeedbackRoleArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsTopic_LambdaSuccessFeedbackRoleArn(k *SnsTopicParameters, p *SnsTopicParameters, md *plugin.MergeDescription) bool {
	if k.LambdaSuccessFeedbackRoleArn != p.LambdaSuccessFeedbackRoleArn {
		p.LambdaSuccessFeedbackRoleArn = k.LambdaSuccessFeedbackRoleArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsTopic_SqsSuccessFeedbackRoleArn(k *SnsTopicParameters, p *SnsTopicParameters, md *plugin.MergeDescription) bool {
	if k.SqsSuccessFeedbackRoleArn != p.SqsSuccessFeedbackRoleArn {
		p.SqsSuccessFeedbackRoleArn = k.SqsSuccessFeedbackRoleArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsTopic_SqsSuccessFeedbackSampleRate(k *SnsTopicParameters, p *SnsTopicParameters, md *plugin.MergeDescription) bool {
	if k.SqsSuccessFeedbackSampleRate != p.SqsSuccessFeedbackSampleRate {
		p.SqsSuccessFeedbackSampleRate = k.SqsSuccessFeedbackSampleRate
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsTopic_DisplayName(k *SnsTopicParameters, p *SnsTopicParameters, md *plugin.MergeDescription) bool {
	if k.DisplayName != p.DisplayName {
		p.DisplayName = k.DisplayName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsTopic_KmsMasterKeyId(k *SnsTopicParameters, p *SnsTopicParameters, md *plugin.MergeDescription) bool {
	if k.KmsMasterKeyId != p.KmsMasterKeyId {
		p.KmsMasterKeyId = k.KmsMasterKeyId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsTopic_LambdaSuccessFeedbackSampleRate(k *SnsTopicParameters, p *SnsTopicParameters, md *plugin.MergeDescription) bool {
	if k.LambdaSuccessFeedbackSampleRate != p.LambdaSuccessFeedbackSampleRate {
		p.LambdaSuccessFeedbackSampleRate = k.LambdaSuccessFeedbackSampleRate
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsTopic_NamePrefix(k *SnsTopicParameters, p *SnsTopicParameters, md *plugin.MergeDescription) bool {
	if k.NamePrefix != p.NamePrefix {
		p.NamePrefix = k.NamePrefix
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsTopic_Policy(k *SnsTopicParameters, p *SnsTopicParameters, md *plugin.MergeDescription) bool {
	if k.Policy != p.Policy {
		p.Policy = k.Policy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsTopic_ApplicationSuccessFeedbackSampleRate(k *SnsTopicParameters, p *SnsTopicParameters, md *plugin.MergeDescription) bool {
	if k.ApplicationSuccessFeedbackSampleRate != p.ApplicationSuccessFeedbackSampleRate {
		p.ApplicationSuccessFeedbackSampleRate = k.ApplicationSuccessFeedbackSampleRate
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsTopic_HttpSuccessFeedbackRoleArn(k *SnsTopicParameters, p *SnsTopicParameters, md *plugin.MergeDescription) bool {
	if k.HttpSuccessFeedbackRoleArn != p.HttpSuccessFeedbackRoleArn {
		p.HttpSuccessFeedbackRoleArn = k.HttpSuccessFeedbackRoleArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsTopic_HttpSuccessFeedbackSampleRate(k *SnsTopicParameters, p *SnsTopicParameters, md *plugin.MergeDescription) bool {
	if k.HttpSuccessFeedbackSampleRate != p.HttpSuccessFeedbackSampleRate {
		p.HttpSuccessFeedbackSampleRate = k.HttpSuccessFeedbackSampleRate
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsTopic_Name(k *SnsTopicParameters, p *SnsTopicParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsTopic_ApplicationSuccessFeedbackRoleArn(k *SnsTopicParameters, p *SnsTopicParameters, md *plugin.MergeDescription) bool {
	if k.ApplicationSuccessFeedbackRoleArn != p.ApplicationSuccessFeedbackRoleArn {
		p.ApplicationSuccessFeedbackRoleArn = k.ApplicationSuccessFeedbackRoleArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsTopic_DeliveryPolicy(k *SnsTopicParameters, p *SnsTopicParameters, md *plugin.MergeDescription) bool {
	if k.DeliveryPolicy != p.DeliveryPolicy {
		p.DeliveryPolicy = k.DeliveryPolicy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsTopic_HttpFailureFeedbackRoleArn(k *SnsTopicParameters, p *SnsTopicParameters, md *plugin.MergeDescription) bool {
	if k.HttpFailureFeedbackRoleArn != p.HttpFailureFeedbackRoleArn {
		p.HttpFailureFeedbackRoleArn = k.HttpFailureFeedbackRoleArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsTopic_SqsFailureFeedbackRoleArn(k *SnsTopicParameters, p *SnsTopicParameters, md *plugin.MergeDescription) bool {
	if k.SqsFailureFeedbackRoleArn != p.SqsFailureFeedbackRoleArn {
		p.SqsFailureFeedbackRoleArn = k.SqsFailureFeedbackRoleArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeSnsTopic_Tags(k *SnsTopicParameters, p *SnsTopicParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsTopic_ApplicationFailureFeedbackRoleArn(k *SnsTopicParameters, p *SnsTopicParameters, md *plugin.MergeDescription) bool {
	if k.ApplicationFailureFeedbackRoleArn != p.ApplicationFailureFeedbackRoleArn {
		p.ApplicationFailureFeedbackRoleArn = k.ApplicationFailureFeedbackRoleArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeSnsTopic_Arn(k *SnsTopicObservation, p *SnsTopicObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}