// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/loxilb-io/loxilb/api/models"
)

// GetStatusFilesystemHandlerFunc turns a function with the right signature into a get status filesystem handler
type GetStatusFilesystemHandlerFunc func(GetStatusFilesystemParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetStatusFilesystemHandlerFunc) Handle(params GetStatusFilesystemParams) middleware.Responder {
	return fn(params)
}

// GetStatusFilesystemHandler interface for that can handle valid get status filesystem params
type GetStatusFilesystemHandler interface {
	Handle(GetStatusFilesystemParams) middleware.Responder
}

// NewGetStatusFilesystem creates a new http.Handler for the get status filesystem operation
func NewGetStatusFilesystem(ctx *middleware.Context, handler GetStatusFilesystemHandler) *GetStatusFilesystem {
	return &GetStatusFilesystem{Context: ctx, Handler: handler}
}

/*
	GetStatusFilesystem swagger:route GET /status/filesystem getStatusFilesystem

# Get a File System info in the device

Get a File system infomation (linux command "df") in the device or system.
*/
type GetStatusFilesystem struct {
	Context *middleware.Context
	Handler GetStatusFilesystemHandler
}

func (o *GetStatusFilesystem) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetStatusFilesystemParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetStatusFilesystemOKBody get status filesystem o k body
//
// swagger:model GetStatusFilesystemOKBody
type GetStatusFilesystemOKBody struct {

	// filesystem attr
	FilesystemAttr []*models.FileSystemInfoEntry `json:"filesystemAttr"`
}

// Validate validates this get status filesystem o k body
func (o *GetStatusFilesystemOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFilesystemAttr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStatusFilesystemOKBody) validateFilesystemAttr(formats strfmt.Registry) error {
	if swag.IsZero(o.FilesystemAttr) { // not required
		return nil
	}

	for i := 0; i < len(o.FilesystemAttr); i++ {
		if swag.IsZero(o.FilesystemAttr[i]) { // not required
			continue
		}

		if o.FilesystemAttr[i] != nil {
			if err := o.FilesystemAttr[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getStatusFilesystemOK" + "." + "filesystemAttr" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getStatusFilesystemOK" + "." + "filesystemAttr" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get status filesystem o k body based on the context it is used
func (o *GetStatusFilesystemOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFilesystemAttr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStatusFilesystemOKBody) contextValidateFilesystemAttr(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.FilesystemAttr); i++ {

		if o.FilesystemAttr[i] != nil {
			if err := o.FilesystemAttr[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getStatusFilesystemOK" + "." + "filesystemAttr" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getStatusFilesystemOK" + "." + "filesystemAttr" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetStatusFilesystemOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStatusFilesystemOKBody) UnmarshalBinary(b []byte) error {
	var res GetStatusFilesystemOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
