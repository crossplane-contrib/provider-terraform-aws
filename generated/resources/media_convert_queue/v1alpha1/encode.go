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
	r, ok := mr.(*MediaConvertQueue)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a MediaConvertQueue.")
	}
	return EncodeMediaConvertQueue(*r), nil
}

func EncodeMediaConvertQueue(r MediaConvertQueue) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeMediaConvertQueue_Name(r.Spec.ForProvider, ctyVal)
	EncodeMediaConvertQueue_PricingPlan(r.Spec.ForProvider, ctyVal)
	EncodeMediaConvertQueue_Status(r.Spec.ForProvider, ctyVal)
	EncodeMediaConvertQueue_Tags(r.Spec.ForProvider, ctyVal)
	EncodeMediaConvertQueue_Description(r.Spec.ForProvider, ctyVal)
	EncodeMediaConvertQueue_Id(r.Spec.ForProvider, ctyVal)
	EncodeMediaConvertQueue_ReservationPlanSettings(r.Spec.ForProvider.ReservationPlanSettings, ctyVal)
	EncodeMediaConvertQueue_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeMediaConvertQueue_Name(p MediaConvertQueueParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeMediaConvertQueue_PricingPlan(p MediaConvertQueueParameters, vals map[string]cty.Value) {
	vals["pricing_plan"] = cty.StringVal(p.PricingPlan)
}

func EncodeMediaConvertQueue_Status(p MediaConvertQueueParameters, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeMediaConvertQueue_Tags(p MediaConvertQueueParameters, vals map[string]cty.Value) {
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

func EncodeMediaConvertQueue_Description(p MediaConvertQueueParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeMediaConvertQueue_Id(p MediaConvertQueueParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeMediaConvertQueue_ReservationPlanSettings(p ReservationPlanSettings, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeMediaConvertQueue_ReservationPlanSettings_ReservedSlots(p, ctyVal)
	EncodeMediaConvertQueue_ReservationPlanSettings_Commitment(p, ctyVal)
	EncodeMediaConvertQueue_ReservationPlanSettings_RenewalType(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["reservation_plan_settings"] = cty.ListVal(valsForCollection)
}

func EncodeMediaConvertQueue_ReservationPlanSettings_ReservedSlots(p ReservationPlanSettings, vals map[string]cty.Value) {
	vals["reserved_slots"] = cty.NumberIntVal(p.ReservedSlots)
}

func EncodeMediaConvertQueue_ReservationPlanSettings_Commitment(p ReservationPlanSettings, vals map[string]cty.Value) {
	vals["commitment"] = cty.StringVal(p.Commitment)
}

func EncodeMediaConvertQueue_ReservationPlanSettings_RenewalType(p ReservationPlanSettings, vals map[string]cty.Value) {
	vals["renewal_type"] = cty.StringVal(p.RenewalType)
}

func EncodeMediaConvertQueue_Arn(p MediaConvertQueueObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}