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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*ElasticacheReplicationGroup)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ElasticacheReplicationGroup.")
	}
	return EncodeElasticacheReplicationGroup(*r), nil
}

func EncodeElasticacheReplicationGroup(r ElasticacheReplicationGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeElasticacheReplicationGroup_AtRestEncryptionEnabled(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_MaintenanceWindow(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_SnapshotRetentionLimit(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_TransitEncryptionEnabled(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_SecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_SecurityGroupNames(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_SubnetGroupName(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_AutoMinorVersionUpgrade(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_Engine(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_Port(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_ParameterGroupName(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_ApplyImmediately(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_AvailabilityZones(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_KmsKeyId(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_ReplicationGroupId(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_NumberCacheClusters(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_SnapshotWindow(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_AutomaticFailoverEnabled(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_EngineVersion(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_ReplicationGroupDescription(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_SnapshotArns(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_SnapshotName(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_AuthToken(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_NodeType(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_NotificationTopicArn(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheReplicationGroup_ClusterMode(r.Spec.ForProvider.ClusterMode, ctyVal)
	EncodeElasticacheReplicationGroup_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeElasticacheReplicationGroup_MemberClusters(r.Status.AtProvider, ctyVal)
	EncodeElasticacheReplicationGroup_PrimaryEndpointAddress(r.Status.AtProvider, ctyVal)
	EncodeElasticacheReplicationGroup_ConfigurationEndpointAddress(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeElasticacheReplicationGroup_AtRestEncryptionEnabled(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	vals["at_rest_encryption_enabled"] = cty.BoolVal(p.AtRestEncryptionEnabled)
}

func EncodeElasticacheReplicationGroup_MaintenanceWindow(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	vals["maintenance_window"] = cty.StringVal(p.MaintenanceWindow)
}

func EncodeElasticacheReplicationGroup_SnapshotRetentionLimit(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	vals["snapshot_retention_limit"] = cty.NumberIntVal(p.SnapshotRetentionLimit)
}

func EncodeElasticacheReplicationGroup_TransitEncryptionEnabled(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	vals["transit_encryption_enabled"] = cty.BoolVal(p.TransitEncryptionEnabled)
}

func EncodeElasticacheReplicationGroup_SecurityGroupIds(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_group_ids"] = cty.SetVal(colVals)
}

func EncodeElasticacheReplicationGroup_SecurityGroupNames(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroupNames {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_group_names"] = cty.SetVal(colVals)
}

func EncodeElasticacheReplicationGroup_SubnetGroupName(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	vals["subnet_group_name"] = cty.StringVal(p.SubnetGroupName)
}

func EncodeElasticacheReplicationGroup_AutoMinorVersionUpgrade(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	vals["auto_minor_version_upgrade"] = cty.BoolVal(p.AutoMinorVersionUpgrade)
}

func EncodeElasticacheReplicationGroup_Engine(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	vals["engine"] = cty.StringVal(p.Engine)
}

func EncodeElasticacheReplicationGroup_Port(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeElasticacheReplicationGroup_ParameterGroupName(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	vals["parameter_group_name"] = cty.StringVal(p.ParameterGroupName)
}

func EncodeElasticacheReplicationGroup_ApplyImmediately(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	vals["apply_immediately"] = cty.BoolVal(p.ApplyImmediately)
}

func EncodeElasticacheReplicationGroup_AvailabilityZones(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AvailabilityZones {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["availability_zones"] = cty.SetVal(colVals)
}

func EncodeElasticacheReplicationGroup_KmsKeyId(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeElasticacheReplicationGroup_ReplicationGroupId(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	vals["replication_group_id"] = cty.StringVal(p.ReplicationGroupId)
}

func EncodeElasticacheReplicationGroup_Tags(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeElasticacheReplicationGroup_NumberCacheClusters(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	vals["number_cache_clusters"] = cty.NumberIntVal(p.NumberCacheClusters)
}

func EncodeElasticacheReplicationGroup_SnapshotWindow(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	vals["snapshot_window"] = cty.StringVal(p.SnapshotWindow)
}

func EncodeElasticacheReplicationGroup_AutomaticFailoverEnabled(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	vals["automatic_failover_enabled"] = cty.BoolVal(p.AutomaticFailoverEnabled)
}

func EncodeElasticacheReplicationGroup_Id(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeElasticacheReplicationGroup_EngineVersion(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	vals["engine_version"] = cty.StringVal(p.EngineVersion)
}

func EncodeElasticacheReplicationGroup_ReplicationGroupDescription(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	vals["replication_group_description"] = cty.StringVal(p.ReplicationGroupDescription)
}

func EncodeElasticacheReplicationGroup_SnapshotArns(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SnapshotArns {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["snapshot_arns"] = cty.SetVal(colVals)
}

func EncodeElasticacheReplicationGroup_SnapshotName(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	vals["snapshot_name"] = cty.StringVal(p.SnapshotName)
}

func EncodeElasticacheReplicationGroup_AuthToken(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	vals["auth_token"] = cty.StringVal(p.AuthToken)
}

func EncodeElasticacheReplicationGroup_NodeType(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	vals["node_type"] = cty.StringVal(p.NodeType)
}

func EncodeElasticacheReplicationGroup_NotificationTopicArn(p ElasticacheReplicationGroupParameters, vals map[string]cty.Value) {
	vals["notification_topic_arn"] = cty.StringVal(p.NotificationTopicArn)
}

func EncodeElasticacheReplicationGroup_ClusterMode(p ClusterMode, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeElasticacheReplicationGroup_ClusterMode_ReplicasPerNodeGroup(p, ctyVal)
	EncodeElasticacheReplicationGroup_ClusterMode_NumNodeGroups(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["cluster_mode"] = cty.ListVal(valsForCollection)
}

func EncodeElasticacheReplicationGroup_ClusterMode_ReplicasPerNodeGroup(p ClusterMode, vals map[string]cty.Value) {
	vals["replicas_per_node_group"] = cty.NumberIntVal(p.ReplicasPerNodeGroup)
}

func EncodeElasticacheReplicationGroup_ClusterMode_NumNodeGroups(p ClusterMode, vals map[string]cty.Value) {
	vals["num_node_groups"] = cty.NumberIntVal(p.NumNodeGroups)
}

func EncodeElasticacheReplicationGroup_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeElasticacheReplicationGroup_Timeouts_Create(p, ctyVal)
	EncodeElasticacheReplicationGroup_Timeouts_Delete(p, ctyVal)
	EncodeElasticacheReplicationGroup_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeElasticacheReplicationGroup_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeElasticacheReplicationGroup_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeElasticacheReplicationGroup_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeElasticacheReplicationGroup_MemberClusters(p ElasticacheReplicationGroupObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.MemberClusters {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["member_clusters"] = cty.SetVal(colVals)
}

func EncodeElasticacheReplicationGroup_PrimaryEndpointAddress(p ElasticacheReplicationGroupObservation, vals map[string]cty.Value) {
	vals["primary_endpoint_address"] = cty.StringVal(p.PrimaryEndpointAddress)
}

func EncodeElasticacheReplicationGroup_ConfigurationEndpointAddress(p ElasticacheReplicationGroupObservation, vals map[string]cty.Value) {
	vals["configuration_endpoint_address"] = cty.StringVal(p.ConfigurationEndpointAddress)
}