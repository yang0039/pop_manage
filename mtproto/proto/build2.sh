#!/bin/sh
SRC_DIR=.
DST_DIR=..


# 按照上面的说明，我想试试gogofaster，里面有描述without XXX_unrecognized。在第二部分使用官方库时，生成的这个XXX_属性，在初始化时很烦。
# (1)go get -u github.com/gogo/protobuf/proto 这一步特别慢，差点强行关掉了。
# (2)go get -u github.com/gogo/protobuf/protoc-gen-gogofaster
# (3)go get github.com/gogo/protobuf/gogoproto
# 然后就能看到bin下面多了一个protoc-gen-gogofaster.exe了。
# 使用protoc .\testCui.proto --gogofaster_out ./导出新的go文件，果然那一堆XXX不见了，然后之前的main文件可以直接运行。当然把import中的"github.com/golang/protobuf/proto"改成"github.com/gogo/protobuf/proto"也是可以的，使用下面的基准测试发现性能差不多啊，目前还不确定两者的区别。

# 作者：懒皮
# 链接：https://www.jianshu.com/p/7e3dcfbc5fd7
# 来源：简书
# 简书著作权归作者所有，任何形式的转载都请联系作者获得授权并注明出处。



# go get github.com/gogo/protobuf/proto
# go get github.com/gogo/protobuf/jsonpb
# go get github.com/gogo/protobuf/protoc-gen-gogo
# go get github.com/gogo/protobuf/gogoproto

# 作者：懒皮
# 链接：https://www.jianshu.com/p/7e3dcfbc5fd7
# 来源：简书
# 简书著作权归作者所有，任何形式的转载都请联系作者获得授权并注明出处。

#./codegen.sh
protoc -I=$SRC_DIR --gogofaster_out=plugins=grpc:$DST_DIR/ $SRC_DIR/*.proto
# protoc -I=$SRC_DIR --gogo_out=plugins=grpc:$DST_DIR/ $SRC_DIR/*.proto
# protoc -I=$SRC_DIR --gofast_out=plugins=grpc:$DST_DIR/ $SRC_DIR/*.proto


# //gogo
# protoc --gogo_out=. *.proto
# //gofast
# protoc --gofast_out=. *.proto

# echo "a:b" | sed 's/:/\'$'\n/g'
# echo "a;b" | sed $'s/;/\\\n/g'
# 修改

# 这个在.shk 执行才会生效，下面的在命令行中执行才会生效
sed -i "" $'s/math \"math\"/math \"math\"\\\n\\\t\"runtime\/debug\"\\\n\\\t\"winkim\/baselib\/logger\"/g' ../schema.tl.sync.pb.go
# 下面这个在控制参中才会生效
# sed -i "" $'s/math \"math\"/math \"math\"\\\n\\\t\"runtime\\\/debug\"\\\n\\\t\"winkim\\\/baselib\\\/logger\"/g' ../schema.tl.sync.pb.go

sed -i "" 's/func (m \*Chat) MarshalTo(dAtA \[\]byte) (int, error) {/func (m \*Chat) MarshalTo(dAtA \[\]byte) (int, error) {'$'\\\n\\\tdefer func() {'$'\\\n\\\t\\\tif err := recover(); err != nil {'$'\\\n\\\t\\\t\\\tlogger.LogSugar.Errorf(\"recover error MarshalTo panic: %v\\\\n%s\", err, string(debug.Stack()))'$'\\\n\\\t\\\t\\\tlogger.LogSugar.Errorf(\"recover error Chat MarshalTo: %v\\\\n\", m)'$'\\\n\\\t\\\t}\\\n\\\t}()/g' ../schema.tl.sync.pb.go
# func (m *Chat) MarshalTo(dAtA []byte) (int, error) {
# 	defer func() {
# 		if err := recover(); err != nil {
# 			logger.LogSugar.Errorf("recover error MarshalTo panic: %v\n%s", err, string(debug.Stack()))
# 			logger.LogSugar.Errorf("recover error Chat MarshalTo: %v\n", m)
# 		}
# 	}()
sed -i "" 's/func (m \*Message) MarshalTo(dAtA \[\]byte) (int, error) {/func (m \*Message) MarshalTo(dAtA \[\]byte) (int, error) {'$'\\\n\\\tdefer func() {'$'\\\n\\\t\\\tif err := recover(); err != nil {'$'\\\n\\\t\\\t\\\tlogger.LogSugar.Errorf(\"recover error MarshalTo panic: %v\\\\n%s\", err, string(debug.Stack()))'$'\\\n\\\t\\\t\\\tlogger.LogSugar.Errorf(\"recover error Message MarshalTo: %v\\\\n\", m)'$'\\\n\\\t\\\t}\\\n\\\t}()/g' ../schema.tl.sync.pb.go
# func (m *Message) MarshalTo(dAtA []byte) (int, error) {
# 	defer func() {
# 			if err := recover(); err != nil {
# 					logger.LogSugar.Errorf("recover error MarshalTo panic: %v\n%s", err, string(debug.Stack()))
# 					logger.LogSugar.Errorf("recover error Message MarshalTo: %v\n", m)
# 			}
# 	}(
sed -i "" 's/func (m \*Updates) Marshal() (dAtA \[\]byte, err error) {/func (m \*Updates) Marshal() (dAtA \[\]byte, err error) {'$'\\\n\\\tdefer func() {'$'\\\n\\\t\\\tif err := recover(); err != nil {'$'\\\n\\\t\\\t\\\tlogger.LogSugar.Errorf(\"recover error Marshal panic: %v\\\\n%s\", err, string(debug.Stack()))'$'\\\n\\\t\\\t\\\tlogger.LogSugar.Errorf(\"recover error Updates Marshal: %v\\\\n\", m)'$'\\\n\\\t\\\t}\\\n\\\t}()/g' ../schema.tl.sync.pb.go
# func (m *Updates) Marshal() (dAtA []byte, err error) {
# 	defer func() {
# 		if err := recover(); err != nil {
# 			logger.LogSugar.Errorf("recover error Marshal panic: %v\n%s", err, string(debug.Stack()))
# 			logger.LogSugar.Errorf("recover error Updates Marshal:%v", m)
# 		}
# 	}()

