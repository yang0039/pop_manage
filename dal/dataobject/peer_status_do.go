package dataobject

type PeerStatus struct {
	Id       int32  `db:"id" json:"id"`
	PeerType int32  `db:"peer_type" json:"peer_type"`
	PeerId   int32  `db:"peer_id" json:"peer_id"`
	Status   int32  `db:"status" json:"status"`
	Util     int64  `db:"util" json:"util"`
	Note     string `db:"note" json:"note"`
	Operator string `db:"operator" json:"operator"`
	AddTime  int64  `db:"add_time" json:"add_time"`
}

type PeerStatusRecord struct {
	Id        int32  `db:"id" json:"id"`
	PeerType  int32  `db:"peer_type" json:"peer_type"`
	PeerId    int32  `db:"peer_id" json:"peer_id"`
	Status    int32  `db:"status" json:"status"`
	Util      int64  `db:"util" json:"util"`
	Note      string `db:"note" json:"note"`
	Opera     string `db:"opera" json:"opera"`
	OperaTime int64  `db:"opera_time" json:"opera_time"`
}
