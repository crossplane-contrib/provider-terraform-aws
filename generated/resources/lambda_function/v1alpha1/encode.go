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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*LambdaFunction)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a LambdaFunction.")
	}
	return EncodeLambdaFunction(*r), nil
}

func EncodeLambdaFunction(r LambdaFunction) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeLambdaFunction_MemorySize(r.Spec.ForProvider, ctyVal)
	EncodeLambdaFunction_S3Bucket(r.Spec.ForProvider, ctyVal)
	EncodeLambdaFunction_Tags(r.Spec.ForProvider, ctyVal)
	EncodeLambdaFunction_Filename(r.Spec.ForProvider, ctyVal)
	EncodeLambdaFunction_Id(r.Spec.ForProvider, ctyVal)
	EncodeLambdaFunction_ReservedConcurrentExecutions(r.Spec.ForProvider, ctyVal)
	EncodeLambdaFunction_S3ObjectVersion(r.Spec.ForProvider, ctyVal)
	EncodeLambdaFunction_Description(r.Spec.ForProvider, ctyVal)
	EncodeLambdaFunction_KmsKeyArn(r.Spec.ForProvider, ctyVal)
	EncodeLambdaFunction_S3Key(r.Spec.ForProvider, ctyVal)
	EncodeLambdaFunction_SourceCodeHash(r.Spec.ForProvider, ctyVal)
	EncodeLambdaFunction_Timeout(r.Spec.ForProvider, ctyVal)
	EncodeLambdaFunction_Runtime(r.Spec.ForProvider, ctyVal)
	EncodeLambdaFunction_FunctionName(r.Spec.ForProvider, ctyVal)
	EncodeLambdaFunction_Handler(r.Spec.ForProvider, ctyVal)
	EncodeLambdaFunction_Layers(r.Spec.ForProvider, ctyVal)
	EncodeLambdaFunction_Publish(r.Spec.ForProvider, ctyVal)
	EncodeLambdaFunction_Role(r.Spec.ForProvider, ctyVal)
	EncodeLambdaFunction_DeadLetterConfig(r.Spec.ForProvider.DeadLetterConfig, ctyVal)
	EncodeLambdaFunction_Environment(r.Spec.ForProvider.Environment, ctyVal)
	EncodeLambdaFunction_FileSystemConfig(r.Spec.ForProvider.FileSystemConfig, ctyVal)
	EncodeLambdaFunction_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeLambdaFunction_TracingConfig(r.Spec.ForProvider.TracingConfig, ctyVal)
	EncodeLambdaFunction_VpcConfig(r.Spec.ForProvider.VpcConfig, ctyVal)
	EncodeLambdaFunction_InvokeArn(r.Status.AtProvider, ctyVal)
	EncodeLambdaFunction_Version(r.Status.AtProvider, ctyVal)
	EncodeLambdaFunction_LastModified(r.Status.AtProvider, ctyVal)
	EncodeLambdaFunction_Arn(r.Status.AtProvider, ctyVal)
	EncodeLambdaFunction_SourceCodeSize(r.Status.AtProvider, ctyVal)
	EncodeLambdaFunction_QualifiedArn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeLambdaFunction_MemorySize(p LambdaFunctionParameters, vals map[string]cty.Value) {
	vals["memory_size"] = cty.NumberIntVal(p.MemorySize)
}

func EncodeLambdaFunction_S3Bucket(p LambdaFunctionParameters, vals map[string]cty.Value) {
	vals["s3_bucket"] = cty.StringVal(p.S3Bucket)
}

func EncodeLambdaFunction_Tags(p LambdaFunctionParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeLambdaFunction_Filename(p LambdaFunctionParameters, vals map[string]cty.Value) {
	vals["filename"] = cty.StringVal(p.Filename)
}

func EncodeLambdaFunction_Id(p LambdaFunctionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLambdaFunction_ReservedConcurrentExecutions(p LambdaFunctionParameters, vals map[string]cty.Value) {
	vals["reserved_concurrent_executions"] = cty.NumberIntVal(p.ReservedConcurrentExecutions)
}

func EncodeLambdaFunction_S3ObjectVersion(p LambdaFunctionParameters, vals map[string]cty.Value) {
	vals["s3_object_version"] = cty.StringVal(p.S3ObjectVersion)
}

func EncodeLambdaFunction_Description(p LambdaFunctionParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeLambdaFunction_KmsKeyArn(p LambdaFunctionParameters, vals map[string]cty.Value) {
	vals["kms_key_arn"] = cty.StringVal(p.KmsKeyArn)
}

func EncodeLambdaFunction_S3Key(p LambdaFunctionParameters, vals map[string]cty.Value) {
	vals["s3_key"] = cty.StringVal(p.S3Key)
}

func EncodeLambdaFunction_SourceCodeHash(p LambdaFunctionParameters, vals map[string]cty.Value) {
	vals["source_code_hash"] = cty.StringVal(p.SourceCodeHash)
}

func EncodeLambdaFunction_Timeout(p LambdaFunctionParameters, vals map[string]cty.Value) {
	vals["timeout"] = cty.NumberIntVal(p.Timeout)
}

func EncodeLambdaFunction_Runtime(p LambdaFunctionParameters, vals map[string]cty.Value) {
	vals["runtime"] = cty.StringVal(p.Runtime)
}

func EncodeLambdaFunction_FunctionName(p LambdaFunctionParameters, vals map[string]cty.Value) {
	vals["function_name"] = cty.StringVal(p.FunctionName)
}

func EncodeLambdaFunction_Handler(p LambdaFunctionParameters, vals map[string]cty.Value) {
	vals["handler"] = cty.StringVal(p.Handler)
}

func EncodeLambdaFunction_Layers(p LambdaFunctionParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Layers {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["layers"] = cty.ListVal(colVals)
}

func EncodeLambdaFunction_Publish(p LambdaFunctionParameters, vals map[string]cty.Value) {
	vals["publish"] = cty.BoolVal(p.Publish)
}

func EncodeLambdaFunction_Role(p LambdaFunctionParameters, vals map[string]cty.Value) {
	vals["role"] = cty.StringVal(p.Role)
}

func EncodeLambdaFunction_DeadLetterConfig(p DeadLetterConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLambdaFunction_DeadLetterConfig_TargetArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["dead_letter_config"] = cty.ListVal(valsForCollection)
}

func EncodeLambdaFunction_DeadLetterConfig_TargetArn(p DeadLetterConfig, vals map[string]cty.Value) {
	vals["target_arn"] = cty.StringVal(p.TargetArn)
}

func EncodeLambdaFunction_Environment(p Environment, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLambdaFunction_Environment_Variables(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["environment"] = cty.ListVal(valsForCollection)
}

func EncodeLambdaFunction_Environment_Variables(p Environment, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Variables {
		mVals[key] = cty.StringVal(value)
	}
	vals["variables"] = cty.MapVal(mVals)
}

func EncodeLambdaFunction_FileSystemConfig(p FileSystemConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLambdaFunction_FileSystemConfig_Arn(p, ctyVal)
	EncodeLambdaFunction_FileSystemConfig_LocalMountPath(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["file_system_config"] = cty.ListVal(valsForCollection)
}

func EncodeLambdaFunction_FileSystemConfig_Arn(p FileSystemConfig, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeLambdaFunction_FileSystemConfig_LocalMountPath(p FileSystemConfig, vals map[string]cty.Value) {
	vals["local_mount_path"] = cty.StringVal(p.LocalMountPath)
}

func EncodeLambdaFunction_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeLambdaFunction_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeLambdaFunction_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeLambdaFunction_TracingConfig(p TracingConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLambdaFunction_TracingConfig_Mode(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["tracing_config"] = cty.ListVal(valsForCollection)
}

func EncodeLambdaFunction_TracingConfig_Mode(p TracingConfig, vals map[string]cty.Value) {
	vals["mode"] = cty.StringVal(p.Mode)
}

func EncodeLambdaFunction_VpcConfig(p VpcConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLambdaFunction_VpcConfig_SecurityGroupIds(p, ctyVal)
	EncodeLambdaFunction_VpcConfig_SubnetIds(p, ctyVal)
	EncodeLambdaFunction_VpcConfig_VpcId(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["vpc_config"] = cty.ListVal(valsForCollection)
}

func EncodeLambdaFunction_VpcConfig_SecurityGroupIds(p VpcConfig, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_group_ids"] = cty.SetVal(colVals)
}

func EncodeLambdaFunction_VpcConfig_SubnetIds(p VpcConfig, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.SetVal(colVals)
}

func EncodeLambdaFunction_VpcConfig_VpcId(p VpcConfig, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeLambdaFunction_InvokeArn(p LambdaFunctionObservation, vals map[string]cty.Value) {
	vals["invoke_arn"] = cty.StringVal(p.InvokeArn)
}

func EncodeLambdaFunction_Version(p LambdaFunctionObservation, vals map[string]cty.Value) {
	vals["version"] = cty.StringVal(p.Version)
}

func EncodeLambdaFunction_LastModified(p LambdaFunctionObservation, vals map[string]cty.Value) {
	vals["last_modified"] = cty.StringVal(p.LastModified)
}

func EncodeLambdaFunction_Arn(p LambdaFunctionObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeLambdaFunction_SourceCodeSize(p LambdaFunctionObservation, vals map[string]cty.Value) {
	vals["source_code_size"] = cty.NumberIntVal(p.SourceCodeSize)
}

func EncodeLambdaFunction_QualifiedArn(p LambdaFunctionObservation, vals map[string]cty.Value) {
	vals["qualified_arn"] = cty.StringVal(p.QualifiedArn)
}