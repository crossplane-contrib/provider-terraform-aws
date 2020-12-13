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
	r, ok := mr.(*SsmResourceDataSync)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a SsmResourceDataSync.")
	}
	return EncodeSsmResourceDataSync(*r), nil
}

func EncodeSsmResourceDataSync(r SsmResourceDataSync) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSsmResourceDataSync_Id(r.Spec.ForProvider, ctyVal)
	EncodeSsmResourceDataSync_Name(r.Spec.ForProvider, ctyVal)
	EncodeSsmResourceDataSync_S3Destination(r.Spec.ForProvider.S3Destination, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeSsmResourceDataSync_Id(p SsmResourceDataSyncParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSsmResourceDataSync_Name(p SsmResourceDataSyncParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSsmResourceDataSync_S3Destination(p S3Destination, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeSsmResourceDataSync_S3Destination_BucketName(p, ctyVal)
	EncodeSsmResourceDataSync_S3Destination_KmsKeyArn(p, ctyVal)
	EncodeSsmResourceDataSync_S3Destination_Prefix(p, ctyVal)
	EncodeSsmResourceDataSync_S3Destination_Region(p, ctyVal)
	EncodeSsmResourceDataSync_S3Destination_SyncFormat(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["s3_destination"] = cty.ListVal(valsForCollection)
}

func EncodeSsmResourceDataSync_S3Destination_BucketName(p S3Destination, vals map[string]cty.Value) {
	vals["bucket_name"] = cty.StringVal(p.BucketName)
}

func EncodeSsmResourceDataSync_S3Destination_KmsKeyArn(p S3Destination, vals map[string]cty.Value) {
	vals["kms_key_arn"] = cty.StringVal(p.KmsKeyArn)
}

func EncodeSsmResourceDataSync_S3Destination_Prefix(p S3Destination, vals map[string]cty.Value) {
	vals["prefix"] = cty.StringVal(p.Prefix)
}

func EncodeSsmResourceDataSync_S3Destination_Region(p S3Destination, vals map[string]cty.Value) {
	vals["region"] = cty.StringVal(p.Region)
}

func EncodeSsmResourceDataSync_S3Destination_SyncFormat(p S3Destination, vals map[string]cty.Value) {
	vals["sync_format"] = cty.StringVal(p.SyncFormat)
}