.PHONY: gen

gen:
	protoc --go_out=plugins=grpc:. api/*.proto