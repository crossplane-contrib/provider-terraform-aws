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
	k := kube.(*ApiGatewayResource)
	p := prov.(*ApiGatewayResource)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeApiGatewayResource_RestApiId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayResource_ParentId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayResource_PathPart(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeApiGatewayResource_Path(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeApiGatewayResource_RestApiId(k *ApiGatewayResourceParameters, p *ApiGatewayResourceParameters, md *plugin.MergeDescription) bool {
	if k.RestApiId != p.RestApiId {
		p.RestApiId = k.RestApiId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayResource_ParentId(k *ApiGatewayResourceParameters, p *ApiGatewayResourceParameters, md *plugin.MergeDescription) bool {
	if k.ParentId != p.ParentId {
		p.ParentId = k.ParentId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeApiGatewayResource_PathPart(k *ApiGatewayResourceParameters, p *ApiGatewayResourceParameters, md *plugin.MergeDescription) bool {
	if k.PathPart != p.PathPart {
		p.PathPart = k.PathPart
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeApiGatewayResource_Path(k *ApiGatewayResourceObservation, p *ApiGatewayResourceObservation, md *plugin.MergeDescription) bool {
	if k.Path != p.Path {
		k.Path = p.Path
		md.StatusUpdated = true
		return true
	}
	return false
}