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
	r, ok := mr.(*CloudwatchLogSubscriptionFilter)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a CloudwatchLogSubscriptionFilter.")
	}
	return EncodeCloudwatchLogSubscriptionFilter(*r), nil
}

func EncodeCloudwatchLogSubscriptionFilter(r CloudwatchLogSubscriptionFilter) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCloudwatchLogSubscriptionFilter_LogGroupName(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogSubscriptionFilter_Name(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogSubscriptionFilter_RoleArn(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogSubscriptionFilter_DestinationArn(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogSubscriptionFilter_Distribution(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogSubscriptionFilter_FilterPattern(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeCloudwatchLogSubscriptionFilter_LogGroupName(p CloudwatchLogSubscriptionFilterParameters, vals map[string]cty.Value) {
	vals["log_group_name"] = cty.StringVal(p.LogGroupName)
}

func EncodeCloudwatchLogSubscriptionFilter_Name(p CloudwatchLogSubscriptionFilterParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCloudwatchLogSubscriptionFilter_RoleArn(p CloudwatchLogSubscriptionFilterParameters, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeCloudwatchLogSubscriptionFilter_DestinationArn(p CloudwatchLogSubscriptionFilterParameters, vals map[string]cty.Value) {
	vals["destination_arn"] = cty.StringVal(p.DestinationArn)
}

func EncodeCloudwatchLogSubscriptionFilter_Distribution(p CloudwatchLogSubscriptionFilterParameters, vals map[string]cty.Value) {
	vals["distribution"] = cty.StringVal(p.Distribution)
}

func EncodeCloudwatchLogSubscriptionFilter_FilterPattern(p CloudwatchLogSubscriptionFilterParameters, vals map[string]cty.Value) {
	vals["filter_pattern"] = cty.StringVal(p.FilterPattern)
}