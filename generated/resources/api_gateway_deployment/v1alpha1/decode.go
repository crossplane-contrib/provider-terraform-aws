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
	r, ok := mr.(*ApiGatewayDeployment)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeApiGatewayDeployment(r, ctyValue)
}

func DecodeApiGatewayDeployment(prev *ApiGatewayDeployment, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeApiGatewayDeployment_RestApiId(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayDeployment_StageDescription(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayDeployment_StageName(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayDeployment_Triggers(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayDeployment_Description(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayDeployment_Variables(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayDeployment_CreatedDate(&new.Status.AtProvider, valMap)
	DecodeApiGatewayDeployment_ExecutionArn(&new.Status.AtProvider, valMap)
	DecodeApiGatewayDeployment_InvokeUrl(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayDeployment_RestApiId(p *ApiGatewayDeploymentParameters, vals map[string]cty.Value) {
	p.RestApiId = ctwhy.ValueAsString(vals["rest_api_id"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayDeployment_StageDescription(p *ApiGatewayDeploymentParameters, vals map[string]cty.Value) {
	p.StageDescription = ctwhy.ValueAsString(vals["stage_description"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayDeployment_StageName(p *ApiGatewayDeploymentParameters, vals map[string]cty.Value) {
	p.StageName = ctwhy.ValueAsString(vals["stage_name"])
}

//primitiveMapTypeDecodeTemplate
func DecodeApiGatewayDeployment_Triggers(p *ApiGatewayDeploymentParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["triggers"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Triggers = vMap
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayDeployment_Description(p *ApiGatewayDeploymentParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveMapTypeDecodeTemplate
func DecodeApiGatewayDeployment_Variables(p *ApiGatewayDeploymentParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["variables"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Variables = vMap
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayDeployment_CreatedDate(p *ApiGatewayDeploymentObservation, vals map[string]cty.Value) {
	p.CreatedDate = ctwhy.ValueAsString(vals["created_date"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayDeployment_ExecutionArn(p *ApiGatewayDeploymentObservation, vals map[string]cty.Value) {
	p.ExecutionArn = ctwhy.ValueAsString(vals["execution_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayDeployment_InvokeUrl(p *ApiGatewayDeploymentObservation, vals map[string]cty.Value) {
	p.InvokeUrl = ctwhy.ValueAsString(vals["invoke_url"])
}