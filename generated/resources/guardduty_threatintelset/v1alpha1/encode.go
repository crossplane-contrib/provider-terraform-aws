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
	r, ok := mr.(*GuarddutyThreatintelset)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a GuarddutyThreatintelset.")
	}
	return EncodeGuarddutyThreatintelset(*r), nil
}

func EncodeGuarddutyThreatintelset(r GuarddutyThreatintelset) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGuarddutyThreatintelset_Location(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyThreatintelset_Name(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyThreatintelset_Tags(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyThreatintelset_Activate(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyThreatintelset_DetectorId(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyThreatintelset_Format(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyThreatintelset_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeGuarddutyThreatintelset_Location(p GuarddutyThreatintelsetParameters, vals map[string]cty.Value) {
	vals["location"] = cty.StringVal(p.Location)
}

func EncodeGuarddutyThreatintelset_Name(p GuarddutyThreatintelsetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGuarddutyThreatintelset_Tags(p GuarddutyThreatintelsetParameters, vals map[string]cty.Value) {
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

func EncodeGuarddutyThreatintelset_Activate(p GuarddutyThreatintelsetParameters, vals map[string]cty.Value) {
	vals["activate"] = cty.BoolVal(p.Activate)
}

func EncodeGuarddutyThreatintelset_DetectorId(p GuarddutyThreatintelsetParameters, vals map[string]cty.Value) {
	vals["detector_id"] = cty.StringVal(p.DetectorId)
}

func EncodeGuarddutyThreatintelset_Format(p GuarddutyThreatintelsetParameters, vals map[string]cty.Value) {
	vals["format"] = cty.StringVal(p.Format)
}

func EncodeGuarddutyThreatintelset_Arn(p GuarddutyThreatintelsetObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}