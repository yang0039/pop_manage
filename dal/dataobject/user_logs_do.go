package dataobject

// select user_id, auth_id, ip, device_model, system_version, app_version, created_at from user_logs2;

type UserLogsDo struct {
	UserId        int32  `db:"user_id" json:"user_id"`
	AuthId        int64  `db:"auth_id" json:"auth_id"`
	Ip            string `db:"ip" json:"ip"`
	DeviceModel   string `db:"device_model" json:"device_model"`
	SystemVersion string `db:"system_version" json:"system_version"`
	AppVersion    string `db:"app_version" json:"app_version"`
	CreatedAt     string `db:"created_at" json:"created_at"`
}
