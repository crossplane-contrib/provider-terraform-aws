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
	r, ok := mr.(*RedshiftCluster)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a RedshiftCluster.")
	}
	return EncodeRedshiftCluster(*r), nil
}

func EncodeRedshiftCluster(r RedshiftCluster) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRedshiftCluster_ClusterSubnetGroupName(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_MasterPassword(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_MasterUsername(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_Port(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_AllowVersionUpgrade(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_ClusterParameterGroupName(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_ClusterSecurityGroups(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_PreferredMaintenanceWindow(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_PubliclyAccessible(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_Tags(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_FinalSnapshotIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_IamRoles(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_KmsKeyId(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_SnapshotClusterIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_NodeType(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_VpcSecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_ClusterType(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_ClusterVersion(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_DatabaseName(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_Endpoint(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_EnhancedVpcRouting(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_Id(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_AvailabilityZone(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_ClusterIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_ClusterRevisionNumber(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_Encrypted(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_OwnerAccount(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_AutomatedSnapshotRetentionPeriod(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_ClusterPublicKey(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_NumberOfNodes(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_SkipFinalSnapshot(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_SnapshotIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_ElasticIp(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftCluster_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeRedshiftCluster_Logging(r.Spec.ForProvider.Logging, ctyVal)
	EncodeRedshiftCluster_SnapshotCopy(r.Spec.ForProvider.SnapshotCopy, ctyVal)
	EncodeRedshiftCluster_DnsName(r.Status.AtProvider, ctyVal)
	EncodeRedshiftCluster_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeRedshiftCluster_ClusterSubnetGroupName(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["cluster_subnet_group_name"] = cty.StringVal(p.ClusterSubnetGroupName)
}

func EncodeRedshiftCluster_MasterPassword(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["master_password"] = cty.StringVal(p.MasterPassword)
}

func EncodeRedshiftCluster_MasterUsername(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["master_username"] = cty.StringVal(p.MasterUsername)
}

func EncodeRedshiftCluster_Port(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeRedshiftCluster_AllowVersionUpgrade(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["allow_version_upgrade"] = cty.BoolVal(p.AllowVersionUpgrade)
}

func EncodeRedshiftCluster_ClusterParameterGroupName(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["cluster_parameter_group_name"] = cty.StringVal(p.ClusterParameterGroupName)
}

func EncodeRedshiftCluster_ClusterSecurityGroups(p RedshiftClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ClusterSecurityGroups {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["cluster_security_groups"] = cty.SetVal(colVals)
}

func EncodeRedshiftCluster_PreferredMaintenanceWindow(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["preferred_maintenance_window"] = cty.StringVal(p.PreferredMaintenanceWindow)
}

func EncodeRedshiftCluster_PubliclyAccessible(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["publicly_accessible"] = cty.BoolVal(p.PubliclyAccessible)
}

func EncodeRedshiftCluster_Tags(p RedshiftClusterParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeRedshiftCluster_FinalSnapshotIdentifier(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["final_snapshot_identifier"] = cty.StringVal(p.FinalSnapshotIdentifier)
}

func EncodeRedshiftCluster_IamRoles(p RedshiftClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.IamRoles {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["iam_roles"] = cty.SetVal(colVals)
}

func EncodeRedshiftCluster_KmsKeyId(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeRedshiftCluster_SnapshotClusterIdentifier(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["snapshot_cluster_identifier"] = cty.StringVal(p.SnapshotClusterIdentifier)
}

func EncodeRedshiftCluster_NodeType(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["node_type"] = cty.StringVal(p.NodeType)
}

func EncodeRedshiftCluster_VpcSecurityGroupIds(p RedshiftClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.VpcSecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["vpc_security_group_ids"] = cty.SetVal(colVals)
}

func EncodeRedshiftCluster_ClusterType(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["cluster_type"] = cty.StringVal(p.ClusterType)
}

func EncodeRedshiftCluster_ClusterVersion(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["cluster_version"] = cty.StringVal(p.ClusterVersion)
}

func EncodeRedshiftCluster_DatabaseName(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["database_name"] = cty.StringVal(p.DatabaseName)
}

func EncodeRedshiftCluster_Endpoint(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["endpoint"] = cty.StringVal(p.Endpoint)
}

func EncodeRedshiftCluster_EnhancedVpcRouting(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["enhanced_vpc_routing"] = cty.BoolVal(p.EnhancedVpcRouting)
}

func EncodeRedshiftCluster_Id(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRedshiftCluster_AvailabilityZone(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["availability_zone"] = cty.StringVal(p.AvailabilityZone)
}

func EncodeRedshiftCluster_ClusterIdentifier(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["cluster_identifier"] = cty.StringVal(p.ClusterIdentifier)
}

func EncodeRedshiftCluster_ClusterRevisionNumber(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["cluster_revision_number"] = cty.StringVal(p.ClusterRevisionNumber)
}

func EncodeRedshiftCluster_Encrypted(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeRedshiftCluster_OwnerAccount(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["owner_account"] = cty.StringVal(p.OwnerAccount)
}

func EncodeRedshiftCluster_AutomatedSnapshotRetentionPeriod(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["automated_snapshot_retention_period"] = cty.NumberIntVal(p.AutomatedSnapshotRetentionPeriod)
}

func EncodeRedshiftCluster_ClusterPublicKey(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["cluster_public_key"] = cty.StringVal(p.ClusterPublicKey)
}

func EncodeRedshiftCluster_NumberOfNodes(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["number_of_nodes"] = cty.NumberIntVal(p.NumberOfNodes)
}

func EncodeRedshiftCluster_SkipFinalSnapshot(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["skip_final_snapshot"] = cty.BoolVal(p.SkipFinalSnapshot)
}

func EncodeRedshiftCluster_SnapshotIdentifier(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["snapshot_identifier"] = cty.StringVal(p.SnapshotIdentifier)
}

func EncodeRedshiftCluster_ElasticIp(p RedshiftClusterParameters, vals map[string]cty.Value) {
	vals["elastic_ip"] = cty.StringVal(p.ElasticIp)
}

func EncodeRedshiftCluster_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeRedshiftCluster_Timeouts_Create(p, ctyVal)
	EncodeRedshiftCluster_Timeouts_Delete(p, ctyVal)
	EncodeRedshiftCluster_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeRedshiftCluster_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeRedshiftCluster_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeRedshiftCluster_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeRedshiftCluster_Logging(p Logging, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeRedshiftCluster_Logging_S3KeyPrefix(p, ctyVal)
	EncodeRedshiftCluster_Logging_BucketName(p, ctyVal)
	EncodeRedshiftCluster_Logging_Enable(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["logging"] = cty.ListVal(valsForCollection)
}

func EncodeRedshiftCluster_Logging_S3KeyPrefix(p Logging, vals map[string]cty.Value) {
	vals["s3_key_prefix"] = cty.StringVal(p.S3KeyPrefix)
}

func EncodeRedshiftCluster_Logging_BucketName(p Logging, vals map[string]cty.Value) {
	vals["bucket_name"] = cty.StringVal(p.BucketName)
}

func EncodeRedshiftCluster_Logging_Enable(p Logging, vals map[string]cty.Value) {
	vals["enable"] = cty.BoolVal(p.Enable)
}

func EncodeRedshiftCluster_SnapshotCopy(p SnapshotCopy, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeRedshiftCluster_SnapshotCopy_RetentionPeriod(p, ctyVal)
	EncodeRedshiftCluster_SnapshotCopy_DestinationRegion(p, ctyVal)
	EncodeRedshiftCluster_SnapshotCopy_GrantName(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["snapshot_copy"] = cty.ListVal(valsForCollection)
}

func EncodeRedshiftCluster_SnapshotCopy_RetentionPeriod(p SnapshotCopy, vals map[string]cty.Value) {
	vals["retention_period"] = cty.NumberIntVal(p.RetentionPeriod)
}

func EncodeRedshiftCluster_SnapshotCopy_DestinationRegion(p SnapshotCopy, vals map[string]cty.Value) {
	vals["destination_region"] = cty.StringVal(p.DestinationRegion)
}

func EncodeRedshiftCluster_SnapshotCopy_GrantName(p SnapshotCopy, vals map[string]cty.Value) {
	vals["grant_name"] = cty.StringVal(p.GrantName)
}

func EncodeRedshiftCluster_DnsName(p RedshiftClusterObservation, vals map[string]cty.Value) {
	vals["dns_name"] = cty.StringVal(p.DnsName)
}

func EncodeRedshiftCluster_Arn(p RedshiftClusterObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}