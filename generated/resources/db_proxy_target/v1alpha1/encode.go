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

func EncodeDbProxyTarget(r DbProxyTarget) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDbProxyTarget_Id(r.Spec.ForProvider, ctyVal)
	EncodeDbProxyTarget_TargetGroupName(r.Spec.ForProvider, ctyVal)
	EncodeDbProxyTarget_DbClusterIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeDbProxyTarget_DbInstanceIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeDbProxyTarget_DbProxyName(r.Spec.ForProvider, ctyVal)
	EncodeDbProxyTarget_Port(r.Status.AtProvider, ctyVal)
	EncodeDbProxyTarget_TargetArn(r.Status.AtProvider, ctyVal)
	EncodeDbProxyTarget_Type(r.Status.AtProvider, ctyVal)
	EncodeDbProxyTarget_Endpoint(r.Status.AtProvider, ctyVal)
	EncodeDbProxyTarget_RdsResourceId(r.Status.AtProvider, ctyVal)
	EncodeDbProxyTarget_TrackedClusterId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDbProxyTarget_Id(p DbProxyTargetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDbProxyTarget_TargetGroupName(p DbProxyTargetParameters, vals map[string]cty.Value) {
	vals["target_group_name"] = cty.StringVal(p.TargetGroupName)
}

func EncodeDbProxyTarget_DbClusterIdentifier(p DbProxyTargetParameters, vals map[string]cty.Value) {
	vals["db_cluster_identifier"] = cty.StringVal(p.DbClusterIdentifier)
}

func EncodeDbProxyTarget_DbInstanceIdentifier(p DbProxyTargetParameters, vals map[string]cty.Value) {
	vals["db_instance_identifier"] = cty.StringVal(p.DbInstanceIdentifier)
}

func EncodeDbProxyTarget_DbProxyName(p DbProxyTargetParameters, vals map[string]cty.Value) {
	vals["db_proxy_name"] = cty.StringVal(p.DbProxyName)
}

func EncodeDbProxyTarget_Port(p DbProxyTargetObservation, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeDbProxyTarget_TargetArn(p DbProxyTargetObservation, vals map[string]cty.Value) {
	vals["target_arn"] = cty.StringVal(p.TargetArn)
}

func EncodeDbProxyTarget_Type(p DbProxyTargetObservation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeDbProxyTarget_Endpoint(p DbProxyTargetObservation, vals map[string]cty.Value) {
	vals["endpoint"] = cty.StringVal(p.Endpoint)
}

func EncodeDbProxyTarget_RdsResourceId(p DbProxyTargetObservation, vals map[string]cty.Value) {
	vals["rds_resource_id"] = cty.StringVal(p.RdsResourceId)
}

func EncodeDbProxyTarget_TrackedClusterId(p DbProxyTargetObservation, vals map[string]cty.Value) {
	vals["tracked_cluster_id"] = cty.StringVal(p.TrackedClusterId)
}