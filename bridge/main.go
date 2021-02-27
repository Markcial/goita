package bridge

import (
	"os/user"
	"runtime"
)

// Kind ...
type Kind = string

const (
	GetOsDetails string = "GetOsDetails"
	GetUserName         = "GetUserName"
)

// Request ...
type Request struct {
	Kind Kind                   `json:"kind"`
	Data map[string]interface{} `json:"data"`
}

// Response ...
type Response struct {
	Kind Kind                   `json:"kind"`
	Data map[string]interface{} `json:"data"`
}

// Process ...
func Process(r *Request) *Response {
	switch r.Kind {
	case GetOsDetails:
		return &Response{
			Kind: GetOsDetails,
			Data: map[string]interface{}{
				"Arch": runtime.GOARCH,
				"Os":   runtime.GOOS,
			},
		}
	case GetUserName:
		u, _ := user.Current()
		return &Response{
			Kind: GetUserName,
			Data: map[string]interface{}{
				"username": u.Name,
			},
		}
	}
	return nil
}
