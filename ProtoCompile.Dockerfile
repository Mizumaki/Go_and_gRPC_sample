FROM grpc/go

WORKDIR /

CMD protoc \
    --proto_path="./grpc/proto" \
    --go_out=plugins=grpc:"./grpc/proto" \
    ./grpc/proto/**/*.proto
