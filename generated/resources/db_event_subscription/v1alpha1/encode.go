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
	r, ok := mr.(*DbEventSubscription)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DbEventSubscription.")
	}
	return EncodeDbEventSubscription(*r), nil
}

func EncodeDbEventSubscription(r DbEventSubscription) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDbEventSubscription_SnsTopic(r.Spec.ForProvider, ctyVal)
	EncodeDbEventSubscription_SourceIds(r.Spec.ForProvider, ctyVal)
	EncodeDbEventSubscription_SourceType(r.Spec.ForProvider, ctyVal)
	EncodeDbEventSubscription_Id(r.Spec.ForProvider, ctyVal)
	EncodeDbEventSubscription_Name(r.Spec.ForProvider, ctyVal)
	EncodeDbEventSubscription_Enabled(r.Spec.ForProvider, ctyVal)
	EncodeDbEventSubscription_EventCategories(r.Spec.ForProvider, ctyVal)
	EncodeDbEventSubscription_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeDbEventSubscription_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDbEventSubscription_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDbEventSubscription_Arn(r.Status.AtProvider, ctyVal)
	EncodeDbEventSubscription_CustomerAwsId(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeDbEventSubscription_SnsTopic(p DbEventSubscriptionParameters, vals map[string]cty.Value) {
	vals["sns_topic"] = cty.StringVal(p.SnsTopic)
}

func EncodeDbEventSubscription_SourceIds(p DbEventSubscriptionParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SourceIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["source_ids"] = cty.SetVal(colVals)
}

func EncodeDbEventSubscription_SourceType(p DbEventSubscriptionParameters, vals map[string]cty.Value) {
	vals["source_type"] = cty.StringVal(p.SourceType)
}

func EncodeDbEventSubscription_Id(p DbEventSubscriptionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDbEventSubscription_Name(p DbEventSubscriptionParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDbEventSubscription_Enabled(p DbEventSubscriptionParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeDbEventSubscription_EventCategories(p DbEventSubscriptionParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.EventCategories {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["event_categories"] = cty.SetVal(colVals)
}

func EncodeDbEventSubscription_NamePrefix(p DbEventSubscriptionParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeDbEventSubscription_Tags(p DbEventSubscriptionParameters, vals map[string]cty.Value) {
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

func EncodeDbEventSubscription_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeDbEventSubscription_Timeouts_Delete(p, ctyVal)
	EncodeDbEventSubscription_Timeouts_Update(p, ctyVal)
	EncodeDbEventSubscription_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDbEventSubscription_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeDbEventSubscription_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeDbEventSubscription_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDbEventSubscription_Arn(p DbEventSubscriptionObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeDbEventSubscription_CustomerAwsId(p DbEventSubscriptionObservation, vals map[string]cty.Value) {
	vals["customer_aws_id"] = cty.StringVal(p.CustomerAwsId)
}