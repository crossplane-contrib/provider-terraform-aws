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
	r, ok := mr.(*CloudwatchLogGroup)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a CloudwatchLogGroup.")
	}
	return EncodeCloudwatchLogGroup(*r), nil
}

func EncodeCloudwatchLogGroup(r CloudwatchLogGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCloudwatchLogGroup_KmsKeyId(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogGroup_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogGroup_RetentionInDays(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogGroup_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeCloudwatchLogGroup_KmsKeyId(p CloudwatchLogGroupParameters, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeCloudwatchLogGroup_Name(p CloudwatchLogGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCloudwatchLogGroup_NamePrefix(p CloudwatchLogGroupParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeCloudwatchLogGroup_RetentionInDays(p CloudwatchLogGroupParameters, vals map[string]cty.Value) {
	vals["retention_in_days"] = cty.NumberIntVal(p.RetentionInDays)
}

func EncodeCloudwatchLogGroup_Tags(p CloudwatchLogGroupParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeCloudwatchLogGroup_Arn(p CloudwatchLogGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}