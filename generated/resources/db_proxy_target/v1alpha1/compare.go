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
	k := kube.(*DbProxyTarget)
	p := prov.(*DbProxyTarget)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDbProxyTarget_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbProxyTarget_TargetGroupName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbProxyTarget_DbClusterIdentifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbProxyTarget_DbInstanceIdentifier(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbProxyTarget_DbProxyName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbProxyTarget_Type(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbProxyTarget_Endpoint(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbProxyTarget_Port(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbProxyTarget_TargetArn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbProxyTarget_TrackedClusterId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDbProxyTarget_RdsResourceId(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeDbProxyTarget_Id(k *DbProxyTargetParameters, p *DbProxyTargetParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDbProxyTarget_TargetGroupName(k *DbProxyTargetParameters, p *DbProxyTargetParameters, md *plugin.MergeDescription) bool {
	if k.TargetGroupName != p.TargetGroupName {
		p.TargetGroupName = k.TargetGroupName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDbProxyTarget_DbClusterIdentifier(k *DbProxyTargetParameters, p *DbProxyTargetParameters, md *plugin.MergeDescription) bool {
	if k.DbClusterIdentifier != p.DbClusterIdentifier {
		p.DbClusterIdentifier = k.DbClusterIdentifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDbProxyTarget_DbInstanceIdentifier(k *DbProxyTargetParameters, p *DbProxyTargetParameters, md *plugin.MergeDescription) bool {
	if k.DbInstanceIdentifier != p.DbInstanceIdentifier {
		p.DbInstanceIdentifier = k.DbInstanceIdentifier
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDbProxyTarget_DbProxyName(k *DbProxyTargetParameters, p *DbProxyTargetParameters, md *plugin.MergeDescription) bool {
	if k.DbProxyName != p.DbProxyName {
		p.DbProxyName = k.DbProxyName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDbProxyTarget_Type(k *DbProxyTargetObservation, p *DbProxyTargetObservation, md *plugin.MergeDescription) bool {
	if k.Type != p.Type {
		k.Type = p.Type
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDbProxyTarget_Endpoint(k *DbProxyTargetObservation, p *DbProxyTargetObservation, md *plugin.MergeDescription) bool {
	if k.Endpoint != p.Endpoint {
		k.Endpoint = p.Endpoint
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDbProxyTarget_Port(k *DbProxyTargetObservation, p *DbProxyTargetObservation, md *plugin.MergeDescription) bool {
	if k.Port != p.Port {
		k.Port = p.Port
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDbProxyTarget_TargetArn(k *DbProxyTargetObservation, p *DbProxyTargetObservation, md *plugin.MergeDescription) bool {
	if k.TargetArn != p.TargetArn {
		k.TargetArn = p.TargetArn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDbProxyTarget_TrackedClusterId(k *DbProxyTargetObservation, p *DbProxyTargetObservation, md *plugin.MergeDescription) bool {
	if k.TrackedClusterId != p.TrackedClusterId {
		k.TrackedClusterId = p.TrackedClusterId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDbProxyTarget_RdsResourceId(k *DbProxyTargetObservation, p *DbProxyTargetObservation, md *plugin.MergeDescription) bool {
	if k.RdsResourceId != p.RdsResourceId {
		k.RdsResourceId = p.RdsResourceId
		md.StatusUpdated = true
		return true
	}
	return false
}