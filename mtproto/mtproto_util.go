package mtproto

func ToBool(b bool) *Bool {
	if b {
		return NewTLBoolTrue().To_Bool()
	} else {
		return NewTLBoolFalse().To_Bool()
	}
}

func ToError(code int32, message string) *Error {
	tlerr := NewTLError()
	tlerr.SetCode(code)
	tlerr.SetText(message)
	return tlerr.To_Error()
}

func ToIntResult(val int32) *IntResult {
	tl := NewTLIntResult()
	tl.SetValue(val)
	return tl.To_IntResult()
}

func FromBool(b *Bool) bool {
	return TLConstructor_CRC32_boolTrue == b.GetConstructor()
}

/*
 * Updates扩展方法
 */

func (self *TLUpdates) AddUpdate(update *Update) {
	self.Data2.Updates = append(self.Data2.Updates, update)
}

func (self *TLUpdates) PrependUpdate(update *Update) {
	var out []*Update
	out = append(out, update)
	out = append(out, self.Data2.Updates...)
	self.Data2.Updates = out
}

func (self *TLUpdates) AddUser(user *User) bool {
	if user == nil {
		return false
	}
	for _, u := range self.Data2.Users {
		if u.Data2.Id == user.Data2.Id {
			return false
		}
	}
	self.Data2.Users = append(self.Data2.Users, user)
	return true
}

func (self *TLUpdates) AddChat(chat *Chat) bool {
	if chat == nil {
		return false
	}
	for _, c := range self.Data2.Chats {
		if c.Data2.Id == chat.Data2.Id {
			return false
		}
	}
	self.Data2.Chats = append(self.Data2.Chats, chat)
	return true
}

func (self *Updates) AddUpdate(update *Update) {
	self.Data2.Updates = append(self.Data2.Updates, update)
}

func (self *Updates) PrependUpdate(update *Update) {
	var out []*Update
	out = append(out, update)
	out = append(out, self.Data2.Updates...)
	self.Data2.Updates = out
}

func (self *Updates) AddUser(user *User) bool {
	if user == nil {
		return false
	}
	for _, u := range self.Data2.Users {
		if u.Data2.Id == user.Data2.Id {
			return false
		}
	}
	self.Data2.Users = append(self.Data2.Users, user)
	return true
}

func (self *Updates) AddChat(chat *Chat) bool {
	if chat == nil {
		return false
	}
	for _, c := range self.Data2.Chats {
		if c.Data2.Id == chat.Data2.Id {
			return false
		}
	}
	self.Data2.Chats = append(self.Data2.Chats, chat)
	return true
}
