compile:
	protoc --gogo_out=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:sql  *.proto

test:
	go get ./protoc-gen-sql
	protoc --gogo_out=Msql.proto=github.com/travisjeffery/proto-go-sql:. --sql_out=Msql.proto=github.com/travisjeffery/proto-go-sql:. integration/proto/*.proto
	go test ./...
