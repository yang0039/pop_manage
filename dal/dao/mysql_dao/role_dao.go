package mysql_dao

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
	"time"
)

type RoleDAO struct {
	db *sqlx.DB
}

func NewRoleDAO(db *sqlx.DB) *RoleDAO {
	return &RoleDAO{db}
}

func (dao *RoleDAO) AddRole(name, operator string, permissionIds []int32) int64 {
	now := time.Now().Unix()
	sqlStr := "insert into manage_role (name, operator, is_effect, add_time) values (?, ?, ?, ?);"
	tx,err := dao.db.Begin()
	raiseTx(err, tx)
	result,err := tx.Exec(sqlStr, name, operator, 1, now)
	raiseTx(err, tx)
	roleId,_ := result.LastInsertId()

	if len(permissionIds) == 0 {
		tx.Commit()
		return roleId
	}
	var perType string
	for i, id := range permissionIds {
		q := fmt.Sprintf("(%d, %d)", roleId, id)
		if i == len(permissionIds) - 1 {
			q += ";"
		} else {
			q += ","
		}
		perType += q
	}
	sqlStr2 := fmt.Sprintf("insert into manage_role_permissions (role_id, permissions_id) values %s", perType)
	_,err = tx.Exec(sqlStr2)
	raiseTx(err, tx)
	tx.Commit()
	return roleId
}

func (dao *RoleDAO) RoleIsEffect(id int32) bool {
	qry := "select id from manage_role where id = ? and is_effect = 1 limit 1;"
	row := dao.db.QueryRowx(qry, id)
	var rId int32
	row.Scan(&rId)
	return id == rId
}

func (dao *RoleDAO) EditRole(roleId int32, permissionIds []int32) {
	tx,err := dao.db.Begin()
	raiseTx(err, tx)
	sqlStr := "delete from manage_role_permissions where role_id = ?;"
	_,err = tx.Exec(sqlStr, roleId)
	raiseTx(err, tx)
	if len(permissionIds) == 0 {
		tx.Commit()
		return
	}
	var perType string
	for i, id := range permissionIds {
		q := fmt.Sprintf("(%d, %d)", roleId, id)
		if i == len(permissionIds) - 1 {
			q += ";"
		} else {
			q += ","
		}
		perType += q
	}
	sqlStr2 := fmt.Sprintf("insert into manage_role_permissions (role_id, permissions_id) values %s", perType)
	_,err = tx.Exec(sqlStr2)
	raiseTx(err, tx)
	tx.Commit()
}

func (dao *RoleDAO) DeleteRole(roleId int32) {
	sqlStr := "update manage_role set is_effect = 0 where id = ?;"
	_, err := dao.db.Exec(sqlStr, roleId)
	raise(err)
}


func (dao *RoleDAO) GetAllRole() []map[string]interface{} {
	res := make([]map[string]interface{}, 0)
	qry := "select id, name from manage_role where is_effect = 1;"
	rows, err := dao.db.Queryx(qry)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var id int32
		var name string
		err = rows.Scan(&id, &name)
		raise(err)
		m := map[string]interface{}{
			"role_id": id,
			"name": name,
		}
		res = append(res, m)
	}
	return res
}

func (dao *RoleDAO) GetAllRolePermission() map[int32][]int32{
	res := make(map[int32][]int32, 0)
	qry := "select role_id, permissions_id from manage_role_permissions;"
	rows, err := dao.db.Queryx(qry)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var rId, pId int32
		err = rows.Scan(&rId, &pId)
		raise(err)
		_, ok :=  res[rId]
		if !ok {
			res[rId] = make([]int32, 0)
		}
		res[rId] = append(res[rId], pId)
	}
	return res
}

// 过滤角色id
func (dao *RoleDAO) GetEffectRoleIds(ids []int32) []int32 {
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
	qry := fmt.Sprintf("select id from manage_role where is_effect= 1 and id in %s;", queryType)
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

// 获取角色的权限
func (dao *RoleDAO) GetRolePermissionIds(ids []int32) []int32 {
	res := make([]int32, 0)
	if len(ids) == 0 {
		return res
	}
	qryType := "("

	for i, id := range ids {
		qryType += strconv.Itoa(int(id))
		if i == len(ids) - 1 {
			qryType += ")"
		} else {
			qryType += ","
		}
	}

	sqlStr := fmt.Sprintf("select permissions_id from manage_role_permissions where role_id in %s;", qryType)
	rows, err := dao.db.Queryx(sqlStr)
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