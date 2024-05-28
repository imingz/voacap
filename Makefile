## 指定应用使用的 version 包，会通过 `-ldflags -X` 向该包中指定的变量注入值
VERSION_PACKAGE=voacap/pkg/version

## 定义 VERSION 语义化版本号
ifeq ($(origin VERSION), undefined)
VERSION := $(shell git describe --tags --always --match='v*')
endif

## 检查代码仓库是否是 dirty（默认dirty）
GIT_TREE_STATE:="dirty"
ifeq (, $(shell git status --porcelain 2>/dev/null))
	GIT_TREE_STATE="clean"
endif
GIT_COMMIT:=$(shell git rev-parse HEAD)

GO_LDFLAGS += \
	-X $(VERSION_PACKAGE).GitVersion=$(VERSION) \
	-X $(VERSION_PACKAGE).GitCommit=$(GIT_COMMIT) \
	-X $(VERSION_PACKAGE).GitTreeState=$(GIT_TREE_STATE) \
	-X $(VERSION_PACKAGE).BuildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

.PHONY: run
run: build # 运行服务.
	@./output/linux/voacap_go

.PHONY: build
build: tidy # 编译服务.
	@go build -v -ldflags "$(GO_LDFLAGS)" -o output/linux/voacap_go ./cmd/voacap.go

.PHONY: build-windows
build-windows: tidy # 编译 windows 版本.
	@GOOS=windows go build -v -ldflags "$(GO_LDFLAGS)" -o output/windows/voacap_go.exe ./cmd/voacap.go

.PHONY: tidy
tidy: # 自动添加/移除依赖包.
	@go mod tidy

.PHONY: hz_update_antennas
hz_update_antennas: # 生成 antennas 代码.
	@hz update -idl idl/antennas.proto

.PHONY: hz_update_station
hz_update_station: # 生成 antennas 代码.
	@hz update -idl idl/station.proto

.PHONY: hz_update_link
hz_update_link: # 生成 link 代码.
	@hz update -idl idl/link.proto