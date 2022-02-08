package dto

// 举报功能的返回
type ReportRes struct {
	ReportUser User       `json:"report_user"`
	Peer       Peer       `json:"peer"`
	Reason     int32      `json:"reason"`
	Content    string     `json:"content"`
	ReportTime int64      `json:"report_time"`
	Messages   []*Message `json:"messages"`
}

type User struct {
	UserId    int32  `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Peer struct {
	PeerType  int32  `json:"peer_type"`
	PeerId    int32  `json:"peer_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

type Message struct {
	MsgId   int32  `json:"msg_id"`
	From    User   `json:"from"` // 消息发送者id
	Peer    Peer   `json:"peer"`
	Date    int32  `json:"date"`     // 消息发送时间
	MsgType int32  `json:"msg_type"` // 消息类型
	Message string `json:"message"`  // 消息内容
	Url     string `json:"url"`
}
