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
	k := kube.(*GlacierVaultLock)
	p := prov.(*GlacierVaultLock)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeGlacierVaultLock_IgnoreDeletionError(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeGlacierVaultLock_Policy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeGlacierVaultLock_VaultName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeGlacierVaultLock_CompleteLock(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeGlacierVaultLock_IgnoreDeletionError(k *GlacierVaultLockParameters, p *GlacierVaultLockParameters, md *plugin.MergeDescription) bool {
	if k.IgnoreDeletionError != p.IgnoreDeletionError {
		p.IgnoreDeletionError = k.IgnoreDeletionError
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeGlacierVaultLock_Policy(k *GlacierVaultLockParameters, p *GlacierVaultLockParameters, md *plugin.MergeDescription) bool {
	if k.Policy != p.Policy {
		p.Policy = k.Policy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeGlacierVaultLock_VaultName(k *GlacierVaultLockParameters, p *GlacierVaultLockParameters, md *plugin.MergeDescription) bool {
	if k.VaultName != p.VaultName {
		p.VaultName = k.VaultName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeGlacierVaultLock_CompleteLock(k *GlacierVaultLockParameters, p *GlacierVaultLockParameters, md *plugin.MergeDescription) bool {
	if k.CompleteLock != p.CompleteLock {
		p.CompleteLock = k.CompleteLock
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}