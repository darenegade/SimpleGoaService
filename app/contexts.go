// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=SimpleGoaService/design
// --out=$(GOPATH)/src/SimpleGoaService
// --version=v1.1.0-dirty
//
// API "HelloWorld": Application Contexts
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// HelloHelloContext provides the hello hello action context.
type HelloHelloContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *HelloHelloPayload
}

// NewHelloHelloContext parses the incoming request URL and body, performs validations and creates the
// context used by the hello controller hello action.
func NewHelloHelloContext(ctx context.Context, r *http.Request, service *goa.Service) (*HelloHelloContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := HelloHelloContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// helloHelloPayload is the hello hello action payload.
type helloHelloPayload struct {
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Publicize creates HelloHelloPayload from helloHelloPayload
func (payload *helloHelloPayload) Publicize() *HelloHelloPayload {
	var pub HelloHelloPayload
	if payload.Name != nil {
		pub.Name = payload.Name
	}
	return &pub
}

// HelloHelloPayload is the hello hello action payload.
type HelloHelloPayload struct {
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// OK sends a HTTP response with status code 200.
func (ctx *HelloHelloContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *HelloHelloContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}