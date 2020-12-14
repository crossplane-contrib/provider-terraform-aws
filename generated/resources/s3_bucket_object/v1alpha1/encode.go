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
	r, ok := mr.(*S3BucketObject)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a S3BucketObject.")
	}
	return EncodeS3BucketObject(*r), nil
}

func EncodeS3BucketObject(r S3BucketObject) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeS3BucketObject_Key(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketObject_ObjectLockRetainUntilDate(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketObject_ServerSideEncryption(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketObject_Source(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketObject_StorageClass(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketObject_Tags(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketObject_Content(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketObject_Etag(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketObject_Metadata(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketObject_ObjectLockMode(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketObject_WebsiteRedirect(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketObject_ContentEncoding(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketObject_ForceDestroy(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketObject_ContentDisposition(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketObject_ContentLanguage(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketObject_ContentType(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketObject_Id(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketObject_ObjectLockLegalHoldStatus(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketObject_Acl(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketObject_Bucket(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketObject_KmsKeyId(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketObject_CacheControl(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketObject_ContentBase64(r.Spec.ForProvider, ctyVal)
	EncodeS3BucketObject_VersionId(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeS3BucketObject_Key(p S3BucketObjectParameters, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}

func EncodeS3BucketObject_ObjectLockRetainUntilDate(p S3BucketObjectParameters, vals map[string]cty.Value) {
	vals["object_lock_retain_until_date"] = cty.StringVal(p.ObjectLockRetainUntilDate)
}

func EncodeS3BucketObject_ServerSideEncryption(p S3BucketObjectParameters, vals map[string]cty.Value) {
	vals["server_side_encryption"] = cty.StringVal(p.ServerSideEncryption)
}

func EncodeS3BucketObject_Source(p S3BucketObjectParameters, vals map[string]cty.Value) {
	vals["source"] = cty.StringVal(p.Source)
}

func EncodeS3BucketObject_StorageClass(p S3BucketObjectParameters, vals map[string]cty.Value) {
	vals["storage_class"] = cty.StringVal(p.StorageClass)
}

func EncodeS3BucketObject_Tags(p S3BucketObjectParameters, vals map[string]cty.Value) {
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

func EncodeS3BucketObject_Content(p S3BucketObjectParameters, vals map[string]cty.Value) {
	vals["content"] = cty.StringVal(p.Content)
}

func EncodeS3BucketObject_Etag(p S3BucketObjectParameters, vals map[string]cty.Value) {
	vals["etag"] = cty.StringVal(p.Etag)
}

func EncodeS3BucketObject_Metadata(p S3BucketObjectParameters, vals map[string]cty.Value) {
	if len(p.Metadata) == 0 {
		vals["metadata"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Metadata {
		mVals[key] = cty.StringVal(value)
	}
	vals["metadata"] = cty.MapVal(mVals)
}

func EncodeS3BucketObject_ObjectLockMode(p S3BucketObjectParameters, vals map[string]cty.Value) {
	vals["object_lock_mode"] = cty.StringVal(p.ObjectLockMode)
}

func EncodeS3BucketObject_WebsiteRedirect(p S3BucketObjectParameters, vals map[string]cty.Value) {
	vals["website_redirect"] = cty.StringVal(p.WebsiteRedirect)
}

func EncodeS3BucketObject_ContentEncoding(p S3BucketObjectParameters, vals map[string]cty.Value) {
	vals["content_encoding"] = cty.StringVal(p.ContentEncoding)
}

func EncodeS3BucketObject_ForceDestroy(p S3BucketObjectParameters, vals map[string]cty.Value) {
	vals["force_destroy"] = cty.BoolVal(p.ForceDestroy)
}

func EncodeS3BucketObject_ContentDisposition(p S3BucketObjectParameters, vals map[string]cty.Value) {
	vals["content_disposition"] = cty.StringVal(p.ContentDisposition)
}

func EncodeS3BucketObject_ContentLanguage(p S3BucketObjectParameters, vals map[string]cty.Value) {
	vals["content_language"] = cty.StringVal(p.ContentLanguage)
}

func EncodeS3BucketObject_ContentType(p S3BucketObjectParameters, vals map[string]cty.Value) {
	vals["content_type"] = cty.StringVal(p.ContentType)
}

func EncodeS3BucketObject_Id(p S3BucketObjectParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeS3BucketObject_ObjectLockLegalHoldStatus(p S3BucketObjectParameters, vals map[string]cty.Value) {
	vals["object_lock_legal_hold_status"] = cty.StringVal(p.ObjectLockLegalHoldStatus)
}

func EncodeS3BucketObject_Acl(p S3BucketObjectParameters, vals map[string]cty.Value) {
	vals["acl"] = cty.StringVal(p.Acl)
}

func EncodeS3BucketObject_Bucket(p S3BucketObjectParameters, vals map[string]cty.Value) {
	vals["bucket"] = cty.StringVal(p.Bucket)
}

func EncodeS3BucketObject_KmsKeyId(p S3BucketObjectParameters, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeS3BucketObject_CacheControl(p S3BucketObjectParameters, vals map[string]cty.Value) {
	vals["cache_control"] = cty.StringVal(p.CacheControl)
}

func EncodeS3BucketObject_ContentBase64(p S3BucketObjectParameters, vals map[string]cty.Value) {
	vals["content_base64"] = cty.StringVal(p.ContentBase64)
}

func EncodeS3BucketObject_VersionId(p S3BucketObjectObservation, vals map[string]cty.Value) {
	vals["version_id"] = cty.StringVal(p.VersionId)
}