package dataobject

type DocumentDO struct {
	Id               int64  `db:"id"`
	UserId           int32  `db:"user_id"`
	DocumentId       int64  `db:"document_id"`
	AccessHash       int64  `db:"access_hash"`
	DcId             int32  `db:"dc_id"`
	FilePath         string `db:"file_path"`
	FileSize         int32  `db:"file_size"`
	UploadedFileName string `db:"uploaded_file_name"`
	Ext              string `db:"ext"`
	MimeType         string `db:"mime_type"`
	ThumbId          int64  `db:"thumb_id"`
	Version          int32  `db:"version"`
	Attributes       string `db:"attributes"`
	CreatedAt        string `db:"created_at"`
	UpdatedAt        string `db:"updated_at"`

	AuthId int64 `db:"auth_id"`
	Status int8  `db:"status"`
	Sse    bool  `db:"sse"` //Server-Side Encryption
}
