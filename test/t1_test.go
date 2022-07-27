package test

import (
	"fmt"
	"pop-api/baselib/logger"
	"pop-api/baselib/minio_client"
	"testing"
)



func init() {
	minio_client.InitData("192.168.31.205")
	minio_client.InitClientConfig(minio_client.MinioIp + ":9123", minio_client.MinioIp + ":9000", "Wink@YaMyB2GmOEetkib6O#+KRfuze6T", "DQMYMM5HIJ4EF2XROGRK", "UooDmD1HwHvv67fjuVHYFpQcMGmyUCjyJt+B+n24")
}

func TestGetObject(t *testing.T) {
	filePath := "/photo/y/20220620/1538807250355900416.jpg"

	bytes,err := minio_client.GetObjectByLimit(filePath, false, 0, 1000)
	if err != nil {
		logger.LogSugar.Errorf("GetFile err:%v", err)
	}
	fmt.Println("bytes=", bytes)
}

func TestRemoveObject(t *testing.T) {
	filePath := "/photo/y/20220620/1538807250355900416.jpg"
	err := minio_client.RemoveObject(filePath)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("success")
	}
}





