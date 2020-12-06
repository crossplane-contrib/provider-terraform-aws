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
	"github.com/zclconf/go-cty/cty"
)

func EncodeDaxCluster(r DaxCluster) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDaxCluster_Description(r.Spec.ForProvider, ctyVal)
	EncodeDaxCluster_ReplicationFactor(r.Spec.ForProvider, ctyVal)
	EncodeDaxCluster_ClusterName(r.Spec.ForProvider, ctyVal)
	EncodeDaxCluster_SubnetGroupName(r.Spec.ForProvider, ctyVal)
	EncodeDaxCluster_MaintenanceWindow(r.Spec.ForProvider, ctyVal)
	EncodeDaxCluster_NodeType(r.Spec.ForProvider, ctyVal)
	EncodeDaxCluster_SecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeDaxCluster_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDaxCluster_IamRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeDaxCluster_AvailabilityZones(r.Spec.ForProvider, ctyVal)
	EncodeDaxCluster_Id(r.Spec.ForProvider, ctyVal)
	EncodeDaxCluster_NotificationTopicArn(r.Spec.ForProvider, ctyVal)
	EncodeDaxCluster_ParameterGroupName(r.Spec.ForProvider, ctyVal)
	EncodeDaxCluster_ServerSideEncryption(r.Spec.ForProvider.ServerSideEncryption, ctyVal)
	EncodeDaxCluster_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDaxCluster_Nodes(r.Status.AtProvider.Nodes, ctyVal)
	EncodeDaxCluster_ClusterAddress(r.Status.AtProvider, ctyVal)
	EncodeDaxCluster_Port(r.Status.AtProvider, ctyVal)
	EncodeDaxCluster_ConfigurationEndpoint(r.Status.AtProvider, ctyVal)
	EncodeDaxCluster_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDaxCluster_Description(p DaxClusterParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeDaxCluster_ReplicationFactor(p DaxClusterParameters, vals map[string]cty.Value) {
	vals["replication_factor"] = cty.NumberIntVal(p.ReplicationFactor)
}

func EncodeDaxCluster_ClusterName(p DaxClusterParameters, vals map[string]cty.Value) {
	vals["cluster_name"] = cty.StringVal(p.ClusterName)
}

func EncodeDaxCluster_SubnetGroupName(p DaxClusterParameters, vals map[string]cty.Value) {
	vals["subnet_group_name"] = cty.StringVal(p.SubnetGroupName)
}

func EncodeDaxCluster_MaintenanceWindow(p DaxClusterParameters, vals map[string]cty.Value) {
	vals["maintenance_window"] = cty.StringVal(p.MaintenanceWindow)
}

func EncodeDaxCluster_NodeType(p DaxClusterParameters, vals map[string]cty.Value) {
	vals["node_type"] = cty.StringVal(p.NodeType)
}

func EncodeDaxCluster_SecurityGroupIds(p DaxClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_group_ids"] = cty.SetVal(colVals)
}

func EncodeDaxCluster_Tags(p DaxClusterParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDaxCluster_IamRoleArn(p DaxClusterParameters, vals map[string]cty.Value) {
	vals["iam_role_arn"] = cty.StringVal(p.IamRoleArn)
}

func EncodeDaxCluster_AvailabilityZones(p DaxClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AvailabilityZones {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["availability_zones"] = cty.SetVal(colVals)
}

func EncodeDaxCluster_Id(p DaxClusterParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDaxCluster_NotificationTopicArn(p DaxClusterParameters, vals map[string]cty.Value) {
	vals["notification_topic_arn"] = cty.StringVal(p.NotificationTopicArn)
}

func EncodeDaxCluster_ParameterGroupName(p DaxClusterParameters, vals map[string]cty.Value) {
	vals["parameter_group_name"] = cty.StringVal(p.ParameterGroupName)
}

func EncodeDaxCluster_ServerSideEncryption(p ServerSideEncryption, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDaxCluster_ServerSideEncryption_Enabled(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["server_side_encryption"] = cty.ListVal(valsForCollection)
}

func EncodeDaxCluster_ServerSideEncryption_Enabled(p ServerSideEncryption, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeDaxCluster_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeDaxCluster_Timeouts_Create(p, ctyVal)
	EncodeDaxCluster_Timeouts_Delete(p, ctyVal)
	EncodeDaxCluster_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDaxCluster_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDaxCluster_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeDaxCluster_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeDaxCluster_Nodes(p []Nodes, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeDaxCluster_Nodes_Address(v, ctyVal)
		EncodeDaxCluster_Nodes_AvailabilityZone(v, ctyVal)
		EncodeDaxCluster_Nodes_Id(v, ctyVal)
		EncodeDaxCluster_Nodes_Port(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["nodes"] = cty.ListVal(valsForCollection)
}

func EncodeDaxCluster_Nodes_Address(p Nodes, vals map[string]cty.Value) {
	vals["address"] = cty.StringVal(p.Address)
}

func EncodeDaxCluster_Nodes_AvailabilityZone(p Nodes, vals map[string]cty.Value) {
	vals["availability_zone"] = cty.StringVal(p.AvailabilityZone)
}

func EncodeDaxCluster_Nodes_Id(p Nodes, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDaxCluster_Nodes_Port(p Nodes, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeDaxCluster_ClusterAddress(p DaxClusterObservation, vals map[string]cty.Value) {
	vals["cluster_address"] = cty.StringVal(p.ClusterAddress)
}

func EncodeDaxCluster_Port(p DaxClusterObservation, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeDaxCluster_ConfigurationEndpoint(p DaxClusterObservation, vals map[string]cty.Value) {
	vals["configuration_endpoint"] = cty.StringVal(p.ConfigurationEndpoint)
}

func EncodeDaxCluster_Arn(p DaxClusterObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}