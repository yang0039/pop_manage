package mysql_dao

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"pop-api/dal/dataobject"
)

type MenuDAO struct {
	db *sqlx.DB
}

func NewMenuDAO(db *sqlx.DB) *MenuDAO {
	return &MenuDAO{db}
}

func (dao *MenuDAO) GetAllMenu() []dataobject.Menu {
	res := make([]dataobject.Menu, 0)
	qry := `select id, name, title, add_time from manage_menu;`
	rows, err := dao.db.Queryx(qry)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var m dataobject.Menu
		err = rows.StructScan(&m)
		raise(err)
		res = append(res, m)
	}
	return res
}

func (dao *MenuDAO) GetMenuMap() map[int32]dataobject.Menu {
	res := make(map[int32]dataobject.Menu, 0)
	qry := `select id, name, title, add_time from manage_menu;`
	rows, err := dao.db.Queryx(qry)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var m dataobject.Menu
		err = rows.StructScan(&m)
		raise(err)
		//res = append(res, m)
		res[m.Id] = m
	}
	return res
}