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
	r, ok := mr.(*SagemakerEndpointConfiguration)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a SagemakerEndpointConfiguration.")
	}
	return EncodeSagemakerEndpointConfiguration(*r), nil
}

func EncodeSagemakerEndpointConfiguration(r SagemakerEndpointConfiguration) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSagemakerEndpointConfiguration_Name(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerEndpointConfiguration_Tags(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerEndpointConfiguration_Id(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerEndpointConfiguration_KmsKeyArn(r.Spec.ForProvider, ctyVal)
	EncodeSagemakerEndpointConfiguration_ProductionVariants(r.Spec.ForProvider.ProductionVariants, ctyVal)
	EncodeSagemakerEndpointConfiguration_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeSagemakerEndpointConfiguration_Name(p SagemakerEndpointConfigurationParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSagemakerEndpointConfiguration_Tags(p SagemakerEndpointConfigurationParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeSagemakerEndpointConfiguration_Id(p SagemakerEndpointConfigurationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSagemakerEndpointConfiguration_KmsKeyArn(p SagemakerEndpointConfigurationParameters, vals map[string]cty.Value) {
	vals["kms_key_arn"] = cty.StringVal(p.KmsKeyArn)
}

func EncodeSagemakerEndpointConfiguration_ProductionVariants(p []ProductionVariants, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeSagemakerEndpointConfiguration_ProductionVariants_ModelName(v, ctyVal)
		EncodeSagemakerEndpointConfiguration_ProductionVariants_VariantName(v, ctyVal)
		EncodeSagemakerEndpointConfiguration_ProductionVariants_AcceleratorType(v, ctyVal)
		EncodeSagemakerEndpointConfiguration_ProductionVariants_InitialInstanceCount(v, ctyVal)
		EncodeSagemakerEndpointConfiguration_ProductionVariants_InitialVariantWeight(v, ctyVal)
		EncodeSagemakerEndpointConfiguration_ProductionVariants_InstanceType(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["production_variants"] = cty.ListVal(valsForCollection)
}

func EncodeSagemakerEndpointConfiguration_ProductionVariants_ModelName(p ProductionVariants, vals map[string]cty.Value) {
	vals["model_name"] = cty.StringVal(p.ModelName)
}

func EncodeSagemakerEndpointConfiguration_ProductionVariants_VariantName(p ProductionVariants, vals map[string]cty.Value) {
	vals["variant_name"] = cty.StringVal(p.VariantName)
}

func EncodeSagemakerEndpointConfiguration_ProductionVariants_AcceleratorType(p ProductionVariants, vals map[string]cty.Value) {
	vals["accelerator_type"] = cty.StringVal(p.AcceleratorType)
}

func EncodeSagemakerEndpointConfiguration_ProductionVariants_InitialInstanceCount(p ProductionVariants, vals map[string]cty.Value) {
	vals["initial_instance_count"] = cty.NumberIntVal(p.InitialInstanceCount)
}

func EncodeSagemakerEndpointConfiguration_ProductionVariants_InitialVariantWeight(p ProductionVariants, vals map[string]cty.Value) {
	vals["initial_variant_weight"] = cty.NumberIntVal(p.InitialVariantWeight)
}

func EncodeSagemakerEndpointConfiguration_ProductionVariants_InstanceType(p ProductionVariants, vals map[string]cty.Value) {
	vals["instance_type"] = cty.StringVal(p.InstanceType)
}

func EncodeSagemakerEndpointConfiguration_Arn(p SagemakerEndpointConfigurationObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}