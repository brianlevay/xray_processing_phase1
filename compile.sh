#!/bin/bash

echo "BASH: Started building version $1 at: $(date)"
go install fileExplorer
go install app

export GOOS=linux
go build -o xrayImgProcessing_linux_$1 app

export GOOS=windows
go build -o xrayImgProcessing_windows_$1.exe app

echo "BASH: Finished build version $1 at: $(date)"