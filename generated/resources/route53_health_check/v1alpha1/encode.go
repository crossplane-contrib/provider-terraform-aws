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
	r, ok := mr.(*Route53HealthCheck)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Route53HealthCheck.")
	}
	return EncodeRoute53HealthCheck(*r), nil
}

func EncodeRoute53HealthCheck(r Route53HealthCheck) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRoute53HealthCheck_Id(r.Spec.ForProvider, ctyVal)
	EncodeRoute53HealthCheck_InvertHealthcheck(r.Spec.ForProvider, ctyVal)
	EncodeRoute53HealthCheck_Port(r.Spec.ForProvider, ctyVal)
	EncodeRoute53HealthCheck_FailureThreshold(r.Spec.ForProvider, ctyVal)
	EncodeRoute53HealthCheck_Fqdn(r.Spec.ForProvider, ctyVal)
	EncodeRoute53HealthCheck_RequestInterval(r.Spec.ForProvider, ctyVal)
	EncodeRoute53HealthCheck_ChildHealthThreshold(r.Spec.ForProvider, ctyVal)
	EncodeRoute53HealthCheck_CloudwatchAlarmName(r.Spec.ForProvider, ctyVal)
	EncodeRoute53HealthCheck_Disabled(r.Spec.ForProvider, ctyVal)
	EncodeRoute53HealthCheck_IpAddress(r.Spec.ForProvider, ctyVal)
	EncodeRoute53HealthCheck_ResourcePath(r.Spec.ForProvider, ctyVal)
	EncodeRoute53HealthCheck_SearchString(r.Spec.ForProvider, ctyVal)
	EncodeRoute53HealthCheck_ChildHealthchecks(r.Spec.ForProvider, ctyVal)
	EncodeRoute53HealthCheck_CloudwatchAlarmRegion(r.Spec.ForProvider, ctyVal)
	EncodeRoute53HealthCheck_MeasureLatency(r.Spec.ForProvider, ctyVal)
	EncodeRoute53HealthCheck_ReferenceName(r.Spec.ForProvider, ctyVal)
	EncodeRoute53HealthCheck_Regions(r.Spec.ForProvider, ctyVal)
	EncodeRoute53HealthCheck_Tags(r.Spec.ForProvider, ctyVal)
	EncodeRoute53HealthCheck_Type(r.Spec.ForProvider, ctyVal)
	EncodeRoute53HealthCheck_EnableSni(r.Spec.ForProvider, ctyVal)
	EncodeRoute53HealthCheck_InsufficientDataHealthStatus(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeRoute53HealthCheck_Id(p Route53HealthCheckParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRoute53HealthCheck_InvertHealthcheck(p Route53HealthCheckParameters, vals map[string]cty.Value) {
	vals["invert_healthcheck"] = cty.BoolVal(p.InvertHealthcheck)
}

func EncodeRoute53HealthCheck_Port(p Route53HealthCheckParameters, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeRoute53HealthCheck_FailureThreshold(p Route53HealthCheckParameters, vals map[string]cty.Value) {
	vals["failure_threshold"] = cty.NumberIntVal(p.FailureThreshold)
}

func EncodeRoute53HealthCheck_Fqdn(p Route53HealthCheckParameters, vals map[string]cty.Value) {
	vals["fqdn"] = cty.StringVal(p.Fqdn)
}

func EncodeRoute53HealthCheck_RequestInterval(p Route53HealthCheckParameters, vals map[string]cty.Value) {
	vals["request_interval"] = cty.NumberIntVal(p.RequestInterval)
}

func EncodeRoute53HealthCheck_ChildHealthThreshold(p Route53HealthCheckParameters, vals map[string]cty.Value) {
	vals["child_health_threshold"] = cty.NumberIntVal(p.ChildHealthThreshold)
}

func EncodeRoute53HealthCheck_CloudwatchAlarmName(p Route53HealthCheckParameters, vals map[string]cty.Value) {
	vals["cloudwatch_alarm_name"] = cty.StringVal(p.CloudwatchAlarmName)
}

func EncodeRoute53HealthCheck_Disabled(p Route53HealthCheckParameters, vals map[string]cty.Value) {
	vals["disabled"] = cty.BoolVal(p.Disabled)
}

func EncodeRoute53HealthCheck_IpAddress(p Route53HealthCheckParameters, vals map[string]cty.Value) {
	vals["ip_address"] = cty.StringVal(p.IpAddress)
}

func EncodeRoute53HealthCheck_ResourcePath(p Route53HealthCheckParameters, vals map[string]cty.Value) {
	vals["resource_path"] = cty.StringVal(p.ResourcePath)
}

func EncodeRoute53HealthCheck_SearchString(p Route53HealthCheckParameters, vals map[string]cty.Value) {
	vals["search_string"] = cty.StringVal(p.SearchString)
}

func EncodeRoute53HealthCheck_ChildHealthchecks(p Route53HealthCheckParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ChildHealthchecks {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["child_healthchecks"] = cty.SetVal(colVals)
}

func EncodeRoute53HealthCheck_CloudwatchAlarmRegion(p Route53HealthCheckParameters, vals map[string]cty.Value) {
	vals["cloudwatch_alarm_region"] = cty.StringVal(p.CloudwatchAlarmRegion)
}

func EncodeRoute53HealthCheck_MeasureLatency(p Route53HealthCheckParameters, vals map[string]cty.Value) {
	vals["measure_latency"] = cty.BoolVal(p.MeasureLatency)
}

func EncodeRoute53HealthCheck_ReferenceName(p Route53HealthCheckParameters, vals map[string]cty.Value) {
	vals["reference_name"] = cty.StringVal(p.ReferenceName)
}

func EncodeRoute53HealthCheck_Regions(p Route53HealthCheckParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Regions {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["regions"] = cty.SetVal(colVals)
}

func EncodeRoute53HealthCheck_Tags(p Route53HealthCheckParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeRoute53HealthCheck_Type(p Route53HealthCheckParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeRoute53HealthCheck_EnableSni(p Route53HealthCheckParameters, vals map[string]cty.Value) {
	vals["enable_sni"] = cty.BoolVal(p.EnableSni)
}

func EncodeRoute53HealthCheck_InsufficientDataHealthStatus(p Route53HealthCheckParameters, vals map[string]cty.Value) {
	vals["insufficient_data_health_status"] = cty.StringVal(p.InsufficientDataHealthStatus)
}