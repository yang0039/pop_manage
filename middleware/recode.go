package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pop-api/baselib/logger"
	"pop-api/baselib/util"
	"pop-api/dto"
	"pop-api/public"
	"strconv"
	"strings"
	"time"
)

func getUrlData(c *gin.Context) (interface{}, string, error) {
	switch c.FullPath() {
	case "/auth/login":
		data := &dto.Account{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("账户名:%s", data.Name)
		return data, reqData, nil
	case "/dashboard/total_data":
		return nil, "", nil
	case "/dashboard/active_data":
		return nil, "", nil
	case "/dashboard/max_member_chat":
		data := &dto.QryType{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		if data.Limit == 0 {
			data.Limit = 10
		}
		reqData := fmt.Sprintf("请求页码:%d, 请求数量:%d", (data.Offset/data.Limit)+1, data.Limit)
		return data, reqData, nil

	case "/chat_manage/update_note":
		data := &dto.NoteObj{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("群id:%d, 标签id:%s, 备注:%s", data.PeerId, data.LabelId, data.Note)
		return data, reqData, nil
	case "/chat_manage/query_chat":
		data := &dto.QryType{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := qryChatInfo(data)
		return data, reqData, nil
	case "/chat_manage/query_chat_member":
		data := &dto.QryChat{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("群id:%d", data.ChatId)
		return data, reqData, nil
	case "/chat_manage/query_chat_msg":
		data := &dto.ChatMsg{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("群id:%d", data.ChatId)
		if data.Limit == 0 {
			data.Limit = 20
		}
		reqData += fmt.Sprintf(", 请求页码:%d, 请求数量:%d", (data.Offset/data.Limit)+1, data.Limit)
		return data, reqData, nil
	case "/chat_manage/update_chat_status":
		data := &dto.ChatStatus{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("群id:%d", data.ChatId)
		var status string
		if data.Status == 1 {
			status = "正常"
		} else if data.Status == 2 {
			status = "警告"
		} else if data.Status == 3 {
			status = "短期禁言"
			status += fmt.Sprintf(", 天数:%d", data.Days)
		} else if data.Status == 4 {
			status = "长期禁言"
		}
		reqData += fmt.Sprintf(", 类型:%s", status)
		if data.Note != "" {
			reqData += fmt.Sprintf(", 备注:%s", data.Note)
		}
		return data, reqData, nil
	case "/chat_manage/qry_status_record":
		data := &dto.ChatStatus{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("群id:%d", data.ChatId)
		if data.Limit == 0 {
			data.Limit = 20
		}
		reqData += fmt.Sprintf(", 请求页码:%d, 请求数量:%d", (data.Offset/data.Limit)+1, data.Limit)
		return data, reqData, nil
	case "/chat_manage/delete_chat_history":
		data := &dto.QryChat{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("群id:%d", data.ChatId)
		return data, reqData, nil
	case "/chat_manage/delete_chat":
		data := &dto.QryChat{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("群id:%d", data.ChatId)
		return data, reqData, nil

	case "/user_manage/update_note":
		data := &dto.NoteObj{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("用户id:%d, 标签id:%s, 备注:%s", data.PeerId, data.LabelId, data.Note)
		return data, reqData, nil
	case "/user_manage/query_user":
		data := &dto.QryType{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := qryUserInfo(data)
		return data, reqData, nil
	case "/user_manage/query_user_chat":
		data := &dto.UserChat{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		var chatStatus, partStatus string
		if data.ChatType == 1 {
			chatStatus = "群"
		} else if data.ChatType == 2 {
			chatStatus = "频道"
		} else if data.ChatType == 3 {
			chatStatus = "群发"
		} else {
			chatStatus = "全部"
		}

		if data.Type == 0 {
			partStatus = "拥有的群"
		} else if data.Type == 1 {
			partStatus = "管理的群"
		} else if data.Type == 2 {
			partStatus = "参与的群"
		}
		reqData := fmt.Sprintf("用户id:%d, 群类型:%s, 参与类型:%s", data.UserId, chatStatus, partStatus)
		if data.Limit == 0 {
			data.Limit = 20
		}
		reqData += fmt.Sprintf(", 请求页码:%d, 请求数量:%d", (data.Offset/data.Limit)+1, data.Limit)
		return data, reqData, nil
	case "/user_manage/delete_user":
		data := &dto.UpdateUserPhone{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("用户id:%d", data.UserId)
		return data, reqData, nil
	case "/user_manage/delete_user_pwd":
		data := &dto.QryUser{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("用户id:%d", data.UserId)
		return data, reqData, nil
	case "/user_manage/banned_user":
		data := &dto.Banned{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		var operaString string
		if data.OperaType == 0 {
			operaString = "解封"
		} else if data.OperaType == 1 {
			operaString = "封号"
		}
		reqData := fmt.Sprintf("用户id:%d, 操作类型:%s, 原因:%s", data.UserId, operaString, data.Reason)
		return data, reqData, nil
	case "/user_manage/query_phone_transaction":
		data := &dto.PhoneTransaction{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		var selectType, qry string
		if data.Type == 0 {  // 手机号
			selectType = "手机号"
			qry = data.Qry
		} else {  //时间
			selectType = "时间"
			t,_ := strconv.Atoi(data.Qry)
			qry = time.Unix(int64(t), 0).Format("2006-01-02 15:04:05")
		}
		reqData := fmt.Sprintf("请求类型:%s, 请求参数:%s", selectType, qry)
		if data.Limit == 0 {
			data.Limit = 20
		}
		reqData += fmt.Sprintf(", 请求页码:%d, 请求数量:%d", (data.Offset/data.Limit)+1, data.Limit)
		return data, reqData, nil
	case "/user_manage/query_user_report": // 举报管理
		data := &dto.QryType{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		//  0:全部 1：垃圾 2：暴力 3：色情 4：虐待 5：版权 6：其他
		var repType string
		if data.Type == 0 {
			repType = "全部"
		} else if data.Type == 1 {
			repType = "垃圾信息"
		} else if data.Type == 2 {
			repType = "暴力"
		} else if data.Type == 3 {
			repType = "色情"
		} else if data.Type == 4 {
			repType = "儿童虐待"
		} else if data.Type == 5 {
			repType = "版权"
		} else if data.Type == 6 {
			repType = "其他"
		}
		reqData := fmt.Sprintf("举报类型:%s", repType)
		if data.Limit == 0 {
			data.Limit = 20
		}
		reqData += fmt.Sprintf(", 请求页码:%d, 请求数量:%d", (data.Offset/data.Limit)+1, data.Limit)
		return data, reqData, nil
	case "/user_manage/query_user_contact":
		data := &dto.QryUser{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("用户id:%d", data.UserId)
		if data.Limit == 0 {
			data.Limit = 20
		}
		reqData += fmt.Sprintf(", 请求页码:%d, 请求数量:%d", (data.Offset/data.Limit)+1, data.Limit)
		return data, reqData, nil
	case "/user_manage/query_user_dialogs":
		data := &dto.QryUser{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("用户id:%d", data.UserId)
		if data.Limit == 0 {
			data.Limit = 20
		}
		reqData += fmt.Sprintf(", 请求页码:%d, 请求数量:%d", (data.Offset/data.Limit)+1, data.Limit)
		return data, reqData, nil
	case "/user_manage/query_user_relation":
		data := &dto.QryUser{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("用户id:%d, 对方id:%d", data.UserId, data.PeerId)
		if data.Limit == 0 {
			data.Limit = 20
		}
		reqData += fmt.Sprintf(", 请求页码:%d, 请求数量:%d", (data.Offset/data.Limit)+1, data.Limit)
		return data, reqData, nil
	case "/user_manage/set_official_user":
		data := &dto.SetUserOfficial{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		var setType string
		if data.OperaType == 1 {
			setType = "设置客服"
		} else if data.OperaType == 2 {
			setType = "取消客服"
		}
		reqData := fmt.Sprintf("用户id:%d, 操作类型:%s", data.UserId, setType)
		return data, reqData, nil
	case "/user_manage/query_user_store":
		data := &dto.QryUser{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("用户id:%d", data.UserId)
		return data, reqData, nil
	case "/user_manage/query_user_file":
		data := &dto.QryUser{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("用户id:%d", data.UserId)
		if data.Limit == 0 {
			data.Limit = 20
		}
		reqData += fmt.Sprintf(", 请求页码:%d, 请求数量:%d", (data.Offset/data.Limit)+1, data.Limit)
		return data, reqData, nil
	case "/user_manage/query_user_login":
		data := &dto.QryUser{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("用户id:%d", data.UserId)
		if data.Limit == 0 {
			data.Limit = 20
		}
		reqData += fmt.Sprintf(", 请求页码:%d, 请求数量:%d", (data.Offset/data.Limit)+1, data.Limit)
		return data, reqData, nil
	case "/user_manage/update_user_username":
		data := &dto.UpdateUserName{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("用户id:%d, POPID:%s", data.UserId, data.UserName)
		return data, reqData, nil
	case "/user_manage/update_user_phone":
		data := &dto.UpdateUserPhone{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("用户id:%d, 手机号:%s", data.UserId, data.Phone)
		return data, reqData, nil
	case "/user_manage/query_user_opera":
		data := &dto.QryUser{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("用户id:%d", data.UserId)
		if data.Limit == 0 {
			data.Limit = 20
		}
		reqData += fmt.Sprintf(", 请求页码:%d, 请求数量:%d", (data.Offset/data.Limit)+1, data.Limit)
		return data, reqData, nil

	case "/record/query_message":
		data := &dto.MessageRecord{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := qryMsgInfo(data)
		return data, reqData, nil
	case "/record/query_dialogs":
		data := &dto.MessageRecord{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("用户id:%d", data.FromId)
		if data.PeerId != 0 {
			reqData += fmt.Sprintf("对方id:%d", data.FromId)
		}
		if data.Limit == 0 {
			data.Limit = 20
		}
		reqData += fmt.Sprintf(", 请求页码:%d, 请求数量:%d", (data.Offset/data.Limit)+1, data.Limit)
		return data, reqData, nil
	case "/record/del_file":
		data := &dto.RemoveFileMessage{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("用户id:%d, 对方类型:%d, 对方id:%d, 消息id:%v", data.UserId, data.PeerType, data.PeerId, data.MsgIds)
		return data, reqData, nil
	case "/record/del_peer_file":
		data := &dto.RemovePeerFile{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("对方类型:%d, 对方id:%d, 开始时间:%d, 结束:%d", data.PeerType, data.PeerId, data.Start, data.End)
		return data, reqData, nil

	case "/store_manage/all_store":
		data := &dto.QryUser{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		var reqData string
		if data.UserId != 0 {
			reqData = fmt.Sprintf("用户id:%d", data.UserId)
		}
		return data, reqData, nil
	case "/store_manage/user_store":
		data := &dto.StoreType{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		// 0:所有 1:图片 2:视频 3:音频 4:文件 5:其他
		var reqData string
		if data.Type == 0 {
			reqData = "文件类型:所有"
		} else if data.Type == 1 {
			reqData = "文件类型:图片"
		} else if data.Type == 2 {
			reqData = "文件类型:视频"
		} else if data.Type == 3 {
			reqData = "文件类型:音频"
		} else if data.Type == 4 {
			reqData = "文件类型:文件"
		} else if data.Type == 5 {
			reqData = "文件类型:其他"
		}
		if data.Limit == 0 {
			data.Limit = 20
		}
		reqData += fmt.Sprintf(", 请求页码:%d, 请求数量:%d", (data.Offset/data.Limit)+1, data.Limit)
		return data, reqData, nil

	case "/store_manage/last_upload":
		data := &dto.StoreType{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		var reqData string
		if data.UserId != 0 {
			reqData = fmt.Sprintf("用户id:%d, ", data.UserId)
		}
		if data.Type == 0 {
			reqData += "文件类型:所有"
		} else if data.Type == 1 {
			reqData += "文件类型:图片"
		} else if data.Type == 2 {
			reqData += "文件类型:视频"
		} else if data.Type == 3 {
			reqData += "文件类型:音频"
		} else if data.Type == 4 {
			reqData += "文件类型:文件"
		} else if data.Type == 5 {
			reqData += "文件类型:其他"
		}
		if data.Limit == 0 {
			data.Limit = 20
		}
		reqData += fmt.Sprintf(", 请求页码:%d, 请求数量:%d", (data.Offset/data.Limit)+1, data.Limit)
		return data, reqData, nil

	case "/system/query_all_permission":
		return nil, "", nil
	case "/system/add_role":
		data := &dto.Role{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("角色名称:%s, 角色权限id:%v", data.Name, data.PermissionIds)
		return data, reqData, nil
	case "/system/edit_role_permission":
		data := &dto.Role{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("角色id:%d, 角色权限id:%v", data.Id, data.PermissionIds)
		return data, reqData, nil
	case "/system/delete_role":
		data := &dto.Role{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("角色id:%d", data.Id)
		return data, reqData, nil
	case "/system/query_all_role":
		return nil, "", nil
	case "/system/add_account":
		data := &dto.Acount{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("账号名:%s, 姓名:%s, 角色id:%v", data.AccountName, data.UserName, data.RoleIds)
		if data.PwdUtil != 0 {
			reqData += fmt.Sprintf(", 密码有效期:%d天", data.PwdUtil)
		} else {
			reqData += ", 密码有效期:永久"
		}
		return data, reqData, nil

// ==============================================
	case "/system/edit_account":
		data := &dto.Acount{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("账号id:%d", data.Id)
		reqData += fmt.Sprintf(", 老权限id:%v, 新权限id:%v",  public.AcountRole[data.Id], data.RoleIds)
		if len(data.NewPwd) > 0 {
			reqData += ", 更新密码"
			if data.PwdUtil != 0 {
				reqData += fmt.Sprintf(", 密码有效期:%d天", data.PwdUtil)
			} else {
				reqData += ", 密码有效期:永久"
			}
		}
		return data, reqData, nil
	case "/system/edit_account_state":
		data := &dto.Acount{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("账号id:%d", data.Id)
		if data.ForbiddenType == 0 {
			reqData += ", 操作类型:禁用"
		} else if data.ForbiddenType == 1 {
			reqData += ", 操作类型:解禁"
		}
		return data, reqData, nil
	case "/system/edit_account_pwd":
		data := &dto.Acount{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := "用户修改密码"
		return data, reqData, nil
	case "/system/query_account":
		return nil, "", nil
	case "/system/query_all_account":
		return nil, "", nil
	case "/system/query_login_log":
		data := &dto.Acount{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("账号id:%d", data.Id)
		if data.Limit == 0 {
			data.Limit = 20
		}
		reqData += fmt.Sprintf(", 请求页码:%d, 请求数量:%d", (data.Offset/data.Limit)+1, data.Limit)
		return data, reqData, nil
	case "/system/del_account":
		data := &dto.Acount{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("账号id:%d", data.Id)
		return data, reqData, nil
	case "/system/add_white_list":
		data := &dto.AllowIp{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("ip地址:%s", data.Ip)
		return data, reqData, nil
	case "/system/del_white_list":
		data := &dto.AllowIp{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("白名单id:%d", data.Id)
		return data, reqData, nil
	case "/system/query_white_list":
		return nil, "", nil
	case "/system/edit_config":
		data := &dto.Config{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("key:%s, value:%s", data.Key, data.Value)
		return data, reqData, nil
	case "/system/get_config":
		data := &dto.Config{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("key:%s", data.Key)
		return data, reqData, nil
	case "/system/add_label":
		data := &dto.Lable{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("标签名:%s", data.Name)
		return data, reqData, nil
	case "/system/del_label":
		data := &dto.Lable{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("标签id:%d", data.Id)
		return data, reqData, nil
	case "/system/update_label":
		data := &dto.Lable{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		reqData := fmt.Sprintf("标签id:%d, 标签名:%s", data.Id, data.Name)
		return data, reqData, nil
	case "/system/get_label":
		return nil, "", nil
	case "/system/get_label_note":
		return nil, "", nil
	case "/system/get_request_record":
		data := &dto.Acount{}
		if err := c.ShouldBind(data); err != nil {
			return data, "", err
		}
		var reqData string
		if data.Id == 0 {
			reqData = "账户id:所有"
		} else {
			reqData = fmt.Sprintf("账户id:%d", data.Id)
		}
		if data.Limit == 0 {
			data.Limit = 20
		}
		reqData += fmt.Sprintf(", 请求页码:%d, 请求数量:%d", (data.Offset/data.Limit)+1, data.Limit)
		return data, reqData, nil
	case "":
	default:
		logger.LogSugar.Errorf("getUrlData url:%s", c.FullPath())
	}
	return nil, "", nil
}

func qryChatInfo(qt *dto.QryType) string {
	var reqData string
	if qt.ChatType == 1 {
		reqData = "请求群类型:群"
	} else if qt.ChatType == 2 {
		reqData = "请求群类型:频道"
	} else if qt.ChatType == 3 {
		reqData = "请求群类型:群发"
	} else {
		reqData = "请求群类型:所有"
	}
	switch qt.Type {
	case util.QryChatByName:
		reqData += fmt.Sprintf(", 请求类型:群名称, 请求参数:%s", qt.Qry)
	case util.QryChatByCountry:
		reqData += fmt.Sprintf(", 请求类型:国家, 请求参数:%s", qt.Qry)
	case util.QryChatByNum:
		reqData += fmt.Sprintf(", 请求类型:人数范围, 请求参数:%s", qt.Qry)
	case util.QryChatByCreator:
		reqData += fmt.Sprintf(", 请求类型:拥有者id, 请求参数:%s", qt.Qry)
	case util.QryChatByManage:
		reqData += fmt.Sprintf(", 请求类型:管理者id, 请求参数:%s", qt.Qry)
	case util.QryChatByNote:
		reqData += fmt.Sprintf(", 请求类型:标注关键字, 请求参数:%s", qt.Qry)
	case util.QryChatById:
		reqData += fmt.Sprintf(", 请求类型:群ID, 请求参数:%s", qt.Qry)
	case util.QryChatByStatus:
		var cStatus string
		if qt.Qry == "1" {
			cStatus = "正常"
		} else if qt.Qry == "2" {
			cStatus = "警告"
		} else if qt.Qry == "3" {
			cStatus = "短期禁言"
		} else if qt.Qry == "4" {
			cStatus = "长期禁言"
		}
		reqData += fmt.Sprintf(", 请求类型:群状态, 请求参数:%s", cStatus)
	case util.QryChatByUserId:
		reqData += fmt.Sprintf(", 请求类型:群成员, 请求参数:%s", qt.Qry)
	case util.QryChatByLabel:
		reqData += fmt.Sprintf(", 请求类型:群标签, 请求参数:%s", qt.Qry)
	case util.QryChatByActiveDate:
		nums := strings.Split(qt.Qry, ",")
		date := qt.Qry
		if len(nums) == 2 {
			s,_ := strconv.Atoi(nums[0])
			e,_ := strconv.Atoi(nums[1])
			start := time.Unix(int64(s), 0).Format("2006-01-02 15:04:05")
			end := time.Unix(int64(e), 0).Format("2006-01-02 15:04:05")
			date = start + "," + end
		}
		reqData += fmt.Sprintf(", 请求类型:活跃日期, 请求参数:%s", date)
	case util.QryChatByCreateDate:
		nums := strings.Split(qt.Qry, ",")
		date := qt.Qry
		if len(nums) == 2 {
			s,_ := strconv.Atoi(nums[0])
			e,_ := strconv.Atoi(nums[1])
			start := time.Unix(int64(s), 0).Format("2006-01-02 15:04:05")
			end := time.Unix(int64(e), 0).Format("2006-01-02 15:04:05")
			date = start + "," + end
		}
		reqData += fmt.Sprintf(", 请求类型:创建日期, 请求参数:%s", date)
	default:
		reqData += ", 请求类型:全部"
	}
	if qt.Limit == 0 {
		qt.Limit = 20
	}
	reqData += fmt.Sprintf(", 请求页码:%d, 请求数量:%d", (qt.Offset/qt.Limit)+1, qt.Limit)
	return reqData
}

func qryUserInfo(qt *dto.QryType) string {
	var reqData string
	switch qt.Type {
	case util.QryUserByName:
		reqData += fmt.Sprintf("请求类型:昵称, 请求参数:%s", qt.Qry)
	case util.QryUserByUserName:
		reqData += fmt.Sprintf("请求类型:popid, 请求参数:%s", qt.Qry)
	case util.QryUserByPhone:
		reqData += fmt.Sprintf("请求类型:手机号, 请求参数:%s", qt.Qry)
	case util.QryUserByCountry:
		reqData += fmt.Sprintf("请求类型:国家码, 请求参数:%s", qt.Qry)
	case util.QryUserByEmail:
		reqData += fmt.Sprintf("请求类型:邮箱, 请求参数:%s", qt.Qry)
	case util.QryUserByNote:
		reqData += fmt.Sprintf("请求类型:标注关键字, 请求参数:%s", qt.Qry)
	case util.QryUserById:
		reqData += fmt.Sprintf("请求类型:用户id, 请求参数:%s", qt.Qry)
	case util.QryUserByBanned:
		var status string
		if  qt.Qry == "1" {
			status = "已封禁"
		} else {
			status = "正常"
		}
		reqData += fmt.Sprintf("请求类型:封禁状态, 请求参数:%s", status)
	case util.QryUserByLabel:
		reqData += fmt.Sprintf("请求类型:用户标签, 请求参数:%s", qt.Qry)
	case util.QryUserByActiveDate:
		nums := strings.Split(qt.Qry, ",")
		date := qt.Qry
		if len(nums) == 2 {
			s,_ := strconv.Atoi(nums[0])
			e,_ := strconv.Atoi(nums[1])
			start := time.Unix(int64(s), 0).Format("2006-01-02 15:04:05")
			end := time.Unix(int64(e), 0).Format("2006-01-02 15:04:05")
			date = start + "," + end
		}
		reqData += fmt.Sprintf("请求类型:活跃日期, 请求参数:%s", date)
	case util.QryUserByCreateDate:
		nums := strings.Split(qt.Qry, ",")
		date := qt.Qry
		if len(nums) == 2 {
			s,_ := strconv.Atoi(nums[0])
			e,_ := strconv.Atoi(nums[1])
			start := time.Unix(int64(s), 0).Format("2006-01-02 15:04:05")
			end := time.Unix(int64(e), 0).Format("2006-01-02 15:04:05")
			date = start + "," + end
		}
		reqData += fmt.Sprintf("请求类型:注册日期, 请求参数:%s", date)
	case util.QryUserByDevice:
		// device 0:android 1:ios 2:mac 3:windows
		var device string
		if qt.Qry == "0" {
			device = "安卓"
		} else if qt.Qry == "1" {
			device = "IOS"
		} else if qt.Qry == "2" {
			device = "MAC"
		} else if qt.Qry == "3" {
			device = "Windows"
		}
		reqData += fmt.Sprintf("请求类型:设备类型, 请求参数:%s", device)
	case util.QryUserByOnline:
		var status string
		if qt.Qry == "1" {
			status = "在线"
		} else {
			status = "离线"
		}
		reqData += fmt.Sprintf("请求类型:在线状态, 请求参数:%s", status)
	case util.QryUserByOfficial:
		reqData += fmt.Sprintf("请求类型:客服")
	default:
		reqData += "请求类型:全部"
	}
	if qt.Limit == 0 {
		qt.Limit = 20
	}
	reqData += fmt.Sprintf(", 请求页码:%d, 请求数量:%d", (qt.Offset/qt.Limit)+1, qt.Limit)
	return reqData
}

func qryMsgInfo(data *dto.MessageRecord) string {
	var reqData string
	if data.PeerId > 100000 {
		reqData = "类型:单聊"
		if data.FromId != 0 {
			reqData += fmt.Sprintf(", 用户Aid:%d", data.FromId)
		}
		if data.PeerId != 0 {
			reqData += fmt.Sprintf(", 用户Bid:%d", data.PeerId)
		}
	} else {
		reqData = "类型:群聊"
		if data.PeerId != 0 {
			reqData += fmt.Sprintf(", 群id:%d", data.PeerId)
		}
	}
	if data.Start != 0 {
		start := time.Unix(int64(data.Start), 0).Format("2006-01-02 15:04:05")
		reqData += fmt.Sprintf(", 开始时间:%s", start)
	}
	if data.End != 0 {
		end := time.Unix(int64(data.End), 0).Format("2006-01-02 15:04:05")
		reqData += fmt.Sprintf(", 结束时间:%s", end)
	}

	//msgType := util.ToOriginMsgType(data.MessageType)
	switch data.MessageType {
	case util.ALL:
		reqData += ", 文件类型:全部"
	case util.MESSAGE:
		reqData += ", 文件类型:文本"
	case util.DOCUMENT:
		reqData += ", 文件类型:文件"
	case util.PHOTO:
		reqData += ", 文件类型:图片"
	case util.VIDEO:
		reqData += ", 文件类型:视频"
	case util.URL:
		reqData += ", 文件类型:链接"
	case util.GIF:
		reqData += ", 文件类型:GIF"
	case util.VOICE:
		reqData += ", 文件类型:语音"
	case util.MUSIC:
		reqData += ", 文件类型:音乐"
	case util.ROUND_VIDEO:
		reqData += ", 文件类型:圆形图片"
	case util.GEO:
		reqData += ", 文件类型:位置"
	case util.CONTACT:
		reqData += ", 文件类型:联系人"
	case util.OTHER:
		reqData += ", 文件类型:其他"
	}
	if data.Limit == 0 {
		data.Limit = 20
	}
	reqData += fmt.Sprintf(", 请求页码:%d, 请求数量:%d", (data.Offset/data.Limit)+1, data.Limit)
	return reqData
}