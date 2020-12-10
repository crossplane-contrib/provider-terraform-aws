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
	r, ok := mr.(*SsmPatchBaseline)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a SsmPatchBaseline.")
	}
	return EncodeSsmPatchBaseline(*r), nil
}

func EncodeSsmPatchBaseline(r SsmPatchBaseline) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSsmPatchBaseline_OperatingSystem(r.Spec.ForProvider, ctyVal)
	EncodeSsmPatchBaseline_RejectedPatches(r.Spec.ForProvider, ctyVal)
	EncodeSsmPatchBaseline_Tags(r.Spec.ForProvider, ctyVal)
	EncodeSsmPatchBaseline_ApprovedPatches(r.Spec.ForProvider, ctyVal)
	EncodeSsmPatchBaseline_ApprovedPatchesComplianceLevel(r.Spec.ForProvider, ctyVal)
	EncodeSsmPatchBaseline_Description(r.Spec.ForProvider, ctyVal)
	EncodeSsmPatchBaseline_Id(r.Spec.ForProvider, ctyVal)
	EncodeSsmPatchBaseline_Name(r.Spec.ForProvider, ctyVal)
	EncodeSsmPatchBaseline_ApprovalRule(r.Spec.ForProvider.ApprovalRule, ctyVal)
	EncodeSsmPatchBaseline_GlobalFilter(r.Spec.ForProvider.GlobalFilter, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeSsmPatchBaseline_OperatingSystem(p SsmPatchBaselineParameters, vals map[string]cty.Value) {
	vals["operating_system"] = cty.StringVal(p.OperatingSystem)
}

func EncodeSsmPatchBaseline_RejectedPatches(p SsmPatchBaselineParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.RejectedPatches {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["rejected_patches"] = cty.SetVal(colVals)
}

func EncodeSsmPatchBaseline_Tags(p SsmPatchBaselineParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeSsmPatchBaseline_ApprovedPatches(p SsmPatchBaselineParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ApprovedPatches {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["approved_patches"] = cty.SetVal(colVals)
}

func EncodeSsmPatchBaseline_ApprovedPatchesComplianceLevel(p SsmPatchBaselineParameters, vals map[string]cty.Value) {
	vals["approved_patches_compliance_level"] = cty.StringVal(p.ApprovedPatchesComplianceLevel)
}

func EncodeSsmPatchBaseline_Description(p SsmPatchBaselineParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeSsmPatchBaseline_Id(p SsmPatchBaselineParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSsmPatchBaseline_Name(p SsmPatchBaselineParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSsmPatchBaseline_ApprovalRule(p ApprovalRule, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeSsmPatchBaseline_ApprovalRule_ApproveAfterDays(p, ctyVal)
	EncodeSsmPatchBaseline_ApprovalRule_ComplianceLevel(p, ctyVal)
	EncodeSsmPatchBaseline_ApprovalRule_EnableNonSecurity(p, ctyVal)
	EncodeSsmPatchBaseline_ApprovalRule_PatchFilter(p.PatchFilter, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["approval_rule"] = cty.ListVal(valsForCollection)
}

func EncodeSsmPatchBaseline_ApprovalRule_ApproveAfterDays(p ApprovalRule, vals map[string]cty.Value) {
	vals["approve_after_days"] = cty.NumberIntVal(p.ApproveAfterDays)
}

func EncodeSsmPatchBaseline_ApprovalRule_ComplianceLevel(p ApprovalRule, vals map[string]cty.Value) {
	vals["compliance_level"] = cty.StringVal(p.ComplianceLevel)
}

func EncodeSsmPatchBaseline_ApprovalRule_EnableNonSecurity(p ApprovalRule, vals map[string]cty.Value) {
	vals["enable_non_security"] = cty.BoolVal(p.EnableNonSecurity)
}

func EncodeSsmPatchBaseline_ApprovalRule_PatchFilter(p []PatchFilter, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeSsmPatchBaseline_ApprovalRule_PatchFilter_Key(v, ctyVal)
		EncodeSsmPatchBaseline_ApprovalRule_PatchFilter_Values(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["patch_filter"] = cty.ListVal(valsForCollection)
}

func EncodeSsmPatchBaseline_ApprovalRule_PatchFilter_Key(p PatchFilter, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}

func EncodeSsmPatchBaseline_ApprovalRule_PatchFilter_Values(p PatchFilter, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Values {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["values"] = cty.ListVal(colVals)
}

func EncodeSsmPatchBaseline_GlobalFilter(p []GlobalFilter, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeSsmPatchBaseline_GlobalFilter_Key(v, ctyVal)
		EncodeSsmPatchBaseline_GlobalFilter_Values(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["global_filter"] = cty.ListVal(valsForCollection)
}

func EncodeSsmPatchBaseline_GlobalFilter_Key(p GlobalFilter, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}

func EncodeSsmPatchBaseline_GlobalFilter_Values(p GlobalFilter, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Values {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["values"] = cty.ListVal(colVals)
}