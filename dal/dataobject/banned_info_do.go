package dataobject

type BannedInfo struct {
	Id             int32  `db:"id" json:"id"`
	UserId         int32  `db:"user_id" json:"user_id"`
	AuthId         int64  `db:"auth_id" json:"auth_id"`
	Model          string `db:"model" json:"model"`
	SystemVersion  string `db:"system_version" json:"system_version"`
	AppVersion     string `db:"app_version" json:"app_version"`
	SystemLangCode string `db:"system_lang_code" json:"system_lang_code"`
	LangPack       string `db:"lang_pack" json:"lang_pack"`
	LangCode       string `db:"lang_code" json:"lang_code"`
	Ip             string `db:"ip" json:"ip"`
	Layer          int32  `db:"layer" json:"layer"`
	DateCreated    int64  `db:"date_created" json:"date_created"`
	DateActivate   int64  `db:"date_activate" json:"date_activate"`
}
