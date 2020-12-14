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
	r, ok := mr.(*NeptuneEventSubscription)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeNeptuneEventSubscription(r, ctyValue)
}

func DecodeNeptuneEventSubscription(prev *NeptuneEventSubscription, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeNeptuneEventSubscription_Enabled(&new.Spec.ForProvider, valMap)
	DecodeNeptuneEventSubscription_EventCategories(&new.Spec.ForProvider, valMap)
	DecodeNeptuneEventSubscription_SnsTopicArn(&new.Spec.ForProvider, valMap)
	DecodeNeptuneEventSubscription_SourceIds(&new.Spec.ForProvider, valMap)
	DecodeNeptuneEventSubscription_SourceType(&new.Spec.ForProvider, valMap)
	DecodeNeptuneEventSubscription_Tags(&new.Spec.ForProvider, valMap)
	DecodeNeptuneEventSubscription_Id(&new.Spec.ForProvider, valMap)
	DecodeNeptuneEventSubscription_Name(&new.Spec.ForProvider, valMap)
	DecodeNeptuneEventSubscription_NamePrefix(&new.Spec.ForProvider, valMap)
	DecodeNeptuneEventSubscription_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeNeptuneEventSubscription_Arn(&new.Status.AtProvider, valMap)
	DecodeNeptuneEventSubscription_CustomerAwsId(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeNeptuneEventSubscription_Enabled(p *NeptuneEventSubscriptionParameters, vals map[string]cty.Value) {
	p.Enabled = ctwhy.ValueAsBool(vals["enabled"])
}

func DecodeNeptuneEventSubscription_EventCategories(p *NeptuneEventSubscriptionParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["event_categories"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.EventCategories = goVals
}

func DecodeNeptuneEventSubscription_SnsTopicArn(p *NeptuneEventSubscriptionParameters, vals map[string]cty.Value) {
	p.SnsTopicArn = ctwhy.ValueAsString(vals["sns_topic_arn"])
}

func DecodeNeptuneEventSubscription_SourceIds(p *NeptuneEventSubscriptionParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["source_ids"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.SourceIds = goVals
}

func DecodeNeptuneEventSubscription_SourceType(p *NeptuneEventSubscriptionParameters, vals map[string]cty.Value) {
	p.SourceType = ctwhy.ValueAsString(vals["source_type"])
}

func DecodeNeptuneEventSubscription_Tags(p *NeptuneEventSubscriptionParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeNeptuneEventSubscription_Id(p *NeptuneEventSubscriptionParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeNeptuneEventSubscription_Name(p *NeptuneEventSubscriptionParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

func DecodeNeptuneEventSubscription_NamePrefix(p *NeptuneEventSubscriptionParameters, vals map[string]cty.Value) {
	p.NamePrefix = ctwhy.ValueAsString(vals["name_prefix"])
}

func DecodeNeptuneEventSubscription_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeNeptuneEventSubscription_Timeouts_Create(p, valMap)
	DecodeNeptuneEventSubscription_Timeouts_Delete(p, valMap)
	DecodeNeptuneEventSubscription_Timeouts_Update(p, valMap)
}

func DecodeNeptuneEventSubscription_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

func DecodeNeptuneEventSubscription_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

func DecodeNeptuneEventSubscription_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}

func DecodeNeptuneEventSubscription_Arn(p *NeptuneEventSubscriptionObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

func DecodeNeptuneEventSubscription_CustomerAwsId(p *NeptuneEventSubscriptionObservation, vals map[string]cty.Value) {
	p.CustomerAwsId = ctwhy.ValueAsString(vals["customer_aws_id"])
}