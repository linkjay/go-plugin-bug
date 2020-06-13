#!/bin/bash

# Compiles our plugins then builds

#compile plugin
go build -buildmode=plugin -o plugins/test-plugin.so plugins/test-plugin.go

#compile main
go build main.go