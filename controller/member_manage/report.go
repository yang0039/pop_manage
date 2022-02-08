package member_manage

import (
	"github.com/gin-gonic/gin"
	"pop-api/controller/_dummy/recordapi"
	"pop-api/dal/dao"
	"pop-api/dto"
	"pop-api/middleware"
	"strconv"
	"strings"
)

func (service *UserController) GetReport(c *gin.Context) {
	params := &dto.QryType{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if params.Limit == 0 {
		params.Limit = 20
	}

	reportDao := dao.GetReportDAO()
	msgDao := dao.GetChannelMsgRowDAO()
	userMsgDao := dao.GetUserMsgRowDAO()
	reports := reportDao.GetReports(params.Limit, params.Offset)
	total := reportDao.GetReportCount()
	reportRes := make([]*dto.ReportRes, 0, len(reports))

	for _, report := range reports {

		r := &dto.ReportRes{
			ReportUser: recordapi.GetUser(report.UserId),
			Peer:       recordapi.GetPeer(report.PeerType, report.PeerId),
			Reason:     report.Reason,
			Content:    report.Content,
			ReportTime: report.AddTime,
		}

		if report.MsgIds == "" { // 举报的群
			r.Messages = make([]*dto.Message, 0)
			reportRes = append(reportRes, r)
			continue
		}
		msgStrs := strings.Split(report.MsgIds, ",")
		msgIds := make([]int32, 0, len(msgStrs))
		for _, idStr := range msgStrs {
			id, _ := strconv.Atoi(idStr)
			msgIds = append(msgIds, int32(id))
		}

		if report.PeerType == 4 { // 超级群和频道
			msgs := msgDao.GetChannelMsgRowsById(report.PeerId, msgIds)
			r.Messages = recordapi.ChannelMsgToMessage(msgs)
		} else {
			msgs := userMsgDao.GetUserMsgRowsById(report.UserId, msgIds)
			r.Messages = recordapi.UserMsgToMessage(msgs)
		}
		reportRes = append(reportRes, r)
	}

	data := map[string]interface{}{
		"report": reportRes,
		"count":  total,
	}
	middleware.ResponseSuccess(c, data)
}
