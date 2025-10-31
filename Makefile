PROTO_SRC_DIR=proto
PROTO_FILES:=$(shell find $(PROTO_SRC_DIR) -name '*.proto')

.PHONY: proto
proto:
	protoc -I $(PROTO_SRC_DIR) \
		--go_out=services/app --go_opt=module=go-starter-app \
		--go-grpc_out=services/app --go-grpc_opt=module=go-starter-app \
		$(PROTO_FILES)


