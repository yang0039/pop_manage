package mtproto

/*
import (
	//  "reflect"
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"winkim/baselib/crypto"
	"pop-api/baselib/logger"
	"runtime/debug"
)

const (
	QUICK_ACKID = iota
	UNENCRYPTED_MESSAGE
	ENCRYPTED_MESSAGE
)

type MTProtoMessage interface {
	// encode([]byte) ([]byte, error)
	// decode([]byte) error
	// MessageType() int
}

type QuickAckMessage struct {
	ackId int32
}

func (m *QuickAckMessage) MessageType() int {
	return QUICK_ACKID
}

func (m *QuickAckMessage) encode() ([]byte, error) {
	return nil, nil
}

func (m *QuickAckMessage) decode(b []byte) error {
	if len(b) != 4 {
		return fmt.Errorf("Message len: %d (need 4)", len(b))
	}
	m.ackId = int32(binary.LittleEndian.Uint32(b))
	return nil
}

type UnencryptedMessage struct {
	NeedAck bool

	// authKeyId int64
	MessageId int64
	// messageDataLength int32
	// messageData []byte

	// classID int32
	Object TLObject
}

func (m *UnencryptedMessage) MessageType() int {
	return UNENCRYPTED_MESSAGE
}

func (m *UnencryptedMessage) Encode() []byte {
	buf, _ := m.encode()
	return buf
}

func (m *UnencryptedMessage) encode() ([]byte, error) {
	x := NewEncodeBuf(512)
	x.Long(0)
	m.MessageId = GenerateMessageId()
	x.Long(m.MessageId)

	if m.Object == nil {
		x.Int(0)
	} else {
		b := m.Object.Encode()
		x.Int(int32(len(b)))
		x.Bytes(b)
	}
	return x.buf, nil
}

func (m *UnencryptedMessage) Decode(b []byte) error {
	return m.decode(b)
}

func (m *UnencryptedMessage) decode(b []byte) error {
	dbuf := NewDecodeBuf(b)
	// m.authKeyId = dbuf.Long()
	m.MessageId = dbuf.Long()

	// mod := m.messageId & 3
	// if mod != 1 && mod != 3 {
	// 	return fmt.Errorf("Wrong bits of message_id: %d", mod)
	// }

	messageLen := dbuf.Int()
	if messageLen < 4 {
		return fmt.Errorf("message len(%d) < 4", messageLen)
	}

	if int(messageLen) != dbuf.size-12 {
		return fmt.Errorf("message len: %d (need %d)", messageLen, dbuf.size-12)
	}

	m.Object = dbuf.Object()
	//  fmt.Println("m.Object", m.Object, reflect.TypeOf(m.Object))
	if m.Object == nil {
		return fmt.Errorf("decode object is nil")
	}

	return dbuf.err
}

////////////////////////////////////////////////////////////////////////////////////////////
// MsgDetailedInfo
type MsgDetailedInfoContainer struct {
	Message *EncryptedMessage2
}

////////////////////////////////////////////////////////////////////////////////////////////
// TODO(@work): 将Encrypt和Descrypt移到底层
type EncryptedMessage2 struct {
	authKeyId int64
	NeedAck   bool

	msgKey    []byte
	Salt      int64
	SessionId int64
	MessageId int64
	SeqNo     int32
	Object    TLObject
	ObjBuf    []byte

	QuickAckBuf []byte
}

func NewEncryptedMessage2(authKeyId int64) *EncryptedMessage2 {
	return &EncryptedMessage2{
		authKeyId: authKeyId,
	}
}

func (m *EncryptedMessage2) String() string {
	return fmt.Sprintf("{auth_key_id: %d. salt: %d, session_id: %d, message_id: %d, seq_no: %d}",
		m.authKeyId, m.Salt, m.SessionId, m.MessageId, m.SeqNo)
}

func (m *EncryptedMessage2) MessageType() int {
	return ENCRYPTED_MESSAGE
}

func (m *EncryptedMessage2) Encode(authKeyId int64, authKey []byte) ([]byte, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover error Encode authKeyId:%d, len(authKey):%d, panic: %v\n%s", authKeyId, len(authKey), err, string(debug.Stack()))
		}
	}()

	buf, err := m.encode(authKeyId, authKey)
	return buf, err
}

func (m *EncryptedMessage2) encode(authKeyId int64, authKey []byte) ([]byte, error) {
	if m.ObjBuf == nil {
		m.ObjBuf = m.Object.Encode()
	}
	var additionalSize = (32 + len(m.ObjBuf)) % 16
	if additionalSize != 0 {
		additionalSize = 16 - additionalSize
	}
	if MTPROTO_VERSION == 2 && additionalSize < 12 {
		additionalSize += 16
	}

	x := NewEncodeBuf(32 + len(m.ObjBuf) + additionalSize)
	// x.Long(authKeyId)
	// msgKey := make([]byte, 16)
	// x.Bytes(msgKey)
	x.Long(m.Salt)
	x.Long(m.SessionId)
	if m.MessageId == 0 {
		m.MessageId = GenerateMessageId()
	}
	x.Long(m.MessageId)
	x.Int(m.SeqNo)
	x.Int(int32(len(m.ObjBuf)))
	x.Bytes(m.ObjBuf)
	x.Bytes(crypto.GenerateNonce(additionalSize))

	encryptedData, _ := m.encrypt(authKey, x.buf, len(m.ObjBuf))
	x2 := NewEncodeBuf(56 + len(m.ObjBuf) + additionalSize)
	x2.Long(authKeyId)
	x2.Bytes(m.msgKey)
	x2.Bytes(encryptedData)

	return x2.buf, nil
}

func (m *EncryptedMessage2) Decode(authKeyId int64, authKey, b []byte) (uint32, error) {
	_ = authKeyId
	return m.decode(authKey, b)
}

func (m *EncryptedMessage2) decode(authKey []byte, b []byte) (quickAckId uint32, e error) {
	defer func() {
		if err := recover(); err != nil {
			logger.LogSugar.Errorf("recover error EncryptedMessage2 decode: %v", err)
			e = errors.New("recover error EncryptedMessage2 decode")
		}
	}()

	msgKey := b[:16]
	// aesKey, aesIV := generateMessageKey(msgKey, authKey, false)
	// x, err := doAES256IGEdecrypt(b[16:], aesKey, aesIV)

	var x []byte
	x, quickAckId, e = m.decrypt(msgKey, authKey, b[16:])
	if e != nil {
		return
	}

	dbuf := NewDecodeBuf(x)

	m.Salt = dbuf.Long()      // salt
	m.SessionId = dbuf.Long() // session_id
	m.MessageId = dbuf.Long()

	// mod := m.messageId & 3
	// if mod != 1 && mod != 3 {
	//	return fmt.Errorf("Wrong bits of message_id: %d", mod)
	// }

	m.SeqNo = dbuf.Int()
	messageLen := dbuf.Int()
	if int(messageLen) > dbuf.size-32 {
		e = fmt.Errorf("Message len: %d (need less than %d)", messageLen, dbuf.size-32)
		return
	}

	// logger.LogSugar.Infof("salt: %d, sessionId: %d, messageId: %d, seqNo: %d, messageLen: %d", m.salt, m.SessionId, m.MessageId, m.SeqNo, messageLen)
	m.Object = dbuf.Object()
	if m.Object == nil {
		e = fmt.Errorf("Decode object is nil")
	}

	return
}

func (m *EncryptedMessage2) decrypt(msgKey, authKey, data []byte) ([]byte, uint32, error) {
	// dbuf := NewDecodeBuf(data)
	// m.authKeyId = dbuf.Long()
	// msgKey := dbuf.Bytes(16)

	var dataLen = uint32(len(data))
	// 创建aesKey, aesIV
	aesKey, aesIV := generateMessageKey(msgKey, authKey, false)
	d := crypto.NewAES256IGECryptor(aesKey, aesIV)

	x, err := d.Decrypt(data)
	if err != nil {
		logger.LogSugar.Errorf("descrypted data error: %v", err)
		return nil, 0, err
	}

	//// 校验解密后的数据合法性
	messageLen := binary.LittleEndian.Uint32(x[28:])
	if messageLen+32 > dataLen {
		// 	return fmt.Errorf("Message len: %d (need less than %d)", messageLen, dbuf.size-32)
		err = fmt.Errorf("descrypted data error: Wrong message length %d", messageLen)
		logger.Logger.Error(err.Error())
		return nil, 0, err
	}

	messageKey := make([]byte, 96)
	switch MTPROTO_VERSION {
	case 2:
		t_d := make([]byte, 0, 32+dataLen)
		t_d = append(t_d, authKey[88:88+32]...)
		t_d = append(t_d, x[:dataLen]...)
		copy(messageKey, crypto.Sha256Digest(t_d))
	default:
		copy(messageKey[4:], crypto.Sha1Digest(x[:32+messageLen]))
	}

	QuickAckId := (1 << 31) | binary.LittleEndian.Uint32(messageKey[:4])

	if !bytes.Equal(messageKey[8:8+16], msgKey[:16]) {
		err = fmt.Errorf("descrypted data error: (data: %s, aesKey: %s, aseIV: %s, authKeyId: %d, authKey: %s), msgKey verify error, sign: %s, msgKey: %s",
			hex.EncodeToString(data[:64]),
			hex.EncodeToString(aesKey),
			hex.EncodeToString(aesIV),
			m.authKeyId,
			hex.EncodeToString(authKey[88:88+32]),
			hex.EncodeToString(messageKey[8:8+16]),
			hex.EncodeToString(msgKey[:16]))
		logger.Logger.Error(err.Error())
		return nil, 0, err
	}

	return x, QuickAckId, nil
}

// for a gateway error: error line: t_d = append(t_d, authKey[88+8:88+8+32]...)
//panic: runtime error: slice bounds out of range
// goroutine 26 [running]:
// winkim/mtproto.(*EncryptedMessage2).encrypt(0xc00095dbb0, 0x2ac9d48, 0x0, 0x0, 0xc001e2d2f0, 0x30, 0x30, 0x4, 0x0, 0x0, ...)
//         /Users/lollipop/go/src/winkim/mtproto/message.go:303 +0x474
// winkim/mtproto.(*EncryptedMessage2).encode(0xc00095dbb0, 0x61ee8c97d37544fd, 0x2ac9d48, 0x0, 0x0, 0xc00095dc38, 0xb58a6d, 0x15f3500, 0x0, 0x200)
//         /Users/lollipop/go/src/winkim/mtproto/message.go:185 +0x3f0
// winkim/mtproto.(*EncryptedMessage2).Encode(...)
//         /Users/lollipop/go/src/winkim/mtproto/message.go:154
// main.encodeMessage(0xc000846460, 0x61ee8c97d37544fd, 0x2ac9d48, 0x0, 0x0, 0x52aac20efa97e12a, 0xc00a5ace00, 0x4, 0x200, 0x0, ...)
//         /Users/lollipop/go/src/winkim/gateway/server.go:374 +0x162
// main.(*AuthStruct).WakeUpPushSession(0xc000e587e0)
//         /Users/lollipop/go/src/winkim/gateway/server.go:147 +0x289
// main.(*Server).pushRpcRunLoop(0xc000125230)
//         /Users/lollipop/go/src/winkim/gateway/server.go:752 +0x359
// created by main.(*Server).RunLoop
// 		/Users/lollipop/go/src/winkim/gateway/server.go:639 +0x101
func (m *EncryptedMessage2) encrypt(authKey []byte, data []byte, messageSize int) ([]byte, error) {
	messageKey := make([]byte, 32)
	switch MTPROTO_VERSION {
	case 2:
		t_d := make([]byte, 0, 32+len(data))
		t_d = append(t_d, authKey[88+8:88+8+32]...)
		t_d = append(t_d, data...)
		copy(messageKey, crypto.Sha256Digest(t_d))
	default:
		copy(messageKey[4:], crypto.Sha1Digest(data[:32+messageSize]))
	}

	// copy(message_key[8:], )
	// memcpy(p_data + 8, message_key + 8, 16);

	aesKey, aesIV := generateMessageKey(messageKey[8:8+16], authKey, true)
	e := crypto.NewAES256IGECryptor(aesKey, aesIV)

	x, err := e.Encrypt(data)
	if err != nil {
		logger.LogSugar.Errorf("Encrypt data error: %v", err)
		return nil, err
	}

	m.msgKey = messageKey[8 : 8+16]
	return x, nil
}
*/
////////////////////////////////////////////////////////////////////////////////////////////
//
