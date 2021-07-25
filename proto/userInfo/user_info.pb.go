// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: proto/userInfo/user_info.proto

package pbUserInfo

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

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid            string  `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	OpsAccount     string  `protobuf:"bytes,2,opt,name=ops_account,proto3" json:"ops_account,omitempty"`
	NickName       string  `protobuf:"bytes,3,opt,name=nick_name,proto3" json:"nick_name,omitempty"`
	Description    string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Link           string  `protobuf:"bytes,5,opt,name=link,proto3" json:"link,omitempty"`
	RegisterTime   string  `protobuf:"bytes,6,opt,name=register_time,proto3" json:"register_time,omitempty"`
	Profile        string  `protobuf:"bytes,7,opt,name=profile,proto3" json:"profile,omitempty"`
	Banner         string  `protobuf:"bytes,8,opt,name=banner,proto3" json:"banner,omitempty"`
	OldProfileLink string  `protobuf:"bytes,9,opt,name=old_profile_link,proto3" json:"old_profile_link,omitempty"`
	OldBannerLink  string  `protobuf:"bytes,10,opt,name=old_banner_link,proto3" json:"old_banner_link,omitempty"`
	Location       string  `protobuf:"bytes,11,opt,name=location,proto3" json:"location,omitempty"`
	IsOfficialUser bool    `protobuf:"varint,12,opt,name=is_official_user,proto3" json:"is_official_user,omitempty"`
	Label          string  `protobuf:"bytes,13,opt,name=label,proto3" json:"label,omitempty"`
	IsPrivate      bool    `protobuf:"varint,14,opt,name=is_private,proto3" json:"is_private,omitempty"`
	OopNumber      int64   `protobuf:"varint,15,opt,name=oop_number,proto3" json:"oop_number,omitempty"`
	TotalAmount    int64   `protobuf:"varint,16,opt,name=total_amount,proto3" json:"total_amount,omitempty"`
	Price          float64 `protobuf:"fixed64,17,opt,name=price,proto3" json:"price,omitempty"`
	Following      int64   `protobuf:"varint,18,opt,name=following,proto3" json:"following,omitempty"`
	Followed       int64   `protobuf:"varint,19,opt,name=followed,proto3" json:"followed,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_userInfo_user_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_userInfo_user_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_proto_userInfo_user_info_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *UserInfo) GetOpsAccount() string {
	if x != nil {
		return x.OpsAccount
	}
	return ""
}

func (x *UserInfo) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *UserInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UserInfo) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *UserInfo) GetRegisterTime() string {
	if x != nil {
		return x.RegisterTime
	}
	return ""
}

func (x *UserInfo) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *UserInfo) GetBanner() string {
	if x != nil {
		return x.Banner
	}
	return ""
}

func (x *UserInfo) GetOldProfileLink() string {
	if x != nil {
		return x.OldProfileLink
	}
	return ""
}

func (x *UserInfo) GetOldBannerLink() string {
	if x != nil {
		return x.OldBannerLink
	}
	return ""
}

func (x *UserInfo) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *UserInfo) GetIsOfficialUser() bool {
	if x != nil {
		return x.IsOfficialUser
	}
	return false
}

func (x *UserInfo) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *UserInfo) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

func (x *UserInfo) GetOopNumber() int64 {
	if x != nil {
		return x.OopNumber
	}
	return 0
}

func (x *UserInfo) GetTotalAmount() int64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *UserInfo) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UserInfo) GetFollowing() int64 {
	if x != nil {
		return x.Following
	}
	return 0
}

func (x *UserInfo) GetFollowed() int64 {
	if x != nil {
		return x.Followed
	}
	return 0
}

type SetPrivate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid       string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	IsPrivate bool   `protobuf:"varint,2,opt,name=is_private,proto3" json:"is_private,omitempty"`
}

func (x *SetPrivate) Reset() {
	*x = SetPrivate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_userInfo_user_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPrivate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPrivate) ProtoMessage() {}

func (x *SetPrivate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_userInfo_user_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPrivate.ProtoReflect.Descriptor instead.
func (*SetPrivate) Descriptor() ([]byte, []int) {
	return file_proto_userInfo_user_info_proto_rawDescGZIP(), []int{1}
}

func (x *SetPrivate) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *SetPrivate) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

type OperateResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *OperateResult) Reset() {
	*x = OperateResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_userInfo_user_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperateResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperateResult) ProtoMessage() {}

func (x *OperateResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_userInfo_user_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperateResult.ProtoReflect.Descriptor instead.
func (*OperateResult) Descriptor() ([]byte, []int) {
	return file_proto_userInfo_user_info_proto_rawDescGZIP(), []int{2}
}

func (x *OperateResult) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type QueryAndDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *QueryAndDelete) Reset() {
	*x = QueryAndDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_userInfo_user_info_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAndDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAndDelete) ProtoMessage() {}

func (x *QueryAndDelete) ProtoReflect() protoreflect.Message {
	mi := &file_proto_userInfo_user_info_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAndDelete.ProtoReflect.Descriptor instead.
func (*QueryAndDelete) Descriptor() ([]byte, []int) {
	return file_proto_userInfo_user_info_proto_rawDescGZIP(), []int{3}
}

func (x *QueryAndDelete) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type CheckOpsAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpsAccount string `protobuf:"bytes,1,opt,name=ops_account,proto3" json:"ops_account,omitempty"`
}

func (x *CheckOpsAccount) Reset() {
	*x = CheckOpsAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_userInfo_user_info_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckOpsAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckOpsAccount) ProtoMessage() {}

func (x *CheckOpsAccount) ProtoReflect() protoreflect.Message {
	mi := &file_proto_userInfo_user_info_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckOpsAccount.ProtoReflect.Descriptor instead.
func (*CheckOpsAccount) Descriptor() ([]byte, []int) {
	return file_proto_userInfo_user_info_proto_rawDescGZIP(), []int{4}
}

func (x *CheckOpsAccount) GetOpsAccount() string {
	if x != nil {
		return x.OpsAccount
	}
	return ""
}

type OopNum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Num int64  `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *OopNum) Reset() {
	*x = OopNum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_userInfo_user_info_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OopNum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OopNum) ProtoMessage() {}

func (x *OopNum) ProtoReflect() protoreflect.Message {
	mi := &file_proto_userInfo_user_info_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OopNum.ProtoReflect.Descriptor instead.
func (*OopNum) Descriptor() ([]byte, []int) {
	return file_proto_userInfo_user_info_proto_rawDescGZIP(), []int{5}
}

func (x *OopNum) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *OopNum) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid   string  `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Price float64 `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_userInfo_user_info_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_proto_userInfo_user_info_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_proto_userInfo_user_info_proto_rawDescGZIP(), []int{6}
}

func (x *Price) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Price) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type TotalReleaseOCardNum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid            string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ReleaseCardNum int64  `protobuf:"varint,2,opt,name=release_card_num,proto3" json:"release_card_num,omitempty"`
}

func (x *TotalReleaseOCardNum) Reset() {
	*x = TotalReleaseOCardNum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_userInfo_user_info_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TotalReleaseOCardNum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TotalReleaseOCardNum) ProtoMessage() {}

func (x *TotalReleaseOCardNum) ProtoReflect() protoreflect.Message {
	mi := &file_proto_userInfo_user_info_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TotalReleaseOCardNum.ProtoReflect.Descriptor instead.
func (*TotalReleaseOCardNum) Descriptor() ([]byte, []int) {
	return file_proto_userInfo_user_info_proto_rawDescGZIP(), []int{7}
}

func (x *TotalReleaseOCardNum) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *TotalReleaseOCardNum) GetReleaseCardNum() int64 {
	if x != nil {
		return x.ReleaseCardNum
	}
	return 0
}

type OpsAccounts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpsAccounts []string              `protobuf:"bytes,1,rep,name=ops_accounts,proto3" json:"ops_accounts,omitempty"`
	OpsAndUid   []*OpsAccountsAndUids `protobuf:"bytes,2,rep,name=opsAndUid,proto3" json:"opsAndUid,omitempty"`
}

func (x *OpsAccounts) Reset() {
	*x = OpsAccounts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_userInfo_user_info_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpsAccounts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpsAccounts) ProtoMessage() {}

func (x *OpsAccounts) ProtoReflect() protoreflect.Message {
	mi := &file_proto_userInfo_user_info_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpsAccounts.ProtoReflect.Descriptor instead.
func (*OpsAccounts) Descriptor() ([]byte, []int) {
	return file_proto_userInfo_user_info_proto_rawDescGZIP(), []int{8}
}

func (x *OpsAccounts) GetOpsAccounts() []string {
	if x != nil {
		return x.OpsAccounts
	}
	return nil
}

func (x *OpsAccounts) GetOpsAndUid() []*OpsAccountsAndUids {
	if x != nil {
		return x.OpsAndUid
	}
	return nil
}

type OpsAccountsAndUids struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpsAccount string `protobuf:"bytes,1,opt,name=ops_account,proto3" json:"ops_account,omitempty"`
	Uid        string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *OpsAccountsAndUids) Reset() {
	*x = OpsAccountsAndUids{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_userInfo_user_info_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpsAccountsAndUids) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpsAccountsAndUids) ProtoMessage() {}

func (x *OpsAccountsAndUids) ProtoReflect() protoreflect.Message {
	mi := &file_proto_userInfo_user_info_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpsAccountsAndUids.ProtoReflect.Descriptor instead.
func (*OpsAccountsAndUids) Descriptor() ([]byte, []int) {
	return file_proto_userInfo_user_info_proto_rawDescGZIP(), []int{9}
}

func (x *OpsAccountsAndUids) GetOpsAccount() string {
	if x != nil {
		return x.OpsAccount
	}
	return ""
}

func (x *OpsAccountsAndUids) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type UidAndPbkaddr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Pubkaddr string `protobuf:"bytes,2,opt,name=pubkaddr,proto3" json:"pubkaddr,omitempty"`
	ErrCode  string `protobuf:"bytes,3,opt,name=err_code,proto3" json:"err_code,omitempty"`
}

func (x *UidAndPbkaddr) Reset() {
	*x = UidAndPbkaddr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_userInfo_user_info_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UidAndPbkaddr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UidAndPbkaddr) ProtoMessage() {}

func (x *UidAndPbkaddr) ProtoReflect() protoreflect.Message {
	mi := &file_proto_userInfo_user_info_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UidAndPbkaddr.ProtoReflect.Descriptor instead.
func (*UidAndPbkaddr) Descriptor() ([]byte, []int) {
	return file_proto_userInfo_user_info_proto_rawDescGZIP(), []int{10}
}

func (x *UidAndPbkaddr) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *UidAndPbkaddr) GetPubkaddr() string {
	if x != nil {
		return x.Pubkaddr
	}
	return ""
}

func (x *UidAndPbkaddr) GetErrCode() string {
	if x != nil {
		return x.ErrCode
	}
	return ""
}

var File_proto_userInfo_user_info_proto protoreflect.FileDescriptor

var file_proto_userInfo_user_info_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xd2, 0x04, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x70, 0x73,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6f, 0x70, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x69, 0x6e, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12,
	0x24, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x6f, 0x6c, 0x64, 0x5f, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x6f, 0x6c, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6c,
	0x69, 0x6e, 0x6b, 0x12, 0x28, 0x0a, 0x0f, 0x6f, 0x6c, 0x64, 0x5f, 0x62, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x6c,
	0x64, 0x5f, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x69, 0x73, 0x5f,
	0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x69,
	0x73, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6f,
	0x6f, 0x70, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x6f, 0x6f, 0x70, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69,
	0x6e, 0x67, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x22,
	0x3e, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x22,
	0x23, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x22, 0x22, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6e, 0x64,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x33, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x4f, 0x70, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6f,
	0x70, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2c, 0x0a,
	0x06, 0x4f, 0x6f, 0x70, 0x4e, 0x75, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x2f, 0x0a, 0x05, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x54, 0x0a, 0x14,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4f, 0x43, 0x61, 0x72,
	0x64, 0x4e, 0x75, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x10, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6e,
	0x75, 0x6d, 0x22, 0x6d, 0x0a, 0x0b, 0x4f, 0x70, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x3a, 0x0a, 0x09, 0x6f, 0x70, 0x73, 0x41, 0x6e, 0x64, 0x55,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x4f, 0x70, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x41,
	0x6e, 0x64, 0x55, 0x69, 0x64, 0x73, 0x52, 0x09, 0x6f, 0x70, 0x73, 0x41, 0x6e, 0x64, 0x55, 0x69,
	0x64, 0x22, 0x48, 0x0a, 0x12, 0x4f, 0x70, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x41, 0x6e, 0x64, 0x55, 0x69, 0x64, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x70, 0x73, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70,
	0x73, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x59, 0x0a, 0x0d, 0x55,
	0x69, 0x64, 0x41, 0x6e, 0x64, 0x50, 0x62, 0x6b, 0x61, 0x64, 0x64, 0x72, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x75, 0x62, 0x6b, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x75, 0x62, 0x6b, 0x61, 0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x72,
	0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x72,
	0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x32, 0xf6, 0x05, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3e, 0x0a, 0x0d, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a,
	0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6e, 0x64, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x3f, 0x0a,
	0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41,
	0x6e, 0x64, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x4c,
	0x0a, 0x14, 0x49, 0x73, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x70, 0x73, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x70, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0e,
	0x4f, 0x6e, 0x65, 0x53, 0x74, 0x65, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x12, 0x14,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12,
	0x3b, 0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x4f, 0x6f, 0x70, 0x4e, 0x75, 0x6d, 0x12,
	0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4f, 0x6f, 0x70, 0x4e, 0x75,
	0x6d, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x08,
	0x53, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x1a, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x17,
	0x53, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4f,
	0x43, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x12, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4f,
	0x43, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x00, 0x12, 0x44, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x69, 0x64, 0x42, 0x79, 0x4f, 0x70,
	0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x4f, 0x70, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x1a,
	0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4f, 0x70, 0x73, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50,
	0x62, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x42, 0x79, 0x55, 0x69, 0x64, 0x12, 0x17, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x69, 0x64, 0x41, 0x6e, 0x64, 0x50, 0x62, 0x6b,
	0x61, 0x64, 0x64, 0x72, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x55, 0x69, 0x64, 0x41, 0x6e, 0x64, 0x50, 0x62, 0x6b, 0x61, 0x64, 0x64, 0x72, 0x22, 0x00, 0x42,
	0x1d, 0x5a, 0x1b, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x3b, 0x70, 0x62, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_userInfo_user_info_proto_rawDescOnce sync.Once
	file_proto_userInfo_user_info_proto_rawDescData = file_proto_userInfo_user_info_proto_rawDesc
)

func file_proto_userInfo_user_info_proto_rawDescGZIP() []byte {
	file_proto_userInfo_user_info_proto_rawDescOnce.Do(func() {
		file_proto_userInfo_user_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_userInfo_user_info_proto_rawDescData)
	})
	return file_proto_userInfo_user_info_proto_rawDescData
}

var file_proto_userInfo_user_info_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_userInfo_user_info_proto_goTypes = []interface{}{
	(*UserInfo)(nil),             // 0: userInfo.UserInfo
	(*SetPrivate)(nil),           // 1: userInfo.SetPrivate
	(*OperateResult)(nil),        // 2: userInfo.OperateResult
	(*QueryAndDelete)(nil),       // 3: userInfo.QueryAndDelete
	(*CheckOpsAccount)(nil),      // 4: userInfo.CheckOpsAccount
	(*OopNum)(nil),               // 5: userInfo.OopNum
	(*Price)(nil),                // 6: userInfo.Price
	(*TotalReleaseOCardNum)(nil), // 7: userInfo.TotalReleaseOCardNum
	(*OpsAccounts)(nil),          // 8: userInfo.OpsAccounts
	(*OpsAccountsAndUids)(nil),   // 9: userInfo.OpsAccountsAndUids
	(*UidAndPbkaddr)(nil),        // 10: userInfo.UidAndPbkaddr
}
var file_proto_userInfo_user_info_proto_depIdxs = []int32{
	9,  // 0: userInfo.OpsAccounts.opsAndUid:type_name -> userInfo.OpsAccountsAndUids
	0,  // 1: userInfo.OperateUserInfo.StoreUserInfo:input_type -> userInfo.UserInfo
	0,  // 2: userInfo.OperateUserInfo.UpdateUserInfo:input_type -> userInfo.UserInfo
	3,  // 3: userInfo.OperateUserInfo.DeleteUserInfo:input_type -> userInfo.QueryAndDelete
	3,  // 4: userInfo.OperateUserInfo.QueryUserInfo:input_type -> userInfo.QueryAndDelete
	4,  // 5: userInfo.OperateUserInfo.IsRepeatedOpsAccount:input_type -> userInfo.CheckOpsAccount
	1,  // 6: userInfo.OperateUserInfo.OneStepProtect:input_type -> userInfo.SetPrivate
	5,  // 7: userInfo.OperateUserInfo.ModifyOopNum:input_type -> userInfo.OopNum
	6,  // 8: userInfo.OperateUserInfo.SetPrice:input_type -> userInfo.Price
	7,  // 9: userInfo.OperateUserInfo.SetTotalReleaseOCardNum:input_type -> userInfo.TotalReleaseOCardNum
	8,  // 10: userInfo.OperateUserInfo.GetUidByOpsAccount:input_type -> userInfo.OpsAccounts
	10, // 11: userInfo.OperateUserInfo.GetPbkAddrByUid:input_type -> userInfo.UidAndPbkaddr
	2,  // 12: userInfo.OperateUserInfo.StoreUserInfo:output_type -> userInfo.OperateResult
	0,  // 13: userInfo.OperateUserInfo.UpdateUserInfo:output_type -> userInfo.UserInfo
	2,  // 14: userInfo.OperateUserInfo.DeleteUserInfo:output_type -> userInfo.OperateResult
	0,  // 15: userInfo.OperateUserInfo.QueryUserInfo:output_type -> userInfo.UserInfo
	2,  // 16: userInfo.OperateUserInfo.IsRepeatedOpsAccount:output_type -> userInfo.OperateResult
	2,  // 17: userInfo.OperateUserInfo.OneStepProtect:output_type -> userInfo.OperateResult
	2,  // 18: userInfo.OperateUserInfo.ModifyOopNum:output_type -> userInfo.OperateResult
	6,  // 19: userInfo.OperateUserInfo.SetPrice:output_type -> userInfo.Price
	2,  // 20: userInfo.OperateUserInfo.SetTotalReleaseOCardNum:output_type -> userInfo.OperateResult
	8,  // 21: userInfo.OperateUserInfo.GetUidByOpsAccount:output_type -> userInfo.OpsAccounts
	10, // 22: userInfo.OperateUserInfo.GetPbkAddrByUid:output_type -> userInfo.UidAndPbkaddr
	12, // [12:23] is the sub-list for method output_type
	1,  // [1:12] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_proto_userInfo_user_info_proto_init() }
func file_proto_userInfo_user_info_proto_init() {
	if File_proto_userInfo_user_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_userInfo_user_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_proto_userInfo_user_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPrivate); i {
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
		file_proto_userInfo_user_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperateResult); i {
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
		file_proto_userInfo_user_info_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAndDelete); i {
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
		file_proto_userInfo_user_info_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckOpsAccount); i {
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
		file_proto_userInfo_user_info_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OopNum); i {
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
		file_proto_userInfo_user_info_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Price); i {
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
		file_proto_userInfo_user_info_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TotalReleaseOCardNum); i {
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
		file_proto_userInfo_user_info_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpsAccounts); i {
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
		file_proto_userInfo_user_info_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpsAccountsAndUids); i {
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
		file_proto_userInfo_user_info_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UidAndPbkaddr); i {
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
			RawDescriptor: file_proto_userInfo_user_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_userInfo_user_info_proto_goTypes,
		DependencyIndexes: file_proto_userInfo_user_info_proto_depIdxs,
		MessageInfos:      file_proto_userInfo_user_info_proto_msgTypes,
	}.Build()
	File_proto_userInfo_user_info_proto = out.File
	file_proto_userInfo_user_info_proto_rawDesc = nil
	file_proto_userInfo_user_info_proto_goTypes = nil
	file_proto_userInfo_user_info_proto_depIdxs = nil
}
