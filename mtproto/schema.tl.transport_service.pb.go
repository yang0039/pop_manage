// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema.tl.transport_service.proto

package mtproto

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

///////////////////////////////////////////////////////////////////////////////
// rpc_drop_answer#58e4a740 req_msg_id:long = RpcDropAnswer;
type TLRpcDropAnswer struct {
	ReqMsgId int64 `protobuf:"varint,1,opt,name=req_msg_id,json=reqMsgId,proto3" json:"req_msg_id,omitempty"`
}

func (m *TLRpcDropAnswer) Reset()         { *m = TLRpcDropAnswer{} }
func (m *TLRpcDropAnswer) String() string { return proto.CompactTextString(m) }
func (*TLRpcDropAnswer) ProtoMessage()    {}
func (*TLRpcDropAnswer) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dbb7f3fce79bd46, []int{0}
}
func (m *TLRpcDropAnswer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLRpcDropAnswer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLRpcDropAnswer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TLRpcDropAnswer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLRpcDropAnswer.Merge(m, src)
}
func (m *TLRpcDropAnswer) XXX_Size() int {
	return m.Size()
}
func (m *TLRpcDropAnswer) XXX_DiscardUnknown() {
	xxx_messageInfo_TLRpcDropAnswer.DiscardUnknown(m)
}

var xxx_messageInfo_TLRpcDropAnswer proto.InternalMessageInfo

func (m *TLRpcDropAnswer) GetReqMsgId() int64 {
	if m != nil {
		return m.ReqMsgId
	}
	return 0
}

///////////////////////////////////////////////////////////////////////////////
// get_future_salts#b921bd04 num:int = FutureSalts;
type TLGetFutureSalts struct {
	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (m *TLGetFutureSalts) Reset()         { *m = TLGetFutureSalts{} }
func (m *TLGetFutureSalts) String() string { return proto.CompactTextString(m) }
func (*TLGetFutureSalts) ProtoMessage()    {}
func (*TLGetFutureSalts) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dbb7f3fce79bd46, []int{1}
}
func (m *TLGetFutureSalts) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLGetFutureSalts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLGetFutureSalts.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TLGetFutureSalts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLGetFutureSalts.Merge(m, src)
}
func (m *TLGetFutureSalts) XXX_Size() int {
	return m.Size()
}
func (m *TLGetFutureSalts) XXX_DiscardUnknown() {
	xxx_messageInfo_TLGetFutureSalts.DiscardUnknown(m)
}

var xxx_messageInfo_TLGetFutureSalts proto.InternalMessageInfo

func (m *TLGetFutureSalts) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

///////////////////////////////////////////////////////////////////////////////
// ping#7abe77ec ping_id:long = Pong;
type TLPing struct {
	PingId int64 `protobuf:"varint,1,opt,name=ping_id,json=pingId,proto3" json:"ping_id,omitempty"`
}

func (m *TLPing) Reset()         { *m = TLPing{} }
func (m *TLPing) String() string { return proto.CompactTextString(m) }
func (*TLPing) ProtoMessage()    {}
func (*TLPing) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dbb7f3fce79bd46, []int{2}
}
func (m *TLPing) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLPing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLPing.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TLPing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLPing.Merge(m, src)
}
func (m *TLPing) XXX_Size() int {
	return m.Size()
}
func (m *TLPing) XXX_DiscardUnknown() {
	xxx_messageInfo_TLPing.DiscardUnknown(m)
}

var xxx_messageInfo_TLPing proto.InternalMessageInfo

func (m *TLPing) GetPingId() int64 {
	if m != nil {
		return m.PingId
	}
	return 0
}

///////////////////////////////////////////////////////////////////////////////
// ping_delay_disconnect#f3427b8c ping_id:long disconnect_delay:int = Pong;
type TLPingDelayDisconnect struct {
	PingId          int64 `protobuf:"varint,1,opt,name=ping_id,json=pingId,proto3" json:"ping_id,omitempty"`
	DisconnectDelay int32 `protobuf:"varint,2,opt,name=disconnect_delay,json=disconnectDelay,proto3" json:"disconnect_delay,omitempty"`
}

func (m *TLPingDelayDisconnect) Reset()         { *m = TLPingDelayDisconnect{} }
func (m *TLPingDelayDisconnect) String() string { return proto.CompactTextString(m) }
func (*TLPingDelayDisconnect) ProtoMessage()    {}
func (*TLPingDelayDisconnect) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dbb7f3fce79bd46, []int{3}
}
func (m *TLPingDelayDisconnect) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLPingDelayDisconnect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLPingDelayDisconnect.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TLPingDelayDisconnect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLPingDelayDisconnect.Merge(m, src)
}
func (m *TLPingDelayDisconnect) XXX_Size() int {
	return m.Size()
}
func (m *TLPingDelayDisconnect) XXX_DiscardUnknown() {
	xxx_messageInfo_TLPingDelayDisconnect.DiscardUnknown(m)
}

var xxx_messageInfo_TLPingDelayDisconnect proto.InternalMessageInfo

func (m *TLPingDelayDisconnect) GetPingId() int64 {
	if m != nil {
		return m.PingId
	}
	return 0
}

func (m *TLPingDelayDisconnect) GetDisconnectDelay() int32 {
	if m != nil {
		return m.DisconnectDelay
	}
	return 0
}

///////////////////////////////////////////////////////////////////////////////
// destroy_session#e7512126 session_id:long = DestroySessionRes;
type TLDestroySession struct {
	SessionId int64 `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (m *TLDestroySession) Reset()         { *m = TLDestroySession{} }
func (m *TLDestroySession) String() string { return proto.CompactTextString(m) }
func (*TLDestroySession) ProtoMessage()    {}
func (*TLDestroySession) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dbb7f3fce79bd46, []int{4}
}
func (m *TLDestroySession) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLDestroySession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLDestroySession.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TLDestroySession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLDestroySession.Merge(m, src)
}
func (m *TLDestroySession) XXX_Size() int {
	return m.Size()
}
func (m *TLDestroySession) XXX_DiscardUnknown() {
	xxx_messageInfo_TLDestroySession.DiscardUnknown(m)
}

var xxx_messageInfo_TLDestroySession proto.InternalMessageInfo

func (m *TLDestroySession) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

///////////////////////////////////////////////////////////////////////////////
// contest.saveDeveloperInfo#9a5f6e95 vk_id:int name:string phone_number:string age:int city:string = Bool;
type TLContestSaveDeveloperInfo struct {
	VkId        int32  `protobuf:"varint,1,opt,name=vk_id,json=vkId,proto3" json:"vk_id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PhoneNumber string `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Age         int32  `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	City        string `protobuf:"bytes,5,opt,name=city,proto3" json:"city,omitempty"`
}

func (m *TLContestSaveDeveloperInfo) Reset()         { *m = TLContestSaveDeveloperInfo{} }
func (m *TLContestSaveDeveloperInfo) String() string { return proto.CompactTextString(m) }
func (*TLContestSaveDeveloperInfo) ProtoMessage()    {}
func (*TLContestSaveDeveloperInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dbb7f3fce79bd46, []int{5}
}
func (m *TLContestSaveDeveloperInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLContestSaveDeveloperInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLContestSaveDeveloperInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TLContestSaveDeveloperInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLContestSaveDeveloperInfo.Merge(m, src)
}
func (m *TLContestSaveDeveloperInfo) XXX_Size() int {
	return m.Size()
}
func (m *TLContestSaveDeveloperInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TLContestSaveDeveloperInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TLContestSaveDeveloperInfo proto.InternalMessageInfo

func (m *TLContestSaveDeveloperInfo) GetVkId() int32 {
	if m != nil {
		return m.VkId
	}
	return 0
}

func (m *TLContestSaveDeveloperInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TLContestSaveDeveloperInfo) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *TLContestSaveDeveloperInfo) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *TLContestSaveDeveloperInfo) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func init() {
	proto.RegisterType((*TLRpcDropAnswer)(nil), "mtproto.TL_rpc_drop_answer")
	proto.RegisterType((*TLGetFutureSalts)(nil), "mtproto.TL_get_future_salts")
	proto.RegisterType((*TLPing)(nil), "mtproto.TL_ping")
	proto.RegisterType((*TLPingDelayDisconnect)(nil), "mtproto.TL_ping_delay_disconnect")
	proto.RegisterType((*TLDestroySession)(nil), "mtproto.TL_destroy_session")
	proto.RegisterType((*TLContestSaveDeveloperInfo)(nil), "mtproto.TL_contest_saveDeveloperInfo")
}

func init() { proto.RegisterFile("schema.tl.transport_service.proto", fileDescriptor_2dbb7f3fce79bd46) }

var fileDescriptor_2dbb7f3fce79bd46 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xb1, 0xae, 0xd3, 0x30,
	0x14, 0x86, 0x9b, 0xdb, 0xe6, 0x96, 0x7b, 0x40, 0xe2, 0xca, 0x1d, 0x88, 0x50, 0x89, 0x68, 0x16,
	0x60, 0xe9, 0x40, 0xdf, 0x00, 0x75, 0x20, 0x52, 0x60, 0x88, 0x32, 0x73, 0xe4, 0xc6, 0xa7, 0x69,
	0xd4, 0xc4, 0x4e, 0x6d, 0x27, 0xa8, 0x6f, 0xd1, 0xc7, 0x62, 0xec, 0xc8, 0x88, 0xda, 0x17, 0x41,
	0x71, 0x8b, 0x32, 0x31, 0xe5, 0xcf, 0xe7, 0xff, 0xff, 0x6d, 0x1f, 0xc3, 0xc2, 0xe4, 0x3b, 0xaa,
	0xf9, 0xd2, 0x56, 0x4b, 0xab, 0xb9, 0x34, 0x8d, 0xd2, 0x16, 0x0d, 0xe9, 0xae, 0xcc, 0x69, 0xd9,
	0x68, 0x65, 0x15, 0x9b, 0xd6, 0xd6, 0x89, 0xe8, 0x33, 0xb0, 0x2c, 0x41, 0xdd, 0xe4, 0x28, 0xb4,
	0x6a, 0x90, 0x4b, 0xf3, 0x93, 0x34, 0x9b, 0x03, 0x68, 0x3a, 0x60, 0x6d, 0x0a, 0x2c, 0x45, 0xe0,
	0xbd, 0xf7, 0x3e, 0x8e, 0xd3, 0x17, 0x9a, 0x0e, 0xdf, 0x4c, 0x11, 0x8b, 0xe8, 0x03, 0xcc, 0xb2,
	0x04, 0x0b, 0xb2, 0xb8, 0x6d, 0x6d, 0xab, 0x09, 0x0d, 0xaf, 0xac, 0x61, 0xcf, 0x30, 0x96, 0x6d,
	0xed, 0xdc, 0x7e, 0xda, 0xcb, 0x28, 0x82, 0x69, 0x96, 0x60, 0x53, 0xca, 0x82, 0xbd, 0x81, 0x69,
	0xff, 0x1d, 0xea, 0x1e, 0xfb, 0xdf, 0x58, 0x44, 0x3f, 0x20, 0xb8, 0x7b, 0x50, 0x50, 0xc5, 0x8f,
	0x28, 0x4a, 0x93, 0x2b, 0x29, 0x29, 0xb7, 0xff, 0x0d, 0xb1, 0x4f, 0xf0, 0x3c, 0xd8, 0x6e, 0xb9,
	0xe0, 0xc1, 0xed, 0xfb, 0x7a, 0xe0, 0xeb, 0x1e, 0x47, 0x2b, 0x77, 0x41, 0x41, 0xc6, 0x6a, 0x75,
	0x44, 0x43, 0xc6, 0x94, 0x4a, 0xb2, 0x77, 0x00, 0x77, 0x39, 0x94, 0x3f, 0xdd, 0x49, 0x2c, 0xa2,
	0x93, 0x07, 0xf3, 0x2c, 0xc1, 0x5c, 0x49, 0x4b, 0xc6, 0xa2, 0xe1, 0x1d, 0xad, 0xa9, 0xa3, 0x4a,
	0x35, 0xa4, 0x63, 0xb9, 0x55, 0x6c, 0x06, 0x7e, 0xb7, 0xff, 0x17, 0xf5, 0xd3, 0x49, 0xb7, 0x8f,
	0x05, 0x63, 0x30, 0x91, 0xbc, 0x26, 0x77, 0x92, 0xa7, 0xd4, 0x69, 0xb6, 0x80, 0x57, 0xcd, 0x4e,
	0x49, 0x42, 0xd9, 0xd6, 0x1b, 0xd2, 0xc1, 0xd8, 0xad, 0xbd, 0x74, 0xec, 0xbb, 0x43, 0xfd, 0xdc,
	0x78, 0x41, 0xc1, 0xe4, 0x36, 0x37, 0x5e, 0x50, 0x5f, 0x94, 0x97, 0xf6, 0x18, 0xf8, 0xb7, 0xa2,
	0x5e, 0x7f, 0x79, 0xfb, 0xeb, 0x12, 0x7a, 0xe7, 0x4b, 0xe8, 0xfd, 0xb9, 0x84, 0xde, 0xe9, 0x1a,
	0x8e, 0xce, 0xd7, 0x70, 0xf4, 0xfb, 0x1a, 0x8e, 0xbe, 0x3e, 0x6c, 0x1e, 0xdd, 0x5b, 0xae, 0xfe,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x76, 0xe7, 0xa5, 0x63, 0xf9, 0x01, 0x00, 0x00,
}

func (m *TLRpcDropAnswer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLRpcDropAnswer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TLRpcDropAnswer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ReqMsgId != 0 {
		i = encodeVarintSchemaTlTransportService(dAtA, i, uint64(m.ReqMsgId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TLGetFutureSalts) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLGetFutureSalts) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TLGetFutureSalts) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Num != 0 {
		i = encodeVarintSchemaTlTransportService(dAtA, i, uint64(m.Num))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TLPing) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLPing) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TLPing) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PingId != 0 {
		i = encodeVarintSchemaTlTransportService(dAtA, i, uint64(m.PingId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TLPingDelayDisconnect) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLPingDelayDisconnect) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TLPingDelayDisconnect) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DisconnectDelay != 0 {
		i = encodeVarintSchemaTlTransportService(dAtA, i, uint64(m.DisconnectDelay))
		i--
		dAtA[i] = 0x10
	}
	if m.PingId != 0 {
		i = encodeVarintSchemaTlTransportService(dAtA, i, uint64(m.PingId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TLDestroySession) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLDestroySession) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TLDestroySession) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SessionId != 0 {
		i = encodeVarintSchemaTlTransportService(dAtA, i, uint64(m.SessionId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TLContestSaveDeveloperInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLContestSaveDeveloperInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TLContestSaveDeveloperInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.City) > 0 {
		i -= len(m.City)
		copy(dAtA[i:], m.City)
		i = encodeVarintSchemaTlTransportService(dAtA, i, uint64(len(m.City)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Age != 0 {
		i = encodeVarintSchemaTlTransportService(dAtA, i, uint64(m.Age))
		i--
		dAtA[i] = 0x20
	}
	if len(m.PhoneNumber) > 0 {
		i -= len(m.PhoneNumber)
		copy(dAtA[i:], m.PhoneNumber)
		i = encodeVarintSchemaTlTransportService(dAtA, i, uint64(len(m.PhoneNumber)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintSchemaTlTransportService(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.VkId != 0 {
		i = encodeVarintSchemaTlTransportService(dAtA, i, uint64(m.VkId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSchemaTlTransportService(dAtA []byte, offset int, v uint64) int {
	offset -= sovSchemaTlTransportService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TLRpcDropAnswer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ReqMsgId != 0 {
		n += 1 + sovSchemaTlTransportService(uint64(m.ReqMsgId))
	}
	return n
}

func (m *TLGetFutureSalts) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Num != 0 {
		n += 1 + sovSchemaTlTransportService(uint64(m.Num))
	}
	return n
}

func (m *TLPing) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PingId != 0 {
		n += 1 + sovSchemaTlTransportService(uint64(m.PingId))
	}
	return n
}

func (m *TLPingDelayDisconnect) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PingId != 0 {
		n += 1 + sovSchemaTlTransportService(uint64(m.PingId))
	}
	if m.DisconnectDelay != 0 {
		n += 1 + sovSchemaTlTransportService(uint64(m.DisconnectDelay))
	}
	return n
}

func (m *TLDestroySession) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SessionId != 0 {
		n += 1 + sovSchemaTlTransportService(uint64(m.SessionId))
	}
	return n
}

func (m *TLContestSaveDeveloperInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VkId != 0 {
		n += 1 + sovSchemaTlTransportService(uint64(m.VkId))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSchemaTlTransportService(uint64(l))
	}
	l = len(m.PhoneNumber)
	if l > 0 {
		n += 1 + l + sovSchemaTlTransportService(uint64(l))
	}
	if m.Age != 0 {
		n += 1 + sovSchemaTlTransportService(uint64(m.Age))
	}
	l = len(m.City)
	if l > 0 {
		n += 1 + l + sovSchemaTlTransportService(uint64(l))
	}
	return n
}

func sovSchemaTlTransportService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSchemaTlTransportService(x uint64) (n int) {
	return sovSchemaTlTransportService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TLRpcDropAnswer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchemaTlTransportService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_rpc_drop_answer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_rpc_drop_answer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReqMsgId", wireType)
			}
			m.ReqMsgId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchemaTlTransportService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReqMsgId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSchemaTlTransportService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchemaTlTransportService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSchemaTlTransportService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLGetFutureSalts) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchemaTlTransportService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_get_future_salts: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_get_future_salts: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Num", wireType)
			}
			m.Num = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchemaTlTransportService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Num |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSchemaTlTransportService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchemaTlTransportService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSchemaTlTransportService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLPing) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchemaTlTransportService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_ping: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_ping: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PingId", wireType)
			}
			m.PingId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchemaTlTransportService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PingId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSchemaTlTransportService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchemaTlTransportService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSchemaTlTransportService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLPingDelayDisconnect) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchemaTlTransportService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_ping_delay_disconnect: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_ping_delay_disconnect: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PingId", wireType)
			}
			m.PingId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchemaTlTransportService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PingId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisconnectDelay", wireType)
			}
			m.DisconnectDelay = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchemaTlTransportService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DisconnectDelay |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSchemaTlTransportService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchemaTlTransportService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSchemaTlTransportService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLDestroySession) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchemaTlTransportService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_destroy_session: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_destroy_session: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionId", wireType)
			}
			m.SessionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchemaTlTransportService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SessionId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSchemaTlTransportService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchemaTlTransportService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSchemaTlTransportService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLContestSaveDeveloperInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchemaTlTransportService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TL_contest_saveDeveloperInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TL_contest_saveDeveloperInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VkId", wireType)
			}
			m.VkId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchemaTlTransportService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VkId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchemaTlTransportService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchemaTlTransportService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchemaTlTransportService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PhoneNumber", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchemaTlTransportService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchemaTlTransportService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchemaTlTransportService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PhoneNumber = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Age", wireType)
			}
			m.Age = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchemaTlTransportService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Age |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field City", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchemaTlTransportService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchemaTlTransportService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchemaTlTransportService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.City = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchemaTlTransportService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchemaTlTransportService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSchemaTlTransportService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSchemaTlTransportService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSchemaTlTransportService
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSchemaTlTransportService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSchemaTlTransportService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSchemaTlTransportService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSchemaTlTransportService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSchemaTlTransportService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSchemaTlTransportService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSchemaTlTransportService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSchemaTlTransportService = fmt.Errorf("proto: unexpected end of group")
)
