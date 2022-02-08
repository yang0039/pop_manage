package dataobject

// user_id, phone, opera, banned_time, banned_reason, ip, device, state

type Banned struct {
	Id             int32  `db:"id" json:"id"`
	UserId         int32  `db:"user_id" json:"user_id"`
	Phone          string `db:"phone" json:"phone"`
	Opera          string `db:"opera" json:"opera"`
	BannedTime     int64  `db:"banned_time" json:"banned_time"`
	UnbannedTime   int64  `db:"unbanned_time" json:"unbanned_time"`
	BannedReason   string `db:"banned_reason" json:"banned_reason"`
	UnbannedReason string `db:"unbanned_reason" json:"unbanned_reason"`
	Ip             string `db:"ip" json:"ip"`
	Device         string `db:"device" json:"device"`
	State          int32  `db:"state" json:"state"`
}
