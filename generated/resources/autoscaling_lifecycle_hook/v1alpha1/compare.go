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
	k := kube.(*AutoscalingLifecycleHook)
	p := prov.(*AutoscalingLifecycleHook)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeAutoscalingLifecycleHook_LifecycleTransition(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAutoscalingLifecycleHook_RoleArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAutoscalingLifecycleHook_AutoscalingGroupName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAutoscalingLifecycleHook_DefaultResult(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAutoscalingLifecycleHook_HeartbeatTimeout(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAutoscalingLifecycleHook_NotificationTargetArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAutoscalingLifecycleHook_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeAutoscalingLifecycleHook_NotificationMetadata(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeAutoscalingLifecycleHook_LifecycleTransition(k *AutoscalingLifecycleHookParameters, p *AutoscalingLifecycleHookParameters, md *plugin.MergeDescription) bool {
	if k.LifecycleTransition != p.LifecycleTransition {
		p.LifecycleTransition = k.LifecycleTransition
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAutoscalingLifecycleHook_RoleArn(k *AutoscalingLifecycleHookParameters, p *AutoscalingLifecycleHookParameters, md *plugin.MergeDescription) bool {
	if k.RoleArn != p.RoleArn {
		p.RoleArn = k.RoleArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAutoscalingLifecycleHook_AutoscalingGroupName(k *AutoscalingLifecycleHookParameters, p *AutoscalingLifecycleHookParameters, md *plugin.MergeDescription) bool {
	if k.AutoscalingGroupName != p.AutoscalingGroupName {
		p.AutoscalingGroupName = k.AutoscalingGroupName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAutoscalingLifecycleHook_DefaultResult(k *AutoscalingLifecycleHookParameters, p *AutoscalingLifecycleHookParameters, md *plugin.MergeDescription) bool {
	if k.DefaultResult != p.DefaultResult {
		p.DefaultResult = k.DefaultResult
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAutoscalingLifecycleHook_HeartbeatTimeout(k *AutoscalingLifecycleHookParameters, p *AutoscalingLifecycleHookParameters, md *plugin.MergeDescription) bool {
	if k.HeartbeatTimeout != p.HeartbeatTimeout {
		p.HeartbeatTimeout = k.HeartbeatTimeout
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAutoscalingLifecycleHook_NotificationTargetArn(k *AutoscalingLifecycleHookParameters, p *AutoscalingLifecycleHookParameters, md *plugin.MergeDescription) bool {
	if k.NotificationTargetArn != p.NotificationTargetArn {
		p.NotificationTargetArn = k.NotificationTargetArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAutoscalingLifecycleHook_Name(k *AutoscalingLifecycleHookParameters, p *AutoscalingLifecycleHookParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeAutoscalingLifecycleHook_NotificationMetadata(k *AutoscalingLifecycleHookParameters, p *AutoscalingLifecycleHookParameters, md *plugin.MergeDescription) bool {
	if k.NotificationMetadata != p.NotificationMetadata {
		p.NotificationMetadata = k.NotificationMetadata
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}