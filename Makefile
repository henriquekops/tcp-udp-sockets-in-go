# Author: Henrique Kops

.PHONY: setup

#
# Golang environment and dependencies
#

setup:
	rm -rf .env
	cp .env.sample .env
	go get github.com/joho/godotenv
	go build main.go