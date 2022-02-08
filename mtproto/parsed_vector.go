package mtproto

////////////////////////////////////////////////////////////////////////////////
// message Vector_SecureValue {
//     repeated SecureValue datas = 1;
// }
func NewVector_SecureValue() *Vector_SecureValue {
	return &Vector_SecureValue{}
}
func (m *Vector_SecureValue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *Vector_SecureValue) Decode(dbuf *DecodeBuf) error {
	dbuf.Int() // TODO(@work): Check crc32 invalid
	l1 := dbuf.Int()
	m.Datas = make([]*SecureValue, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &SecureValue{}
		(*m.Datas[i]).Decode(dbuf)
	}

	return dbuf.err
}

////////////////////////////////////////////////////////////////////////////////
// message Vector_SavedContact {
//     repeated SavedContact datas = 1;
// }
func NewVector_SavedContact() *Vector_SavedContact {
	return &Vector_SavedContact{}
}
func (m *Vector_SavedContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *Vector_SavedContact) Decode(dbuf *DecodeBuf) error {
	dbuf.Int() // TODO(@work): Check crc32 invalid
	l1 := dbuf.Int()
	m.Datas = make([]*SavedContact, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &SavedContact{}
		(*m.Datas[i]).Decode(dbuf)
	}

	return dbuf.err
}

////////////////////////////////////////////////////////////////////////////////
// message Vector_MessageRange {
//     repeated MessageRange datas = 1;
// }
func NewVector_MessageRange() *Vector_MessageRange {
	return &Vector_MessageRange{}
}
func (m *Vector_MessageRange) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *Vector_MessageRange) Decode(dbuf *DecodeBuf) error {
	dbuf.Int() // TODO(@work): Check crc32 invalid
	l1 := dbuf.Int()
	m.Datas = make([]*MessageRange, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &MessageRange{}
		(*m.Datas[i]).Decode(dbuf)
	}

	return dbuf.err
}

////////////////////////////////////////////////////////////////////////////////
// message Vector_DialogPeer {
//     repeated DialogPeer datas = 1;
// }
func NewVector_DialogPeer() *Vector_DialogPeer {
	return &Vector_DialogPeer{}
}
func (m *Vector_DialogPeer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *Vector_DialogPeer) Decode(dbuf *DecodeBuf) error {
	dbuf.Int() // TODO(@work): Check crc32 invalid
	l1 := dbuf.Int()
	m.Datas = make([]*DialogPeer, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &DialogPeer{}
		(*m.Datas[i]).Decode(dbuf)
	}

	return dbuf.err
}

////////////////////////////////////////////////////////////////////////////////
// message Vector_EmojiLanguage {
//     repeated EmojiLanguage datas = 1;
// }
func NewVector_EmojiLanguage() *Vector_EmojiLanguage {
	return &Vector_EmojiLanguage{}
}
func (m *Vector_EmojiLanguage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *Vector_EmojiLanguage) Decode(dbuf *DecodeBuf) error {
	dbuf.Int() // TODO(@work): Check crc32 invalid
	l1 := dbuf.Int()
	m.Datas = make([]*EmojiLanguage, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &EmojiLanguage{}
		(*m.Datas[i]).Decode(dbuf)
	}

	return dbuf.err
}

////////////////////////////////////////////////////////////////////////////////
// message Vector_Messages_SearchCounter {
//     repeated TL_messages_SearchCounter datas = 1;
// }
func NewVector_Messages_SearchCounter() *Vector_Messages_SearchCounter {
	return &Vector_Messages_SearchCounter{}
}
func (m *Vector_Messages_SearchCounter) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *Vector_Messages_SearchCounter) Decode(dbuf *DecodeBuf) error {
	dbuf.Int() // TODO(@work): Check crc32 invalid
	l1 := dbuf.Int()
	m.Datas = make([]*Messages_SearchCounter, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &Messages_SearchCounter{}
		(*m.Datas[i]).Decode(dbuf)
	}

	return dbuf.err
}

////////////////////////////////////////////////////////////////////////////////
// message Vector_DialogFilter {
//     repeated DialogFilter datas = 1;
// }
func NewVector_DialogFilter() *Vector_DialogFilter {
	return &Vector_DialogFilter{}
}
func (m *Vector_DialogFilter) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *Vector_DialogFilter) Decode(dbuf *DecodeBuf) error {
	dbuf.Int() // TODO(@work): Check crc32 invalid
	l1 := dbuf.Int()
	m.Datas = make([]*DialogFilter, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &DialogFilter{}
		(*m.Datas[i]).Decode(dbuf)
	}

	return dbuf.err
}

////////////////////////////////////////////////////////////////////////////////
// message Vector_DialogFilterSuggested {
//     repeated DialogFilterSuggested datas = 1;
// }
func NewVector_DialogFilterSuggested() *Vector_DialogFilterSuggested {
	return &Vector_DialogFilterSuggested{}
}
func (m *Vector_DialogFilterSuggested) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *Vector_DialogFilterSuggested) Decode(dbuf *DecodeBuf) error {
	dbuf.Int() // TODO(@work): Check crc32 invalid
	l1 := dbuf.Int()
	m.Datas = make([]*DialogFilterSuggested, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &DialogFilterSuggested{}
		(*m.Datas[i]).Decode(dbuf)
	}

	return dbuf.err
}

////////////////////////////////////////////////////////////////////////////////
// message Vector_FileHash {
//     repeated FileHash datas = 1;
// }
func NewVector_FileHash() *Vector_FileHash {
	return &Vector_FileHash{}
}
func (m *Vector_FileHash) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *Vector_FileHash) Decode(dbuf *DecodeBuf) error {
	dbuf.Int() // TODO(@work): Check crc32 invalid
	l1 := dbuf.Int()
	m.Datas = make([]*FileHash, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &FileHash{}
		(*m.Datas[i]).Decode(dbuf)
	}

	return dbuf.err
}
