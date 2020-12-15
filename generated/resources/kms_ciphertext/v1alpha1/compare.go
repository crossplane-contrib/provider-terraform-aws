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
	k := kube.(*KmsCiphertext)
	p := prov.(*KmsCiphertext)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeKmsCiphertext_Context(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKmsCiphertext_KeyId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKmsCiphertext_Plaintext(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKmsCiphertext_CiphertextBlob(&k.Status.AtProvider, &p.Status.AtProvider, md)
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

//mergePrimitiveContainerTemplateSpec
func MergeKmsCiphertext_Context(k *KmsCiphertextParameters, p *KmsCiphertextParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Context, p.Context) {
		p.Context = k.Context
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeKmsCiphertext_KeyId(k *KmsCiphertextParameters, p *KmsCiphertextParameters, md *plugin.MergeDescription) bool {
	if k.KeyId != p.KeyId {
		p.KeyId = k.KeyId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeKmsCiphertext_Plaintext(k *KmsCiphertextParameters, p *KmsCiphertextParameters, md *plugin.MergeDescription) bool {
	if k.Plaintext != p.Plaintext {
		p.Plaintext = k.Plaintext
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeKmsCiphertext_CiphertextBlob(k *KmsCiphertextObservation, p *KmsCiphertextObservation, md *plugin.MergeDescription) bool {
	if k.CiphertextBlob != p.CiphertextBlob {
		k.CiphertextBlob = p.CiphertextBlob
		md.StatusUpdated = true
		return true
	}
	return false
}