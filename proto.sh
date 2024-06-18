protoc --go_out=./services/tasks --go_opt=paths=source_relative \
    --go-grpc_out=./services/tasks --go-grpc_opt=paths=source_relative \
    proto/tasks.proto

protoc --go_out=./services/main --go_opt=paths=source_relative \
    --go-grpc_out=./services/main --go-grpc_opt=paths=source_relative \
    proto/tasks.proto