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
	r, ok := mr.(*ServicequotasServiceQuota)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeServicequotasServiceQuota(r, ctyValue)
}

func DecodeServicequotasServiceQuota(prev *ServicequotasServiceQuota, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeServicequotasServiceQuota_ServiceCode(&new.Spec.ForProvider, valMap)
	DecodeServicequotasServiceQuota_Value(&new.Spec.ForProvider, valMap)
	DecodeServicequotasServiceQuota_QuotaCode(&new.Spec.ForProvider, valMap)
	DecodeServicequotasServiceQuota_QuotaName(&new.Status.AtProvider, valMap)
	DecodeServicequotasServiceQuota_RequestId(&new.Status.AtProvider, valMap)
	DecodeServicequotasServiceQuota_RequestStatus(&new.Status.AtProvider, valMap)
	DecodeServicequotasServiceQuota_Arn(&new.Status.AtProvider, valMap)
	DecodeServicequotasServiceQuota_DefaultValue(&new.Status.AtProvider, valMap)
	DecodeServicequotasServiceQuota_ServiceName(&new.Status.AtProvider, valMap)
	DecodeServicequotasServiceQuota_Adjustable(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeServicequotasServiceQuota_ServiceCode(p *ServicequotasServiceQuotaParameters, vals map[string]cty.Value) {
	p.ServiceCode = ctwhy.ValueAsString(vals["service_code"])
}

//primitiveTypeDecodeTemplate
func DecodeServicequotasServiceQuota_Value(p *ServicequotasServiceQuotaParameters, vals map[string]cty.Value) {
	p.Value = ctwhy.ValueAsInt64(vals["value"])
}

//primitiveTypeDecodeTemplate
func DecodeServicequotasServiceQuota_QuotaCode(p *ServicequotasServiceQuotaParameters, vals map[string]cty.Value) {
	p.QuotaCode = ctwhy.ValueAsString(vals["quota_code"])
}

//primitiveTypeDecodeTemplate
func DecodeServicequotasServiceQuota_QuotaName(p *ServicequotasServiceQuotaObservation, vals map[string]cty.Value) {
	p.QuotaName = ctwhy.ValueAsString(vals["quota_name"])
}

//primitiveTypeDecodeTemplate
func DecodeServicequotasServiceQuota_RequestId(p *ServicequotasServiceQuotaObservation, vals map[string]cty.Value) {
	p.RequestId = ctwhy.ValueAsString(vals["request_id"])
}

//primitiveTypeDecodeTemplate
func DecodeServicequotasServiceQuota_RequestStatus(p *ServicequotasServiceQuotaObservation, vals map[string]cty.Value) {
	p.RequestStatus = ctwhy.ValueAsString(vals["request_status"])
}

//primitiveTypeDecodeTemplate
func DecodeServicequotasServiceQuota_Arn(p *ServicequotasServiceQuotaObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeServicequotasServiceQuota_DefaultValue(p *ServicequotasServiceQuotaObservation, vals map[string]cty.Value) {
	p.DefaultValue = ctwhy.ValueAsInt64(vals["default_value"])
}

//primitiveTypeDecodeTemplate
func DecodeServicequotasServiceQuota_ServiceName(p *ServicequotasServiceQuotaObservation, vals map[string]cty.Value) {
	p.ServiceName = ctwhy.ValueAsString(vals["service_name"])
}

//primitiveTypeDecodeTemplate
func DecodeServicequotasServiceQuota_Adjustable(p *ServicequotasServiceQuotaObservation, vals map[string]cty.Value) {
	p.Adjustable = ctwhy.ValueAsBool(vals["adjustable"])
}