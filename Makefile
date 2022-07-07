run-db:
	docker-compose up -d

run-server:
	go run ./cmd/server

run-client:
	go run ./cmd/client/http

generate-proto:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        ./api/protobuf/calendar/event.proto
