// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account_block_db.proto

package vitepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AccountBlockDb struct {
	To                   []byte   `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	PrevHash             []byte   `protobuf:"bytes,2,opt,name=prevHash,proto3" json:"prevHash,omitempty"`
	FromHash             []byte   `protobuf:"bytes,3,opt,name=fromHash,proto3" json:"fromHash,omitempty"`
	TokenId              []byte   `protobuf:"bytes,4,opt,name=tokenId,proto3" json:"tokenId,omitempty"`
	Balance              []byte   `protobuf:"bytes,5,opt,name=balance,proto3" json:"balance,omitempty"`
	Data                 string   `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	SnapshotTimestamp    []byte   `protobuf:"bytes,7,opt,name=snapshotTimestamp,proto3" json:"snapshotTimestamp,omitempty"`
	Timestamp            uint64   `protobuf:"varint,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Signature            []byte   `protobuf:"bytes,9,opt,name=signature,proto3" json:"signature,omitempty"`
	Nounce               []byte   `protobuf:"bytes,10,opt,name=nounce,proto3" json:"nounce,omitempty"`
	Difficulty           []byte   `protobuf:"bytes,11,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	FAmount              []byte   `protobuf:"bytes,12,opt,name=fAmount,proto3" json:"fAmount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountBlockDb) Reset()         { *m = AccountBlockDb{} }
func (m *AccountBlockDb) String() string { return proto.CompactTextString(m) }
func (*AccountBlockDb) ProtoMessage()    {}
func (*AccountBlockDb) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_block_db_aaf2d4e3b2c35eec, []int{0}
}
func (m *AccountBlockDb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountBlockDb.Unmarshal(m, b)
}
func (m *AccountBlockDb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountBlockDb.Marshal(b, m, deterministic)
}
func (dst *AccountBlockDb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountBlockDb.Merge(dst, src)
}
func (m *AccountBlockDb) XXX_Size() int {
	return xxx_messageInfo_AccountBlockDb.Size(m)
}
func (m *AccountBlockDb) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountBlockDb.DiscardUnknown(m)
}

var xxx_messageInfo_AccountBlockDb proto.InternalMessageInfo

func (m *AccountBlockDb) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *AccountBlockDb) GetPrevHash() []byte {
	if m != nil {
		return m.PrevHash
	}
	return nil
}

func (m *AccountBlockDb) GetFromHash() []byte {
	if m != nil {
		return m.FromHash
	}
	return nil
}

func (m *AccountBlockDb) GetTokenId() []byte {
	if m != nil {
		return m.TokenId
	}
	return nil
}

func (m *AccountBlockDb) GetBalance() []byte {
	if m != nil {
		return m.Balance
	}
	return nil
}

func (m *AccountBlockDb) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *AccountBlockDb) GetSnapshotTimestamp() []byte {
	if m != nil {
		return m.SnapshotTimestamp
	}
	return nil
}

func (m *AccountBlockDb) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *AccountBlockDb) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *AccountBlockDb) GetNounce() []byte {
	if m != nil {
		return m.Nounce
	}
	return nil
}

func (m *AccountBlockDb) GetDifficulty() []byte {
	if m != nil {
		return m.Difficulty
	}
	return nil
}

func (m *AccountBlockDb) GetFAmount() []byte {
	if m != nil {
		return m.FAmount
	}
	return nil
}

type AccountBlockMeta struct {
	AccountId            []byte   `protobuf:"bytes,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
	Height               []byte   `protobuf:"bytes,2,opt,name=height,proto3" json:"height,omitempty"`
	Status               uint32   `protobuf:"varint,3,opt,name=Status,proto3" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountBlockMeta) Reset()         { *m = AccountBlockMeta{} }
func (m *AccountBlockMeta) String() string { return proto.CompactTextString(m) }
func (*AccountBlockMeta) ProtoMessage()    {}
func (*AccountBlockMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_block_db_aaf2d4e3b2c35eec, []int{1}
}
func (m *AccountBlockMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountBlockMeta.Unmarshal(m, b)
}
func (m *AccountBlockMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountBlockMeta.Marshal(b, m, deterministic)
}
func (dst *AccountBlockMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountBlockMeta.Merge(dst, src)
}
func (m *AccountBlockMeta) XXX_Size() int {
	return xxx_messageInfo_AccountBlockMeta.Size(m)
}
func (m *AccountBlockMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountBlockMeta.DiscardUnknown(m)
}

var xxx_messageInfo_AccountBlockMeta proto.InternalMessageInfo

func (m *AccountBlockMeta) GetAccountId() []byte {
	if m != nil {
		return m.AccountId
	}
	return nil
}

func (m *AccountBlockMeta) GetHeight() []byte {
	if m != nil {
		return m.Height
	}
	return nil
}

func (m *AccountBlockMeta) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*AccountBlockDb)(nil), "vitepb.AccountBlockDb")
	proto.RegisterType((*AccountBlockMeta)(nil), "vitepb.AccountBlockMeta")
}

func init() {
	proto.RegisterFile("account_block_db.proto", fileDescriptor_account_block_db_aaf2d4e3b2c35eec)
}

var fileDescriptor_account_block_db_aaf2d4e3b2c35eec = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0x95, 0xdc, 0xdc, 0xb4, 0x39, 0x94, 0x0a, 0x3c, 0x54, 0x16, 0x42, 0x28, 0xea, 0x94,
	0x01, 0xb1, 0xf0, 0x04, 0x45, 0x0c, 0x74, 0x60, 0x09, 0xec, 0xc5, 0x4e, 0x9c, 0xc6, 0x6a, 0x62,
	0x47, 0xf1, 0x49, 0x25, 0xde, 0x95, 0x87, 0x41, 0x76, 0xdc, 0xa6, 0x12, 0x5b, 0xbe, 0xff, 0x3b,
	0xb1, 0x8e, 0x7f, 0xc3, 0x8a, 0x15, 0x85, 0x1e, 0x14, 0xee, 0x78, 0xa3, 0x8b, 0xc3, 0xae, 0xe4,
	0x4f, 0x5d, 0xaf, 0x51, 0x93, 0xf8, 0x28, 0x51, 0x74, 0x7c, 0xfd, 0x13, 0xc2, 0x72, 0x33, 0x8e,
	0xbc, 0xd8, 0x89, 0x57, 0x4e, 0x96, 0x10, 0xa2, 0xa6, 0x41, 0x1a, 0x64, 0x8b, 0x3c, 0x44, 0x4d,
	0xee, 0x60, 0xde, 0xf5, 0xe2, 0xf8, 0xc6, 0x4c, 0x4d, 0x43, 0x97, 0x9e, 0xd9, 0xba, 0xaa, 0xd7,
	0xad, 0x73, 0xff, 0x46, 0x77, 0x62, 0x42, 0x61, 0x86, 0xfa, 0x20, 0xd4, 0xb6, 0xa4, 0x91, 0x53,
	0x27, 0xb4, 0x86, 0xb3, 0x86, 0xa9, 0x42, 0xd0, 0xff, 0xa3, 0xf1, 0x48, 0x08, 0x44, 0x25, 0x43,
	0x46, 0xe3, 0x34, 0xc8, 0x92, 0xdc, 0x7d, 0x93, 0x47, 0xb8, 0x35, 0x8a, 0x75, 0xa6, 0xd6, 0xf8,
	0x29, 0x5b, 0x61, 0x90, 0xb5, 0x1d, 0x9d, 0xb9, 0xff, 0xfe, 0x0a, 0x72, 0x0f, 0x09, 0x9e, 0xa7,
	0xe6, 0x69, 0x90, 0x45, 0xf9, 0x14, 0x58, 0x6b, 0xe4, 0x5e, 0x31, 0x1c, 0x7a, 0x41, 0x13, 0x77,
	0xc6, 0x14, 0x90, 0x15, 0xc4, 0x4a, 0x0f, 0x76, 0x2d, 0x70, 0xca, 0x13, 0x79, 0x00, 0x28, 0x65,
	0x55, 0xc9, 0x62, 0x68, 0xf0, 0x9b, 0x5e, 0x39, 0x77, 0x91, 0xd8, 0xfb, 0x54, 0x9b, 0xd6, 0x76,
	0x48, 0x17, 0xe3, 0x7d, 0x3c, 0xae, 0xbf, 0xe0, 0xe6, 0xb2, 0xdd, 0x77, 0x81, 0xcc, 0xee, 0xe0,
	0x1f, 0x65, 0x5b, 0xfa, 0x9a, 0xa7, 0xc0, 0xee, 0x50, 0x0b, 0xb9, 0xaf, 0xd1, 0x77, 0xed, 0xc9,
	0xe6, 0x1f, 0xc8, 0x70, 0x30, 0xae, 0xe7, 0xeb, 0xdc, 0x13, 0x8f, 0xdd, 0x7b, 0x3e, 0xff, 0x06,
	0x00, 0x00, 0xff, 0xff, 0xf5, 0x14, 0xee, 0x54, 0xe9, 0x01, 0x00, 0x00,
}