package minio_client

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"github.com/minio/minio-go"
	"github.com/minio/minio-go/pkg/encrypt"
	"io/ioutil"
	"net/http"
	"net/url"
	"pop-api/baselib/logger"
	"strings"
	"time"
)

var Initialized = false
var minioClientHttp *minio.Client
var minioClient *minio.Client
var encryption = ""
var BUCK_PRE = ""
var MinioIp = ""

func InitData(ip string) {
	MinioIp = ip
}

func InitClientConfig(endpoint, endpointHttp, sseKey, accessKeyID, secretAccessKey string) {
	if Initialized {
		return
	}
	encryption = sseKey

	var err error
	// Initialize minio client object.
	minioClient, err = minio.New(endpoint, accessKeyID, secretAccessKey, true)
	if endpoint == endpointHttp {
		logger.LogSugar.Infof("endpoint:%v, endpointHttp:%v, sseKey:%v, accessKeyID:%v, secretAccessKey:%v", endpoint, endpointHttp, sseKey, accessKeyID, secretAccessKey)
		minioClientHttp = minioClient
	} else {
		minioClientHttp, err = minio.New(endpointHttp, accessKeyID, secretAccessKey, false)
	}
	if err != nil {
		logger.LogSugar.Fatalf("minioClient.InitClientConfig:%v, endpoint:%v, endpointHttp:%v, sseKey:%v, accessKeyID:%v, secretAccessKey:%v", err, endpoint, endpointHttp, sseKey, accessKeyID, secretAccessKey)
		return
	}

	// 这里不要验证证书
	t := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	minioClient.SetCustomTransport(t)
	minioClientHttp.SetCustomTransport(t)

	Initialized = true
}

func PresignedGetObject(bucket, object string) string {
	reqParams := make(url.Values)
	reqParams.Set("response-content-type", "application/octet-stream")

	//getBucketPolicy
	//key := "Wink@YaMyB2GmOEetkib6O#+KRfuze6T"
	has := md5.Sum([]byte(encryption))
	md5Key := fmt.Sprintf("%x", has)
	baseKey := base64.StdEncoding.EncodeToString([]byte(encryption))
	fmt.Println("md5Key=", md5Key)
	fmt.Println("baseKey=", baseKey)

	url, err := minioClient.PresignedGetObject(bucket, object, time.Second*24*60*60, reqParams)
	if err != nil {
		logger.LogSugar.Errorf("PresignedGetObject err:%v", err)
		return ""
	}
	fmt.Println("url=", url)
	url2, _ := minioClient.PresignedHeadObject(bucket, object, time.Second*24*60*60, reqParams)
	fmt.Println("url2=", url2)

	return url.String()
}

func trimObjectName(objectName string) string {
	return strings.TrimLeft(objectName, "/")
}

func split(path string) (string, string) {
	/* 将第一个目录分离出来, 例如: x/y/z -> (x, y/z) */
	rows := strings.SplitN(path, "/", 2)
	if len(rows) == 2 {
		return rows[0], rows[1]
	}
	if len(rows) == 1 {
		return rows[0], ""
	}
	return "", ""
}

func splitbuck(path string) (string, string) {
	/* 将路径分为 (buckname, path) */
	path = trimObjectName(path)
	buck, path2 := split(path)
	return buck, path2
}

func GetObjectByLimit(path string, enc bool, offset int, limit int) ([]byte, error) {
	buckname, objectName := splitbuck(path)
	buckname = BUCK_PRE + buckname
	opt := minio.GetObjectOptions{}

	var reader *minio.Object
	var err error

	if enc {
		// New SSE-C where the cryptographic key is derived from a password and the objectname + bucketname as salt
		opt.ServerSideEncryption = encrypt.DefaultPBKDF([]byte(encryption), []byte(buckname+objectName))
		// Encrypt file content and upload to the server
		reader, err = minioClient.GetObject(buckname, objectName, opt)
	} else {
		reader, err = minioClientHttp.GetObject(buckname, objectName, opt)
	}

	if err != nil {
		return nil, err
	}
	defer reader.Close()

	if offset >= 0 && limit > 0 {
		data := make([]byte, limit, limit)
		if n, err := reader.ReadAt(data, int64(offset)); n > 0 {
			return data[:n], nil
		} else {
			return nil, err
		}
	} else {
		return ioutil.ReadAll(reader)
	}
}
