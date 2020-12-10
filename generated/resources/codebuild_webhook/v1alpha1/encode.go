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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*CodebuildWebhook)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a CodebuildWebhook.")
	}
	return EncodeCodebuildWebhook(*r), nil
}

func EncodeCodebuildWebhook(r CodebuildWebhook) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCodebuildWebhook_Id(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildWebhook_ProjectName(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildWebhook_BranchFilter(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildWebhook_FilterGroup(r.Spec.ForProvider.FilterGroup, ctyVal)
	EncodeCodebuildWebhook_PayloadUrl(r.Status.AtProvider, ctyVal)
	EncodeCodebuildWebhook_Secret(r.Status.AtProvider, ctyVal)
	EncodeCodebuildWebhook_Url(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeCodebuildWebhook_Id(p CodebuildWebhookParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCodebuildWebhook_ProjectName(p CodebuildWebhookParameters, vals map[string]cty.Value) {
	vals["project_name"] = cty.StringVal(p.ProjectName)
}

func EncodeCodebuildWebhook_BranchFilter(p CodebuildWebhookParameters, vals map[string]cty.Value) {
	vals["branch_filter"] = cty.StringVal(p.BranchFilter)
}

func EncodeCodebuildWebhook_FilterGroup(p FilterGroup, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCodebuildWebhook_FilterGroup_Filter(p.Filter, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["filter_group"] = cty.SetVal(valsForCollection)
}

func EncodeCodebuildWebhook_FilterGroup_Filter(p Filter, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCodebuildWebhook_FilterGroup_Filter_Pattern(p, ctyVal)
	EncodeCodebuildWebhook_FilterGroup_Filter_Type(p, ctyVal)
	EncodeCodebuildWebhook_FilterGroup_Filter_ExcludeMatchedPattern(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["filter"] = cty.ListVal(valsForCollection)
}

func EncodeCodebuildWebhook_FilterGroup_Filter_Pattern(p Filter, vals map[string]cty.Value) {
	vals["pattern"] = cty.StringVal(p.Pattern)
}

func EncodeCodebuildWebhook_FilterGroup_Filter_Type(p Filter, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeCodebuildWebhook_FilterGroup_Filter_ExcludeMatchedPattern(p Filter, vals map[string]cty.Value) {
	vals["exclude_matched_pattern"] = cty.BoolVal(p.ExcludeMatchedPattern)
}

func EncodeCodebuildWebhook_PayloadUrl(p CodebuildWebhookObservation, vals map[string]cty.Value) {
	vals["payload_url"] = cty.StringVal(p.PayloadUrl)
}

func EncodeCodebuildWebhook_Secret(p CodebuildWebhookObservation, vals map[string]cty.Value) {
	vals["secret"] = cty.StringVal(p.Secret)
}

func EncodeCodebuildWebhook_Url(p CodebuildWebhookObservation, vals map[string]cty.Value) {
	vals["url"] = cty.StringVal(p.Url)
}