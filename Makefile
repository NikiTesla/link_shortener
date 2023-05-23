code-gen:
	@ sudo apt install protobuf-compiler
	@ sudo apt install protoc-gen-go
	@ sudo apt install protoc-gen-go-grpc
	@ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	@ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	@ export PATH="$PATH:$(go env GOPATH)/bin"
	@ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/shortener.proto

run-server:
	@ echo "  >  running server"
	@ go run cmd/server/main.go

run-client:
	@ echo "  >  running client"
	@ go run cmd/client/main.go