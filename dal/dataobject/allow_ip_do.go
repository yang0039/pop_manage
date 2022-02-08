package dataobject

type AllowIp struct {
	Id       int32  `db:"id" json:"id"`
	Operator string `db:"operator" json:"operator"`
	Ip       string `db:"ip" json:"ip"`
	AddTime  int32  `db:"add_time" json:"add_time"`
}
