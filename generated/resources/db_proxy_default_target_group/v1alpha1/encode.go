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
	r, ok := mr.(*DbProxyDefaultTargetGroup)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DbProxyDefaultTargetGroup.")
	}
	return EncodeDbProxyDefaultTargetGroup(*r), nil
}

func EncodeDbProxyDefaultTargetGroup(r DbProxyDefaultTargetGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDbProxyDefaultTargetGroup_DbProxyName(r.Spec.ForProvider, ctyVal)
	EncodeDbProxyDefaultTargetGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeDbProxyDefaultTargetGroup_ConnectionPoolConfig(r.Spec.ForProvider.ConnectionPoolConfig, ctyVal)
	EncodeDbProxyDefaultTargetGroup_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDbProxyDefaultTargetGroup_Name(r.Status.AtProvider, ctyVal)
	EncodeDbProxyDefaultTargetGroup_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDbProxyDefaultTargetGroup_DbProxyName(p DbProxyDefaultTargetGroupParameters, vals map[string]cty.Value) {
	vals["db_proxy_name"] = cty.StringVal(p.DbProxyName)
}

func EncodeDbProxyDefaultTargetGroup_Id(p DbProxyDefaultTargetGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDbProxyDefaultTargetGroup_ConnectionPoolConfig(p ConnectionPoolConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDbProxyDefaultTargetGroup_ConnectionPoolConfig_InitQuery(p, ctyVal)
	EncodeDbProxyDefaultTargetGroup_ConnectionPoolConfig_MaxConnectionsPercent(p, ctyVal)
	EncodeDbProxyDefaultTargetGroup_ConnectionPoolConfig_MaxIdleConnectionsPercent(p, ctyVal)
	EncodeDbProxyDefaultTargetGroup_ConnectionPoolConfig_SessionPinningFilters(p, ctyVal)
	EncodeDbProxyDefaultTargetGroup_ConnectionPoolConfig_ConnectionBorrowTimeout(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["connection_pool_config"] = cty.ListVal(valsForCollection)
}

func EncodeDbProxyDefaultTargetGroup_ConnectionPoolConfig_InitQuery(p ConnectionPoolConfig, vals map[string]cty.Value) {
	vals["init_query"] = cty.StringVal(p.InitQuery)
}

func EncodeDbProxyDefaultTargetGroup_ConnectionPoolConfig_MaxConnectionsPercent(p ConnectionPoolConfig, vals map[string]cty.Value) {
	vals["max_connections_percent"] = cty.NumberIntVal(p.MaxConnectionsPercent)
}

func EncodeDbProxyDefaultTargetGroup_ConnectionPoolConfig_MaxIdleConnectionsPercent(p ConnectionPoolConfig, vals map[string]cty.Value) {
	vals["max_idle_connections_percent"] = cty.NumberIntVal(p.MaxIdleConnectionsPercent)
}

func EncodeDbProxyDefaultTargetGroup_ConnectionPoolConfig_SessionPinningFilters(p ConnectionPoolConfig, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SessionPinningFilters {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["session_pinning_filters"] = cty.SetVal(colVals)
}

func EncodeDbProxyDefaultTargetGroup_ConnectionPoolConfig_ConnectionBorrowTimeout(p ConnectionPoolConfig, vals map[string]cty.Value) {
	vals["connection_borrow_timeout"] = cty.NumberIntVal(p.ConnectionBorrowTimeout)
}

func EncodeDbProxyDefaultTargetGroup_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeDbProxyDefaultTargetGroup_Timeouts_Create(p, ctyVal)
	EncodeDbProxyDefaultTargetGroup_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDbProxyDefaultTargetGroup_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDbProxyDefaultTargetGroup_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeDbProxyDefaultTargetGroup_Name(p DbProxyDefaultTargetGroupObservation, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDbProxyDefaultTargetGroup_Arn(p DbProxyDefaultTargetGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}