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

func EncodeRdsGlobalCluster(r RdsGlobalCluster) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRdsGlobalCluster_SourceDbClusterIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeRdsGlobalCluster_StorageEncrypted(r.Spec.ForProvider, ctyVal)
	EncodeRdsGlobalCluster_ForceDestroy(r.Spec.ForProvider, ctyVal)
	EncodeRdsGlobalCluster_GlobalClusterIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeRdsGlobalCluster_Id(r.Spec.ForProvider, ctyVal)
	EncodeRdsGlobalCluster_DatabaseName(r.Spec.ForProvider, ctyVal)
	EncodeRdsGlobalCluster_DeletionProtection(r.Spec.ForProvider, ctyVal)
	EncodeRdsGlobalCluster_Engine(r.Spec.ForProvider, ctyVal)
	EncodeRdsGlobalCluster_EngineVersion(r.Spec.ForProvider, ctyVal)
	EncodeRdsGlobalCluster_Arn(r.Status.AtProvider, ctyVal)
	EncodeRdsGlobalCluster_GlobalClusterMembers(r.Status.AtProvider.GlobalClusterMembers, ctyVal)
	EncodeRdsGlobalCluster_GlobalClusterResourceId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeRdsGlobalCluster_SourceDbClusterIdentifier(p RdsGlobalClusterParameters, vals map[string]cty.Value) {
	vals["source_db_cluster_identifier"] = cty.StringVal(p.SourceDbClusterIdentifier)
}

func EncodeRdsGlobalCluster_StorageEncrypted(p RdsGlobalClusterParameters, vals map[string]cty.Value) {
	vals["storage_encrypted"] = cty.BoolVal(p.StorageEncrypted)
}

func EncodeRdsGlobalCluster_ForceDestroy(p RdsGlobalClusterParameters, vals map[string]cty.Value) {
	vals["force_destroy"] = cty.BoolVal(p.ForceDestroy)
}

func EncodeRdsGlobalCluster_GlobalClusterIdentifier(p RdsGlobalClusterParameters, vals map[string]cty.Value) {
	vals["global_cluster_identifier"] = cty.StringVal(p.GlobalClusterIdentifier)
}

func EncodeRdsGlobalCluster_Id(p RdsGlobalClusterParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRdsGlobalCluster_DatabaseName(p RdsGlobalClusterParameters, vals map[string]cty.Value) {
	vals["database_name"] = cty.StringVal(p.DatabaseName)
}

func EncodeRdsGlobalCluster_DeletionProtection(p RdsGlobalClusterParameters, vals map[string]cty.Value) {
	vals["deletion_protection"] = cty.BoolVal(p.DeletionProtection)
}

func EncodeRdsGlobalCluster_Engine(p RdsGlobalClusterParameters, vals map[string]cty.Value) {
	vals["engine"] = cty.StringVal(p.Engine)
}

func EncodeRdsGlobalCluster_EngineVersion(p RdsGlobalClusterParameters, vals map[string]cty.Value) {
	vals["engine_version"] = cty.StringVal(p.EngineVersion)
}

func EncodeRdsGlobalCluster_Arn(p RdsGlobalClusterObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeRdsGlobalCluster_GlobalClusterMembers(p []GlobalClusterMembers, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeRdsGlobalCluster_GlobalClusterMembers_DbClusterArn(v, ctyVal)
		EncodeRdsGlobalCluster_GlobalClusterMembers_IsWriter(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["global_cluster_members"] = cty.SetVal(valsForCollection)
}

func EncodeRdsGlobalCluster_GlobalClusterMembers_DbClusterArn(p GlobalClusterMembers, vals map[string]cty.Value) {
	vals["db_cluster_arn"] = cty.StringVal(p.DbClusterArn)
}

func EncodeRdsGlobalCluster_GlobalClusterMembers_IsWriter(p GlobalClusterMembers, vals map[string]cty.Value) {
	vals["is_writer"] = cty.BoolVal(p.IsWriter)
}

func EncodeRdsGlobalCluster_GlobalClusterResourceId(p RdsGlobalClusterObservation, vals map[string]cty.Value) {
	vals["global_cluster_resource_id"] = cty.StringVal(p.GlobalClusterResourceId)
}