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
	k := kube.(*S3BucketPublicAccessBlock)
	p := prov.(*S3BucketPublicAccessBlock)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeS3BucketPublicAccessBlock_BlockPublicAcls(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeS3BucketPublicAccessBlock_BlockPublicPolicy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeS3BucketPublicAccessBlock_Bucket(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeS3BucketPublicAccessBlock_IgnorePublicAcls(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeS3BucketPublicAccessBlock_RestrictPublicBuckets(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeS3BucketPublicAccessBlock_BlockPublicAcls(k *S3BucketPublicAccessBlockParameters, p *S3BucketPublicAccessBlockParameters, md *plugin.MergeDescription) bool {
	if k.BlockPublicAcls != p.BlockPublicAcls {
		p.BlockPublicAcls = k.BlockPublicAcls
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeS3BucketPublicAccessBlock_BlockPublicPolicy(k *S3BucketPublicAccessBlockParameters, p *S3BucketPublicAccessBlockParameters, md *plugin.MergeDescription) bool {
	if k.BlockPublicPolicy != p.BlockPublicPolicy {
		p.BlockPublicPolicy = k.BlockPublicPolicy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeS3BucketPublicAccessBlock_Bucket(k *S3BucketPublicAccessBlockParameters, p *S3BucketPublicAccessBlockParameters, md *plugin.MergeDescription) bool {
	if k.Bucket != p.Bucket {
		p.Bucket = k.Bucket
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeS3BucketPublicAccessBlock_IgnorePublicAcls(k *S3BucketPublicAccessBlockParameters, p *S3BucketPublicAccessBlockParameters, md *plugin.MergeDescription) bool {
	if k.IgnorePublicAcls != p.IgnorePublicAcls {
		p.IgnorePublicAcls = k.IgnorePublicAcls
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeS3BucketPublicAccessBlock_RestrictPublicBuckets(k *S3BucketPublicAccessBlockParameters, p *S3BucketPublicAccessBlockParameters, md *plugin.MergeDescription) bool {
	if k.RestrictPublicBuckets != p.RestrictPublicBuckets {
		p.RestrictPublicBuckets = k.RestrictPublicBuckets
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}