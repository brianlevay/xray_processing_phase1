#!/bin/bash

echo "BASH: Started build at: $(date)"
go install fileExplorer
go install histogram
go install app

export GOOS=linux
go build -o xray-processing-linux app

echo "BASH: Started running program at: $(date)"
./xray-processing-linux