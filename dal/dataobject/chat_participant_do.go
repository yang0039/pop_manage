package dataobject

type ParticipantRow struct {
	ChatId         int32  `db:"chat_id"`
	UserID         int32  `db:"user_id"`
	Type           int    `db:"type"`
	Rank           string `db:"rank"`
	Bot            bool   `db:"bot"`
	InviterId      int32  `db:"inviter_id"`
	RightsMask     int32  `db:"rights_mask"`
	UntilDate      int32  `db:"until_date"`
	Kicked         bool   `db:"kicked"`
	AvailableMinId int32  `db:"available_min_id"`
	PromotedBy     int32  `db:"promoted_by"`
	AddTime        int32  `db:"add_time"`
	UpdateTime     int32  `db:"update_time"`
}

type ChatParticipantInfo struct {
	ChatId       int32  `db:"chat_id" json:"chat_id"`
	UserId       int32  `db:"user_id" json:"user_id"`
	Type         int32  `db:"type" json:"type"`
	AddTime      int64  `db:"add_time" json:"add_time"`
	FirstName    string `db:"first_name" json:"first_name"`
	LastName     string `db:"last_name" json:"last_name"`
	Username     string `db:"username" json:"username"`
	Phone        string `db:"phone" json:"phone"`
	CountryCode  string `db:"country_code" json:"country_code"`
	RegisterTime int64  `db:"register_time" json:"register_time"`
}
