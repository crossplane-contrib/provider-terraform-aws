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
	k := kube.(*LightsailKeyPair)
	p := prov.(*LightsailKeyPair)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeLightsailKeyPair_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLightsailKeyPair_NamePrefix(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLightsailKeyPair_PgpKey(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLightsailKeyPair_PublicKey(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLightsailKeyPair_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLightsailKeyPair_EncryptedFingerprint(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLightsailKeyPair_EncryptedPrivateKey(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLightsailKeyPair_PrivateKey(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLightsailKeyPair_Fingerprint(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeLightsailKeyPair_Name(k *LightsailKeyPairParameters, p *LightsailKeyPairParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLightsailKeyPair_NamePrefix(k *LightsailKeyPairParameters, p *LightsailKeyPairParameters, md *plugin.MergeDescription) bool {
	if k.NamePrefix != p.NamePrefix {
		p.NamePrefix = k.NamePrefix
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLightsailKeyPair_PgpKey(k *LightsailKeyPairParameters, p *LightsailKeyPairParameters, md *plugin.MergeDescription) bool {
	if k.PgpKey != p.PgpKey {
		p.PgpKey = k.PgpKey
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLightsailKeyPair_PublicKey(k *LightsailKeyPairParameters, p *LightsailKeyPairParameters, md *plugin.MergeDescription) bool {
	if k.PublicKey != p.PublicKey {
		p.PublicKey = k.PublicKey
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeLightsailKeyPair_Arn(k *LightsailKeyPairObservation, p *LightsailKeyPairObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeLightsailKeyPair_EncryptedFingerprint(k *LightsailKeyPairObservation, p *LightsailKeyPairObservation, md *plugin.MergeDescription) bool {
	if k.EncryptedFingerprint != p.EncryptedFingerprint {
		k.EncryptedFingerprint = p.EncryptedFingerprint
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeLightsailKeyPair_EncryptedPrivateKey(k *LightsailKeyPairObservation, p *LightsailKeyPairObservation, md *plugin.MergeDescription) bool {
	if k.EncryptedPrivateKey != p.EncryptedPrivateKey {
		k.EncryptedPrivateKey = p.EncryptedPrivateKey
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeLightsailKeyPair_PrivateKey(k *LightsailKeyPairObservation, p *LightsailKeyPairObservation, md *plugin.MergeDescription) bool {
	if k.PrivateKey != p.PrivateKey {
		k.PrivateKey = p.PrivateKey
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeLightsailKeyPair_Fingerprint(k *LightsailKeyPairObservation, p *LightsailKeyPairObservation, md *plugin.MergeDescription) bool {
	if k.Fingerprint != p.Fingerprint {
		k.Fingerprint = p.Fingerprint
		md.StatusUpdated = true
		return true
	}
	return false
}