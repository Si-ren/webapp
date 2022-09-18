1.create protobuf gofile
protoc -I . --go_out ./v1  --go_opt paths=source_relative ./auth.proto

protoc -I . --grpc-gateway_out ./v1 
   --grpc-gateway_opt logtostderr=true \
   --grpc-gateway_opt paths=source_relative \
   --grpc-gateway_opt grpc_api_configuration=./auth.yaml \
   --grpc-gateway_opt standalone=false \
   ./auth.proto

protoc -I . --grpc-gateway_out ./v1    --grpc-gateway_opt logtostderr=true    --grpc-gateway_opt paths=source_relative    --grpc-gateway_opt grpc_api_configuration=./auth.yaml    --grpc-gateway_opt standalone=false    ./auth.proto


protoc -I . --go-grpc_out=require_unimplemented_servers=false:./v1 --go-grpc_opt  paths=source_relative ./auth.proto