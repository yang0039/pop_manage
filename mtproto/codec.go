package mtproto

/*
import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net"
	"pop-api/baselib/logger"
	//net2 "winkim/baselib/net2"
)

const (
	CODEC_UNKNOWN = iota
	CODEC_CONNECTED
	CODEC_req_pq
	CODEC_resPQ
	CODEC_req_DH_params
	CODEC_server_DH_params_ok
	CODEC_server_DH_params_fail
	CODEC_set_client_DH_params
	CODEC_dh_gen_ok
	CODEC_dh_gen_retry
	CODEC_dh_gen_fail
	CODEC_AUTH_KEY_OK
	CODEC_ERROR
)

func NewMTProto() *MTProto {
	return &MTProto{}
}

type MTProto struct {
}

type AuthKeyStorager interface {
	GetAuthKey(int64) []byte
	// PutAuthKey(int64, []byte) error
}

func (m *MTProto) NewCodec(rw io.ReadWriter) error {
	logger.Logger.Error("MtProto NewCodec")
	codec := &MTProtoCodec{}
	codec.rw, _ = rw.(io.ReadWriteCloser)
	codec.State = CODEC_CONNECTED
	codec.UserId = 0
	//return codec, nil
	return nil
}

type MTProtoCodec struct {
	rw io.ReadWriteCloser

	// 缓存AuthKey
	AuthKeyStorager

	State int

	//
	AuthKeyId int64
	AuthKey   []byte
	UserId    int32

	Salt      int64
	SessionId int64
	SeqNo     int32
}

func (m *MTProtoCodec) RemoteAddr() net.Addr {
	return m.rw.(net.Conn).RemoteAddr()
}

func (m *MTProtoCodec) LocalAddr() net.Addr {
	return m.rw.(net.Conn).LocalAddr()
}

func (m *MTProtoCodec) getSeqNo(increment bool) int32 {
	value := m.SeqNo
	if increment {
		m.SeqNo += 1
	}
	v := int32(0)
	if increment {
		v = int32(1)
	}

	return value*2 + v
}

func (m *MTProtoCodec) Receive() (interface{}, error) {
	var size int
	var n int
	var err error

	b := make([]byte, 1)
	n, err = io.ReadFull(m.rw, b)
	if err != nil {
		return nil, err
	}

	needAck := bool(b[0]>>7 == 1)

	b[0] = b[0] & 0x7f

	if b[0] < 0x7f {
		size = int(b[0]) << 2
	} else {
		logger.LogSugar.Infof("first_byte2: %v", hex.EncodeToString(b[:1]))
		b := make([]byte, 3)
		n, err = io.ReadFull(m.rw, b)
		if err != nil {
			return nil, err
		}
		size = (int(b[0]) | int(b[1])<<8 | int(b[2])<<16) << 2
	}

	left := size
	buf := make([]byte, size)
	for left > 0 {
		n, err = io.ReadFull(m.rw, buf[size-left:])
		if err != nil {
			return nil, err
		}
		left -= n
	}

	// 截断QuickAck消息，客户端有问题
	if size == 4 {
		logger.LogSugar.Errorf("Server response error: %d", int32(binary.LittleEndian.Uint32(buf)))
		return nil, fmt.Errorf("Recv QuickAckMessage, ignore!!!! sessionId: ", m.SessionId, ", by client ", m.RemoteAddr())
	}

	authKeyId := int64(binary.LittleEndian.Uint64(buf))
	if authKeyId == 0 {
		var message = &UnencryptedMessage{}
		message.NeedAck = needAck
		err = message.decode(buf[8:])
		if err != nil {
			return nil, err
		}
		return message, nil
	} else {
		logger.Logger.Info("Recv authKeyId not 0")

		// TODO(@work): 检查m.State状态，authKeyId不为0时codec状态必须是CODEC_AUTH_KEY_OK或CODEC_resPQ
		if m.State != CODEC_CONNECTED && m.State != CODEC_AUTH_KEY_OK && m.State != CODEC_resPQ && m.State != CODEC_dh_gen_ok {
			// 连接需要断开
			return nil, fmt.Errorf("Invalid state, is CODEC_AUTH_KEY_OK or CODEC_resPQ, but is %d", m.State)
		}
		if m.AuthKeyId == 0 {
			key := m.GetAuthKey(authKeyId)
			if key == nil {
				return nil, fmt.Errorf("Can't find authKey by authKeyId %d", authKeyId)
			}
			m.AuthKeyId = authKeyId
			m.AuthKey = key
			logger.LogSugar.Infof("Found key, keyId: %d key:%v", authKeyId, hex.EncodeToString(key))
		} else if m.AuthKeyId != authKeyId {
			return nil, fmt.Errorf("Key error, cacheKey is ", m.AuthKeyId, ", recved keyId is ", authKeyId)
		}

		var message = &EncryptedMessage2{}
		quickAckId, err := message.decode(m.AuthKey, buf[8:])
		logger.LogSugar.Warnf("(m *MTProtoCodec) Receive quickAckId :%d", quickAckId)
		if err != nil {
			logger.LogSugar.Errorf("Decode encrypted message error: %v", err)
			return nil, err
		}

		if m.State != CODEC_AUTH_KEY_OK {
			// m.SessionId = message.SessionId
			m.State = CODEC_AUTH_KEY_OK
		}

		return message, nil
	}
}

func (m *MTProtoCodec) Send(msg interface{}) error {
	switch msg.(type) {
	case *UnencryptedMessage:
		b, _ := msg.(*UnencryptedMessage).encode()

		sb := make([]byte, 4)
		// minus padding
		size := len(b) / 4

		if size < 127 {
			sb = []byte{byte(size)}
		} else {
			binary.LittleEndian.PutUint32(sb, uint32(size<<8|127))
		}

		b = append(sb, b...)

		logger.LogSugar.Warnf("MTProtoCodec - len b: %d, x:%x", len(b), b)

		_, err := m.rw.Write(b)
		if err != nil {
			logger.LogSugar.Warnf("MTProtoCodec - Send UnencryptedMessage errror: %v", err)
			return err
		}
		return nil

	case *EncryptedMessage2:
		encrypedMessage, _ := msg.(*EncryptedMessage2)
		encrypedMessage.SessionId = m.SessionId
		encrypedMessage.Salt = m.Salt
		encrypedMessage.SeqNo = m.getSeqNo(true)
		// switch encrypedMessage.Object.(type) {
		// case *TLUpdates:
		// 	glog.Info("send message: %v", encrypedMessage)
		// }
		b, _ := encrypedMessage.encode(int64(m.AuthKeyId), m.AuthKey)

		//switch encrypedMessage.Object.(type) {
		//case *TLRpcResult:
		//	result := encrypedMessage.Object.(*TLRpcResult)
		//	switch result.Result.(type) {
		//	case
		//	}
		//
		//	msgDetailedInfo := NewTLMsgDetailedInfo()
		//
		//	msgDetailedInfo.SetBytes()
		//	msgDetailedInfo.SetAnswerMsgId()
		//	msgDetailedInfo.SetMsgId()
		//	msgDetailedInfo.SetStatus(0)
		//default:
		//}

		sb := make([]byte, 4)
		// minus padding
		size := len(b) / 4

		if size < 127 {
			sb = []byte{byte(size)}
		} else {
			binary.LittleEndian.PutUint32(sb, uint32(size<<8|127))
		}

		b = append(sb, b...)

		logger.LogSugar.Warnf("MTProtoCodec - len b: %d, x:%x", len(b), b)

		_, err := m.rw.Write(b)
		if err != nil {
			logger.LogSugar.Warnf("MTProtoCodec - Send EncryptedMessage2 errror: %v", err)
			return err
		}
		return nil

	case *MsgDetailedInfoContainer:
		// TODO(@work): 蹩脚的实现
		encrypedMessage := msg.(*MsgDetailedInfoContainer).Message
		encrypedMessage.SessionId = m.SessionId
		encrypedMessage.Salt = m.Salt
		encrypedMessage.SeqNo = m.getSeqNo(true)
		b, _ := encrypedMessage.encode(int64(m.AuthKeyId), m.AuthKey)

		msgDetailedInfo := NewTLMsgDetailedInfo()
		objData := encrypedMessage.Object.Encode()
		// TODO(@work): ReqMsgId置入MsgDetailedInfoContainer内
		msgDetailedInfo.SetMsgId(encrypedMessage.Object.(*TLRpcResult).ReqMsgId)
		msgDetailedInfo.SetAnswerMsgId(encrypedMessage.MessageId)
		msgDetailedInfo.SetBytes(int32(len(objData)))
		msgDetailedInfo.SetStatus(0)

		msgDetailedInfoMessage := &EncryptedMessage2{
			NeedAck: false,
			// SeqNo:	  seqNo,
			Object: msgDetailedInfo,
		}
		m.Send(msgDetailedInfoMessage)

		sb := make([]byte, 4)
		// minus padding
		size := len(b) / 4

		if size < 127 {
			sb = []byte{byte(size)}
		} else {
			binary.LittleEndian.PutUint32(sb, uint32(size<<8|127))
		}

		b = append(sb, b...)

		logger.LogSugar.Warnf("MTProtoCodec - len b: %d, x:%x", len(b), b)

		_, err := m.rw.Write(b)
		if err != nil {
			logger.LogSugar.Warnf("MTProtoCodec - Send EncryptedMessage2 errror: %v", err)
			return err
		}
		return nil
	}

	return errors.New("msg type error, only UnencryptedMessage or EncryptedMessage2, but recv invalid type")
}

func (m *MTProtoCodec) Close() error {
	return m.rw.Close()
}
*/