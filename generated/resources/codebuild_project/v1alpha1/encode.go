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

package v1alpha1func EncodeCodebuildProject(r CodebuildProject) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeCodebuildProject_EncryptionKey(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildProject_Id(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildProject_ServiceRole(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildProject_BadgeEnabled(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildProject_Description(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildProject_Tags(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildProject_BuildTimeout(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildProject_Name(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildProject_QueuedTimeout(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildProject_SourceVersion(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildProject_Artifacts(r.Spec.ForProvider.Artifacts, ctyVal)
	EncodeCodebuildProject_Cache(r.Spec.ForProvider.Cache, ctyVal)
	EncodeCodebuildProject_Environment(r.Spec.ForProvider.Environment, ctyVal)
	EncodeCodebuildProject_LogsConfig(r.Spec.ForProvider.LogsConfig, ctyVal)
	EncodeCodebuildProject_SecondaryArtifacts(r.Spec.ForProvider.SecondaryArtifacts, ctyVal)
	EncodeCodebuildProject_SecondarySources(r.Spec.ForProvider.SecondarySources, ctyVal)
	EncodeCodebuildProject_Source(r.Spec.ForProvider.Source, ctyVal)
	EncodeCodebuildProject_VpcConfig(r.Spec.ForProvider.VpcConfig, ctyVal)
	EncodeCodebuildProject_Arn(r.Status.AtProvider, ctyVal)
	EncodeCodebuildProject_BadgeUrl(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeCodebuildProject_EncryptionKey(p *CodebuildProjectParameters, vals map[string]cty.Value) {
	vals["encryption_key"] = cty.StringVal(p.EncryptionKey)
}

func EncodeCodebuildProject_Id(p *CodebuildProjectParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCodebuildProject_ServiceRole(p *CodebuildProjectParameters, vals map[string]cty.Value) {
	vals["service_role"] = cty.StringVal(p.ServiceRole)
}

func EncodeCodebuildProject_BadgeEnabled(p *CodebuildProjectParameters, vals map[string]cty.Value) {
	vals["badge_enabled"] = cty.BoolVal(p.BadgeEnabled)
}

func EncodeCodebuildProject_Description(p *CodebuildProjectParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeCodebuildProject_Tags(p *CodebuildProjectParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeCodebuildProject_BuildTimeout(p *CodebuildProjectParameters, vals map[string]cty.Value) {
	vals["build_timeout"] = cty.IntVal(p.BuildTimeout)
}

func EncodeCodebuildProject_Name(p *CodebuildProjectParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCodebuildProject_QueuedTimeout(p *CodebuildProjectParameters, vals map[string]cty.Value) {
	vals["queued_timeout"] = cty.IntVal(p.QueuedTimeout)
}

func EncodeCodebuildProject_SourceVersion(p *CodebuildProjectParameters, vals map[string]cty.Value) {
	vals["source_version"] = cty.StringVal(p.SourceVersion)
}

func EncodeCodebuildProject_Artifacts(p *Artifacts, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Artifacts {
		ctyVal = make(map[string]cty.Value)
		EncodeCodebuildProject_Artifacts_Location(v, ctyVal)
		EncodeCodebuildProject_Artifacts_Name(v, ctyVal)
		EncodeCodebuildProject_Artifacts_Packaging(v, ctyVal)
		EncodeCodebuildProject_Artifacts_Type(v, ctyVal)
		EncodeCodebuildProject_Artifacts_Path(v, ctyVal)
		EncodeCodebuildProject_Artifacts_ArtifactIdentifier(v, ctyVal)
		EncodeCodebuildProject_Artifacts_EncryptionDisabled(v, ctyVal)
		EncodeCodebuildProject_Artifacts_NamespaceType(v, ctyVal)
		EncodeCodebuildProject_Artifacts_OverrideArtifactName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["artifacts"] = cty.ListVal(valsForCollection)
}

func EncodeCodebuildProject_Artifacts_Location(p *Artifacts, vals map[string]cty.Value) {
	vals["location"] = cty.StringVal(p.Location)
}

func EncodeCodebuildProject_Artifacts_Name(p *Artifacts, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCodebuildProject_Artifacts_Packaging(p *Artifacts, vals map[string]cty.Value) {
	vals["packaging"] = cty.StringVal(p.Packaging)
}

func EncodeCodebuildProject_Artifacts_Type(p *Artifacts, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeCodebuildProject_Artifacts_Path(p *Artifacts, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeCodebuildProject_Artifacts_ArtifactIdentifier(p *Artifacts, vals map[string]cty.Value) {
	vals["artifact_identifier"] = cty.StringVal(p.ArtifactIdentifier)
}

func EncodeCodebuildProject_Artifacts_EncryptionDisabled(p *Artifacts, vals map[string]cty.Value) {
	vals["encryption_disabled"] = cty.BoolVal(p.EncryptionDisabled)
}

func EncodeCodebuildProject_Artifacts_NamespaceType(p *Artifacts, vals map[string]cty.Value) {
	vals["namespace_type"] = cty.StringVal(p.NamespaceType)
}

func EncodeCodebuildProject_Artifacts_OverrideArtifactName(p *Artifacts, vals map[string]cty.Value) {
	vals["override_artifact_name"] = cty.BoolVal(p.OverrideArtifactName)
}

func EncodeCodebuildProject_Cache(p *Cache, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Cache {
		ctyVal = make(map[string]cty.Value)
		EncodeCodebuildProject_Cache_Location(v, ctyVal)
		EncodeCodebuildProject_Cache_Modes(v, ctyVal)
		EncodeCodebuildProject_Cache_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["cache"] = cty.ListVal(valsForCollection)
}

func EncodeCodebuildProject_Cache_Location(p *Cache, vals map[string]cty.Value) {
	vals["location"] = cty.StringVal(p.Location)
}

func EncodeCodebuildProject_Cache_Modes(p *Cache, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Modes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["modes"] = cty.ListVal(colVals)
}

func EncodeCodebuildProject_Cache_Type(p *Cache, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeCodebuildProject_Environment(p *Environment, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Environment {
		ctyVal = make(map[string]cty.Value)
		EncodeCodebuildProject_Environment_Image(v, ctyVal)
		EncodeCodebuildProject_Environment_ImagePullCredentialsType(v, ctyVal)
		EncodeCodebuildProject_Environment_PrivilegedMode(v, ctyVal)
		EncodeCodebuildProject_Environment_Type(v, ctyVal)
		EncodeCodebuildProject_Environment_Certificate(v, ctyVal)
		EncodeCodebuildProject_Environment_ComputeType(v, ctyVal)
		EncodeCodebuildProject_Environment_EnvironmentVariable(v.EnvironmentVariable, ctyVal)
		EncodeCodebuildProject_Environment_RegistryCredential(v.RegistryCredential, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["environment"] = cty.ListVal(valsForCollection)
}

func EncodeCodebuildProject_Environment_Image(p *Environment, vals map[string]cty.Value) {
	vals["image"] = cty.StringVal(p.Image)
}

func EncodeCodebuildProject_Environment_ImagePullCredentialsType(p *Environment, vals map[string]cty.Value) {
	vals["image_pull_credentials_type"] = cty.StringVal(p.ImagePullCredentialsType)
}

func EncodeCodebuildProject_Environment_PrivilegedMode(p *Environment, vals map[string]cty.Value) {
	vals["privileged_mode"] = cty.BoolVal(p.PrivilegedMode)
}

func EncodeCodebuildProject_Environment_Type(p *Environment, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeCodebuildProject_Environment_Certificate(p *Environment, vals map[string]cty.Value) {
	vals["certificate"] = cty.StringVal(p.Certificate)
}

func EncodeCodebuildProject_Environment_ComputeType(p *Environment, vals map[string]cty.Value) {
	vals["compute_type"] = cty.StringVal(p.ComputeType)
}

func EncodeCodebuildProject_Environment_EnvironmentVariable(p *EnvironmentVariable, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.EnvironmentVariable {
		ctyVal = make(map[string]cty.Value)
		EncodeCodebuildProject_Environment_EnvironmentVariable_Name(v, ctyVal)
		EncodeCodebuildProject_Environment_EnvironmentVariable_Type(v, ctyVal)
		EncodeCodebuildProject_Environment_EnvironmentVariable_Value(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["environment_variable"] = cty.ListVal(valsForCollection)
}

func EncodeCodebuildProject_Environment_EnvironmentVariable_Name(p *EnvironmentVariable, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCodebuildProject_Environment_EnvironmentVariable_Type(p *EnvironmentVariable, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeCodebuildProject_Environment_EnvironmentVariable_Value(p *EnvironmentVariable, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeCodebuildProject_Environment_RegistryCredential(p *RegistryCredential, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RegistryCredential {
		ctyVal = make(map[string]cty.Value)
		EncodeCodebuildProject_Environment_RegistryCredential_Credential(v, ctyVal)
		EncodeCodebuildProject_Environment_RegistryCredential_CredentialProvider(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["registry_credential"] = cty.ListVal(valsForCollection)
}

func EncodeCodebuildProject_Environment_RegistryCredential_Credential(p *RegistryCredential, vals map[string]cty.Value) {
	vals["credential"] = cty.StringVal(p.Credential)
}

func EncodeCodebuildProject_Environment_RegistryCredential_CredentialProvider(p *RegistryCredential, vals map[string]cty.Value) {
	vals["credential_provider"] = cty.StringVal(p.CredentialProvider)
}

func EncodeCodebuildProject_LogsConfig(p *LogsConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.LogsConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeCodebuildProject_LogsConfig_CloudwatchLogs(v.CloudwatchLogs, ctyVal)
		EncodeCodebuildProject_LogsConfig_S3Logs(v.S3Logs, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["logs_config"] = cty.ListVal(valsForCollection)
}

func EncodeCodebuildProject_LogsConfig_CloudwatchLogs(p *CloudwatchLogs, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.CloudwatchLogs {
		ctyVal = make(map[string]cty.Value)
		EncodeCodebuildProject_LogsConfig_CloudwatchLogs_StreamName(v, ctyVal)
		EncodeCodebuildProject_LogsConfig_CloudwatchLogs_GroupName(v, ctyVal)
		EncodeCodebuildProject_LogsConfig_CloudwatchLogs_Status(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["cloudwatch_logs"] = cty.ListVal(valsForCollection)
}

func EncodeCodebuildProject_LogsConfig_CloudwatchLogs_StreamName(p *CloudwatchLogs, vals map[string]cty.Value) {
	vals["stream_name"] = cty.StringVal(p.StreamName)
}

func EncodeCodebuildProject_LogsConfig_CloudwatchLogs_GroupName(p *CloudwatchLogs, vals map[string]cty.Value) {
	vals["group_name"] = cty.StringVal(p.GroupName)
}

func EncodeCodebuildProject_LogsConfig_CloudwatchLogs_Status(p *CloudwatchLogs, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeCodebuildProject_LogsConfig_S3Logs(p *S3Logs, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.S3Logs {
		ctyVal = make(map[string]cty.Value)
		EncodeCodebuildProject_LogsConfig_S3Logs_EncryptionDisabled(v, ctyVal)
		EncodeCodebuildProject_LogsConfig_S3Logs_Location(v, ctyVal)
		EncodeCodebuildProject_LogsConfig_S3Logs_Status(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["s3_logs"] = cty.ListVal(valsForCollection)
}

func EncodeCodebuildProject_LogsConfig_S3Logs_EncryptionDisabled(p *S3Logs, vals map[string]cty.Value) {
	vals["encryption_disabled"] = cty.BoolVal(p.EncryptionDisabled)
}

func EncodeCodebuildProject_LogsConfig_S3Logs_Location(p *S3Logs, vals map[string]cty.Value) {
	vals["location"] = cty.StringVal(p.Location)
}

func EncodeCodebuildProject_LogsConfig_S3Logs_Status(p *S3Logs, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeCodebuildProject_SecondaryArtifacts(p *SecondaryArtifacts, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SecondaryArtifacts {
		ctyVal = make(map[string]cty.Value)
		EncodeCodebuildProject_SecondaryArtifacts_EncryptionDisabled(v, ctyVal)
		EncodeCodebuildProject_SecondaryArtifacts_OverrideArtifactName(v, ctyVal)
		EncodeCodebuildProject_SecondaryArtifacts_Type(v, ctyVal)
		EncodeCodebuildProject_SecondaryArtifacts_ArtifactIdentifier(v, ctyVal)
		EncodeCodebuildProject_SecondaryArtifacts_Location(v, ctyVal)
		EncodeCodebuildProject_SecondaryArtifacts_Name(v, ctyVal)
		EncodeCodebuildProject_SecondaryArtifacts_NamespaceType(v, ctyVal)
		EncodeCodebuildProject_SecondaryArtifacts_Packaging(v, ctyVal)
		EncodeCodebuildProject_SecondaryArtifacts_Path(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["secondary_artifacts"] = cty.SetVal(valsForCollection)
}

func EncodeCodebuildProject_SecondaryArtifacts_EncryptionDisabled(p *SecondaryArtifacts, vals map[string]cty.Value) {
	vals["encryption_disabled"] = cty.BoolVal(p.EncryptionDisabled)
}

func EncodeCodebuildProject_SecondaryArtifacts_OverrideArtifactName(p *SecondaryArtifacts, vals map[string]cty.Value) {
	vals["override_artifact_name"] = cty.BoolVal(p.OverrideArtifactName)
}

func EncodeCodebuildProject_SecondaryArtifacts_Type(p *SecondaryArtifacts, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeCodebuildProject_SecondaryArtifacts_ArtifactIdentifier(p *SecondaryArtifacts, vals map[string]cty.Value) {
	vals["artifact_identifier"] = cty.StringVal(p.ArtifactIdentifier)
}

func EncodeCodebuildProject_SecondaryArtifacts_Location(p *SecondaryArtifacts, vals map[string]cty.Value) {
	vals["location"] = cty.StringVal(p.Location)
}

func EncodeCodebuildProject_SecondaryArtifacts_Name(p *SecondaryArtifacts, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCodebuildProject_SecondaryArtifacts_NamespaceType(p *SecondaryArtifacts, vals map[string]cty.Value) {
	vals["namespace_type"] = cty.StringVal(p.NamespaceType)
}

func EncodeCodebuildProject_SecondaryArtifacts_Packaging(p *SecondaryArtifacts, vals map[string]cty.Value) {
	vals["packaging"] = cty.StringVal(p.Packaging)
}

func EncodeCodebuildProject_SecondaryArtifacts_Path(p *SecondaryArtifacts, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeCodebuildProject_SecondarySources(p *SecondarySources, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SecondarySources {
		ctyVal = make(map[string]cty.Value)
		EncodeCodebuildProject_SecondarySources_SourceIdentifier(v, ctyVal)
		EncodeCodebuildProject_SecondarySources_Type(v, ctyVal)
		EncodeCodebuildProject_SecondarySources_Buildspec(v, ctyVal)
		EncodeCodebuildProject_SecondarySources_GitCloneDepth(v, ctyVal)
		EncodeCodebuildProject_SecondarySources_InsecureSsl(v, ctyVal)
		EncodeCodebuildProject_SecondarySources_Location(v, ctyVal)
		EncodeCodebuildProject_SecondarySources_ReportBuildStatus(v, ctyVal)
		EncodeCodebuildProject_SecondarySources_GitSubmodulesConfig(v.GitSubmodulesConfig, ctyVal)
		EncodeCodebuildProject_SecondarySources_Auth(v.Auth, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["secondary_sources"] = cty.SetVal(valsForCollection)
}

func EncodeCodebuildProject_SecondarySources_SourceIdentifier(p *SecondarySources, vals map[string]cty.Value) {
	vals["source_identifier"] = cty.StringVal(p.SourceIdentifier)
}

func EncodeCodebuildProject_SecondarySources_Type(p *SecondarySources, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeCodebuildProject_SecondarySources_Buildspec(p *SecondarySources, vals map[string]cty.Value) {
	vals["buildspec"] = cty.StringVal(p.Buildspec)
}

func EncodeCodebuildProject_SecondarySources_GitCloneDepth(p *SecondarySources, vals map[string]cty.Value) {
	vals["git_clone_depth"] = cty.IntVal(p.GitCloneDepth)
}

func EncodeCodebuildProject_SecondarySources_InsecureSsl(p *SecondarySources, vals map[string]cty.Value) {
	vals["insecure_ssl"] = cty.BoolVal(p.InsecureSsl)
}

func EncodeCodebuildProject_SecondarySources_Location(p *SecondarySources, vals map[string]cty.Value) {
	vals["location"] = cty.StringVal(p.Location)
}

func EncodeCodebuildProject_SecondarySources_ReportBuildStatus(p *SecondarySources, vals map[string]cty.Value) {
	vals["report_build_status"] = cty.BoolVal(p.ReportBuildStatus)
}

func EncodeCodebuildProject_SecondarySources_GitSubmodulesConfig(p *GitSubmodulesConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.GitSubmodulesConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeCodebuildProject_SecondarySources_GitSubmodulesConfig_FetchSubmodules(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["git_submodules_config"] = cty.ListVal(valsForCollection)
}

func EncodeCodebuildProject_SecondarySources_GitSubmodulesConfig_FetchSubmodules(p *GitSubmodulesConfig, vals map[string]cty.Value) {
	vals["fetch_submodules"] = cty.BoolVal(p.FetchSubmodules)
}

func EncodeCodebuildProject_SecondarySources_Auth(p *Auth, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Auth {
		ctyVal = make(map[string]cty.Value)
		EncodeCodebuildProject_SecondarySources_Auth_Resource(v, ctyVal)
		EncodeCodebuildProject_SecondarySources_Auth_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["auth"] = cty.ListVal(valsForCollection)
}

func EncodeCodebuildProject_SecondarySources_Auth_Resource(p *Auth, vals map[string]cty.Value) {
	vals["resource"] = cty.StringVal(p.Resource)
}

func EncodeCodebuildProject_SecondarySources_Auth_Type(p *Auth, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeCodebuildProject_Source(p *Source, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Source {
		ctyVal = make(map[string]cty.Value)
		EncodeCodebuildProject_Source_GitCloneDepth(v, ctyVal)
		EncodeCodebuildProject_Source_InsecureSsl(v, ctyVal)
		EncodeCodebuildProject_Source_Location(v, ctyVal)
		EncodeCodebuildProject_Source_ReportBuildStatus(v, ctyVal)
		EncodeCodebuildProject_Source_Type(v, ctyVal)
		EncodeCodebuildProject_Source_Buildspec(v, ctyVal)
		EncodeCodebuildProject_Source_GitSubmodulesConfig(v.GitSubmodulesConfig, ctyVal)
		EncodeCodebuildProject_Source_Auth(v.Auth, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["source"] = cty.ListVal(valsForCollection)
}

func EncodeCodebuildProject_Source_GitCloneDepth(p *Source, vals map[string]cty.Value) {
	vals["git_clone_depth"] = cty.IntVal(p.GitCloneDepth)
}

func EncodeCodebuildProject_Source_InsecureSsl(p *Source, vals map[string]cty.Value) {
	vals["insecure_ssl"] = cty.BoolVal(p.InsecureSsl)
}

func EncodeCodebuildProject_Source_Location(p *Source, vals map[string]cty.Value) {
	vals["location"] = cty.StringVal(p.Location)
}

func EncodeCodebuildProject_Source_ReportBuildStatus(p *Source, vals map[string]cty.Value) {
	vals["report_build_status"] = cty.BoolVal(p.ReportBuildStatus)
}

func EncodeCodebuildProject_Source_Type(p *Source, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeCodebuildProject_Source_Buildspec(p *Source, vals map[string]cty.Value) {
	vals["buildspec"] = cty.StringVal(p.Buildspec)
}

func EncodeCodebuildProject_Source_GitSubmodulesConfig(p *GitSubmodulesConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.GitSubmodulesConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeCodebuildProject_Source_GitSubmodulesConfig_FetchSubmodules(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["git_submodules_config"] = cty.ListVal(valsForCollection)
}

func EncodeCodebuildProject_Source_GitSubmodulesConfig_FetchSubmodules(p *GitSubmodulesConfig, vals map[string]cty.Value) {
	vals["fetch_submodules"] = cty.BoolVal(p.FetchSubmodules)
}

func EncodeCodebuildProject_Source_Auth(p *Auth, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Auth {
		ctyVal = make(map[string]cty.Value)
		EncodeCodebuildProject_Source_Auth_Resource(v, ctyVal)
		EncodeCodebuildProject_Source_Auth_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["auth"] = cty.ListVal(valsForCollection)
}

func EncodeCodebuildProject_Source_Auth_Resource(p *Auth, vals map[string]cty.Value) {
	vals["resource"] = cty.StringVal(p.Resource)
}

func EncodeCodebuildProject_Source_Auth_Type(p *Auth, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeCodebuildProject_VpcConfig(p *VpcConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.VpcConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeCodebuildProject_VpcConfig_SecurityGroupIds(v, ctyVal)
		EncodeCodebuildProject_VpcConfig_Subnets(v, ctyVal)
		EncodeCodebuildProject_VpcConfig_VpcId(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["vpc_config"] = cty.ListVal(valsForCollection)
}

func EncodeCodebuildProject_VpcConfig_SecurityGroupIds(p *VpcConfig, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_group_ids"] = cty.SetVal(colVals)
}

func EncodeCodebuildProject_VpcConfig_Subnets(p *VpcConfig, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Subnets {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnets"] = cty.SetVal(colVals)
}

func EncodeCodebuildProject_VpcConfig_VpcId(p *VpcConfig, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeCodebuildProject_Arn(p *CodebuildProjectObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeCodebuildProject_BadgeUrl(p *CodebuildProjectObservation, vals map[string]cty.Value) {
	vals["badge_url"] = cty.StringVal(p.BadgeUrl)
}