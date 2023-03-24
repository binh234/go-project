package ping

import (
	"context"
	"fmt"
	"log"
)

type Ping struct{}

func New() *Ping {
	return &Ping{}
}

// Default responds on root path `/ping`
// example of method with minimal arguments
// test with:
//
//	mantil invoke ping
//	mantil invoke ping/default
func (p *Ping) Default() string {
	return "pong"
}

// example simple mwthod with string input and output
// test with:
//
//	mantil invoke ping/hello --data "World"
func (p *Ping) Hello(ctx context.Context, name string) (string, error) {
	return "Hello, " + name, nil
}

type Request struct {
	Name string
}

type Response struct {
	Response string
}

// example of JSON api method
// test with:
//
//	mantil invoke ping/reqrsp --data '{"name":"World"}'
func (p *Ping) ReqRsp(ctx context.Context, req Request) (Response, error) {
	return Response{Response: "Hello, " + req.Name}, nil
}

// example same method with pointer arguments
// test with:
//
//	mantil invoke ping/reqrsp2
//	mantil invoke ping/reqrsp2 --data '{"name":"World"}'
func (p *Ping) ReqRsp2(ctx context.Context, req *Request) (*Response, error) {
	if req == nil {
		return nil, fmt.Errorf("request not found")
	}
	return &Response{Response: "Hello, " + req.Name}, nil
}

// example of getting function logs in invoke
// test with:
//
//	mantil invoke ping/logs --data '{"name":"Foo"}'
//
// You should see each log line before result.
func (p *Ping) Logs(ctx context.Context, req Request) (Response, error) {
	log.Printf("start Logs method")
	defer log.Printf("end")
	log.Printf("req.Name: '%s'", req.Name)
	return Response{Response: "Hello, " + req.Name}, nil
}
