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
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*DmsReplicationInstance)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDmsReplicationInstance(r, ctyValue)
}

func DecodeDmsReplicationInstance(prev *DmsReplicationInstance, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDmsReplicationInstance_EngineVersion(&new.Spec.ForProvider, valMap)
	DecodeDmsReplicationInstance_KmsKeyArn(&new.Spec.ForProvider, valMap)
	DecodeDmsReplicationInstance_PubliclyAccessible(&new.Spec.ForProvider, valMap)
	DecodeDmsReplicationInstance_Tags(&new.Spec.ForProvider, valMap)
	DecodeDmsReplicationInstance_VpcSecurityGroupIds(&new.Spec.ForProvider, valMap)
	DecodeDmsReplicationInstance_AvailabilityZone(&new.Spec.ForProvider, valMap)
	DecodeDmsReplicationInstance_MultiAz(&new.Spec.ForProvider, valMap)
	DecodeDmsReplicationInstance_AllocatedStorage(&new.Spec.ForProvider, valMap)
	DecodeDmsReplicationInstance_AllowMajorVersionUpgrade(&new.Spec.ForProvider, valMap)
	DecodeDmsReplicationInstance_ApplyImmediately(&new.Spec.ForProvider, valMap)
	DecodeDmsReplicationInstance_AutoMinorVersionUpgrade(&new.Spec.ForProvider, valMap)
	DecodeDmsReplicationInstance_ReplicationInstanceId(&new.Spec.ForProvider, valMap)
	DecodeDmsReplicationInstance_ReplicationSubnetGroupId(&new.Spec.ForProvider, valMap)
	DecodeDmsReplicationInstance_Id(&new.Spec.ForProvider, valMap)
	DecodeDmsReplicationInstance_PreferredMaintenanceWindow(&new.Spec.ForProvider, valMap)
	DecodeDmsReplicationInstance_ReplicationInstanceClass(&new.Spec.ForProvider, valMap)
	DecodeDmsReplicationInstance_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeDmsReplicationInstance_ReplicationInstancePublicIps(&new.Status.AtProvider, valMap)
	DecodeDmsReplicationInstance_ReplicationInstancePrivateIps(&new.Status.AtProvider, valMap)
	DecodeDmsReplicationInstance_ReplicationInstanceArn(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeDmsReplicationInstance_EngineVersion(p *DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	p.EngineVersion = ctwhy.ValueAsString(vals["engine_version"])
}

func DecodeDmsReplicationInstance_KmsKeyArn(p *DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	p.KmsKeyArn = ctwhy.ValueAsString(vals["kms_key_arn"])
}

func DecodeDmsReplicationInstance_PubliclyAccessible(p *DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	p.PubliclyAccessible = ctwhy.ValueAsBool(vals["publicly_accessible"])
}

func DecodeDmsReplicationInstance_Tags(p *DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeDmsReplicationInstance_VpcSecurityGroupIds(p *DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["vpc_security_group_ids"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.VpcSecurityGroupIds = goVals
}

func DecodeDmsReplicationInstance_AvailabilityZone(p *DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	p.AvailabilityZone = ctwhy.ValueAsString(vals["availability_zone"])
}

func DecodeDmsReplicationInstance_MultiAz(p *DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	p.MultiAz = ctwhy.ValueAsBool(vals["multi_az"])
}

func DecodeDmsReplicationInstance_AllocatedStorage(p *DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	p.AllocatedStorage = ctwhy.ValueAsInt64(vals["allocated_storage"])
}

func DecodeDmsReplicationInstance_AllowMajorVersionUpgrade(p *DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	p.AllowMajorVersionUpgrade = ctwhy.ValueAsBool(vals["allow_major_version_upgrade"])
}

func DecodeDmsReplicationInstance_ApplyImmediately(p *DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	p.ApplyImmediately = ctwhy.ValueAsBool(vals["apply_immediately"])
}

func DecodeDmsReplicationInstance_AutoMinorVersionUpgrade(p *DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	p.AutoMinorVersionUpgrade = ctwhy.ValueAsBool(vals["auto_minor_version_upgrade"])
}

func DecodeDmsReplicationInstance_ReplicationInstanceId(p *DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	p.ReplicationInstanceId = ctwhy.ValueAsString(vals["replication_instance_id"])
}

func DecodeDmsReplicationInstance_ReplicationSubnetGroupId(p *DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	p.ReplicationSubnetGroupId = ctwhy.ValueAsString(vals["replication_subnet_group_id"])
}

func DecodeDmsReplicationInstance_Id(p *DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeDmsReplicationInstance_PreferredMaintenanceWindow(p *DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	p.PreferredMaintenanceWindow = ctwhy.ValueAsString(vals["preferred_maintenance_window"])
}

func DecodeDmsReplicationInstance_ReplicationInstanceClass(p *DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	p.ReplicationInstanceClass = ctwhy.ValueAsString(vals["replication_instance_class"])
}

func DecodeDmsReplicationInstance_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeDmsReplicationInstance_Timeouts_Create(p, valMap)
	DecodeDmsReplicationInstance_Timeouts_Delete(p, valMap)
	DecodeDmsReplicationInstance_Timeouts_Update(p, valMap)
}

func DecodeDmsReplicationInstance_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

func DecodeDmsReplicationInstance_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

func DecodeDmsReplicationInstance_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}

func DecodeDmsReplicationInstance_ReplicationInstancePublicIps(p *DmsReplicationInstanceObservation, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["replication_instance_public_ips"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.ReplicationInstancePublicIps = goVals
}

func DecodeDmsReplicationInstance_ReplicationInstancePrivateIps(p *DmsReplicationInstanceObservation, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["replication_instance_private_ips"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.ReplicationInstancePrivateIps = goVals
}

func DecodeDmsReplicationInstance_ReplicationInstanceArn(p *DmsReplicationInstanceObservation, vals map[string]cty.Value) {
	p.ReplicationInstanceArn = ctwhy.ValueAsString(vals["replication_instance_arn"])
}