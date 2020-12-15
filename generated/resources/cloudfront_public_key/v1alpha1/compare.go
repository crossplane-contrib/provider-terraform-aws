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
	k := kube.(*CloudfrontPublicKey)
	p := prov.(*CloudfrontPublicKey)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeCloudfrontPublicKey_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudfrontPublicKey_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudfrontPublicKey_NamePrefix(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudfrontPublicKey_Comment(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudfrontPublicKey_EncodedKey(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudfrontPublicKey_CallerReference(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCloudfrontPublicKey_Etag(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeCloudfrontPublicKey_Id(k *CloudfrontPublicKeyParameters, p *CloudfrontPublicKeyParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudfrontPublicKey_Name(k *CloudfrontPublicKeyParameters, p *CloudfrontPublicKeyParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudfrontPublicKey_NamePrefix(k *CloudfrontPublicKeyParameters, p *CloudfrontPublicKeyParameters, md *plugin.MergeDescription) bool {
	if k.NamePrefix != p.NamePrefix {
		p.NamePrefix = k.NamePrefix
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudfrontPublicKey_Comment(k *CloudfrontPublicKeyParameters, p *CloudfrontPublicKeyParameters, md *plugin.MergeDescription) bool {
	if k.Comment != p.Comment {
		p.Comment = k.Comment
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCloudfrontPublicKey_EncodedKey(k *CloudfrontPublicKeyParameters, p *CloudfrontPublicKeyParameters, md *plugin.MergeDescription) bool {
	if k.EncodedKey != p.EncodedKey {
		p.EncodedKey = k.EncodedKey
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCloudfrontPublicKey_CallerReference(k *CloudfrontPublicKeyObservation, p *CloudfrontPublicKeyObservation, md *plugin.MergeDescription) bool {
	if k.CallerReference != p.CallerReference {
		k.CallerReference = p.CallerReference
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCloudfrontPublicKey_Etag(k *CloudfrontPublicKeyObservation, p *CloudfrontPublicKeyObservation, md *plugin.MergeDescription) bool {
	if k.Etag != p.Etag {
		k.Etag = p.Etag
		md.StatusUpdated = true
		return true
	}
	return false
}