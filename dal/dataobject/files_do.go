package dataobject

type FilesDo struct {
	Id         int64  `db:"id"`
	UserId     int32  `db:"user_id"`
	FileId     int64  `db:"file_id"`
	AccessHash int64  `db:"access_hash"`
	FilePartId int64  `db:"file_part_id"`
	FileParts  int32  `db:"file_parts"`
	FileSize   int32  `db:"file_size"`
	FilePath   string `db:"file_path"`
	Ext        string `db:"ext"`
	UploadName string `db:"upload_name"`
	CreatedAt  string `db:"created_at"`
	AuthId     int64  `db:"auth_id"`
	Status     int8   `db:"status"`
	Sse        bool   `db:"sse"` //Server-Side Encryption
}

type UserFilesDo struct {
	UserId int32   `db:"user_id"`
	Ext    string  `db:"ext"`
	Size   float64 `db:"size"`
	Count  int32   `db:"count"`
}
