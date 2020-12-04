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
	"github.com/zclconf/go-cty/cty"
)

func EncodeCodepipelineWebhook(r CodepipelineWebhook) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCodepipelineWebhook_Authentication(r.Spec.ForProvider, ctyVal)
	EncodeCodepipelineWebhook_Id(r.Spec.ForProvider, ctyVal)
	EncodeCodepipelineWebhook_Name(r.Spec.ForProvider, ctyVal)
	EncodeCodepipelineWebhook_Tags(r.Spec.ForProvider, ctyVal)
	EncodeCodepipelineWebhook_TargetAction(r.Spec.ForProvider, ctyVal)
	EncodeCodepipelineWebhook_TargetPipeline(r.Spec.ForProvider, ctyVal)
	EncodeCodepipelineWebhook_AuthenticationConfiguration(r.Spec.ForProvider.AuthenticationConfiguration, ctyVal)
	EncodeCodepipelineWebhook_Filter(r.Spec.ForProvider.Filter, ctyVal)
	EncodeCodepipelineWebhook_Url(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeCodepipelineWebhook_Authentication(p CodepipelineWebhookParameters, vals map[string]cty.Value) {
	vals["authentication"] = cty.StringVal(p.Authentication)
}

func EncodeCodepipelineWebhook_Id(p CodepipelineWebhookParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCodepipelineWebhook_Name(p CodepipelineWebhookParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCodepipelineWebhook_Tags(p CodepipelineWebhookParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeCodepipelineWebhook_TargetAction(p CodepipelineWebhookParameters, vals map[string]cty.Value) {
	vals["target_action"] = cty.StringVal(p.TargetAction)
}

func EncodeCodepipelineWebhook_TargetPipeline(p CodepipelineWebhookParameters, vals map[string]cty.Value) {
	vals["target_pipeline"] = cty.StringVal(p.TargetPipeline)
}

func EncodeCodepipelineWebhook_AuthenticationConfiguration(p AuthenticationConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCodepipelineWebhook_AuthenticationConfiguration_AllowedIpRange(p, ctyVal)
	EncodeCodepipelineWebhook_AuthenticationConfiguration_SecretToken(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["authentication_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeCodepipelineWebhook_AuthenticationConfiguration_AllowedIpRange(p AuthenticationConfiguration, vals map[string]cty.Value) {
	vals["allowed_ip_range"] = cty.StringVal(p.AllowedIpRange)
}

func EncodeCodepipelineWebhook_AuthenticationConfiguration_SecretToken(p AuthenticationConfiguration, vals map[string]cty.Value) {
	vals["secret_token"] = cty.StringVal(p.SecretToken)
}

func EncodeCodepipelineWebhook_Filter(p []Filter, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeCodepipelineWebhook_Filter_JsonPath(v, ctyVal)
		EncodeCodepipelineWebhook_Filter_MatchEquals(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["filter"] = cty.SetVal(valsForCollection)
}

func EncodeCodepipelineWebhook_Filter_JsonPath(p Filter, vals map[string]cty.Value) {
	vals["json_path"] = cty.StringVal(p.JsonPath)
}

func EncodeCodepipelineWebhook_Filter_MatchEquals(p Filter, vals map[string]cty.Value) {
	vals["match_equals"] = cty.StringVal(p.MatchEquals)
}

func EncodeCodepipelineWebhook_Url(p CodepipelineWebhookObservation, vals map[string]cty.Value) {
	vals["url"] = cty.StringVal(p.Url)
}