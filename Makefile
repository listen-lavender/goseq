# 版权 © 2020 云族佳科技有限公司版权所有
.PHONY: default
default:

# 定义一个制作镜像的函数
define docker_build_image
	@# 第一个参数是程序名称
	@# 第二个参数是镜像的tag
	@# 第三个参数Dockerfile文件路径
	@# 第四个参数Docker制作镜像的路径
	docker build ${BUILD_ARGS} -t clouderwork/${1}:${2} -f ${3} ${4}
endef

# 编译可执行程序 打包基础镜像
.PHONY: build-exe build-base-images
build-exe:
	@mkdir -p _output/
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cmd/goseq/goseq cmd/goseq/main.go

build-base-images: build-exe
	$(call docker_build_image,goseq,latest,./cmd/goseq/Dockerfile,./cmd/goseq)

