#!/bin/bash

go test -v *.go

#
# go test .
# main.test
#/tmp/go-build999206704/b001/_testmain.go:13:2: cannot import "main"
#/tmp/go-build999206704/b001/_testmain.go:21:14: undefined: _test
#FAIL    main [build failed]
#FAIL

