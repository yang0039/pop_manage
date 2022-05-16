package dto

type NoteObj struct {
	LabelId string `json:"label_id" form:"label_id"`
	PeerId  int32  `json:"peer_id" form:"peer_id"`
	Note    string `json:"note" form:"note"`
}

type QryType struct {
	Type     int32  `json:"type" form:"type"`
	ChatType int32  `json:"chat_type" form:"chat_type"`
	Qry      string `json:"qry" form:"qry"`
	Limit    int32  `json:"limit" form:"limit"`
	Offset   int32  `json:"offset" form:"offset"`
}

type QryUser struct {
	UserId int32 `json:"user_id" form:"user_id"`
	PeerId int32 `json:"peer_id" form:"peer_id"`
	Limit  int32 `json:"limit" form:"limit"`
	Offset int32 `json:"offset" form:"offset"`
}

type UserChat struct {
	Type     int32 `json:"type" form:"type"`
	ChatType int32 `json:"chat_type" form:"chat_type"`
	UserId   int32 `json:"user_id" form:"user_id"`
	Limit    int32 `json:"limit" form:"limit"`
	Offset   int32 `json:"offset" form:"offset"`
}

type Account struct {
	Name string `json:"name" form:"name"`
	Pwd  string `json:"pwd" form:"pwd"`
}

type PhoneTransaction struct {
	Type   int32  `json:"type" form:"type"`
	Qry    string `json:"qry" form:"qry"`
	Limit  int32  `json:"limit" form:"limit"`
	Offset int32  `json:"offset" form:"offset"`
}

type SetUserOfficial struct {
	UserId    int32 `json:"user_id"`
	OperaType int32 `json:"opera_type"` // 1:设为客服， 2：解除客服
}

type Banned struct {
	UserId    int32  `json:"user_id"`
	OperaType int32  `json:"opera_type"`
	Reason    string `json:"reason"`
}

type BannedReq struct {
	UserId    int32  `json:"user_id"`
	OperaType int32  `json:"opera_type"`
	Opera     string `json:"opera"`
	Reason    string `json:"reason"`
}

type ApiReq struct {
	From   string      `json:"from"`
	Method int32       `json:"method"`
	Data   interface{} `json:"data"`
}

type Role struct {
	Id       int32  `json:"id"`
	Name     string `json:"name"`
	Operator string `json:"operator"`
	//IsEffect      int32   `json:"is_effect"`
	//AddTime       int64   `json:"add_time"`
	PermissionIds []int32 `json:"permission_ids"`
}

type Acount struct {
	Id            int32   `json:"id" form:"id"`
	AccountName   string  `json:"account_name"`
	UserName      string  `json:"user_name"`
	Pwd           string  `json:"pwd"`
	NewPwd        string  `json:"new_pwd"`
	PwdUtil       int32   `json:"pwd_util"`
	RoleIds       []int32 `json:"role_ids"`
	ForbiddenType int32   `json:"forbidden_type"`
	Limit         int32   `json:"limit" form:"limit"`
	Offset        int32   `json:"offset" form:"offset"`
}

type ChatMsg struct {
	ChatId  int32 `json:"chat_id" form:"chat_id"`
	MaxTime int32 `json:"max_time" form:"max_time"`
	Limit   int32 `json:"limit" form:"limit"`
	Offset  int32 `json:"offset" form:"offset"`
}

type AllowIp struct {
	Id  int32  `json:"id" form:"id"`
	Ip  string `json:"ip" form:"ip"`
	Pwd string `json:"pwd" form:"pwd"`
}

type Config struct {
	Key   string `json:"key" form:"key"`
	Value string `json:"value" form:"value"`
}

type QryFile struct {
	FileId int64 `json:"file_id" form:"file_id"`
}

type Lable struct {
	Id   int32  `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Pwd  string `json:"pwd" form:"pwd"`
}

type MessageRecord struct {
	FromId      int32 `json:"from_id" form:"from_id"`
	PeerId      int32 `json:"peer_id" form:"peer_id"`
	MessageType int32 `json:"message_type" form:"message_type"`
	Start       int64 `json:"start" form:"start"`
	End         int64 `json:"end" form:"end"`
	Limit       int32 `json:"limit" form:"limit"`
	Offset      int32 `json:"offset" form:"offset"`
}

type ChatStatus struct {
	ChatId int32  `json:"chat_id" form:"chat_id"`
	Status int32  `json:"status" form:"status"`
	Days   int32  `json:"days" form:"days"`
	Note   string `json:"note" form:"note"`
	Limit  int32  `json:"limit" form:"limit"`
	Offset int32  `json:"offset" form:"offset"`
}

type StoreType struct {
	UserId int32 `json:"user_id" form:"user_id"`
	Type   int32 `json:"type" form:"type"` // 0:所有 1:图片 2:视频 3:音频 4:文件 5:其他
	Limit  int32 `json:"limit" form:"limit"`
	Offset int32 `json:"offset" form:"offset"`
}

type UpdateUserName struct {
	UserId   int32  `json:"user_id" form:"user_id"`
	UserName string `json:"username" form:"username"`
}

type UpdateUserPhone struct {
	UserId int32  `json:"user_id" form:"user_id"`
	Phone  string `json:"phone" form:"phone"`
}

type QryChat struct {
	ChatId int32 `form:"chat_id"`
}