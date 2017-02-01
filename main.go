//go:generate goagen bootstrap -d SimpleGoaService/design

package main

import (

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/darenegade/SimpleGoaService/app"
)

func main() {
	// Create service
	service := goa.New("HelloWorld")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "hello" controller
	c := NewHelloController(service)
	app.MountHelloController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}

