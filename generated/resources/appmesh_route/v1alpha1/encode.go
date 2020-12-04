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

func EncodeAppmeshRoute(r AppmeshRoute) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshRoute_MeshOwner(r.Spec.ForProvider, ctyVal)
	EncodeAppmeshRoute_Tags(r.Spec.ForProvider, ctyVal)
	EncodeAppmeshRoute_VirtualRouterName(r.Spec.ForProvider, ctyVal)
	EncodeAppmeshRoute_Id(r.Spec.ForProvider, ctyVal)
	EncodeAppmeshRoute_MeshName(r.Spec.ForProvider, ctyVal)
	EncodeAppmeshRoute_Name(r.Spec.ForProvider, ctyVal)
	EncodeAppmeshRoute_Spec(r.Spec.ForProvider.Spec, ctyVal)
	EncodeAppmeshRoute_ResourceOwner(r.Status.AtProvider, ctyVal)
	EncodeAppmeshRoute_Arn(r.Status.AtProvider, ctyVal)
	EncodeAppmeshRoute_CreatedDate(r.Status.AtProvider, ctyVal)
	EncodeAppmeshRoute_LastUpdatedDate(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeAppmeshRoute_MeshOwner(p AppmeshRouteParameters, vals map[string]cty.Value) {
	vals["mesh_owner"] = cty.StringVal(p.MeshOwner)
}

func EncodeAppmeshRoute_Tags(p AppmeshRouteParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeAppmeshRoute_VirtualRouterName(p AppmeshRouteParameters, vals map[string]cty.Value) {
	vals["virtual_router_name"] = cty.StringVal(p.VirtualRouterName)
}

func EncodeAppmeshRoute_Id(p AppmeshRouteParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAppmeshRoute_MeshName(p AppmeshRouteParameters, vals map[string]cty.Value) {
	vals["mesh_name"] = cty.StringVal(p.MeshName)
}

func EncodeAppmeshRoute_Name(p AppmeshRouteParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAppmeshRoute_Spec(p Spec, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshRoute_Spec_Priority(p, ctyVal)
	EncodeAppmeshRoute_Spec_GrpcRoute(p.GrpcRoute, ctyVal)
	EncodeAppmeshRoute_Spec_Http2Route(p.Http2Route, ctyVal)
	EncodeAppmeshRoute_Spec_HttpRoute(p.HttpRoute, ctyVal)
	EncodeAppmeshRoute_Spec_TcpRoute(p.TcpRoute, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["spec"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_Priority(p Spec, vals map[string]cty.Value) {
	vals["priority"] = cty.NumberIntVal(p.Priority)
}

func EncodeAppmeshRoute_Spec_GrpcRoute(p GrpcRoute, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshRoute_Spec_GrpcRoute_RetryPolicy(p.RetryPolicy, ctyVal)
	EncodeAppmeshRoute_Spec_GrpcRoute_Action(p.Action, ctyVal)
	EncodeAppmeshRoute_Spec_GrpcRoute_Match(p.Match, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["grpc_route"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_GrpcRoute_RetryPolicy(p RetryPolicy, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshRoute_Spec_GrpcRoute_RetryPolicy_GrpcRetryEvents(p, ctyVal)
	EncodeAppmeshRoute_Spec_GrpcRoute_RetryPolicy_HttpRetryEvents(p, ctyVal)
	EncodeAppmeshRoute_Spec_GrpcRoute_RetryPolicy_MaxRetries(p, ctyVal)
	EncodeAppmeshRoute_Spec_GrpcRoute_RetryPolicy_TcpRetryEvents(p, ctyVal)
	EncodeAppmeshRoute_Spec_GrpcRoute_RetryPolicy_PerRetryTimeout(p.PerRetryTimeout, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["retry_policy"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_GrpcRoute_RetryPolicy_GrpcRetryEvents(p RetryPolicy, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.GrpcRetryEvents {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["grpc_retry_events"] = cty.SetVal(colVals)
}

func EncodeAppmeshRoute_Spec_GrpcRoute_RetryPolicy_HttpRetryEvents(p RetryPolicy, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.HttpRetryEvents {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["http_retry_events"] = cty.SetVal(colVals)
}

func EncodeAppmeshRoute_Spec_GrpcRoute_RetryPolicy_MaxRetries(p RetryPolicy, vals map[string]cty.Value) {
	vals["max_retries"] = cty.NumberIntVal(p.MaxRetries)
}

func EncodeAppmeshRoute_Spec_GrpcRoute_RetryPolicy_TcpRetryEvents(p RetryPolicy, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.TcpRetryEvents {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["tcp_retry_events"] = cty.SetVal(colVals)
}

func EncodeAppmeshRoute_Spec_GrpcRoute_RetryPolicy_PerRetryTimeout(p PerRetryTimeout, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshRoute_Spec_GrpcRoute_RetryPolicy_PerRetryTimeout_Unit(p, ctyVal)
	EncodeAppmeshRoute_Spec_GrpcRoute_RetryPolicy_PerRetryTimeout_Value(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["per_retry_timeout"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_GrpcRoute_RetryPolicy_PerRetryTimeout_Unit(p PerRetryTimeout, vals map[string]cty.Value) {
	vals["unit"] = cty.StringVal(p.Unit)
}

func EncodeAppmeshRoute_Spec_GrpcRoute_RetryPolicy_PerRetryTimeout_Value(p PerRetryTimeout, vals map[string]cty.Value) {
	vals["value"] = cty.NumberIntVal(p.Value)
}

func EncodeAppmeshRoute_Spec_GrpcRoute_Action(p Action, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshRoute_Spec_GrpcRoute_Action_WeightedTarget(p.WeightedTarget, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["action"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_GrpcRoute_Action_WeightedTarget(p []WeightedTarget, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeAppmeshRoute_Spec_GrpcRoute_Action_WeightedTarget_VirtualNode(v, ctyVal)
		EncodeAppmeshRoute_Spec_GrpcRoute_Action_WeightedTarget_Weight(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["weighted_target"] = cty.SetVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_GrpcRoute_Action_WeightedTarget_VirtualNode(p WeightedTarget, vals map[string]cty.Value) {
	vals["virtual_node"] = cty.StringVal(p.VirtualNode)
}

func EncodeAppmeshRoute_Spec_GrpcRoute_Action_WeightedTarget_Weight(p WeightedTarget, vals map[string]cty.Value) {
	vals["weight"] = cty.NumberIntVal(p.Weight)
}

func EncodeAppmeshRoute_Spec_GrpcRoute_Match(p Match, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshRoute_Spec_GrpcRoute_Match_MethodName(p, ctyVal)
	EncodeAppmeshRoute_Spec_GrpcRoute_Match_ServiceName(p, ctyVal)
	EncodeAppmeshRoute_Spec_GrpcRoute_Match_Metadata(p.Metadata, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["match"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_GrpcRoute_Match_MethodName(p Match, vals map[string]cty.Value) {
	vals["method_name"] = cty.StringVal(p.MethodName)
}

func EncodeAppmeshRoute_Spec_GrpcRoute_Match_ServiceName(p Match, vals map[string]cty.Value) {
	vals["service_name"] = cty.StringVal(p.ServiceName)
}

func EncodeAppmeshRoute_Spec_GrpcRoute_Match_Metadata(p []Metadata, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeAppmeshRoute_Spec_GrpcRoute_Match_Metadata_Invert(v, ctyVal)
		EncodeAppmeshRoute_Spec_GrpcRoute_Match_Metadata_Name(v, ctyVal)
		EncodeAppmeshRoute_Spec_GrpcRoute_Match_Metadata_Match(v.Match, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["metadata"] = cty.SetVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_GrpcRoute_Match_Metadata_Invert(p Metadata, vals map[string]cty.Value) {
	vals["invert"] = cty.BoolVal(p.Invert)
}

func EncodeAppmeshRoute_Spec_GrpcRoute_Match_Metadata_Name(p Metadata, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAppmeshRoute_Spec_GrpcRoute_Match_Metadata_Match(p Match, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshRoute_Spec_GrpcRoute_Match_Metadata_Match_Prefix(p, ctyVal)
	EncodeAppmeshRoute_Spec_GrpcRoute_Match_Metadata_Match_Regex(p, ctyVal)
	EncodeAppmeshRoute_Spec_GrpcRoute_Match_Metadata_Match_Suffix(p, ctyVal)
	EncodeAppmeshRoute_Spec_GrpcRoute_Match_Metadata_Match_Exact(p, ctyVal)
	EncodeAppmeshRoute_Spec_GrpcRoute_Match_Metadata_Match_Range(p.Range, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["match"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_GrpcRoute_Match_Metadata_Match_Prefix(p Match, vals map[string]cty.Value) {
	vals["prefix"] = cty.StringVal(p.Prefix)
}

func EncodeAppmeshRoute_Spec_GrpcRoute_Match_Metadata_Match_Regex(p Match, vals map[string]cty.Value) {
	vals["regex"] = cty.StringVal(p.Regex)
}

func EncodeAppmeshRoute_Spec_GrpcRoute_Match_Metadata_Match_Suffix(p Match, vals map[string]cty.Value) {
	vals["suffix"] = cty.StringVal(p.Suffix)
}

func EncodeAppmeshRoute_Spec_GrpcRoute_Match_Metadata_Match_Exact(p Match, vals map[string]cty.Value) {
	vals["exact"] = cty.StringVal(p.Exact)
}

func EncodeAppmeshRoute_Spec_GrpcRoute_Match_Metadata_Match_Range(p Range, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshRoute_Spec_GrpcRoute_Match_Metadata_Match_Range_End(p, ctyVal)
	EncodeAppmeshRoute_Spec_GrpcRoute_Match_Metadata_Match_Range_Start(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["range"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_GrpcRoute_Match_Metadata_Match_Range_End(p Range, vals map[string]cty.Value) {
	vals["end"] = cty.NumberIntVal(p.End)
}

func EncodeAppmeshRoute_Spec_GrpcRoute_Match_Metadata_Match_Range_Start(p Range, vals map[string]cty.Value) {
	vals["start"] = cty.NumberIntVal(p.Start)
}

func EncodeAppmeshRoute_Spec_Http2Route(p Http2Route, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshRoute_Spec_Http2Route_Action(p.Action, ctyVal)
	EncodeAppmeshRoute_Spec_Http2Route_Match(p.Match, ctyVal)
	EncodeAppmeshRoute_Spec_Http2Route_RetryPolicy(p.RetryPolicy, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["http2_route"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_Http2Route_Action(p Action, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshRoute_Spec_Http2Route_Action_WeightedTarget(p.WeightedTarget, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["action"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_Http2Route_Action_WeightedTarget(p []WeightedTarget, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeAppmeshRoute_Spec_Http2Route_Action_WeightedTarget_VirtualNode(v, ctyVal)
		EncodeAppmeshRoute_Spec_Http2Route_Action_WeightedTarget_Weight(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["weighted_target"] = cty.SetVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_Http2Route_Action_WeightedTarget_VirtualNode(p WeightedTarget, vals map[string]cty.Value) {
	vals["virtual_node"] = cty.StringVal(p.VirtualNode)
}

func EncodeAppmeshRoute_Spec_Http2Route_Action_WeightedTarget_Weight(p WeightedTarget, vals map[string]cty.Value) {
	vals["weight"] = cty.NumberIntVal(p.Weight)
}

func EncodeAppmeshRoute_Spec_Http2Route_Match(p Match, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshRoute_Spec_Http2Route_Match_Scheme(p, ctyVal)
	EncodeAppmeshRoute_Spec_Http2Route_Match_Method(p, ctyVal)
	EncodeAppmeshRoute_Spec_Http2Route_Match_Prefix(p, ctyVal)
	EncodeAppmeshRoute_Spec_Http2Route_Match_Header(p.Header, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["match"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_Http2Route_Match_Scheme(p Match, vals map[string]cty.Value) {
	vals["scheme"] = cty.StringVal(p.Scheme)
}

func EncodeAppmeshRoute_Spec_Http2Route_Match_Method(p Match, vals map[string]cty.Value) {
	vals["method"] = cty.StringVal(p.Method)
}

func EncodeAppmeshRoute_Spec_Http2Route_Match_Prefix(p Match, vals map[string]cty.Value) {
	vals["prefix"] = cty.StringVal(p.Prefix)
}

func EncodeAppmeshRoute_Spec_Http2Route_Match_Header(p []Header, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeAppmeshRoute_Spec_Http2Route_Match_Header_Name(v, ctyVal)
		EncodeAppmeshRoute_Spec_Http2Route_Match_Header_Invert(v, ctyVal)
		EncodeAppmeshRoute_Spec_Http2Route_Match_Header_Match(v.Match, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["header"] = cty.SetVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_Http2Route_Match_Header_Name(p Header, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAppmeshRoute_Spec_Http2Route_Match_Header_Invert(p Header, vals map[string]cty.Value) {
	vals["invert"] = cty.BoolVal(p.Invert)
}

func EncodeAppmeshRoute_Spec_Http2Route_Match_Header_Match(p Match, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshRoute_Spec_Http2Route_Match_Header_Match_Regex(p, ctyVal)
	EncodeAppmeshRoute_Spec_Http2Route_Match_Header_Match_Suffix(p, ctyVal)
	EncodeAppmeshRoute_Spec_Http2Route_Match_Header_Match_Exact(p, ctyVal)
	EncodeAppmeshRoute_Spec_Http2Route_Match_Header_Match_Prefix(p, ctyVal)
	EncodeAppmeshRoute_Spec_Http2Route_Match_Header_Match_Range(p.Range, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["match"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_Http2Route_Match_Header_Match_Regex(p Match, vals map[string]cty.Value) {
	vals["regex"] = cty.StringVal(p.Regex)
}

func EncodeAppmeshRoute_Spec_Http2Route_Match_Header_Match_Suffix(p Match, vals map[string]cty.Value) {
	vals["suffix"] = cty.StringVal(p.Suffix)
}

func EncodeAppmeshRoute_Spec_Http2Route_Match_Header_Match_Exact(p Match, vals map[string]cty.Value) {
	vals["exact"] = cty.StringVal(p.Exact)
}

func EncodeAppmeshRoute_Spec_Http2Route_Match_Header_Match_Prefix(p Match, vals map[string]cty.Value) {
	vals["prefix"] = cty.StringVal(p.Prefix)
}

func EncodeAppmeshRoute_Spec_Http2Route_Match_Header_Match_Range(p Range, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshRoute_Spec_Http2Route_Match_Header_Match_Range_End(p, ctyVal)
	EncodeAppmeshRoute_Spec_Http2Route_Match_Header_Match_Range_Start(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["range"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_Http2Route_Match_Header_Match_Range_End(p Range, vals map[string]cty.Value) {
	vals["end"] = cty.NumberIntVal(p.End)
}

func EncodeAppmeshRoute_Spec_Http2Route_Match_Header_Match_Range_Start(p Range, vals map[string]cty.Value) {
	vals["start"] = cty.NumberIntVal(p.Start)
}

func EncodeAppmeshRoute_Spec_Http2Route_RetryPolicy(p RetryPolicy, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshRoute_Spec_Http2Route_RetryPolicy_MaxRetries(p, ctyVal)
	EncodeAppmeshRoute_Spec_Http2Route_RetryPolicy_TcpRetryEvents(p, ctyVal)
	EncodeAppmeshRoute_Spec_Http2Route_RetryPolicy_HttpRetryEvents(p, ctyVal)
	EncodeAppmeshRoute_Spec_Http2Route_RetryPolicy_PerRetryTimeout(p.PerRetryTimeout, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["retry_policy"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_Http2Route_RetryPolicy_MaxRetries(p RetryPolicy, vals map[string]cty.Value) {
	vals["max_retries"] = cty.NumberIntVal(p.MaxRetries)
}

func EncodeAppmeshRoute_Spec_Http2Route_RetryPolicy_TcpRetryEvents(p RetryPolicy, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.TcpRetryEvents {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["tcp_retry_events"] = cty.SetVal(colVals)
}

func EncodeAppmeshRoute_Spec_Http2Route_RetryPolicy_HttpRetryEvents(p RetryPolicy, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.HttpRetryEvents {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["http_retry_events"] = cty.SetVal(colVals)
}

func EncodeAppmeshRoute_Spec_Http2Route_RetryPolicy_PerRetryTimeout(p PerRetryTimeout, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshRoute_Spec_Http2Route_RetryPolicy_PerRetryTimeout_Unit(p, ctyVal)
	EncodeAppmeshRoute_Spec_Http2Route_RetryPolicy_PerRetryTimeout_Value(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["per_retry_timeout"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_Http2Route_RetryPolicy_PerRetryTimeout_Unit(p PerRetryTimeout, vals map[string]cty.Value) {
	vals["unit"] = cty.StringVal(p.Unit)
}

func EncodeAppmeshRoute_Spec_Http2Route_RetryPolicy_PerRetryTimeout_Value(p PerRetryTimeout, vals map[string]cty.Value) {
	vals["value"] = cty.NumberIntVal(p.Value)
}

func EncodeAppmeshRoute_Spec_HttpRoute(p HttpRoute, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshRoute_Spec_HttpRoute_Action(p.Action, ctyVal)
	EncodeAppmeshRoute_Spec_HttpRoute_Match(p.Match, ctyVal)
	EncodeAppmeshRoute_Spec_HttpRoute_RetryPolicy(p.RetryPolicy, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["http_route"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_HttpRoute_Action(p Action, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshRoute_Spec_HttpRoute_Action_WeightedTarget(p.WeightedTarget, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["action"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_HttpRoute_Action_WeightedTarget(p []WeightedTarget, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeAppmeshRoute_Spec_HttpRoute_Action_WeightedTarget_VirtualNode(v, ctyVal)
		EncodeAppmeshRoute_Spec_HttpRoute_Action_WeightedTarget_Weight(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["weighted_target"] = cty.SetVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_HttpRoute_Action_WeightedTarget_VirtualNode(p WeightedTarget, vals map[string]cty.Value) {
	vals["virtual_node"] = cty.StringVal(p.VirtualNode)
}

func EncodeAppmeshRoute_Spec_HttpRoute_Action_WeightedTarget_Weight(p WeightedTarget, vals map[string]cty.Value) {
	vals["weight"] = cty.NumberIntVal(p.Weight)
}

func EncodeAppmeshRoute_Spec_HttpRoute_Match(p Match, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshRoute_Spec_HttpRoute_Match_Prefix(p, ctyVal)
	EncodeAppmeshRoute_Spec_HttpRoute_Match_Scheme(p, ctyVal)
	EncodeAppmeshRoute_Spec_HttpRoute_Match_Method(p, ctyVal)
	EncodeAppmeshRoute_Spec_HttpRoute_Match_Header(p.Header, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["match"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_HttpRoute_Match_Prefix(p Match, vals map[string]cty.Value) {
	vals["prefix"] = cty.StringVal(p.Prefix)
}

func EncodeAppmeshRoute_Spec_HttpRoute_Match_Scheme(p Match, vals map[string]cty.Value) {
	vals["scheme"] = cty.StringVal(p.Scheme)
}

func EncodeAppmeshRoute_Spec_HttpRoute_Match_Method(p Match, vals map[string]cty.Value) {
	vals["method"] = cty.StringVal(p.Method)
}

func EncodeAppmeshRoute_Spec_HttpRoute_Match_Header(p []Header, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeAppmeshRoute_Spec_HttpRoute_Match_Header_Invert(v, ctyVal)
		EncodeAppmeshRoute_Spec_HttpRoute_Match_Header_Name(v, ctyVal)
		EncodeAppmeshRoute_Spec_HttpRoute_Match_Header_Match(v.Match, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["header"] = cty.SetVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_HttpRoute_Match_Header_Invert(p Header, vals map[string]cty.Value) {
	vals["invert"] = cty.BoolVal(p.Invert)
}

func EncodeAppmeshRoute_Spec_HttpRoute_Match_Header_Name(p Header, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAppmeshRoute_Spec_HttpRoute_Match_Header_Match(p Match, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshRoute_Spec_HttpRoute_Match_Header_Match_Regex(p, ctyVal)
	EncodeAppmeshRoute_Spec_HttpRoute_Match_Header_Match_Suffix(p, ctyVal)
	EncodeAppmeshRoute_Spec_HttpRoute_Match_Header_Match_Exact(p, ctyVal)
	EncodeAppmeshRoute_Spec_HttpRoute_Match_Header_Match_Prefix(p, ctyVal)
	EncodeAppmeshRoute_Spec_HttpRoute_Match_Header_Match_Range(p.Range, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["match"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_HttpRoute_Match_Header_Match_Regex(p Match, vals map[string]cty.Value) {
	vals["regex"] = cty.StringVal(p.Regex)
}

func EncodeAppmeshRoute_Spec_HttpRoute_Match_Header_Match_Suffix(p Match, vals map[string]cty.Value) {
	vals["suffix"] = cty.StringVal(p.Suffix)
}

func EncodeAppmeshRoute_Spec_HttpRoute_Match_Header_Match_Exact(p Match, vals map[string]cty.Value) {
	vals["exact"] = cty.StringVal(p.Exact)
}

func EncodeAppmeshRoute_Spec_HttpRoute_Match_Header_Match_Prefix(p Match, vals map[string]cty.Value) {
	vals["prefix"] = cty.StringVal(p.Prefix)
}

func EncodeAppmeshRoute_Spec_HttpRoute_Match_Header_Match_Range(p Range, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshRoute_Spec_HttpRoute_Match_Header_Match_Range_End(p, ctyVal)
	EncodeAppmeshRoute_Spec_HttpRoute_Match_Header_Match_Range_Start(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["range"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_HttpRoute_Match_Header_Match_Range_End(p Range, vals map[string]cty.Value) {
	vals["end"] = cty.NumberIntVal(p.End)
}

func EncodeAppmeshRoute_Spec_HttpRoute_Match_Header_Match_Range_Start(p Range, vals map[string]cty.Value) {
	vals["start"] = cty.NumberIntVal(p.Start)
}

func EncodeAppmeshRoute_Spec_HttpRoute_RetryPolicy(p RetryPolicy, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshRoute_Spec_HttpRoute_RetryPolicy_HttpRetryEvents(p, ctyVal)
	EncodeAppmeshRoute_Spec_HttpRoute_RetryPolicy_MaxRetries(p, ctyVal)
	EncodeAppmeshRoute_Spec_HttpRoute_RetryPolicy_TcpRetryEvents(p, ctyVal)
	EncodeAppmeshRoute_Spec_HttpRoute_RetryPolicy_PerRetryTimeout(p.PerRetryTimeout, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["retry_policy"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_HttpRoute_RetryPolicy_HttpRetryEvents(p RetryPolicy, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.HttpRetryEvents {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["http_retry_events"] = cty.SetVal(colVals)
}

func EncodeAppmeshRoute_Spec_HttpRoute_RetryPolicy_MaxRetries(p RetryPolicy, vals map[string]cty.Value) {
	vals["max_retries"] = cty.NumberIntVal(p.MaxRetries)
}

func EncodeAppmeshRoute_Spec_HttpRoute_RetryPolicy_TcpRetryEvents(p RetryPolicy, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.TcpRetryEvents {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["tcp_retry_events"] = cty.SetVal(colVals)
}

func EncodeAppmeshRoute_Spec_HttpRoute_RetryPolicy_PerRetryTimeout(p PerRetryTimeout, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshRoute_Spec_HttpRoute_RetryPolicy_PerRetryTimeout_Unit(p, ctyVal)
	EncodeAppmeshRoute_Spec_HttpRoute_RetryPolicy_PerRetryTimeout_Value(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["per_retry_timeout"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_HttpRoute_RetryPolicy_PerRetryTimeout_Unit(p PerRetryTimeout, vals map[string]cty.Value) {
	vals["unit"] = cty.StringVal(p.Unit)
}

func EncodeAppmeshRoute_Spec_HttpRoute_RetryPolicy_PerRetryTimeout_Value(p PerRetryTimeout, vals map[string]cty.Value) {
	vals["value"] = cty.NumberIntVal(p.Value)
}

func EncodeAppmeshRoute_Spec_TcpRoute(p TcpRoute, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshRoute_Spec_TcpRoute_Action(p.Action, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["tcp_route"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_TcpRoute_Action(p Action, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppmeshRoute_Spec_TcpRoute_Action_WeightedTarget(p.WeightedTarget, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["action"] = cty.ListVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_TcpRoute_Action_WeightedTarget(p []WeightedTarget, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeAppmeshRoute_Spec_TcpRoute_Action_WeightedTarget_VirtualNode(v, ctyVal)
		EncodeAppmeshRoute_Spec_TcpRoute_Action_WeightedTarget_Weight(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["weighted_target"] = cty.SetVal(valsForCollection)
}

func EncodeAppmeshRoute_Spec_TcpRoute_Action_WeightedTarget_VirtualNode(p WeightedTarget, vals map[string]cty.Value) {
	vals["virtual_node"] = cty.StringVal(p.VirtualNode)
}

func EncodeAppmeshRoute_Spec_TcpRoute_Action_WeightedTarget_Weight(p WeightedTarget, vals map[string]cty.Value) {
	vals["weight"] = cty.NumberIntVal(p.Weight)
}

func EncodeAppmeshRoute_ResourceOwner(p AppmeshRouteObservation, vals map[string]cty.Value) {
	vals["resource_owner"] = cty.StringVal(p.ResourceOwner)
}

func EncodeAppmeshRoute_Arn(p AppmeshRouteObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeAppmeshRoute_CreatedDate(p AppmeshRouteObservation, vals map[string]cty.Value) {
	vals["created_date"] = cty.StringVal(p.CreatedDate)
}

func EncodeAppmeshRoute_LastUpdatedDate(p AppmeshRouteObservation, vals map[string]cty.Value) {
	vals["last_updated_date"] = cty.StringVal(p.LastUpdatedDate)
}