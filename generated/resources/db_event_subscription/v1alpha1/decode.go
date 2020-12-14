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
	r, ok := mr.(*DbEventSubscription)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDbEventSubscription(r, ctyValue)
}

func DecodeDbEventSubscription(prev *DbEventSubscription, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDbEventSubscription_EventCategories(&new.Spec.ForProvider, valMap)
	DecodeDbEventSubscription_Name(&new.Spec.ForProvider, valMap)
	DecodeDbEventSubscription_NamePrefix(&new.Spec.ForProvider, valMap)
	DecodeDbEventSubscription_SnsTopic(&new.Spec.ForProvider, valMap)
	DecodeDbEventSubscription_Tags(&new.Spec.ForProvider, valMap)
	DecodeDbEventSubscription_Enabled(&new.Spec.ForProvider, valMap)
	DecodeDbEventSubscription_Id(&new.Spec.ForProvider, valMap)
	DecodeDbEventSubscription_SourceIds(&new.Spec.ForProvider, valMap)
	DecodeDbEventSubscription_SourceType(&new.Spec.ForProvider, valMap)
	DecodeDbEventSubscription_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeDbEventSubscription_Arn(&new.Status.AtProvider, valMap)
	DecodeDbEventSubscription_CustomerAwsId(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeDbEventSubscription_EventCategories(p *DbEventSubscriptionParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["event_categories"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.EventCategories = goVals
}

func DecodeDbEventSubscription_Name(p *DbEventSubscriptionParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

func DecodeDbEventSubscription_NamePrefix(p *DbEventSubscriptionParameters, vals map[string]cty.Value) {
	p.NamePrefix = ctwhy.ValueAsString(vals["name_prefix"])
}

func DecodeDbEventSubscription_SnsTopic(p *DbEventSubscriptionParameters, vals map[string]cty.Value) {
	p.SnsTopic = ctwhy.ValueAsString(vals["sns_topic"])
}

func DecodeDbEventSubscription_Tags(p *DbEventSubscriptionParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeDbEventSubscription_Enabled(p *DbEventSubscriptionParameters, vals map[string]cty.Value) {
	p.Enabled = ctwhy.ValueAsBool(vals["enabled"])
}

func DecodeDbEventSubscription_Id(p *DbEventSubscriptionParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeDbEventSubscription_SourceIds(p *DbEventSubscriptionParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["source_ids"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.SourceIds = goVals
}

func DecodeDbEventSubscription_SourceType(p *DbEventSubscriptionParameters, vals map[string]cty.Value) {
	p.SourceType = ctwhy.ValueAsString(vals["source_type"])
}

func DecodeDbEventSubscription_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeDbEventSubscription_Timeouts_Create(p, valMap)
	DecodeDbEventSubscription_Timeouts_Delete(p, valMap)
	DecodeDbEventSubscription_Timeouts_Update(p, valMap)
}

func DecodeDbEventSubscription_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

func DecodeDbEventSubscription_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

func DecodeDbEventSubscription_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}

func DecodeDbEventSubscription_Arn(p *DbEventSubscriptionObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

func DecodeDbEventSubscription_CustomerAwsId(p *DbEventSubscriptionObservation, vals map[string]cty.Value) {
	p.CustomerAwsId = ctwhy.ValueAsString(vals["customer_aws_id"])
}