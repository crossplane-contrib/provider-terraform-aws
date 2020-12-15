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
	r, ok := mr.(*RedshiftSnapshotSchedule)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeRedshiftSnapshotSchedule(r, ctyValue)
}

func DecodeRedshiftSnapshotSchedule(prev *RedshiftSnapshotSchedule, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeRedshiftSnapshotSchedule_ForceDestroy(&new.Spec.ForProvider, valMap)
	DecodeRedshiftSnapshotSchedule_Identifier(&new.Spec.ForProvider, valMap)
	DecodeRedshiftSnapshotSchedule_IdentifierPrefix(&new.Spec.ForProvider, valMap)
	DecodeRedshiftSnapshotSchedule_Tags(&new.Spec.ForProvider, valMap)
	DecodeRedshiftSnapshotSchedule_Definitions(&new.Spec.ForProvider, valMap)
	DecodeRedshiftSnapshotSchedule_Description(&new.Spec.ForProvider, valMap)
	DecodeRedshiftSnapshotSchedule_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeRedshiftSnapshotSchedule_ForceDestroy(p *RedshiftSnapshotScheduleParameters, vals map[string]cty.Value) {
	p.ForceDestroy = ctwhy.ValueAsBool(vals["force_destroy"])
}

//primitiveTypeDecodeTemplate
func DecodeRedshiftSnapshotSchedule_Identifier(p *RedshiftSnapshotScheduleParameters, vals map[string]cty.Value) {
	p.Identifier = ctwhy.ValueAsString(vals["identifier"])
}

//primitiveTypeDecodeTemplate
func DecodeRedshiftSnapshotSchedule_IdentifierPrefix(p *RedshiftSnapshotScheduleParameters, vals map[string]cty.Value) {
	p.IdentifierPrefix = ctwhy.ValueAsString(vals["identifier_prefix"])
}

//primitiveMapTypeDecodeTemplate
func DecodeRedshiftSnapshotSchedule_Tags(p *RedshiftSnapshotScheduleParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveCollectionTypeDecodeTemplate
func DecodeRedshiftSnapshotSchedule_Definitions(p *RedshiftSnapshotScheduleParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["definitions"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.Definitions = goVals
}

//primitiveTypeDecodeTemplate
func DecodeRedshiftSnapshotSchedule_Description(p *RedshiftSnapshotScheduleParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeRedshiftSnapshotSchedule_Arn(p *RedshiftSnapshotScheduleObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}