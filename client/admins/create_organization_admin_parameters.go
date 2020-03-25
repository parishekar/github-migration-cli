// Code generated by go-swagger; DO NOT EDIT.

package admins

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

	"github.com/cisco-sso/meraki-cli/models"
)

// NewCreateOrganizationAdminParams creates a new CreateOrganizationAdminParams object
// with the default values initialized.
func NewCreateOrganizationAdminParams() *CreateOrganizationAdminParams {
	var ()
	return &CreateOrganizationAdminParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrganizationAdminParamsWithTimeout creates a new CreateOrganizationAdminParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateOrganizationAdminParamsWithTimeout(timeout time.Duration) *CreateOrganizationAdminParams {
	var ()
	return &CreateOrganizationAdminParams{

		timeout: timeout,
	}
}

// NewCreateOrganizationAdminParamsWithContext creates a new CreateOrganizationAdminParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateOrganizationAdminParamsWithContext(ctx context.Context) *CreateOrganizationAdminParams {
	var ()
	return &CreateOrganizationAdminParams{

		Context: ctx,
	}
}

// NewCreateOrganizationAdminParamsWithHTTPClient creates a new CreateOrganizationAdminParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateOrganizationAdminParamsWithHTTPClient(client *http.Client) *CreateOrganizationAdminParams {
	var ()
	return &CreateOrganizationAdminParams{
		HTTPClient: client,
	}
}

/*CreateOrganizationAdminParams contains all the parameters to send to the API endpoint
for the create organization admin operation typically these are written to a http.Request
*/
type CreateOrganizationAdminParams struct {

	/*CreateOrganizationAdmin*/
	CreateOrganizationAdmin *models.CreateOrganizationAdmin
	/*OrganizationID*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create organization admin params
func (o *CreateOrganizationAdminParams) WithTimeout(timeout time.Duration) *CreateOrganizationAdminParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create organization admin params
func (o *CreateOrganizationAdminParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create organization admin params
func (o *CreateOrganizationAdminParams) WithContext(ctx context.Context) *CreateOrganizationAdminParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create organization admin params
func (o *CreateOrganizationAdminParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create organization admin params
func (o *CreateOrganizationAdminParams) WithHTTPClient(client *http.Client) *CreateOrganizationAdminParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create organization admin params
func (o *CreateOrganizationAdminParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateOrganizationAdmin adds the createOrganizationAdmin to the create organization admin params
func (o *CreateOrganizationAdminParams) WithCreateOrganizationAdmin(createOrganizationAdmin *models.CreateOrganizationAdmin) *CreateOrganizationAdminParams {
	o.SetCreateOrganizationAdmin(createOrganizationAdmin)
	return o
}

// SetCreateOrganizationAdmin adds the createOrganizationAdmin to the create organization admin params
func (o *CreateOrganizationAdminParams) SetCreateOrganizationAdmin(createOrganizationAdmin *models.CreateOrganizationAdmin) {
	o.CreateOrganizationAdmin = createOrganizationAdmin
}

// WithOrganizationID adds the organizationID to the create organization admin params
func (o *CreateOrganizationAdminParams) WithOrganizationID(organizationID string) *CreateOrganizationAdminParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create organization admin params
func (o *CreateOrganizationAdminParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrganizationAdminParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateOrganizationAdmin != nil {
		if err := r.SetBodyParam(o.CreateOrganizationAdmin); err != nil {
			return err
		}
	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}