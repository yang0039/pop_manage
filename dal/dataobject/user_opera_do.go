package dataobject

type UserOpera struct {
	Id           int32  `db:"id"`
	AccountId    int32  `db:"account_id"`
	UserId       int32  `db:"user_id"`
	OperaType    int32  `db:"opera_type"`
	OperaContent string `db:"opera_content"`
	AddTime      int64  `db:"add_time"`
}
