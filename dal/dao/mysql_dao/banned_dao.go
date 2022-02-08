package mysql_dao

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"pop-api/dal/dataobject"
)

type BannedDAO struct {
	db *sqlx.DB
}

func NewBannedDAO(db *sqlx.DB) *BannedDAO {
	return &BannedDAO{db}
}

func (dao *BannedDAO) GetBanneds(userIds []int32) map[int32]dataobject.Banned {

	res := make(map[int32]dataobject.Banned, 0)
	if len(userIds) == 0 {
		return res
	}

	sql2 := "("
	for i, id := range userIds {
		sql2 = sql2 + fmt.Sprintf("%d", id)
		if i == len(userIds)-1 {
			sql2 += ")"
		} else {
			sql2 += ","
		}
	}

	qry := fmt.Sprintf(`
	select user_id, phone, opera, banned_time, banned_reason, ip, device, state 
	from banned where state = 1 and user_id in %s;
	`, sql2)
	rows, err := dao.db.Queryx(qry)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res
	}
	raise(err)
	for rows.Next() {
		var b dataobject.Banned
		err = rows.StructScan(&b)
		raise(err)
		res[b.UserId] = b
	}
	return res
}

func (dao *BannedDAO) GetUserByBanned(limit, offset int32) ([]int32, int32) {
	uIds := make([]int32, 0)
	qry := `
	select user_id from banned where state = 1 limit ? offset ?;
	`
	rows, err := dao.db.Queryx(qry, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return uIds, 0
	}
	raise(err)
	for rows.Next() {
		var userId int32
		err = rows.Scan(&userId)
		raise(err)
		uIds = append(uIds, userId)
	}

	qryCount := "select count(*) from banned where state = 1;"
	row := dao.db.QueryRow(qryCount)
	var count int32
	err = row.Scan(&count)
	raise(err)
	return uIds, count
}

func (dao *BannedDAO) GetUserBanneds(userId, limit, offset int32) ([]map[string]interface{}, int32) {
	res := make([]map[string]interface{}, 0)
	qry := `
	select id, user_id, phone, opera,
	time, reason, ip, device, state
	from
	(
		select
			id, user_id, phone, opera,
			banned_time time, banned_reason reason,
			ip, device, 1 state
		from banned where user_id = ?
		union all
		select
			id, user_id, phone, opera,
			unbanned_time time, unbanned_reason reason,
			ip, device, 0 state
		from banned where user_id = ? and state = 0
	) data
	order by id, state limit ? offset ?;
	`
	rows, err := dao.db.Queryx(qry, userId, userId, limit, offset)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return res, 0
	}
	raise(err)
	for rows.Next() {
		var id, user_id, time, state int32
		var phone, opera, reason, ip, device string
		err = rows.Scan(&id, &user_id, &phone, &opera, &time, &reason, &ip, &device, &state)
		raise(err)
		banMap := map[string]interface{}{
			"id":         id,
			"user_id":    user_id,
			"opera":      opera,
			"opera_time": time,
			"reason":     reason,
			"ip":         ip,
			"device":     device,
			"state":      state,
		}
		res = append(res, banMap)
	}

	qryCount := `
		select count1+count2 count
		from (
			select 1 as id, count(*) count1 from banned where state = 1 and user_id = ?
		) d1
		left join (
			select 1 as id, count(*)*2 count2 from banned where state = 0 and user_id = ?
		) d2 on d1.id = d2.id;
	`
	row := dao.db.QueryRow(qryCount, userId, userId)
	var count int32
	err = row.Scan(&count)
	return res, count
}
