// Code generated by go-swagger; DO NOT EDIT.

package dashboard_branding_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateOrganizationBrandingPolicyReader is a Reader for the CreateOrganizationBrandingPolicy structure.
type CreateOrganizationBrandingPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrganizationBrandingPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateOrganizationBrandingPolicyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateOrganizationBrandingPolicyCreated creates a CreateOrganizationBrandingPolicyCreated with default headers values
func NewCreateOrganizationBrandingPolicyCreated() *CreateOrganizationBrandingPolicyCreated {
	return &CreateOrganizationBrandingPolicyCreated{}
}

/*CreateOrganizationBrandingPolicyCreated handles this case with default header values.

Successful operation
*/
type CreateOrganizationBrandingPolicyCreated struct {
	Payload interface{}
}

func (o *CreateOrganizationBrandingPolicyCreated) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/brandingPolicies][%d] createOrganizationBrandingPolicyCreated  %+v", 201, o.Payload)
}

func (o *CreateOrganizationBrandingPolicyCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateOrganizationBrandingPolicyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}