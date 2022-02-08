package mysql_dao

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"pop-api/dal/dataobject"
	"time"
)

type AccountDAO struct {
	db *sqlx.DB
}

func NewAccountDAO(db *sqlx.DB) *AccountDAO {
	return &AccountDAO{db}
}

func (dao *AccountDAO) AccountIsExit(name string) bool {
	qry := "select account_name from manage_account where account_name = ? and is_effect = 1 limit 1;"
	row := dao.db.QueryRowx(qry, name)
	var n string
	row.Scan(&n)
	return name == n
}

func (dao *AccountDAO) AccountIsExitById(id int32) bool {
	qry := "select id from manage_account where id = ? and is_effect = 1 limit 1;"
	row := dao.db.QueryRowx(qry, id)
	var n int32
	row.Scan(&n)
	return id == n
}

func (dao *AccountDAO) AccountIsExitById2(id int32) bool {
	qry := "select id from manage_account where id = ? limit 1;"
	row := dao.db.QueryRowx(qry, id)
	var n int32
	row.Scan(&n)
	return id == n
}

// 新增角色
func (dao *AccountDAO) AddAccount(accName, userName, pwd, operator string, pwd_util int32, roleIds []int32)  int32 {
	now := time.Now().Unix()
	tx, err := dao.db.Begin()
	raiseTx(err, tx)
	sqlStr := "insert into manage_account (account_name, user_name, pwd, pwd_util, operator, is_effect, add_time) values (?, ?, ?, ?, ?, 1, ?);"
	result, err := tx.Exec(sqlStr, accName, userName, pwd, pwd_util, operator, now)
	raiseTx(err, tx)
	accId,_ := result.LastInsertId()
	if len(roleIds) == 0 {
		tx.Commit()
		return int32(accId)
	}

	var perType string
	for i, id := range roleIds {
		q := fmt.Sprintf("(%d, %d)", accId, id)
		if i == len(roleIds) - 1 {
			q += ";"
		} else {
			q += ","
		}
		perType += q
	}

	sqlStr2 := fmt.Sprintf("insert into manage_account_role (account_id, role_id) values %s", perType)
	_,err = tx.Exec(sqlStr2)
	raiseTx(err, tx)
	tx.Commit()
	return int32(accId)
}

// 获取密码
func (dao *AccountDAO) QryAccountPwd(accId int32)  (string, int32) {
	qry := "select pwd, pwd_util from manage_account where id = ? and is_effect = 1 limit 1;"
	row := dao.db.QueryRowx(qry, accId)
	var pwd string
	var pwdUtil int32
	err := row.Scan(&pwd, &pwdUtil)
	raise(err)
	return pwd, pwdUtil
}

// 编辑角色
func (dao *AccountDAO) EditAccountRole(accId int32, roleIds[] int32) {
	tx,err := dao.db.Begin()
	raiseTx(err, tx)
	sqlStr := "delete from manage_account_role where account_id = ?;"
	_,err = tx.Exec(sqlStr, accId)
	raiseTx(err, tx)
	if len(roleIds) == 0 {
		tx.Commit()
		return
	}
	var perType string
	for i, id := range roleIds {
		q := fmt.Sprintf("(%d, %d)", accId, id)
		if i == len(roleIds) - 1 {
			q += ";"
		} else {
			q += ","
		}
		perType += q
	}
	sqlStr2 := fmt.Sprintf("insert into manage_account_role (account_id, role_id) values %s", perType)
	_,err = tx.Exec(sqlStr2)
	raiseTx(err, tx)
	tx.Commit()
}

func (dao *AccountDAO) EditAccountState(accId, opera int32) {
	sqlStr := "update manage_account set is_effect = ? where id = ?;"
	_, err := dao.db.Exec(sqlStr, opera, accId)
	raise(err)
}

func (dao *AccountDAO) EditAccountPwd(accId int32, newPwd string) {
	sqlStr := "update manage_account set pwd = ? where id = ?;"
	_, err := dao.db.Exec(sqlStr, newPwd, accId)
	raise(err)
}

func (dao *AccountDAO) EditAccountPwdUtil(accId, pwdUtil int32) {
	sqlStr := "update manage_account set pwd_util = ? where id = ?;"
	_, err := dao.db.Exec(sqlStr, pwdUtil, accId)
	raise(err)
}

func (dao *AccountDAO) GetAccountRoleIds(accId int32) []int32 {
	res :=make([]int32, 0)
	sqlStr := "select role_id from manage_account_role where account_id = ?;"
	rows, err := dao.db.Queryx(sqlStr, accId)
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

func (dao *AccountDAO) GetAccountByName(name string) dataobject.Account {
	sqlStr := "select id, account_name, user_name, pwd, pwd_util, is_effect, add_time from manage_account where account_name = ? and is_effect = 1;"
	row := dao.db.QueryRowx(sqlStr, name)
	var account dataobject.Account
	err := row.StructScan(&account)
	if err == sql.ErrNoRows {
		return account
	}
	raise(err)
	return account
}

func (dao *AccountDAO) GetAccountById(id int32) dataobject.Account {
	sqlStr := "select id, account_name, user_name, pwd, pwd_util, is_effect, add_time from manage_account where id = ? and is_effect = 1;"
	row := dao.db.QueryRowx(sqlStr, id)
	var account dataobject.Account
	err := row.StructScan(&account)
	if err == sql.ErrNoRows {
		return account
	}
	raise(err)
	return account
}


func (dao *AccountDAO) GetAllAccount() []dataobject.Account {
	res :=make([]dataobject.Account, 0)
	sqlStr := "select id, account_name, user_name, pwd_util, is_effect, add_time from manage_account;"
	rows, err := dao.db.Queryx(sqlStr)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var account dataobject.Account
		err = rows.StructScan(&account)
		raise(err)
		res = append(res, account)
	}
	return res
}

func (dao *AccountDAO) GetAccountRoleName() (map[int32][]int32, map[int32][]string) {
	resId := make(map[int32][]int32, 0)
	resName := make(map[int32][]string, 0)
	sqlStr := `
	select a.id, ifnull(r.id,'') r_id, ifnull(r.name,'') name
	from manage_account a
	left join manage_account_role ar on a.id = ar.account_id
	left join manage_role r on ar.role_id = r.id
	where r.is_effect = 1;
	`
	rows, err := dao.db.Queryx(sqlStr)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return resId, resName
	}
	raise(err)
	for rows.Next() {
		var id, rId int32
		var name string
		err = rows.Scan(&id, &rId, &name)
		raise(err)
		_, ok := resId[id]
		if !ok {
			resId[id] = make([]int32, 0)
			resName[id] = make([]string, 0)
		}
		resId[id] = append(resId[id], rId)
		resName[id] = append(resName[id], name)
	}
	return resId, resName
}

func (dao *AccountDAO) GetAccountRole(id int32) []map[string]interface{} {
	res :=make([]map[string]interface{}, 0)
	sqlStr := `
	select ifnull(r.id,0) id,
		ifnull(r.name,'') name
	from manage_account a
	left join manage_account_role ar on a.id = ar.account_id
	left join manage_role r on ar.role_id = r.id
	where r.is_effect = 1 and a.id = ?;
	`
	rows, err := dao.db.Queryx(sqlStr, id)
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
			"id": id,
			"name": name,
		}
		res = append(res, m)
	}
	return res
}

// 删除账户
func (dao *AccountDAO) DelAccount(accId int32) {
	tx,err := dao.db.Begin()
	raiseTx(err, tx)
	sqlStr := "delete from manage_account_role where account_id = ?;"
	_,err = tx.Exec(sqlStr, accId)
	raiseTx(err, tx)

	sqlStr2 := "delete from manage_account where id = ?;"
	_,err = tx.Exec(sqlStr2, accId)
	raiseTx(err, tx)
	tx.Commit()
}

// 获取超级管理员id
func (dao *AccountDAO) QryAdminId()  int32 {
	qry := "select id from manage_account where account_name = 'admin';"
	row := dao.db.QueryRowx(qry)
	var adminId int32
	err := row.Scan(&adminId)
	raise(err)
	return adminId
}