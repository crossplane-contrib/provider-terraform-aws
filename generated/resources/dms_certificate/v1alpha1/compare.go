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
	k := kube.(*DmsCertificate)
	p := prov.(*DmsCertificate)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDmsCertificate_CertificatePem(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsCertificate_CertificateWallet(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsCertificate_CertificateId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDmsCertificate_CertificateArn(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeDmsCertificate_CertificatePem(k *DmsCertificateParameters, p *DmsCertificateParameters, md *plugin.MergeDescription) bool {
	if k.CertificatePem != p.CertificatePem {
		p.CertificatePem = k.CertificatePem
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDmsCertificate_CertificateWallet(k *DmsCertificateParameters, p *DmsCertificateParameters, md *plugin.MergeDescription) bool {
	if k.CertificateWallet != p.CertificateWallet {
		p.CertificateWallet = k.CertificateWallet
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDmsCertificate_CertificateId(k *DmsCertificateParameters, p *DmsCertificateParameters, md *plugin.MergeDescription) bool {
	if k.CertificateId != p.CertificateId {
		p.CertificateId = k.CertificateId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDmsCertificate_CertificateArn(k *DmsCertificateObservation, p *DmsCertificateObservation, md *plugin.MergeDescription) bool {
	if k.CertificateArn != p.CertificateArn {
		k.CertificateArn = p.CertificateArn
		md.StatusUpdated = true
		return true
	}
	return false
}