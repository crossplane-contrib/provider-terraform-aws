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
	k := kube.(*DmsReplicationSubnetGroup)
	p := prov.(*DmsReplicationSubnetGroup)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDmsReplicationSubnetGroup_ReplicationSubnetGroupId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationSubnetGroup_SubnetIds(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationSubnetGroup_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationSubnetGroup_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationSubnetGroup_ReplicationSubnetGroupDescription(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationSubnetGroup_VpcId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsReplicationSubnetGroup_ReplicationSubnetGroupArn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeDmsReplicationSubnetGroup_ReplicationSubnetGroupId(k *DmsReplicationSubnetGroupParameters, p *DmsReplicationSubnetGroupParameters, md *plugin.MergeDescription) bool {
	if k.ReplicationSubnetGroupId != p.ReplicationSubnetGroupId {
		p.ReplicationSubnetGroupId = k.ReplicationSubnetGroupId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDmsReplicationSubnetGroup_SubnetIds(k *DmsReplicationSubnetGroupParameters, p *DmsReplicationSubnetGroupParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(p.SubnetIds, p.SubnetIds) {
		p.SubnetIds = k.SubnetIds
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDmsReplicationSubnetGroup_Tags(k *DmsReplicationSubnetGroupParameters, p *DmsReplicationSubnetGroupParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(p.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDmsReplicationSubnetGroup_Id(k *DmsReplicationSubnetGroupParameters, p *DmsReplicationSubnetGroupParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDmsReplicationSubnetGroup_ReplicationSubnetGroupDescription(k *DmsReplicationSubnetGroupParameters, p *DmsReplicationSubnetGroupParameters, md *plugin.MergeDescription) bool {
	if k.ReplicationSubnetGroupDescription != p.ReplicationSubnetGroupDescription {
		p.ReplicationSubnetGroupDescription = k.ReplicationSubnetGroupDescription
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDmsReplicationSubnetGroup_VpcId(k *DmsReplicationSubnetGroupObservation, p *DmsReplicationSubnetGroupObservation, md *plugin.MergeDescription) bool {
	if k.VpcId != p.VpcId {
		k.VpcId = p.VpcId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDmsReplicationSubnetGroup_ReplicationSubnetGroupArn(k *DmsReplicationSubnetGroupObservation, p *DmsReplicationSubnetGroupObservation, md *plugin.MergeDescription) bool {
	if k.ReplicationSubnetGroupArn != p.ReplicationSubnetGroupArn {
		k.ReplicationSubnetGroupArn = p.ReplicationSubnetGroupArn
		md.StatusUpdated = true
		return true
	}
	return false
}