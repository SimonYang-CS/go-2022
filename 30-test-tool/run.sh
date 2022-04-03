#!/bin/bash

go test -v -coverprofile=c.out *.go

#
# go test .
# main.test
#/tmp/go-build999206704/b001/_testmain.go:13:2: cannot import "main"
#/tmp/go-build999206704/b001/_testmain.go:21:14: undefined: _test
#FAIL    main [build failed]
#FAIL

go tool cover -html=c.out

# run benchmark
# go test -v -bench=. -benchtime=10s *.go
