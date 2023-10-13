.PHONY: install1
install1:
	# download protoc
	curl -L https://github.com/protocolbuffers/protobuf/releases/download/v24.3/protoc-24.3-linux-x86_64.zip -o protoc-24.3-linux.zip  && \
    unzip -d protoc protoc-24.3-linux.zip

.PHONY: install2
install2:
	go mod download
	go install -v google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install -v google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


PROTOC := $(shell pwd)/protoc/bin/protoc

.PHONY: gen
gen:
	echo protoc=$(PROTOC)
	$(PROTOC) --proto_path=api/proto api/proto/*.proto --go_out=. --go-grpc_out=. --grpc-gateway_out=. --openapiv2_out=./api/swagger

.PHONY: tls
tls:
	mkdir -p tls
	cd tls && \
	openssl genrsa -out ca.key 2048 && \
	openssl req -new -key ca.key -out ca.csr && \
	openssl req -new -x509 -key ca.key -out ca.crt -days 36500 && \
	openssl genpkey -algorithm RSA -out server.key && \
	openssl req -new -nodes -key server.key -out server.csr -days 3650 -subj "/C=cn/OU=myorg/O=mycomp/CN=myname" -config openssl.cnf -extensions v3_req && \
	openssl x509 -req -days 3650 -in server.csr -out server.pem -CA ca.crt -CAkey ca.key -CAcreateserial -extfile openssl.cnf -extensions v3_req


.PHONY: up
up:
	cd deployments/docker-compose && \
	mkdir -p -m 0777 etcd/data && \
	docker-compose up

.PHONY: down
down:
	cd deployments/docker-compose && \
	docker-compose down

.PHONY: rm
rm:
	cd deployments/docker-compose && sudo rm -rf etcd

.PHONY: server
server:
	go run cmd/server/main.go

.PHONY: gateway
gateway:
	go run cmd/grpc-gateway/main.go

.PHONY: test
test:
	go run cmd/client/main.go

.PHONY: test_unary_rpc_rest  # 需要先启动etcd、grpc server 和 grpc gateway
test_unary_rpc_rest:
	curl -X POST -H "Content-Type: application/json" -d  '{"msg":{"fromUsername":"1","toUsername":"2","content":"3"}}' http://localhost:9091/v1/example/hello

.PHONY: test_server_stream_rest
test_server_stream_rest:
	curl -X POST -H "Content-Type: application/json" -d  '{"TimeStamp": 12314125}' http://localhost:9091/v1/example/hello-server-stream
