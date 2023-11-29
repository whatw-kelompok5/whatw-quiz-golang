protos:
	protoc --go_out=. \
    --go-grpc_out=.\
    proto/avatar_proto.proto

protos question:
	protoc --go_out=. \
    --go-grpc_out=.\
    proto/question_proto.proto

protos gateway avatar:
	protoc -I=proto --grpc-gateway_out=. proto/avatar_proto.proto

protos gateway question:
	protoc -I=proto --grpc-gateway_out=. proto/question_proto.proto