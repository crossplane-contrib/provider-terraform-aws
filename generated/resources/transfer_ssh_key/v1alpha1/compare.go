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
	k := kube.(*TransferSshKey)
	p := prov.(*TransferSshKey)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeTransferSshKey_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeTransferSshKey_ServerId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeTransferSshKey_UserName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeTransferSshKey_Body(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeTransferSshKey_Id(k *TransferSshKeyParameters, p *TransferSshKeyParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeTransferSshKey_ServerId(k *TransferSshKeyParameters, p *TransferSshKeyParameters, md *plugin.MergeDescription) bool {
	if k.ServerId != p.ServerId {
		p.ServerId = k.ServerId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeTransferSshKey_UserName(k *TransferSshKeyParameters, p *TransferSshKeyParameters, md *plugin.MergeDescription) bool {
	if k.UserName != p.UserName {
		p.UserName = k.UserName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeTransferSshKey_Body(k *TransferSshKeyParameters, p *TransferSshKeyParameters, md *plugin.MergeDescription) bool {
	if k.Body != p.Body {
		p.Body = k.Body
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}