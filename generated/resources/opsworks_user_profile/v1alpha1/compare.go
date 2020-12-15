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
	k := kube.(*OpsworksUserProfile)
	p := prov.(*OpsworksUserProfile)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeOpsworksUserProfile_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeOpsworksUserProfile_SshPublicKey(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeOpsworksUserProfile_SshUsername(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeOpsworksUserProfile_UserArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeOpsworksUserProfile_AllowSelfManagement(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeOpsworksUserProfile_Id(k *OpsworksUserProfileParameters, p *OpsworksUserProfileParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeOpsworksUserProfile_SshPublicKey(k *OpsworksUserProfileParameters, p *OpsworksUserProfileParameters, md *plugin.MergeDescription) bool {
	if k.SshPublicKey != p.SshPublicKey {
		p.SshPublicKey = k.SshPublicKey
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeOpsworksUserProfile_SshUsername(k *OpsworksUserProfileParameters, p *OpsworksUserProfileParameters, md *plugin.MergeDescription) bool {
	if k.SshUsername != p.SshUsername {
		p.SshUsername = k.SshUsername
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeOpsworksUserProfile_UserArn(k *OpsworksUserProfileParameters, p *OpsworksUserProfileParameters, md *plugin.MergeDescription) bool {
	if k.UserArn != p.UserArn {
		p.UserArn = k.UserArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeOpsworksUserProfile_AllowSelfManagement(k *OpsworksUserProfileParameters, p *OpsworksUserProfileParameters, md *plugin.MergeDescription) bool {
	if k.AllowSelfManagement != p.AllowSelfManagement {
		p.AllowSelfManagement = k.AllowSelfManagement
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}