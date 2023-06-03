#!/usr/bin/env bash
set -eux
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" assets
cd cmd/wasm
if [[ $# -ge 1 && "$1"="-w" ]]; then
	if [[ ! $(command -v inotifywait) ]]; then
		echo "inotifywait is not installed"
		exit 1
	fi
	inotifywait -mq -e close_write main.go | while read events; do
		echo "$(date) recompiling"
		GOOS=js GOARCH=wasm go build -o ../../assets/app.wasm || true
	done
else
	GOOS=js GOARCH=wasm go build -o ../../assets/app.wasm
fi
cd ../..
