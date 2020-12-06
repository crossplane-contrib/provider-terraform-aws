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

func EncodeRdsClusterEndpoint(r RdsClusterEndpoint) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRdsClusterEndpoint_CustomEndpointType(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterEndpoint_Id(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterEndpoint_StaticMembers(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterEndpoint_ClusterEndpointIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterEndpoint_ClusterIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterEndpoint_ExcludedMembers(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterEndpoint_Tags(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterEndpoint_Arn(r.Status.AtProvider, ctyVal)
	EncodeRdsClusterEndpoint_Endpoint(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeRdsClusterEndpoint_CustomEndpointType(p RdsClusterEndpointParameters, vals map[string]cty.Value) {
	vals["custom_endpoint_type"] = cty.StringVal(p.CustomEndpointType)
}

func EncodeRdsClusterEndpoint_Id(p RdsClusterEndpointParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRdsClusterEndpoint_StaticMembers(p RdsClusterEndpointParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.StaticMembers {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["static_members"] = cty.SetVal(colVals)
}

func EncodeRdsClusterEndpoint_ClusterEndpointIdentifier(p RdsClusterEndpointParameters, vals map[string]cty.Value) {
	vals["cluster_endpoint_identifier"] = cty.StringVal(p.ClusterEndpointIdentifier)
}

func EncodeRdsClusterEndpoint_ClusterIdentifier(p RdsClusterEndpointParameters, vals map[string]cty.Value) {
	vals["cluster_identifier"] = cty.StringVal(p.ClusterIdentifier)
}

func EncodeRdsClusterEndpoint_ExcludedMembers(p RdsClusterEndpointParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ExcludedMembers {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["excluded_members"] = cty.SetVal(colVals)
}

func EncodeRdsClusterEndpoint_Tags(p RdsClusterEndpointParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeRdsClusterEndpoint_Arn(p RdsClusterEndpointObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeRdsClusterEndpoint_Endpoint(p RdsClusterEndpointObservation, vals map[string]cty.Value) {
	vals["endpoint"] = cty.StringVal(p.Endpoint)
}