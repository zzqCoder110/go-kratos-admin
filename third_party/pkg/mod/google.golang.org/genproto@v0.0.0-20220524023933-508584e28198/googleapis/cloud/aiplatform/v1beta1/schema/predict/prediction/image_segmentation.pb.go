// Copyright 2020 Google LLC
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
// source: google/cloud/aiplatform/v1beta1/schema/predict/prediction/image_segmentation.proto

package prediction

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Prediction output format for Image Segmentation.
type ImageSegmentationPredictionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A PNG image where each pixel in the mask represents the category in which
	// the pixel in the original image was predicted to belong to. The size of
	// this image will be the same as the original image. The mapping between the
	// AnntoationSpec and the color can be found in model's metadata. The model
	// will choose the most likely category and if none of the categories reach
	// the confidence threshold, the pixel will be marked as background.
	CategoryMask string `protobuf:"bytes,1,opt,name=category_mask,json=categoryMask,proto3" json:"category_mask,omitempty"`
	// A one channel image which is encoded as an 8bit lossless PNG. The size of
	// the image will be the same as the original image. For a specific pixel,
	// darker color means less confidence in correctness of the cateogry in the
	// categoryMask for the corresponding pixel. Black means no confidence and
	// white means complete confidence.
	ConfidenceMask string `protobuf:"bytes,2,opt,name=confidence_mask,json=confidenceMask,proto3" json:"confidence_mask,omitempty"`
}

func (x *ImageSegmentationPredictionResult) Reset() {
	*x = ImageSegmentationPredictionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageSegmentationPredictionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageSegmentationPredictionResult) ProtoMessage() {}

func (x *ImageSegmentationPredictionResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageSegmentationPredictionResult.ProtoReflect.Descriptor instead.
func (*ImageSegmentationPredictionResult) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto_rawDescGZIP(), []int{0}
}

func (x *ImageSegmentationPredictionResult) GetCategoryMask() string {
	if x != nil {
		return x.CategoryMask
	}
	return ""
}

func (x *ImageSegmentationPredictionResult) GetConfidenceMask() string {
	if x != nil {
		return x.ConfidenceMask
	}
	return ""
}

var File_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto_rawDesc = []byte{
	0x0a, 0x52, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x65,
	0x64, 0x69, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a,
	0x21, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x4d, 0x61, 0x73, 0x6b,
	0x42, 0x88, 0x03, 0x0a, 0x3d, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x26, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x63, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0xaa, 0x02, 0x39, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0x42, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69,
	0x63, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0xca, 0x02, 0x39,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x49, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5c, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x5c, 0x50,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0xea, 0x02, 0x3f, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3a, 0x3a,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x3a, 0x3a, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x3a,
	0x3a, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto_goTypes = []interface{}{
	(*ImageSegmentationPredictionResult)(nil), // 0: google.cloud.aiplatform.v1beta1.schema.predict.prediction.ImageSegmentationPredictionResult
}
var file_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto_init()
}
func file_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageSegmentationPredictionResult); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto = out.File
	file_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_schema_predict_prediction_image_segmentation_proto_depIdxs = nil
}
