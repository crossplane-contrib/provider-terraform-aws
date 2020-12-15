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
	k := kube.(*VpcEndpointConnectionNotification)
	p := prov.(*VpcEndpointConnectionNotification)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeVpcEndpointConnectionNotification_VpcEndpointId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcEndpointConnectionNotification_VpcEndpointServiceId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcEndpointConnectionNotification_ConnectionEvents(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcEndpointConnectionNotification_ConnectionNotificationArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcEndpointConnectionNotification_NotificationType(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVpcEndpointConnectionNotification_State(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeVpcEndpointConnectionNotification_VpcEndpointId(k *VpcEndpointConnectionNotificationParameters, p *VpcEndpointConnectionNotificationParameters, md *plugin.MergeDescription) bool {
	if k.VpcEndpointId != p.VpcEndpointId {
		p.VpcEndpointId = k.VpcEndpointId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVpcEndpointConnectionNotification_VpcEndpointServiceId(k *VpcEndpointConnectionNotificationParameters, p *VpcEndpointConnectionNotificationParameters, md *plugin.MergeDescription) bool {
	if k.VpcEndpointServiceId != p.VpcEndpointServiceId {
		p.VpcEndpointServiceId = k.VpcEndpointServiceId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVpcEndpointConnectionNotification_ConnectionEvents(k *VpcEndpointConnectionNotificationParameters, p *VpcEndpointConnectionNotificationParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.ConnectionEvents, p.ConnectionEvents) {
		p.ConnectionEvents = k.ConnectionEvents
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVpcEndpointConnectionNotification_ConnectionNotificationArn(k *VpcEndpointConnectionNotificationParameters, p *VpcEndpointConnectionNotificationParameters, md *plugin.MergeDescription) bool {
	if k.ConnectionNotificationArn != p.ConnectionNotificationArn {
		p.ConnectionNotificationArn = k.ConnectionNotificationArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVpcEndpointConnectionNotification_NotificationType(k *VpcEndpointConnectionNotificationObservation, p *VpcEndpointConnectionNotificationObservation, md *plugin.MergeDescription) bool {
	if k.NotificationType != p.NotificationType {
		k.NotificationType = p.NotificationType
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVpcEndpointConnectionNotification_State(k *VpcEndpointConnectionNotificationObservation, p *VpcEndpointConnectionNotificationObservation, md *plugin.MergeDescription) bool {
	if k.State != p.State {
		k.State = p.State
		md.StatusUpdated = true
		return true
	}
	return false
}