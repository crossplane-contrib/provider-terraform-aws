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
	r, ok := mr.(*LambdaProvisionedConcurrencyConfig)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a LambdaProvisionedConcurrencyConfig.")
	}
	return EncodeLambdaProvisionedConcurrencyConfig(*r), nil
}

func EncodeLambdaProvisionedConcurrencyConfig(r LambdaProvisionedConcurrencyConfig) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeLambdaProvisionedConcurrencyConfig_FunctionName(r.Spec.ForProvider, ctyVal)
	EncodeLambdaProvisionedConcurrencyConfig_Id(r.Spec.ForProvider, ctyVal)
	EncodeLambdaProvisionedConcurrencyConfig_ProvisionedConcurrentExecutions(r.Spec.ForProvider, ctyVal)
	EncodeLambdaProvisionedConcurrencyConfig_Qualifier(r.Spec.ForProvider, ctyVal)
	EncodeLambdaProvisionedConcurrencyConfig_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeLambdaProvisionedConcurrencyConfig_FunctionName(p LambdaProvisionedConcurrencyConfigParameters, vals map[string]cty.Value) {
	vals["function_name"] = cty.StringVal(p.FunctionName)
}

func EncodeLambdaProvisionedConcurrencyConfig_Id(p LambdaProvisionedConcurrencyConfigParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLambdaProvisionedConcurrencyConfig_ProvisionedConcurrentExecutions(p LambdaProvisionedConcurrencyConfigParameters, vals map[string]cty.Value) {
	vals["provisioned_concurrent_executions"] = cty.NumberIntVal(p.ProvisionedConcurrentExecutions)
}

func EncodeLambdaProvisionedConcurrencyConfig_Qualifier(p LambdaProvisionedConcurrencyConfigParameters, vals map[string]cty.Value) {
	vals["qualifier"] = cty.StringVal(p.Qualifier)
}

func EncodeLambdaProvisionedConcurrencyConfig_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeLambdaProvisionedConcurrencyConfig_Timeouts_Create(p, ctyVal)
	EncodeLambdaProvisionedConcurrencyConfig_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeLambdaProvisionedConcurrencyConfig_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeLambdaProvisionedConcurrencyConfig_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}