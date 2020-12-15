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
	k := kube.(*Apigatewayv2Route)
	p := prov.(*Apigatewayv2Route)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeApigatewayv2Route_RouteKey(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApigatewayv2Route_Target(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApigatewayv2Route_ModelSelectionExpression(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApigatewayv2Route_ApiKeyRequired(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApigatewayv2Route_AuthorizationScopes(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApigatewayv2Route_AuthorizationType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApigatewayv2Route_AuthorizerId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApigatewayv2Route_OperationName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApigatewayv2Route_RequestModels(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApigatewayv2Route_ApiId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApigatewayv2Route_RouteResponseSelectionExpression(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeApigatewayv2Route_RouteKey(k *Apigatewayv2RouteParameters, p *Apigatewayv2RouteParameters, md *plugin.MergeDescription) bool {
	if k.RouteKey != p.RouteKey {
		p.RouteKey = k.RouteKey
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApigatewayv2Route_Target(k *Apigatewayv2RouteParameters, p *Apigatewayv2RouteParameters, md *plugin.MergeDescription) bool {
	if k.Target != p.Target {
		p.Target = k.Target
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApigatewayv2Route_ModelSelectionExpression(k *Apigatewayv2RouteParameters, p *Apigatewayv2RouteParameters, md *plugin.MergeDescription) bool {
	if k.ModelSelectionExpression != p.ModelSelectionExpression {
		p.ModelSelectionExpression = k.ModelSelectionExpression
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApigatewayv2Route_ApiKeyRequired(k *Apigatewayv2RouteParameters, p *Apigatewayv2RouteParameters, md *plugin.MergeDescription) bool {
	if k.ApiKeyRequired != p.ApiKeyRequired {
		p.ApiKeyRequired = k.ApiKeyRequired
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeApigatewayv2Route_AuthorizationScopes(k *Apigatewayv2RouteParameters, p *Apigatewayv2RouteParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.AuthorizationScopes, p.AuthorizationScopes) {
		p.AuthorizationScopes = k.AuthorizationScopes
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApigatewayv2Route_AuthorizationType(k *Apigatewayv2RouteParameters, p *Apigatewayv2RouteParameters, md *plugin.MergeDescription) bool {
	if k.AuthorizationType != p.AuthorizationType {
		p.AuthorizationType = k.AuthorizationType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApigatewayv2Route_AuthorizerId(k *Apigatewayv2RouteParameters, p *Apigatewayv2RouteParameters, md *plugin.MergeDescription) bool {
	if k.AuthorizerId != p.AuthorizerId {
		p.AuthorizerId = k.AuthorizerId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApigatewayv2Route_OperationName(k *Apigatewayv2RouteParameters, p *Apigatewayv2RouteParameters, md *plugin.MergeDescription) bool {
	if k.OperationName != p.OperationName {
		p.OperationName = k.OperationName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeApigatewayv2Route_RequestModels(k *Apigatewayv2RouteParameters, p *Apigatewayv2RouteParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.RequestModels, p.RequestModels) {
		p.RequestModels = k.RequestModels
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApigatewayv2Route_ApiId(k *Apigatewayv2RouteParameters, p *Apigatewayv2RouteParameters, md *plugin.MergeDescription) bool {
	if k.ApiId != p.ApiId {
		p.ApiId = k.ApiId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApigatewayv2Route_RouteResponseSelectionExpression(k *Apigatewayv2RouteParameters, p *Apigatewayv2RouteParameters, md *plugin.MergeDescription) bool {
	if k.RouteResponseSelectionExpression != p.RouteResponseSelectionExpression {
		p.RouteResponseSelectionExpression = k.RouteResponseSelectionExpression
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}