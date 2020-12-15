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
	r, ok := mr.(*RedshiftSnapshotCopyGrant)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeRedshiftSnapshotCopyGrant(r, ctyValue)
}

func DecodeRedshiftSnapshotCopyGrant(prev *RedshiftSnapshotCopyGrant, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeRedshiftSnapshotCopyGrant_Id(&new.Spec.ForProvider, valMap)
	DecodeRedshiftSnapshotCopyGrant_KmsKeyId(&new.Spec.ForProvider, valMap)
	DecodeRedshiftSnapshotCopyGrant_SnapshotCopyGrantName(&new.Spec.ForProvider, valMap)
	DecodeRedshiftSnapshotCopyGrant_Tags(&new.Spec.ForProvider, valMap)
	DecodeRedshiftSnapshotCopyGrant_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeRedshiftSnapshotCopyGrant_Id(p *RedshiftSnapshotCopyGrantParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeRedshiftSnapshotCopyGrant_KmsKeyId(p *RedshiftSnapshotCopyGrantParameters, vals map[string]cty.Value) {
	p.KmsKeyId = ctwhy.ValueAsString(vals["kms_key_id"])
}

//primitiveTypeDecodeTemplate
func DecodeRedshiftSnapshotCopyGrant_SnapshotCopyGrantName(p *RedshiftSnapshotCopyGrantParameters, vals map[string]cty.Value) {
	p.SnapshotCopyGrantName = ctwhy.ValueAsString(vals["snapshot_copy_grant_name"])
}

//primitiveMapTypeDecodeTemplate
func DecodeRedshiftSnapshotCopyGrant_Tags(p *RedshiftSnapshotCopyGrantParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeRedshiftSnapshotCopyGrant_Arn(p *RedshiftSnapshotCopyGrantObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}