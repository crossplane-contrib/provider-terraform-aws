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
	r, ok := mr.(*RedshiftSnapshotSchedule)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a RedshiftSnapshotSchedule.")
	}
	return EncodeRedshiftSnapshotSchedule(*r), nil
}

func EncodeRedshiftSnapshotSchedule(r RedshiftSnapshotSchedule) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRedshiftSnapshotSchedule_Tags(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSnapshotSchedule_Definitions(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSnapshotSchedule_Description(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSnapshotSchedule_ForceDestroy(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSnapshotSchedule_Id(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSnapshotSchedule_Identifier(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSnapshotSchedule_IdentifierPrefix(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSnapshotSchedule_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeRedshiftSnapshotSchedule_Tags(p RedshiftSnapshotScheduleParameters, vals map[string]cty.Value) {
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

func EncodeRedshiftSnapshotSchedule_Definitions(p RedshiftSnapshotScheduleParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Definitions {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["definitions"] = cty.SetVal(colVals)
}

func EncodeRedshiftSnapshotSchedule_Description(p RedshiftSnapshotScheduleParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeRedshiftSnapshotSchedule_ForceDestroy(p RedshiftSnapshotScheduleParameters, vals map[string]cty.Value) {
	vals["force_destroy"] = cty.BoolVal(p.ForceDestroy)
}

func EncodeRedshiftSnapshotSchedule_Id(p RedshiftSnapshotScheduleParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRedshiftSnapshotSchedule_Identifier(p RedshiftSnapshotScheduleParameters, vals map[string]cty.Value) {
	vals["identifier"] = cty.StringVal(p.Identifier)
}

func EncodeRedshiftSnapshotSchedule_IdentifierPrefix(p RedshiftSnapshotScheduleParameters, vals map[string]cty.Value) {
	vals["identifier_prefix"] = cty.StringVal(p.IdentifierPrefix)
}

func EncodeRedshiftSnapshotSchedule_Arn(p RedshiftSnapshotScheduleObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}