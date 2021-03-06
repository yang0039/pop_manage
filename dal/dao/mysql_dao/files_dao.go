package mysql_dao

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"pop-api/dal/dataobject"
)

type FilesDAO struct {
	db *sqlx.DB
}

func NewFilesDAO(db *sqlx.DB) *FilesDAO {
	return &FilesDAO{db}
}

func (dao *FilesDAO) GetUserAllFile(userId int32) []*dataobject.UserFilesDo {
	res := make([]*dataobject.UserFilesDo, 0)
	var query = "select ext, sum(file_size)/1024/1024 size, count(*) count from files where user_id = ? group by ext;"
	rows, err := dao.db.Queryx(query, userId)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var files dataobject.UserFilesDo
		err = rows.StructScan(&files)
		raise(err)
		res = append(res, &files)
	}
	return res
}

func (dao *FilesDAO) GetUserFiles(userId, offset, limit int32) ([]*dataobject.FilesDo, int32) {
	res := make([]*dataobject.FilesDo, 0)
	var query = "select id, user_id, file_id, access_hash, file_part_id, file_parts, file_size, file_path, ext, upload_name, created_at, auth_id, status, sse " +
		"from files where user_id = ? order by id desc limit ? offset ?;"
	rows, err := dao.db.Queryx(query, userId, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var files dataobject.FilesDo
		err = rows.StructScan(&files)
		raise(err)
		res = append(res, &files)
	}
	qryCount := "select count(*) from files where user_id = ?;"
	row := dao.db.QueryRow(qryCount, userId)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return res, count
}

func (dao *FilesDAO) GetAllStore(uId int32) map[string]*dataobject.UserFilesDo {
	var userQuery string
	if uId != 0 {
		userQuery = fmt.Sprintf("where user_id = %d", uId)
	}


	res := make(map[string]*dataobject.UserFilesDo, 0)
	var query = "select ext, sum(file_size)/1024/1024 size, count(*) count from files " + userQuery + " group by ext;"
	rows, err := dao.db.Queryx(query)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var files dataobject.UserFilesDo
		err = rows.StructScan(&files)
		raise(err)
		if files.Ext == "" {
			files.Ext = "other"
		}
		res[files.Ext] = &files
	}
	return res
}

func (dao *FilesDAO) GetUserRank(qType, offset, limit int32) ([]*dataobject.UserFilesDo, int32) {
	var queryType string
	// 0:?????? 1:?????? 2:?????? 3:?????? 4:?????? 5:??????
	if qType == 0 {
	} else if qType == 1 {
		queryType = "where ext in ('.jpg','.jepg','.png')"
	} else if qType == 2 {
		queryType = "where ext in ('.mp4','.avi')"
	} else if qType == 3 {
		queryType = "where ext in ('.ogg')"
	} else if qType == 4 {
		queryType = "where ext not in ('','.ogg','.jpg','.jepg','.png','.mp4','.avi')"
	} else if qType == 5 {
		queryType = "where ext in ('')"
	}


	res := make([]*dataobject.UserFilesDo, 0)
	var query = "select user_id,  sum(file_size)/1024/1024 size, count(*) count from files " + queryType + " group by user_id order by size desc limit ? offset ?;"
	rows, err := dao.db.Queryx(query, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var files dataobject.UserFilesDo
		err = rows.StructScan(&files)
		//res[files.UserId] = &files
		res  = append(res, &files)
	}

	qryCount := "select count(*) from (select user_id from files " + queryType + " group by user_id)u;"
	row := dao.db.QueryRow(qryCount)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return res, count
}

func (dao *FilesDAO) GetLastUpload(uId, qType, offset, limit int32) ([]*dataobject.FilesDo, int32) {
	var queryType string
	// 0:?????? 1:?????? 2:?????? 3:?????? 4:?????? 5:??????
	if qType == 0 {
	} else if qType == 1 {
		queryType = "where ext in ('.jpg','.jepg','.png')"
	} else if qType == 2 {
		queryType = "where ext in ('.mp4','.avi')"
	} else if qType == 3 {
		queryType = "where ext in ('.ogg')"
	} else if qType == 4 {
		queryType = "where ext not in ('','.ogg','.jpg','.jepg','.png','.mp4','.avi')"
	} else if qType == 5 {
		queryType = "where ext in ('')"
	}
	if uId != 0 {
		if queryType == "" {
			queryType += fmt.Sprintf("where user_id = %d", uId)
		} else {
			queryType += fmt.Sprintf(" and user_id = %d", uId)
		}
	}


	res := make([]*dataobject.FilesDo, 0)
	var query = "select id, user_id, file_id, access_hash, file_part_id, file_parts, file_size, file_path, ext, upload_name, created_at, auth_id, status, sse " +
		"from files " + queryType + " order by id desc limit ? offset ?;"
	rows, err := dao.db.Queryx(query, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var files dataobject.FilesDo
		err = rows.StructScan(&files)
		raise(err)
		//res[files.UserId] = &files
		res = append(res, &files)
	}

	qryCount := "select count(*) from files " + queryType
	row := dao.db.QueryRow(qryCount)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return res, count
}