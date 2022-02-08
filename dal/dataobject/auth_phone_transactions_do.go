package dataobject

type AuthPhoneTransactions struct {
	Id              int32  `db:"id" json:"id"`
	AuthKeyId       int64  `db:"auth_key_id" json:"auth_key_id"`
	PhoneNumber     string `db:"phone_number" json:"phone_number"`
	Code            string `db:"code" json:"code"`
	CodeExpired     int64  `db:"code_expired" json:"code_expired"`
	//TransactionHash string `db:"transaction_hash"`
	//SentCodeType    int32  `db:"sent_code_type"`
	State           int32  `db:"state" json:"state"`
	//ApiId           int32  `db:"api_id"`
	//ApiHash         string `db:"api_hash"`
	Attempts        int32  `db:"attempts" json:"attempts"`
	CreatedTime     int64  `db:"created_time" json:"created_time"`
	IsDeleted       int32  `db:"is_deleted" json:"is_deleted"`
}




