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
	k := kube.(*KeyPair)
	p := prov.(*KeyPair)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeKeyPair_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKeyPair_KeyName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKeyPair_KeyNamePrefix(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKeyPair_PublicKey(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKeyPair_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKeyPair_Fingerprint(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKeyPair_KeyPairId(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeKeyPair_Tags(k *KeyPairParameters, p *KeyPairParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeKeyPair_KeyName(k *KeyPairParameters, p *KeyPairParameters, md *plugin.MergeDescription) bool {
	if k.KeyName != p.KeyName {
		p.KeyName = k.KeyName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeKeyPair_KeyNamePrefix(k *KeyPairParameters, p *KeyPairParameters, md *plugin.MergeDescription) bool {
	if k.KeyNamePrefix != p.KeyNamePrefix {
		p.KeyNamePrefix = k.KeyNamePrefix
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeKeyPair_PublicKey(k *KeyPairParameters, p *KeyPairParameters, md *plugin.MergeDescription) bool {
	if k.PublicKey != p.PublicKey {
		p.PublicKey = k.PublicKey
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeKeyPair_Arn(k *KeyPairObservation, p *KeyPairObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeKeyPair_Fingerprint(k *KeyPairObservation, p *KeyPairObservation, md *plugin.MergeDescription) bool {
	if k.Fingerprint != p.Fingerprint {
		k.Fingerprint = p.Fingerprint
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeKeyPair_KeyPairId(k *KeyPairObservation, p *KeyPairObservation, md *plugin.MergeDescription) bool {
	if k.KeyPairId != p.KeyPairId {
		k.KeyPairId = p.KeyPairId
		md.StatusUpdated = true
		return true
	}
	return false
}