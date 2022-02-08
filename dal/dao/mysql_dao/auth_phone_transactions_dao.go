package mysql_dao

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"pop-api/dal/dataobject"
)

type PhoneTransactionsDAO struct {
	db *sqlx.DB
}

func NewPhoneTransactionsDAO(db *sqlx.DB) *PhoneTransactionsDAO {
	return &PhoneTransactionsDAO{db}
}

func (dao *PhoneTransactionsDAO) GetPhoneTransactionsByPhone(phone string, limit, offset int32) ([]*dataobject.AuthPhoneTransactions, int32) {
	res := make([]*dataobject.AuthPhoneTransactions, 0)
	qry := `
	select id, auth_key_id, phone_number , code, code_expired,
	  state, attempts, created_time, is_deleted
	from auth_phone_transactions
	where phone_number = ? and is_deleted = 0 order by created_time desc limit ? offset ?;
	`
	rows, err := dao.db.Queryx(qry, phone, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		pt := &dataobject.AuthPhoneTransactions{}
		err = rows.StructScan(pt)
		raise(err)
		res = append(res, pt)
	}

	qryCount := "select count(*) from auth_phone_transactions where phone_number = ? and is_deleted = 0;"
	row := dao.db.QueryRow(qryCount, phone)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return res, count
}

func (dao *PhoneTransactionsDAO) GetPhoneTransactionsByTime(start, end int64, limit, offset int32) ([]*dataobject.AuthPhoneTransactions, int32) {
	res := make([]*dataobject.AuthPhoneTransactions, 0)
	qry := `
	select id, auth_key_id, phone_number , code, code_expired,
	  state, attempts, created_time, is_deleted
	from auth_phone_transactions
	where created_time between ? and ? and is_deleted = 0 order by created_time desc limit ? offset ?;
	`
	rows, err := dao.db.Queryx(qry, start, end, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		pt := &dataobject.AuthPhoneTransactions{}
		err = rows.StructScan(pt)
		raise(err)
		res = append(res, pt)
	}

	qryCount := "select count(*) from auth_phone_transactions where created_time >= ? and is_deleted = 0;"
	row := dao.db.QueryRow(qryCount, start)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return res, count
}



