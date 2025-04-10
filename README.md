基于项目[lumina37/aiotieba](https://github.com/lumina37/aiotieba)
### 生成protobuf
```shell
(cd protobuf && \
protoc \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    *.proto)
```
