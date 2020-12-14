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
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*AcmCertificateValidation)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeAcmCertificateValidation(r, ctyValue)
}

func DecodeAcmCertificateValidation(prev *AcmCertificateValidation, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeAcmCertificateValidation_Id(&new.Spec.ForProvider, valMap)
	DecodeAcmCertificateValidation_ValidationRecordFqdns(&new.Spec.ForProvider, valMap)
	DecodeAcmCertificateValidation_CertificateArn(&new.Spec.ForProvider, valMap)
	DecodeAcmCertificateValidation_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)

	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeAcmCertificateValidation_Id(p *AcmCertificateValidationParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeAcmCertificateValidation_ValidationRecordFqdns(p *AcmCertificateValidationParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["validation_record_fqdns"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.ValidationRecordFqdns = goVals
}

func DecodeAcmCertificateValidation_CertificateArn(p *AcmCertificateValidationParameters, vals map[string]cty.Value) {
	p.CertificateArn = ctwhy.ValueAsString(vals["certificate_arn"])
}

func DecodeAcmCertificateValidation_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeAcmCertificateValidation_Timeouts_Create(p, valMap)
}

func DecodeAcmCertificateValidation_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}