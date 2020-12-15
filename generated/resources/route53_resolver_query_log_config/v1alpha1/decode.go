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
	r, ok := mr.(*Route53ResolverQueryLogConfig)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeRoute53ResolverQueryLogConfig(r, ctyValue)
}

func DecodeRoute53ResolverQueryLogConfig(prev *Route53ResolverQueryLogConfig, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeRoute53ResolverQueryLogConfig_DestinationArn(&new.Spec.ForProvider, valMap)
	DecodeRoute53ResolverQueryLogConfig_Id(&new.Spec.ForProvider, valMap)
	DecodeRoute53ResolverQueryLogConfig_Name(&new.Spec.ForProvider, valMap)
	DecodeRoute53ResolverQueryLogConfig_Tags(&new.Spec.ForProvider, valMap)
	DecodeRoute53ResolverQueryLogConfig_Arn(&new.Status.AtProvider, valMap)
	DecodeRoute53ResolverQueryLogConfig_OwnerId(&new.Status.AtProvider, valMap)
	DecodeRoute53ResolverQueryLogConfig_ShareStatus(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeRoute53ResolverQueryLogConfig_DestinationArn(p *Route53ResolverQueryLogConfigParameters, vals map[string]cty.Value) {
	p.DestinationArn = ctwhy.ValueAsString(vals["destination_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeRoute53ResolverQueryLogConfig_Id(p *Route53ResolverQueryLogConfigParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeRoute53ResolverQueryLogConfig_Name(p *Route53ResolverQueryLogConfigParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveMapTypeDecodeTemplate
func DecodeRoute53ResolverQueryLogConfig_Tags(p *Route53ResolverQueryLogConfigParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeRoute53ResolverQueryLogConfig_Arn(p *Route53ResolverQueryLogConfigObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeRoute53ResolverQueryLogConfig_OwnerId(p *Route53ResolverQueryLogConfigObservation, vals map[string]cty.Value) {
	p.OwnerId = ctwhy.ValueAsString(vals["owner_id"])
}

//primitiveTypeDecodeTemplate
func DecodeRoute53ResolverQueryLogConfig_ShareStatus(p *Route53ResolverQueryLogConfigObservation, vals map[string]cty.Value) {
	p.ShareStatus = ctwhy.ValueAsString(vals["share_status"])
}