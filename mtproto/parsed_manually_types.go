package mtproto

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"reflect"

	"pop-api/baselib/logger"
)

//const (
//	TLConstructor_CRC32_message2  		= 0x5bb8e511
//	TLConstructor_CRC32_msg_container  	= 0x73f1f8dc
//	TLConstructor_CRC32_msg_copy  		= 0xe06046b2
//	TLConstructor_CRC32_gzip_packed 	= 0x3072cfa1
//)

///////////////////////////////////////////////////////////////////////////////
//message2 msg_id:long seqno:int bytes:int body:Object = Message2; // parsed manually
type TLMessage2 struct {
	MsgId          int64
	Seqno          int32
	Bytes          int32
	Object         TLObject
	WithoutUpdates bool
}

func (m *TLMessage2) String() string {
	return "{message2#5bb8e511}"
}

func (m *TLMessage2) Encode() []byte {
	x := NewEncodeBuf(512)
	// x.Int(int32(TLConstructor_CRC32_message2))
	x.Long(m.MsgId)
	x.Int(m.Seqno)
	x.Int(m.Bytes)
	x.StringBytes(m.Object.Encode())
	return x.buf
}

func (m *TLMessage2) Encode2() []byte {
	x := NewEncodeBuf(512)
	// x.Int(int32(TLConstructor_CRC32_message2))
	x.Long(m.MsgId)
	x.Int(m.Seqno)

	buf := m.Object.Encode()
	fmt.Println("TLMessage2 buf", len(buf), reflect.TypeOf(m.Object), m.MsgId, m.Seqno)
	// fmt.Println("buf", hex.EncodeToString(buf))

	x.Int(int32(len(buf)))
	x.Bytes(buf)
	// fmt.Println("x.buf", hex.EncodeToString(x.buf))

	return x.buf
}

func (m *TLMessage2) Decode(dbuf *DecodeBuf) error {
	m.MsgId = dbuf.Long()
	m.Seqno = dbuf.Int()
	m.Bytes = dbuf.Int()
	b := dbuf.Bytes(int(m.Bytes))

	dbuf2 := NewDecodeBuf(b)
	m.Object = dbuf2.Object()
	if m.Object == nil {
		err := fmt.Errorf("Decode core_message error: %s", hex.EncodeToString(b))
		logger.Logger.Error(err.Error())
		return err
	}

	return dbuf2.err
}

///////////////////////////////////////////////////////////////////////////////
//msg_container#73f1f8dc messages:vector<message2> = MessageContainer; // parsed manually
type TLMsgContainer struct {
	Messages []TLMessage2
}

func (m *TLMsgContainer) String() string {
	return "{msg_container#73f1f8dc}"
}

func (m *TLMsgContainer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_msg_container))

	x.Int(int32(len(m.Messages)))
	for _, m := range m.Messages {
		x.Bytes(m.Encode2())
	}
	return x.buf
}

func (m *TLMsgContainer) Decode(dbuf *DecodeBuf) error {
	len := dbuf.Int()
	logger.LogSugar.Infof("TLMsgContainer: messages len: %d", len)
	for i := 0; i < int(len); i++ {
		// logger.LogSugar.Infof("TLMsgContainer: messages[%d]: ", i)
		// classID := dbuf.Int()
		// if classID != (int32)(TLConstructor_CRC32_message2) {
		// 	err := fmt.Errorf("Decode TL_message2 error, invalid TL_message2 classID, classID: 0x%x", uint32(classID))
		// 	return err
		// }
		message2 := &TLMessage2{}
		err := message2.Decode(dbuf)
		if err != nil {
			logger.LogSugar.Errorf("Decode message2 error: %v", err)
			return err
		}

		m.Messages = append(m.Messages, *message2)
	}
	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
//msg_copy#e06046b2 orig_message:Message2 = MessageCopy; // parsed manually, not used - use msg_container
type TLMsgCopy struct {
	OrigMessage TLMessage2
}

func (m *TLMsgCopy) String() string {
	return "{msg_copy#e06046b2}"
}

func (m *TLMsgCopy) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_msg_copy))
	x.Bytes(m.OrigMessage.Encode())
	return x.buf
}

func (m *TLMsgCopy) Decode(dbuf *DecodeBuf) error {
	o := dbuf.Object()
	message2, _ := o.(*TLMessage2)
	m.OrigMessage = *message2
	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
//gzip_packed#3072cfa1 packed_data:string = Object; // parsed manually
type TLGzipPacked struct {
	PackedData []byte
}

func (m *TLGzipPacked) String() string {
	return "{gzip_packed#3072cfa1}"
}

func (m *TLGzipPacked) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_gzip_packed))
	x.Bytes(m.PackedData)
	return x.buf
}

func (m *TLGzipPacked) Decode(dbuf *DecodeBuf) error {
	//gzip消息解析错误，待处理 TODO
	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x288 pc=0x6db6a7]

	// goroutine 122 [running]:
	// compress/gzip.(*Reader).Read(0x0, 0xc420659000, 0x1000, 0x1000, 0x1000, 0x1000, 0x0)
	// 	/usr/local/go/src/compress/gzip/gunzip.go:247 +0x37
	// winkim/mtproto.(*TLGzipPacked).Decode(0xc420652520, 0xc42062aa80, 0xc420652520, 0x11b8940)
	// 	/Users/pcname/go/src/winkim/mtproto/parsed_manually_types.go:169 +0x155
	// winkim/mtproto.(*DecodeBuf).Object(0xc42062aa80, 0xc42062aa80, 0x47903f)
	// 	/Users/pcname/go/src/winkim/mtproto/decode.go:321 +0x13a
	// main.NewInvokeAfterMsgExt(0xc4200e1900, 0xc4200dc100)
	// 	/Users/pcname/go/src/winkim/biz_server/mtproto_session_ext.go:32 +0xb5
	// main.(*clientSession).onClientMessage(0xc4202b9490, 0xf7ad8f54b08ca49b, 0x5b7f69862e978c08, 0x8d, 0x11bd640, 0xc4200e1900, 0xc42024fd10)
	// 	/Users/pcname/go/src/winkim/biz_server/client_session.go:583 +0x6a4
	// main.(*clientSession).onClientMessage(0xc4202b9490, 0xf7ad8f54b08ca49b, 0x5b7f69862ed91820, 0x1000000a4, 0x11c3fc0, 0xc4200e15c0, 0xc42024fd10)
	// 	/Users/pcname/go/src/winkim/biz_server/client_session.go:561 +0x390
	// main.(*clientSessionManager).onSessionData(0xc4204dccf0, 0xc4200e15a0)
	// 	/Users/pcname/go/src/winkim/biz_server/client_session_manager.go:382 +0x3cf
	// main.(*clientSessionManager).runLoop(0xc4204dccf0)
	// 	/Users/pcname/go/src/winkim/biz_server/client_session_manager.go:181 +0x247
	// created by main.(*clientSessionManager).Start
	// 	/Users/pcname/go/src/winkim/biz_server/client_session_manager.go:144 +0x8e
	// defer func() {
	// 	if r := recover(); r != nil {
	// 	}
	// }()

	// m.PackedData = make([]byte, 0, 4096)

	// var buf bytes.Buffer
	// buf.Write(dbuf.StringBytes())
	// gz, err := gzip.NewReader(&buf)
	// if err != nil {
	// 	return err
	// }

	// b := make([]byte, 4096)
	// for true {
	// 	n, err := gz.Read(b)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	m.PackedData = append(m.PackedData, b[0:n]...)
	// 	if n <= 0 {
	// 		break
	// 	}
	// }

	in := dbuf.StringBytes()
	if dbuf.err != nil {
		logger.Logger.Error(dbuf.err.Error())
		return dbuf.err
	}
	reader, err := gzip.NewReader(bytes.NewReader(in))
	if err != nil {
		logger.LogSugar.Errorf("TLGzipPacked decode err:%v", err.Error())
		if len(in) > 4 {
			m.PackedData = in[4:]
		} else {
			m.PackedData = in
		}
		logger.LogSugar.Debugf("classId:%d, claddId:%d", binary.LittleEndian.Uint32(in), binary.LittleEndian.Uint32(m.PackedData))
		return nil
		// return err
	}
	defer reader.Close()

	out, err := ioutil.ReadAll(reader)
	if err != nil {
		logger.Logger.Error(err.Error())
		return err
	}

	m.PackedData = out

	return nil
}

///////////////////////////////////////////////////////////////////////////////
//rpc_result#f35c6d01 req_msg_id:long result:Object = RpcResult; // parsed manually
type TLRpcResult struct {
	ReqMsgId int64
	Result   TLObject
}

func (m *TLRpcResult) String() string {
	return "{rpc_result#f35c6d01: req_msg_id:" + string(m.ReqMsgId) + "}"
}

func (m *TLRpcResult) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_rpc_result))
	x.Long(m.ReqMsgId)
	x.Bytes(m.Result.Encode())
	return x.buf
}

func (m *TLRpcResult) Decode(dbuf *DecodeBuf) error {
	m.ReqMsgId = dbuf.Long()
	m.Result = dbuf.Object()
	return dbuf.err
}

// ///////////////////////////////////////////////////////////////////////////////
// // contacts.getContactsLayer70#22c6aa08 hash:string = contacts.Contacts;
// func NewTLContactsGetContactsLayer70() *TLContactsGetContactsLayer70 {
// 	return &TLContactsGetContactsLayer70{}
// }

// func (m *TLContactsGetContactsLayer70) Encode() []byte {
// 	x := NewEncodeBuf(512)
// 	x.Int(int32(TLConstructor_CRC32_contacts_getContactsLayer70))

// 	x.String(m.Hash)

// 	return x.buf
// }

// func (m *TLContactsGetContactsLayer70) Decode(dbuf *DecodeBuf) error {
// 	m.Hash = dbuf.String()

// 	return dbuf.err
// }

////////////////////////////////////////////////////////////////////////////////
// Vector

////////////////////////////////////////////////////////////////////////////////
//// Vector api result type
//message Vector_WallPaper {
//    repeated WallPaper datas = 1;
//}
func NewVector_WallPaper() *Vector_WallPaper {
	return &Vector_WallPaper{}
}

func (m *Vector_WallPaper) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *Vector_WallPaper) Decode(dbuf *DecodeBuf) error {
	dbuf.Int() // TODO(@work): Check crc32 invalid
	l1 := dbuf.Int()
	m.Datas = make([]*WallPaper, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &WallPaper{}
		(*m.Datas[i]).Decode(dbuf)
	}

	return dbuf.err
}

////////////////////////////////////////////////////////////////////////////////
//message Vector_User {
//    repeated User datas = 1;
//}
func NewVector_User() *Vector_User {
	return &Vector_User{}
}
func (m *Vector_User) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *Vector_User) Decode(dbuf *DecodeBuf) error {
	dbuf.Int() // TODO(@work): Check crc32 invalid
	l1 := dbuf.Int()
	m.Datas = make([]*User, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &User{}
		(*m.Datas[i]).Decode(dbuf)
	}

	return dbuf.err
}

////////////////////////////////////////////////////////////////////////////////
//message Vector_ContactStatus {
//    repeated ContactStatus datas = 1;
//}
func NewVector_ContactStatus() *Vector_ContactStatus {
	return &Vector_ContactStatus{}
}
func (m *Vector_ContactStatus) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *Vector_ContactStatus) Decode(dbuf *DecodeBuf) error {
	dbuf.Int() // TODO(@work): Check crc32 invalid
	l1 := dbuf.Int()
	m.Datas = make([]*ContactStatus, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &ContactStatus{}
		(*m.Datas[i]).Decode(dbuf)
	}

	return dbuf.err
}

////////////////////////////////////////////////////////////////////////////////
//message Vector_int {
//    repeated int32 datas = 1;
//}
func NewVectorInt() *VectorInt {
	return &VectorInt{}
}

func (m *VectorInt) Encode() []byte {
	x := NewEncodeBuf(512)
	x.VectorInt(m.Datas)
	return x.buf
}

func (m *VectorInt) Decode(dbuf *DecodeBuf) error {
	// dbuf.Int() // TODO(@work): Check crc32 invalid
	m.Datas = dbuf.VectorInt()
	return dbuf.err
}

//message Vector_ReceivedNotifyMessage {
//    repeated ReceivedNotifyMessage datas = 1;
//}
func NewVector_ReceivedNotifyMessage() *Vector_ReceivedNotifyMessage {
	return &Vector_ReceivedNotifyMessage{}
}
func (m *Vector_ReceivedNotifyMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *Vector_ReceivedNotifyMessage) Decode(dbuf *DecodeBuf) error {
	dbuf.Int() // TODO(@work): Check crc32 invalid
	l1 := dbuf.Int()
	m.Datas = make([]*ReceivedNotifyMessage, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &ReceivedNotifyMessage{}
		(*m.Datas[i]).Decode(dbuf)
	}

	return dbuf.err
}

////////////////////////////////////////////////////////////////////////////////
//message Vector_long {
//    repeated int64 datas = 1;
//}
func NewVectorLong() *VectorLong {
	return &VectorLong{}
}
func (m *VectorLong) Encode() []byte {
	x := NewEncodeBuf(512)
	x.VectorLong(m.Datas)
	return x.buf
}

func (m *VectorLong) Decode(dbuf *DecodeBuf) error {
	m.Datas = dbuf.VectorLong()

	return dbuf.err
}

////////////////////////////////////////////////////////////////////////////////
//message Vector_StickerSetCovered {
//    repeated StickerSetCovered datas = 1;
//}
func NewVector_StickerSetCovered() *Vector_StickerSetCovered {
	return &Vector_StickerSetCovered{}
}
func (m *Vector_StickerSetCovered) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *Vector_StickerSetCovered) Decode(dbuf *DecodeBuf) error {
	dbuf.Int() // TODO(@work): Check crc32 invalid
	l1 := dbuf.Int()
	m.Datas = make([]*StickerSetCovered, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &StickerSetCovered{}
		(*m.Datas[i]).Decode(dbuf)
	}

	return dbuf.err
}

////////////////////////////////////////////////////////////////////////////////
//message Vector_CdnFileHash {
//    repeated CdnFileHash datas = 1;
//}
func NewVector_CdnFileHash() *Vector_CdnFileHash {
	return &Vector_CdnFileHash{}
}
func (m *Vector_CdnFileHash) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *Vector_CdnFileHash) Decode(dbuf *DecodeBuf) error {
	dbuf.Int() // TODO(@work): Check crc32 invalid
	l1 := dbuf.Int()
	m.Datas = make([]*CdnFileHash, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &CdnFileHash{}
		(*m.Datas[i]).Decode(dbuf)
	}

	return dbuf.err
}

////////////////////////////////////////////////////////////////////////////////
//message Vector_LangPackString {
//    repeated LangPackString datas = 1;
//}
func NewVector_LangPackString() *Vector_LangPackString {
	return &Vector_LangPackString{}
}
func (m *Vector_LangPackString) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *Vector_LangPackString) Decode(dbuf *DecodeBuf) error {
	dbuf.Int() // TODO(@work): Check crc32 invalid
	l1 := dbuf.Int()
	m.Datas = make([]*LangPackString, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &LangPackString{}
		(*m.Datas[i]).Decode(dbuf)
	}

	return dbuf.err
}

////////////////////////////////////////////////////////////////////////////////
//message Vector_LangPackLanguage {
//    repeated LangPackLanguage datas = 1;
//}
func NewVector_LangPackLanguage() *Vector_LangPackLanguage {
	return &Vector_LangPackLanguage{}
}
func (m *Vector_LangPackLanguage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *Vector_LangPackLanguage) Decode(dbuf *DecodeBuf) error {
	dbuf.Int() // TODO(@work): Check crc32 invalid
	l1 := dbuf.Int()
	m.Datas = make([]*LangPackLanguage, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &LangPackLanguage{}
		(*m.Datas[i]).Decode(dbuf)
	}

	return dbuf.err
}

func Copy(src TLObject) TLObject {
	encoded := src.Encode()
	dbuf := NewDecodeBuf(encoded)
	if dbuf.GetError() != nil {
		logger.Logger.Error(dbuf.GetError().Error())
	}
	return dbuf.Object()
}
