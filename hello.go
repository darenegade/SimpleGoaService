package main

import (
	"github.com/goadesign/goa"
	"github.com/darenegade/SimpleGoaService/app"
)

// HelloController implements the hello resource.
type HelloController struct {
	*goa.Controller
}

// NewHelloController creates a hello controller.
func NewHelloController(service *goa.Service) *HelloController {
	return &HelloController{Controller: service.NewController("HelloController")}
}

// Hello runs the hello action.
func (c *HelloController) Hello(ctx *app.HelloHelloContext) error {

	return ctx.OK([]byte("Hello " + *ctx.Payload.Name))
}
