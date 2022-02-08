package mtproto

const (
	MTPROTO_VERSION = 2
)

type TLObject interface {
	Encode() []byte
	Decode(dbuf *DecodeBuf) error
	String() string
}
