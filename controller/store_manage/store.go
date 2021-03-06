package store_manage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pop-api/baselib/logger"
	"pop-api/baselib/minio_client"
	"pop-api/baselib/util"
	"pop-api/dal/dao"
	"pop-api/dal/dataobject"
	"pop-api/dto"
	"pop-api/middleware"
	"strings"
	"time"
)

func (service *StoreController) AllStore(c *gin.Context) {
	//user := &dto.QryUser{}
	//if err := c.ShouldBind(user); err != nil {
	//	middleware.ResponseError(c, 500, "系统错误", err)
	//	return
	//}

	bindData, err := middleware.ShouldBind(c)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	user, _ := bindData.(*dto.QryUser)
	filesDao := dao.GetFilesDAO()
	files := filesDao.GetAllStore(user.UserId)

	fileStore := map[string]float64{
		"other": 0.00,
		"file":  0.00,
		"photo": 0.00,
		"audio": 0.00,
		"video": 0.00,
	}
	fileCount := map[string]int32{
		"other": 0,
		"file":  0,
		"photo": 0,
		"audio": 0,
		"video": 0,
	}

	for _, f := range files {
		fileType := util.FileType(f.Ext)
		fileStore[fileType] += f.Size
		fileCount[fileType] += f.Count
	}

	data := []map[string]interface{}{
		{"type": "photo", "size": util.Folat4(fileStore["photo"]), "count": fileCount["photo"]},
		{"type": "audio", "size": util.Folat4(fileStore["audio"]), "count": fileCount["audio"]},
		{"type": "video", "size": util.Folat4(fileStore["video"]), "count": fileCount["video"]},
		{"type": "file", "size": util.Folat4(fileStore["file"]), "count": fileCount["file"]},
		{"type": "other", "size": util.Folat4(fileStore["other"]), "count": fileCount["other"]},
	}

	middleware.ResponseSuccess(c, data)
}

func (service *StoreController) UserStore(c *gin.Context) {
	//sType := &dto.StoreType{}
	//if err := c.ShouldBind(sType); err != nil {
	//	middleware.ResponseError(c, 500, "系统错误", err)
	//	return
	//}

	bindData, err := middleware.ShouldBind(c)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	sType, _ := bindData.(*dto.StoreType)
	if sType.Limit == 0 {
		sType.Limit = 20
	}

	filesDao := dao.GetFilesDAO()
	userDao := dao.GetUserDAO()

	userStore, count := filesDao.GetUserRank(sType.Type, sType.Offset, sType.Limit)
	var userIds []int32
	for _, s := range userStore {
		userIds = append(userIds, s.UserId)
	}
	users := userDao.GetUsers(userIds)
	uMap := make(map[int32]*dataobject.UserDO, len(users))
	for _, u := range users {
		uMap[u.Id] = u
	}

	res := make([]map[string]interface{}, 0, len(userStore))
	for _, f := range userStore {
		var uName, fName, lName string
		u, ok := uMap[f.UserId]
		if ok {
			uName = u.Username
			fName = u.FirstName
			lName = u.LastName
		}

		file := map[string]interface{}{
			"user_id":    f.UserId,
			"username":   uName,
			"first_name": fName,
			"last_name":  lName,
			"size":       util.Folat4(f.Size),
			"count":      f.Count,
		}
		res = append(res, file)
	}

	data := map[string]interface{}{
		"files": res,
		"count": count,
	}
	middleware.ResponseSuccess(c, data)
}

func (service *StoreController) LastUpload(c *gin.Context) {
	//sType := &dto.StoreType{}
	//if err := c.ShouldBind(sType); err != nil {
	//	middleware.ResponseError(c, 500, "系统错误", err)
	//	return
	//}

	bindData, err := middleware.ShouldBind(c)
	if err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	sType, _ := bindData.(*dto.StoreType)
	if sType.Limit == 0 {
		sType.Limit = 20
	}

	filesDao := dao.GetFilesDAO()
	userDao := dao.GetUserDAO()

	files, count := filesDao.GetLastUpload(sType.UserId, sType.Type, sType.Offset, sType.Limit)
	var userIds []int32
	for _, f := range files {
		userIds = append(userIds, f.UserId)
	}

	users := userDao.GetUsers(userIds)
	uMap := make(map[int32]*dataobject.UserDO, len(users))
	for _, u := range users {
		uMap[u.Id] = u
	}

	res := make([]map[string]interface{}, 0, len(files))
	for _, f := range files {
		var timeInt int64
		loc, _ := time.LoadLocation("Local")
		if strings.Contains(f.CreatedAt,"+") {
			t, _ := time.ParseInLocation("2006-01-02T15:04:05+08:00", f.CreatedAt, loc)
			timeInt = t.Unix()
		} else {
			t, _ := time.ParseInLocation("2006-01-02 15:04:05", f.CreatedAt, loc)
			timeInt = t.Unix()
		}
		logger.LogSugar.Debugf("---------- CreatedAt=%s", f.CreatedAt)
		file := map[string]interface{}{
			"user_id":     f.UserId,
			"username":    uMap[f.UserId].Username,
			"first_name":  uMap[f.UserId].FirstName,
			"last_name":   uMap[f.UserId].LastName,
			"id":          f.FileId,
			"type":        util.FileType(f.Ext),
			"size":        util.Folat4(float64(f.FileSize) / 1024 / 1024),
			"file_name":   f.UploadName,
			//"upload_time": f.CreatedAt,
			"upload_time": timeInt,
			"url":         fmt.Sprintf("http://%s%s", minio_client.MinioIp, f.FilePath),
		}
		res = append(res, file)
	}

	data := map[string]interface{}{
		"files": res,
		"count": count,
	}
	middleware.ResponseSuccess(c, data)
}
