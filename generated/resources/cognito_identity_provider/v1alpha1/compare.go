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
	k := kube.(*CognitoIdentityProvider)
	p := prov.(*CognitoIdentityProvider)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeCognitoIdentityProvider_IdpIdentifiers(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCognitoIdentityProvider_ProviderDetails(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCognitoIdentityProvider_ProviderName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCognitoIdentityProvider_ProviderType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCognitoIdentityProvider_UserPoolId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCognitoIdentityProvider_AttributeMapping(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeCognitoIdentityProvider_IdpIdentifiers(k *CognitoIdentityProviderParameters, p *CognitoIdentityProviderParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.IdpIdentifiers, p.IdpIdentifiers) {
		p.IdpIdentifiers = k.IdpIdentifiers
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeCognitoIdentityProvider_ProviderDetails(k *CognitoIdentityProviderParameters, p *CognitoIdentityProviderParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.ProviderDetails, p.ProviderDetails) {
		p.ProviderDetails = k.ProviderDetails
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCognitoIdentityProvider_ProviderName(k *CognitoIdentityProviderParameters, p *CognitoIdentityProviderParameters, md *plugin.MergeDescription) bool {
	if k.ProviderName != p.ProviderName {
		p.ProviderName = k.ProviderName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCognitoIdentityProvider_ProviderType(k *CognitoIdentityProviderParameters, p *CognitoIdentityProviderParameters, md *plugin.MergeDescription) bool {
	if k.ProviderType != p.ProviderType {
		p.ProviderType = k.ProviderType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCognitoIdentityProvider_UserPoolId(k *CognitoIdentityProviderParameters, p *CognitoIdentityProviderParameters, md *plugin.MergeDescription) bool {
	if k.UserPoolId != p.UserPoolId {
		p.UserPoolId = k.UserPoolId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeCognitoIdentityProvider_AttributeMapping(k *CognitoIdentityProviderParameters, p *CognitoIdentityProviderParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.AttributeMapping, p.AttributeMapping) {
		p.AttributeMapping = k.AttributeMapping
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}