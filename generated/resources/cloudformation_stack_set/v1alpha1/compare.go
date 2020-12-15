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
	k := kube.(*CloudformationStackSet)
	p := prov.(*CloudformationStackSet)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeCloudformationStackSet_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStackSet_Capabilities(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStackSet_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStackSet_Parameters(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStackSet_TemplateBody(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStackSet_TemplateUrl(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStackSet_AdministrationRoleArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStackSet_ExecutionRoleName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStackSet_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStackSet_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStackSet_StackSetId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStackSet_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeCloudformationStackSet_Tags(k *CloudformationStackSetParameters, p *CloudformationStackSetParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeCloudformationStackSet_Capabilities(k *CloudformationStackSetParameters, p *CloudformationStackSetParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.Capabilities, p.Capabilities) {
		p.Capabilities = k.Capabilities
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudformationStackSet_Description(k *CloudformationStackSetParameters, p *CloudformationStackSetParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeCloudformationStackSet_Parameters(k *CloudformationStackSetParameters, p *CloudformationStackSetParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Parameters, p.Parameters) {
		p.Parameters = k.Parameters
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudformationStackSet_TemplateBody(k *CloudformationStackSetParameters, p *CloudformationStackSetParameters, md *plugin.MergeDescription) bool {
	if k.TemplateBody != p.TemplateBody {
		p.TemplateBody = k.TemplateBody
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudformationStackSet_TemplateUrl(k *CloudformationStackSetParameters, p *CloudformationStackSetParameters, md *plugin.MergeDescription) bool {
	if k.TemplateUrl != p.TemplateUrl {
		p.TemplateUrl = k.TemplateUrl
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudformationStackSet_AdministrationRoleArn(k *CloudformationStackSetParameters, p *CloudformationStackSetParameters, md *plugin.MergeDescription) bool {
	if k.AdministrationRoleArn != p.AdministrationRoleArn {
		p.AdministrationRoleArn = k.AdministrationRoleArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudformationStackSet_ExecutionRoleName(k *CloudformationStackSetParameters, p *CloudformationStackSetParameters, md *plugin.MergeDescription) bool {
	if k.ExecutionRoleName != p.ExecutionRoleName {
		p.ExecutionRoleName = k.ExecutionRoleName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudformationStackSet_Name(k *CloudformationStackSetParameters, p *CloudformationStackSetParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeCloudformationStackSet_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeCloudformationStackSet_Timeouts_Update(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeCloudformationStackSet_Timeouts_Update(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Update != p.Update {
		p.Update = k.Update
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCloudformationStackSet_StackSetId(k *CloudformationStackSetObservation, p *CloudformationStackSetObservation, md *plugin.MergeDescription) bool {
	if k.StackSetId != p.StackSetId {
		k.StackSetId = p.StackSetId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCloudformationStackSet_Arn(k *CloudformationStackSetObservation, p *CloudformationStackSetObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}