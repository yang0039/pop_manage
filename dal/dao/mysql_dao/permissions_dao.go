package mysql_dao

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"pop-api/dal/dataobject"
	"strconv"
)

type PermissionsDAO struct {
	db *sqlx.DB
}

func NewPermissionsDAO(db *sqlx.DB) *PermissionsDAO {
	return &PermissionsDAO{db}
}

func (dao *PermissionsDAO) GetAllPermissions() []dataobject.Permissions {
	res := make([]dataobject.Permissions, 0)
	qry := `select id, menu_id, func_name, func_title, name, title, add_time from manage_permissions;`
	rows, err := dao.db.Queryx(qry)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var p dataobject.Permissions
		err = rows.StructScan(&p)
		raise(err)
		res = append(res, p)
	}
	return res
}

func (dao *PermissionsDAO) GetPermissionsByIds(ids []int32) []dataobject.Permissions {
	res := make([]dataobject.Permissions, 0)
	if len(ids) == 0 {
		return res
	}
	queryType := "("
	for i, id := range ids {
		queryType += strconv.Itoa(int(id))
		if i == len(ids)-1 {
			queryType += ")"
		} else {
			queryType += ","
		}
	}
	qry := fmt.Sprintf(`select id, menu_id, func_name, func_title, name, title, add_time from manage_permissions where id in %s;`, queryType)
	rows, err := dao.db.Queryx(qry)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var p dataobject.Permissions
		err = rows.StructScan(&p)
		raise(err)
		res = append(res, p)
	}
	return res
}

func (dao *PermissionsDAO) GetAllMenuFunc() []dataobject.MenuFunc {
	res := make([]dataobject.MenuFunc, 0)
	qry := `select menu_id, func_name, func_title from manage_permissions group by menu_id, func_name, func_title;`
	rows, err := dao.db.Queryx(qry)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var m dataobject.MenuFunc
		err = rows.StructScan(&m)
		raise(err)
		res = append(res, m)
	}
	return res
}

// 过滤权限id
func (dao *PermissionsDAO) GetEffectPermissionIds(ids []int32) []int32 {
	res := make([]int32, 0)
	if len(ids) == 0 {
		return res
	}
	queryType := "("
	for i, id := range ids {
		queryType += strconv.Itoa(int(id))
		if i == len(ids)-1 {
			queryType += ")"
		} else {
			queryType += ","
		}
	}
	qry := fmt.Sprintf("select id from manage_permissions where id in %s;", queryType)
	rows, err := dao.db.Queryx(qry)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var id int32
		err = rows.Scan(&id)
		raise(err)
		res = append(res, id)
	}
	return res
}
