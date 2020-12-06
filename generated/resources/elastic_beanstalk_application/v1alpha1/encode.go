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
	"github.com/zclconf/go-cty/cty"
)

func EncodeElasticBeanstalkApplication(r ElasticBeanstalkApplication) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeElasticBeanstalkApplication_Name(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkApplication_Tags(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkApplication_Description(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkApplication_Id(r.Spec.ForProvider, ctyVal)
	EncodeElasticBeanstalkApplication_AppversionLifecycle(r.Spec.ForProvider.AppversionLifecycle, ctyVal)
	EncodeElasticBeanstalkApplication_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeElasticBeanstalkApplication_Name(p ElasticBeanstalkApplicationParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeElasticBeanstalkApplication_Tags(p ElasticBeanstalkApplicationParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeElasticBeanstalkApplication_Description(p ElasticBeanstalkApplicationParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeElasticBeanstalkApplication_Id(p ElasticBeanstalkApplicationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeElasticBeanstalkApplication_AppversionLifecycle(p AppversionLifecycle, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeElasticBeanstalkApplication_AppversionLifecycle_MaxCount(p, ctyVal)
	EncodeElasticBeanstalkApplication_AppversionLifecycle_ServiceRole(p, ctyVal)
	EncodeElasticBeanstalkApplication_AppversionLifecycle_DeleteSourceFromS3(p, ctyVal)
	EncodeElasticBeanstalkApplication_AppversionLifecycle_MaxAgeInDays(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["appversion_lifecycle"] = cty.ListVal(valsForCollection)
}

func EncodeElasticBeanstalkApplication_AppversionLifecycle_MaxCount(p AppversionLifecycle, vals map[string]cty.Value) {
	vals["max_count"] = cty.NumberIntVal(p.MaxCount)
}

func EncodeElasticBeanstalkApplication_AppversionLifecycle_ServiceRole(p AppversionLifecycle, vals map[string]cty.Value) {
	vals["service_role"] = cty.StringVal(p.ServiceRole)
}

func EncodeElasticBeanstalkApplication_AppversionLifecycle_DeleteSourceFromS3(p AppversionLifecycle, vals map[string]cty.Value) {
	vals["delete_source_from_s3"] = cty.BoolVal(p.DeleteSourceFromS3)
}

func EncodeElasticBeanstalkApplication_AppversionLifecycle_MaxAgeInDays(p AppversionLifecycle, vals map[string]cty.Value) {
	vals["max_age_in_days"] = cty.NumberIntVal(p.MaxAgeInDays)
}

func EncodeElasticBeanstalkApplication_Arn(p ElasticBeanstalkApplicationObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}