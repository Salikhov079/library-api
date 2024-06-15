// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.4
// source: books.proto

package genprotos

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

type BookReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	AuthorId string `protobuf:"bytes,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	GenreId  string `protobuf:"bytes,4,opt,name=genre_id,json=genreId,proto3" json:"genre_id,omitempty"`
	Summary  string `protobuf:"bytes,5,opt,name=summary,proto3" json:"summary,omitempty"`
}

func (x *BookReq) Reset() {
	*x = BookReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookReq) ProtoMessage() {}

func (x *BookReq) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookReq.ProtoReflect.Descriptor instead.
func (*BookReq) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{0}
}

func (x *BookReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BookReq) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *BookReq) GetGenreId() string {
	if x != nil {
		return x.GenreId
	}
	return ""
}

func (x *BookReq) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

type BookRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string     `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author  *AuthorRes `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Genre   *Genre     `protobuf:"bytes,4,opt,name=genre,proto3" json:"genre,omitempty"`
	Summary string     `protobuf:"bytes,5,opt,name=summary,proto3" json:"summary,omitempty"`
}

func (x *BookRes) Reset() {
	*x = BookRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookRes) ProtoMessage() {}

func (x *BookRes) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookRes.ProtoReflect.Descriptor instead.
func (*BookRes) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{1}
}

func (x *BookRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BookRes) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BookRes) GetAuthor() *AuthorRes {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *BookRes) GetGenre() *Genre {
	if x != nil {
		return x.Genre
	}
	return nil
}

func (x *BookRes) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

type AllBooks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*BookRes `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *AllBooks) Reset() {
	*x = AllBooks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllBooks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllBooks) ProtoMessage() {}

func (x *AllBooks) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllBooks.ProtoReflect.Descriptor instead.
func (*AllBooks) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{2}
}

func (x *AllBooks) GetBooks() []*BookRes {
	if x != nil {
		return x.Books
	}
	return nil
}

var File_books_proto protoreflect.FileDescriptor

var file_books_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0c, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x71, 0x0a, 0x07, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x22, 0x99, 0x01, 0x0a, 0x07, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12,
	0x23, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x52, 0x05, 0x67,
	0x65, 0x6e, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x31,
	0x0a, 0x08, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x25, 0x0a, 0x05, 0x62, 0x6f,
	0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x32, 0xfb, 0x01, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71,
	0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00,
	0x12, 0x2d, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x1a,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12,
	0x2a, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x00, 0x42,
	0x0c, 0x5a, 0x0a, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_books_proto_rawDescOnce sync.Once
	file_books_proto_rawDescData = file_books_proto_rawDesc
)

func file_books_proto_rawDescGZIP() []byte {
	file_books_proto_rawDescOnce.Do(func() {
		file_books_proto_rawDescData = protoimpl.X.CompressGZIP(file_books_proto_rawDescData)
	})
	return file_books_proto_rawDescData
}

var file_books_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_books_proto_goTypes = []interface{}{
	(*BookReq)(nil),   // 0: protos.BookReq
	(*BookRes)(nil),   // 1: protos.BookRes
	(*AllBooks)(nil),  // 2: protos.AllBooks
	(*AuthorRes)(nil), // 3: protos.AuthorRes
	(*Genre)(nil),     // 4: protos.Genre
	(*ById)(nil),      // 5: protos.ById
	(*Void)(nil),      // 6: protos.Void
}
var file_books_proto_depIdxs = []int32{
	3, // 0: protos.BookRes.author:type_name -> protos.AuthorRes
	4, // 1: protos.BookRes.genre:type_name -> protos.Genre
	1, // 2: protos.AllBooks.books:type_name -> protos.BookRes
	0, // 3: protos.BookService.CreateBook:input_type -> protos.BookReq
	1, // 4: protos.BookService.UpdateBook:input_type -> protos.BookRes
	5, // 5: protos.BookService.DeleteBook:input_type -> protos.ById
	5, // 6: protos.BookService.GetByIdBook:input_type -> protos.ById
	0, // 7: protos.BookService.GetAllBooks:input_type -> protos.BookReq
	6, // 8: protos.BookService.CreateBook:output_type -> protos.Void
	6, // 9: protos.BookService.UpdateBook:output_type -> protos.Void
	6, // 10: protos.BookService.DeleteBook:output_type -> protos.Void
	1, // 11: protos.BookService.GetByIdBook:output_type -> protos.BookRes
	2, // 12: protos.BookService.GetAllBooks:output_type -> protos.AllBooks
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_books_proto_init() }
func file_books_proto_init() {
	if File_books_proto != nil {
		return
	}
	file_users_proto_init()
	file_authors_proto_init()
	file_genres_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_books_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookReq); i {
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
		file_books_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookRes); i {
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
		file_books_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllBooks); i {
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
			RawDescriptor: file_books_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_books_proto_goTypes,
		DependencyIndexes: file_books_proto_depIdxs,
		MessageInfos:      file_books_proto_msgTypes,
	}.Build()
	File_books_proto = out.File
	file_books_proto_rawDesc = nil
	file_books_proto_goTypes = nil
	file_books_proto_depIdxs = nil
}
