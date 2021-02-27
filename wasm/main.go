// +build js

package main

import (
	"github.com/markcial/goita/bridge"
	dom "honnef.co/go/js/dom/v2"
)

//go:generate cp $GOROOT/misc/wasm/wasm_exec.js ../storage/wasm_exec.js
//go:generate env GOOS=js GOARCH=wasm go build -o ../storage/main.wasm ./

var (
	window = dom.GetWindow()
	doc    = window.Document()
	body   = doc.GetElementsByTagName("body")[0]
)

func createSelect() *dom.HTMLSelectElement {
	s := doc.CreateElement("select").(*dom.HTMLSelectElement)
	for _, v := range []bridge.Kind{bridge.GetOsDetails, bridge.GetUserName} {
		opt := doc.CreateElement("option").(*dom.HTMLOptionElement)
		opt.SetValue(string(v))
		opt.SetTextContent(string(v))
		s.AppendChild(opt)
	}
	return s
}

func createButton() *dom.HTMLButtonElement {
	bt := doc.CreateElement("button").(*dom.HTMLButtonElement)
	bt.SetInnerHTML("Click me")
	return bt
}

func createLogger() *dom.HTMLPreElement {
	pre := doc.CreateElement("pre").(*dom.HTMLPreElement)
	return pre
}

func main() {
	c := make(chan struct{}, 0)
	ws := connectWs()
	//
	sel := createSelect()
	body.AppendChild(sel)
	//
	bt := createButton()
	bt.AddEventListener("click", false, func(e dom.Event) {
		ws.command(sel.Value(), map[string]interface{}{})
	})
	body.AppendChild(bt)
	//
	lg := createLogger()
	body.AppendChild(lg)
	//
	ws.onMessage(func(e dom.MessageEvent) {
		lg.SetInnerHTML(e.Data().String())
	})
	<-c
}
