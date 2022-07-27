package mysql_dao

import (
	"database/sql"
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
	var query = "select id, photo_id, file_part_id, photo_type, dc_id, volume_id, local_id, access_hash, width, height, file_size, file_path, ext, sse from photo_datas where dc_id = 2 and photo_id = ? and local_id = ?"
	var do dataobject.PhotoDatasDO
	row := dao.db.QueryRowx(query, photo_id, local_id)
	err := row.StructScan(&do)
	raise(err)
	return &do
}

func (dao *PhotoDAO) SelectPhotosById(photo_id int64) []*dataobject.PhotoDatasDO {
	res := make([]*dataobject.PhotoDatasDO, 0)
	var query = "select id, photo_id, photo_type, dc_id, volume_id, local_id, access_hash, width, height, file_size, file_path, ext, sse from photo_datas where photo_id = ?;"
	rows, err := dao.db.Queryx(query, photo_id)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var p dataobject.PhotoDatasDO
		err = rows.StructScan(&p)
		raise(err)
		res = append(res, &p)
	}
	return res
}