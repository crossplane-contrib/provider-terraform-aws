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
	k := kube.(*IamOpenidConnectProvider)
	p := prov.(*IamOpenidConnectProvider)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeIamOpenidConnectProvider_ClientIdList(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamOpenidConnectProvider_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamOpenidConnectProvider_ThumbprintList(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamOpenidConnectProvider_Url(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeIamOpenidConnectProvider_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeIamOpenidConnectProvider_ClientIdList(k *IamOpenidConnectProviderParameters, p *IamOpenidConnectProviderParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(p.ClientIdList, p.ClientIdList) {
		p.ClientIdList = k.ClientIdList
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamOpenidConnectProvider_Id(k *IamOpenidConnectProviderParameters, p *IamOpenidConnectProviderParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeIamOpenidConnectProvider_ThumbprintList(k *IamOpenidConnectProviderParameters, p *IamOpenidConnectProviderParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(p.ThumbprintList, p.ThumbprintList) {
		p.ThumbprintList = k.ThumbprintList
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeIamOpenidConnectProvider_Url(k *IamOpenidConnectProviderParameters, p *IamOpenidConnectProviderParameters, md *plugin.MergeDescription) bool {
	if k.Url != p.Url {
		p.Url = k.Url
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeIamOpenidConnectProvider_Arn(k *IamOpenidConnectProviderObservation, p *IamOpenidConnectProviderObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}