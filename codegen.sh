#!/bin/bash

rm -rf gen/

buf generate

go mod tidy
