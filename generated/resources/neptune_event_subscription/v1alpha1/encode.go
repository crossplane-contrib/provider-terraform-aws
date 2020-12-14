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
	r, ok := mr.(*NeptuneEventSubscription)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a NeptuneEventSubscription.")
	}
	return EncodeNeptuneEventSubscription(*r), nil
}

func EncodeNeptuneEventSubscription(r NeptuneEventSubscription) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeNeptuneEventSubscription_Enabled(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneEventSubscription_EventCategories(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneEventSubscription_SnsTopicArn(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneEventSubscription_SourceIds(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneEventSubscription_SourceType(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneEventSubscription_Tags(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneEventSubscription_Id(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneEventSubscription_Name(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneEventSubscription_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneEventSubscription_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeNeptuneEventSubscription_Arn(r.Status.AtProvider, ctyVal)
	EncodeNeptuneEventSubscription_CustomerAwsId(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeNeptuneEventSubscription_Enabled(p NeptuneEventSubscriptionParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeNeptuneEventSubscription_EventCategories(p NeptuneEventSubscriptionParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.EventCategories {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["event_categories"] = cty.SetVal(colVals)
}

func EncodeNeptuneEventSubscription_SnsTopicArn(p NeptuneEventSubscriptionParameters, vals map[string]cty.Value) {
	vals["sns_topic_arn"] = cty.StringVal(p.SnsTopicArn)
}

func EncodeNeptuneEventSubscription_SourceIds(p NeptuneEventSubscriptionParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SourceIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["source_ids"] = cty.SetVal(colVals)
}

func EncodeNeptuneEventSubscription_SourceType(p NeptuneEventSubscriptionParameters, vals map[string]cty.Value) {
	vals["source_type"] = cty.StringVal(p.SourceType)
}

func EncodeNeptuneEventSubscription_Tags(p NeptuneEventSubscriptionParameters, vals map[string]cty.Value) {
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

func EncodeNeptuneEventSubscription_Id(p NeptuneEventSubscriptionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeNeptuneEventSubscription_Name(p NeptuneEventSubscriptionParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeNeptuneEventSubscription_NamePrefix(p NeptuneEventSubscriptionParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeNeptuneEventSubscription_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeNeptuneEventSubscription_Timeouts_Create(p, ctyVal)
	EncodeNeptuneEventSubscription_Timeouts_Delete(p, ctyVal)
	EncodeNeptuneEventSubscription_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeNeptuneEventSubscription_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeNeptuneEventSubscription_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeNeptuneEventSubscription_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeNeptuneEventSubscription_Arn(p NeptuneEventSubscriptionObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeNeptuneEventSubscription_CustomerAwsId(p NeptuneEventSubscriptionObservation, vals map[string]cty.Value) {
	vals["customer_aws_id"] = cty.StringVal(p.CustomerAwsId)
}