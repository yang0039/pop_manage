package memberapi

var (
	GetUserDialog func(uid, limit, offset int32) ([]map[string]int32, int32)
)
