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
	r, ok := mr.(*Apigatewayv2Deployment)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeApigatewayv2Deployment(r, ctyValue)
}

func DecodeApigatewayv2Deployment(prev *Apigatewayv2Deployment, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeApigatewayv2Deployment_ApiId(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2Deployment_Description(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2Deployment_Id(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2Deployment_Triggers(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2Deployment_AutoDeployed(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeApigatewayv2Deployment_ApiId(p *Apigatewayv2DeploymentParameters, vals map[string]cty.Value) {
	p.ApiId = ctwhy.ValueAsString(vals["api_id"])
}

func DecodeApigatewayv2Deployment_Description(p *Apigatewayv2DeploymentParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

func DecodeApigatewayv2Deployment_Id(p *Apigatewayv2DeploymentParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeApigatewayv2Deployment_Triggers(p *Apigatewayv2DeploymentParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["triggers"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Triggers = vMap
}

func DecodeApigatewayv2Deployment_AutoDeployed(p *Apigatewayv2DeploymentObservation, vals map[string]cty.Value) {
	p.AutoDeployed = ctwhy.ValueAsBool(vals["auto_deployed"])
}