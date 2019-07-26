.PHONY: build docker clean help doc ck genpb

all: doc genpb build docker clean

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o catdogs
	@echo "------ build go success ------"

docker:
	docker build -t catdogs .
	docker tag catdogs yokooll/catdogs
	docker push yokooll/catdogs
	@echo "------ docker push success ------"

genpb:
	protoc -I=$$GOPATH/src/catdogs-proto --micro_out=./pb --go_out=./pb $$GOPATH/src/catdogs-proto/*proto
	
doc:
	@swag init

clean:
	@rm -rf catdogs
	@go clean -i .
	@echo "------ clean done ------"
	@docker images

ck:
	@docker images|grep none|awk '{print $3}'|xargs docker rmi -f

help:
	@echo "make: compile packages and dependencies"
	@echo "make docker: build,tag,push docker image"
	@echo "make clean: remove object files and cached files"
	@echo "make doc: swag init"
	@echo "make ck: clean docker images"
	@echo "make genpb: generate pb files"
