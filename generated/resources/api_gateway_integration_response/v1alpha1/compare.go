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
	k := kube.(*ApiGatewayIntegrationResponse)
	p := prov.(*ApiGatewayIntegrationResponse)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeApiGatewayIntegrationResponse_StatusCode(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayIntegrationResponse_ResponseParameters(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayIntegrationResponse_HttpMethod(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayIntegrationResponse_ResourceId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayIntegrationResponse_ResponseTemplates(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayIntegrationResponse_RestApiId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayIntegrationResponse_SelectionPattern(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayIntegrationResponse_ContentHandling(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeApiGatewayIntegrationResponse_StatusCode(k *ApiGatewayIntegrationResponseParameters, p *ApiGatewayIntegrationResponseParameters, md *plugin.MergeDescription) bool {
	if k.StatusCode != p.StatusCode {
		p.StatusCode = k.StatusCode
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeApiGatewayIntegrationResponse_ResponseParameters(k *ApiGatewayIntegrationResponseParameters, p *ApiGatewayIntegrationResponseParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.ResponseParameters, p.ResponseParameters) {
		p.ResponseParameters = k.ResponseParameters
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayIntegrationResponse_HttpMethod(k *ApiGatewayIntegrationResponseParameters, p *ApiGatewayIntegrationResponseParameters, md *plugin.MergeDescription) bool {
	if k.HttpMethod != p.HttpMethod {
		p.HttpMethod = k.HttpMethod
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayIntegrationResponse_ResourceId(k *ApiGatewayIntegrationResponseParameters, p *ApiGatewayIntegrationResponseParameters, md *plugin.MergeDescription) bool {
	if k.ResourceId != p.ResourceId {
		p.ResourceId = k.ResourceId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeApiGatewayIntegrationResponse_ResponseTemplates(k *ApiGatewayIntegrationResponseParameters, p *ApiGatewayIntegrationResponseParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.ResponseTemplates, p.ResponseTemplates) {
		p.ResponseTemplates = k.ResponseTemplates
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayIntegrationResponse_RestApiId(k *ApiGatewayIntegrationResponseParameters, p *ApiGatewayIntegrationResponseParameters, md *plugin.MergeDescription) bool {
	if k.RestApiId != p.RestApiId {
		p.RestApiId = k.RestApiId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayIntegrationResponse_SelectionPattern(k *ApiGatewayIntegrationResponseParameters, p *ApiGatewayIntegrationResponseParameters, md *plugin.MergeDescription) bool {
	if k.SelectionPattern != p.SelectionPattern {
		p.SelectionPattern = k.SelectionPattern
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayIntegrationResponse_ContentHandling(k *ApiGatewayIntegrationResponseParameters, p *ApiGatewayIntegrationResponseParameters, md *plugin.MergeDescription) bool {
	if k.ContentHandling != p.ContentHandling {
		p.ContentHandling = k.ContentHandling
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}