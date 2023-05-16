// Copyright 2021 Google LLC
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
// source: google/spanner/v1/result_set.proto

package spanner

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Results from [Read][google.spanner.v1.Spanner.Read] or
// [ExecuteSql][google.spanner.v1.Spanner.ExecuteSql].
type ResultSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Metadata about the result set, such as row type information.
	Metadata *ResultSetMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Each element in `rows` is a row whose format is defined by
	// [metadata.row_type][google.spanner.v1.ResultSetMetadata.row_type]. The ith element
	// in each row matches the ith field in
	// [metadata.row_type][google.spanner.v1.ResultSetMetadata.row_type]. Elements are
	// encoded based on type as described
	// [here][google.spanner.v1.TypeCode].
	Rows []*structpb.ListValue `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
	// Query plan and execution statistics for the SQL statement that
	// produced this result set. These can be requested by setting
	// [ExecuteSqlRequest.query_mode][google.spanner.v1.ExecuteSqlRequest.query_mode].
	// DML statements always produce stats containing the number of rows
	// modified, unless executed using the
	// [ExecuteSqlRequest.QueryMode.PLAN][google.spanner.v1.ExecuteSqlRequest.QueryMode.PLAN] [ExecuteSqlRequest.query_mode][google.spanner.v1.ExecuteSqlRequest.query_mode].
	// Other fields may or may not be populated, based on the
	// [ExecuteSqlRequest.query_mode][google.spanner.v1.ExecuteSqlRequest.query_mode].
	Stats *ResultSetStats `protobuf:"bytes,3,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (x *ResultSet) Reset() {
	*x = ResultSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_spanner_v1_result_set_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultSet) ProtoMessage() {}

func (x *ResultSet) ProtoReflect() protoreflect.Message {
	mi := &file_google_spanner_v1_result_set_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultSet.ProtoReflect.Descriptor instead.
func (*ResultSet) Descriptor() ([]byte, []int) {
	return file_google_spanner_v1_result_set_proto_rawDescGZIP(), []int{0}
}

func (x *ResultSet) GetMetadata() *ResultSetMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *ResultSet) GetRows() []*structpb.ListValue {
	if x != nil {
		return x.Rows
	}
	return nil
}

func (x *ResultSet) GetStats() *ResultSetStats {
	if x != nil {
		return x.Stats
	}
	return nil
}

// Partial results from a streaming read or SQL query. Streaming reads and
// SQL queries better tolerate large result sets, large rows, and large
// values, but are a little trickier to consume.
type PartialResultSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Metadata about the result set, such as row type information.
	// Only present in the first response.
	Metadata *ResultSetMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// A streamed result set consists of a stream of values, which might
	// be split into many `PartialResultSet` messages to accommodate
	// large rows and/or large values. Every N complete values defines a
	// row, where N is equal to the number of entries in
	// [metadata.row_type.fields][google.spanner.v1.StructType.fields].
	//
	// Most values are encoded based on type as described
	// [here][google.spanner.v1.TypeCode].
	//
	// It is possible that the last value in values is "chunked",
	// meaning that the rest of the value is sent in subsequent
	// `PartialResultSet`(s). This is denoted by the [chunked_value][google.spanner.v1.PartialResultSet.chunked_value]
	// field. Two or more chunked values can be merged to form a
	// complete value as follows:
	//
	//   * `bool/number/null`: cannot be chunked
	//   * `string`: concatenate the strings
	//   * `list`: concatenate the lists. If the last element in a list is a
	//     `string`, `list`, or `object`, merge it with the first element in
	//     the next list by applying these rules recursively.
	//   * `object`: concatenate the (field name, field value) pairs. If a
	//     field name is duplicated, then apply these rules recursively
	//     to merge the field values.
	//
	// Some examples of merging:
	//
	//     # Strings are concatenated.
	//     "foo", "bar" => "foobar"
	//
	//     # Lists of non-strings are concatenated.
	//     [2, 3], [4] => [2, 3, 4]
	//
	//     # Lists are concatenated, but the last and first elements are merged
	//     # because they are strings.
	//     ["a", "b"], ["c", "d"] => ["a", "bc", "d"]
	//
	//     # Lists are concatenated, but the last and first elements are merged
	//     # because they are lists. Recursively, the last and first elements
	//     # of the inner lists are merged because they are strings.
	//     ["a", ["b", "c"]], [["d"], "e"] => ["a", ["b", "cd"], "e"]
	//
	//     # Non-overlapping object fields are combined.
	//     {"a": "1"}, {"b": "2"} => {"a": "1", "b": 2"}
	//
	//     # Overlapping object fields are merged.
	//     {"a": "1"}, {"a": "2"} => {"a": "12"}
	//
	//     # Examples of merging objects containing lists of strings.
	//     {"a": ["1"]}, {"a": ["2"]} => {"a": ["12"]}
	//
	// For a more complete example, suppose a streaming SQL query is
	// yielding a result set whose rows contain a single string
	// field. The following `PartialResultSet`s might be yielded:
	//
	//     {
	//       "metadata": { ... }
	//       "values": ["Hello", "W"]
	//       "chunked_value": true
	//       "resume_token": "Af65..."
	//     }
	//     {
	//       "values": ["orl"]
	//       "chunked_value": true
	//       "resume_token": "Bqp2..."
	//     }
	//     {
	//       "values": ["d"]
	//       "resume_token": "Zx1B..."
	//     }
	//
	// This sequence of `PartialResultSet`s encodes two rows, one
	// containing the field value `"Hello"`, and a second containing the
	// field value `"World" = "W" + "orl" + "d"`.
	Values []*structpb.Value `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	// If true, then the final value in [values][google.spanner.v1.PartialResultSet.values] is chunked, and must
	// be combined with more values from subsequent `PartialResultSet`s
	// to obtain a complete field value.
	ChunkedValue bool `protobuf:"varint,3,opt,name=chunked_value,json=chunkedValue,proto3" json:"chunked_value,omitempty"`
	// Streaming calls might be interrupted for a variety of reasons, such
	// as TCP connection loss. If this occurs, the stream of results can
	// be resumed by re-sending the original request and including
	// `resume_token`. Note that executing any other transaction in the
	// same session invalidates the token.
	ResumeToken []byte `protobuf:"bytes,4,opt,name=resume_token,json=resumeToken,proto3" json:"resume_token,omitempty"`
	// Query plan and execution statistics for the statement that produced this
	// streaming result set. These can be requested by setting
	// [ExecuteSqlRequest.query_mode][google.spanner.v1.ExecuteSqlRequest.query_mode] and are sent
	// only once with the last response in the stream.
	// This field will also be present in the last response for DML
	// statements.
	Stats *ResultSetStats `protobuf:"bytes,5,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (x *PartialResultSet) Reset() {
	*x = PartialResultSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_spanner_v1_result_set_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartialResultSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartialResultSet) ProtoMessage() {}

func (x *PartialResultSet) ProtoReflect() protoreflect.Message {
	mi := &file_google_spanner_v1_result_set_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartialResultSet.ProtoReflect.Descriptor instead.
func (*PartialResultSet) Descriptor() ([]byte, []int) {
	return file_google_spanner_v1_result_set_proto_rawDescGZIP(), []int{1}
}

func (x *PartialResultSet) GetMetadata() *ResultSetMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *PartialResultSet) GetValues() []*structpb.Value {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *PartialResultSet) GetChunkedValue() bool {
	if x != nil {
		return x.ChunkedValue
	}
	return false
}

func (x *PartialResultSet) GetResumeToken() []byte {
	if x != nil {
		return x.ResumeToken
	}
	return nil
}

func (x *PartialResultSet) GetStats() *ResultSetStats {
	if x != nil {
		return x.Stats
	}
	return nil
}

// Metadata about a [ResultSet][google.spanner.v1.ResultSet] or [PartialResultSet][google.spanner.v1.PartialResultSet].
type ResultSetMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indicates the field names and types for the rows in the result
	// set.  For example, a SQL query like `"SELECT UserId, UserName FROM
	// Users"` could return a `row_type` value like:
	//
	//     "fields": [
	//       { "name": "UserId", "type": { "code": "INT64" } },
	//       { "name": "UserName", "type": { "code": "STRING" } },
	//     ]
	RowType *StructType `protobuf:"bytes,1,opt,name=row_type,json=rowType,proto3" json:"row_type,omitempty"`
	// If the read or SQL query began a transaction as a side-effect, the
	// information about the new transaction is yielded here.
	Transaction *Transaction `protobuf:"bytes,2,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *ResultSetMetadata) Reset() {
	*x = ResultSetMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_spanner_v1_result_set_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultSetMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultSetMetadata) ProtoMessage() {}

func (x *ResultSetMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_spanner_v1_result_set_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultSetMetadata.ProtoReflect.Descriptor instead.
func (*ResultSetMetadata) Descriptor() ([]byte, []int) {
	return file_google_spanner_v1_result_set_proto_rawDescGZIP(), []int{2}
}

func (x *ResultSetMetadata) GetRowType() *StructType {
	if x != nil {
		return x.RowType
	}
	return nil
}

func (x *ResultSetMetadata) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

// Additional statistics about a [ResultSet][google.spanner.v1.ResultSet] or [PartialResultSet][google.spanner.v1.PartialResultSet].
type ResultSetStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// [QueryPlan][google.spanner.v1.QueryPlan] for the query associated with this result.
	QueryPlan *QueryPlan `protobuf:"bytes,1,opt,name=query_plan,json=queryPlan,proto3" json:"query_plan,omitempty"`
	// Aggregated statistics from the execution of the query. Only present when
	// the query is profiled. For example, a query could return the statistics as
	// follows:
	//
	//     {
	//       "rows_returned": "3",
	//       "elapsed_time": "1.22 secs",
	//       "cpu_time": "1.19 secs"
	//     }
	QueryStats *structpb.Struct `protobuf:"bytes,2,opt,name=query_stats,json=queryStats,proto3" json:"query_stats,omitempty"`
	// The number of rows modified by the DML statement.
	//
	// Types that are assignable to RowCount:
	//	*ResultSetStats_RowCountExact
	//	*ResultSetStats_RowCountLowerBound
	RowCount isResultSetStats_RowCount `protobuf_oneof:"row_count"`
}

func (x *ResultSetStats) Reset() {
	*x = ResultSetStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_spanner_v1_result_set_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultSetStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultSetStats) ProtoMessage() {}

func (x *ResultSetStats) ProtoReflect() protoreflect.Message {
	mi := &file_google_spanner_v1_result_set_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultSetStats.ProtoReflect.Descriptor instead.
func (*ResultSetStats) Descriptor() ([]byte, []int) {
	return file_google_spanner_v1_result_set_proto_rawDescGZIP(), []int{3}
}

func (x *ResultSetStats) GetQueryPlan() *QueryPlan {
	if x != nil {
		return x.QueryPlan
	}
	return nil
}

func (x *ResultSetStats) GetQueryStats() *structpb.Struct {
	if x != nil {
		return x.QueryStats
	}
	return nil
}

func (m *ResultSetStats) GetRowCount() isResultSetStats_RowCount {
	if m != nil {
		return m.RowCount
	}
	return nil
}

func (x *ResultSetStats) GetRowCountExact() int64 {
	if x, ok := x.GetRowCount().(*ResultSetStats_RowCountExact); ok {
		return x.RowCountExact
	}
	return 0
}

func (x *ResultSetStats) GetRowCountLowerBound() int64 {
	if x, ok := x.GetRowCount().(*ResultSetStats_RowCountLowerBound); ok {
		return x.RowCountLowerBound
	}
	return 0
}

type isResultSetStats_RowCount interface {
	isResultSetStats_RowCount()
}

type ResultSetStats_RowCountExact struct {
	// Standard DML returns an exact count of rows that were modified.
	RowCountExact int64 `protobuf:"varint,3,opt,name=row_count_exact,json=rowCountExact,proto3,oneof"`
}

type ResultSetStats_RowCountLowerBound struct {
	// Partitioned DML does not offer exactly-once semantics, so it
	// returns a lower bound of the rows modified.
	RowCountLowerBound int64 `protobuf:"varint,4,opt,name=row_count_lower_bound,json=rowCountLowerBound,proto3,oneof"`
}

func (*ResultSetStats_RowCountExact) isResultSetStats_RowCount() {}

func (*ResultSetStats_RowCountLowerBound) isResultSetStats_RowCount() {}

var File_google_spanner_v1_result_set_proto protoreflect.FileDescriptor

var file_google_spanner_v1_result_set_proto_rawDesc = []byte{
	0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x73, 0x70,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x53, 0x65, 0x74, 0x12, 0x40, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73,
	0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x53, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x04, 0x72, 0x6f, 0x77, 0x73, 0x12, 0x37, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22, 0x85,
	0x02, 0x0a, 0x10, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x53, 0x65, 0x74, 0x12, 0x40, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73,
	0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x53, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x65, 0x64,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x73, 0x75, 0x6d, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x37, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22, 0x8f, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x08,
	0x72, 0x6f, 0x77, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x72,
	0x6f, 0x77, 0x54, 0x79, 0x70, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xf3, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x0a, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x09, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x38, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x65, 0x78, 0x61, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0d, 0x72,
	0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x78, 0x61, 0x63, 0x74, 0x12, 0x33, 0x0a, 0x15,
	0x72, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x12, 0x72,
	0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x77, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e,
	0x64, 0x42, 0x0b, 0x0a, 0x09, 0x72, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0xb7,
	0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x53, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x70, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0xf8, 0x01, 0x01, 0xaa, 0x02, 0x17, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x17, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x53, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x70, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_spanner_v1_result_set_proto_rawDescOnce sync.Once
	file_google_spanner_v1_result_set_proto_rawDescData = file_google_spanner_v1_result_set_proto_rawDesc
)

func file_google_spanner_v1_result_set_proto_rawDescGZIP() []byte {
	file_google_spanner_v1_result_set_proto_rawDescOnce.Do(func() {
		file_google_spanner_v1_result_set_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_spanner_v1_result_set_proto_rawDescData)
	})
	return file_google_spanner_v1_result_set_proto_rawDescData
}

var file_google_spanner_v1_result_set_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_spanner_v1_result_set_proto_goTypes = []interface{}{
	(*ResultSet)(nil),          // 0: google.spanner.v1.ResultSet
	(*PartialResultSet)(nil),   // 1: google.spanner.v1.PartialResultSet
	(*ResultSetMetadata)(nil),  // 2: google.spanner.v1.ResultSetMetadata
	(*ResultSetStats)(nil),     // 3: google.spanner.v1.ResultSetStats
	(*structpb.ListValue)(nil), // 4: google.protobuf.ListValue
	(*structpb.Value)(nil),     // 5: google.protobuf.Value
	(*StructType)(nil),         // 6: google.spanner.v1.StructType
	(*Transaction)(nil),        // 7: google.spanner.v1.Transaction
	(*QueryPlan)(nil),          // 8: google.spanner.v1.QueryPlan
	(*structpb.Struct)(nil),    // 9: google.protobuf.Struct
}
var file_google_spanner_v1_result_set_proto_depIdxs = []int32{
	2,  // 0: google.spanner.v1.ResultSet.metadata:type_name -> google.spanner.v1.ResultSetMetadata
	4,  // 1: google.spanner.v1.ResultSet.rows:type_name -> google.protobuf.ListValue
	3,  // 2: google.spanner.v1.ResultSet.stats:type_name -> google.spanner.v1.ResultSetStats
	2,  // 3: google.spanner.v1.PartialResultSet.metadata:type_name -> google.spanner.v1.ResultSetMetadata
	5,  // 4: google.spanner.v1.PartialResultSet.values:type_name -> google.protobuf.Value
	3,  // 5: google.spanner.v1.PartialResultSet.stats:type_name -> google.spanner.v1.ResultSetStats
	6,  // 6: google.spanner.v1.ResultSetMetadata.row_type:type_name -> google.spanner.v1.StructType
	7,  // 7: google.spanner.v1.ResultSetMetadata.transaction:type_name -> google.spanner.v1.Transaction
	8,  // 8: google.spanner.v1.ResultSetStats.query_plan:type_name -> google.spanner.v1.QueryPlan
	9,  // 9: google.spanner.v1.ResultSetStats.query_stats:type_name -> google.protobuf.Struct
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_google_spanner_v1_result_set_proto_init() }
func file_google_spanner_v1_result_set_proto_init() {
	if File_google_spanner_v1_result_set_proto != nil {
		return
	}
	file_google_spanner_v1_commit_response_proto_init()
	file_google_spanner_v1_query_plan_proto_init()
	file_google_spanner_v1_transaction_proto_init()
	file_google_spanner_v1_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_spanner_v1_result_set_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultSet); i {
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
		file_google_spanner_v1_result_set_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartialResultSet); i {
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
		file_google_spanner_v1_result_set_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultSetMetadata); i {
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
		file_google_spanner_v1_result_set_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultSetStats); i {
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
	file_google_spanner_v1_result_set_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*ResultSetStats_RowCountExact)(nil),
		(*ResultSetStats_RowCountLowerBound)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_spanner_v1_result_set_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_spanner_v1_result_set_proto_goTypes,
		DependencyIndexes: file_google_spanner_v1_result_set_proto_depIdxs,
		MessageInfos:      file_google_spanner_v1_result_set_proto_msgTypes,
	}.Build()
	File_google_spanner_v1_result_set_proto = out.File
	file_google_spanner_v1_result_set_proto_rawDesc = nil
	file_google_spanner_v1_result_set_proto_goTypes = nil
	file_google_spanner_v1_result_set_proto_depIdxs = nil
}
