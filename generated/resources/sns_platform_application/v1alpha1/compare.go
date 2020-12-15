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
	k := kube.(*SnsPlatformApplication)
	p := prov.(*SnsPlatformApplication)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeSnsPlatformApplication_EventDeliveryFailureTopicArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsPlatformApplication_EventEndpointUpdatedTopicArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsPlatformApplication_FailureFeedbackRoleArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsPlatformApplication_Platform(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsPlatformApplication_PlatformCredential(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsPlatformApplication_SuccessFeedbackSampleRate(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsPlatformApplication_EventEndpointCreatedTopicArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsPlatformApplication_EventEndpointDeletedTopicArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsPlatformApplication_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsPlatformApplication_PlatformPrincipal(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsPlatformApplication_SuccessFeedbackRoleArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSnsPlatformApplication_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeSnsPlatformApplication_EventDeliveryFailureTopicArn(k *SnsPlatformApplicationParameters, p *SnsPlatformApplicationParameters, md *plugin.MergeDescription) bool {
	if k.EventDeliveryFailureTopicArn != p.EventDeliveryFailureTopicArn {
		p.EventDeliveryFailureTopicArn = k.EventDeliveryFailureTopicArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsPlatformApplication_EventEndpointUpdatedTopicArn(k *SnsPlatformApplicationParameters, p *SnsPlatformApplicationParameters, md *plugin.MergeDescription) bool {
	if k.EventEndpointUpdatedTopicArn != p.EventEndpointUpdatedTopicArn {
		p.EventEndpointUpdatedTopicArn = k.EventEndpointUpdatedTopicArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsPlatformApplication_FailureFeedbackRoleArn(k *SnsPlatformApplicationParameters, p *SnsPlatformApplicationParameters, md *plugin.MergeDescription) bool {
	if k.FailureFeedbackRoleArn != p.FailureFeedbackRoleArn {
		p.FailureFeedbackRoleArn = k.FailureFeedbackRoleArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsPlatformApplication_Platform(k *SnsPlatformApplicationParameters, p *SnsPlatformApplicationParameters, md *plugin.MergeDescription) bool {
	if k.Platform != p.Platform {
		p.Platform = k.Platform
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsPlatformApplication_PlatformCredential(k *SnsPlatformApplicationParameters, p *SnsPlatformApplicationParameters, md *plugin.MergeDescription) bool {
	if k.PlatformCredential != p.PlatformCredential {
		p.PlatformCredential = k.PlatformCredential
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsPlatformApplication_SuccessFeedbackSampleRate(k *SnsPlatformApplicationParameters, p *SnsPlatformApplicationParameters, md *plugin.MergeDescription) bool {
	if k.SuccessFeedbackSampleRate != p.SuccessFeedbackSampleRate {
		p.SuccessFeedbackSampleRate = k.SuccessFeedbackSampleRate
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsPlatformApplication_EventEndpointCreatedTopicArn(k *SnsPlatformApplicationParameters, p *SnsPlatformApplicationParameters, md *plugin.MergeDescription) bool {
	if k.EventEndpointCreatedTopicArn != p.EventEndpointCreatedTopicArn {
		p.EventEndpointCreatedTopicArn = k.EventEndpointCreatedTopicArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsPlatformApplication_EventEndpointDeletedTopicArn(k *SnsPlatformApplicationParameters, p *SnsPlatformApplicationParameters, md *plugin.MergeDescription) bool {
	if k.EventEndpointDeletedTopicArn != p.EventEndpointDeletedTopicArn {
		p.EventEndpointDeletedTopicArn = k.EventEndpointDeletedTopicArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsPlatformApplication_Name(k *SnsPlatformApplicationParameters, p *SnsPlatformApplicationParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsPlatformApplication_PlatformPrincipal(k *SnsPlatformApplicationParameters, p *SnsPlatformApplicationParameters, md *plugin.MergeDescription) bool {
	if k.PlatformPrincipal != p.PlatformPrincipal {
		p.PlatformPrincipal = k.PlatformPrincipal
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSnsPlatformApplication_SuccessFeedbackRoleArn(k *SnsPlatformApplicationParameters, p *SnsPlatformApplicationParameters, md *plugin.MergeDescription) bool {
	if k.SuccessFeedbackRoleArn != p.SuccessFeedbackRoleArn {
		p.SuccessFeedbackRoleArn = k.SuccessFeedbackRoleArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeSnsPlatformApplication_Arn(k *SnsPlatformApplicationObservation, p *SnsPlatformApplicationObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}