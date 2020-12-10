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
	r, ok := mr.(*AccessanalyzerAnalyzer)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a AccessanalyzerAnalyzer.")
	}
	return EncodeAccessanalyzerAnalyzer(*r), nil
}

func EncodeAccessanalyzerAnalyzer(r AccessanalyzerAnalyzer) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeAccessanalyzerAnalyzer_Type(r.Spec.ForProvider, ctyVal)
	EncodeAccessanalyzerAnalyzer_AnalyzerName(r.Spec.ForProvider, ctyVal)
	EncodeAccessanalyzerAnalyzer_Id(r.Spec.ForProvider, ctyVal)
	EncodeAccessanalyzerAnalyzer_Tags(r.Spec.ForProvider, ctyVal)
	EncodeAccessanalyzerAnalyzer_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeAccessanalyzerAnalyzer_Type(p AccessanalyzerAnalyzerParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeAccessanalyzerAnalyzer_AnalyzerName(p AccessanalyzerAnalyzerParameters, vals map[string]cty.Value) {
	vals["analyzer_name"] = cty.StringVal(p.AnalyzerName)
}

func EncodeAccessanalyzerAnalyzer_Id(p AccessanalyzerAnalyzerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAccessanalyzerAnalyzer_Tags(p AccessanalyzerAnalyzerParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeAccessanalyzerAnalyzer_Arn(p AccessanalyzerAnalyzerObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}