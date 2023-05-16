// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/cloud/aiplatform/v1beta1/dataset.proto

package aiplatform

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A collection of DataItems and Annotations on them.
type Dataset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the Dataset.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The user-defined name of the Dataset.
	// The name can be up to 128 characters long and can be consist of any UTF-8
	// characters.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Optional. The description of the Dataset.
	Description string `protobuf:"bytes,16,opt,name=description,proto3" json:"description,omitempty"`
	// Required. Points to a YAML file stored on Google Cloud Storage describing additional
	// information about the Dataset.
	// The schema is defined as an OpenAPI 3.0.2 Schema Object.
	// The schema files that can be used here are found in
	// gs://google-cloud-aiplatform/schema/dataset/metadata/.
	MetadataSchemaUri string `protobuf:"bytes,3,opt,name=metadata_schema_uri,json=metadataSchemaUri,proto3" json:"metadata_schema_uri,omitempty"`
	// Required. Additional information about the Dataset.
	Metadata *structpb.Value `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Output only. Timestamp when this Dataset was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Timestamp when this Dataset was last updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Used to perform consistent read-modify-write updates. If not set, a blind
	// "overwrite" update happens.
	Etag string `protobuf:"bytes,6,opt,name=etag,proto3" json:"etag,omitempty"`
	// The labels with user-defined metadata to organize your Datasets.
	//
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	// No more than 64 user labels can be associated with one Dataset (System
	// labels are excluded).
	//
	// See https://goo.gl/xmQnxf for more information and examples of labels.
	// System reserved label keys are prefixed with "aiplatform.googleapis.com/"
	// and are immutable. Following system labels exist for each Dataset:
	//
	// * "aiplatform.googleapis.com/dataset_metadata_schema": output only, its
	//   value is the [metadata_schema's][google.cloud.aiplatform.v1beta1.Dataset.metadata_schema_uri] title.
	Labels map[string]string `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Customer-managed encryption key spec for a Dataset. If set, this Dataset
	// and all sub-resources of this Dataset will be secured by this key.
	EncryptionSpec *EncryptionSpec `protobuf:"bytes,11,opt,name=encryption_spec,json=encryptionSpec,proto3" json:"encryption_spec,omitempty"`
}

func (x *Dataset) Reset() {
	*x = Dataset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_dataset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dataset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dataset) ProtoMessage() {}

func (x *Dataset) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_dataset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dataset.ProtoReflect.Descriptor instead.
func (*Dataset) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_dataset_proto_rawDescGZIP(), []int{0}
}

func (x *Dataset) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Dataset) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Dataset) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Dataset) GetMetadataSchemaUri() string {
	if x != nil {
		return x.MetadataSchemaUri
	}
	return ""
}

func (x *Dataset) GetMetadata() *structpb.Value {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Dataset) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Dataset) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Dataset) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *Dataset) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Dataset) GetEncryptionSpec() *EncryptionSpec {
	if x != nil {
		return x.EncryptionSpec
	}
	return nil
}

// Describes the location from where we import data into a Dataset, together
// with the labels that will be applied to the DataItems and the Annotations.
type ImportDataConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The source of the input.
	//
	// Types that are assignable to Source:
	//	*ImportDataConfig_GcsSource
	Source isImportDataConfig_Source `protobuf_oneof:"source"`
	// Labels that will be applied to newly imported DataItems. If an identical
	// DataItem as one being imported already exists in the Dataset, then these
	// labels will be appended to these of the already existing one, and if labels
	// with identical key is imported before, the old label value will be
	// overwritten. If two DataItems are identical in the same import data
	// operation, the labels will be combined and if key collision happens in this
	// case, one of the values will be picked randomly. Two DataItems are
	// considered identical if their content bytes are identical (e.g. image bytes
	// or pdf bytes).
	// These labels will be overridden by Annotation labels specified inside index
	// file referenced by [import_schema_uri][google.cloud.aiplatform.v1beta1.ImportDataConfig.import_schema_uri], e.g. jsonl file.
	DataItemLabels map[string]string `protobuf:"bytes,2,rep,name=data_item_labels,json=dataItemLabels,proto3" json:"data_item_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Required. Points to a YAML file stored on Google Cloud Storage describing the import
	// format. Validation will be done against the schema. The schema is defined
	// as an [OpenAPI 3.0.2 Schema
	// Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject).
	ImportSchemaUri string `protobuf:"bytes,4,opt,name=import_schema_uri,json=importSchemaUri,proto3" json:"import_schema_uri,omitempty"`
}

func (x *ImportDataConfig) Reset() {
	*x = ImportDataConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_dataset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportDataConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportDataConfig) ProtoMessage() {}

func (x *ImportDataConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_dataset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportDataConfig.ProtoReflect.Descriptor instead.
func (*ImportDataConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_dataset_proto_rawDescGZIP(), []int{1}
}

func (m *ImportDataConfig) GetSource() isImportDataConfig_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *ImportDataConfig) GetGcsSource() *GcsSource {
	if x, ok := x.GetSource().(*ImportDataConfig_GcsSource); ok {
		return x.GcsSource
	}
	return nil
}

func (x *ImportDataConfig) GetDataItemLabels() map[string]string {
	if x != nil {
		return x.DataItemLabels
	}
	return nil
}

func (x *ImportDataConfig) GetImportSchemaUri() string {
	if x != nil {
		return x.ImportSchemaUri
	}
	return ""
}

type isImportDataConfig_Source interface {
	isImportDataConfig_Source()
}

type ImportDataConfig_GcsSource struct {
	// The Google Cloud Storage location for the input content.
	GcsSource *GcsSource `protobuf:"bytes,1,opt,name=gcs_source,json=gcsSource,proto3,oneof"`
}

func (*ImportDataConfig_GcsSource) isImportDataConfig_Source() {}

// Describes what part of the Dataset is to be exported, the destination of
// the export and how to export.
type ExportDataConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The destination of the output.
	//
	// Types that are assignable to Destination:
	//	*ExportDataConfig_GcsDestination
	Destination isExportDataConfig_Destination `protobuf_oneof:"destination"`
	// A filter on Annotations of the Dataset. Only Annotations on to-be-exported
	// DataItems(specified by [data_items_filter][]) that match this filter will
	// be exported. The filter syntax is the same as in
	// [ListAnnotations][google.cloud.aiplatform.v1beta1.DatasetService.ListAnnotations].
	AnnotationsFilter string `protobuf:"bytes,2,opt,name=annotations_filter,json=annotationsFilter,proto3" json:"annotations_filter,omitempty"`
}

func (x *ExportDataConfig) Reset() {
	*x = ExportDataConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_dataset_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportDataConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportDataConfig) ProtoMessage() {}

func (x *ExportDataConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_dataset_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportDataConfig.ProtoReflect.Descriptor instead.
func (*ExportDataConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_dataset_proto_rawDescGZIP(), []int{2}
}

func (m *ExportDataConfig) GetDestination() isExportDataConfig_Destination {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (x *ExportDataConfig) GetGcsDestination() *GcsDestination {
	if x, ok := x.GetDestination().(*ExportDataConfig_GcsDestination); ok {
		return x.GcsDestination
	}
	return nil
}

func (x *ExportDataConfig) GetAnnotationsFilter() string {
	if x != nil {
		return x.AnnotationsFilter
	}
	return ""
}

type isExportDataConfig_Destination interface {
	isExportDataConfig_Destination()
}

type ExportDataConfig_GcsDestination struct {
	// The Google Cloud Storage location where the output is to be written to.
	// In the given directory a new directory will be created with name:
	// `export-data-<dataset-display-name>-<timestamp-of-export-call>` where
	// timestamp is in YYYY-MM-DDThh:mm:ss.sssZ ISO-8601 format. All export
	// output will be written into that directory. Inside that directory,
	// annotations with the same schema will be grouped into sub directories
	// which are named with the corresponding annotations' schema title. Inside
	// these sub directories, a schema.yaml will be created to describe the
	// output format.
	GcsDestination *GcsDestination `protobuf:"bytes,1,opt,name=gcs_destination,json=gcsDestination,proto3,oneof"`
}

func (*ExportDataConfig_GcsDestination) isExportDataConfig_Destination() {}

var File_google_cloud_aiplatform_v1beta1_dataset_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_dataset_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x05, 0x0a,
	0x07, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x33, 0x0a, 0x13, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x11, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x55, 0x72, 0x69, 0x12, 0x37, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x40,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x12, 0x4c, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x58, 0x0a, 0x0f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0e,
	0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x1a, 0x39,
	0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x62, 0xea, 0x41, 0x5f, 0x0a, 0x21,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x12, 0x3a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x7d, 0x22, 0xce, 0x02,
	0x0a, 0x10, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x4b, 0x0a, 0x0a, 0x67, 0x63, 0x73, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x63, 0x73, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x48, 0x00, 0x52, 0x09, 0x67, 0x63, 0x73, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x6f, 0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x49, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x12, 0x2f, 0x0a, 0x11, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x0f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x55, 0x72,
	0x69, 0x1a, 0x41, 0x0a, 0x13, 0x44, 0x61, 0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0xac,
	0x01, 0x0a, 0x10, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x5a, 0x0a, 0x0f, 0x67, 0x63, 0x73, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47,
	0x63, 0x73, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52,
	0x0e, 0x67, 0x63, 0x73, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2d, 0x0a, 0x12, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x0d,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0xe9, 0x01,
	0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74,
	0x61, 0x31, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_dataset_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_dataset_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_dataset_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_dataset_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_dataset_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_dataset_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_dataset_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_dataset_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_dataset_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_aiplatform_v1beta1_dataset_proto_goTypes = []interface{}{
	(*Dataset)(nil),               // 0: google.cloud.aiplatform.v1beta1.Dataset
	(*ImportDataConfig)(nil),      // 1: google.cloud.aiplatform.v1beta1.ImportDataConfig
	(*ExportDataConfig)(nil),      // 2: google.cloud.aiplatform.v1beta1.ExportDataConfig
	nil,                           // 3: google.cloud.aiplatform.v1beta1.Dataset.LabelsEntry
	nil,                           // 4: google.cloud.aiplatform.v1beta1.ImportDataConfig.DataItemLabelsEntry
	(*structpb.Value)(nil),        // 5: google.protobuf.Value
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*EncryptionSpec)(nil),        // 7: google.cloud.aiplatform.v1beta1.EncryptionSpec
	(*GcsSource)(nil),             // 8: google.cloud.aiplatform.v1beta1.GcsSource
	(*GcsDestination)(nil),        // 9: google.cloud.aiplatform.v1beta1.GcsDestination
}
var file_google_cloud_aiplatform_v1beta1_dataset_proto_depIdxs = []int32{
	5, // 0: google.cloud.aiplatform.v1beta1.Dataset.metadata:type_name -> google.protobuf.Value
	6, // 1: google.cloud.aiplatform.v1beta1.Dataset.create_time:type_name -> google.protobuf.Timestamp
	6, // 2: google.cloud.aiplatform.v1beta1.Dataset.update_time:type_name -> google.protobuf.Timestamp
	3, // 3: google.cloud.aiplatform.v1beta1.Dataset.labels:type_name -> google.cloud.aiplatform.v1beta1.Dataset.LabelsEntry
	7, // 4: google.cloud.aiplatform.v1beta1.Dataset.encryption_spec:type_name -> google.cloud.aiplatform.v1beta1.EncryptionSpec
	8, // 5: google.cloud.aiplatform.v1beta1.ImportDataConfig.gcs_source:type_name -> google.cloud.aiplatform.v1beta1.GcsSource
	4, // 6: google.cloud.aiplatform.v1beta1.ImportDataConfig.data_item_labels:type_name -> google.cloud.aiplatform.v1beta1.ImportDataConfig.DataItemLabelsEntry
	9, // 7: google.cloud.aiplatform.v1beta1.ExportDataConfig.gcs_destination:type_name -> google.cloud.aiplatform.v1beta1.GcsDestination
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1beta1_dataset_proto_init() }
func file_google_cloud_aiplatform_v1beta1_dataset_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_dataset_proto != nil {
		return
	}
	file_google_cloud_aiplatform_v1beta1_encryption_spec_proto_init()
	file_google_cloud_aiplatform_v1beta1_io_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1beta1_dataset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dataset); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_aiplatform_v1beta1_dataset_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportDataConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_aiplatform_v1beta1_dataset_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportDataConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_google_cloud_aiplatform_v1beta1_dataset_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ImportDataConfig_GcsSource)(nil),
	}
	file_google_cloud_aiplatform_v1beta1_dataset_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ExportDataConfig_GcsDestination)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_dataset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_dataset_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_dataset_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1beta1_dataset_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_dataset_proto = out.File
	file_google_cloud_aiplatform_v1beta1_dataset_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_dataset_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_dataset_proto_depIdxs = nil
}
