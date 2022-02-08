package mtproto

import (

)

// <<effective-go>>的接口检查章节中对这种用法做了解释：
// 当需要确保包中的类型确实满足某接口时，就会使用这种方式。
// 如果像json.RawMessage这样的类型需要一个自定义的JSON表示，
// 它应该实现json.Marshaler，但这里不存在可导致编译器自动进行验证的静态转换。
// 如果类型非故意地不能满足此接口，JSON编码器将仍能工作，
// 但却不使用自定义的实现。要确保实现是正确的，在此包中可以使用一个具有空白标识符的全局声明：
//	var _ json.Marshaler = (*RawMessage)(nil)
//


//var _ net.Conn = &MTProtoConn{}

/*
type Config struct {
	EnableCrypt        bool
	HandshakeTimeout   time.Duration
	RewriterBufferSize int
	ReconnWaitTimeout  time.Duration
}
*/

/*
type Dialer func() (net.Conn, error)

type MTProtoConn struct {
	base      net.Conn

	// reomoteAddr	net.Addr
	// aes_key
	encryptor *crypto.AesCTR128Encrypt
	decryptor *crypto.AesCTR128Encrypt
	listener *Listener
	id       uint64

	closed    bool
	closeChan chan struct{}
	closeOnce sync.Once
}

func newMTProtoConn(base net.Conn, id uint64, encryptor *crypto.AesCTR128Encrypt, decryptor *crypto.AesCTR128Encrypt) (conn *MTProtoConn) {
	return &MTProtoConn{
		base:           base,
		id:             id,
		encryptor:		encryptor,
		decryptor:		decryptor,
		closeChan:      make(chan struct{}),
	}
}

func (c *MTProtoConn) WrapBaseForTest(wrap func(net.Conn) net.Conn) {
	c.base = wrap(c.base)
}

func (c *MTProtoConn) RemoteAddr() net.Addr {
	return c.base.RemoteAddr()
}

func (c *MTProtoConn) LocalAddr() net.Addr {
	return c.base.LocalAddr()
}

func (c *MTProtoConn) SetDeadline(t time.Time) error {
	return c.base.SetDeadline(t)
}

func (c *MTProtoConn) SetReadDeadline(t time.Time) error {
	return c.base.SetReadDeadline(t)
}

func (c *MTProtoConn) SetWriteDeadline(t time.Time) error {
	return c.base.SetWriteDeadline(t)
}

func (c *MTProtoConn) Close() error {
	logger.Logger.Info("Close()")
	c.closeOnce.Do(func() {
		c.closed = true
		if c.listener != nil {
			c.listener.delConn(c.id)
		}
		close(c.closeChan)
	})
	return c.base.Close()
}

func (c *MTProtoConn) Read(b []byte) (n int, err error) {
	n, err = c.base.Read(b)
	if err == nil {
		c.decryptor.Encrypt(b[:])
		return
	}

	logger.LogSugar.Warnf("MTProtoConn - Will close conn by %v, reason:%v", c.base.RemoteAddr(), err)
	c.base.Close()
	return
}

func (c *MTProtoConn) Write(b []byte) (n int, err error) {
	c.encryptor.Encrypt(b[:])
	return c.base.Write(b)
}

*/