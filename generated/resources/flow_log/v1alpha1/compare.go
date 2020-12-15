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
	k := kube.(*FlowLog)
	p := prov.(*FlowLog)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeFlowLog_EniId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFlowLog_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFlowLog_LogDestination(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFlowLog_LogFormat(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFlowLog_TrafficType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFlowLog_LogDestinationType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFlowLog_LogGroupName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFlowLog_MaxAggregationInterval(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFlowLog_SubnetId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFlowLog_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFlowLog_VpcId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFlowLog_IamRoleArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFlowLog_Arn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeFlowLog_EniId(k *FlowLogParameters, p *FlowLogParameters, md *plugin.MergeDescription) bool {
	if k.EniId != p.EniId {
		p.EniId = k.EniId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFlowLog_Id(k *FlowLogParameters, p *FlowLogParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFlowLog_LogDestination(k *FlowLogParameters, p *FlowLogParameters, md *plugin.MergeDescription) bool {
	if k.LogDestination != p.LogDestination {
		p.LogDestination = k.LogDestination
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFlowLog_LogFormat(k *FlowLogParameters, p *FlowLogParameters, md *plugin.MergeDescription) bool {
	if k.LogFormat != p.LogFormat {
		p.LogFormat = k.LogFormat
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFlowLog_TrafficType(k *FlowLogParameters, p *FlowLogParameters, md *plugin.MergeDescription) bool {
	if k.TrafficType != p.TrafficType {
		p.TrafficType = k.TrafficType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFlowLog_LogDestinationType(k *FlowLogParameters, p *FlowLogParameters, md *plugin.MergeDescription) bool {
	if k.LogDestinationType != p.LogDestinationType {
		p.LogDestinationType = k.LogDestinationType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFlowLog_LogGroupName(k *FlowLogParameters, p *FlowLogParameters, md *plugin.MergeDescription) bool {
	if k.LogGroupName != p.LogGroupName {
		p.LogGroupName = k.LogGroupName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFlowLog_MaxAggregationInterval(k *FlowLogParameters, p *FlowLogParameters, md *plugin.MergeDescription) bool {
	if k.MaxAggregationInterval != p.MaxAggregationInterval {
		p.MaxAggregationInterval = k.MaxAggregationInterval
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFlowLog_SubnetId(k *FlowLogParameters, p *FlowLogParameters, md *plugin.MergeDescription) bool {
	if k.SubnetId != p.SubnetId {
		p.SubnetId = k.SubnetId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeFlowLog_Tags(k *FlowLogParameters, p *FlowLogParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(p.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFlowLog_VpcId(k *FlowLogParameters, p *FlowLogParameters, md *plugin.MergeDescription) bool {
	if k.VpcId != p.VpcId {
		p.VpcId = k.VpcId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFlowLog_IamRoleArn(k *FlowLogParameters, p *FlowLogParameters, md *plugin.MergeDescription) bool {
	if k.IamRoleArn != p.IamRoleArn {
		p.IamRoleArn = k.IamRoleArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeFlowLog_Arn(k *FlowLogObservation, p *FlowLogObservation, md *plugin.MergeDescription) bool {
	if k.Arn != p.Arn {
		k.Arn = p.Arn
		md.StatusUpdated = true
		return true
	}
	return false
}