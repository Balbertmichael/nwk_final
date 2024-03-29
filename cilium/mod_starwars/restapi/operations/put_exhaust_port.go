package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
  "log"
	"os"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutExhaustPortHandlerFunc turns a function with the right signature into a put exhaust port handler
type PutExhaustPortHandlerFunc func(PutExhaustPortParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutExhaustPortHandlerFunc) Handle(params PutExhaustPortParams) middleware.Responder {
	return fn(params)
}

// PutExhaustPortHandler interface for that can handle valid put exhaust port params
type PutExhaustPortHandler interface {
	Handle(PutExhaustPortParams) middleware.Responder
}

// NewPutExhaustPort creates a new http.Handler for the put exhaust port operation
func NewPutExhaustPort(ctx *middleware.Context, handler PutExhaustPortHandler) *PutExhaustPort {
	return &PutExhaustPort{Context: ctx, Handler: handler}
}

/*PutExhaustPort swagger:route PUT /exhaust-port putExhaustPort

Put something into the thermal exhaust port of the deathstar

*/
type PutExhaustPort struct {
	Context *middleware.Context
	Handler PutExhaustPortHandler
}

func (o *PutExhaustPort) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPutExhaustPortParams()

	f, err := os.OpenFile("request.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
					log.Fatal(err)
	}   

	//defer to close when you're done with it, not because you think it's idiomatic!
	defer f.Close()

	//set output of logs to f
	log.SetOutput(f)

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
    log.Printf("%s %s %s", Params.HTTPRequest.Method, Params.HTTPRequest.URL, Params.HTTPRequest.Body)
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	log.Printf("%s %s %s", Params.HTTPRequest.Method, Params.HTTPRequest.URL, Params.HTTPRequest.Body)
	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
