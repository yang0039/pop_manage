package member_manage

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"pop-api/baselib/minio_client"
	"pop-api/baselib/util"
	"pop-api/dal/dao"
	"pop-api/dto"
	"pop-api/middleware"
)

func (service *UserController) UserFile(c *gin.Context) {
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
	if user.UserId == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", user)))
		return
	}
	if user.Limit == 0 {
		user.Limit = 20
	}
	filesDao := dao.GetFilesDAO()
	files, count := filesDao.GetUserFiles(user.UserId, user.Offset, user.Limit)
	res := make([]map[string]interface{}, 0, len(files))
	for _, f := range files {
		file := map[string]interface{}{
			"id":          f.FileId,
			"type":        util.FileType(f.Ext),
			"size":        util.Folat4(float64(f.FileSize) / 1024 / 1024),
			"file_name":   f.UploadName,
			"upload_time": f.CreatedAt,
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

func (service *UserController) UserStore(c *gin.Context) {
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
	if user.UserId == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", user)))
		return
	}
	filesDao := dao.GetFilesDAO()
	files := filesDao.GetUserAllFile(user.UserId)

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
		{"type": "video", "video": util.Folat4(fileStore["video"]), "count": fileCount["video"]},
		{"type": "file", "size": util.Folat4(fileStore["file"]), "count": fileCount["file"]},
		{"type": "other", "size": util.Folat4(fileStore["other"]), "count": fileCount["other"]},
	}

	middleware.ResponseSuccess(c, data)
}
