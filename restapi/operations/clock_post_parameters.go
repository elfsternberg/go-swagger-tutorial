// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "github.com/elfsternberg/timeofday/models"
)

// NewClockPostParams creates a new ClockPostParams object
// no default values defined in spec.
func NewClockPostParams() ClockPostParams {

	return ClockPostParams{}
}

// ClockPostParams contains all the bound params for the clock post operation
// typically these are obtained from a http.Request
//
// swagger:parameters ClockPost
type ClockPostParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Timezone to return
	  In: body
	*/
	Timezone *models.Timezone
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewClockPostParams() beforehand.
func (o *ClockPostParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Timezone
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("timezone", "body", "", err))
		} else {

			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Timezone = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}