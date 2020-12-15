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
	k := kube.(*SagemakerNotebookInstance)
	p := prov.(*SagemakerNotebookInstance)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeSagemakerNotebookInstance_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSagemakerNotebookInstance_RootAccess(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSagemakerNotebookInstance_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSagemakerNotebookInstance_SubnetId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSagemakerNotebookInstance_DirectInternetAccess(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSagemakerNotebookInstance_InstanceType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSagemakerNotebookInstance_KmsKeyId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSagemakerNotebookInstance_LifecycleConfigName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSagemakerNotebookInstance_RoleArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSagemakerNotebookInstance_SecurityGroups(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSagemakerNotebookInstance_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeSagemakerNotebookInstance_Name(k *SagemakerNotebookInstanceParameters, p *SagemakerNotebookInstanceParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSagemakerNotebookInstance_RootAccess(k *SagemakerNotebookInstanceParameters, p *SagemakerNotebookInstanceParameters, md *plugin.MergeDescription) bool {
	if k.RootAccess != p.RootAccess {
		p.RootAccess = k.RootAccess
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeSagemakerNotebookInstance_Tags(k *SagemakerNotebookInstanceParameters, p *SagemakerNotebookInstanceParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSagemakerNotebookInstance_SubnetId(k *SagemakerNotebookInstanceParameters, p *SagemakerNotebookInstanceParameters, md *plugin.MergeDescription) bool {
	if k.SubnetId != p.SubnetId {
		p.SubnetId = k.SubnetId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSagemakerNotebookInstance_DirectInternetAccess(k *SagemakerNotebookInstanceParameters, p *SagemakerNotebookInstanceParameters, md *plugin.MergeDescription) bool {
	if k.DirectInternetAccess != p.DirectInternetAccess {
		p.DirectInternetAccess = k.DirectInternetAccess
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSagemakerNotebookInstance_InstanceType(k *SagemakerNotebookInstanceParameters, p *SagemakerNotebookInstanceParameters, md *plugin.MergeDescription) bool {
	if k.InstanceType != p.InstanceType {
		p.InstanceType = k.InstanceType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSagemakerNotebookInstance_KmsKeyId(k *SagemakerNotebookInstanceParameters, p *SagemakerNotebookInstanceParameters, md *plugin.MergeDescription) bool {
	if k.KmsKeyId != p.KmsKeyId {
		p.KmsKeyId = k.KmsKeyId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSagemakerNotebookInstance_LifecycleConfigName(k *SagemakerNotebookInstanceParameters, p *SagemakerNotebookInstanceParameters, md *plugin.MergeDescription) bool {
	if k.LifecycleConfigName != p.LifecycleConfigName {
		p.LifecycleConfigName = k.LifecycleConfigName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSagemakerNotebookInstance_RoleArn(k *SagemakerNotebookInstanceParameters, p *SagemakerNotebookInstanceParameters, md *plugin.MergeDescription) bool {
	if k.RoleArn != p.RoleArn {
		p.RoleArn = k.RoleArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeSagemakerNotebookInstance_SecurityGroups(k *SagemakerNotebookInstanceParameters, p *SagemakerNotebookInstanceParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.SecurityGroups, p.SecurityGroups) {
		p.SecurityGroups = k.SecurityGroups
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeSagemakerNotebookInstance_Arn(k *SagemakerNotebookInstanceObservation, p *SagemakerNotebookInstanceObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}