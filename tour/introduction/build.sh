#!/usr/bin/env bash
# go > v1.16
go install fyne.io/fyne/v2/cmd/fyne@latest

# go < v1.15 
#go get fyne.io/fyne/v2/cmd/fyne

go build -O bin/introduction
fyne package -icon ../../icon.png

# result is a platform specific package
# for the current operating system.
