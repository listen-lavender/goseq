# 版权 © 2020 云族佳科技有限公司版权所有

.PHONY: all prepare
all: prepare seq

ifeq (, $(GOPATH))
  $(error 请配置 GOPATH 环境变量)
endif
ifeq (, $(GOBIN))
  $(error 请配置 GOBIN 环境变量 并加入到系统PATH中)
endif
# 取$GOPATH的第一个路径
GO_PATH=$(shell echo ${GOPATH} | awk -F : '{print $$1}')
ifeq (, $(GO_PATH))
  GO_PATH = ${GOPATH}
endif

# 自动生成对应的go文件
.PHONY: seq
seq:
	@protoc \
            -I . \
            -I /usr/local/include \
            -I ${GO_PATH}/src \
            -I ${GO_PATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
            -I ${GO_PATH}/src/github.com/grpc-ecosystem/grpc-gateway \
            --go_out=plugins=grpc:./_output \
            $(sort $(wildcard ./${@}/*.proto))
	@mv -f _output/github.com/listen-lavender/goseq/api/pb${@} ../pb${@}

.PHONY: prepare
prepare: depend
	@rm -rf ./_output/*
	@mkdir -p ./_output
	@rm -rf ../pb*

.PHONY: depend protoc protoc-gen-go googleapis
depend: protoc protoc-gen-go googleapis
protoc:
ifeq (, $(shell which protoc))
	@echo "[请自行安装protoc](http://google.github.io/proto-lens/installing-protoc.html)"
	@exit 1
endif
protoc-gen-go:
ifeq (, $(shell which protoc-gen-go))
	cd ~ ; GO111MODULE=on go get -u -v github.com/golang/protobuf/protoc-gen-go@v1.4.2
endif
googleapis:
ifeq (,$(wildcard ${GO_PATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis))
	cd ~ ; go get -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
endif
