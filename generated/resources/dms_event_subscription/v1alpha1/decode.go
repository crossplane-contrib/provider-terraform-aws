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
	r, ok := mr.(*DmsEventSubscription)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDmsEventSubscription(r, ctyValue)
}

func DecodeDmsEventSubscription(prev *DmsEventSubscription, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDmsEventSubscription_Enabled(&new.Spec.ForProvider, valMap)
	DecodeDmsEventSubscription_EventCategories(&new.Spec.ForProvider, valMap)
	DecodeDmsEventSubscription_Id(&new.Spec.ForProvider, valMap)
	DecodeDmsEventSubscription_SourceType(&new.Spec.ForProvider, valMap)
	DecodeDmsEventSubscription_Name(&new.Spec.ForProvider, valMap)
	DecodeDmsEventSubscription_SnsTopicArn(&new.Spec.ForProvider, valMap)
	DecodeDmsEventSubscription_SourceIds(&new.Spec.ForProvider, valMap)
	DecodeDmsEventSubscription_Tags(&new.Spec.ForProvider, valMap)
	DecodeDmsEventSubscription_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeDmsEventSubscription_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeDmsEventSubscription_Enabled(p *DmsEventSubscriptionParameters, vals map[string]cty.Value) {
	p.Enabled = ctwhy.ValueAsBool(vals["enabled"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeDmsEventSubscription_EventCategories(p *DmsEventSubscriptionParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["event_categories"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.EventCategories = goVals
}

//primitiveTypeDecodeTemplate
func DecodeDmsEventSubscription_Id(p *DmsEventSubscriptionParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeDmsEventSubscription_SourceType(p *DmsEventSubscriptionParameters, vals map[string]cty.Value) {
	p.SourceType = ctwhy.ValueAsString(vals["source_type"])
}

//primitiveTypeDecodeTemplate
func DecodeDmsEventSubscription_Name(p *DmsEventSubscriptionParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeDmsEventSubscription_SnsTopicArn(p *DmsEventSubscriptionParameters, vals map[string]cty.Value) {
	p.SnsTopicArn = ctwhy.ValueAsString(vals["sns_topic_arn"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeDmsEventSubscription_SourceIds(p *DmsEventSubscriptionParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["source_ids"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.SourceIds = goVals
}

//primitiveMapTypeDecodeTemplate
func DecodeDmsEventSubscription_Tags(p *DmsEventSubscriptionParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//containerTypeDecodeTemplate
func DecodeDmsEventSubscription_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeDmsEventSubscription_Timeouts_Create(p, valMap)
	DecodeDmsEventSubscription_Timeouts_Delete(p, valMap)
	DecodeDmsEventSubscription_Timeouts_Update(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeDmsEventSubscription_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveTypeDecodeTemplate
func DecodeDmsEventSubscription_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

//primitiveTypeDecodeTemplate
func DecodeDmsEventSubscription_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}

//primitiveTypeDecodeTemplate
func DecodeDmsEventSubscription_Arn(p *DmsEventSubscriptionObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}