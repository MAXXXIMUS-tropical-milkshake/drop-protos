// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.29.1
// source: audio/audio.proto

package audiov1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type NoteUpload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Note  string `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
	Scale string `protobuf:"bytes,2,opt,name=scale,proto3" json:"scale,omitempty"`
}

func (x *NoteUpload) Reset() {
	*x = NoteUpload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteUpload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteUpload) ProtoMessage() {}

func (x *NoteUpload) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteUpload.ProtoReflect.Descriptor instead.
func (*NoteUpload) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{0}
}

func (x *NoteUpload) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *NoteUpload) GetScale() string {
	if x != nil {
		return x.Scale
	}
	return ""
}

type UploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BeatId      int64       `protobuf:"varint,1,opt,name=beat_id,json=beatId,proto3" json:"beat_id,omitempty"`
	BeatmakerId int64       `protobuf:"varint,2,opt,name=beatmaker_id,json=beatmakerId,proto3" json:"beatmaker_id,omitempty"`
	Name        string      `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string      `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	BeatGenre   []int64     `protobuf:"varint,5,rep,packed,name=beat_genre,json=beatGenre,proto3" json:"beat_genre,omitempty"`
	BeatTag     []int64     `protobuf:"varint,6,rep,packed,name=beat_tag,json=beatTag,proto3" json:"beat_tag,omitempty"`
	BeatMood    []int64     `protobuf:"varint,7,rep,packed,name=beat_mood,json=beatMood,proto3" json:"beat_mood,omitempty"`
	Note        *NoteUpload `protobuf:"bytes,8,opt,name=note,proto3" json:"note,omitempty"`
	Bpm         int64       `protobuf:"varint,9,opt,name=bpm,proto3" json:"bpm,omitempty"`
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{1}
}

func (x *UploadRequest) GetBeatId() int64 {
	if x != nil {
		return x.BeatId
	}
	return 0
}

func (x *UploadRequest) GetBeatmakerId() int64 {
	if x != nil {
		return x.BeatmakerId
	}
	return 0
}

func (x *UploadRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UploadRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UploadRequest) GetBeatGenre() []int64 {
	if x != nil {
		return x.BeatGenre
	}
	return nil
}

func (x *UploadRequest) GetBeatTag() []int64 {
	if x != nil {
		return x.BeatTag
	}
	return nil
}

func (x *UploadRequest) GetBeatMood() []int64 {
	if x != nil {
		return x.BeatMood
	}
	return nil
}

func (x *UploadRequest) GetNote() *NoteUpload {
	if x != nil {
		return x.Note
	}
	return nil
}

func (x *UploadRequest) GetBpm() int64 {
	if x != nil {
		return x.Bpm
	}
	return 0
}

type UploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileUploadUrl  string `protobuf:"bytes,1,opt,name=file_upload_url,json=fileUploadUrl,proto3" json:"file_upload_url,omitempty"`
	ImageUploadUrl string `protobuf:"bytes,2,opt,name=image_upload_url,json=imageUploadUrl,proto3" json:"image_upload_url,omitempty"`
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponse.ProtoReflect.Descriptor instead.
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{2}
}

func (x *UploadResponse) GetFileUploadUrl() string {
	if x != nil {
		return x.FileUploadUrl
	}
	return ""
}

func (x *UploadResponse) GetImageUploadUrl() string {
	if x != nil {
		return x.ImageUploadUrl
	}
	return ""
}

type GetBeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BeatId int64 `protobuf:"varint,1,opt,name=beat_id,json=beatId,proto3" json:"beat_id,omitempty"`
}

func (x *GetBeatRequest) Reset() {
	*x = GetBeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBeatRequest) ProtoMessage() {}

func (x *GetBeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBeatRequest.ProtoReflect.Descriptor instead.
func (*GetBeatRequest) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{3}
}

func (x *GetBeatRequest) GetBeatId() int64 {
	if x != nil {
		return x.BeatId
	}
	return 0
}

type Beatmaker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Pseudonym string `protobuf:"bytes,3,opt,name=pseudonym,proto3" json:"pseudonym,omitempty"`
}

func (x *Beatmaker) Reset() {
	*x = Beatmaker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Beatmaker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Beatmaker) ProtoMessage() {}

func (x *Beatmaker) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Beatmaker.ProtoReflect.Descriptor instead.
func (*Beatmaker) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{4}
}

func (x *Beatmaker) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Beatmaker) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Beatmaker) GetPseudonym() string {
	if x != nil {
		return x.Pseudonym
	}
	return ""
}

type GetBeatNote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Scale string `protobuf:"bytes,2,opt,name=scale,proto3" json:"scale,omitempty"`
}

func (x *GetBeatNote) Reset() {
	*x = GetBeatNote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBeatNote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBeatNote) ProtoMessage() {}

func (x *GetBeatNote) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBeatNote.ProtoReflect.Descriptor instead.
func (*GetBeatNote) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{5}
}

func (x *GetBeatNote) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetBeatNote) GetScale() string {
	if x != nil {
		return x.Scale
	}
	return ""
}

type GetBeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string       `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Beatmaker   *Beatmaker   `protobuf:"bytes,4,opt,name=beatmaker,proto3" json:"beatmaker,omitempty"`
	BeatGenre   []string     `protobuf:"bytes,5,rep,name=beat_genre,json=beatGenre,proto3" json:"beat_genre,omitempty"`
	BeatTag     []string     `protobuf:"bytes,6,rep,name=beat_tag,json=beatTag,proto3" json:"beat_tag,omitempty"`
	BeatMood    []string     `protobuf:"bytes,7,rep,name=beat_mood,json=beatMood,proto3" json:"beat_mood,omitempty"`
	Note        *GetBeatNote `protobuf:"bytes,8,opt,name=note,proto3" json:"note,omitempty"`
	Bpm         int64        `protobuf:"varint,9,opt,name=bpm,proto3" json:"bpm,omitempty"`
}

func (x *GetBeatResponse) Reset() {
	*x = GetBeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBeatResponse) ProtoMessage() {}

func (x *GetBeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBeatResponse.ProtoReflect.Descriptor instead.
func (*GetBeatResponse) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{6}
}

func (x *GetBeatResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetBeatResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetBeatResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetBeatResponse) GetBeatmaker() *Beatmaker {
	if x != nil {
		return x.Beatmaker
	}
	return nil
}

func (x *GetBeatResponse) GetBeatGenre() []string {
	if x != nil {
		return x.BeatGenre
	}
	return nil
}

func (x *GetBeatResponse) GetBeatTag() []string {
	if x != nil {
		return x.BeatTag
	}
	return nil
}

func (x *GetBeatResponse) GetBeatMood() []string {
	if x != nil {
		return x.BeatMood
	}
	return nil
}

func (x *GetBeatResponse) GetNote() *GetBeatNote {
	if x != nil {
		return x.Note
	}
	return nil
}

func (x *GetBeatResponse) GetBpm() int64 {
	if x != nil {
		return x.Bpm
	}
	return 0
}

type GetFiltersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFiltersRequest) Reset() {
	*x = GetFiltersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFiltersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFiltersRequest) ProtoMessage() {}

func (x *GetFiltersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFiltersRequest.ProtoReflect.Descriptor instead.
func (*GetFiltersRequest) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{7}
}

type GetFiltersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Genres []*Genre `protobuf:"bytes,1,rep,name=genres,proto3" json:"genres,omitempty"`
	Tags   []*Tag   `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	Moods  []*Mood  `protobuf:"bytes,3,rep,name=moods,proto3" json:"moods,omitempty"`
	Notes  []*Note  `protobuf:"bytes,4,rep,name=notes,proto3" json:"notes,omitempty"`
}

func (x *GetFiltersResponse) Reset() {
	*x = GetFiltersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFiltersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFiltersResponse) ProtoMessage() {}

func (x *GetFiltersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFiltersResponse.ProtoReflect.Descriptor instead.
func (*GetFiltersResponse) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{8}
}

func (x *GetFiltersResponse) GetGenres() []*Genre {
	if x != nil {
		return x.Genres
	}
	return nil
}

func (x *GetFiltersResponse) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *GetFiltersResponse) GetMoods() []*Mood {
	if x != nil {
		return x.Moods
	}
	return nil
}

func (x *GetFiltersResponse) GetNotes() []*Note {
	if x != nil {
		return x.Notes
	}
	return nil
}

type Genre struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GenreId int64  `protobuf:"varint,1,opt,name=genre_id,json=genreId,proto3" json:"genre_id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Genre) Reset() {
	*x = Genre{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Genre) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Genre) ProtoMessage() {}

func (x *Genre) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Genre.ProtoReflect.Descriptor instead.
func (*Genre) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{9}
}

func (x *Genre) GetGenreId() int64 {
	if x != nil {
		return x.GenreId
	}
	return 0
}

func (x *Genre) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TagId int64  `protobuf:"varint,1,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{10}
}

func (x *Tag) GetTagId() int64 {
	if x != nil {
		return x.TagId
	}
	return 0
}

func (x *Tag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Mood struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MoodId int64  `protobuf:"varint,1,opt,name=mood_id,json=moodId,proto3" json:"mood_id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Mood) Reset() {
	*x = Mood{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mood) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mood) ProtoMessage() {}

func (x *Mood) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mood.ProtoReflect.Descriptor instead.
func (*Mood) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{11}
}

func (x *Mood) GetMoodId() int64 {
	if x != nil {
		return x.MoodId
	}
	return 0
}

func (x *Mood) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Note struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoteId int64  `protobuf:"varint,1,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Note) Reset() {
	*x = Note{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audio_audio_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Note) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Note) ProtoMessage() {}

func (x *Note) ProtoReflect() protoreflect.Message {
	mi := &file_audio_audio_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Note.ProtoReflect.Descriptor instead.
func (*Note) Descriptor() ([]byte, []int) {
	return file_audio_audio_proto_rawDescGZIP(), []int{12}
}

func (x *Note) GetNoteId() int64 {
	if x != nil {
		return x.NoteId
	}
	return 0
}

func (x *Note) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_audio_audio_proto protoreflect.FileDescriptor

var file_audio_audio_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x0a, 0x4e, 0x6f, 0x74, 0x65,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x22, 0x91, 0x02, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x65, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x62, 0x65, 0x61, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x62,
	0x65, 0x61, 0x74, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x62, 0x65, 0x61, 0x74, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x61, 0x74, 0x5f, 0x67, 0x65, 0x6e,
	0x72, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x03, 0x52, 0x09, 0x62, 0x65, 0x61, 0x74, 0x47, 0x65,
	0x6e, 0x72, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x65, 0x61, 0x74, 0x5f, 0x74, 0x61, 0x67, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x62, 0x65, 0x61, 0x74, 0x54, 0x61, 0x67, 0x12, 0x1b,
	0x0a, 0x09, 0x62, 0x65, 0x61, 0x74, 0x5f, 0x6d, 0x6f, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x08, 0x62, 0x65, 0x61, 0x74, 0x4d, 0x6f, 0x6f, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x6e,
	0x6f, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x75, 0x64, 0x69,
	0x6f, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x04, 0x6e, 0x6f,
	0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x70, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x62, 0x70, 0x6d, 0x22, 0x62, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x28,
	0x0a, 0x10, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42,
	0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x65,
	0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x62, 0x65, 0x61,
	0x74, 0x49, 0x64, 0x22, 0x55, 0x0a, 0x09, 0x42, 0x65, 0x61, 0x74, 0x6d, 0x61, 0x6b, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x73, 0x65, 0x75, 0x64, 0x6f, 0x6e, 0x79, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x73, 0x65, 0x75, 0x64, 0x6f, 0x6e, 0x79, 0x6d, 0x22, 0x37, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x42, 0x65, 0x61, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x22, 0x98, 0x02, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a,
	0x09, 0x62, 0x65, 0x61, 0x74, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x42, 0x65, 0x61, 0x74, 0x6d, 0x61, 0x6b,
	0x65, 0x72, 0x52, 0x09, 0x62, 0x65, 0x61, 0x74, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x12, 0x1d, 0x0a,
	0x0a, 0x62, 0x65, 0x61, 0x74, 0x5f, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x62, 0x65, 0x61, 0x74, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x62, 0x65, 0x61, 0x74, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x62, 0x65, 0x61, 0x74, 0x54, 0x61, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x65, 0x61, 0x74, 0x5f,
	0x6d, 0x6f, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x62, 0x65, 0x61, 0x74,
	0x4d, 0x6f, 0x6f, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x65,
	0x61, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x62, 0x70, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x62, 0x70, 0x6d, 0x22, 0x13,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0xa0, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x67, 0x65,
	0x6e, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x75, 0x64,
	0x69, 0x6f, 0x2e, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73,
	0x12, 0x1e, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x12, 0x21, 0x0a, 0x05, 0x6d, 0x6f, 0x6f, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x4d, 0x6f, 0x6f, 0x64, 0x52, 0x05, 0x6d, 0x6f,
	0x6f, 0x64, 0x73, 0x12, 0x21, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52,
	0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x36, 0x0a, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x30,
	0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x33, 0x0a, 0x04, 0x4d, 0x6f, 0x6f, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x6f, 0x6f, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6d, 0x6f, 0x6f, 0x64, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x04, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x6e, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x90, 0x02, 0x0a, 0x0c, 0x41,
	0x75, 0x64, 0x69, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x14, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x75,
	0x64, 0x69, 0x6f, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x55, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42,
	0x65, 0x61, 0x74, 0x12, 0x15, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x75, 0x64,
	0x69, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x75, 0x64, 0x69, 0x6f, 0x2f, 0x7b, 0x62, 0x65, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12,
	0x5c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x18, 0x2e,
	0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x75, 0x64, 0x69, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x42, 0x41, 0x5a,
	0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x41, 0x58, 0x58,
	0x58, 0x49, 0x4d, 0x55, 0x53, 0x2d, 0x74, 0x72, 0x6f, 0x70, 0x69, 0x63, 0x61, 0x6c, 0x2d, 0x6d,
	0x69, 0x6c, 0x6b, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x2f, 0x62, 0x65, 0x61, 0x74, 0x66, 0x6c, 0x6f,
	0x77, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x3b, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_audio_audio_proto_rawDescOnce sync.Once
	file_audio_audio_proto_rawDescData = file_audio_audio_proto_rawDesc
)

func file_audio_audio_proto_rawDescGZIP() []byte {
	file_audio_audio_proto_rawDescOnce.Do(func() {
		file_audio_audio_proto_rawDescData = protoimpl.X.CompressGZIP(file_audio_audio_proto_rawDescData)
	})
	return file_audio_audio_proto_rawDescData
}

var file_audio_audio_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_audio_audio_proto_goTypes = []any{
	(*NoteUpload)(nil),         // 0: audio.NoteUpload
	(*UploadRequest)(nil),      // 1: audio.UploadRequest
	(*UploadResponse)(nil),     // 2: audio.UploadResponse
	(*GetBeatRequest)(nil),     // 3: audio.GetBeatRequest
	(*Beatmaker)(nil),          // 4: audio.Beatmaker
	(*GetBeatNote)(nil),        // 5: audio.GetBeatNote
	(*GetBeatResponse)(nil),    // 6: audio.GetBeatResponse
	(*GetFiltersRequest)(nil),  // 7: audio.GetFiltersRequest
	(*GetFiltersResponse)(nil), // 8: audio.GetFiltersResponse
	(*Genre)(nil),              // 9: audio.Genre
	(*Tag)(nil),                // 10: audio.Tag
	(*Mood)(nil),               // 11: audio.Mood
	(*Note)(nil),               // 12: audio.Note
}
var file_audio_audio_proto_depIdxs = []int32{
	0,  // 0: audio.UploadRequest.note:type_name -> audio.NoteUpload
	4,  // 1: audio.GetBeatResponse.beatmaker:type_name -> audio.Beatmaker
	5,  // 2: audio.GetBeatResponse.note:type_name -> audio.GetBeatNote
	9,  // 3: audio.GetFiltersResponse.genres:type_name -> audio.Genre
	10, // 4: audio.GetFiltersResponse.tags:type_name -> audio.Tag
	11, // 5: audio.GetFiltersResponse.moods:type_name -> audio.Mood
	12, // 6: audio.GetFiltersResponse.notes:type_name -> audio.Note
	1,  // 7: audio.AudioService.Upload:input_type -> audio.UploadRequest
	3,  // 8: audio.AudioService.GetBeat:input_type -> audio.GetBeatRequest
	7,  // 9: audio.AudioService.GetFilters:input_type -> audio.GetFiltersRequest
	2,  // 10: audio.AudioService.Upload:output_type -> audio.UploadResponse
	6,  // 11: audio.AudioService.GetBeat:output_type -> audio.GetBeatResponse
	8,  // 12: audio.AudioService.GetFilters:output_type -> audio.GetFiltersResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_audio_audio_proto_init() }
func file_audio_audio_proto_init() {
	if File_audio_audio_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_audio_audio_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*NoteUpload); i {
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
		file_audio_audio_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UploadRequest); i {
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
		file_audio_audio_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UploadResponse); i {
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
		file_audio_audio_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetBeatRequest); i {
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
		file_audio_audio_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Beatmaker); i {
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
		file_audio_audio_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetBeatNote); i {
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
		file_audio_audio_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetBeatResponse); i {
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
		file_audio_audio_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetFiltersRequest); i {
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
		file_audio_audio_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetFiltersResponse); i {
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
		file_audio_audio_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*Genre); i {
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
		file_audio_audio_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*Tag); i {
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
		file_audio_audio_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*Mood); i {
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
		file_audio_audio_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*Note); i {
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
			RawDescriptor: file_audio_audio_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_audio_audio_proto_goTypes,
		DependencyIndexes: file_audio_audio_proto_depIdxs,
		MessageInfos:      file_audio_audio_proto_msgTypes,
	}.Build()
	File_audio_audio_proto = out.File
	file_audio_audio_proto_rawDesc = nil
	file_audio_audio_proto_goTypes = nil
	file_audio_audio_proto_depIdxs = nil
}
