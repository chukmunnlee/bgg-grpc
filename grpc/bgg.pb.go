// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: grpc/bgg.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Game struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId     int32  `protobuf:"varint,1,opt,name=gameId,proto3" json:"gameId,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Year       int32  `protobuf:"varint,3,opt,name=year,proto3" json:"year,omitempty"`
	Ranking    int32  `protobuf:"varint,4,opt,name=ranking,proto3" json:"ranking,omitempty"`
	UsersRated int32  `protobuf:"varint,5,opt,name=usersRated,proto3" json:"usersRated,omitempty"`
	Url        string `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	Image      string `protobuf:"bytes,7,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *Game) Reset() {
	*x = Game{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_bgg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Game) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Game) ProtoMessage() {}

func (x *Game) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_bgg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Game.ProtoReflect.Descriptor instead.
func (*Game) Descriptor() ([]byte, []int) {
	return file_grpc_bgg_proto_rawDescGZIP(), []int{0}
}

func (x *Game) GetGameId() int32 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *Game) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Game) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Game) GetRanking() int32 {
	if x != nil {
		return x.Ranking
	}
	return 0
}

func (x *Game) GetUsersRated() int32 {
	if x != nil {
		return x.UsersRated
	}
	return 0
}

func (x *Game) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Game) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	User   string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Rating int32  `protobuf:"varint,3,opt,name=rating,proto3" json:"rating,omitempty"`
	Text   string `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	GameId int32  `protobuf:"varint,5,opt,name=gameId,proto3" json:"gameId,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_bgg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_bgg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_grpc_bgg_proto_rawDescGZIP(), []int{1}
}

func (x *Comment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Comment) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Comment) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Comment) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Comment) GetGameId() int32 {
	if x != nil {
		return x.GameId
	}
	return 0
}

type FindBoardgameByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query  string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Limit  int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *FindBoardgameByNameRequest) Reset() {
	*x = FindBoardgameByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_bgg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindBoardgameByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindBoardgameByNameRequest) ProtoMessage() {}

func (x *FindBoardgameByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_bgg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindBoardgameByNameRequest.ProtoReflect.Descriptor instead.
func (*FindBoardgameByNameRequest) Descriptor() ([]byte, []int) {
	return file_grpc_bgg_proto_rawDescGZIP(), []int{2}
}

func (x *FindBoardgameByNameRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *FindBoardgameByNameRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FindBoardgameByNameRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type FindBoardgameByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Games  []*Game `protobuf:"bytes,1,rep,name=games,proto3" json:"games,omitempty"`
	Limit  int32   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32   `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Total  int32   `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *FindBoardgameByNameResponse) Reset() {
	*x = FindBoardgameByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_bgg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindBoardgameByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindBoardgameByNameResponse) ProtoMessage() {}

func (x *FindBoardgameByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_bgg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindBoardgameByNameResponse.ProtoReflect.Descriptor instead.
func (*FindBoardgameByNameResponse) Descriptor() ([]byte, []int) {
	return file_grpc_bgg_proto_rawDescGZIP(), []int{3}
}

func (x *FindBoardgameByNameResponse) GetGames() []*Game {
	if x != nil {
		return x.Games
	}
	return nil
}

func (x *FindBoardgameByNameResponse) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FindBoardgameByNameResponse) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *FindBoardgameByNameResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_grpc_bgg_proto protoreflect.FileDescriptor

var file_grpc_bgg_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x67, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0xa8, 0x01, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79,
	0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x52, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x61, 0x74, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x22, 0x71, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x61,
	0x6d, 0x65, 0x49, 0x64, 0x22, 0x60, 0x0a, 0x1a, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x61, 0x72,
	0x64, 0x67, 0x61, 0x6d, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x83, 0x01, 0x0a, 0x1b, 0x46, 0x69, 0x6e, 0x64, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x67, 0x61, 0x6d, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x61, 0x6d,
	0x65, 0x52, 0x05, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x6a, 0x0a, 0x0a,
	0x42, 0x47, 0x47, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x13, 0x46, 0x69,
	0x6e, 0x64, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x67, 0x61, 0x6d, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x61,
	0x72, 0x64, 0x67, 0x61, 0x6d, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x67, 0x61, 0x6d, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x75, 0x6b, 0x6d, 0x75, 0x6e, 0x6e, 0x6c,
	0x65, 0x65, 0x2f, 0x62, 0x67, 0x67, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_bgg_proto_rawDescOnce sync.Once
	file_grpc_bgg_proto_rawDescData = file_grpc_bgg_proto_rawDesc
)

func file_grpc_bgg_proto_rawDescGZIP() []byte {
	file_grpc_bgg_proto_rawDescOnce.Do(func() {
		file_grpc_bgg_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_bgg_proto_rawDescData)
	})
	return file_grpc_bgg_proto_rawDescData
}

var file_grpc_bgg_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpc_bgg_proto_goTypes = []interface{}{
	(*Game)(nil),                        // 0: grpc.Game
	(*Comment)(nil),                     // 1: grpc.Comment
	(*FindBoardgameByNameRequest)(nil),  // 2: grpc.FindBoardgameByNameRequest
	(*FindBoardgameByNameResponse)(nil), // 3: grpc.FindBoardgameByNameResponse
}
var file_grpc_bgg_proto_depIdxs = []int32{
	0, // 0: grpc.FindBoardgameByNameResponse.games:type_name -> grpc.Game
	2, // 1: grpc.BGGService.FindBoardgameByName:input_type -> grpc.FindBoardgameByNameRequest
	3, // 2: grpc.BGGService.FindBoardgameByName:output_type -> grpc.FindBoardgameByNameResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_bgg_proto_init() }
func file_grpc_bgg_proto_init() {
	if File_grpc_bgg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_bgg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Game); i {
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
		file_grpc_bgg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_grpc_bgg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindBoardgameByNameRequest); i {
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
		file_grpc_bgg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindBoardgameByNameResponse); i {
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
			RawDescriptor: file_grpc_bgg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_bgg_proto_goTypes,
		DependencyIndexes: file_grpc_bgg_proto_depIdxs,
		MessageInfos:      file_grpc_bgg_proto_msgTypes,
	}.Build()
	File_grpc_bgg_proto = out.File
	file_grpc_bgg_proto_rawDesc = nil
	file_grpc_bgg_proto_goTypes = nil
	file_grpc_bgg_proto_depIdxs = nil
}
