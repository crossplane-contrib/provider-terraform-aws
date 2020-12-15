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
	k := kube.(*XraySamplingRule)
	p := prov.(*XraySamplingRule)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeXraySamplingRule_Attributes(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeXraySamplingRule_HttpMethod(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeXraySamplingRule_Version(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeXraySamplingRule_Host(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeXraySamplingRule_FixedRate(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeXraySamplingRule_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeXraySamplingRule_ReservoirSize(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeXraySamplingRule_ServiceType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeXraySamplingRule_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeXraySamplingRule_Priority(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeXraySamplingRule_ResourceArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeXraySamplingRule_RuleName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeXraySamplingRule_ServiceName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeXraySamplingRule_UrlPath(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeXraySamplingRule_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeXraySamplingRule_Attributes(k *XraySamplingRuleParameters, p *XraySamplingRuleParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(p.Attributes, p.Attributes) {
		p.Attributes = k.Attributes
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeXraySamplingRule_HttpMethod(k *XraySamplingRuleParameters, p *XraySamplingRuleParameters, md *plugin.MergeDescription) bool {
	if k.HttpMethod != p.HttpMethod {
		p.HttpMethod = k.HttpMethod
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeXraySamplingRule_Version(k *XraySamplingRuleParameters, p *XraySamplingRuleParameters, md *plugin.MergeDescription) bool {
	if k.Version != p.Version {
		p.Version = k.Version
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeXraySamplingRule_Host(k *XraySamplingRuleParameters, p *XraySamplingRuleParameters, md *plugin.MergeDescription) bool {
	if k.Host != p.Host {
		p.Host = k.Host
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeXraySamplingRule_FixedRate(k *XraySamplingRuleParameters, p *XraySamplingRuleParameters, md *plugin.MergeDescription) bool {
	if k.FixedRate != p.FixedRate {
		p.FixedRate = k.FixedRate
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeXraySamplingRule_Id(k *XraySamplingRuleParameters, p *XraySamplingRuleParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeXraySamplingRule_ReservoirSize(k *XraySamplingRuleParameters, p *XraySamplingRuleParameters, md *plugin.MergeDescription) bool {
	if k.ReservoirSize != p.ReservoirSize {
		p.ReservoirSize = k.ReservoirSize
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeXraySamplingRule_ServiceType(k *XraySamplingRuleParameters, p *XraySamplingRuleParameters, md *plugin.MergeDescription) bool {
	if k.ServiceType != p.ServiceType {
		p.ServiceType = k.ServiceType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeXraySamplingRule_Tags(k *XraySamplingRuleParameters, p *XraySamplingRuleParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(p.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeXraySamplingRule_Priority(k *XraySamplingRuleParameters, p *XraySamplingRuleParameters, md *plugin.MergeDescription) bool {
	if k.Priority != p.Priority {
		p.Priority = k.Priority
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeXraySamplingRule_ResourceArn(k *XraySamplingRuleParameters, p *XraySamplingRuleParameters, md *plugin.MergeDescription) bool {
	if k.ResourceArn != p.ResourceArn {
		p.ResourceArn = k.ResourceArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeXraySamplingRule_RuleName(k *XraySamplingRuleParameters, p *XraySamplingRuleParameters, md *plugin.MergeDescription) bool {
	if k.RuleName != p.RuleName {
		p.RuleName = k.RuleName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeXraySamplingRule_ServiceName(k *XraySamplingRuleParameters, p *XraySamplingRuleParameters, md *plugin.MergeDescription) bool {
	if k.ServiceName != p.ServiceName {
		p.ServiceName = k.ServiceName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeXraySamplingRule_UrlPath(k *XraySamplingRuleParameters, p *XraySamplingRuleParameters, md *plugin.MergeDescription) bool {
	if k.UrlPath != p.UrlPath {
		p.UrlPath = k.UrlPath
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeXraySamplingRule_Arn(k *XraySamplingRuleObservation, p *XraySamplingRuleObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}