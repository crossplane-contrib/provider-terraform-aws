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
	k := kube.(*CodebuildSourceCredential)
	p := prov.(*CodebuildSourceCredential)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeCodebuildSourceCredential_AuthType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCodebuildSourceCredential_ServerType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCodebuildSourceCredential_Token(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCodebuildSourceCredential_UserName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCodebuildSourceCredential_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeCodebuildSourceCredential_AuthType(k *CodebuildSourceCredentialParameters, p *CodebuildSourceCredentialParameters, md *plugin.MergeDescription) bool {
	if k.AuthType != p.AuthType {
		p.AuthType = k.AuthType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCodebuildSourceCredential_ServerType(k *CodebuildSourceCredentialParameters, p *CodebuildSourceCredentialParameters, md *plugin.MergeDescription) bool {
	if k.ServerType != p.ServerType {
		p.ServerType = k.ServerType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCodebuildSourceCredential_Token(k *CodebuildSourceCredentialParameters, p *CodebuildSourceCredentialParameters, md *plugin.MergeDescription) bool {
	if k.Token != p.Token {
		p.Token = k.Token
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCodebuildSourceCredential_UserName(k *CodebuildSourceCredentialParameters, p *CodebuildSourceCredentialParameters, md *plugin.MergeDescription) bool {
	if k.UserName != p.UserName {
		p.UserName = k.UserName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCodebuildSourceCredential_Arn(k *CodebuildSourceCredentialObservation, p *CodebuildSourceCredentialObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}