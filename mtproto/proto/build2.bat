set currdir=%cd%

protoc -I. --go_out=plugins=grpc:.. api.proto
protoc -I. --go_out=plugins=grpc:.. api_service.proto
protoc -I. --go_out=plugins=grpc:.. cfs_service.proto
protoc -I. --go_out=plugins=grpc:.. rpc_error_codes.proto
protoc -I. --go_out=plugins=grpc:.. schema.tl.core_types.proto
protoc -I. --go_out=plugins=grpc:.. schema.tl.crc32.proto
protoc -I. --go_out=plugins=grpc:.. schema.tl.handshake.proto
protoc -I. --go_out=plugins=grpc:.. schema.tl.sync.proto
protoc -I. --go_out=plugins=grpc:.. schema.tl.sync_service.proto
protoc -I. --go_out=plugins=grpc:.. schema.tl.transport.proto
protoc -I. --go_out=plugins=grpc:.. schema.tl.transport_service.proto
protoc -I. --go_out=plugins=grpc:.. zproto_auth_key.proto
protoc -I. --go_out=plugins=grpc:.. zproto_sync.proto