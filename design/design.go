package design                                     // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("HelloWorld", func() {                     // API defines the microservice endpoint and
	Title("Hello World Service")           // other global properties. There should be one
	Description("A simple Hello World service")        // and exactly one API definition appearing in
	Scheme("http")                             // the design.
	Host("localhost:8080")
})

// JWT defines a security scheme using JWT.  The scheme uses the "Authorization" header to lookup
// the token.  It also defines then scope "api".
var JWT = JWTSecurity("jwt", func() {
	Header("Authorization")
	Scope("api:access", "API access") // Define "api:access" scope
})

var _ = Resource("hello", func() {                // Resources group related API endpoints
	BasePath("/hello")                       // together. They map to REST resources for REST

	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	})

	Action("hello", func() {                    // Actions define a single API endpoint together
		Description("Say Hello to the service")    // with its path, parameters (both path
		Routing(POST(""))         // parameters and querystring values) and payload
		Payload(func() {
			Member("name")
		})
		Response(OK)                       // Responses define the shape and status code
		NoSecurity()
	})
})