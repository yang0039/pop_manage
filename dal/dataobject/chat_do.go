package dataobject

type Chat struct {
	Id                 int32  `db:"id"`
	AccessHash         int64  `db:"access_hash"`
	Democracy          bool   `db:"democracy"`
	CreatorId          int32  `db:"creator_id"`
	PinnedMsgId        int32  `db:"pinned_msg_id"`
	About              string `db:"about"`
	Title              string `db:"title"`
	Type               int    `db:"type"`
	PhotoId            int64  `db:"photo_id"`
	AdminsEnabled      bool   `db:"admins_enabled"`
	MigratedFromChatId int32  `db:"migrated_from_chat_id"`
	MigratedFromMaxId  int32  `db:"migrated_from_max_id"`
	MigratedTo         int32  `db:"migrated_to"`
	Username           string `db:"username"`
	HiddenPrehistory   bool   `db:"hidden_prehistory"`
	Signatures         bool   `db:"signatures"`
	Deactivated        bool   `db:"deactivated"`
	Verified           bool   `db:"verified"`
	RightsMask         int32  `db:"rights_mask"`
	UntilDate          int32  `db:"until_date"`
	SlowmodeSeconds    int32  `db:"slowmode_seconds"`
	Version            int32  `db:"version"`
	StickerSetId       int64  `db:"sticker_set_id"`
	AddTime            int32  `db:"add_time"`
}
