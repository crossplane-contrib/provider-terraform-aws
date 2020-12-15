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
	k := kube.(*PinpointApnsVoipSandboxChannel)
	p := prov.(*PinpointApnsVoipSandboxChannel)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergePinpointApnsVoipSandboxChannel_BundleId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergePinpointApnsVoipSandboxChannel_Enabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergePinpointApnsVoipSandboxChannel_TokenKey(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergePinpointApnsVoipSandboxChannel_TokenKeyId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergePinpointApnsVoipSandboxChannel_TeamId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergePinpointApnsVoipSandboxChannel_ApplicationId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergePinpointApnsVoipSandboxChannel_Certificate(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergePinpointApnsVoipSandboxChannel_DefaultAuthenticationMethod(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergePinpointApnsVoipSandboxChannel_PrivateKey(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergePinpointApnsVoipSandboxChannel_BundleId(k *PinpointApnsVoipSandboxChannelParameters, p *PinpointApnsVoipSandboxChannelParameters, md *plugin.MergeDescription) bool {
	if k.BundleId != p.BundleId {
		p.BundleId = k.BundleId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergePinpointApnsVoipSandboxChannel_Enabled(k *PinpointApnsVoipSandboxChannelParameters, p *PinpointApnsVoipSandboxChannelParameters, md *plugin.MergeDescription) bool {
	if k.Enabled != p.Enabled {
		p.Enabled = k.Enabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergePinpointApnsVoipSandboxChannel_TokenKey(k *PinpointApnsVoipSandboxChannelParameters, p *PinpointApnsVoipSandboxChannelParameters, md *plugin.MergeDescription) bool {
	if k.TokenKey != p.TokenKey {
		p.TokenKey = k.TokenKey
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergePinpointApnsVoipSandboxChannel_TokenKeyId(k *PinpointApnsVoipSandboxChannelParameters, p *PinpointApnsVoipSandboxChannelParameters, md *plugin.MergeDescription) bool {
	if k.TokenKeyId != p.TokenKeyId {
		p.TokenKeyId = k.TokenKeyId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergePinpointApnsVoipSandboxChannel_TeamId(k *PinpointApnsVoipSandboxChannelParameters, p *PinpointApnsVoipSandboxChannelParameters, md *plugin.MergeDescription) bool {
	if k.TeamId != p.TeamId {
		p.TeamId = k.TeamId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergePinpointApnsVoipSandboxChannel_ApplicationId(k *PinpointApnsVoipSandboxChannelParameters, p *PinpointApnsVoipSandboxChannelParameters, md *plugin.MergeDescription) bool {
	if k.ApplicationId != p.ApplicationId {
		p.ApplicationId = k.ApplicationId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergePinpointApnsVoipSandboxChannel_Certificate(k *PinpointApnsVoipSandboxChannelParameters, p *PinpointApnsVoipSandboxChannelParameters, md *plugin.MergeDescription) bool {
	if k.Certificate != p.Certificate {
		p.Certificate = k.Certificate
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergePinpointApnsVoipSandboxChannel_DefaultAuthenticationMethod(k *PinpointApnsVoipSandboxChannelParameters, p *PinpointApnsVoipSandboxChannelParameters, md *plugin.MergeDescription) bool {
	if k.DefaultAuthenticationMethod != p.DefaultAuthenticationMethod {
		p.DefaultAuthenticationMethod = k.DefaultAuthenticationMethod
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergePinpointApnsVoipSandboxChannel_PrivateKey(k *PinpointApnsVoipSandboxChannelParameters, p *PinpointApnsVoipSandboxChannelParameters, md *plugin.MergeDescription) bool {
	if k.PrivateKey != p.PrivateKey {
		p.PrivateKey = k.PrivateKey
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}