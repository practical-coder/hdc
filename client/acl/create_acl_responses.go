// Code generated by go-swagger; DO NOT EDIT.

package acl

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v2/models"
)

// CreateACLReader is a Reader for the CreateACL structure.
type CreateACLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateACLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateACLCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateACLAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateACLBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateACLConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateACLDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateACLCreated creates a CreateACLCreated with default headers values
func NewCreateACLCreated() *CreateACLCreated {
	return &CreateACLCreated{}
}

/* CreateACLCreated describes a response with status code 201, with default header values.

ACL line created
*/
type CreateACLCreated struct {
	Payload *models.ACL
}

func (o *CreateACLCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/acls][%d] createAclCreated  %+v", 201, o.Payload)
}
func (o *CreateACLCreated) GetPayload() *models.ACL {
	return o.Payload
}

func (o *CreateACLCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ACL)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateACLAccepted creates a CreateACLAccepted with default headers values
func NewCreateACLAccepted() *CreateACLAccepted {
	return &CreateACLAccepted{}
}

/* CreateACLAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateACLAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.ACL
}

func (o *CreateACLAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/acls][%d] createAclAccepted  %+v", 202, o.Payload)
}
func (o *CreateACLAccepted) GetPayload() *models.ACL {
	return o.Payload
}

func (o *CreateACLAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.ACL)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateACLBadRequest creates a CreateACLBadRequest with default headers values
func NewCreateACLBadRequest() *CreateACLBadRequest {
	return &CreateACLBadRequest{}
}

/* CreateACLBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateACLBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateACLBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/acls][%d] createAclBadRequest  %+v", 400, o.Payload)
}
func (o *CreateACLBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateACLBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateACLConflict creates a CreateACLConflict with default headers values
func NewCreateACLConflict() *CreateACLConflict {
	return &CreateACLConflict{}
}

/* CreateACLConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateACLConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateACLConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/acls][%d] createAclConflict  %+v", 409, o.Payload)
}
func (o *CreateACLConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateACLConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateACLDefault creates a CreateACLDefault with default headers values
func NewCreateACLDefault(code int) *CreateACLDefault {
	return &CreateACLDefault{
		_statusCode: code,
	}
}

/* CreateACLDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateACLDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create Acl default response
func (o *CreateACLDefault) Code() int {
	return o._statusCode
}

func (o *CreateACLDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/acls][%d] createAcl default  %+v", o._statusCode, o.Payload)
}
func (o *CreateACLDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateACLDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
