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
	k := kube.(*BatchJobQueue)
	p := prov.(*BatchJobQueue)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeBatchJobQueue_ComputeEnvironments(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeBatchJobQueue_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeBatchJobQueue_Priority(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeBatchJobQueue_State(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeBatchJobQueue_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeBatchJobQueue_ComputeEnvironments(k *BatchJobQueueParameters, p *BatchJobQueueParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.ComputeEnvironments, p.ComputeEnvironments) {
		p.ComputeEnvironments = k.ComputeEnvironments
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeBatchJobQueue_Name(k *BatchJobQueueParameters, p *BatchJobQueueParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeBatchJobQueue_Priority(k *BatchJobQueueParameters, p *BatchJobQueueParameters, md *plugin.MergeDescription) bool {
	if k.Priority != p.Priority {
		p.Priority = k.Priority
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeBatchJobQueue_State(k *BatchJobQueueParameters, p *BatchJobQueueParameters, md *plugin.MergeDescription) bool {
	if k.State != p.State {
		p.State = k.State
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeBatchJobQueue_Arn(k *BatchJobQueueObservation, p *BatchJobQueueObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}