// Code generated by go-swagger; DO NOT EDIT.

package admins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateOrganizationAdminReader is a Reader for the CreateOrganizationAdmin structure.
type CreateOrganizationAdminReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrganizationAdminReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateOrganizationAdminCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateOrganizationAdminCreated creates a CreateOrganizationAdminCreated with default headers values
func NewCreateOrganizationAdminCreated() *CreateOrganizationAdminCreated {
	return &CreateOrganizationAdminCreated{}
}

/*CreateOrganizationAdminCreated handles this case with default header values.

Successful operation
*/
type CreateOrganizationAdminCreated struct {
	Payload interface{}
}

func (o *CreateOrganizationAdminCreated) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/admins][%d] createOrganizationAdminCreated  %+v", 201, o.Payload)
}

func (o *CreateOrganizationAdminCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateOrganizationAdminCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}