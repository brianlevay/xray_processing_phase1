#!/bin/bash

echo "BASH: Started building at: $(date)"
go install fileExplorer
go install app

export GOARCH=amd64

export GOOS=linux
go build -o xrayImgProcessing_linux app

export GOOS=windows
go build -o xrayImgProcessing_windows.exe app

export GOOS=darwin
go build -o xrayImgProcessing_mac app

echo "BASH: Finished build at: $(date)"