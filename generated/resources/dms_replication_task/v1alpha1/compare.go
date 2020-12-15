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
	k := kube.(*DmsReplicationTask)
	p := prov.(*DmsReplicationTask)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDmsReplicationTask_ReplicationTaskSettings(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationTask_SourceEndpointArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationTask_TableMappings(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationTask_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationTask_CdcStartTime(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationTask_MigrationType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationTask_ReplicationInstanceArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationTask_ReplicationTaskId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationTask_TargetEndpointArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationTask_ReplicationTaskArn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeDmsReplicationTask_ReplicationTaskSettings(k *DmsReplicationTaskParameters, p *DmsReplicationTaskParameters, md *plugin.MergeDescription) bool {
	if k.ReplicationTaskSettings != p.ReplicationTaskSettings {
		p.ReplicationTaskSettings = k.ReplicationTaskSettings
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDmsReplicationTask_SourceEndpointArn(k *DmsReplicationTaskParameters, p *DmsReplicationTaskParameters, md *plugin.MergeDescription) bool {
	if k.SourceEndpointArn != p.SourceEndpointArn {
		p.SourceEndpointArn = k.SourceEndpointArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDmsReplicationTask_TableMappings(k *DmsReplicationTaskParameters, p *DmsReplicationTaskParameters, md *plugin.MergeDescription) bool {
	if k.TableMappings != p.TableMappings {
		p.TableMappings = k.TableMappings
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDmsReplicationTask_Tags(k *DmsReplicationTaskParameters, p *DmsReplicationTaskParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDmsReplicationTask_CdcStartTime(k *DmsReplicationTaskParameters, p *DmsReplicationTaskParameters, md *plugin.MergeDescription) bool {
	if k.CdcStartTime != p.CdcStartTime {
		p.CdcStartTime = k.CdcStartTime
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDmsReplicationTask_MigrationType(k *DmsReplicationTaskParameters, p *DmsReplicationTaskParameters, md *plugin.MergeDescription) bool {
	if k.MigrationType != p.MigrationType {
		p.MigrationType = k.MigrationType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDmsReplicationTask_ReplicationInstanceArn(k *DmsReplicationTaskParameters, p *DmsReplicationTaskParameters, md *plugin.MergeDescription) bool {
	if k.ReplicationInstanceArn != p.ReplicationInstanceArn {
		p.ReplicationInstanceArn = k.ReplicationInstanceArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDmsReplicationTask_ReplicationTaskId(k *DmsReplicationTaskParameters, p *DmsReplicationTaskParameters, md *plugin.MergeDescription) bool {
	if k.ReplicationTaskId != p.ReplicationTaskId {
		p.ReplicationTaskId = k.ReplicationTaskId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDmsReplicationTask_TargetEndpointArn(k *DmsReplicationTaskParameters, p *DmsReplicationTaskParameters, md *plugin.MergeDescription) bool {
	if k.TargetEndpointArn != p.TargetEndpointArn {
		p.TargetEndpointArn = k.TargetEndpointArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDmsReplicationTask_ReplicationTaskArn(k *DmsReplicationTaskObservation, p *DmsReplicationTaskObservation, md *plugin.MergeDescription) bool {
	if k.ReplicationTaskArn != p.ReplicationTaskArn {
		k.ReplicationTaskArn = p.ReplicationTaskArn
		md.StatusUpdated = true
		return true
	}
	return false
}