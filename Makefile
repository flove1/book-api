include .env

.PHONY: cover
cover:
	go test -v ./... -coverprofile=./cover.html $@ && go tool cover -html=./cover.html