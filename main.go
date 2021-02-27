// +build !js

package main

import (
	_ "embed"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/markcial/goita/bridge"
	"github.com/webview/webview"
)

const (
	wasmExecPath = "/wasm_exec.js"
	mainWasmPath = "/main.wasm"
	wsPath       = "/ws"
)

var indexTemplate = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
		<script type="text/javascript" src="` + wasmExecPath + `"></script>
		<script type="text/javascript">
		function getEnv() {
			return {
				wsURL: "` + wsURL + `",
			}
		}
		async function loadWasm() {
			const source = await fetch("` + mainWasmPath + `");
			const go = new Go();
			const { instance } = await WebAssembly.instantiate(
				await source.arrayBuffer(),
				go.importObject
			);
			await go.run(instance);
		}
		loadWasm();
		</script>
	</body>
</html>
`

//go:embed storage/wasm_exec.js
var wasmExecJs []byte

//go:embed storage/main.wasm
var mainWasm []byte

// websocket configuration
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

const (
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

var (
	listener, _ = net.Listen("tcp", ":0")
	port        = listener.Addr().(*net.TCPAddr).Port
	url         = fmt.Sprintf("http://localhost:%d", port)
	wsURL       = fmt.Sprintf("ws://localhost:%d%s", port, wsPath)
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, string(indexTemplate))
}

func wasmExecHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, string(wasmExecJs))
}

func wasmAppHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/wasm")
	fmt.Fprint(w, string(mainWasm))
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}
	defer c.Close()
	for {
		req := &bridge.Request{}
		err := c.ReadJSON(req)
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %#v", req)
		res := bridge.Process(req)
		log.Printf("sent: %#v", res)
		err = c.WriteJSON(res)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func serve() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc(wasmExecPath, wasmExecHandler)
	http.HandleFunc(mainWasmPath, wasmAppHandler)
	http.HandleFunc(wsPath, websocketHandler)
	log.Fatal(http.Serve(listener, nil))
}

func main() {
	go serve()
	log.Printf("Server running on %s", url)
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate(url)
	w.Run()
}
