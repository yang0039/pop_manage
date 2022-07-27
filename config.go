// Copyright © 2014 Terry Mao, LiuDing All rights reserved.
// This file is part of gopush-cluster.

// gopush-cluster is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// gopush-cluster is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with gopush-cluster.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"pop-api/baselib/mysql_client"
	"pop-api/baselib/redis_client"
	"runtime"
)

var (
	Conf     *Config
	confFile string
)

func init() {
	flag.StringVar(&confFile, "c", "./api.toml", " set api config file path")
}

type RpcServerConfig struct {
	Addr string
}

type CoinHubConfig struct {
	Addr string `toml:"addr"`
}

type MinioConfig struct {
	Endpoint        string
	EndpointHttp    string
	SseKey          string
	AccessKeyID     string
	SecretAccessKey string
}

type Config struct {
	Ver       string
	PidFile   string
	LogFile   string
	PprofBind []string
	StatBind  []string

	ServerId int
	MaxProc  int

	FuncType int  // 作为功能使用 0:后台接口管理，其他作为数据修复

	Dbindex int
	Minio string
	Cfsdata string
	Mysql *mysql_client.MySQLConfig
	Redis *redis_client.RedisConfig

}

type MysqlConf struct {
	UserCount    int32
	ChannelCount int32
	Mysql        []mysql_client.MySQLConfig
}

func NewConfig() *Config {
	return &Config{
		Ver:     "0.0.1",
		PidFile: "./api.pid",
		LogFile: "./api-log.xml",

		PprofBind: []string{"localhost:8101"},
		StatBind:  []string{"localhost:8102"},
		MaxProc: runtime.NumCPU(),
	}
}

// InitConfig init the global config.
func InitConfig() (err error) {
	Conf = NewConfig()
	if _, err := toml.DecodeFile(confFile, Conf); err != nil {
		panic(fmt.Errorf("Decode Config Error: %s", err.Error()))
		//Sugar.Fatalf("conf init err:%v", err)
	}
	return nil
}

func ReloadConfig() (*Config, error) {
	conf := NewConfig()
	if _, err := toml.DecodeFile(confFile, conf); err != nil {
		return nil, err
	}
	return conf, nil
}
