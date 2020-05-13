# Hostgator hostgator-challenge

## Stack of tech

Here is a list of stack of tech used in this project

* Golang version 1.14.1 linux/amd64
* Docker CE (todo)

## Dependencies of project

You need install some libs for running this project.

### Swagger API doc

`go get -u github.com/swaggo/swag/cmd/swag`

for update the documentation just run command `cd api/cmd` and after run the swag for find and write a swagger documentation. `swag init -g main.go`

## Test project

Just run `go test ./... -v` (todo)

## Start project

just run `go run api/cmd/main.go` or `Makefile build` (todo)

## Generate docker image

Run `docker build . -t thiagozs/challenge` (todo)
