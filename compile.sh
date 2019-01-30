#!/bin/bash

echo "BASH: Started building at: $(date)"
go install fileExplorer
go install app

export GOOS=linux
go build -o xrayImgProcessing_linux app

export GOOS=windows
go build -o xrayImgProcessing_windows.exe app

echo "BASH: Finished build at: $(date)"