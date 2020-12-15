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
	k := kube.(*CloudfrontOriginAccessIdentity)
	p := prov.(*CloudfrontOriginAccessIdentity)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeCloudfrontOriginAccessIdentity_Comment(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudfrontOriginAccessIdentity_Etag(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudfrontOriginAccessIdentity_IamArn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudfrontOriginAccessIdentity_S3CanonicalUserId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudfrontOriginAccessIdentity_CallerReference(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudfrontOriginAccessIdentity_CloudfrontAccessIdentityPath(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeCloudfrontOriginAccessIdentity_Comment(k *CloudfrontOriginAccessIdentityParameters, p *CloudfrontOriginAccessIdentityParameters, md *plugin.MergeDescription) bool {
	if k.Comment != p.Comment {
		p.Comment = k.Comment
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCloudfrontOriginAccessIdentity_Etag(k *CloudfrontOriginAccessIdentityObservation, p *CloudfrontOriginAccessIdentityObservation, md *plugin.MergeDescription) bool {
	if k.Etag != p.Etag {
		k.Etag = p.Etag
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCloudfrontOriginAccessIdentity_IamArn(k *CloudfrontOriginAccessIdentityObservation, p *CloudfrontOriginAccessIdentityObservation, md *plugin.MergeDescription) bool {
	if k.IamArn != p.IamArn {
		k.IamArn = p.IamArn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCloudfrontOriginAccessIdentity_S3CanonicalUserId(k *CloudfrontOriginAccessIdentityObservation, p *CloudfrontOriginAccessIdentityObservation, md *plugin.MergeDescription) bool {
	if k.S3CanonicalUserId != p.S3CanonicalUserId {
		k.S3CanonicalUserId = p.S3CanonicalUserId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCloudfrontOriginAccessIdentity_CallerReference(k *CloudfrontOriginAccessIdentityObservation, p *CloudfrontOriginAccessIdentityObservation, md *plugin.MergeDescription) bool {
	if k.CallerReference != p.CallerReference {
		k.CallerReference = p.CallerReference
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCloudfrontOriginAccessIdentity_CloudfrontAccessIdentityPath(k *CloudfrontOriginAccessIdentityObservation, p *CloudfrontOriginAccessIdentityObservation, md *plugin.MergeDescription) bool {
	if k.CloudfrontAccessIdentityPath != p.CloudfrontAccessIdentityPath {
		k.CloudfrontAccessIdentityPath = p.CloudfrontAccessIdentityPath
		md.StatusUpdated = true
		return true
	}
	return false
}