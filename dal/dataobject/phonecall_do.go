package dataobject

type Call struct {
	CallId          int64  `db:"call_id"`
	AId             int32  `db:"a_id"`
	AAuthId         int64  `db:"a_auth_id"`
	AAccessHash     int64  `db:"a_access_hash"`
	ACode           string `db:"a_code"`
	BId             int32  `db:"b_id"`
	BAuthId         int64  `db:"b_auth_id"`
	BAccessHash     int64  `db:"b_access_hash"`
	BCode           string `db:"b_code"`
	UdpP2P          bool   `db:"udp_p2p"`
	UdpReflector    bool   `db:"udp_reflector"`
	MinLayer        int32  `db:"min_layer"`
	MaxLayer        int32  `db:"max_layer"`
	BMaxLayer       int32  `db:"b_max_layer"`
	GA              []byte `db:"g_a"`
	GB              []byte `db:"g_b"`
	Stage           int    `db:"stage"`
	LibraryVersions string `db:"library_versions"`
	ReceiveDate     int32  `db:"receive_date"`
	AddTime         int32  `db:"add_time"`
}
