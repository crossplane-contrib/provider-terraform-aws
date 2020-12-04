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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
)

// +kubebuilder:object:root=true

// ElastictranscoderPreset is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ElastictranscoderPreset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ElastictranscoderPresetSpec   `json:"spec"`
	Status ElastictranscoderPresetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElastictranscoderPreset contains a list of ElastictranscoderPresetList
type ElastictranscoderPresetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElastictranscoderPreset `json:"items"`
}

// A ElastictranscoderPresetSpec defines the desired state of a ElastictranscoderPreset
type ElastictranscoderPresetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ElastictranscoderPresetParameters `json:",inline"`
}

// A ElastictranscoderPresetParameters defines the desired state of a ElastictranscoderPreset
type ElastictranscoderPresetParameters struct {
	Container         string            `json:"container"`
	Description       string            `json:"description"`
	Id                string            `json:"id"`
	Name              string            `json:"name"`
	Type              string            `json:"type"`
	VideoCodecOptions map[string]string `json:"video_codec_options"`
	Video             Video             `json:"video"`
	VideoWatermarks   VideoWatermarks   `json:"video_watermarks"`
	Audio             Audio             `json:"audio"`
	AudioCodecOptions AudioCodecOptions `json:"audio_codec_options"`
	Thumbnails        Thumbnails        `json:"thumbnails"`
}

type Video struct {
	MaxFrameRate       string `json:"max_frame_rate"`
	SizingPolicy       string `json:"sizing_policy"`
	DisplayAspectRatio string `json:"display_aspect_ratio"`
	FrameRate          string `json:"frame_rate"`
	Codec              string `json:"codec"`
	FixedGop           string `json:"fixed_gop"`
	KeyframesMaxDist   string `json:"keyframes_max_dist"`
	MaxHeight          string `json:"max_height"`
	MaxWidth           string `json:"max_width"`
	PaddingPolicy      string `json:"padding_policy"`
	AspectRatio        string `json:"aspect_ratio"`
	BitRate            string `json:"bit_rate"`
	Resolution         string `json:"resolution"`
}

type VideoWatermarks struct {
	Target           string `json:"target"`
	HorizontalAlign  string `json:"horizontal_align"`
	HorizontalOffset string `json:"horizontal_offset"`
	SizingPolicy     string `json:"sizing_policy"`
	Opacity          string `json:"opacity"`
	VerticalAlign    string `json:"vertical_align"`
	VerticalOffset   string `json:"vertical_offset"`
	Id               string `json:"id"`
	MaxHeight        string `json:"max_height"`
	MaxWidth         string `json:"max_width"`
}

type Audio struct {
	AudioPackingMode string `json:"audio_packing_mode"`
	BitRate          string `json:"bit_rate"`
	Channels         string `json:"channels"`
	Codec            string `json:"codec"`
	SampleRate       string `json:"sample_rate"`
}

type AudioCodecOptions struct {
	BitOrder string `json:"bit_order"`
	Profile  string `json:"profile"`
	Signed   string `json:"signed"`
	BitDepth string `json:"bit_depth"`
}

type Thumbnails struct {
	SizingPolicy  string `json:"sizing_policy"`
	AspectRatio   string `json:"aspect_ratio"`
	Format        string `json:"format"`
	Interval      string `json:"interval"`
	MaxHeight     string `json:"max_height"`
	MaxWidth      string `json:"max_width"`
	PaddingPolicy string `json:"padding_policy"`
	Resolution    string `json:"resolution"`
}

// A ElastictranscoderPresetStatus defines the observed state of a ElastictranscoderPreset
type ElastictranscoderPresetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ElastictranscoderPresetObservation `json:",inline"`
}

// A ElastictranscoderPresetObservation records the observed state of a ElastictranscoderPreset
type ElastictranscoderPresetObservation struct {
	Arn string `json:"arn"`
}