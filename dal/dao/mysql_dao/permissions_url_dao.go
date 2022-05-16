package mysql_dao

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"pop-api/dal/dataobject"
)

type PermissionsUrlDAO struct {
	db *sqlx.DB
}

func NewPermissionsUrlDAO(db *sqlx.DB) *PermissionsUrlDAO {
	return &PermissionsUrlDAO{db}
}

func (dao *PermissionsUrlDAO) GetAllPermissionsUrl() []dataobject.PermissionsUrl {
	res := make([]dataobject.PermissionsUrl, 0)
	qry := `select id, permissions_id, url, method, method_name, is_effect, add_time from manage_permissions_url;`
	rows, err := dao.db.Queryx(qry)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var p dataobject.PermissionsUrl
		err = rows.StructScan(&p)
		raise(err)
		res = append(res, p)
	}
	return res
}