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
	k := kube.(*Apigatewayv2IntegrationResponse)
	p := prov.(*Apigatewayv2IntegrationResponse)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeApigatewayv2IntegrationResponse_IntegrationId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApigatewayv2IntegrationResponse_IntegrationResponseKey(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApigatewayv2IntegrationResponse_ResponseTemplates(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApigatewayv2IntegrationResponse_TemplateSelectionExpression(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApigatewayv2IntegrationResponse_ApiId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApigatewayv2IntegrationResponse_ContentHandlingStrategy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeApigatewayv2IntegrationResponse_IntegrationId(k *Apigatewayv2IntegrationResponseParameters, p *Apigatewayv2IntegrationResponseParameters, md *plugin.MergeDescription) bool {
	if k.IntegrationId != p.IntegrationId {
		p.IntegrationId = k.IntegrationId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApigatewayv2IntegrationResponse_IntegrationResponseKey(k *Apigatewayv2IntegrationResponseParameters, p *Apigatewayv2IntegrationResponseParameters, md *plugin.MergeDescription) bool {
	if k.IntegrationResponseKey != p.IntegrationResponseKey {
		p.IntegrationResponseKey = k.IntegrationResponseKey
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeApigatewayv2IntegrationResponse_ResponseTemplates(k *Apigatewayv2IntegrationResponseParameters, p *Apigatewayv2IntegrationResponseParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.ResponseTemplates, p.ResponseTemplates) {
		p.ResponseTemplates = k.ResponseTemplates
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApigatewayv2IntegrationResponse_TemplateSelectionExpression(k *Apigatewayv2IntegrationResponseParameters, p *Apigatewayv2IntegrationResponseParameters, md *plugin.MergeDescription) bool {
	if k.TemplateSelectionExpression != p.TemplateSelectionExpression {
		p.TemplateSelectionExpression = k.TemplateSelectionExpression
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApigatewayv2IntegrationResponse_ApiId(k *Apigatewayv2IntegrationResponseParameters, p *Apigatewayv2IntegrationResponseParameters, md *plugin.MergeDescription) bool {
	if k.ApiId != p.ApiId {
		p.ApiId = k.ApiId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApigatewayv2IntegrationResponse_ContentHandlingStrategy(k *Apigatewayv2IntegrationResponseParameters, p *Apigatewayv2IntegrationResponseParameters, md *plugin.MergeDescription) bool {
	if k.ContentHandlingStrategy != p.ContentHandlingStrategy {
		p.ContentHandlingStrategy = k.ContentHandlingStrategy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}