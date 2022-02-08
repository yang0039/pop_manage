package mtproto

import (
	"fmt"
)

type MessageBase interface {
	Encode() ([]byte)
	Decode(b []byte) error
}

//
//type CodecAble interface {
//	Encode() ([]byte, error)
//	Decode(dbuf *DecodeBuf) error
//}

func NewMTPRawMessage(protocolType int, authKeyId int64, needAck bool, payload []byte) *MTPRawMessage {
	return &MTPRawMessage{
		protocolType:   protocolType,
		authKeyId: authKeyId,
		needAck: needAck,
		Payload: payload,
	}
}

////////////////////////////////////////////////////////////////////////////
// 代理使用
type MTPRawMessage struct {
	protocolType   int
	authKeyId int64
	needAck bool
	// 原始数据	前8个字节含有AuthKeyId
	Payload    []byte
	ClientIp	string	//通过https请求时，真实的IP只能通过这里传递了
}

func (m *MTPRawMessage) ProtocolType() int {
	return m.protocolType
}

func (m *MTPRawMessage) AuthKeyId() int64 {
	return m.authKeyId
}

func (m *MTPRawMessage) NeedAck() bool {
	return m.needAck
}


func (m *MTPRawMessage) Encode() []byte {
	x := NewEncodeBuf(2+4+len(m.Payload))
	if m.needAck {
		x.Int16(1)
	} else {
		x.Int16(0)
	}
	x.Int(int32(len(m.Payload)))
	x.Bytes(m.Payload)
	
	return x.GetBuf()
}

func (m *MTPRawMessage) Decode(b []byte) error {
	dbuf := NewDecodeBuf(b)
	if dbuf.Uint16() == 1 {
		m.needAck = true
	}
	messageLen := dbuf.Int()
	if int(messageLen) != dbuf.size-6 {
		return fmt.Errorf("invalid UnencryptedRawMessage len: %d (need %d)", messageLen, dbuf.size-12)
	}
	m.Payload = dbuf.Bytes(int(messageLen))
	return dbuf.err
}

////////////////////////////////////////////////////////////////////////////
func NewUnencryptedRawMessage() *UnencryptedRawMessage {
	return &UnencryptedRawMessage{
		AuthKeyId: 0,
	}
}

type UnencryptedRawMessage struct {
	// TODO(@work): reportAck and quickAck
	// NeedAck bool
	AuthKeyId int64
	MessageId int64
	MessageData []byte
}

func (m *UnencryptedRawMessage) Encode() []byte {
	// 一次性分配
	x := NewEncodeBuf(20+len(m.MessageData))
	x.Long(0)
	m.MessageId = GenerateMessageId()
	x.Long(m.MessageId)
	x.Int(int32(len(m.MessageData)))
	x.Bytes(m.MessageData)
	return x.buf
}

func (m *UnencryptedRawMessage) Decode(b []byte) error {
	dbuf := NewDecodeBuf(b)
	m.MessageId = dbuf.Long()
	messageLen := dbuf.Int()
	if int(messageLen) != dbuf.size-12 {
		return fmt.Errorf("invalid UnencryptedRawMessage len: %d (need %d)", messageLen, dbuf.size-12)
	}
	m.MessageData = dbuf.Bytes(int(messageLen))
	return dbuf.err
}

type EncryptedRawMessage struct {
	// TODO(@work): reportAck and quickAck
	// NeedAck bool
	AuthKeyId     int64
	MsgKey        []byte
	EncryptedData []byte
}

func NewEncryptedRawMessage(authKeyId int64) *EncryptedRawMessage {
	return &EncryptedRawMessage{
		AuthKeyId: authKeyId,
	}
}

func (m *EncryptedRawMessage) Encode() []byte {
	// 一次性分配
	x := NewEncodeBuf(24+len(m.EncryptedData))
	x.Long(m.AuthKeyId)
	x.Bytes(m.MsgKey)
	x.Bytes(m.EncryptedData)
	return x.buf
}

func (m *EncryptedRawMessage) Decode(b []byte) error {
	dbuf := NewDecodeBuf(b)
	m.MsgKey = dbuf.Bytes(16)
	m.EncryptedData = dbuf.Bytes(len(b)-16)
	return dbuf.err
}
