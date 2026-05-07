LOCAL_BIN := C:/Users/836D~1/GOLang/mic_curce/week_2/config/bin

LOCAL_MIGRATION_DIR=$(MIGRATiON_DIR)
LOCAL_MIGRATION_DSN="host=localhost port=$(PG_PORT) dbname=$(PG_DATABASE_NAME) user=$(PG_USER) password=$(PG_PASSWORD) sslmode=disable"


install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/protobuf/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-note-api

generate-note-api:
	@if not exist pkg\note_v1 mkdir pkg\note_v1
	protoc --proto_path api/note_v1 \
		--go_out=pkg/note_v1 --go_opt=paths=source_relative \
		--plugin=protoc-gen-go="$(LOCAL_BIN)/protoc-gen-go.exe" \
		--go-grpc_out=pkg/note_v1 --go-grpc_opt=paths=source_relative \
		--plugin=protoc-gen-go-grpc="$(LOCAL_BIN)/protoc-gen-go-grpc.exe" \
		api/note_v1/note.proto

local-migration-status:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	b$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v