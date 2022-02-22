#!/bin/bash
set -e

rm -rf gen/

buf generate

go mod tidy
