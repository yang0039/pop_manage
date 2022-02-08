package dataobject

type LoginLog struct {
	Id        int32  `db:"id" json:"id"`
	AccountId int32  `db:"account_id" json:"account_id"`
	LoginTime int32  `db:"login_time" json:"login_time"`
	LoginIp   string `db:"login_ip" json:"login_ip"`
}
