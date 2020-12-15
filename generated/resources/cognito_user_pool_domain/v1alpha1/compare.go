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
	k := kube.(*CognitoUserPoolDomain)
	p := prov.(*CognitoUserPoolDomain)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeCognitoUserPoolDomain_CertificateArn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCognitoUserPoolDomain_Domain(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCognitoUserPoolDomain_Id(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCognitoUserPoolDomain_UserPoolId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCognitoUserPoolDomain_Version(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCognitoUserPoolDomain_AwsAccountId(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCognitoUserPoolDomain_CloudfrontDistributionArn(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeCognitoUserPoolDomain_S3Bucket(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeCognitoUserPoolDomain_CertificateArn(k *CognitoUserPoolDomainParameters, p *CognitoUserPoolDomainParameters, md *plugin.MergeDescription) bool {
	if k.CertificateArn != p.CertificateArn {
		p.CertificateArn = k.CertificateArn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCognitoUserPoolDomain_Domain(k *CognitoUserPoolDomainParameters, p *CognitoUserPoolDomainParameters, md *plugin.MergeDescription) bool {
	if k.Domain != p.Domain {
		p.Domain = k.Domain
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCognitoUserPoolDomain_Id(k *CognitoUserPoolDomainParameters, p *CognitoUserPoolDomainParameters, md *plugin.MergeDescription) bool {
	if k.Id != p.Id {
		p.Id = k.Id
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeCognitoUserPoolDomain_UserPoolId(k *CognitoUserPoolDomainParameters, p *CognitoUserPoolDomainParameters, md *plugin.MergeDescription) bool {
	if k.UserPoolId != p.UserPoolId {
		p.UserPoolId = k.UserPoolId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCognitoUserPoolDomain_Version(k *CognitoUserPoolDomainObservation, p *CognitoUserPoolDomainObservation, md *plugin.MergeDescription) bool {
	if k.Version != p.Version {
		k.Version = p.Version
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCognitoUserPoolDomain_AwsAccountId(k *CognitoUserPoolDomainObservation, p *CognitoUserPoolDomainObservation, md *plugin.MergeDescription) bool {
	if k.AwsAccountId != p.AwsAccountId {
		k.AwsAccountId = p.AwsAccountId
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCognitoUserPoolDomain_CloudfrontDistributionArn(k *CognitoUserPoolDomainObservation, p *CognitoUserPoolDomainObservation, md *plugin.MergeDescription) bool {
	if k.CloudfrontDistributionArn != p.CloudfrontDistributionArn {
		k.CloudfrontDistributionArn = p.CloudfrontDistributionArn
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeCognitoUserPoolDomain_S3Bucket(k *CognitoUserPoolDomainObservation, p *CognitoUserPoolDomainObservation, md *plugin.MergeDescription) bool {
	if k.S3Bucket != p.S3Bucket {
		k.S3Bucket = p.S3Bucket
		md.StatusUpdated = true
		return true
	}
	return false
}