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

	"github.com/zclconf/go-cty/cty"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*LicensemanagerLicenseConfiguration)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a LicensemanagerLicenseConfiguration.")
	}
	return EncodeLicensemanagerLicenseConfiguration(*r), nil
}

func EncodeLicensemanagerLicenseConfiguration(r LicensemanagerLicenseConfiguration) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeLicensemanagerLicenseConfiguration_Tags(r.Spec.ForProvider, ctyVal)
	EncodeLicensemanagerLicenseConfiguration_Description(r.Spec.ForProvider, ctyVal)
	EncodeLicensemanagerLicenseConfiguration_Id(r.Spec.ForProvider, ctyVal)
	EncodeLicensemanagerLicenseConfiguration_LicenseCount(r.Spec.ForProvider, ctyVal)
	EncodeLicensemanagerLicenseConfiguration_LicenseCountHardLimit(r.Spec.ForProvider, ctyVal)
	EncodeLicensemanagerLicenseConfiguration_LicenseCountingType(r.Spec.ForProvider, ctyVal)
	EncodeLicensemanagerLicenseConfiguration_LicenseRules(r.Spec.ForProvider, ctyVal)
	EncodeLicensemanagerLicenseConfiguration_Name(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeLicensemanagerLicenseConfiguration_Tags(p LicensemanagerLicenseConfigurationParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeLicensemanagerLicenseConfiguration_Description(p LicensemanagerLicenseConfigurationParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeLicensemanagerLicenseConfiguration_Id(p LicensemanagerLicenseConfigurationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLicensemanagerLicenseConfiguration_LicenseCount(p LicensemanagerLicenseConfigurationParameters, vals map[string]cty.Value) {
	vals["license_count"] = cty.NumberIntVal(p.LicenseCount)
}

func EncodeLicensemanagerLicenseConfiguration_LicenseCountHardLimit(p LicensemanagerLicenseConfigurationParameters, vals map[string]cty.Value) {
	vals["license_count_hard_limit"] = cty.BoolVal(p.LicenseCountHardLimit)
}

func EncodeLicensemanagerLicenseConfiguration_LicenseCountingType(p LicensemanagerLicenseConfigurationParameters, vals map[string]cty.Value) {
	vals["license_counting_type"] = cty.StringVal(p.LicenseCountingType)
}

func EncodeLicensemanagerLicenseConfiguration_LicenseRules(p LicensemanagerLicenseConfigurationParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.LicenseRules {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["license_rules"] = cty.ListVal(colVals)
}

func EncodeLicensemanagerLicenseConfiguration_Name(p LicensemanagerLicenseConfigurationParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}