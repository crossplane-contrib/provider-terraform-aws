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
	r, ok := mr.(*RedshiftEventSubscription)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeRedshiftEventSubscription(r, ctyValue)
}

func DecodeRedshiftEventSubscription(prev *RedshiftEventSubscription, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeRedshiftEventSubscription_Tags(&new.Spec.ForProvider, valMap)
	DecodeRedshiftEventSubscription_Id(&new.Spec.ForProvider, valMap)
	DecodeRedshiftEventSubscription_Severity(&new.Spec.ForProvider, valMap)
	DecodeRedshiftEventSubscription_SnsTopicArn(&new.Spec.ForProvider, valMap)
	DecodeRedshiftEventSubscription_SourceIds(&new.Spec.ForProvider, valMap)
	DecodeRedshiftEventSubscription_Enabled(&new.Spec.ForProvider, valMap)
	DecodeRedshiftEventSubscription_EventCategories(&new.Spec.ForProvider, valMap)
	DecodeRedshiftEventSubscription_Name(&new.Spec.ForProvider, valMap)
	DecodeRedshiftEventSubscription_SourceType(&new.Spec.ForProvider, valMap)
	DecodeRedshiftEventSubscription_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeRedshiftEventSubscription_CustomerAwsId(&new.Status.AtProvider, valMap)
	DecodeRedshiftEventSubscription_Status(&new.Status.AtProvider, valMap)
	DecodeRedshiftEventSubscription_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveMapTypeDecodeTemplate
func DecodeRedshiftEventSubscription_Tags(p *RedshiftEventSubscriptionParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeRedshiftEventSubscription_Id(p *RedshiftEventSubscriptionParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeRedshiftEventSubscription_Severity(p *RedshiftEventSubscriptionParameters, vals map[string]cty.Value) {
	p.Severity = ctwhy.ValueAsString(vals["severity"])
}

//primitiveTypeDecodeTemplate
func DecodeRedshiftEventSubscription_SnsTopicArn(p *RedshiftEventSubscriptionParameters, vals map[string]cty.Value) {
	p.SnsTopicArn = ctwhy.ValueAsString(vals["sns_topic_arn"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeRedshiftEventSubscription_SourceIds(p *RedshiftEventSubscriptionParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["source_ids"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.SourceIds = goVals
}

//primitiveTypeDecodeTemplate
func DecodeRedshiftEventSubscription_Enabled(p *RedshiftEventSubscriptionParameters, vals map[string]cty.Value) {
	p.Enabled = ctwhy.ValueAsBool(vals["enabled"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeRedshiftEventSubscription_EventCategories(p *RedshiftEventSubscriptionParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["event_categories"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.EventCategories = goVals
}

//primitiveTypeDecodeTemplate
func DecodeRedshiftEventSubscription_Name(p *RedshiftEventSubscriptionParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeRedshiftEventSubscription_SourceType(p *RedshiftEventSubscriptionParameters, vals map[string]cty.Value) {
	p.SourceType = ctwhy.ValueAsString(vals["source_type"])
}

//containerTypeDecodeTemplate
func DecodeRedshiftEventSubscription_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeRedshiftEventSubscription_Timeouts_Delete(p, valMap)
	DecodeRedshiftEventSubscription_Timeouts_Update(p, valMap)
	DecodeRedshiftEventSubscription_Timeouts_Create(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeRedshiftEventSubscription_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

//primitiveTypeDecodeTemplate
func DecodeRedshiftEventSubscription_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}

//primitiveTypeDecodeTemplate
func DecodeRedshiftEventSubscription_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveTypeDecodeTemplate
func DecodeRedshiftEventSubscription_CustomerAwsId(p *RedshiftEventSubscriptionObservation, vals map[string]cty.Value) {
	p.CustomerAwsId = ctwhy.ValueAsString(vals["customer_aws_id"])
}

//primitiveTypeDecodeTemplate
func DecodeRedshiftEventSubscription_Status(p *RedshiftEventSubscriptionObservation, vals map[string]cty.Value) {
	p.Status = ctwhy.ValueAsString(vals["status"])
}

//primitiveTypeDecodeTemplate
func DecodeRedshiftEventSubscription_Arn(p *RedshiftEventSubscriptionObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}