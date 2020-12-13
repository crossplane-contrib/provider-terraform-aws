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
	r, ok := mr.(*ServicequotasServiceQuota)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ServicequotasServiceQuota.")
	}
	return EncodeServicequotasServiceQuota(*r), nil
}

func EncodeServicequotasServiceQuota(r ServicequotasServiceQuota) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeServicequotasServiceQuota_Id(r.Spec.ForProvider, ctyVal)
	EncodeServicequotasServiceQuota_QuotaCode(r.Spec.ForProvider, ctyVal)
	EncodeServicequotasServiceQuota_Value(r.Spec.ForProvider, ctyVal)
	EncodeServicequotasServiceQuota_ServiceCode(r.Spec.ForProvider, ctyVal)
	EncodeServicequotasServiceQuota_DefaultValue(r.Status.AtProvider, ctyVal)
	EncodeServicequotasServiceQuota_ServiceName(r.Status.AtProvider, ctyVal)
	EncodeServicequotasServiceQuota_Adjustable(r.Status.AtProvider, ctyVal)
	EncodeServicequotasServiceQuota_Arn(r.Status.AtProvider, ctyVal)
	EncodeServicequotasServiceQuota_QuotaName(r.Status.AtProvider, ctyVal)
	EncodeServicequotasServiceQuota_RequestId(r.Status.AtProvider, ctyVal)
	EncodeServicequotasServiceQuota_RequestStatus(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeServicequotasServiceQuota_Id(p ServicequotasServiceQuotaParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeServicequotasServiceQuota_QuotaCode(p ServicequotasServiceQuotaParameters, vals map[string]cty.Value) {
	vals["quota_code"] = cty.StringVal(p.QuotaCode)
}

func EncodeServicequotasServiceQuota_Value(p ServicequotasServiceQuotaParameters, vals map[string]cty.Value) {
	vals["value"] = cty.NumberIntVal(p.Value)
}

func EncodeServicequotasServiceQuota_ServiceCode(p ServicequotasServiceQuotaParameters, vals map[string]cty.Value) {
	vals["service_code"] = cty.StringVal(p.ServiceCode)
}

func EncodeServicequotasServiceQuota_DefaultValue(p ServicequotasServiceQuotaObservation, vals map[string]cty.Value) {
	vals["default_value"] = cty.NumberIntVal(p.DefaultValue)
}

func EncodeServicequotasServiceQuota_ServiceName(p ServicequotasServiceQuotaObservation, vals map[string]cty.Value) {
	vals["service_name"] = cty.StringVal(p.ServiceName)
}

func EncodeServicequotasServiceQuota_Adjustable(p ServicequotasServiceQuotaObservation, vals map[string]cty.Value) {
	vals["adjustable"] = cty.BoolVal(p.Adjustable)
}

func EncodeServicequotasServiceQuota_Arn(p ServicequotasServiceQuotaObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeServicequotasServiceQuota_QuotaName(p ServicequotasServiceQuotaObservation, vals map[string]cty.Value) {
	vals["quota_name"] = cty.StringVal(p.QuotaName)
}

func EncodeServicequotasServiceQuota_RequestId(p ServicequotasServiceQuotaObservation, vals map[string]cty.Value) {
	vals["request_id"] = cty.StringVal(p.RequestId)
}

func EncodeServicequotasServiceQuota_RequestStatus(p ServicequotasServiceQuotaObservation, vals map[string]cty.Value) {
	vals["request_status"] = cty.StringVal(p.RequestStatus)
}