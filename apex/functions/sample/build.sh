#!/bin/bash

GOPATH=/opt/golang/src/github.com/jonee316/lambda_golang_mongodb_apex/vendor:/opt/golang GO15VENDOREXPERIMENT=0 go build github.com/jonee316/lambda_golang_mongodb_apex/apex/functions/sample/

