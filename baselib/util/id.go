package util

import (
	"pop-api/baselib/snowflake"
)

var id *snowflake.IdWorker

// = &snowflake.IdWorker{
//
//}

const (
	workerId     = int64(1)
	dataCenterId = int64(1)
	twepoch      = int64(1288834974657)
)

// func init() {
// 	id, _ = snowflake.NewIdWorker(workerId, dataCenterId, twepoch)
// }

func InitSnowFlakeId(workerId, dataCenterId int64) {
	id, _ = snowflake.NewIdWorker(workerId, dataCenterId, twepoch)
}

func NextSnowflakeId() int64 {
	r, _ := id.NextId()
	return r
}

