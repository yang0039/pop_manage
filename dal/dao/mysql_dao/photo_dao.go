package mysql_dao

import (
	"github.com/jmoiron/sqlx"
	"pop-api/dal/dataobject"
)

type PhotoDAO struct {
	db *sqlx.DB
}

func NewPhotoDAO(db *sqlx.DB) *PhotoDAO {
	return &PhotoDAO{db}
}

func (dao *PhotoDAO) SelectByPhotoId(photo_id int64, local_id int32) *dataobject.PhotoDatasDO {
	var query = "select id, photo_id, photo_type, dc_id, volume_id, local_id, access_hash, width, height, file_size, file_path, ext, sse from photo_datas where dc_id = 2 and photo_id = ? and local_id = ?"
	var do dataobject.PhotoDatasDO
	row := dao.db.QueryRowx(query, photo_id, local_id)
	err := row.StructScan(&do)
	raise(err)
	return &do
}