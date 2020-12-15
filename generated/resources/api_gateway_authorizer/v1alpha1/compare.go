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
	k := kube.(*ApiGatewayAuthorizer)
	p := prov.(*ApiGatewayAuthorizer)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeApiGatewayAuthorizer_AuthorizerResultTtlInSeconds(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayAuthorizer_IdentityValidationExpression(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayAuthorizer_ProviderArns(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayAuthorizer_RestApiId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayAuthorizer_AuthorizerCredentials(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayAuthorizer_AuthorizerUri(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayAuthorizer_IdentitySource(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayAuthorizer_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayAuthorizer_Type(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeApiGatewayAuthorizer_AuthorizerResultTtlInSeconds(k *ApiGatewayAuthorizerParameters, p *ApiGatewayAuthorizerParameters, md *plugin.MergeDescription) bool {
	if k.AuthorizerResultTtlInSeconds != p.AuthorizerResultTtlInSeconds {
		p.AuthorizerResultTtlInSeconds = k.AuthorizerResultTtlInSeconds
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayAuthorizer_IdentityValidationExpression(k *ApiGatewayAuthorizerParameters, p *ApiGatewayAuthorizerParameters, md *plugin.MergeDescription) bool {
	if k.IdentityValidationExpression != p.IdentityValidationExpression {
		p.IdentityValidationExpression = k.IdentityValidationExpression
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeApiGatewayAuthorizer_ProviderArns(k *ApiGatewayAuthorizerParameters, p *ApiGatewayAuthorizerParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.ProviderArns, p.ProviderArns) {
		p.ProviderArns = k.ProviderArns
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayAuthorizer_RestApiId(k *ApiGatewayAuthorizerParameters, p *ApiGatewayAuthorizerParameters, md *plugin.MergeDescription) bool {
	if k.RestApiId != p.RestApiId {
		p.RestApiId = k.RestApiId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayAuthorizer_AuthorizerCredentials(k *ApiGatewayAuthorizerParameters, p *ApiGatewayAuthorizerParameters, md *plugin.MergeDescription) bool {
	if k.AuthorizerCredentials != p.AuthorizerCredentials {
		p.AuthorizerCredentials = k.AuthorizerCredentials
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayAuthorizer_AuthorizerUri(k *ApiGatewayAuthorizerParameters, p *ApiGatewayAuthorizerParameters, md *plugin.MergeDescription) bool {
	if k.AuthorizerUri != p.AuthorizerUri {
		p.AuthorizerUri = k.AuthorizerUri
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayAuthorizer_IdentitySource(k *ApiGatewayAuthorizerParameters, p *ApiGatewayAuthorizerParameters, md *plugin.MergeDescription) bool {
	if k.IdentitySource != p.IdentitySource {
		p.IdentitySource = k.IdentitySource
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayAuthorizer_Name(k *ApiGatewayAuthorizerParameters, p *ApiGatewayAuthorizerParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayAuthorizer_Type(k *ApiGatewayAuthorizerParameters, p *ApiGatewayAuthorizerParameters, md *plugin.MergeDescription) bool {
	if k.Type != p.Type {
		p.Type = k.Type
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}