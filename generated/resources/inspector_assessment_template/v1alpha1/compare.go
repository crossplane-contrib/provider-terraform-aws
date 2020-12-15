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
	k := kube.(*InspectorAssessmentTemplate)
	p := prov.(*InspectorAssessmentTemplate)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeInspectorAssessmentTemplate_TargetArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeInspectorAssessmentTemplate_Duration(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeInspectorAssessmentTemplate_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeInspectorAssessmentTemplate_RulesPackageArns(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeInspectorAssessmentTemplate_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeInspectorAssessmentTemplate_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeInspectorAssessmentTemplate_TargetArn(k *InspectorAssessmentTemplateParameters, p *InspectorAssessmentTemplateParameters, md *plugin.MergeDescription) bool {
	if k.TargetArn != p.TargetArn {
		p.TargetArn = k.TargetArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeInspectorAssessmentTemplate_Duration(k *InspectorAssessmentTemplateParameters, p *InspectorAssessmentTemplateParameters, md *plugin.MergeDescription) bool {
	if k.Duration != p.Duration {
		p.Duration = k.Duration
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeInspectorAssessmentTemplate_Name(k *InspectorAssessmentTemplateParameters, p *InspectorAssessmentTemplateParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeInspectorAssessmentTemplate_RulesPackageArns(k *InspectorAssessmentTemplateParameters, p *InspectorAssessmentTemplateParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.RulesPackageArns, p.RulesPackageArns) {
		p.RulesPackageArns = k.RulesPackageArns
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeInspectorAssessmentTemplate_Tags(k *InspectorAssessmentTemplateParameters, p *InspectorAssessmentTemplateParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeInspectorAssessmentTemplate_Arn(k *InspectorAssessmentTemplateObservation, p *InspectorAssessmentTemplateObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}