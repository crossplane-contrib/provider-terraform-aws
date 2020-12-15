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
	k := kube.(*Apigatewayv2ApiMapping)
	p := prov.(*Apigatewayv2ApiMapping)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeApigatewayv2ApiMapping_ApiMappingKey(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApigatewayv2ApiMapping_DomainName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApigatewayv2ApiMapping_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApigatewayv2ApiMapping_Stage(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApigatewayv2ApiMapping_ApiId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeApigatewayv2ApiMapping_ApiMappingKey(k *Apigatewayv2ApiMappingParameters, p *Apigatewayv2ApiMappingParameters, md *plugin.MergeDescription) bool {
	if k.ApiMappingKey != p.ApiMappingKey {
		p.ApiMappingKey = k.ApiMappingKey
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApigatewayv2ApiMapping_DomainName(k *Apigatewayv2ApiMappingParameters, p *Apigatewayv2ApiMappingParameters, md *plugin.MergeDescription) bool {
	if k.DomainName != p.DomainName {
		p.DomainName = k.DomainName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApigatewayv2ApiMapping_Id(k *Apigatewayv2ApiMappingParameters, p *Apigatewayv2ApiMappingParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApigatewayv2ApiMapping_Stage(k *Apigatewayv2ApiMappingParameters, p *Apigatewayv2ApiMappingParameters, md *plugin.MergeDescription) bool {
	if k.Stage != p.Stage {
		p.Stage = k.Stage
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApigatewayv2ApiMapping_ApiId(k *Apigatewayv2ApiMappingParameters, p *Apigatewayv2ApiMappingParameters, md *plugin.MergeDescription) bool {
	if k.ApiId != p.ApiId {
		p.ApiId = k.ApiId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}