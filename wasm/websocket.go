// +build js,wasm

package main

import (
	"encoding/json"
	"syscall/js"

	"github.com/markcial/goita/bridge"
	"honnef.co/go/js/dom/v2"
)

// Websocket ...
type Websocket struct {
	*js.Value
}

func (w *Websocket) onOpen(handler func(e dom.Event)) {
	w.Call("addEventListener", "open", js.FuncOf(
		func(this js.Value, args []js.Value) interface{} {
			ev := &dom.BasicEvent{Value: args[0]}
			handler(ev)
			return nil
		},
	))
}

func (w *Websocket) onMessage(handler func(e dom.MessageEvent)) {
	w.Call("addEventListener", "message", js.FuncOf(
		func(this js.Value, args []js.Value) interface{} {
			ev := dom.MessageEvent{
				BasicEvent: &dom.BasicEvent{Value: args[0]},
			}
			handler(ev)
			return nil
		},
	))
}

func (w *Websocket) onClose(handler func(e dom.CloseEvent)) {
	w.Call("addEventListener", "close", js.FuncOf(
		func(this js.Value, args []js.Value) interface{} {
			ev := dom.CloseEvent{
				BasicEvent: &dom.BasicEvent{Value: args[0]},
			}
			handler(ev)
			return nil
		},
	))
}

func (w *Websocket) send(data []byte) {
	dt := js.Global().Get("Uint8Array").New(len(data))
	js.CopyBytesToJS(dt, data)
	w.Call("send", dt)
}

func (w *Websocket) command(kind bridge.Kind, data map[string]interface{}) {
	req := &bridge.Request{
		Kind: kind,
		Data: data,
	}
	js, _ := json.Marshal(req)
	println(string(js))
	w.send(js)
}

func connectWs() Websocket {
	params := js.Global().Call("getEnv")
	url := params.Get("wsURL").String()
	return newWs(url)
}

func newWs(url string) Websocket {
	w := js.Global().Get("WebSocket").New(url)
	return Websocket{&w}
}
