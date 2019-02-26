SHELL := /bin/bash

image:
	GOOS=linux GOARCH=amd64 go build -o ./docker/http-test .
	cd docker && docker build -t reg.qiniu.com/hao/http-test:latest .
	docker push reg.qiniu.com/hao/http-test:latest
	rm ./docker/http-test
