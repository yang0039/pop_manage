package cache

import "fmt"

func BotTokenId(token string) *Key {
	/* bot_token对应的bot_id */
	return &Key{
		DBName: "cache",
		Key:    "bot:token:map",
		Field:  token,
	}
}

func BotMap(bot_id int32) *Key {
	/* bot map */
	return &Key{
		DBName: "cache",
		Key:    fmt.Sprintf("bot:%d:map", bot_id),
	}
}

func BotWebhook(bot_id int32) *Key {
	/* bot webhook */
	return &Key{
		DBName: "cache",
		Key:    fmt.Sprintf("bot:%d:map", bot_id),
		Field:  "webhook",
	}
}

func BotUidSet(bot_id int32) *Key {
	/* bot map */
	return &Key{
		DBName: "cache",
		Key:    fmt.Sprintf("bot:%d:uid_set", bot_id),
	}
}
