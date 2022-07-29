GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)

.PHONY: config
# generate internal proto
# protoc -I=$GOPATH/pkg/mod --proto_path=./internal --proto_path=./third_party --gogofaster_out=paths=source_relative:./internal conf/conf.proto
# protoc --proto_path=./internal --proto_path=./third_party --go_out=paths=source_relative:./internal conf/conf.proto
config:
	protoc --proto_path=./internal \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./internal \
	       $(INTERNAL_PROTO_FILES)




