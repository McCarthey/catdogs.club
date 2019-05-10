.PHONY: build docker clean help

all: build docker clean

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o catdogs
	@echo "------ build go success ------"

docker:
	docker build -t catdogs .
	docker tag catdogs yokooll/catdogs
	docker push yokooll/catdogs
	@echo "------ docker push success ------"

clean:
	@docker images|grep none|awk '{print $3}'|xargs docker rmi -f
	@rm -rf catdogs
	@go clean -i .
	@echo "------ clean done ------"
	@docker images

help:
	@echo "make: compile packages and dependencies"
	@echo "make docker: build,tag,push docker image"
	@echo "make clean: remove object files and cached files"
