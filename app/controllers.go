// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=SimpleGoaService/design
// --out=$(GOPATH)/src/SimpleGoaService
// --version=v1.1.0-dirty
//
// API "HelloWorld": Application Controllers
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// HelloController is the controller interface for the Hello actions.
type HelloController interface {
	goa.Muxer
	Hello(*HelloHelloContext) error
}

// MountHelloController "mounts" a Hello resource controller on the given service.
func MountHelloController(service *goa.Service, ctrl HelloController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewHelloHelloContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*HelloHelloPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Hello(rctx)
	}
	service.Mux.Handle("POST", "/hello", ctrl.MuxHandler("Hello", h, unmarshalHelloHelloPayload))
	service.LogInfo("mount", "ctrl", "Hello", "action", "Hello", "route", "POST /hello")
}

// unmarshalHelloHelloPayload unmarshals the request body into the context request data Payload field.
func unmarshalHelloHelloPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &helloHelloPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}