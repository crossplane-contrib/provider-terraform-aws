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
	r, ok := mr.(*Route53Record)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Route53Record.")
	}
	return EncodeRoute53Record(*r), nil
}

func EncodeRoute53Record(r Route53Record) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRoute53Record_HealthCheckId(r.Spec.ForProvider, ctyVal)
	EncodeRoute53Record_MultivalueAnswerRoutingPolicy(r.Spec.ForProvider, ctyVal)
	EncodeRoute53Record_Name(r.Spec.ForProvider, ctyVal)
	EncodeRoute53Record_Records(r.Spec.ForProvider, ctyVal)
	EncodeRoute53Record_Ttl(r.Spec.ForProvider, ctyVal)
	EncodeRoute53Record_Type(r.Spec.ForProvider, ctyVal)
	EncodeRoute53Record_ZoneId(r.Spec.ForProvider, ctyVal)
	EncodeRoute53Record_Id(r.Spec.ForProvider, ctyVal)
	EncodeRoute53Record_SetIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeRoute53Record_AllowOverwrite(r.Spec.ForProvider, ctyVal)
	EncodeRoute53Record_GeolocationRoutingPolicy(r.Spec.ForProvider.GeolocationRoutingPolicy, ctyVal)
	EncodeRoute53Record_LatencyRoutingPolicy(r.Spec.ForProvider.LatencyRoutingPolicy, ctyVal)
	EncodeRoute53Record_WeightedRoutingPolicy(r.Spec.ForProvider.WeightedRoutingPolicy, ctyVal)
	EncodeRoute53Record_Alias(r.Spec.ForProvider.Alias, ctyVal)
	EncodeRoute53Record_FailoverRoutingPolicy(r.Spec.ForProvider.FailoverRoutingPolicy, ctyVal)
	EncodeRoute53Record_Fqdn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeRoute53Record_HealthCheckId(p Route53RecordParameters, vals map[string]cty.Value) {
	vals["health_check_id"] = cty.StringVal(p.HealthCheckId)
}

func EncodeRoute53Record_MultivalueAnswerRoutingPolicy(p Route53RecordParameters, vals map[string]cty.Value) {
	vals["multivalue_answer_routing_policy"] = cty.BoolVal(p.MultivalueAnswerRoutingPolicy)
}

func EncodeRoute53Record_Name(p Route53RecordParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeRoute53Record_Records(p Route53RecordParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Records {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["records"] = cty.SetVal(colVals)
}

func EncodeRoute53Record_Ttl(p Route53RecordParameters, vals map[string]cty.Value) {
	vals["ttl"] = cty.NumberIntVal(p.Ttl)
}

func EncodeRoute53Record_Type(p Route53RecordParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeRoute53Record_ZoneId(p Route53RecordParameters, vals map[string]cty.Value) {
	vals["zone_id"] = cty.StringVal(p.ZoneId)
}

func EncodeRoute53Record_Id(p Route53RecordParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRoute53Record_SetIdentifier(p Route53RecordParameters, vals map[string]cty.Value) {
	vals["set_identifier"] = cty.StringVal(p.SetIdentifier)
}

func EncodeRoute53Record_AllowOverwrite(p Route53RecordParameters, vals map[string]cty.Value) {
	vals["allow_overwrite"] = cty.BoolVal(p.AllowOverwrite)
}

func EncodeRoute53Record_GeolocationRoutingPolicy(p GeolocationRoutingPolicy, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeRoute53Record_GeolocationRoutingPolicy_Continent(p, ctyVal)
	EncodeRoute53Record_GeolocationRoutingPolicy_Country(p, ctyVal)
	EncodeRoute53Record_GeolocationRoutingPolicy_Subdivision(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["geolocation_routing_policy"] = cty.ListVal(valsForCollection)
}

func EncodeRoute53Record_GeolocationRoutingPolicy_Continent(p GeolocationRoutingPolicy, vals map[string]cty.Value) {
	vals["continent"] = cty.StringVal(p.Continent)
}

func EncodeRoute53Record_GeolocationRoutingPolicy_Country(p GeolocationRoutingPolicy, vals map[string]cty.Value) {
	vals["country"] = cty.StringVal(p.Country)
}

func EncodeRoute53Record_GeolocationRoutingPolicy_Subdivision(p GeolocationRoutingPolicy, vals map[string]cty.Value) {
	vals["subdivision"] = cty.StringVal(p.Subdivision)
}

func EncodeRoute53Record_LatencyRoutingPolicy(p LatencyRoutingPolicy, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeRoute53Record_LatencyRoutingPolicy_Region(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["latency_routing_policy"] = cty.ListVal(valsForCollection)
}

func EncodeRoute53Record_LatencyRoutingPolicy_Region(p LatencyRoutingPolicy, vals map[string]cty.Value) {
	vals["region"] = cty.StringVal(p.Region)
}

func EncodeRoute53Record_WeightedRoutingPolicy(p WeightedRoutingPolicy, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeRoute53Record_WeightedRoutingPolicy_Weight(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["weighted_routing_policy"] = cty.ListVal(valsForCollection)
}

func EncodeRoute53Record_WeightedRoutingPolicy_Weight(p WeightedRoutingPolicy, vals map[string]cty.Value) {
	vals["weight"] = cty.NumberIntVal(p.Weight)
}

func EncodeRoute53Record_Alias(p Alias, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeRoute53Record_Alias_EvaluateTargetHealth(p, ctyVal)
	EncodeRoute53Record_Alias_Name(p, ctyVal)
	EncodeRoute53Record_Alias_ZoneId(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["alias"] = cty.SetVal(valsForCollection)
}

func EncodeRoute53Record_Alias_EvaluateTargetHealth(p Alias, vals map[string]cty.Value) {
	vals["evaluate_target_health"] = cty.BoolVal(p.EvaluateTargetHealth)
}

func EncodeRoute53Record_Alias_Name(p Alias, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeRoute53Record_Alias_ZoneId(p Alias, vals map[string]cty.Value) {
	vals["zone_id"] = cty.StringVal(p.ZoneId)
}

func EncodeRoute53Record_FailoverRoutingPolicy(p FailoverRoutingPolicy, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeRoute53Record_FailoverRoutingPolicy_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["failover_routing_policy"] = cty.ListVal(valsForCollection)
}

func EncodeRoute53Record_FailoverRoutingPolicy_Type(p FailoverRoutingPolicy, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeRoute53Record_Fqdn(p Route53RecordObservation, vals map[string]cty.Value) {
	vals["fqdn"] = cty.StringVal(p.Fqdn)
}