#!/bin/sh

outname=./eword2decoded2047.wasm
mainpat=./cmd/eword2decoded2047/main.go

GOOS=wasip1 GOARCH=wasm go \
	build \
	-o "${outname}" \
	-ldflags="-s -w" \
	"${mainpat}"
