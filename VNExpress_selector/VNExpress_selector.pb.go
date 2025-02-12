// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.1
// source: VNExpress_selector.proto

package VNExpress_selector

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

type Range struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MainCategories []string `protobuf:"bytes,1,rep,name=main_categories,json=mainCategories,proto3" json:"main_categories,omitempty"`
	SubCategories  []string `protobuf:"bytes,2,rep,name=sub_categories,json=subCategories,proto3" json:"sub_categories,omitempty"`
	Author         []string `protobuf:"bytes,3,rep,name=author,proto3" json:"author,omitempty"`
	Day            []string `protobuf:"bytes,4,rep,name=day,proto3" json:"day,omitempty"`
	Time           []string `protobuf:"bytes,5,rep,name=time,proto3" json:"time,omitempty"`
	WholeDay       bool     `protobuf:"varint,6,opt,name=whole_day,json=wholeDay,proto3" json:"whole_day,omitempty"`
	DayComparisor  []string `protobuf:"bytes,7,rep,name=day_comparisor,json=dayComparisor,proto3" json:"day_comparisor,omitempty"`
	TimeComparisor string   `protobuf:"bytes,8,opt,name=time_comparisor,json=timeComparisor,proto3" json:"time_comparisor,omitempty"`
	Limit          string   `protobuf:"bytes,9,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *Range) Reset() {
	*x = Range{}
	if protoimpl.UnsafeEnabled {
		mi := &file_VNExpress_selector_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Range) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Range) ProtoMessage() {}

func (x *Range) ProtoReflect() protoreflect.Message {
	mi := &file_VNExpress_selector_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Range.ProtoReflect.Descriptor instead.
func (*Range) Descriptor() ([]byte, []int) {
	return file_VNExpress_selector_proto_rawDescGZIP(), []int{0}
}

func (x *Range) GetMainCategories() []string {
	if x != nil {
		return x.MainCategories
	}
	return nil
}

func (x *Range) GetSubCategories() []string {
	if x != nil {
		return x.SubCategories
	}
	return nil
}

func (x *Range) GetAuthor() []string {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Range) GetDay() []string {
	if x != nil {
		return x.Day
	}
	return nil
}

func (x *Range) GetTime() []string {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Range) GetWholeDay() bool {
	if x != nil {
		return x.WholeDay
	}
	return false
}

func (x *Range) GetDayComparisor() []string {
	if x != nil {
		return x.DayComparisor
	}
	return nil
}

func (x *Range) GetTimeComparisor() string {
	if x != nil {
		return x.TimeComparisor
	}
	return ""
}

func (x *Range) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

type News struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url           string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	MainCategory  string   `protobuf:"bytes,2,opt,name=main_category,json=mainCategory,proto3" json:"main_category,omitempty"`
	SubCategory   string   `protobuf:"bytes,3,opt,name=sub_category,json=subCategory,proto3" json:"sub_category,omitempty"`
	Title         string   `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Day           string   `protobuf:"bytes,5,opt,name=day,proto3" json:"day,omitempty"`
	Time          string   `protobuf:"bytes,6,opt,name=time,proto3" json:"time,omitempty"`
	TimeZone      string   `protobuf:"bytes,7,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
	Description   string   `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	NewsContent   string   `protobuf:"bytes,9,opt,name=news_content,json=newsContent,proto3" json:"news_content,omitempty"`
	RelatingImage []string `protobuf:"bytes,10,rep,name=relating_image,json=relatingImage,proto3" json:"relating_image,omitempty"`
	Author        string   `protobuf:"bytes,11,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *News) Reset() {
	*x = News{}
	if protoimpl.UnsafeEnabled {
		mi := &file_VNExpress_selector_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *News) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*News) ProtoMessage() {}

func (x *News) ProtoReflect() protoreflect.Message {
	mi := &file_VNExpress_selector_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use News.ProtoReflect.Descriptor instead.
func (*News) Descriptor() ([]byte, []int) {
	return file_VNExpress_selector_proto_rawDescGZIP(), []int{1}
}

func (x *News) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *News) GetMainCategory() string {
	if x != nil {
		return x.MainCategory
	}
	return ""
}

func (x *News) GetSubCategory() string {
	if x != nil {
		return x.SubCategory
	}
	return ""
}

func (x *News) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *News) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

func (x *News) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *News) GetTimeZone() string {
	if x != nil {
		return x.TimeZone
	}
	return ""
}

func (x *News) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *News) GetNewsContent() string {
	if x != nil {
		return x.NewsContent
	}
	return ""
}

func (x *News) GetRelatingImage() []string {
	if x != nil {
		return x.RelatingImage
	}
	return nil
}

func (x *News) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type Podcast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url             string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	MainCategory    string `protobuf:"bytes,2,opt,name=main_category,json=mainCategory,proto3" json:"main_category,omitempty"`
	SubCategory     string `protobuf:"bytes,3,opt,name=sub_category,json=subCategory,proto3" json:"sub_category,omitempty"`
	Title           string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Day             string `protobuf:"bytes,5,opt,name=day,proto3" json:"day,omitempty"`
	Time            string `protobuf:"bytes,6,opt,name=time,proto3" json:"time,omitempty"`
	TimeZone        string `protobuf:"bytes,7,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
	Description     string `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	RelatingPodcast string `protobuf:"bytes,9,opt,name=relating_podcast,json=relatingPodcast,proto3" json:"relating_podcast,omitempty"`
	Author          string `protobuf:"bytes,10,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *Podcast) Reset() {
	*x = Podcast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_VNExpress_selector_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Podcast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Podcast) ProtoMessage() {}

func (x *Podcast) ProtoReflect() protoreflect.Message {
	mi := &file_VNExpress_selector_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Podcast.ProtoReflect.Descriptor instead.
func (*Podcast) Descriptor() ([]byte, []int) {
	return file_VNExpress_selector_proto_rawDescGZIP(), []int{2}
}

func (x *Podcast) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Podcast) GetMainCategory() string {
	if x != nil {
		return x.MainCategory
	}
	return ""
}

func (x *Podcast) GetSubCategory() string {
	if x != nil {
		return x.SubCategory
	}
	return ""
}

func (x *Podcast) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Podcast) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

func (x *Podcast) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *Podcast) GetTimeZone() string {
	if x != nil {
		return x.TimeZone
	}
	return ""
}

func (x *Podcast) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Podcast) GetRelatingPodcast() string {
	if x != nil {
		return x.RelatingPodcast
	}
	return ""
}

func (x *Podcast) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

var File_VNExpress_selector_proto protoreflect.FileDescriptor

var file_VNExpress_selector_proto_rawDesc = []byte{
	0x0a, 0x18, 0x56, 0x4e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x76, 0x6e, 0x65, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x98,
	0x02, 0x0a, 0x05, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x61, 0x69, 0x6e,
	0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0e, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x64,
	0x61, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x68, 0x6f, 0x6c, 0x65, 0x5f,
	0x64, 0x61, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x77, 0x68, 0x6f, 0x6c, 0x65,
	0x44, 0x61, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x61, 0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x72, 0x69, 0x73, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x61, 0x79,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x72, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69,
	0x73, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0xbd, 0x02, 0x0a, 0x04, 0x4e, 0x65,
	0x77, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x69,
	0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x62,
	0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x75, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x64, 0x61, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x73, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e,
	0x65, 0x77, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0xa1, 0x02, 0x0a, 0x07, 0x50, 0x6f,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x69, 0x6e, 0x5f,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6d, 0x61, 0x69, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x75, 0x62, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6f,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x32, 0xaa, 0x01,
	0x0a, 0x12, 0x56, 0x4e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x12, 0x46, 0x0a, 0x0b, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x6e,
	0x65, 0x77, 0x73, 0x12, 0x19, 0x2e, 0x76, 0x6e, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x5f,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x1a, 0x18,
	0x2e, 0x76, 0x6e, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4c, 0x0a, 0x0e,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x19,
	0x2e, 0x76, 0x6e, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x1a, 0x1b, 0x2e, 0x76, 0x6e, 0x65, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50,
	0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x22, 0x00, 0x30, 0x01, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x6f, 0x6e, 0x67, 0x51, 0x75, 0x61,
	0x6e, 0x31, 0x38, 0x31, 0x32, 0x2f, 0x56, 0x4e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x56, 0x4e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_VNExpress_selector_proto_rawDescOnce sync.Once
	file_VNExpress_selector_proto_rawDescData = file_VNExpress_selector_proto_rawDesc
)

func file_VNExpress_selector_proto_rawDescGZIP() []byte {
	file_VNExpress_selector_proto_rawDescOnce.Do(func() {
		file_VNExpress_selector_proto_rawDescData = protoimpl.X.CompressGZIP(file_VNExpress_selector_proto_rawDescData)
	})
	return file_VNExpress_selector_proto_rawDescData
}

var file_VNExpress_selector_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_VNExpress_selector_proto_goTypes = []interface{}{
	(*Range)(nil),   // 0: vnexpress_selector.Range
	(*News)(nil),    // 1: vnexpress_selector.News
	(*Podcast)(nil), // 2: vnexpress_selector.Podcast
}
var file_VNExpress_selector_proto_depIdxs = []int32{
	0, // 0: vnexpress_selector.VNExpress_selector.Select_news:input_type -> vnexpress_selector.Range
	0, // 1: vnexpress_selector.VNExpress_selector.Select_podcast:input_type -> vnexpress_selector.Range
	1, // 2: vnexpress_selector.VNExpress_selector.Select_news:output_type -> vnexpress_selector.News
	2, // 3: vnexpress_selector.VNExpress_selector.Select_podcast:output_type -> vnexpress_selector.Podcast
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_VNExpress_selector_proto_init() }
func file_VNExpress_selector_proto_init() {
	if File_VNExpress_selector_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_VNExpress_selector_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Range); i {
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
		file_VNExpress_selector_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*News); i {
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
		file_VNExpress_selector_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Podcast); i {
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
			RawDescriptor: file_VNExpress_selector_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_VNExpress_selector_proto_goTypes,
		DependencyIndexes: file_VNExpress_selector_proto_depIdxs,
		MessageInfos:      file_VNExpress_selector_proto_msgTypes,
	}.Build()
	File_VNExpress_selector_proto = out.File
	file_VNExpress_selector_proto_rawDesc = nil
	file_VNExpress_selector_proto_goTypes = nil
	file_VNExpress_selector_proto_depIdxs = nil
}
