package member_manage

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"pop-api/baselib/logger"
	"pop-api/baselib/minio_client"
	"pop-api/dal/dao"
	"pop-api/dto"
	"pop-api/middleware"
)

func (service *UserController) GetFile(c *gin.Context) {
	params := &dto.QryFile{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c, 500, "系统错误", err)
		return
	}
	if params.FileId == 0 {
		middleware.ResponseError(c, 400, "参数错误", errors.New(fmt.Sprintf("invalid param, param:%v", params)))
		return
	}

	photoDao := dao.GetPhotoDAO()
	photo := photoDao.SelectByPhotoId(params.FileId, 0)

	bytes,err := minio_client.GetObjectByLimit(photo.FilePath, photo.Sse, 0, 1000)
	if err != nil {
		logger.LogSugar.Errorf("GetFile err:%v", err)
	}
	fmt.Println("bytes=", bytes)

	middleware.ResponseSuccess(c, "")
}
