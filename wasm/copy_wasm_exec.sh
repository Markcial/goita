#!/bin/sh
# Copy the wasm_exec.js that matches this Go toolchain into ../storage.
# Go 1.24 moved the file from $GOROOT/misc/wasm/ to $GOROOT/lib/wasm/.
set -eu

root="${GOROOT:-}"
if [ -z "$root" ]; then
	root=$(go env GOROOT)
fi

dest="../storage/wasm_exec.js"
for rel in lib/wasm/wasm_exec.js misc/wasm/wasm_exec.js; do
	src="$root/$rel"
	if [ -f "$src" ]; then
		cp "$src" "$dest"
		exit 0
	fi
done

echo "wasm_exec.js not found under $root/lib/wasm or $root/misc/wasm" >&2
exit 1
