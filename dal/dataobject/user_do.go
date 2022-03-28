package dataobject

type UserDO struct {
	Id          int32  `db:"id"`
	InviteUid   int32  `db:"invite_uid"`
	AccessHash  int64  `db:"access_hash"`
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	Username    string `db:"username"`
	Phone       string `db:"phone"`
	Support     bool   `db:"support"`
	CountryCode int32  `db:"country_code"`
	//Bio            string `db:"bio"`
	About string `db:"about"`
	//State          int32  `db:"state"`
	Bot            int8   `db:"bot"`
	Banned         int64  `db:"banned"`
	BannedReason   string `db:"banned_reason"`
	AccountDaysTtl int32  `db:"account_days_ttl"`
	PhotoId        string `db:"photo_id"`
	//Deleted        int8   `db:"deleted"`
	DeletedReason string `db:"deleted_reason"`
	AddTime       int64  `db:"add_time"`
}
