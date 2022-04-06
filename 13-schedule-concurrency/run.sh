#!/bin/bash

export GODEBUG=schedtrace=1000,scheddetail=1
go run main.go
