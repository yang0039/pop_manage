package util

import (
	"fmt"
	"pop-api/mtproto"
)

const (
	ManageToken = "manage_token"
)

const (
	SetUserOfficial = 2001 // 设置客服
	SetUserUserName = 2002 // 更新username
	SetUserPhone    = 2003 // 更新用户手机号
	DelUser         = 2004 // 删除用户
)

const (
	QryChatDefault      = 0
	QryChatByName       = 1
	QryChatByCountry    = 2
	QryChatByNum        = 3  // 人数范围
	QryChatByCreator    = 4  // 创建人
	QryChatByManage     = 5  // 管理者
	QryChatByNote       = 6  // 备注
	QryChatById         = 7  // id
	QryChatByStatus     = 8  // 群状态
	QryChatByUserId     = 9  // 成员id
	QryChatByLabel      = 10 // 标签
	QryChatByActiveDate = 11 // 活跃日期
	QryChatByCreateDate = 12 // 创建日期
)

const (
	QryUserByName       = 1 // 昵称
	QryUserByUserName   = 2 // popid
	QryUserByPhone      = 3 // phone
	QryUserByCountry    = 4
	QryUserByEmail      = 5
	QryUserByNote       = 7
	QryUserById         = 8
	QryUserByBanned     = 9
	QryUserByLabel      = 10
	QryUserByActiveDate = 11
	QryUserByCreateDate = 12
	QryUserByDevice     = 13
	QryUserByOnline     = 14
	QryUserByOfficial   = 15
)

const (
	ALL         = 0
	MESSAGE     = 1
	DOCUMENT    = 2
	PHOTO       = 3
	VIDEO       = 4
	URL         = 5
	GIF         = 6
	VOICE       = 7
	MUSIC       = 8
	ROUND_VIDEO = 9
	GEO         = 10
	CONTACT     = 11
	OTHER       = 12
)

func ToOriginMsgType(msgType int32) int32 {
	switch msgType {
	case ALL:
		return -1
	case MESSAGE:
		return 0
	case DOCUMENT:
		return 1
	case PHOTO:
		return 2
	case VIDEO:
		return 3
	case URL:
		return 4
	case GIF:
		return 5
	case VOICE:
		return 6
	case MUSIC:
		return 7
	case ROUND_VIDEO:
		return 11
	case GEO:
		return 12
	case CONTACT:
		return 13
	case OTHER:
		return -2
	default:
		return -1
	}
}

func ToApiMsgType(msgType int32, message *mtproto.Message) (int32, string) {
	switch msgType {
	case 0:
		return MESSAGE, message.Data2.Message
	case 1:
		return DOCUMENT, message.Data2.Message
	case 2:
		return PHOTO, message.Data2.Message
	case 3:
		return VIDEO, message.Data2.Message
	case 4:
		return URL, message.Data2.Message
	case 5:
		return GIF, message.Data2.Message
	case 6:
		return VOICE, message.Data2.Message
	case 7:
		return MUSIC, message.Data2.Message
	case 8:
		return OTHER, "修改群头像"
	case 9, 10:
		return OTHER, "电话消息"
	case 11:
		return ROUND_VIDEO, message.Data2.Message
	case 12:
		return GEO, message.Data2.Message
	case 13:
		media := message.Data2.Media
		if media != nil {
			return CONTACT, fmt.Sprintf("分享联系人id:%d, 昵称:%s", media.Data2.UserId, media.Data2.FirstName+media.Data2.LastName)
		}
		return CONTACT, message.Data2.Message
	default:
		return OTHER, message.Data2.Message
	}
}

// 1:更新popid, 2:更新手机号，3:更新备注(做标记), 4:更新封号状态, 5:更新客服, 6:删除用户
const (
	UpdatePopId    = 1
	UpdatePhone    = 2
	UpdateNote     = 3 // 更新备注(做标记)
	Updatebanned   = 4 // 封号解封
	UpdateOfficial = 5 // 更新客服状态
	DeleteUser     = 6 // 删除用户
)
