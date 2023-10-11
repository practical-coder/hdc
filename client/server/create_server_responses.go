// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v5/models"
)

// CreateServerReader is a Reader for the CreateServer structure.
type CreateServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateServerCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateServerAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateServerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateServerConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateServerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateServerCreated creates a CreateServerCreated with default headers values
func NewCreateServerCreated() *CreateServerCreated {
	return &CreateServerCreated{}
}

/*
CreateServerCreated describes a response with status code 201, with default header values.

Server created
*/
type CreateServerCreated struct {
	Payload *models.Server
}

// IsSuccess returns true when this create server created response has a 2xx status code
func (o *CreateServerCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create server created response has a 3xx status code
func (o *CreateServerCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create server created response has a 4xx status code
func (o *CreateServerCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create server created response has a 5xx status code
func (o *CreateServerCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create server created response a status code equal to that given
func (o *CreateServerCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create server created response
func (o *CreateServerCreated) Code() int {
	return 201
}

func (o *CreateServerCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/servers][%d] createServerCreated  %+v", 201, o.Payload)
}

func (o *CreateServerCreated) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/servers][%d] createServerCreated  %+v", 201, o.Payload)
}

func (o *CreateServerCreated) GetPayload() *models.Server {
	return o.Payload
}

func (o *CreateServerCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Server)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServerAccepted creates a CreateServerAccepted with default headers values
func NewCreateServerAccepted() *CreateServerAccepted {
	return &CreateServerAccepted{}
}

/*
CreateServerAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateServerAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.Server
}

// IsSuccess returns true when this create server accepted response has a 2xx status code
func (o *CreateServerAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create server accepted response has a 3xx status code
func (o *CreateServerAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create server accepted response has a 4xx status code
func (o *CreateServerAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this create server accepted response has a 5xx status code
func (o *CreateServerAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this create server accepted response a status code equal to that given
func (o *CreateServerAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the create server accepted response
func (o *CreateServerAccepted) Code() int {
	return 202
}

func (o *CreateServerAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/servers][%d] createServerAccepted  %+v", 202, o.Payload)
}

func (o *CreateServerAccepted) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/servers][%d] createServerAccepted  %+v", 202, o.Payload)
}

func (o *CreateServerAccepted) GetPayload() *models.Server {
	return o.Payload
}

func (o *CreateServerAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.Server)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServerBadRequest creates a CreateServerBadRequest with default headers values
func NewCreateServerBadRequest() *CreateServerBadRequest {
	return &CreateServerBadRequest{}
}

/*
CreateServerBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateServerBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create server bad request response has a 2xx status code
func (o *CreateServerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create server bad request response has a 3xx status code
func (o *CreateServerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create server bad request response has a 4xx status code
func (o *CreateServerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create server bad request response has a 5xx status code
func (o *CreateServerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create server bad request response a status code equal to that given
func (o *CreateServerBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create server bad request response
func (o *CreateServerBadRequest) Code() int {
	return 400
}

func (o *CreateServerBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/servers][%d] createServerBadRequest  %+v", 400, o.Payload)
}

func (o *CreateServerBadRequest) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/servers][%d] createServerBadRequest  %+v", 400, o.Payload)
}

func (o *CreateServerBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateServerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateServerConflict creates a CreateServerConflict with default headers values
func NewCreateServerConflict() *CreateServerConflict {
	return &CreateServerConflict{}
}

/*
CreateServerConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateServerConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create server conflict response has a 2xx status code
func (o *CreateServerConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create server conflict response has a 3xx status code
func (o *CreateServerConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create server conflict response has a 4xx status code
func (o *CreateServerConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create server conflict response has a 5xx status code
func (o *CreateServerConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create server conflict response a status code equal to that given
func (o *CreateServerConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create server conflict response
func (o *CreateServerConflict) Code() int {
	return 409
}

func (o *CreateServerConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/servers][%d] createServerConflict  %+v", 409, o.Payload)
}

func (o *CreateServerConflict) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/servers][%d] createServerConflict  %+v", 409, o.Payload)
}

func (o *CreateServerConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateServerConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateServerDefault creates a CreateServerDefault with default headers values
func NewCreateServerDefault(code int) *CreateServerDefault {
	return &CreateServerDefault{
		_statusCode: code,
	}
}

/*
CreateServerDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateServerDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create server default response has a 2xx status code
func (o *CreateServerDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create server default response has a 3xx status code
func (o *CreateServerDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create server default response has a 4xx status code
func (o *CreateServerDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create server default response has a 5xx status code
func (o *CreateServerDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create server default response a status code equal to that given
func (o *CreateServerDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create server default response
func (o *CreateServerDefault) Code() int {
	return o._statusCode
}

func (o *CreateServerDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/servers][%d] createServer default  %+v", o._statusCode, o.Payload)
}

func (o *CreateServerDefault) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/servers][%d] createServer default  %+v", o._statusCode, o.Payload)
}

func (o *CreateServerDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateServerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
