package dataobject

type Account struct {
	Id          int32  `db:"id" json:"id"`
	AccountName string `db:"account_name" json:"account_name"`
	UserName    string `db:"user_name" json:"user_name"`
	Pwd         string `db:"pwd" json:"pwd"`
	PwdUtil     int32  `db:"pwd_util" json:"pwd_util"`
	Operator    string `db:"operator" json:"operator"`
	IsEffect    int32  `db:"is_effect" json:"is_effect"`
	AddTime     int32  `db:"add_time" json:"add_time"`
}
