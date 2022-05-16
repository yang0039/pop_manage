package dataobject

type RequestRecordDO struct {
	Id        int32  `db:"id"`
	AccountId int32  `db:"account_id"`
	Url       string `db:"url"`
	Method    string `db:"method"`
	ClientIp  string `db:"client_ip"`
	ReqData   string `db:"req_data"`
	IsSuccess bool   `db:"is_success"`
	Reason    string `db:"reason"`
	AddTime   int64  `db:"add_time"`
}
