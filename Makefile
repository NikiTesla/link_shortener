PROJECTNAME=$(shell basename "$(PWD)")

# code gen is for protobuf and protoc setting
code-gen:
	@ sudo apt install protobuf-compiler
	@ sudo apt install protoc-gen-go
	@ sudo apt install protoc-gen-go-grpc
	@ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	@ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
    # packages for rest gateway
	@ go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@ go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

	@ export PATH="$PATH:$(go env GOPATH)/bin"
	@ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/shortener.proto

run-server:
	@ echo "  >  running server"
	@ sudo docker start link_shortener_pg
	@ sleep 0.1
	@ go run cmd/main.go

docker:
	@ sudo docker rmi -f $(PROJECTNAME)
	@ echo "  >  making docker container $(PROJECTNAME)..."
	@ sudo docker build -t $(PROJECTNAME) .
	@ sudo docker compose up -d
	@ make migration-up

# making only init migration (may be changed using ARGS="[version]")
migration-up:
	@ echo "  >  making migrations"
	@ sudo docker start link_shortener_pg
	@ sleep 0.1
	@ cat schemas/0001_init.up.sql | sudo docker exec -i link_shortener_pg  psql -U postgres -d postgres

# making only init migration (may be changed using ARGS="[version]")
migration-down:
	@ echo "  >  making migrations"
	@ sudo docker start link_shortener_pg
	@ sleep 0.1
	@ cat schemas/0001_init.down.sql | sudo docker exec -i link_shortener_pg  psql -U postgres -d postgres