# ProtoBuf

```azure
cd ./apps/book
protoc --proto_path=. --proto_path=../../common/pb --go_out=. --go_opt=module="calf-restful-api/apps/book" --go-grpc_out=.  --go-grpc_opt=module="calf-restful-api/apps/book"  pb/book.proto
```