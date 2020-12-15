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
	r, ok := mr.(*RedshiftSnapshotCopyGrant)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a RedshiftSnapshotCopyGrant.")
	}
	return EncodeRedshiftSnapshotCopyGrant(*r), nil
}

func EncodeRedshiftSnapshotCopyGrant(r RedshiftSnapshotCopyGrant) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRedshiftSnapshotCopyGrant_Id(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSnapshotCopyGrant_KmsKeyId(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSnapshotCopyGrant_SnapshotCopyGrantName(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSnapshotCopyGrant_Tags(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftSnapshotCopyGrant_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeRedshiftSnapshotCopyGrant_Id(p RedshiftSnapshotCopyGrantParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRedshiftSnapshotCopyGrant_KmsKeyId(p RedshiftSnapshotCopyGrantParameters, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeRedshiftSnapshotCopyGrant_SnapshotCopyGrantName(p RedshiftSnapshotCopyGrantParameters, vals map[string]cty.Value) {
	vals["snapshot_copy_grant_name"] = cty.StringVal(p.SnapshotCopyGrantName)
}

func EncodeRedshiftSnapshotCopyGrant_Tags(p RedshiftSnapshotCopyGrantParameters, vals map[string]cty.Value) {
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

func EncodeRedshiftSnapshotCopyGrant_Arn(p RedshiftSnapshotCopyGrantObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}