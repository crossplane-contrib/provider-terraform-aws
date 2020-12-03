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

package v1alpha1func EncodeLexSlotType(r LexSlotType) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeLexSlotType_Id(r.Spec.ForProvider, ctyVal)
	EncodeLexSlotType_CreateVersion(r.Spec.ForProvider, ctyVal)
	EncodeLexSlotType_Description(r.Spec.ForProvider, ctyVal)
	EncodeLexSlotType_Name(r.Spec.ForProvider, ctyVal)
	EncodeLexSlotType_ValueSelectionStrategy(r.Spec.ForProvider, ctyVal)
	EncodeLexSlotType_EnumerationValue(r.Spec.ForProvider.EnumerationValue, ctyVal)
	EncodeLexSlotType_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeLexSlotType_CreatedDate(r.Status.AtProvider, ctyVal)
	EncodeLexSlotType_Version(r.Status.AtProvider, ctyVal)
	EncodeLexSlotType_Checksum(r.Status.AtProvider, ctyVal)
	EncodeLexSlotType_LastUpdatedDate(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeLexSlotType_Id(p *LexSlotTypeParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLexSlotType_CreateVersion(p *LexSlotTypeParameters, vals map[string]cty.Value) {
	vals["create_version"] = cty.BoolVal(p.CreateVersion)
}

func EncodeLexSlotType_Description(p *LexSlotTypeParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeLexSlotType_Name(p *LexSlotTypeParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeLexSlotType_ValueSelectionStrategy(p *LexSlotTypeParameters, vals map[string]cty.Value) {
	vals["value_selection_strategy"] = cty.StringVal(p.ValueSelectionStrategy)
}

func EncodeLexSlotType_EnumerationValue(p *EnumerationValue, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.EnumerationValue {
		ctyVal = make(map[string]cty.Value)
		EncodeLexSlotType_EnumerationValue_Value(v, ctyVal)
		EncodeLexSlotType_EnumerationValue_Synonyms(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["enumeration_value"] = cty.SetVal(valsForCollection)
}

func EncodeLexSlotType_EnumerationValue_Value(p *EnumerationValue, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeLexSlotType_EnumerationValue_Synonyms(p *EnumerationValue, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Synonyms {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["synonyms"] = cty.SetVal(colVals)
}

func EncodeLexSlotType_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	ctyVal = make(map[string]cty.Value)
	EncodeLexSlotType_Timeouts_Create(p, ctyVal)
	EncodeLexSlotType_Timeouts_Delete(p, ctyVal)
	EncodeLexSlotType_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeLexSlotType_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeLexSlotType_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeLexSlotType_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeLexSlotType_CreatedDate(p *LexSlotTypeObservation, vals map[string]cty.Value) {
	vals["created_date"] = cty.StringVal(p.CreatedDate)
}

func EncodeLexSlotType_Version(p *LexSlotTypeObservation, vals map[string]cty.Value) {
	vals["version"] = cty.StringVal(p.Version)
}

func EncodeLexSlotType_Checksum(p *LexSlotTypeObservation, vals map[string]cty.Value) {
	vals["checksum"] = cty.StringVal(p.Checksum)
}

func EncodeLexSlotType_LastUpdatedDate(p *LexSlotTypeObservation, vals map[string]cty.Value) {
	vals["last_updated_date"] = cty.StringVal(p.LastUpdatedDate)
}