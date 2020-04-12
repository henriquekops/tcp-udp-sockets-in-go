# Author: Henrique Kops

.PHONY: setup-linux setup-windows

#
# Golang environment and dependencies
#

setup-linux:
	rm -f .env
	cp .env.sample .env
	go get github.com/joho/godotenv
	go build main.go

setup-windows:
	del .env
	copy .env.sample .env
	go get github.com/joho/godotenv
	go build main.go
