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
	k := kube.(*OpsworksPermission)
	p := prov.(*OpsworksPermission)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeOpsworksPermission_Level(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeOpsworksPermission_StackId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeOpsworksPermission_UserArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeOpsworksPermission_AllowSsh(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeOpsworksPermission_AllowSudo(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeOpsworksPermission_Level(k *OpsworksPermissionParameters, p *OpsworksPermissionParameters, md *plugin.MergeDescription) bool {
	if k.Level != p.Level {
		p.Level = k.Level
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeOpsworksPermission_StackId(k *OpsworksPermissionParameters, p *OpsworksPermissionParameters, md *plugin.MergeDescription) bool {
	if k.StackId != p.StackId {
		p.StackId = k.StackId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeOpsworksPermission_UserArn(k *OpsworksPermissionParameters, p *OpsworksPermissionParameters, md *plugin.MergeDescription) bool {
	if k.UserArn != p.UserArn {
		p.UserArn = k.UserArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeOpsworksPermission_AllowSsh(k *OpsworksPermissionParameters, p *OpsworksPermissionParameters, md *plugin.MergeDescription) bool {
	if k.AllowSsh != p.AllowSsh {
		p.AllowSsh = k.AllowSsh
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeOpsworksPermission_AllowSudo(k *OpsworksPermissionParameters, p *OpsworksPermissionParameters, md *plugin.MergeDescription) bool {
	if k.AllowSudo != p.AllowSudo {
		p.AllowSudo = k.AllowSudo
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}