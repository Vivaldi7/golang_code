LOCAL_BIN := C:/Users/836D~1/GOLang/mic_curce/grpc/bin

install-deps:
	go env -w GOBIN=$(LOCAL_BIN)
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go env -u GOBIN

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

