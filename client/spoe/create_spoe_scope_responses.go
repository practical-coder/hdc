// Code generated by go-swagger; DO NOT EDIT.

package spoe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/cloud-native/models"
)

// CreateSpoeScopeReader is a Reader for the CreateSpoeScope structure.
type CreateSpoeScopeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSpoeScopeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateSpoeScopeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateSpoeScopeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateSpoeScopeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateSpoeScopeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSpoeScopeCreated creates a CreateSpoeScopeCreated with default headers values
func NewCreateSpoeScopeCreated() *CreateSpoeScopeCreated {
	return &CreateSpoeScopeCreated{}
}

/* CreateSpoeScopeCreated describes a response with status code 201, with default header values.

Spoe scope created
*/
type CreateSpoeScopeCreated struct {
	Payload models.SpoeScope
}

func (o *CreateSpoeScopeCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_scopes][%d] createSpoeScopeCreated  %+v", 201, o.Payload)
}
func (o *CreateSpoeScopeCreated) GetPayload() models.SpoeScope {
	return o.Payload
}

func (o *CreateSpoeScopeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSpoeScopeBadRequest creates a CreateSpoeScopeBadRequest with default headers values
func NewCreateSpoeScopeBadRequest() *CreateSpoeScopeBadRequest {
	return &CreateSpoeScopeBadRequest{}
}

/* CreateSpoeScopeBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateSpoeScopeBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateSpoeScopeBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_scopes][%d] createSpoeScopeBadRequest  %+v", 400, o.Payload)
}
func (o *CreateSpoeScopeBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSpoeScopeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateSpoeScopeConflict creates a CreateSpoeScopeConflict with default headers values
func NewCreateSpoeScopeConflict() *CreateSpoeScopeConflict {
	return &CreateSpoeScopeConflict{}
}

/* CreateSpoeScopeConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateSpoeScopeConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateSpoeScopeConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_scopes][%d] createSpoeScopeConflict  %+v", 409, o.Payload)
}
func (o *CreateSpoeScopeConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSpoeScopeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateSpoeScopeDefault creates a CreateSpoeScopeDefault with default headers values
func NewCreateSpoeScopeDefault(code int) *CreateSpoeScopeDefault {
	return &CreateSpoeScopeDefault{
		_statusCode: code,
	}
}

/* CreateSpoeScopeDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateSpoeScopeDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create spoe scope default response
func (o *CreateSpoeScopeDefault) Code() int {
	return o._statusCode
}

func (o *CreateSpoeScopeDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_scopes][%d] createSpoeScope default  %+v", o._statusCode, o.Payload)
}
func (o *CreateSpoeScopeDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSpoeScopeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
