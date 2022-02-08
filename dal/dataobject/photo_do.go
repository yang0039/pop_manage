package dataobject

type PhotoDatasDO struct {
	Id     int32 `db:"id"`
	UserId int32 `db:"user_id"`
	PType  int8  `db:"p_type"`

	PhotoId        int64  `db:"photo_id"`
	FilePartId     int64  `db:"file_part_id"`
	FileTotalParts int32  `db:"file_total_parts"`
	Md5CheckSum    string `db:"md5_check_sum"`

	PhotoType  int8   `db:"photo_type"`
	DcId       int32  `db:"dc_id"`
	VolumeId   int64  `db:"volume_id"`
	LocalId    int32  `db:"local_id"`
	AccessHash int64  `db:"access_hash"`
	Width      int32  `db:"width"`
	Height     int32  `db:"height"`
	FileSize   int32  `db:"file_size"`
	FilePath   string `db:"file_path"`
	Ext        string `db:"ext"`
	CreatedAt  string `db:"created_at"`

	Status int8 `db:"status"`
	Sse    bool `db:"sse"` //Server-Side Encryption
}
