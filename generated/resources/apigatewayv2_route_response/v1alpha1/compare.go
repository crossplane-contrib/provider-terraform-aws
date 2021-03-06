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
	k := kube.(*Apigatewayv2RouteResponse)
	p := prov.(*Apigatewayv2RouteResponse)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeApigatewayv2RouteResponse_ApiId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApigatewayv2RouteResponse_ModelSelectionExpression(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApigatewayv2RouteResponse_ResponseModels(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApigatewayv2RouteResponse_RouteId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApigatewayv2RouteResponse_RouteResponseKey(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeApigatewayv2RouteResponse_ApiId(k *Apigatewayv2RouteResponseParameters, p *Apigatewayv2RouteResponseParameters, md *plugin.MergeDescription) bool {
	if k.ApiId != p.ApiId {
		p.ApiId = k.ApiId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApigatewayv2RouteResponse_ModelSelectionExpression(k *Apigatewayv2RouteResponseParameters, p *Apigatewayv2RouteResponseParameters, md *plugin.MergeDescription) bool {
	if k.ModelSelectionExpression != p.ModelSelectionExpression {
		p.ModelSelectionExpression = k.ModelSelectionExpression
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeApigatewayv2RouteResponse_ResponseModels(k *Apigatewayv2RouteResponseParameters, p *Apigatewayv2RouteResponseParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.ResponseModels, p.ResponseModels) {
		p.ResponseModels = k.ResponseModels
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApigatewayv2RouteResponse_RouteId(k *Apigatewayv2RouteResponseParameters, p *Apigatewayv2RouteResponseParameters, md *plugin.MergeDescription) bool {
	if k.RouteId != p.RouteId {
		p.RouteId = k.RouteId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApigatewayv2RouteResponse_RouteResponseKey(k *Apigatewayv2RouteResponseParameters, p *Apigatewayv2RouteResponseParameters, md *plugin.MergeDescription) bool {
	if k.RouteResponseKey != p.RouteResponseKey {
		p.RouteResponseKey = k.RouteResponseKey
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}