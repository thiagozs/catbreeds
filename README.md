# Hostgator challenge

## Stack of tech

Here is a list of stack of tech used in this project

* Golang version 1.14.1 linux/amd64

## Dependencies of project

You need install some libs for running this project.

### Swagger API doc

`go get -u github.com/swaggo/swag/cmd/swag`

## Test project

Just run `go test ./... -v`

## Start project

just run `go run main.go` or `Makefile build`

## Generate docker image

Run `docker build -t thiagozs/challenge .`