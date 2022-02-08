package cache

import (
	"encoding/json"
	"errors"
	"github.com/gomodule/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"pop-api/baselib/logger"
	"pop-api/baselib/redis_client"
)

const (
	Minite = 60
	Hour   = 60 * Minite
	Day    = 24 * Hour
	Month  = 30 * Day
)

var ErrNil = errors.New("cache: nil returned")
var ErrNilEncoder = errors.New("cache: nil encoder")
var ErrNilDecoder = errors.New("cache: nil decoder")

func Warn(err error) {
	if err == nil {
		return
	}
	if err == ErrNil {
		return
	}
	if err == redis.ErrNil {
		return
	}
	logger.Logger.Warn(err.Error())
}

var Encoders map[string]func(v interface{}) ([]byte, error)
var Decoders map[string]func(data []byte, v interface{}) error

func init() {
	Encoders = make(map[string]func(v interface{}) ([]byte, error))
	Decoders = make(map[string]func(data []byte, v interface{}) error)
	Encoders["json"] = json.Marshal
	Decoders["json"] = json.Unmarshal
	Encoders["msgpack"] = msgpack.Marshal
	Decoders["msgpack"] = msgpack.Unmarshal
}

func Encode(encoding string, v interface{}) (interface{}, error) {
	if encoding == "" {
		return v, nil
	}
	encoder, ok := Encoders[encoding]
	if !ok {
		return nil, ErrNilEncoder
	}
	return encoder(v)
}

func Decode(encoding string, b []byte, v interface{}) error {
	if encoding == "" {
		return nil
	}
	decoder, ok := Decoders[encoding]
	if !ok {
		return ErrNilDecoder
	}
	return decoder(b, v)
}

type Result struct {
	Reply    interface{}
	Err      error
	Encoding string
}

func (this *Result) Int() int {
	v, err := redis.Int(this.Reply, this.Err)
	Warn(err)
	return v
}

func (this *Result) Int32() int32 {
	v, err := redis.Int(this.Reply, this.Err)
	Warn(err)
	return int32(v)
}

func (this *Result) Int64() int64 {
	v, err := redis.Int64(this.Reply, this.Err)
	Warn(err)
	return v
}

func (this *Result) Float64() float64 {
	v, err := redis.Float64(this.Reply, this.Err)
	Warn(err)
	return v
}

func (this *Result) String() string {
	v, err := redis.String(this.Reply, this.Err)
	Warn(err)
	return v
}

func (this *Result) Bytes() []byte {
	v, err := redis.Bytes(this.Reply, this.Err)
	Warn(err)
	return v
}

func (this *Result) Unmarshal(v interface{}) error {
	b := this.Bytes()
	if len(b) <= 0 {
		return ErrNil
	}
	return Decode(this.Encoding, b, v)
}

func (this *Result) Bool() bool {
	v, err := redis.Bool(this.Reply, this.Err)
	Warn(err)
	return v
}

func (this *Result) Values() []interface{} {
	v, err := redis.Values(this.Reply, this.Err)
	Warn(err)
	return v
}

func (this *Result) Ints() []int {
	v, err := redis.Ints(this.Reply, this.Err)
	Warn(err)
	return v
}

func (this *Result) Int32s() []int32 {
	v, err := redis.Ints(this.Reply, this.Err)
	Warn(err)
	var out []int32
	for _, i := range v {
		out = append(out, int32(i))
	}
	return out
}

func (this *Result) Int64s() []int64 {
	v, err := redis.Int64s(this.Reply, this.Err)
	Warn(err)
	return v
}

func (this *Result) Float64s() []float64 {
	v, err := redis.Float64s(this.Reply, this.Err)
	Warn(err)
	return v
}

func (this *Result) Strings() []string {
	v, err := redis.Strings(this.Reply, this.Err)
	Warn(err)
	return v
}

func (this *Result) StringMap() map[string]string {
	v, err := redis.StringMap(this.Reply, this.Err)
	Warn(err)
	return v
}

func (this *Result) IntMap() map[string]int {
	v, err := redis.IntMap(this.Reply, this.Err)
	Warn(err)
	return v
}

func (this *Result) Int64Map() map[string]int64 {
	v, err := redis.Int64Map(this.Reply, this.Err)
	Warn(err)
	return v
}

func Do(dbname, command string, args ...interface{}) *Result {
	//st := time.Now()
	conn := redis_client.GetRedisPoolClient(dbname)
	defer conn.Close()
	reply, err := conn.Do(command, args...)
	if err != nil {
		logger.LogSugar.Errorf("[%s]%v, args: %d, %v", command, args, len(args), err)
	}
	return &Result{Reply: reply, Err: err}
}

type Key struct {
	DBName   string      // redis连接池名
	Key      interface{} // 主键名
	Field    interface{} // (可选)二级键名, 哈希操作使用, 比如hget <key> <field>
	Encoding string      // 序列化方式, msgpack/json/自定义
	Timeout  int         // 过期时间(秒)
}

func (this *Key) Do(command string, args ...interface{}) *Result {
	var newargs []interface{}
	if this.Field != nil {
		newargs = make([]interface{}, 0, len(args)+2)
		newargs = append(newargs, this.Key, this.Field)
		newargs = append(newargs, args...)
	} else {
		newargs = make([]interface{}, 0, len(args)+1)
		newargs = append(newargs, this.Key)
		newargs = append(newargs, args...)
	}
	result := Do(this.DBName, command, newargs...)
	result.Encoding = this.Encoding
	return result
}

func (this *Key) Exists() bool {
	return this.Do("EXISTS").Bool()
}

func (this *Key) Get() *Result {
	return this.Do("GET")
}

func (this *Key) Set(value interface{}) *Result {
	value, err := Encode(this.Encoding, value)
	if err != nil {
		return &Result{Reply: nil, Err: err}
	}
	if this.Timeout > 0 {
		return this.Do("SET", value, "EX", this.Timeout)
	}
	return this.Do("SET", value)
}

func (this *Key) Del() *Result {
	return this.Do("DEL")
}

func (this *Key) Incr() *Result {
	return this.Do("INCR")
}

func (this *Key) HExists() bool {
	return this.Do("HEXISTS").Bool()
}

func (this *Key) HGet(fields ...interface{}) *Result {
	return this.Do("HGET", fields...)
}

func (this *Key) HSet(value interface{}) *Result {
	value, err := Encode(this.Encoding, value)
	if err != nil {
		return &Result{Reply: nil, Err: err}
	}
	return this.Do("HSET", value)
}

func (this *Key) HDel(fields ...interface{}) *Result {
	return this.Do("HDEL", fields...)
}

//set
func (this *Key) SAdd(fields ...interface{}) *Result {
	return this.Do("sadd", fields...)
}

func (this *Key) SRem(fields ...interface{}) *Result {
	return this.Do("srem", fields...)
}

func (this *Key) SIsmember(fields ...interface{}) bool {
	logger.LogSugar.Errorf("this:%v, fields:%v", *this, fields)

	return this.Do("sismember", fields...).Bool()
}
