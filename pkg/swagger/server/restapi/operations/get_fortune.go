// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetFortuneHandlerFunc turns a function with the right signature into a get fortune handler
type GetFortuneHandlerFunc func(GetFortuneParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFortuneHandlerFunc) Handle(params GetFortuneParams) middleware.Responder {
	return fn(params)
}

// GetFortuneHandler interface for that can handle valid get fortune params
type GetFortuneHandler interface {
	Handle(GetFortuneParams) middleware.Responder
}

// NewGetFortune creates a new http.Handler for the get fortune operation
func NewGetFortune(ctx *middleware.Context, handler GetFortuneHandler) *GetFortune {
	return &GetFortune{Context: ctx, Handler: handler}
}

/*GetFortune swagger:route GET /fortune getFortune

Returns a random fortune

*/
type GetFortune struct {
	Context *middleware.Context
	Handler GetFortuneHandler
}

func (o *GetFortune) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetFortuneParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetFortuneOKBody get fortune o k body
//
// swagger:model GetFortuneOKBody
type GetFortuneOKBody struct {

	// Fortune message.
	Fortune string `json:"fortune,omitempty"`
}

// Validate validates this get fortune o k body
func (o *GetFortuneOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetFortuneOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFortuneOKBody) UnmarshalBinary(b []byte) error {
	var res GetFortuneOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
