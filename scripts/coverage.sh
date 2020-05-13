#!/bin/bash

file=cover.out

if [ -f $file ]; then
    rm $ffile
fi

go test ./... -coverprofile $file
#go tool cover -html $file
rm $file
