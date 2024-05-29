# 指定应用使用的 version 包，会通过 `-ldflags -X` 向该包中指定的变量注入值
VERSION_PACKAGE=voacap/cmd

# 定义 VERSION 语义化版本号
ifeq ($(origin VERSION), undefined)
VERSION := $(shell git describe --tags --always --match='v*')
endif

GO_LDFLAGS += -X $(VERSION_PACKAGE).GitVersion=$(VERSION)

.PHONY: run
run: build # 运行服务.
	@./output/linux/voacap_go

.PHONY: build
build: tidy # 编译服务.
	go build -v -ldflags "$(GO_LDFLAGS)" -o output/linux/voacap_go main.go

.PHONY: build-windows
build-windows: tidy # 编译 windows 版本.
	GOOS=windows go build -v -ldflags "$(GO_LDFLAGS)" -o output/windows/voacap_go.exe main.go

.PHONY: tidy
tidy: # 自动添加/移除依赖包.
	@go mod tidy

.PHONY: hz_update
hz_update: hz_update_antennas hz_update_station hz_update_link # 生成所有代码.

.PHONY: hz_update_antennas
hz_update_antennas: # 生成 antennas 代码.
	hz update -idl idl/antennas.proto

.PHONY: hz_update_station
hz_update_station: # 生成 antennas 代码.
	hz update -idl idl/station.proto

.PHONY: hz_update_link
hz_update_link: # 生成 link 代码.
	hz update -idl idl/link.proto