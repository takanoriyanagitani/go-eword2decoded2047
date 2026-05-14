#!/bin/sh

wsm="./eword2decoded2047.wasm"

runwasi() {
	cat /dev/stdin |
		wasmtime \
			run \
			"${wsm}"
}

printf \
	'%s\n' \
	'=?utf-8?q?hello,=20world?=' \
	'=?utf-8?q?hello, =e4=b8=96=e7=95=8c?=' \
	'=?utf-8?b?aGVsbG8s5LiW55WM?='  |
	runwasi
