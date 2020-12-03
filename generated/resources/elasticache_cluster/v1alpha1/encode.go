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

package v1alpha1func EncodeElasticacheCluster(r ElasticacheCluster) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeElasticacheCluster_AvailabilityZone(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheCluster_Id(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheCluster_SnapshotWindow(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheCluster_Engine(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheCluster_NodeType(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheCluster_SubnetGroupName(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheCluster_SnapshotArns(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheCluster_EngineVersion(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheCluster_MaintenanceWindow(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheCluster_NotificationTopicArn(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheCluster_NumCacheNodes(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheCluster_SecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheCluster_SnapshotName(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheCluster_Tags(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheCluster_ClusterId(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheCluster_SecurityGroupNames(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheCluster_SnapshotRetentionLimit(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheCluster_ReplicationGroupId(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheCluster_ParameterGroupName(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheCluster_PreferredAvailabilityZones(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheCluster_ApplyImmediately(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheCluster_AzMode(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheCluster_Port(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheCluster_ClusterAddress(r.Status.AtProvider, ctyVal)
	EncodeElasticacheCluster_Arn(r.Status.AtProvider, ctyVal)
	EncodeElasticacheCluster_CacheNodes(r.Status.AtProvider.CacheNodes, ctyVal)
	EncodeElasticacheCluster_ConfigurationEndpoint(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeElasticacheCluster_AvailabilityZone(p *ElasticacheClusterParameters, vals map[string]cty.Value) {
	vals["availability_zone"] = cty.StringVal(p.AvailabilityZone)
}

func EncodeElasticacheCluster_Id(p *ElasticacheClusterParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeElasticacheCluster_SnapshotWindow(p *ElasticacheClusterParameters, vals map[string]cty.Value) {
	vals["snapshot_window"] = cty.StringVal(p.SnapshotWindow)
}

func EncodeElasticacheCluster_Engine(p *ElasticacheClusterParameters, vals map[string]cty.Value) {
	vals["engine"] = cty.StringVal(p.Engine)
}

func EncodeElasticacheCluster_NodeType(p *ElasticacheClusterParameters, vals map[string]cty.Value) {
	vals["node_type"] = cty.StringVal(p.NodeType)
}

func EncodeElasticacheCluster_SubnetGroupName(p *ElasticacheClusterParameters, vals map[string]cty.Value) {
	vals["subnet_group_name"] = cty.StringVal(p.SubnetGroupName)
}

func EncodeElasticacheCluster_SnapshotArns(p *ElasticacheClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SnapshotArns {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["snapshot_arns"] = cty.SetVal(colVals)
}

func EncodeElasticacheCluster_EngineVersion(p *ElasticacheClusterParameters, vals map[string]cty.Value) {
	vals["engine_version"] = cty.StringVal(p.EngineVersion)
}

func EncodeElasticacheCluster_MaintenanceWindow(p *ElasticacheClusterParameters, vals map[string]cty.Value) {
	vals["maintenance_window"] = cty.StringVal(p.MaintenanceWindow)
}

func EncodeElasticacheCluster_NotificationTopicArn(p *ElasticacheClusterParameters, vals map[string]cty.Value) {
	vals["notification_topic_arn"] = cty.StringVal(p.NotificationTopicArn)
}

func EncodeElasticacheCluster_NumCacheNodes(p *ElasticacheClusterParameters, vals map[string]cty.Value) {
	vals["num_cache_nodes"] = cty.IntVal(p.NumCacheNodes)
}

func EncodeElasticacheCluster_SecurityGroupIds(p *ElasticacheClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_group_ids"] = cty.SetVal(colVals)
}

func EncodeElasticacheCluster_SnapshotName(p *ElasticacheClusterParameters, vals map[string]cty.Value) {
	vals["snapshot_name"] = cty.StringVal(p.SnapshotName)
}

func EncodeElasticacheCluster_Tags(p *ElasticacheClusterParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeElasticacheCluster_ClusterId(p *ElasticacheClusterParameters, vals map[string]cty.Value) {
	vals["cluster_id"] = cty.StringVal(p.ClusterId)
}

func EncodeElasticacheCluster_SecurityGroupNames(p *ElasticacheClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroupNames {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_group_names"] = cty.SetVal(colVals)
}

func EncodeElasticacheCluster_SnapshotRetentionLimit(p *ElasticacheClusterParameters, vals map[string]cty.Value) {
	vals["snapshot_retention_limit"] = cty.IntVal(p.SnapshotRetentionLimit)
}

func EncodeElasticacheCluster_ReplicationGroupId(p *ElasticacheClusterParameters, vals map[string]cty.Value) {
	vals["replication_group_id"] = cty.StringVal(p.ReplicationGroupId)
}

func EncodeElasticacheCluster_ParameterGroupName(p *ElasticacheClusterParameters, vals map[string]cty.Value) {
	vals["parameter_group_name"] = cty.StringVal(p.ParameterGroupName)
}

func EncodeElasticacheCluster_PreferredAvailabilityZones(p *ElasticacheClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.PreferredAvailabilityZones {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["preferred_availability_zones"] = cty.ListVal(colVals)
}

func EncodeElasticacheCluster_ApplyImmediately(p *ElasticacheClusterParameters, vals map[string]cty.Value) {
	vals["apply_immediately"] = cty.BoolVal(p.ApplyImmediately)
}

func EncodeElasticacheCluster_AzMode(p *ElasticacheClusterParameters, vals map[string]cty.Value) {
	vals["az_mode"] = cty.StringVal(p.AzMode)
}

func EncodeElasticacheCluster_Port(p *ElasticacheClusterParameters, vals map[string]cty.Value) {
	vals["port"] = cty.IntVal(p.Port)
}

func EncodeElasticacheCluster_ClusterAddress(p *ElasticacheClusterObservation, vals map[string]cty.Value) {
	vals["cluster_address"] = cty.StringVal(p.ClusterAddress)
}

func EncodeElasticacheCluster_Arn(p *ElasticacheClusterObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeElasticacheCluster_CacheNodes(p *CacheNodes, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.CacheNodes {
		ctyVal = make(map[string]cty.Value)
		EncodeElasticacheCluster_CacheNodes_Address(v, ctyVal)
		EncodeElasticacheCluster_CacheNodes_AvailabilityZone(v, ctyVal)
		EncodeElasticacheCluster_CacheNodes_Id(v, ctyVal)
		EncodeElasticacheCluster_CacheNodes_Port(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["cache_nodes"] = cty.ListVal(valsForCollection)
}

func EncodeElasticacheCluster_CacheNodes_Address(p *CacheNodes, vals map[string]cty.Value) {
	vals["address"] = cty.StringVal(p.Address)
}

func EncodeElasticacheCluster_CacheNodes_AvailabilityZone(p *CacheNodes, vals map[string]cty.Value) {
	vals["availability_zone"] = cty.StringVal(p.AvailabilityZone)
}

func EncodeElasticacheCluster_CacheNodes_Id(p *CacheNodes, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeElasticacheCluster_CacheNodes_Port(p *CacheNodes, vals map[string]cty.Value) {
	vals["port"] = cty.IntVal(p.Port)
}

func EncodeElasticacheCluster_ConfigurationEndpoint(p *ElasticacheClusterObservation, vals map[string]cty.Value) {
	vals["configuration_endpoint"] = cty.StringVal(p.ConfigurationEndpoint)
}