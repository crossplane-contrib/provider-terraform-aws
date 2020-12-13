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
	r, ok := mr.(*S3BucketInventory)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a S3BucketInventory.")
	}
	return EncodeS3BucketInventory(*r), nil
}

func EncodeS3BucketInventory(r S3BucketInventory) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeS3BucketInventory_Bucket(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketInventory_Enabled(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketInventory_Id(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketInventory_IncludedObjectVersions(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketInventory_Name(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketInventory_OptionalFields(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketInventory_Destination(r.Spec.ForProvider.Destination, ctyVal)
	EncodeS3BucketInventory_Filter(r.Spec.ForProvider.Filter, ctyVal)
	EncodeS3BucketInventory_Schedule(r.Spec.ForProvider.Schedule, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeS3BucketInventory_Bucket(p S3BucketInventoryParameters, vals map[string]cty.Value) {
	vals["bucket"] = cty.StringVal(p.Bucket)
}

func EncodeS3BucketInventory_Enabled(p S3BucketInventoryParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeS3BucketInventory_Id(p S3BucketInventoryParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeS3BucketInventory_IncludedObjectVersions(p S3BucketInventoryParameters, vals map[string]cty.Value) {
	vals["included_object_versions"] = cty.StringVal(p.IncludedObjectVersions)
}

func EncodeS3BucketInventory_Name(p S3BucketInventoryParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeS3BucketInventory_OptionalFields(p S3BucketInventoryParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.OptionalFields {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["optional_fields"] = cty.SetVal(colVals)
}

func EncodeS3BucketInventory_Destination(p Destination, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeS3BucketInventory_Destination_Bucket(p.Bucket, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["destination"] = cty.ListVal(valsForCollection)
}

func EncodeS3BucketInventory_Destination_Bucket(p Bucket, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeS3BucketInventory_Destination_Bucket_AccountId(p, ctyVal)
	EncodeS3BucketInventory_Destination_Bucket_BucketArn(p, ctyVal)
	EncodeS3BucketInventory_Destination_Bucket_Format(p, ctyVal)
	EncodeS3BucketInventory_Destination_Bucket_Prefix(p, ctyVal)
	EncodeS3BucketInventory_Destination_Bucket_Encryption(p.Encryption, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["bucket"] = cty.ListVal(valsForCollection)
}

func EncodeS3BucketInventory_Destination_Bucket_AccountId(p Bucket, vals map[string]cty.Value) {
	vals["account_id"] = cty.StringVal(p.AccountId)
}

func EncodeS3BucketInventory_Destination_Bucket_BucketArn(p Bucket, vals map[string]cty.Value) {
	vals["bucket_arn"] = cty.StringVal(p.BucketArn)
}

func EncodeS3BucketInventory_Destination_Bucket_Format(p Bucket, vals map[string]cty.Value) {
	vals["format"] = cty.StringVal(p.Format)
}

func EncodeS3BucketInventory_Destination_Bucket_Prefix(p Bucket, vals map[string]cty.Value) {
	vals["prefix"] = cty.StringVal(p.Prefix)
}

func EncodeS3BucketInventory_Destination_Bucket_Encryption(p Encryption, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeS3BucketInventory_Destination_Bucket_Encryption_SseKms(p.SseKms, ctyVal)
	EncodeS3BucketInventory_Destination_Bucket_Encryption_SseS3(p.SseS3, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["encryption"] = cty.ListVal(valsForCollection)
}

func EncodeS3BucketInventory_Destination_Bucket_Encryption_SseKms(p SseKms, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeS3BucketInventory_Destination_Bucket_Encryption_SseKms_KeyId(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["sse_kms"] = cty.ListVal(valsForCollection)
}

func EncodeS3BucketInventory_Destination_Bucket_Encryption_SseKms_KeyId(p SseKms, vals map[string]cty.Value) {
	vals["key_id"] = cty.StringVal(p.KeyId)
}

func EncodeS3BucketInventory_Destination_Bucket_Encryption_SseS3(p SseS3, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)

	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["sse_s3"] = cty.ListVal(valsForCollection)
}

func EncodeS3BucketInventory_Filter(p Filter, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeS3BucketInventory_Filter_Prefix(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["filter"] = cty.ListVal(valsForCollection)
}

func EncodeS3BucketInventory_Filter_Prefix(p Filter, vals map[string]cty.Value) {
	vals["prefix"] = cty.StringVal(p.Prefix)
}

func EncodeS3BucketInventory_Schedule(p Schedule, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeS3BucketInventory_Schedule_Frequency(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["schedule"] = cty.ListVal(valsForCollection)
}

func EncodeS3BucketInventory_Schedule_Frequency(p Schedule, vals map[string]cty.Value) {
	vals["frequency"] = cty.StringVal(p.Frequency)
}