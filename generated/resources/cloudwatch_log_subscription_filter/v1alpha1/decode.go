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
	r, ok := mr.(*CloudwatchLogSubscriptionFilter)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeCloudwatchLogSubscriptionFilter(r, ctyValue)
}

func DecodeCloudwatchLogSubscriptionFilter(prev *CloudwatchLogSubscriptionFilter, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeCloudwatchLogSubscriptionFilter_LogGroupName(&new.Spec.ForProvider, valMap)
	DecodeCloudwatchLogSubscriptionFilter_Name(&new.Spec.ForProvider, valMap)
	DecodeCloudwatchLogSubscriptionFilter_RoleArn(&new.Spec.ForProvider, valMap)
	DecodeCloudwatchLogSubscriptionFilter_DestinationArn(&new.Spec.ForProvider, valMap)
	DecodeCloudwatchLogSubscriptionFilter_Distribution(&new.Spec.ForProvider, valMap)
	DecodeCloudwatchLogSubscriptionFilter_FilterPattern(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeCloudwatchLogSubscriptionFilter_LogGroupName(p *CloudwatchLogSubscriptionFilterParameters, vals map[string]cty.Value) {
	p.LogGroupName = ctwhy.ValueAsString(vals["log_group_name"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudwatchLogSubscriptionFilter_Name(p *CloudwatchLogSubscriptionFilterParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudwatchLogSubscriptionFilter_RoleArn(p *CloudwatchLogSubscriptionFilterParameters, vals map[string]cty.Value) {
	p.RoleArn = ctwhy.ValueAsString(vals["role_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudwatchLogSubscriptionFilter_DestinationArn(p *CloudwatchLogSubscriptionFilterParameters, vals map[string]cty.Value) {
	p.DestinationArn = ctwhy.ValueAsString(vals["destination_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudwatchLogSubscriptionFilter_Distribution(p *CloudwatchLogSubscriptionFilterParameters, vals map[string]cty.Value) {
	p.Distribution = ctwhy.ValueAsString(vals["distribution"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudwatchLogSubscriptionFilter_FilterPattern(p *CloudwatchLogSubscriptionFilterParameters, vals map[string]cty.Value) {
	p.FilterPattern = ctwhy.ValueAsString(vals["filter_pattern"])
}