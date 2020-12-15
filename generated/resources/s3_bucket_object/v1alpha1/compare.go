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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane-contrib/terraform-runtime/pkg/plugin"
)

//mergeManagedResourceEntrypointTemplate
type resourceMerger struct{}

func (r *resourceMerger) MergeResources(kube resource.Managed, prov resource.Managed) plugin.MergeDescription {
	k := kube.(*S3BucketObject)
	p := prov.(*S3BucketObject)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeS3BucketObject_ContentType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeS3BucketObject_Etag(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeS3BucketObject_Bucket(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeS3BucketObject_CacheControl(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeS3BucketObject_ContentBase64(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeS3BucketObject_ContentDisposition(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeS3BucketObject_StorageClass(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeS3BucketObject_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeS3BucketObject_WebsiteRedirect(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeS3BucketObject_Metadata(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeS3BucketObject_ObjectLockLegalHoldStatus(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeS3BucketObject_ObjectLockMode(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeS3BucketObject_ObjectLockRetainUntilDate(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeS3BucketObject_ServerSideEncryption(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeS3BucketObject_Source(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeS3BucketObject_Acl(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeS3BucketObject_ContentLanguage(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeS3BucketObject_KmsKeyId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeS3BucketObject_Content(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeS3BucketObject_ContentEncoding(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeS3BucketObject_ForceDestroy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeS3BucketObject_Key(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeS3BucketObject_VersionId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	for key, v := range p.Annotations {
		if k.Annotations[key] != v {
			k.Annotations[key] = v
			md.AnnotationsUpdated = true
		}
	}
	md.AnyFieldUpdated = anyChildUpdated
	return *md
}

//mergePrimitiveTemplateSpec
func MergeS3BucketObject_ContentType(k *S3BucketObjectParameters, p *S3BucketObjectParameters, md *plugin.MergeDescription) bool {
	if k.ContentType != p.ContentType {
		p.ContentType = k.ContentType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeS3BucketObject_Etag(k *S3BucketObjectParameters, p *S3BucketObjectParameters, md *plugin.MergeDescription) bool {
	if k.Etag != p.Etag {
		p.Etag = k.Etag
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeS3BucketObject_Bucket(k *S3BucketObjectParameters, p *S3BucketObjectParameters, md *plugin.MergeDescription) bool {
	if k.Bucket != p.Bucket {
		p.Bucket = k.Bucket
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeS3BucketObject_CacheControl(k *S3BucketObjectParameters, p *S3BucketObjectParameters, md *plugin.MergeDescription) bool {
	if k.CacheControl != p.CacheControl {
		p.CacheControl = k.CacheControl
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeS3BucketObject_ContentBase64(k *S3BucketObjectParameters, p *S3BucketObjectParameters, md *plugin.MergeDescription) bool {
	if k.ContentBase64 != p.ContentBase64 {
		p.ContentBase64 = k.ContentBase64
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeS3BucketObject_ContentDisposition(k *S3BucketObjectParameters, p *S3BucketObjectParameters, md *plugin.MergeDescription) bool {
	if k.ContentDisposition != p.ContentDisposition {
		p.ContentDisposition = k.ContentDisposition
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeS3BucketObject_StorageClass(k *S3BucketObjectParameters, p *S3BucketObjectParameters, md *plugin.MergeDescription) bool {
	if k.StorageClass != p.StorageClass {
		p.StorageClass = k.StorageClass
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeS3BucketObject_Tags(k *S3BucketObjectParameters, p *S3BucketObjectParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeS3BucketObject_WebsiteRedirect(k *S3BucketObjectParameters, p *S3BucketObjectParameters, md *plugin.MergeDescription) bool {
	if k.WebsiteRedirect != p.WebsiteRedirect {
		p.WebsiteRedirect = k.WebsiteRedirect
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeS3BucketObject_Metadata(k *S3BucketObjectParameters, p *S3BucketObjectParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Metadata, p.Metadata) {
		p.Metadata = k.Metadata
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeS3BucketObject_ObjectLockLegalHoldStatus(k *S3BucketObjectParameters, p *S3BucketObjectParameters, md *plugin.MergeDescription) bool {
	if k.ObjectLockLegalHoldStatus != p.ObjectLockLegalHoldStatus {
		p.ObjectLockLegalHoldStatus = k.ObjectLockLegalHoldStatus
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeS3BucketObject_ObjectLockMode(k *S3BucketObjectParameters, p *S3BucketObjectParameters, md *plugin.MergeDescription) bool {
	if k.ObjectLockMode != p.ObjectLockMode {
		p.ObjectLockMode = k.ObjectLockMode
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeS3BucketObject_ObjectLockRetainUntilDate(k *S3BucketObjectParameters, p *S3BucketObjectParameters, md *plugin.MergeDescription) bool {
	if k.ObjectLockRetainUntilDate != p.ObjectLockRetainUntilDate {
		p.ObjectLockRetainUntilDate = k.ObjectLockRetainUntilDate
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeS3BucketObject_ServerSideEncryption(k *S3BucketObjectParameters, p *S3BucketObjectParameters, md *plugin.MergeDescription) bool {
	if k.ServerSideEncryption != p.ServerSideEncryption {
		p.ServerSideEncryption = k.ServerSideEncryption
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeS3BucketObject_Source(k *S3BucketObjectParameters, p *S3BucketObjectParameters, md *plugin.MergeDescription) bool {
	if k.Source != p.Source {
		p.Source = k.Source
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeS3BucketObject_Acl(k *S3BucketObjectParameters, p *S3BucketObjectParameters, md *plugin.MergeDescription) bool {
	if k.Acl != p.Acl {
		p.Acl = k.Acl
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeS3BucketObject_ContentLanguage(k *S3BucketObjectParameters, p *S3BucketObjectParameters, md *plugin.MergeDescription) bool {
	if k.ContentLanguage != p.ContentLanguage {
		p.ContentLanguage = k.ContentLanguage
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeS3BucketObject_KmsKeyId(k *S3BucketObjectParameters, p *S3BucketObjectParameters, md *plugin.MergeDescription) bool {
	if k.KmsKeyId != p.KmsKeyId {
		p.KmsKeyId = k.KmsKeyId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeS3BucketObject_Content(k *S3BucketObjectParameters, p *S3BucketObjectParameters, md *plugin.MergeDescription) bool {
	if k.Content != p.Content {
		p.Content = k.Content
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeS3BucketObject_ContentEncoding(k *S3BucketObjectParameters, p *S3BucketObjectParameters, md *plugin.MergeDescription) bool {
	if k.ContentEncoding != p.ContentEncoding {
		p.ContentEncoding = k.ContentEncoding
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeS3BucketObject_ForceDestroy(k *S3BucketObjectParameters, p *S3BucketObjectParameters, md *plugin.MergeDescription) bool {
	if k.ForceDestroy != p.ForceDestroy {
		p.ForceDestroy = k.ForceDestroy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeS3BucketObject_Key(k *S3BucketObjectParameters, p *S3BucketObjectParameters, md *plugin.MergeDescription) bool {
	if k.Key != p.Key {
		p.Key = k.Key
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeS3BucketObject_VersionId(k *S3BucketObjectObservation, p *S3BucketObjectObservation, md *plugin.MergeDescription) bool {
	if k.VersionId != p.VersionId {
		k.VersionId = p.VersionId
		md.StatusUpdated = true
		return true
	}
	return false
}