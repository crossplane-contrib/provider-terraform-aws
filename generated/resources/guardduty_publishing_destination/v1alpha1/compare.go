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
	k := kube.(*GuarddutyPublishingDestination)
	p := prov.(*GuarddutyPublishingDestination)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeGuarddutyPublishingDestination_DestinationArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeGuarddutyPublishingDestination_DestinationType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeGuarddutyPublishingDestination_DetectorId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeGuarddutyPublishingDestination_KmsKeyArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeGuarddutyPublishingDestination_DestinationArn(k *GuarddutyPublishingDestinationParameters, p *GuarddutyPublishingDestinationParameters, md *plugin.MergeDescription) bool {
	if k.DestinationArn != p.DestinationArn {
		p.DestinationArn = k.DestinationArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeGuarddutyPublishingDestination_DestinationType(k *GuarddutyPublishingDestinationParameters, p *GuarddutyPublishingDestinationParameters, md *plugin.MergeDescription) bool {
	if k.DestinationType != p.DestinationType {
		p.DestinationType = k.DestinationType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeGuarddutyPublishingDestination_DetectorId(k *GuarddutyPublishingDestinationParameters, p *GuarddutyPublishingDestinationParameters, md *plugin.MergeDescription) bool {
	if k.DetectorId != p.DetectorId {
		p.DetectorId = k.DetectorId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeGuarddutyPublishingDestination_KmsKeyArn(k *GuarddutyPublishingDestinationParameters, p *GuarddutyPublishingDestinationParameters, md *plugin.MergeDescription) bool {
	if k.KmsKeyArn != p.KmsKeyArn {
		p.KmsKeyArn = k.KmsKeyArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}