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
	k := kube.(*ApiGatewayIntegration)
	p := prov.(*ApiGatewayIntegration)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeApiGatewayIntegration_Credentials(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayIntegration_RequestParameters(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayIntegration_ResourceId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayIntegration_CacheNamespace(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayIntegration_ConnectionType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayIntegration_ContentHandling(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayIntegration_TimeoutMilliseconds(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayIntegration_Type(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayIntegration_CacheKeyParameters(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayIntegration_IntegrationHttpMethod(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayIntegration_RestApiId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayIntegration_PassthroughBehavior(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayIntegration_RequestTemplates(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayIntegration_Uri(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayIntegration_ConnectionId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayIntegration_HttpMethod(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeApiGatewayIntegration_Credentials(k *ApiGatewayIntegrationParameters, p *ApiGatewayIntegrationParameters, md *plugin.MergeDescription) bool {
	if k.Credentials != p.Credentials {
		p.Credentials = k.Credentials
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeApiGatewayIntegration_RequestParameters(k *ApiGatewayIntegrationParameters, p *ApiGatewayIntegrationParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.RequestParameters, p.RequestParameters) {
		p.RequestParameters = k.RequestParameters
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayIntegration_ResourceId(k *ApiGatewayIntegrationParameters, p *ApiGatewayIntegrationParameters, md *plugin.MergeDescription) bool {
	if k.ResourceId != p.ResourceId {
		p.ResourceId = k.ResourceId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayIntegration_CacheNamespace(k *ApiGatewayIntegrationParameters, p *ApiGatewayIntegrationParameters, md *plugin.MergeDescription) bool {
	if k.CacheNamespace != p.CacheNamespace {
		p.CacheNamespace = k.CacheNamespace
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayIntegration_ConnectionType(k *ApiGatewayIntegrationParameters, p *ApiGatewayIntegrationParameters, md *plugin.MergeDescription) bool {
	if k.ConnectionType != p.ConnectionType {
		p.ConnectionType = k.ConnectionType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayIntegration_ContentHandling(k *ApiGatewayIntegrationParameters, p *ApiGatewayIntegrationParameters, md *plugin.MergeDescription) bool {
	if k.ContentHandling != p.ContentHandling {
		p.ContentHandling = k.ContentHandling
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayIntegration_TimeoutMilliseconds(k *ApiGatewayIntegrationParameters, p *ApiGatewayIntegrationParameters, md *plugin.MergeDescription) bool {
	if k.TimeoutMilliseconds != p.TimeoutMilliseconds {
		p.TimeoutMilliseconds = k.TimeoutMilliseconds
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayIntegration_Type(k *ApiGatewayIntegrationParameters, p *ApiGatewayIntegrationParameters, md *plugin.MergeDescription) bool {
	if k.Type != p.Type {
		p.Type = k.Type
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeApiGatewayIntegration_CacheKeyParameters(k *ApiGatewayIntegrationParameters, p *ApiGatewayIntegrationParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.CacheKeyParameters, p.CacheKeyParameters) {
		p.CacheKeyParameters = k.CacheKeyParameters
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayIntegration_IntegrationHttpMethod(k *ApiGatewayIntegrationParameters, p *ApiGatewayIntegrationParameters, md *plugin.MergeDescription) bool {
	if k.IntegrationHttpMethod != p.IntegrationHttpMethod {
		p.IntegrationHttpMethod = k.IntegrationHttpMethod
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayIntegration_RestApiId(k *ApiGatewayIntegrationParameters, p *ApiGatewayIntegrationParameters, md *plugin.MergeDescription) bool {
	if k.RestApiId != p.RestApiId {
		p.RestApiId = k.RestApiId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayIntegration_PassthroughBehavior(k *ApiGatewayIntegrationParameters, p *ApiGatewayIntegrationParameters, md *plugin.MergeDescription) bool {
	if k.PassthroughBehavior != p.PassthroughBehavior {
		p.PassthroughBehavior = k.PassthroughBehavior
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeApiGatewayIntegration_RequestTemplates(k *ApiGatewayIntegrationParameters, p *ApiGatewayIntegrationParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.RequestTemplates, p.RequestTemplates) {
		p.RequestTemplates = k.RequestTemplates
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayIntegration_Uri(k *ApiGatewayIntegrationParameters, p *ApiGatewayIntegrationParameters, md *plugin.MergeDescription) bool {
	if k.Uri != p.Uri {
		p.Uri = k.Uri
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayIntegration_ConnectionId(k *ApiGatewayIntegrationParameters, p *ApiGatewayIntegrationParameters, md *plugin.MergeDescription) bool {
	if k.ConnectionId != p.ConnectionId {
		p.ConnectionId = k.ConnectionId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayIntegration_HttpMethod(k *ApiGatewayIntegrationParameters, p *ApiGatewayIntegrationParameters, md *plugin.MergeDescription) bool {
	if k.HttpMethod != p.HttpMethod {
		p.HttpMethod = k.HttpMethod
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}