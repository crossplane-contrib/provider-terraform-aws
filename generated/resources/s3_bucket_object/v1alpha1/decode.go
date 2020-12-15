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
	r, ok := mr.(*S3BucketObject)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeS3BucketObject(r, ctyValue)
}

func DecodeS3BucketObject(prev *S3BucketObject, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeS3BucketObject_Bucket(&new.Spec.ForProvider, valMap)
	DecodeS3BucketObject_ContentBase64(&new.Spec.ForProvider, valMap)
	DecodeS3BucketObject_KmsKeyId(&new.Spec.ForProvider, valMap)
	DecodeS3BucketObject_Metadata(&new.Spec.ForProvider, valMap)
	DecodeS3BucketObject_ObjectLockLegalHoldStatus(&new.Spec.ForProvider, valMap)
	DecodeS3BucketObject_ServerSideEncryption(&new.Spec.ForProvider, valMap)
	DecodeS3BucketObject_CacheControl(&new.Spec.ForProvider, valMap)
	DecodeS3BucketObject_ContentDisposition(&new.Spec.ForProvider, valMap)
	DecodeS3BucketObject_Key(&new.Spec.ForProvider, valMap)
	DecodeS3BucketObject_Source(&new.Spec.ForProvider, valMap)
	DecodeS3BucketObject_Acl(&new.Spec.ForProvider, valMap)
	DecodeS3BucketObject_ContentEncoding(&new.Spec.ForProvider, valMap)
	DecodeS3BucketObject_Etag(&new.Spec.ForProvider, valMap)
	DecodeS3BucketObject_ObjectLockRetainUntilDate(&new.Spec.ForProvider, valMap)
	DecodeS3BucketObject_StorageClass(&new.Spec.ForProvider, valMap)
	DecodeS3BucketObject_Tags(&new.Spec.ForProvider, valMap)
	DecodeS3BucketObject_WebsiteRedirect(&new.Spec.ForProvider, valMap)
	DecodeS3BucketObject_Content(&new.Spec.ForProvider, valMap)
	DecodeS3BucketObject_ContentLanguage(&new.Spec.ForProvider, valMap)
	DecodeS3BucketObject_ContentType(&new.Spec.ForProvider, valMap)
	DecodeS3BucketObject_ForceDestroy(&new.Spec.ForProvider, valMap)
	DecodeS3BucketObject_Id(&new.Spec.ForProvider, valMap)
	DecodeS3BucketObject_ObjectLockMode(&new.Spec.ForProvider, valMap)
	DecodeS3BucketObject_VersionId(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketObject_Bucket(p *S3BucketObjectParameters, vals map[string]cty.Value) {
	p.Bucket = ctwhy.ValueAsString(vals["bucket"])
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketObject_ContentBase64(p *S3BucketObjectParameters, vals map[string]cty.Value) {
	p.ContentBase64 = ctwhy.ValueAsString(vals["content_base64"])
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketObject_KmsKeyId(p *S3BucketObjectParameters, vals map[string]cty.Value) {
	p.KmsKeyId = ctwhy.ValueAsString(vals["kms_key_id"])
}

//primitiveMapTypeDecodeTemplate
func DecodeS3BucketObject_Metadata(p *S3BucketObjectParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["metadata"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Metadata = vMap
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketObject_ObjectLockLegalHoldStatus(p *S3BucketObjectParameters, vals map[string]cty.Value) {
	p.ObjectLockLegalHoldStatus = ctwhy.ValueAsString(vals["object_lock_legal_hold_status"])
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketObject_ServerSideEncryption(p *S3BucketObjectParameters, vals map[string]cty.Value) {
	p.ServerSideEncryption = ctwhy.ValueAsString(vals["server_side_encryption"])
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketObject_CacheControl(p *S3BucketObjectParameters, vals map[string]cty.Value) {
	p.CacheControl = ctwhy.ValueAsString(vals["cache_control"])
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketObject_ContentDisposition(p *S3BucketObjectParameters, vals map[string]cty.Value) {
	p.ContentDisposition = ctwhy.ValueAsString(vals["content_disposition"])
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketObject_Key(p *S3BucketObjectParameters, vals map[string]cty.Value) {
	p.Key = ctwhy.ValueAsString(vals["key"])
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketObject_Source(p *S3BucketObjectParameters, vals map[string]cty.Value) {
	p.Source = ctwhy.ValueAsString(vals["source"])
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketObject_Acl(p *S3BucketObjectParameters, vals map[string]cty.Value) {
	p.Acl = ctwhy.ValueAsString(vals["acl"])
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketObject_ContentEncoding(p *S3BucketObjectParameters, vals map[string]cty.Value) {
	p.ContentEncoding = ctwhy.ValueAsString(vals["content_encoding"])
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketObject_Etag(p *S3BucketObjectParameters, vals map[string]cty.Value) {
	p.Etag = ctwhy.ValueAsString(vals["etag"])
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketObject_ObjectLockRetainUntilDate(p *S3BucketObjectParameters, vals map[string]cty.Value) {
	p.ObjectLockRetainUntilDate = ctwhy.ValueAsString(vals["object_lock_retain_until_date"])
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketObject_StorageClass(p *S3BucketObjectParameters, vals map[string]cty.Value) {
	p.StorageClass = ctwhy.ValueAsString(vals["storage_class"])
}

//primitiveMapTypeDecodeTemplate
func DecodeS3BucketObject_Tags(p *S3BucketObjectParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketObject_WebsiteRedirect(p *S3BucketObjectParameters, vals map[string]cty.Value) {
	p.WebsiteRedirect = ctwhy.ValueAsString(vals["website_redirect"])
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketObject_Content(p *S3BucketObjectParameters, vals map[string]cty.Value) {
	p.Content = ctwhy.ValueAsString(vals["content"])
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketObject_ContentLanguage(p *S3BucketObjectParameters, vals map[string]cty.Value) {
	p.ContentLanguage = ctwhy.ValueAsString(vals["content_language"])
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketObject_ContentType(p *S3BucketObjectParameters, vals map[string]cty.Value) {
	p.ContentType = ctwhy.ValueAsString(vals["content_type"])
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketObject_ForceDestroy(p *S3BucketObjectParameters, vals map[string]cty.Value) {
	p.ForceDestroy = ctwhy.ValueAsBool(vals["force_destroy"])
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketObject_Id(p *S3BucketObjectParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketObject_ObjectLockMode(p *S3BucketObjectParameters, vals map[string]cty.Value) {
	p.ObjectLockMode = ctwhy.ValueAsString(vals["object_lock_mode"])
}

//primitiveTypeDecodeTemplate
func DecodeS3BucketObject_VersionId(p *S3BucketObjectObservation, vals map[string]cty.Value) {
	p.VersionId = ctwhy.ValueAsString(vals["version_id"])
}