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
	k := kube.(*PinpointBaiduChannel)
	p := prov.(*PinpointBaiduChannel)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergePinpointBaiduChannel_ApiKey(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergePinpointBaiduChannel_ApplicationId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergePinpointBaiduChannel_Enabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergePinpointBaiduChannel_SecretKey(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergePinpointBaiduChannel_ApiKey(k *PinpointBaiduChannelParameters, p *PinpointBaiduChannelParameters, md *plugin.MergeDescription) bool {
	if k.ApiKey != p.ApiKey {
		p.ApiKey = k.ApiKey
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergePinpointBaiduChannel_ApplicationId(k *PinpointBaiduChannelParameters, p *PinpointBaiduChannelParameters, md *plugin.MergeDescription) bool {
	if k.ApplicationId != p.ApplicationId {
		p.ApplicationId = k.ApplicationId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergePinpointBaiduChannel_Enabled(k *PinpointBaiduChannelParameters, p *PinpointBaiduChannelParameters, md *plugin.MergeDescription) bool {
	if k.Enabled != p.Enabled {
		p.Enabled = k.Enabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergePinpointBaiduChannel_SecretKey(k *PinpointBaiduChannelParameters, p *PinpointBaiduChannelParameters, md *plugin.MergeDescription) bool {
	if k.SecretKey != p.SecretKey {
		p.SecretKey = k.SecretKey
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}