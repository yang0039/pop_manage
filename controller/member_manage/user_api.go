package member_manage

import (
	"fmt"
	"math/rand"
	"pop-api/baselib/cache"
	"pop-api/baselib/redis_client"
	"pop-api/baselib/util"
	"pop-api/dal/dao"
	"time"
)

// 获取机器人(没有就创建一个)
func GetBot(userId int32) string {
	commonDao := dao.GetCommonDAO()
	token := commonDao.GetBotToken(userId)
	if token != "" {
		return token
	}

	fName := ""
	botName := "manage"
	if fName != "" {
		botName = fName + "_" + botName
	}
	token, botId := createbot(userId, botName, "manage_bot")
	redis_client.RedisCache.SAdd(fmt.Sprintf("bot:%d:uid_set", botId), userId)

	return token
}

func createbot(creatId int32, firstname, username string) (string, int32) {
	/* create bot, return token */
	//commonDao := dao.GetCommonDAO()

	tx := dao.Master.MustBegin()
	var err error
	defer func() {
		if err != nil {
			tx.Rollback()
			util.RaiseDBERR(err)
		} else {
			tx.Commit()
		}
	}()
	// user table
	access_hash := rand.Int63()
	now := int32(time.Now().Unix())
	r, err := tx.Exec(`insert into user set access_hash=?, first_name=?, username=?, phone=?, country_code=0, bot=1, account_days_ttl=0, add_time=?`,
		access_hash, firstname, username, access_hash, now)
	if err != nil {
		return "", 0
	}
	user_id, err := r.LastInsertId()
	if err != nil {
		return "", 0
	}
	_, err = tx.Exec(`update user set phone=? where id=?`, user_id, user_id)
	if err != nil {
		return "", 0
	}
	// bot table
	token := newtoken()
	_, err = tx.Exec(`insert into bot set user_id=?, info_version=1, creator_id=?, token=?, add_time=?`,
		user_id, creatId, token, now)
	if err != nil {
		return "", 0
	}
	cache.BotTokenId(token).HSet(user_id)

	return token, int32(user_id)
}

func  newtoken() string {
	now := int32(time.Now().Unix())
	return fmt.Sprintf("%d:%x", now, util.NextSnowflakeId())
}

//redisConn := redis_client.GetRedisPoolClient("memory")
//defer redisConn.Close()
//redisConn.Do("sadd", fmt.Sprintf("bot:%d:uid_set", peer.PeerId), self_id)
