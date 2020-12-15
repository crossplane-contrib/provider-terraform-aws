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
	r, ok := mr.(*XraySamplingRule)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeXraySamplingRule(r, ctyValue)
}

func DecodeXraySamplingRule(prev *XraySamplingRule, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeXraySamplingRule_Priority(&new.Spec.ForProvider, valMap)
	DecodeXraySamplingRule_ReservoirSize(&new.Spec.ForProvider, valMap)
	DecodeXraySamplingRule_ServiceType(&new.Spec.ForProvider, valMap)
	DecodeXraySamplingRule_FixedRate(&new.Spec.ForProvider, valMap)
	DecodeXraySamplingRule_Host(&new.Spec.ForProvider, valMap)
	DecodeXraySamplingRule_HttpMethod(&new.Spec.ForProvider, valMap)
	DecodeXraySamplingRule_Version(&new.Spec.ForProvider, valMap)
	DecodeXraySamplingRule_RuleName(&new.Spec.ForProvider, valMap)
	DecodeXraySamplingRule_ServiceName(&new.Spec.ForProvider, valMap)
	DecodeXraySamplingRule_Tags(&new.Spec.ForProvider, valMap)
	DecodeXraySamplingRule_Attributes(&new.Spec.ForProvider, valMap)
	DecodeXraySamplingRule_ResourceArn(&new.Spec.ForProvider, valMap)
	DecodeXraySamplingRule_UrlPath(&new.Spec.ForProvider, valMap)
	DecodeXraySamplingRule_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeXraySamplingRule_Priority(p *XraySamplingRuleParameters, vals map[string]cty.Value) {
	p.Priority = ctwhy.ValueAsInt64(vals["priority"])
}

//primitiveTypeDecodeTemplate
func DecodeXraySamplingRule_ReservoirSize(p *XraySamplingRuleParameters, vals map[string]cty.Value) {
	p.ReservoirSize = ctwhy.ValueAsInt64(vals["reservoir_size"])
}

//primitiveTypeDecodeTemplate
func DecodeXraySamplingRule_ServiceType(p *XraySamplingRuleParameters, vals map[string]cty.Value) {
	p.ServiceType = ctwhy.ValueAsString(vals["service_type"])
}

//primitiveTypeDecodeTemplate
func DecodeXraySamplingRule_FixedRate(p *XraySamplingRuleParameters, vals map[string]cty.Value) {
	p.FixedRate = ctwhy.ValueAsInt64(vals["fixed_rate"])
}

//primitiveTypeDecodeTemplate
func DecodeXraySamplingRule_Host(p *XraySamplingRuleParameters, vals map[string]cty.Value) {
	p.Host = ctwhy.ValueAsString(vals["host"])
}

//primitiveTypeDecodeTemplate
func DecodeXraySamplingRule_HttpMethod(p *XraySamplingRuleParameters, vals map[string]cty.Value) {
	p.HttpMethod = ctwhy.ValueAsString(vals["http_method"])
}

//primitiveTypeDecodeTemplate
func DecodeXraySamplingRule_Version(p *XraySamplingRuleParameters, vals map[string]cty.Value) {
	p.Version = ctwhy.ValueAsInt64(vals["version"])
}

//primitiveTypeDecodeTemplate
func DecodeXraySamplingRule_RuleName(p *XraySamplingRuleParameters, vals map[string]cty.Value) {
	p.RuleName = ctwhy.ValueAsString(vals["rule_name"])
}

//primitiveTypeDecodeTemplate
func DecodeXraySamplingRule_ServiceName(p *XraySamplingRuleParameters, vals map[string]cty.Value) {
	p.ServiceName = ctwhy.ValueAsString(vals["service_name"])
}

//primitiveMapTypeDecodeTemplate
func DecodeXraySamplingRule_Tags(p *XraySamplingRuleParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveMapTypeDecodeTemplate
func DecodeXraySamplingRule_Attributes(p *XraySamplingRuleParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["attributes"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Attributes = vMap
}

//primitiveTypeDecodeTemplate
func DecodeXraySamplingRule_ResourceArn(p *XraySamplingRuleParameters, vals map[string]cty.Value) {
	p.ResourceArn = ctwhy.ValueAsString(vals["resource_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeXraySamplingRule_UrlPath(p *XraySamplingRuleParameters, vals map[string]cty.Value) {
	p.UrlPath = ctwhy.ValueAsString(vals["url_path"])
}

//primitiveTypeDecodeTemplate
func DecodeXraySamplingRule_Arn(p *XraySamplingRuleObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}