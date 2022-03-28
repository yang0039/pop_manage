package mysql_dao

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"pop-api/dal/dataobject"
	"strconv"
)

type UserDAO struct {
	db *sqlx.DB
}

func NewUserDAO(db *sqlx.DB) *UserDAO {
	return &UserDAO{db}
}

// 获取总人数
func (dao *UserDAO) GetUserNum(start int64) (num int32) {
	var query = "select count(*) count from user where add_time >= ?;"
	row := dao.db.QueryRowx(query, start)
	err := row.Scan(&num)
	raise(err)
	return
}

// 获取用户基本信息
func (dao *UserDAO) GetUsers(userIds []int32) []*dataobject.UserDO {
	res := make([]*dataobject.UserDO, 0, len(userIds))
	if len(userIds) == 0 {
		return res
	}

	queryType := "("
	for i, id := range userIds {
		queryType += strconv.Itoa(int(id))
		if i == len(userIds)-1 {
			queryType += ")"
		} else {
			queryType += ","
		}
	}

	qry := fmt.Sprintf(`
	select
	  id, access_hash, first_name, last_name, username, phone,
	  country_code, about, bot, photo_id, support, add_time
	from user
	where id in %s;
	`, queryType)
	rows, err := dao.db.Queryx(qry)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		u := &dataobject.UserDO{}
		err = rows.StructScan(u)
		raise(err)
		res = append(res, u)
	}
	return res
}

/*
func (dao *UserDAO) GetUserIdsByName(name string, limit, offset int32) ([]int32, int32) {
	res := make([]int32, 0)
	if name == "" {
		return res, 0
	}
	qry := "select id from user where concat(first_name, ' ', last_name) = ? limit ? offset ?;"
	rows, err := dao.db.Queryx(qry, name, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var id int32
		err = rows.Scan(&id)
		raise(err)
		res = append(res, id)
	}

	qryCount := "select count(*) from user where concat(first_name, last_name) = ?;"
	row := dao.db.QueryRow(qryCount, name)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return res, count
}
*/

func (dao *UserDAO) GetUserIdsByUserName(uname string, limit, offset int32) ([]int32, int32) {
	res := make([]int32, 0)
	if uname == "" {
		return res, 0
	}
	qry := "select id from user where username = ? limit ? offset ?;"
	rows, err := dao.db.Queryx(qry, uname, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var id int32
		err = rows.Scan(&id)
		raise(err)
		res = append(res, id)
	}

	qryCount := "select count(*) from user where username = ?;"
	row := dao.db.QueryRow(qryCount, uname)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return res, count
}

func (dao *UserDAO) GetUserIdsByPhone(phone string, limit, offset int32) ([]int32, int32) {
	res := make([]int32, 0)
	if phone == "" {
		return res, 0
	}
	phone = "%" + phone + "%"
	qry := "select id from user where phone like ? limit ? offset ?;"
	rows, err := dao.db.Queryx(qry, phone, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var id int32
		err = rows.Scan(&id)
		raise(err)
		res = append(res, id)
	}

	qryCount := "select count(*) from user where phone like ?;"
	row := dao.db.QueryRow(qryCount, phone)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return res, count
}

func (dao *UserDAO) GetUserIdsByCountryCode(code string, limit, offset int32) ([]int32, int32) {
	res := make([]int32, 0)
	if code == "" {
		return res, 0
	}
	qry := "select id from user where country_code = ? limit ? offset ?;"
	rows, err := dao.db.Queryx(qry, code, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var id int32
		err = rows.Scan(&id)
		raise(err)
		res = append(res, id)
	}

	qryCount := "select count(*) from user where country_code = ?;"
	row := dao.db.QueryRow(qryCount, code)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return res, count
}

func (dao *UserDAO) GetUserIdsDefault(limit, offset int32) ([]int32, int32) {
	res := make([]int32, 0)
	qry := "select id from user order by add_time desc limit ? offset ?;"
	rows, err := dao.db.Queryx(qry, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var id int32
		err = rows.Scan(&id)
		raise(err)
		res = append(res, id)
	}

	qryCount := "select count(*) from user;"
	row := dao.db.QueryRow(qryCount)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return res, count
}

func (dao *UserDAO) GetUserIdsNoBanned(limit, offset int32) ([]int32, int32) {
	res := make([]int32, 0)
	qry := `
	select id from user  u
	left join (
			select user_id
			from banned where state = 1
	) b on u.id = b.user_id
	where b.user_id is null
	order by id desc limit ? offset ?;
	`
	rows, err := dao.db.Queryx(qry, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var id int32
		err = rows.Scan(&id)
		raise(err)
		res = append(res, id)
	}

	qryCount := "select count(*) from user;"
	row := dao.db.QueryRow(qryCount)
	var count int32
	err = row.Scan(&count)
	raise(err)

	qryBanCount := "select count(*) from banned where state = 1;"
	row2 := dao.db.QueryRow(qryBanCount)
	var count2 int32
	err = row2.Scan(&count2)
	raise(err)

	return res, count - count2
}

// 检查用户是否存在
func (dao *UserDAO) CheckUser(userId int32) bool {
	var query = "select id from user where id = ? and bot = 0;"
	row := dao.db.QueryRowx(query, userId)
	var uId int32
	err := row.Scan(&uId)
	if err == sql.ErrNoRows {
		return false
	}
	raise(err)
	if uId == 0 {
		return false
	}
	return true
}


func (dao *UserDAO) GetUserIdsByName(name string, limit, offset int32) ([]int32, int32) {
	res := make([]int32, 0)
	if name == "" {
		return res, 0
	}
	qry := "select id from user where concat(first_name, ' ', last_name) = ? || concat(first_name, last_name) = ? limit ? offset ?;"
	rows, err := dao.db.Queryx(qry, name, name, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var id int32
		err = rows.Scan(&id)
		raise(err)
		res = append(res, id)
	}

	qryCount := "select count(*) from user where concat(first_name, ' ', last_name) = ? || concat(first_name, last_name) = ?;"
	row := dao.db.QueryRow(qryCount, name, name)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return res, count
}

func (dao *UserDAO) GetUserIdsByCreate(start, end int64, limit, offset int32) ([]int32, int32) {
	res := make([]int32, 0)
	if end == 0 {
		return res, 0
	}
	qry := "select id from user where add_time between ? and ? limit ? offset ?;"
	rows, err := dao.db.Queryx(qry, start, end, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var id int32
		err = rows.Scan(&id)
		raise(err)
		res = append(res, id)
	}

	qryCount := "select count(*) from user where add_time between ? and ?;"
	row := dao.db.QueryRow(qryCount, start, end)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return res, count
}

func (dao *UserDAO) GetUserIdsNotIn(uIds[]int32, limit, offset int32) ([]int32, int32) {
	res := make([]int32, 0)
	uidStr := "("
	if len(uIds) == 0 {
		uidStr = "(0)"
	} else {
		for i, uId := range uIds {
			if i == len(uIds) - 1 {
				uidStr += fmt.Sprintf("%d)", uId)
			}  else {
				uidStr += fmt.Sprintf("%d,", uId)
			}
		}
	}

	qry := fmt.Sprintf("select id from user where id not in %s limit ? offset ?;", uidStr)

	rows, err := dao.db.Queryx(qry, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var id int32
		err = rows.Scan(&id)
		raise(err)
		res = append(res, id)
	}

	qryCount := fmt.Sprintf("select id from user where id not in %s;", uidStr)
	row := dao.db.QueryRow(qryCount)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return res, count
}

// 查询客服
func (dao *UserDAO) GetUserIdsOfficial(limit, offset int32) ([]int32, int32) {
	res := make([]int32, 0)
	qry := "select id from user where support = 1 order by add_time desc limit ? offset ?;"
	rows, err := dao.db.Queryx(qry, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var id int32
		err = rows.Scan(&id)
		raise(err)
		res = append(res, id)
	}

	qryCount := "select count(*) from user where support = 1;"
	row := dao.db.QueryRow(qryCount)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return res, count
}