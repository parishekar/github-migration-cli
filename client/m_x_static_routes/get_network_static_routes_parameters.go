// Code generated by go-swagger; DO NOT EDIT.

package m_x_static_routes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetNetworkStaticRoutesParams creates a new GetNetworkStaticRoutesParams object
// with the default values initialized.
func NewGetNetworkStaticRoutesParams() *GetNetworkStaticRoutesParams {
	var ()
	return &GetNetworkStaticRoutesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkStaticRoutesParamsWithTimeout creates a new GetNetworkStaticRoutesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkStaticRoutesParamsWithTimeout(timeout time.Duration) *GetNetworkStaticRoutesParams {
	var ()
	return &GetNetworkStaticRoutesParams{

		timeout: timeout,
	}
}

// NewGetNetworkStaticRoutesParamsWithContext creates a new GetNetworkStaticRoutesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkStaticRoutesParamsWithContext(ctx context.Context) *GetNetworkStaticRoutesParams {
	var ()
	return &GetNetworkStaticRoutesParams{

		Context: ctx,
	}
}

// NewGetNetworkStaticRoutesParamsWithHTTPClient creates a new GetNetworkStaticRoutesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkStaticRoutesParamsWithHTTPClient(client *http.Client) *GetNetworkStaticRoutesParams {
	var ()
	return &GetNetworkStaticRoutesParams{
		HTTPClient: client,
	}
}

/*GetNetworkStaticRoutesParams contains all the parameters to send to the API endpoint
for the get network static routes operation typically these are written to a http.Request
*/
type GetNetworkStaticRoutesParams struct {

	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network static routes params
func (o *GetNetworkStaticRoutesParams) WithTimeout(timeout time.Duration) *GetNetworkStaticRoutesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network static routes params
func (o *GetNetworkStaticRoutesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network static routes params
func (o *GetNetworkStaticRoutesParams) WithContext(ctx context.Context) *GetNetworkStaticRoutesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network static routes params
func (o *GetNetworkStaticRoutesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network static routes params
func (o *GetNetworkStaticRoutesParams) WithHTTPClient(client *http.Client) *GetNetworkStaticRoutesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network static routes params
func (o *GetNetworkStaticRoutesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network static routes params
func (o *GetNetworkStaticRoutesParams) WithNetworkID(networkID string) *GetNetworkStaticRoutesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network static routes params
func (o *GetNetworkStaticRoutesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkStaticRoutesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}