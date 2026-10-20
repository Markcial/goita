<div align="center">

# 🪟 goita

**Minimal desktop application with webview, websockets and WASM, using only Go.**

[![License: MIT](https://img.shields.io/github/license/Markcial/goita)](LICENSE)
![Go 1.16+](https://img.shields.io/badge/Go-1.16%2B-00ADD8?logo=go&logoColor=white)
![WebAssembly](https://img.shields.io/badge/WebAssembly-654FF0?logo=webassembly&logoColor=white)
![webview](https://img.shields.io/badge/webview-desktop-555)
[![GitHub stars](https://img.shields.io/github/stars/Markcial/goita?style=social)](https://github.com/Markcial/goita/stargazers)

![goita demo](https://user-images.githubusercontent.com/208523/109386098-c1963000-78f8-11eb-84c2-7a204b1eb031.gif)

</div>

This is an experiment on how to create a minimal desktop application using only the Go language.

## ✨ How it works

- 🪟 **Native window** through [webview](https://github.com/webview/webview), pointed at a local server
- 📦 **Embedded WASM**: `main.go` embeds the WebAssembly build (`//go:embed`) and serves it on a random free port
- 🔌 **Websocket bridge**: the UI (Go compiled to WASM) and the app talk through JSON messages
- 🧱 **Typed protocol**: the `bridge` package defines the request/response types on both sides

```text
goita/
├── main.go      # desktop app: embeds the WASM, starts the server, opens the webview
├── bridge/      # shared JSON message types (Request / Response) and Process()
├── wasm/        # UI code compiled to WebAssembly (honnef.co/go/js/dom/v2)
└── storage/     # generated main.wasm + wasm_exec.js land here
```

A message from the UI and its answer:

```json
{ "kind": "GetOsDetails", "data": {} }
```

```json
{ "kind": "GetOsDetails", "data": { "Arch": "amd64", "Os": "linux" } }
```

<sub>Response values depend on your machine. The other supported message is <code>GetUserName</code>.</sub>

## 🧰 Requirements

- 🐹 Go 1.16 or newer
- 🖥️ The system webview libraries required by [webview](https://github.com/webview/webview). On Linux the pinned module links with `pkg-config` packages `gtk+-3.0` and `webkit2gtk-4.0`.
<!-- TODO(Marc): confirm the platforms you actually tested on. -->

## 🚀 Run it

From the repository root, generate the WebAssembly assets:

```bash
go generate wasm/main.go
```

That copies `wasm_exec.js` into `storage/` from `$GOROOT/lib/wasm` (Go 1.24 and newer) or `$GOROOT/misc/wasm` (older Go), then builds `storage/main.wasm`.

Run the desktop app:

```bash
go run main.go
```

A webview window will appear, and there you can start querying for the system platform or the username that started the application.

## 🔗 See also

- 🦀 [rusty](https://github.com/Markcial/rusty): the same idea in pure Rust (wry + WASM)

## 📄 License

[MIT](LICENSE)

---

<div align="center">

Made by **Marc Garcia** · [marcgarcia.dev](https://marcgarcia.dev) · [@_marc_garcia on X](https://x.com/_marc_garcia) · [GitHub](https://github.com/Markcial)

</div>
