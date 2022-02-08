package mtproto


/*
import (
	"io"
	"pop-api/baselib/logger"
	"encoding/binary"
	"fmt"
	// "winkim/baselib/net2"
	"net"
)

// https://core.telegram.org/mtproto#tcp-transport
//
// In case 4-byte data alignment is needed,
// an intermediate version of the original protocol may be used:
// if the client sends 0xeeeeeeee as the first int (four bytes),
// then packet length is encoded always by four bytes as in the original version,
// but the sequence number and CRC32 are omitted,
// thus decreasing total packet size by 8 bytes.
//
type MTProtoIntermediateCodec struct {
	conn net.Conn
}

func NewMTProtoIntermediateCodec(conn net.Conn) *MTProtoIntermediateCodec {
	return &MTProtoIntermediateCodec{
		conn: conn,
	}
}

func (c *MTProtoIntermediateCodec) Receive() (interface{}, error) {
	logger.Logger.Error("MTProtoIntermediateCodec")
	var size int
	var n int
	var err error

	b := make([]byte, 4)
	n, err = io.ReadFull(c.conn, b)
	if err != nil {
		return nil, err
	}

	size = int(binary.LittleEndian.Uint32(b) << 2)

	// logger.LogSugar.Infof("first_byte: %v", hex.EncodeToString(b[:1]))
	// needAck := bool(b[0] >> 7 == 1)
	// _ = needAck

	//b[0] = b[0] & 0x7f
	//// logger.LogSugar.Infof("first_byte2: %v", hex.EncodeToString(b[:1]))
	//
	//if b[0] < 0x7f {
	//	size = int(b[0]) << 2
	//	logger.LogSugar.Infof("size1: %d", size)
	//	if size == 0 {
	//		return nil, nil
	//	}
	//} else {
	//	logger.LogSugar.Infof("first_byte2: %v", hex.EncodeToString(b[:1]))
	//	b2 := make([]byte, 3)
	//	n, err = io.ReadFull(c.conn, b2)
	//	if err != nil {
	//		return nil, err
	//	}
	//	size = (int(b2[0]) | int(b2[1])<<8 | int(b2[2])<<16) << 2
	//	logger.LogSugar.Infof("size2: %d", size)
	//}

	left := size
	buf := make([]byte, size)
	for left > 0 {
		n, err = io.ReadFull(c.conn, buf[size-left:])
		if err != nil {
			logger.LogSugar.Errorf("ReadFull2 error: %v", err)
			return nil, err
		}
		left -= n
	}
	//if size > 10240 {
	//	logger.LogSugar.Infof("ReadFull2: %v", hex.EncodeToString(buf[:256]))
	//}

	// TODO(@work): process report ack and quickack
	// 截断QuickAck消息，客户端有问题
	if size == 4 {
		logger.LogSugar.Errorf("Server response error: %d", int32(binary.LittleEndian.Uint32(buf)))
		// return nil, fmt.Errorf("Recv QuickAckMessage, ignore!!!!") //  connId: ", c.stream, ", by client ", m.RemoteAddr())
		return nil, nil
	}

	authKeyId := int64(binary.LittleEndian.Uint64(buf))
	message := NewMTPRawMessage(TRANSPORT_TCP, authKeyId, false, buf)
	return message, nil
}

func (c *MTProtoIntermediateCodec) Send(msg interface{}) error {
	logger.Logger.Error("MTProtoIntermediateCodec Send")

	message, ok := msg.(*MTPRawMessage)
	if !ok {
		err := fmt.Errorf("msg type error, only MTPRawMessage, msg: {%v}", msg)
		logger.Logger.Error(err.Error())
		return err
	}

	// b := message.Encode()
	b := message.Payload

	sb := make([]byte, 4)
	// minus padding
	size := len(b)/4

	//if size < 127 {
	//	sb = []byte{byte(size)}
	//} else {
	binary.LittleEndian.PutUint32(sb, uint32(size))
	//}

	b = append(sb, b...)
	_, err := c.conn.Write(b)

	if err != nil {
		logger.LogSugar.Errorf("Send msg error: %v", err)
	}

	return err
}

func (c *MTProtoIntermediateCodec) Close() error {
	return c.conn.Close()
}
*/