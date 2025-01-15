# common说明
此目录主要存放微服务中的公共模块

- 创建 api
```
goctl api -o xxx.api # 创建api文件
goctl api go -api xxx.api -dir . # 生成api代码
```
- 创建rpc
```
goctl rpc -o xxx.proto  # 创建rpc文件
goctl rpc protoc xxx.proto --go_out=. --go-grpc_out=. --zrpc_out=.
```
## 注意：api和rpc的名称记得要区分开