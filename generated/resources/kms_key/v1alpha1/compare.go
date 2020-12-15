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
	k := kube.(*KmsKey)
	p := prov.(*KmsKey)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeKmsKey_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKmsKey_EnableKeyRotation(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKmsKey_IsEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKmsKey_KeyUsage(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKmsKey_Policy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKmsKey_CustomerMasterKeySpec(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKmsKey_DeletionWindowInDays(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKmsKey_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKmsKey_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeKmsKey_KeyId(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeKmsKey_Tags(k *KmsKeyParameters, p *KmsKeyParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeKmsKey_EnableKeyRotation(k *KmsKeyParameters, p *KmsKeyParameters, md *plugin.MergeDescription) bool {
	if k.EnableKeyRotation != p.EnableKeyRotation {
		p.EnableKeyRotation = k.EnableKeyRotation
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeKmsKey_IsEnabled(k *KmsKeyParameters, p *KmsKeyParameters, md *plugin.MergeDescription) bool {
	if k.IsEnabled != p.IsEnabled {
		p.IsEnabled = k.IsEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeKmsKey_KeyUsage(k *KmsKeyParameters, p *KmsKeyParameters, md *plugin.MergeDescription) bool {
	if k.KeyUsage != p.KeyUsage {
		p.KeyUsage = k.KeyUsage
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeKmsKey_Policy(k *KmsKeyParameters, p *KmsKeyParameters, md *plugin.MergeDescription) bool {
	if k.Policy != p.Policy {
		p.Policy = k.Policy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeKmsKey_CustomerMasterKeySpec(k *KmsKeyParameters, p *KmsKeyParameters, md *plugin.MergeDescription) bool {
	if k.CustomerMasterKeySpec != p.CustomerMasterKeySpec {
		p.CustomerMasterKeySpec = k.CustomerMasterKeySpec
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeKmsKey_DeletionWindowInDays(k *KmsKeyParameters, p *KmsKeyParameters, md *plugin.MergeDescription) bool {
	if k.DeletionWindowInDays != p.DeletionWindowInDays {
		p.DeletionWindowInDays = k.DeletionWindowInDays
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeKmsKey_Description(k *KmsKeyParameters, p *KmsKeyParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeKmsKey_Arn(k *KmsKeyObservation, p *KmsKeyObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeKmsKey_KeyId(k *KmsKeyObservation, p *KmsKeyObservation, md *plugin.MergeDescription) bool {
	if k.KeyId != p.KeyId {
		k.KeyId = p.KeyId
		md.StatusUpdated = true
		return true
	}
	return false
}