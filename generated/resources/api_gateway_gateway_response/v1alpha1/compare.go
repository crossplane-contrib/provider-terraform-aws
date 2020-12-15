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
	k := kube.(*ApiGatewayGatewayResponse)
	p := prov.(*ApiGatewayGatewayResponse)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeApiGatewayGatewayResponse_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayGatewayResponse_ResponseParameters(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayGatewayResponse_ResponseTemplates(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayGatewayResponse_ResponseType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayGatewayResponse_RestApiId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayGatewayResponse_StatusCode(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeApiGatewayGatewayResponse_Id(k *ApiGatewayGatewayResponseParameters, p *ApiGatewayGatewayResponseParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeApiGatewayGatewayResponse_ResponseParameters(k *ApiGatewayGatewayResponseParameters, p *ApiGatewayGatewayResponseParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(p.ResponseParameters, p.ResponseParameters) {
		p.ResponseParameters = k.ResponseParameters
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeApiGatewayGatewayResponse_ResponseTemplates(k *ApiGatewayGatewayResponseParameters, p *ApiGatewayGatewayResponseParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(p.ResponseTemplates, p.ResponseTemplates) {
		p.ResponseTemplates = k.ResponseTemplates
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayGatewayResponse_ResponseType(k *ApiGatewayGatewayResponseParameters, p *ApiGatewayGatewayResponseParameters, md *plugin.MergeDescription) bool {
	if k.ResponseType != p.ResponseType {
		p.ResponseType = k.ResponseType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayGatewayResponse_RestApiId(k *ApiGatewayGatewayResponseParameters, p *ApiGatewayGatewayResponseParameters, md *plugin.MergeDescription) bool {
	if k.RestApiId != p.RestApiId {
		p.RestApiId = k.RestApiId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayGatewayResponse_StatusCode(k *ApiGatewayGatewayResponseParameters, p *ApiGatewayGatewayResponseParameters, md *plugin.MergeDescription) bool {
	if k.StatusCode != p.StatusCode {
		p.StatusCode = k.StatusCode
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}