.PHONY: run
run: # 运行服务.
	@go run .

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