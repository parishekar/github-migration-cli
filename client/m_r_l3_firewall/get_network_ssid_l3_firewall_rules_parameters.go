// Code generated by go-swagger; DO NOT EDIT.

package m_r_l3_firewall

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

// NewGetNetworkSsidL3FirewallRulesParams creates a new GetNetworkSsidL3FirewallRulesParams object
// with the default values initialized.
func NewGetNetworkSsidL3FirewallRulesParams() *GetNetworkSsidL3FirewallRulesParams {
	var ()
	return &GetNetworkSsidL3FirewallRulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkSsidL3FirewallRulesParamsWithTimeout creates a new GetNetworkSsidL3FirewallRulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkSsidL3FirewallRulesParamsWithTimeout(timeout time.Duration) *GetNetworkSsidL3FirewallRulesParams {
	var ()
	return &GetNetworkSsidL3FirewallRulesParams{

		timeout: timeout,
	}
}

// NewGetNetworkSsidL3FirewallRulesParamsWithContext creates a new GetNetworkSsidL3FirewallRulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkSsidL3FirewallRulesParamsWithContext(ctx context.Context) *GetNetworkSsidL3FirewallRulesParams {
	var ()
	return &GetNetworkSsidL3FirewallRulesParams{

		Context: ctx,
	}
}

// NewGetNetworkSsidL3FirewallRulesParamsWithHTTPClient creates a new GetNetworkSsidL3FirewallRulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkSsidL3FirewallRulesParamsWithHTTPClient(client *http.Client) *GetNetworkSsidL3FirewallRulesParams {
	var ()
	return &GetNetworkSsidL3FirewallRulesParams{
		HTTPClient: client,
	}
}

/*GetNetworkSsidL3FirewallRulesParams contains all the parameters to send to the API endpoint
for the get network ssid l3 firewall rules operation typically these are written to a http.Request
*/
type GetNetworkSsidL3FirewallRulesParams struct {

	/*NetworkID*/
	NetworkID string
	/*Number*/
	Number string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network ssid l3 firewall rules params
func (o *GetNetworkSsidL3FirewallRulesParams) WithTimeout(timeout time.Duration) *GetNetworkSsidL3FirewallRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network ssid l3 firewall rules params
func (o *GetNetworkSsidL3FirewallRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network ssid l3 firewall rules params
func (o *GetNetworkSsidL3FirewallRulesParams) WithContext(ctx context.Context) *GetNetworkSsidL3FirewallRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network ssid l3 firewall rules params
func (o *GetNetworkSsidL3FirewallRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network ssid l3 firewall rules params
func (o *GetNetworkSsidL3FirewallRulesParams) WithHTTPClient(client *http.Client) *GetNetworkSsidL3FirewallRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network ssid l3 firewall rules params
func (o *GetNetworkSsidL3FirewallRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network ssid l3 firewall rules params
func (o *GetNetworkSsidL3FirewallRulesParams) WithNetworkID(networkID string) *GetNetworkSsidL3FirewallRulesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network ssid l3 firewall rules params
func (o *GetNetworkSsidL3FirewallRulesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithNumber adds the number to the get network ssid l3 firewall rules params
func (o *GetNetworkSsidL3FirewallRulesParams) WithNumber(number string) *GetNetworkSsidL3FirewallRulesParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the get network ssid l3 firewall rules params
func (o *GetNetworkSsidL3FirewallRulesParams) SetNumber(number string) {
	o.Number = number
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkSsidL3FirewallRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	// path param number
	if err := r.SetPathParam("number", o.Number); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}