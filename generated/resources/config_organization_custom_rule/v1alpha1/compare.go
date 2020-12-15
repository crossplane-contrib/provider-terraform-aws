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
	k := kube.(*ConfigOrganizationCustomRule)
	p := prov.(*ConfigOrganizationCustomRule)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeConfigOrganizationCustomRule_ExcludedAccounts(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeConfigOrganizationCustomRule_LambdaFunctionArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeConfigOrganizationCustomRule_ResourceIdScope(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeConfigOrganizationCustomRule_TagValueScope(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeConfigOrganizationCustomRule_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeConfigOrganizationCustomRule_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeConfigOrganizationCustomRule_InputParameters(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeConfigOrganizationCustomRule_MaximumExecutionFrequency(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeConfigOrganizationCustomRule_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeConfigOrganizationCustomRule_ResourceTypesScope(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeConfigOrganizationCustomRule_TagKeyScope(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeConfigOrganizationCustomRule_TriggerTypes(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeConfigOrganizationCustomRule_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeConfigOrganizationCustomRule_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeConfigOrganizationCustomRule_ExcludedAccounts(k *ConfigOrganizationCustomRuleParameters, p *ConfigOrganizationCustomRuleParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(p.ExcludedAccounts, p.ExcludedAccounts) {
		p.ExcludedAccounts = k.ExcludedAccounts
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeConfigOrganizationCustomRule_LambdaFunctionArn(k *ConfigOrganizationCustomRuleParameters, p *ConfigOrganizationCustomRuleParameters, md *plugin.MergeDescription) bool {
	if k.LambdaFunctionArn != p.LambdaFunctionArn {
		p.LambdaFunctionArn = k.LambdaFunctionArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeConfigOrganizationCustomRule_ResourceIdScope(k *ConfigOrganizationCustomRuleParameters, p *ConfigOrganizationCustomRuleParameters, md *plugin.MergeDescription) bool {
	if k.ResourceIdScope != p.ResourceIdScope {
		p.ResourceIdScope = k.ResourceIdScope
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeConfigOrganizationCustomRule_TagValueScope(k *ConfigOrganizationCustomRuleParameters, p *ConfigOrganizationCustomRuleParameters, md *plugin.MergeDescription) bool {
	if k.TagValueScope != p.TagValueScope {
		p.TagValueScope = k.TagValueScope
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeConfigOrganizationCustomRule_Description(k *ConfigOrganizationCustomRuleParameters, p *ConfigOrganizationCustomRuleParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeConfigOrganizationCustomRule_Id(k *ConfigOrganizationCustomRuleParameters, p *ConfigOrganizationCustomRuleParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeConfigOrganizationCustomRule_InputParameters(k *ConfigOrganizationCustomRuleParameters, p *ConfigOrganizationCustomRuleParameters, md *plugin.MergeDescription) bool {
	if k.InputParameters != p.InputParameters {
		p.InputParameters = k.InputParameters
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeConfigOrganizationCustomRule_MaximumExecutionFrequency(k *ConfigOrganizationCustomRuleParameters, p *ConfigOrganizationCustomRuleParameters, md *plugin.MergeDescription) bool {
	if k.MaximumExecutionFrequency != p.MaximumExecutionFrequency {
		p.MaximumExecutionFrequency = k.MaximumExecutionFrequency
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeConfigOrganizationCustomRule_Name(k *ConfigOrganizationCustomRuleParameters, p *ConfigOrganizationCustomRuleParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeConfigOrganizationCustomRule_ResourceTypesScope(k *ConfigOrganizationCustomRuleParameters, p *ConfigOrganizationCustomRuleParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(p.ResourceTypesScope, p.ResourceTypesScope) {
		p.ResourceTypesScope = k.ResourceTypesScope
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeConfigOrganizationCustomRule_TagKeyScope(k *ConfigOrganizationCustomRuleParameters, p *ConfigOrganizationCustomRuleParameters, md *plugin.MergeDescription) bool {
	if k.TagKeyScope != p.TagKeyScope {
		p.TagKeyScope = k.TagKeyScope
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeConfigOrganizationCustomRule_TriggerTypes(k *ConfigOrganizationCustomRuleParameters, p *ConfigOrganizationCustomRuleParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(p.TriggerTypes, p.TriggerTypes) {
		p.TriggerTypes = k.TriggerTypes
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeConfigOrganizationCustomRule_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeConfigOrganizationCustomRule_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeConfigOrganizationCustomRule_Timeouts_Delete(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeConfigOrganizationCustomRule_Timeouts_Update(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeConfigOrganizationCustomRule_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeConfigOrganizationCustomRule_Timeouts_Delete(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Delete != p.Delete {
		p.Delete = k.Delete
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeConfigOrganizationCustomRule_Timeouts_Update(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Update != p.Update {
		p.Update = k.Update
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeConfigOrganizationCustomRule_Arn(k *ConfigOrganizationCustomRuleObservation, p *ConfigOrganizationCustomRuleObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}