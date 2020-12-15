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
	k := kube.(*DocdbClusterSnapshot)
	p := prov.(*DocdbClusterSnapshot)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDocdbClusterSnapshot_DbClusterSnapshotIdentifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterSnapshot_DbClusterIdentifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterSnapshot_Timeouts(&k.Spec.ForProvider.Timeouts, &p.Spec.ForProvider.Timeouts, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterSnapshot_EngineVersion(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterSnapshot_DbClusterSnapshotArn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterSnapshot_Status(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterSnapshot_Port(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterSnapshot_SourceDbClusterSnapshotArn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterSnapshot_StorageEncrypted(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterSnapshot_VpcId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterSnapshot_KmsKeyId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterSnapshot_SnapshotType(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterSnapshot_AvailabilityZones(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDocdbClusterSnapshot_Engine(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeDocdbClusterSnapshot_DbClusterSnapshotIdentifier(k *DocdbClusterSnapshotParameters, p *DocdbClusterSnapshotParameters, md *plugin.MergeDescription) bool {
	if k.DbClusterSnapshotIdentifier != p.DbClusterSnapshotIdentifier {
		p.DbClusterSnapshotIdentifier = k.DbClusterSnapshotIdentifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDocdbClusterSnapshot_DbClusterIdentifier(k *DocdbClusterSnapshotParameters, p *DocdbClusterSnapshotParameters, md *plugin.MergeDescription) bool {
	if k.DbClusterIdentifier != p.DbClusterIdentifier {
		p.DbClusterIdentifier = k.DbClusterIdentifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeDocdbClusterSnapshot_Timeouts(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeDocdbClusterSnapshot_Timeouts_Create(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeDocdbClusterSnapshot_Timeouts_Create(k *Timeouts, p *Timeouts, md *plugin.MergeDescription) bool {
	if k.Create != p.Create {
		p.Create = k.Create
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbClusterSnapshot_EngineVersion(k *DocdbClusterSnapshotObservation, p *DocdbClusterSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.EngineVersion != p.EngineVersion {
		k.EngineVersion = p.EngineVersion
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbClusterSnapshot_DbClusterSnapshotArn(k *DocdbClusterSnapshotObservation, p *DocdbClusterSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.DbClusterSnapshotArn != p.DbClusterSnapshotArn {
		k.DbClusterSnapshotArn = p.DbClusterSnapshotArn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbClusterSnapshot_Status(k *DocdbClusterSnapshotObservation, p *DocdbClusterSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.Status != p.Status {
		k.Status = p.Status
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbClusterSnapshot_Port(k *DocdbClusterSnapshotObservation, p *DocdbClusterSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.Port != p.Port {
		k.Port = p.Port
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbClusterSnapshot_SourceDbClusterSnapshotArn(k *DocdbClusterSnapshotObservation, p *DocdbClusterSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.SourceDbClusterSnapshotArn != p.SourceDbClusterSnapshotArn {
		k.SourceDbClusterSnapshotArn = p.SourceDbClusterSnapshotArn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbClusterSnapshot_StorageEncrypted(k *DocdbClusterSnapshotObservation, p *DocdbClusterSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.StorageEncrypted != p.StorageEncrypted {
		k.StorageEncrypted = p.StorageEncrypted
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbClusterSnapshot_VpcId(k *DocdbClusterSnapshotObservation, p *DocdbClusterSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.VpcId != p.VpcId {
		k.VpcId = p.VpcId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbClusterSnapshot_KmsKeyId(k *DocdbClusterSnapshotObservation, p *DocdbClusterSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.KmsKeyId != p.KmsKeyId {
		k.KmsKeyId = p.KmsKeyId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbClusterSnapshot_SnapshotType(k *DocdbClusterSnapshotObservation, p *DocdbClusterSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.SnapshotType != p.SnapshotType {
		k.SnapshotType = p.SnapshotType
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateStatus
func MergeDocdbClusterSnapshot_AvailabilityZones(k *DocdbClusterSnapshotObservation, p *DocdbClusterSnapshotObservation, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.AvailabilityZones, p.AvailabilityZones) {
		k.AvailabilityZones = p.AvailabilityZones
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDocdbClusterSnapshot_Engine(k *DocdbClusterSnapshotObservation, p *DocdbClusterSnapshotObservation, md *plugin.MergeDescription) bool {
	if k.Engine != p.Engine {
		k.Engine = p.Engine
		md.StatusUpdated = true
		return true
	}
	return false
}