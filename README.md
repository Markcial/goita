# goita
Minimal desktop application with webview, websockets and WASM

## Description

Requires go 1.16

This is an experiment on how to create a minimal desktop application using only go language.

The wasm folder contains all the webassembly code that will be generated inside the `storage` folder.

The bridge folder is just a set of types to enforce communication between the go application code and the WASM through JSON messages on websockets.

Inside the `main.go` embeds the WASM files and starts a webserver on a random port to enable communication between WASM and the application.

## Run it

Just generate the WAMS code with:

```
go generate wasm/main.go
```

and then run it

```
go run main.go
```

A webview window will appear and there you can start querying for the system platform or the username that started the application.

![example](https://user-images.githubusercontent.com/208523/109386098-c1963000-78f8-11eb-84c2-7a204b1eb031.gif)

