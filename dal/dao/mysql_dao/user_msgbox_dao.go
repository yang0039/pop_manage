package mysql_dao

import (
	"github.com/jmoiron/sqlx"
)

type UserMsgboxDAO struct {
	db *sqlx.DB
}

func NewUserMsgboxDAO(db *sqlx.DB) *UserMsgboxDAO {
	return &UserMsgboxDAO{db}
}
