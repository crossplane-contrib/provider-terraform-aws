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
	k := kube.(*RedshiftSnapshotSchedule)
	p := prov.(*RedshiftSnapshotSchedule)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeRedshiftSnapshotSchedule_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRedshiftSnapshotSchedule_Identifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRedshiftSnapshotSchedule_IdentifierPrefix(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRedshiftSnapshotSchedule_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRedshiftSnapshotSchedule_Definitions(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRedshiftSnapshotSchedule_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRedshiftSnapshotSchedule_ForceDestroy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeRedshiftSnapshotSchedule_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeRedshiftSnapshotSchedule_Id(k *RedshiftSnapshotScheduleParameters, p *RedshiftSnapshotScheduleParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRedshiftSnapshotSchedule_Identifier(k *RedshiftSnapshotScheduleParameters, p *RedshiftSnapshotScheduleParameters, md *plugin.MergeDescription) bool {
	if k.Identifier != p.Identifier {
		p.Identifier = k.Identifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRedshiftSnapshotSchedule_IdentifierPrefix(k *RedshiftSnapshotScheduleParameters, p *RedshiftSnapshotScheduleParameters, md *plugin.MergeDescription) bool {
	if k.IdentifierPrefix != p.IdentifierPrefix {
		p.IdentifierPrefix = k.IdentifierPrefix
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeRedshiftSnapshotSchedule_Tags(k *RedshiftSnapshotScheduleParameters, p *RedshiftSnapshotScheduleParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(p.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeRedshiftSnapshotSchedule_Definitions(k *RedshiftSnapshotScheduleParameters, p *RedshiftSnapshotScheduleParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(p.Definitions, p.Definitions) {
		p.Definitions = k.Definitions
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRedshiftSnapshotSchedule_Description(k *RedshiftSnapshotScheduleParameters, p *RedshiftSnapshotScheduleParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeRedshiftSnapshotSchedule_ForceDestroy(k *RedshiftSnapshotScheduleParameters, p *RedshiftSnapshotScheduleParameters, md *plugin.MergeDescription) bool {
	if k.ForceDestroy != p.ForceDestroy {
		p.ForceDestroy = k.ForceDestroy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeRedshiftSnapshotSchedule_Arn(k *RedshiftSnapshotScheduleObservation, p *RedshiftSnapshotScheduleObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}