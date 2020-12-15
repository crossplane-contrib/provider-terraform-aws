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
	k := kube.(*CloudformationStack)
	p := prov.(*CloudformationStack)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeCloudformationStack_NotificationArns(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStack_Capabilities(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStack_PolicyBody(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStack_PolicyUrl(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStack_TemplateBody(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStack_TemplateUrl(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStack_DisableRollback(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStack_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStack_OnFailure(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStack_Parameters(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStack_TimeoutInMinutes(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStack_IamRoleArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStack_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStack_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStack_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStack_Outputs(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeCloudformationStack_NotificationArns(k *CloudformationStackParameters, p *CloudformationStackParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(p.NotificationArns, p.NotificationArns) {
		p.NotificationArns = k.NotificationArns
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeCloudformationStack_Capabilities(k *CloudformationStackParameters, p *CloudformationStackParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(p.Capabilities, p.Capabilities) {
		p.Capabilities = k.Capabilities
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudformationStack_PolicyBody(k *CloudformationStackParameters, p *CloudformationStackParameters, md *plugin.MergeDescription) bool {
	if k.PolicyBody != p.PolicyBody {
		p.PolicyBody = k.PolicyBody
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudformationStack_PolicyUrl(k *CloudformationStackParameters, p *CloudformationStackParameters, md *plugin.MergeDescription) bool {
	if k.PolicyUrl != p.PolicyUrl {
		p.PolicyUrl = k.PolicyUrl
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudformationStack_TemplateBody(k *CloudformationStackParameters, p *CloudformationStackParameters, md *plugin.MergeDescription) bool {
	if k.TemplateBody != p.TemplateBody {
		p.TemplateBody = k.TemplateBody
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudformationStack_TemplateUrl(k *CloudformationStackParameters, p *CloudformationStackParameters, md *plugin.MergeDescription) bool {
	if k.TemplateUrl != p.TemplateUrl {
		p.TemplateUrl = k.TemplateUrl
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudformationStack_DisableRollback(k *CloudformationStackParameters, p *CloudformationStackParameters, md *plugin.MergeDescription) bool {
	if k.DisableRollback != p.DisableRollback {
		p.DisableRollback = k.DisableRollback
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudformationStack_Id(k *CloudformationStackParameters, p *CloudformationStackParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudformationStack_OnFailure(k *CloudformationStackParameters, p *CloudformationStackParameters, md *plugin.MergeDescription) bool {
	if k.OnFailure != p.OnFailure {
		p.OnFailure = k.OnFailure
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeCloudformationStack_Parameters(k *CloudformationStackParameters, p *CloudformationStackParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(p.Parameters, p.Parameters) {
		p.Parameters = k.Parameters
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudformationStack_TimeoutInMinutes(k *CloudformationStackParameters, p *CloudformationStackParameters, md *plugin.MergeDescription) bool {
	if k.TimeoutInMinutes != p.TimeoutInMinutes {
		p.TimeoutInMinutes = k.TimeoutInMinutes
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudformationStack_IamRoleArn(k *CloudformationStackParameters, p *CloudformationStackParameters, md *plugin.MergeDescription) bool {
	if k.IamRoleArn != p.IamRoleArn {
		p.IamRoleArn = k.IamRoleArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudformationStack_Name(k *CloudformationStackParameters, p *CloudformationStackParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeCloudformationStack_Tags(k *CloudformationStackParameters, p *CloudformationStackParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(p.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeCloudformationStack_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeCloudformationStack_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStack_Timeouts_Delete(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudformationStack_Timeouts_Update(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeCloudformationStack_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudformationStack_Timeouts_Delete(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Delete != p.Delete {
		p.Delete = k.Delete
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudformationStack_Timeouts_Update(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Update != p.Update {
		p.Update = k.Update
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateStatus
func MergeCloudformationStack_Outputs(k *CloudformationStackObservation, p *CloudformationStackObservation, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(p.Outputs, p.Outputs) {
		k.Outputs = p.Outputs
		md.StatusUpdated = true
		return true
	}
	return false
}