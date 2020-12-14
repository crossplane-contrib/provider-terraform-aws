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

	"github.com/zclconf/go-cty/cty"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*DmsReplicationInstance)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DmsReplicationInstance.")
	}
	return EncodeDmsReplicationInstance(*r), nil
}

func EncodeDmsReplicationInstance(r DmsReplicationInstance) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDmsReplicationInstance_EngineVersion(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationInstance_KmsKeyArn(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationInstance_PubliclyAccessible(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationInstance_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationInstance_VpcSecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationInstance_AvailabilityZone(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationInstance_MultiAz(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationInstance_AllocatedStorage(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationInstance_AllowMajorVersionUpgrade(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationInstance_ApplyImmediately(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationInstance_AutoMinorVersionUpgrade(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationInstance_ReplicationInstanceId(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationInstance_ReplicationSubnetGroupId(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationInstance_Id(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationInstance_PreferredMaintenanceWindow(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationInstance_ReplicationInstanceClass(r.Spec.ForProvider, ctyVal)
	EncodeDmsReplicationInstance_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDmsReplicationInstance_ReplicationInstancePublicIps(r.Status.AtProvider, ctyVal)
	EncodeDmsReplicationInstance_ReplicationInstancePrivateIps(r.Status.AtProvider, ctyVal)
	EncodeDmsReplicationInstance_ReplicationInstanceArn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeDmsReplicationInstance_EngineVersion(p DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	vals["engine_version"] = cty.StringVal(p.EngineVersion)
}

func EncodeDmsReplicationInstance_KmsKeyArn(p DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	vals["kms_key_arn"] = cty.StringVal(p.KmsKeyArn)
}

func EncodeDmsReplicationInstance_PubliclyAccessible(p DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	vals["publicly_accessible"] = cty.BoolVal(p.PubliclyAccessible)
}

func EncodeDmsReplicationInstance_Tags(p DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDmsReplicationInstance_VpcSecurityGroupIds(p DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.VpcSecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["vpc_security_group_ids"] = cty.SetVal(colVals)
}

func EncodeDmsReplicationInstance_AvailabilityZone(p DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	vals["availability_zone"] = cty.StringVal(p.AvailabilityZone)
}

func EncodeDmsReplicationInstance_MultiAz(p DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	vals["multi_az"] = cty.BoolVal(p.MultiAz)
}

func EncodeDmsReplicationInstance_AllocatedStorage(p DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	vals["allocated_storage"] = cty.NumberIntVal(p.AllocatedStorage)
}

func EncodeDmsReplicationInstance_AllowMajorVersionUpgrade(p DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	vals["allow_major_version_upgrade"] = cty.BoolVal(p.AllowMajorVersionUpgrade)
}

func EncodeDmsReplicationInstance_ApplyImmediately(p DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	vals["apply_immediately"] = cty.BoolVal(p.ApplyImmediately)
}

func EncodeDmsReplicationInstance_AutoMinorVersionUpgrade(p DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	vals["auto_minor_version_upgrade"] = cty.BoolVal(p.AutoMinorVersionUpgrade)
}

func EncodeDmsReplicationInstance_ReplicationInstanceId(p DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	vals["replication_instance_id"] = cty.StringVal(p.ReplicationInstanceId)
}

func EncodeDmsReplicationInstance_ReplicationSubnetGroupId(p DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	vals["replication_subnet_group_id"] = cty.StringVal(p.ReplicationSubnetGroupId)
}

func EncodeDmsReplicationInstance_Id(p DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDmsReplicationInstance_PreferredMaintenanceWindow(p DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	vals["preferred_maintenance_window"] = cty.StringVal(p.PreferredMaintenanceWindow)
}

func EncodeDmsReplicationInstance_ReplicationInstanceClass(p DmsReplicationInstanceParameters, vals map[string]cty.Value) {
	vals["replication_instance_class"] = cty.StringVal(p.ReplicationInstanceClass)
}

func EncodeDmsReplicationInstance_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeDmsReplicationInstance_Timeouts_Create(p, ctyVal)
	EncodeDmsReplicationInstance_Timeouts_Delete(p, ctyVal)
	EncodeDmsReplicationInstance_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDmsReplicationInstance_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDmsReplicationInstance_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeDmsReplicationInstance_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeDmsReplicationInstance_ReplicationInstancePublicIps(p DmsReplicationInstanceObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ReplicationInstancePublicIps {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["replication_instance_public_ips"] = cty.ListVal(colVals)
}

func EncodeDmsReplicationInstance_ReplicationInstancePrivateIps(p DmsReplicationInstanceObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ReplicationInstancePrivateIps {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["replication_instance_private_ips"] = cty.ListVal(colVals)
}

func EncodeDmsReplicationInstance_ReplicationInstanceArn(p DmsReplicationInstanceObservation, vals map[string]cty.Value) {
	vals["replication_instance_arn"] = cty.StringVal(p.ReplicationInstanceArn)
}