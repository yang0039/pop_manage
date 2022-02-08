package mysql_dao

import (
	"github.com/jmoiron/sqlx"
)

type ChannelMsgboxDAO struct {
	db *sqlx.DB
}

func NewChannelMsgboxDAO(db *sqlx.DB) *ChannelMsgboxDAO {
	return &ChannelMsgboxDAO{db}
}

// 查看活跃的用户id
func (dao *ChannelMsgboxDAO) GetSendUserIds(start, end int64) []int32 {
	return nil
}