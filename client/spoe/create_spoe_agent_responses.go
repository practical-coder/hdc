// Code generated by go-swagger; DO NOT EDIT.

package spoe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v5/models"
)

// CreateSpoeAgentReader is a Reader for the CreateSpoeAgent structure.
type CreateSpoeAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSpoeAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateSpoeAgentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateSpoeAgentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateSpoeAgentConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateSpoeAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSpoeAgentCreated creates a CreateSpoeAgentCreated with default headers values
func NewCreateSpoeAgentCreated() *CreateSpoeAgentCreated {
	return &CreateSpoeAgentCreated{}
}

/*
CreateSpoeAgentCreated describes a response with status code 201, with default header values.

Spoe agent created
*/
type CreateSpoeAgentCreated struct {
	Payload *models.SpoeAgent
}

// IsSuccess returns true when this create spoe agent created response has a 2xx status code
func (o *CreateSpoeAgentCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create spoe agent created response has a 3xx status code
func (o *CreateSpoeAgentCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create spoe agent created response has a 4xx status code
func (o *CreateSpoeAgentCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create spoe agent created response has a 5xx status code
func (o *CreateSpoeAgentCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create spoe agent created response a status code equal to that given
func (o *CreateSpoeAgentCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create spoe agent created response
func (o *CreateSpoeAgentCreated) Code() int {
	return 201
}

func (o *CreateSpoeAgentCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_agents][%d] createSpoeAgentCreated  %+v", 201, o.Payload)
}

func (o *CreateSpoeAgentCreated) String() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_agents][%d] createSpoeAgentCreated  %+v", 201, o.Payload)
}

func (o *CreateSpoeAgentCreated) GetPayload() *models.SpoeAgent {
	return o.Payload
}

func (o *CreateSpoeAgentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SpoeAgent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSpoeAgentBadRequest creates a CreateSpoeAgentBadRequest with default headers values
func NewCreateSpoeAgentBadRequest() *CreateSpoeAgentBadRequest {
	return &CreateSpoeAgentBadRequest{}
}

/*
CreateSpoeAgentBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateSpoeAgentBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create spoe agent bad request response has a 2xx status code
func (o *CreateSpoeAgentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create spoe agent bad request response has a 3xx status code
func (o *CreateSpoeAgentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create spoe agent bad request response has a 4xx status code
func (o *CreateSpoeAgentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create spoe agent bad request response has a 5xx status code
func (o *CreateSpoeAgentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create spoe agent bad request response a status code equal to that given
func (o *CreateSpoeAgentBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create spoe agent bad request response
func (o *CreateSpoeAgentBadRequest) Code() int {
	return 400
}

func (o *CreateSpoeAgentBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_agents][%d] createSpoeAgentBadRequest  %+v", 400, o.Payload)
}

func (o *CreateSpoeAgentBadRequest) String() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_agents][%d] createSpoeAgentBadRequest  %+v", 400, o.Payload)
}

func (o *CreateSpoeAgentBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSpoeAgentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateSpoeAgentConflict creates a CreateSpoeAgentConflict with default headers values
func NewCreateSpoeAgentConflict() *CreateSpoeAgentConflict {
	return &CreateSpoeAgentConflict{}
}

/*
CreateSpoeAgentConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateSpoeAgentConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create spoe agent conflict response has a 2xx status code
func (o *CreateSpoeAgentConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create spoe agent conflict response has a 3xx status code
func (o *CreateSpoeAgentConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create spoe agent conflict response has a 4xx status code
func (o *CreateSpoeAgentConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create spoe agent conflict response has a 5xx status code
func (o *CreateSpoeAgentConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create spoe agent conflict response a status code equal to that given
func (o *CreateSpoeAgentConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create spoe agent conflict response
func (o *CreateSpoeAgentConflict) Code() int {
	return 409
}

func (o *CreateSpoeAgentConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_agents][%d] createSpoeAgentConflict  %+v", 409, o.Payload)
}

func (o *CreateSpoeAgentConflict) String() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_agents][%d] createSpoeAgentConflict  %+v", 409, o.Payload)
}

func (o *CreateSpoeAgentConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSpoeAgentConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateSpoeAgentDefault creates a CreateSpoeAgentDefault with default headers values
func NewCreateSpoeAgentDefault(code int) *CreateSpoeAgentDefault {
	return &CreateSpoeAgentDefault{
		_statusCode: code,
	}
}

/*
CreateSpoeAgentDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateSpoeAgentDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create spoe agent default response has a 2xx status code
func (o *CreateSpoeAgentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create spoe agent default response has a 3xx status code
func (o *CreateSpoeAgentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create spoe agent default response has a 4xx status code
func (o *CreateSpoeAgentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create spoe agent default response has a 5xx status code
func (o *CreateSpoeAgentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create spoe agent default response a status code equal to that given
func (o *CreateSpoeAgentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create spoe agent default response
func (o *CreateSpoeAgentDefault) Code() int {
	return o._statusCode
}

func (o *CreateSpoeAgentDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_agents][%d] createSpoeAgent default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSpoeAgentDefault) String() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_agents][%d] createSpoeAgent default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSpoeAgentDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSpoeAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
