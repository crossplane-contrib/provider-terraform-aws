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
	r, ok := mr.(*GuarddutyDetector)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a GuarddutyDetector.")
	}
	return EncodeGuarddutyDetector(*r), nil
}

func EncodeGuarddutyDetector(r GuarddutyDetector) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGuarddutyDetector_Enable(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyDetector_FindingPublishingFrequency(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyDetector_Tags(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyDetector_AccountId(r.Status.AtProvider, ctyVal)
	EncodeGuarddutyDetector_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeGuarddutyDetector_Enable(p GuarddutyDetectorParameters, vals map[string]cty.Value) {
	vals["enable"] = cty.BoolVal(p.Enable)
}

func EncodeGuarddutyDetector_FindingPublishingFrequency(p GuarddutyDetectorParameters, vals map[string]cty.Value) {
	vals["finding_publishing_frequency"] = cty.StringVal(p.FindingPublishingFrequency)
}

func EncodeGuarddutyDetector_Tags(p GuarddutyDetectorParameters, vals map[string]cty.Value) {
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

func EncodeGuarddutyDetector_AccountId(p GuarddutyDetectorObservation, vals map[string]cty.Value) {
	vals["account_id"] = cty.StringVal(p.AccountId)
}

func EncodeGuarddutyDetector_Arn(p GuarddutyDetectorObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}