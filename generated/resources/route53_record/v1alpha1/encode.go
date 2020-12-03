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

package v1alpha1func EncodeRoute53Record(r Route53Record) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeRoute53Record_AllowOverwrite(r.Spec.ForProvider, ctyVal)
	EncodeRoute53Record_HealthCheckId(r.Spec.ForProvider, ctyVal)
	EncodeRoute53Record_Id(r.Spec.ForProvider, ctyVal)
	EncodeRoute53Record_ZoneId(r.Spec.ForProvider, ctyVal)
	EncodeRoute53Record_MultivalueAnswerRoutingPolicy(r.Spec.ForProvider, ctyVal)
	EncodeRoute53Record_Name(r.Spec.ForProvider, ctyVal)
	EncodeRoute53Record_Records(r.Spec.ForProvider, ctyVal)
	EncodeRoute53Record_SetIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeRoute53Record_Ttl(r.Spec.ForProvider, ctyVal)
	EncodeRoute53Record_Type(r.Spec.ForProvider, ctyVal)
	EncodeRoute53Record_LatencyRoutingPolicy(r.Spec.ForProvider.LatencyRoutingPolicy, ctyVal)
	EncodeRoute53Record_WeightedRoutingPolicy(r.Spec.ForProvider.WeightedRoutingPolicy, ctyVal)
	EncodeRoute53Record_Alias(r.Spec.ForProvider.Alias, ctyVal)
	EncodeRoute53Record_FailoverRoutingPolicy(r.Spec.ForProvider.FailoverRoutingPolicy, ctyVal)
	EncodeRoute53Record_GeolocationRoutingPolicy(r.Spec.ForProvider.GeolocationRoutingPolicy, ctyVal)
	EncodeRoute53Record_Fqdn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeRoute53Record_AllowOverwrite(p *Route53RecordParameters, vals map[string]cty.Value) {
	vals["allow_overwrite"] = cty.BoolVal(p.AllowOverwrite)
}

func EncodeRoute53Record_HealthCheckId(p *Route53RecordParameters, vals map[string]cty.Value) {
	vals["health_check_id"] = cty.StringVal(p.HealthCheckId)
}

func EncodeRoute53Record_Id(p *Route53RecordParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRoute53Record_ZoneId(p *Route53RecordParameters, vals map[string]cty.Value) {
	vals["zone_id"] = cty.StringVal(p.ZoneId)
}

func EncodeRoute53Record_MultivalueAnswerRoutingPolicy(p *Route53RecordParameters, vals map[string]cty.Value) {
	vals["multivalue_answer_routing_policy"] = cty.BoolVal(p.MultivalueAnswerRoutingPolicy)
}

func EncodeRoute53Record_Name(p *Route53RecordParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeRoute53Record_Records(p *Route53RecordParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Records {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["records"] = cty.SetVal(colVals)
}

func EncodeRoute53Record_SetIdentifier(p *Route53RecordParameters, vals map[string]cty.Value) {
	vals["set_identifier"] = cty.StringVal(p.SetIdentifier)
}

func EncodeRoute53Record_Ttl(p *Route53RecordParameters, vals map[string]cty.Value) {
	vals["ttl"] = cty.IntVal(p.Ttl)
}

func EncodeRoute53Record_Type(p *Route53RecordParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeRoute53Record_LatencyRoutingPolicy(p *LatencyRoutingPolicy, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.LatencyRoutingPolicy {
		ctyVal = make(map[string]cty.Value)
		EncodeRoute53Record_LatencyRoutingPolicy_Region(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["latency_routing_policy"] = cty.ListVal(valsForCollection)
}

func EncodeRoute53Record_LatencyRoutingPolicy_Region(p *LatencyRoutingPolicy, vals map[string]cty.Value) {
	vals["region"] = cty.StringVal(p.Region)
}

func EncodeRoute53Record_WeightedRoutingPolicy(p *WeightedRoutingPolicy, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.WeightedRoutingPolicy {
		ctyVal = make(map[string]cty.Value)
		EncodeRoute53Record_WeightedRoutingPolicy_Weight(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["weighted_routing_policy"] = cty.ListVal(valsForCollection)
}

func EncodeRoute53Record_WeightedRoutingPolicy_Weight(p *WeightedRoutingPolicy, vals map[string]cty.Value) {
	vals["weight"] = cty.IntVal(p.Weight)
}

func EncodeRoute53Record_Alias(p *Alias, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Alias {
		ctyVal = make(map[string]cty.Value)
		EncodeRoute53Record_Alias_EvaluateTargetHealth(v, ctyVal)
		EncodeRoute53Record_Alias_Name(v, ctyVal)
		EncodeRoute53Record_Alias_ZoneId(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["alias"] = cty.SetVal(valsForCollection)
}

func EncodeRoute53Record_Alias_EvaluateTargetHealth(p *Alias, vals map[string]cty.Value) {
	vals["evaluate_target_health"] = cty.BoolVal(p.EvaluateTargetHealth)
}

func EncodeRoute53Record_Alias_Name(p *Alias, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeRoute53Record_Alias_ZoneId(p *Alias, vals map[string]cty.Value) {
	vals["zone_id"] = cty.StringVal(p.ZoneId)
}

func EncodeRoute53Record_FailoverRoutingPolicy(p *FailoverRoutingPolicy, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FailoverRoutingPolicy {
		ctyVal = make(map[string]cty.Value)
		EncodeRoute53Record_FailoverRoutingPolicy_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["failover_routing_policy"] = cty.ListVal(valsForCollection)
}

func EncodeRoute53Record_FailoverRoutingPolicy_Type(p *FailoverRoutingPolicy, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeRoute53Record_GeolocationRoutingPolicy(p *GeolocationRoutingPolicy, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.GeolocationRoutingPolicy {
		ctyVal = make(map[string]cty.Value)
		EncodeRoute53Record_GeolocationRoutingPolicy_Continent(v, ctyVal)
		EncodeRoute53Record_GeolocationRoutingPolicy_Country(v, ctyVal)
		EncodeRoute53Record_GeolocationRoutingPolicy_Subdivision(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["geolocation_routing_policy"] = cty.ListVal(valsForCollection)
}

func EncodeRoute53Record_GeolocationRoutingPolicy_Continent(p *GeolocationRoutingPolicy, vals map[string]cty.Value) {
	vals["continent"] = cty.StringVal(p.Continent)
}

func EncodeRoute53Record_GeolocationRoutingPolicy_Country(p *GeolocationRoutingPolicy, vals map[string]cty.Value) {
	vals["country"] = cty.StringVal(p.Country)
}

func EncodeRoute53Record_GeolocationRoutingPolicy_Subdivision(p *GeolocationRoutingPolicy, vals map[string]cty.Value) {
	vals["subdivision"] = cty.StringVal(p.Subdivision)
}

func EncodeRoute53Record_Fqdn(p *Route53RecordObservation, vals map[string]cty.Value) {
	vals["fqdn"] = cty.StringVal(p.Fqdn)
}