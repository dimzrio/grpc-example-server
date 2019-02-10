#!/bin/bash

protoc model/nginx.proto --go_out=plugins=grpc:.