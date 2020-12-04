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
	"github.com/zclconf/go-cty/cty"
)

func EncodeDatapipelinePipeline(r DatapipelinePipeline) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDatapipelinePipeline_Description(r.Spec.ForProvider, ctyVal)
	EncodeDatapipelinePipeline_Id(r.Spec.ForProvider, ctyVal)
	EncodeDatapipelinePipeline_Name(r.Spec.ForProvider, ctyVal)
	EncodeDatapipelinePipeline_Tags(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeDatapipelinePipeline_Description(p DatapipelinePipelineParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeDatapipelinePipeline_Id(p DatapipelinePipelineParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDatapipelinePipeline_Name(p DatapipelinePipelineParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDatapipelinePipeline_Tags(p DatapipelinePipelineParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}