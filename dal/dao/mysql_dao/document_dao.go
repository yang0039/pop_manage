package mysql_dao

import (
	"github.com/jmoiron/sqlx"
	"pop-api/dal/dataobject"
)

type DocumentDAO struct {
	db *sqlx.DB
}

func NewDocumentDAO(db *sqlx.DB) *DocumentDAO {
	return &DocumentDAO{db}
}


func (dao *DocumentDAO) SelectById(document_id int64) *dataobject.DocumentDO {
	var query = "select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, attributes, version, sse from documents where dc_id = 2 and document_id = ?"
	row := dao.db.QueryRowx(query, document_id)
	var do dataobject.DocumentDO
	err := row.StructScan(&do)
	raise(err)
	return &do
}