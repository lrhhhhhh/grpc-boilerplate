## grpc example
包含常见功能与特性：
- 通信
  - unary
  - client stream
  - server stream
  - bi-directional stream
- 拦截器 interceptor
  - 客户端拦截器
    - unary
    - stream
  - 服务端拦截器
    - unary
    - stream
- 使用tls进行安全通信
- 使用etcd实现服务发现与客户端负载均衡
- grpc-gateway和swagger
  - unary
  - stream 

## 运行
```shell
make up         # 新开终端，启动 etcd
make server     # 新开终端，启动 grpc server
make gateway    # 新开终端，启动 grpc gateway
make test       # 新开终端，用client和grpc server通信
make test_unary_rpc_rest        # 新开终端，使用http post的方式访问grpc gateway测试unary rpc
make test_server_stream_rest    # 使用http post的方式访问grpc gateway测试server stream rpc 
```

