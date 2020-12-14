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

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*Route53HealthCheck)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeRoute53HealthCheck(r, ctyValue)
}

func DecodeRoute53HealthCheck(prev *Route53HealthCheck, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeRoute53HealthCheck_MeasureLatency(&new.Spec.ForProvider, valMap)
	DecodeRoute53HealthCheck_ResourcePath(&new.Spec.ForProvider, valMap)
	DecodeRoute53HealthCheck_Tags(&new.Spec.ForProvider, valMap)
	DecodeRoute53HealthCheck_Type(&new.Spec.ForProvider, valMap)
	DecodeRoute53HealthCheck_ChildHealthchecks(&new.Spec.ForProvider, valMap)
	DecodeRoute53HealthCheck_InvertHealthcheck(&new.Spec.ForProvider, valMap)
	DecodeRoute53HealthCheck_Regions(&new.Spec.ForProvider, valMap)
	DecodeRoute53HealthCheck_SearchString(&new.Spec.ForProvider, valMap)
	DecodeRoute53HealthCheck_ChildHealthThreshold(&new.Spec.ForProvider, valMap)
	DecodeRoute53HealthCheck_EnableSni(&new.Spec.ForProvider, valMap)
	DecodeRoute53HealthCheck_FailureThreshold(&new.Spec.ForProvider, valMap)
	DecodeRoute53HealthCheck_InsufficientDataHealthStatus(&new.Spec.ForProvider, valMap)
	DecodeRoute53HealthCheck_Port(&new.Spec.ForProvider, valMap)
	DecodeRoute53HealthCheck_ReferenceName(&new.Spec.ForProvider, valMap)
	DecodeRoute53HealthCheck_RequestInterval(&new.Spec.ForProvider, valMap)
	DecodeRoute53HealthCheck_CloudwatchAlarmName(&new.Spec.ForProvider, valMap)
	DecodeRoute53HealthCheck_Disabled(&new.Spec.ForProvider, valMap)
	DecodeRoute53HealthCheck_Id(&new.Spec.ForProvider, valMap)
	DecodeRoute53HealthCheck_IpAddress(&new.Spec.ForProvider, valMap)
	DecodeRoute53HealthCheck_CloudwatchAlarmRegion(&new.Spec.ForProvider, valMap)
	DecodeRoute53HealthCheck_Fqdn(&new.Spec.ForProvider, valMap)

	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeRoute53HealthCheck_MeasureLatency(p *Route53HealthCheckParameters, vals map[string]cty.Value) {
	p.MeasureLatency = ctwhy.ValueAsBool(vals["measure_latency"])
}

func DecodeRoute53HealthCheck_ResourcePath(p *Route53HealthCheckParameters, vals map[string]cty.Value) {
	p.ResourcePath = ctwhy.ValueAsString(vals["resource_path"])
}

func DecodeRoute53HealthCheck_Tags(p *Route53HealthCheckParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeRoute53HealthCheck_Type(p *Route53HealthCheckParameters, vals map[string]cty.Value) {
	p.Type = ctwhy.ValueAsString(vals["type"])
}

func DecodeRoute53HealthCheck_ChildHealthchecks(p *Route53HealthCheckParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["child_healthchecks"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.ChildHealthchecks = goVals
}

func DecodeRoute53HealthCheck_InvertHealthcheck(p *Route53HealthCheckParameters, vals map[string]cty.Value) {
	p.InvertHealthcheck = ctwhy.ValueAsBool(vals["invert_healthcheck"])
}

func DecodeRoute53HealthCheck_Regions(p *Route53HealthCheckParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["regions"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.Regions = goVals
}

func DecodeRoute53HealthCheck_SearchString(p *Route53HealthCheckParameters, vals map[string]cty.Value) {
	p.SearchString = ctwhy.ValueAsString(vals["search_string"])
}

func DecodeRoute53HealthCheck_ChildHealthThreshold(p *Route53HealthCheckParameters, vals map[string]cty.Value) {
	p.ChildHealthThreshold = ctwhy.ValueAsInt64(vals["child_health_threshold"])
}

func DecodeRoute53HealthCheck_EnableSni(p *Route53HealthCheckParameters, vals map[string]cty.Value) {
	p.EnableSni = ctwhy.ValueAsBool(vals["enable_sni"])
}

func DecodeRoute53HealthCheck_FailureThreshold(p *Route53HealthCheckParameters, vals map[string]cty.Value) {
	p.FailureThreshold = ctwhy.ValueAsInt64(vals["failure_threshold"])
}

func DecodeRoute53HealthCheck_InsufficientDataHealthStatus(p *Route53HealthCheckParameters, vals map[string]cty.Value) {
	p.InsufficientDataHealthStatus = ctwhy.ValueAsString(vals["insufficient_data_health_status"])
}

func DecodeRoute53HealthCheck_Port(p *Route53HealthCheckParameters, vals map[string]cty.Value) {
	p.Port = ctwhy.ValueAsInt64(vals["port"])
}

func DecodeRoute53HealthCheck_ReferenceName(p *Route53HealthCheckParameters, vals map[string]cty.Value) {
	p.ReferenceName = ctwhy.ValueAsString(vals["reference_name"])
}

func DecodeRoute53HealthCheck_RequestInterval(p *Route53HealthCheckParameters, vals map[string]cty.Value) {
	p.RequestInterval = ctwhy.ValueAsInt64(vals["request_interval"])
}

func DecodeRoute53HealthCheck_CloudwatchAlarmName(p *Route53HealthCheckParameters, vals map[string]cty.Value) {
	p.CloudwatchAlarmName = ctwhy.ValueAsString(vals["cloudwatch_alarm_name"])
}

func DecodeRoute53HealthCheck_Disabled(p *Route53HealthCheckParameters, vals map[string]cty.Value) {
	p.Disabled = ctwhy.ValueAsBool(vals["disabled"])
}

func DecodeRoute53HealthCheck_Id(p *Route53HealthCheckParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeRoute53HealthCheck_IpAddress(p *Route53HealthCheckParameters, vals map[string]cty.Value) {
	p.IpAddress = ctwhy.ValueAsString(vals["ip_address"])
}

func DecodeRoute53HealthCheck_CloudwatchAlarmRegion(p *Route53HealthCheckParameters, vals map[string]cty.Value) {
	p.CloudwatchAlarmRegion = ctwhy.ValueAsString(vals["cloudwatch_alarm_region"])
}

func DecodeRoute53HealthCheck_Fqdn(p *Route53HealthCheckParameters, vals map[string]cty.Value) {
	p.Fqdn = ctwhy.ValueAsString(vals["fqdn"])
}