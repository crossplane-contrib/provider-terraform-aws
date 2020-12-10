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
	r, ok := mr.(*BackupSelection)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a BackupSelection.")
	}
	return EncodeBackupSelection(*r), nil
}

func EncodeBackupSelection(r BackupSelection) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeBackupSelection_IamRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeBackupSelection_Id(r.Spec.ForProvider, ctyVal)
	EncodeBackupSelection_Name(r.Spec.ForProvider, ctyVal)
	EncodeBackupSelection_PlanId(r.Spec.ForProvider, ctyVal)
	EncodeBackupSelection_Resources(r.Spec.ForProvider, ctyVal)
	EncodeBackupSelection_SelectionTag(r.Spec.ForProvider.SelectionTag, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeBackupSelection_IamRoleArn(p BackupSelectionParameters, vals map[string]cty.Value) {
	vals["iam_role_arn"] = cty.StringVal(p.IamRoleArn)
}

func EncodeBackupSelection_Id(p BackupSelectionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeBackupSelection_Name(p BackupSelectionParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeBackupSelection_PlanId(p BackupSelectionParameters, vals map[string]cty.Value) {
	vals["plan_id"] = cty.StringVal(p.PlanId)
}

func EncodeBackupSelection_Resources(p BackupSelectionParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Resources {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["resources"] = cty.SetVal(colVals)
}

func EncodeBackupSelection_SelectionTag(p SelectionTag, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeBackupSelection_SelectionTag_Key(p, ctyVal)
	EncodeBackupSelection_SelectionTag_Type(p, ctyVal)
	EncodeBackupSelection_SelectionTag_Value(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["selection_tag"] = cty.SetVal(valsForCollection)
}

func EncodeBackupSelection_SelectionTag_Key(p SelectionTag, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}

func EncodeBackupSelection_SelectionTag_Type(p SelectionTag, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeBackupSelection_SelectionTag_Value(p SelectionTag, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}