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
	k := kube.(*KmsExternalKey)
	p := prov.(*KmsExternalKey)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeKmsExternalKey_Policy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKmsExternalKey_ValidTo(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKmsExternalKey_Enabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKmsExternalKey_KeyMaterialBase64(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKmsExternalKey_DeletionWindowInDays(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKmsExternalKey_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKmsExternalKey_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKmsExternalKey_KeyUsage(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKmsExternalKey_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKmsExternalKey_KeyState(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKmsExternalKey_ExpirationModel(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeKmsExternalKey_Policy(k *KmsExternalKeyParameters, p *KmsExternalKeyParameters, md *plugin.MergeDescription) bool {
	if k.Policy != p.Policy {
		p.Policy = k.Policy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeKmsExternalKey_ValidTo(k *KmsExternalKeyParameters, p *KmsExternalKeyParameters, md *plugin.MergeDescription) bool {
	if k.ValidTo != p.ValidTo {
		p.ValidTo = k.ValidTo
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeKmsExternalKey_Enabled(k *KmsExternalKeyParameters, p *KmsExternalKeyParameters, md *plugin.MergeDescription) bool {
	if k.Enabled != p.Enabled {
		p.Enabled = k.Enabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeKmsExternalKey_KeyMaterialBase64(k *KmsExternalKeyParameters, p *KmsExternalKeyParameters, md *plugin.MergeDescription) bool {
	if k.KeyMaterialBase64 != p.KeyMaterialBase64 {
		p.KeyMaterialBase64 = k.KeyMaterialBase64
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeKmsExternalKey_DeletionWindowInDays(k *KmsExternalKeyParameters, p *KmsExternalKeyParameters, md *plugin.MergeDescription) bool {
	if k.DeletionWindowInDays != p.DeletionWindowInDays {
		p.DeletionWindowInDays = k.DeletionWindowInDays
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeKmsExternalKey_Description(k *KmsExternalKeyParameters, p *KmsExternalKeyParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeKmsExternalKey_Tags(k *KmsExternalKeyParameters, p *KmsExternalKeyParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeKmsExternalKey_KeyUsage(k *KmsExternalKeyObservation, p *KmsExternalKeyObservation, md *plugin.MergeDescription) bool {
	if k.KeyUsage != p.KeyUsage {
		k.KeyUsage = p.KeyUsage
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeKmsExternalKey_Arn(k *KmsExternalKeyObservation, p *KmsExternalKeyObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeKmsExternalKey_KeyState(k *KmsExternalKeyObservation, p *KmsExternalKeyObservation, md *plugin.MergeDescription) bool {
	if k.KeyState != p.KeyState {
		k.KeyState = p.KeyState
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeKmsExternalKey_ExpirationModel(k *KmsExternalKeyObservation, p *KmsExternalKeyObservation, md *plugin.MergeDescription) bool {
	if k.ExpirationModel != p.ExpirationModel {
		k.ExpirationModel = p.ExpirationModel
		md.StatusUpdated = true
		return true
	}
	return false
}