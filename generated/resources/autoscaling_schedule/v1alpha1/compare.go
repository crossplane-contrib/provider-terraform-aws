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
	k := kube.(*AutoscalingSchedule)
	p := prov.(*AutoscalingSchedule)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeAutoscalingSchedule_MinSize(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAutoscalingSchedule_Recurrence(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAutoscalingSchedule_StartTime(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAutoscalingSchedule_EndTime(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAutoscalingSchedule_MaxSize(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAutoscalingSchedule_ScheduledActionName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAutoscalingSchedule_AutoscalingGroupName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAutoscalingSchedule_DesiredCapacity(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAutoscalingSchedule_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeAutoscalingSchedule_MinSize(k *AutoscalingScheduleParameters, p *AutoscalingScheduleParameters, md *plugin.MergeDescription) bool {
	if k.MinSize != p.MinSize {
		p.MinSize = k.MinSize
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAutoscalingSchedule_Recurrence(k *AutoscalingScheduleParameters, p *AutoscalingScheduleParameters, md *plugin.MergeDescription) bool {
	if k.Recurrence != p.Recurrence {
		p.Recurrence = k.Recurrence
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAutoscalingSchedule_StartTime(k *AutoscalingScheduleParameters, p *AutoscalingScheduleParameters, md *plugin.MergeDescription) bool {
	if k.StartTime != p.StartTime {
		p.StartTime = k.StartTime
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAutoscalingSchedule_EndTime(k *AutoscalingScheduleParameters, p *AutoscalingScheduleParameters, md *plugin.MergeDescription) bool {
	if k.EndTime != p.EndTime {
		p.EndTime = k.EndTime
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAutoscalingSchedule_MaxSize(k *AutoscalingScheduleParameters, p *AutoscalingScheduleParameters, md *plugin.MergeDescription) bool {
	if k.MaxSize != p.MaxSize {
		p.MaxSize = k.MaxSize
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAutoscalingSchedule_ScheduledActionName(k *AutoscalingScheduleParameters, p *AutoscalingScheduleParameters, md *plugin.MergeDescription) bool {
	if k.ScheduledActionName != p.ScheduledActionName {
		p.ScheduledActionName = k.ScheduledActionName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAutoscalingSchedule_AutoscalingGroupName(k *AutoscalingScheduleParameters, p *AutoscalingScheduleParameters, md *plugin.MergeDescription) bool {
	if k.AutoscalingGroupName != p.AutoscalingGroupName {
		p.AutoscalingGroupName = k.AutoscalingGroupName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAutoscalingSchedule_DesiredCapacity(k *AutoscalingScheduleParameters, p *AutoscalingScheduleParameters, md *plugin.MergeDescription) bool {
	if k.DesiredCapacity != p.DesiredCapacity {
		p.DesiredCapacity = k.DesiredCapacity
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeAutoscalingSchedule_Arn(k *AutoscalingScheduleObservation, p *AutoscalingScheduleObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}