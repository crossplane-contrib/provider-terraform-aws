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
	k := kube.(*SsmMaintenanceWindow)
	p := prov.(*SsmMaintenanceWindow)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeSsmMaintenanceWindow_Duration(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSsmMaintenanceWindow_EndDate(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSsmMaintenanceWindow_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSsmMaintenanceWindow_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSsmMaintenanceWindow_Cutoff(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSsmMaintenanceWindow_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSsmMaintenanceWindow_Schedule(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSsmMaintenanceWindow_ScheduleTimezone(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSsmMaintenanceWindow_StartDate(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSsmMaintenanceWindow_AllowUnassociatedTargets(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeSsmMaintenanceWindow_Enabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeSsmMaintenanceWindow_Duration(k *SsmMaintenanceWindowParameters, p *SsmMaintenanceWindowParameters, md *plugin.MergeDescription) bool {
	if k.Duration != p.Duration {
		p.Duration = k.Duration
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSsmMaintenanceWindow_EndDate(k *SsmMaintenanceWindowParameters, p *SsmMaintenanceWindowParameters, md *plugin.MergeDescription) bool {
	if k.EndDate != p.EndDate {
		p.EndDate = k.EndDate
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSsmMaintenanceWindow_Name(k *SsmMaintenanceWindowParameters, p *SsmMaintenanceWindowParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeSsmMaintenanceWindow_Tags(k *SsmMaintenanceWindowParameters, p *SsmMaintenanceWindowParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSsmMaintenanceWindow_Cutoff(k *SsmMaintenanceWindowParameters, p *SsmMaintenanceWindowParameters, md *plugin.MergeDescription) bool {
	if k.Cutoff != p.Cutoff {
		p.Cutoff = k.Cutoff
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSsmMaintenanceWindow_Description(k *SsmMaintenanceWindowParameters, p *SsmMaintenanceWindowParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSsmMaintenanceWindow_Schedule(k *SsmMaintenanceWindowParameters, p *SsmMaintenanceWindowParameters, md *plugin.MergeDescription) bool {
	if k.Schedule != p.Schedule {
		p.Schedule = k.Schedule
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSsmMaintenanceWindow_ScheduleTimezone(k *SsmMaintenanceWindowParameters, p *SsmMaintenanceWindowParameters, md *plugin.MergeDescription) bool {
	if k.ScheduleTimezone != p.ScheduleTimezone {
		p.ScheduleTimezone = k.ScheduleTimezone
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSsmMaintenanceWindow_StartDate(k *SsmMaintenanceWindowParameters, p *SsmMaintenanceWindowParameters, md *plugin.MergeDescription) bool {
	if k.StartDate != p.StartDate {
		p.StartDate = k.StartDate
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSsmMaintenanceWindow_AllowUnassociatedTargets(k *SsmMaintenanceWindowParameters, p *SsmMaintenanceWindowParameters, md *plugin.MergeDescription) bool {
	if k.AllowUnassociatedTargets != p.AllowUnassociatedTargets {
		p.AllowUnassociatedTargets = k.AllowUnassociatedTargets
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeSsmMaintenanceWindow_Enabled(k *SsmMaintenanceWindowParameters, p *SsmMaintenanceWindowParameters, md *plugin.MergeDescription) bool {
	if k.Enabled != p.Enabled {
		p.Enabled = k.Enabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}