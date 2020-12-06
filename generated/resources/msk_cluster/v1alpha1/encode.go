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

func EncodeMskCluster(r MskCluster) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeMskCluster_EnhancedMonitoring(r.Spec.ForProvider, ctyVal)
	EncodeMskCluster_Tags(r.Spec.ForProvider, ctyVal)
	EncodeMskCluster_Id(r.Spec.ForProvider, ctyVal)
	EncodeMskCluster_KafkaVersion(r.Spec.ForProvider, ctyVal)
	EncodeMskCluster_NumberOfBrokerNodes(r.Spec.ForProvider, ctyVal)
	EncodeMskCluster_ClusterName(r.Spec.ForProvider, ctyVal)
	EncodeMskCluster_BrokerNodeGroupInfo(r.Spec.ForProvider.BrokerNodeGroupInfo, ctyVal)
	EncodeMskCluster_ClientAuthentication(r.Spec.ForProvider.ClientAuthentication, ctyVal)
	EncodeMskCluster_ConfigurationInfo(r.Spec.ForProvider.ConfigurationInfo, ctyVal)
	EncodeMskCluster_EncryptionInfo(r.Spec.ForProvider.EncryptionInfo, ctyVal)
	EncodeMskCluster_LoggingInfo(r.Spec.ForProvider.LoggingInfo, ctyVal)
	EncodeMskCluster_OpenMonitoring(r.Spec.ForProvider.OpenMonitoring, ctyVal)
	EncodeMskCluster_CurrentVersion(r.Status.AtProvider, ctyVal)
	EncodeMskCluster_ZookeeperConnectString(r.Status.AtProvider, ctyVal)
	EncodeMskCluster_Arn(r.Status.AtProvider, ctyVal)
	EncodeMskCluster_BootstrapBrokers(r.Status.AtProvider, ctyVal)
	EncodeMskCluster_BootstrapBrokersTls(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeMskCluster_EnhancedMonitoring(p MskClusterParameters, vals map[string]cty.Value) {
	vals["enhanced_monitoring"] = cty.StringVal(p.EnhancedMonitoring)
}

func EncodeMskCluster_Tags(p MskClusterParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeMskCluster_Id(p MskClusterParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeMskCluster_KafkaVersion(p MskClusterParameters, vals map[string]cty.Value) {
	vals["kafka_version"] = cty.StringVal(p.KafkaVersion)
}

func EncodeMskCluster_NumberOfBrokerNodes(p MskClusterParameters, vals map[string]cty.Value) {
	vals["number_of_broker_nodes"] = cty.NumberIntVal(p.NumberOfBrokerNodes)
}

func EncodeMskCluster_ClusterName(p MskClusterParameters, vals map[string]cty.Value) {
	vals["cluster_name"] = cty.StringVal(p.ClusterName)
}

func EncodeMskCluster_BrokerNodeGroupInfo(p BrokerNodeGroupInfo, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeMskCluster_BrokerNodeGroupInfo_EbsVolumeSize(p, ctyVal)
	EncodeMskCluster_BrokerNodeGroupInfo_InstanceType(p, ctyVal)
	EncodeMskCluster_BrokerNodeGroupInfo_SecurityGroups(p, ctyVal)
	EncodeMskCluster_BrokerNodeGroupInfo_AzDistribution(p, ctyVal)
	EncodeMskCluster_BrokerNodeGroupInfo_ClientSubnets(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["broker_node_group_info"] = cty.ListVal(valsForCollection)
}

func EncodeMskCluster_BrokerNodeGroupInfo_EbsVolumeSize(p BrokerNodeGroupInfo, vals map[string]cty.Value) {
	vals["ebs_volume_size"] = cty.NumberIntVal(p.EbsVolumeSize)
}

func EncodeMskCluster_BrokerNodeGroupInfo_InstanceType(p BrokerNodeGroupInfo, vals map[string]cty.Value) {
	vals["instance_type"] = cty.StringVal(p.InstanceType)
}

func EncodeMskCluster_BrokerNodeGroupInfo_SecurityGroups(p BrokerNodeGroupInfo, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroups {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_groups"] = cty.ListVal(colVals)
}

func EncodeMskCluster_BrokerNodeGroupInfo_AzDistribution(p BrokerNodeGroupInfo, vals map[string]cty.Value) {
	vals["az_distribution"] = cty.StringVal(p.AzDistribution)
}

func EncodeMskCluster_BrokerNodeGroupInfo_ClientSubnets(p BrokerNodeGroupInfo, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ClientSubnets {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["client_subnets"] = cty.ListVal(colVals)
}

func EncodeMskCluster_ClientAuthentication(p ClientAuthentication, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeMskCluster_ClientAuthentication_Tls(p.Tls, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["client_authentication"] = cty.ListVal(valsForCollection)
}

func EncodeMskCluster_ClientAuthentication_Tls(p Tls, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeMskCluster_ClientAuthentication_Tls_CertificateAuthorityArns(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["tls"] = cty.ListVal(valsForCollection)
}

func EncodeMskCluster_ClientAuthentication_Tls_CertificateAuthorityArns(p Tls, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CertificateAuthorityArns {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["certificate_authority_arns"] = cty.SetVal(colVals)
}

func EncodeMskCluster_ConfigurationInfo(p ConfigurationInfo, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeMskCluster_ConfigurationInfo_Arn(p, ctyVal)
	EncodeMskCluster_ConfigurationInfo_Revision(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["configuration_info"] = cty.ListVal(valsForCollection)
}

func EncodeMskCluster_ConfigurationInfo_Arn(p ConfigurationInfo, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeMskCluster_ConfigurationInfo_Revision(p ConfigurationInfo, vals map[string]cty.Value) {
	vals["revision"] = cty.NumberIntVal(p.Revision)
}

func EncodeMskCluster_EncryptionInfo(p EncryptionInfo, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeMskCluster_EncryptionInfo_EncryptionAtRestKmsKeyArn(p, ctyVal)
	EncodeMskCluster_EncryptionInfo_EncryptionInTransit(p.EncryptionInTransit, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["encryption_info"] = cty.ListVal(valsForCollection)
}

func EncodeMskCluster_EncryptionInfo_EncryptionAtRestKmsKeyArn(p EncryptionInfo, vals map[string]cty.Value) {
	vals["encryption_at_rest_kms_key_arn"] = cty.StringVal(p.EncryptionAtRestKmsKeyArn)
}

func EncodeMskCluster_EncryptionInfo_EncryptionInTransit(p EncryptionInTransit, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeMskCluster_EncryptionInfo_EncryptionInTransit_ClientBroker(p, ctyVal)
	EncodeMskCluster_EncryptionInfo_EncryptionInTransit_InCluster(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["encryption_in_transit"] = cty.ListVal(valsForCollection)
}

func EncodeMskCluster_EncryptionInfo_EncryptionInTransit_ClientBroker(p EncryptionInTransit, vals map[string]cty.Value) {
	vals["client_broker"] = cty.StringVal(p.ClientBroker)
}

func EncodeMskCluster_EncryptionInfo_EncryptionInTransit_InCluster(p EncryptionInTransit, vals map[string]cty.Value) {
	vals["in_cluster"] = cty.BoolVal(p.InCluster)
}

func EncodeMskCluster_LoggingInfo(p LoggingInfo, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeMskCluster_LoggingInfo_BrokerLogs(p.BrokerLogs, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["logging_info"] = cty.ListVal(valsForCollection)
}

func EncodeMskCluster_LoggingInfo_BrokerLogs(p BrokerLogs, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeMskCluster_LoggingInfo_BrokerLogs_CloudwatchLogs(p.CloudwatchLogs, ctyVal)
	EncodeMskCluster_LoggingInfo_BrokerLogs_Firehose(p.Firehose, ctyVal)
	EncodeMskCluster_LoggingInfo_BrokerLogs_S3(p.S3, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["broker_logs"] = cty.ListVal(valsForCollection)
}

func EncodeMskCluster_LoggingInfo_BrokerLogs_CloudwatchLogs(p CloudwatchLogs, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeMskCluster_LoggingInfo_BrokerLogs_CloudwatchLogs_Enabled(p, ctyVal)
	EncodeMskCluster_LoggingInfo_BrokerLogs_CloudwatchLogs_LogGroup(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["cloudwatch_logs"] = cty.ListVal(valsForCollection)
}

func EncodeMskCluster_LoggingInfo_BrokerLogs_CloudwatchLogs_Enabled(p CloudwatchLogs, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeMskCluster_LoggingInfo_BrokerLogs_CloudwatchLogs_LogGroup(p CloudwatchLogs, vals map[string]cty.Value) {
	vals["log_group"] = cty.StringVal(p.LogGroup)
}

func EncodeMskCluster_LoggingInfo_BrokerLogs_Firehose(p Firehose, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeMskCluster_LoggingInfo_BrokerLogs_Firehose_DeliveryStream(p, ctyVal)
	EncodeMskCluster_LoggingInfo_BrokerLogs_Firehose_Enabled(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["firehose"] = cty.ListVal(valsForCollection)
}

func EncodeMskCluster_LoggingInfo_BrokerLogs_Firehose_DeliveryStream(p Firehose, vals map[string]cty.Value) {
	vals["delivery_stream"] = cty.StringVal(p.DeliveryStream)
}

func EncodeMskCluster_LoggingInfo_BrokerLogs_Firehose_Enabled(p Firehose, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeMskCluster_LoggingInfo_BrokerLogs_S3(p S3, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeMskCluster_LoggingInfo_BrokerLogs_S3_Bucket(p, ctyVal)
	EncodeMskCluster_LoggingInfo_BrokerLogs_S3_Enabled(p, ctyVal)
	EncodeMskCluster_LoggingInfo_BrokerLogs_S3_Prefix(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["s3"] = cty.ListVal(valsForCollection)
}

func EncodeMskCluster_LoggingInfo_BrokerLogs_S3_Bucket(p S3, vals map[string]cty.Value) {
	vals["bucket"] = cty.StringVal(p.Bucket)
}

func EncodeMskCluster_LoggingInfo_BrokerLogs_S3_Enabled(p S3, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeMskCluster_LoggingInfo_BrokerLogs_S3_Prefix(p S3, vals map[string]cty.Value) {
	vals["prefix"] = cty.StringVal(p.Prefix)
}

func EncodeMskCluster_OpenMonitoring(p OpenMonitoring, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeMskCluster_OpenMonitoring_Prometheus(p.Prometheus, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["open_monitoring"] = cty.ListVal(valsForCollection)
}

func EncodeMskCluster_OpenMonitoring_Prometheus(p Prometheus, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeMskCluster_OpenMonitoring_Prometheus_NodeExporter(p.NodeExporter, ctyVal)
	EncodeMskCluster_OpenMonitoring_Prometheus_JmxExporter(p.JmxExporter, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["prometheus"] = cty.ListVal(valsForCollection)
}

func EncodeMskCluster_OpenMonitoring_Prometheus_NodeExporter(p NodeExporter, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeMskCluster_OpenMonitoring_Prometheus_NodeExporter_EnabledInBroker(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["node_exporter"] = cty.ListVal(valsForCollection)
}

func EncodeMskCluster_OpenMonitoring_Prometheus_NodeExporter_EnabledInBroker(p NodeExporter, vals map[string]cty.Value) {
	vals["enabled_in_broker"] = cty.BoolVal(p.EnabledInBroker)
}

func EncodeMskCluster_OpenMonitoring_Prometheus_JmxExporter(p JmxExporter, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeMskCluster_OpenMonitoring_Prometheus_JmxExporter_EnabledInBroker(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["jmx_exporter"] = cty.ListVal(valsForCollection)
}

func EncodeMskCluster_OpenMonitoring_Prometheus_JmxExporter_EnabledInBroker(p JmxExporter, vals map[string]cty.Value) {
	vals["enabled_in_broker"] = cty.BoolVal(p.EnabledInBroker)
}

func EncodeMskCluster_CurrentVersion(p MskClusterObservation, vals map[string]cty.Value) {
	vals["current_version"] = cty.StringVal(p.CurrentVersion)
}

func EncodeMskCluster_ZookeeperConnectString(p MskClusterObservation, vals map[string]cty.Value) {
	vals["zookeeper_connect_string"] = cty.StringVal(p.ZookeeperConnectString)
}

func EncodeMskCluster_Arn(p MskClusterObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeMskCluster_BootstrapBrokers(p MskClusterObservation, vals map[string]cty.Value) {
	vals["bootstrap_brokers"] = cty.StringVal(p.BootstrapBrokers)
}

func EncodeMskCluster_BootstrapBrokersTls(p MskClusterObservation, vals map[string]cty.Value) {
	vals["bootstrap_brokers_tls"] = cty.StringVal(p.BootstrapBrokersTls)
}