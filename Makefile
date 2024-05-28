.PHONY: run
run: build # 运行服务.
	@./output/linux/voacap_go

.PHONY: build
build: # 编译服务.
	@go build -o output/linux/voacap_go .

.PHONY: build-windows
build-windows: # 编译 windows 版本.
	@GOOS=windows go build -o output/windows/voacap_go.exe .

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