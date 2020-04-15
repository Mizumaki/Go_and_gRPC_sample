## Check gRPC works

At first, `docker-compose up`.

Then, `docker-compose exec go_and_grpc sh`
And, inside the container, run `go run ./src/scripts/user/user.go`.

## Generate Go code from `.proto` files

`docker build -f ProtoCompile.Dockerfile -t proto-compile .`

`docker run -v $PWD/server/src/handler/grpc:/grpc proto-compile`
