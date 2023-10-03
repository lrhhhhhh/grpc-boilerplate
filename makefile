.PHONY:
install1:
	# download protoc
	curl -L https://github.com/protocolbuffers/protobuf/releases/download/v24.3/protoc-24.3-linux-x86_64.zip -o protoc-24.3-linux.zip  && \
    unzip -d protoc protoc-24.3-linux.zip

.PHONY:
install2:
	go mod download
	go install -v google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install -v google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest



PROTOC := $(shell pwd)/protoc/bin/protoc


.PHONY:
gen:
	echo protoc=$(PROTOC)
	$(PROTOC) --go_out=. --go-grpc_out=. proto/hello.proto






