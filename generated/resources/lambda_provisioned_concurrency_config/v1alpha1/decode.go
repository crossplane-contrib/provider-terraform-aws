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
	r, ok := mr.(*LambdaProvisionedConcurrencyConfig)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeLambdaProvisionedConcurrencyConfig(r, ctyValue)
}

func DecodeLambdaProvisionedConcurrencyConfig(prev *LambdaProvisionedConcurrencyConfig, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeLambdaProvisionedConcurrencyConfig_FunctionName(&new.Spec.ForProvider, valMap)
	DecodeLambdaProvisionedConcurrencyConfig_ProvisionedConcurrentExecutions(&new.Spec.ForProvider, valMap)
	DecodeLambdaProvisionedConcurrencyConfig_Qualifier(&new.Spec.ForProvider, valMap)
	DecodeLambdaProvisionedConcurrencyConfig_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeLambdaProvisionedConcurrencyConfig_FunctionName(p *LambdaProvisionedConcurrencyConfigParameters, vals map[string]cty.Value) {
	p.FunctionName = ctwhy.ValueAsString(vals["function_name"])
}

//primitiveTypeDecodeTemplate
func DecodeLambdaProvisionedConcurrencyConfig_ProvisionedConcurrentExecutions(p *LambdaProvisionedConcurrencyConfigParameters, vals map[string]cty.Value) {
	p.ProvisionedConcurrentExecutions = ctwhy.ValueAsInt64(vals["provisioned_concurrent_executions"])
}

//primitiveTypeDecodeTemplate
func DecodeLambdaProvisionedConcurrencyConfig_Qualifier(p *LambdaProvisionedConcurrencyConfigParameters, vals map[string]cty.Value) {
	p.Qualifier = ctwhy.ValueAsString(vals["qualifier"])
}

//containerTypeDecodeTemplate
func DecodeLambdaProvisionedConcurrencyConfig_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeLambdaProvisionedConcurrencyConfig_Timeouts_Create(p, valMap)
	DecodeLambdaProvisionedConcurrencyConfig_Timeouts_Update(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeLambdaProvisionedConcurrencyConfig_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveTypeDecodeTemplate
func DecodeLambdaProvisionedConcurrencyConfig_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}