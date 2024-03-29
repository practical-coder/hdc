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

// ReplaceSpoeAgentReader is a Reader for the ReplaceSpoeAgent structure.
type ReplaceSpoeAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceSpoeAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceSpoeAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceSpoeAgentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceSpoeAgentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceSpoeAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceSpoeAgentOK creates a ReplaceSpoeAgentOK with default headers values
func NewReplaceSpoeAgentOK() *ReplaceSpoeAgentOK {
	return &ReplaceSpoeAgentOK{}
}

/*
ReplaceSpoeAgentOK describes a response with status code 200, with default header values.

Spoe agent replaced
*/
type ReplaceSpoeAgentOK struct {
	Payload *models.SpoeAgent
}

// IsSuccess returns true when this replace spoe agent o k response has a 2xx status code
func (o *ReplaceSpoeAgentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this replace spoe agent o k response has a 3xx status code
func (o *ReplaceSpoeAgentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace spoe agent o k response has a 4xx status code
func (o *ReplaceSpoeAgentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this replace spoe agent o k response has a 5xx status code
func (o *ReplaceSpoeAgentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this replace spoe agent o k response a status code equal to that given
func (o *ReplaceSpoeAgentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the replace spoe agent o k response
func (o *ReplaceSpoeAgentOK) Code() int {
	return 200
}

func (o *ReplaceSpoeAgentOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe/spoe_agents/{name}][%d] replaceSpoeAgentOK  %+v", 200, o.Payload)
}

func (o *ReplaceSpoeAgentOK) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe/spoe_agents/{name}][%d] replaceSpoeAgentOK  %+v", 200, o.Payload)
}

func (o *ReplaceSpoeAgentOK) GetPayload() *models.SpoeAgent {
	return o.Payload
}

func (o *ReplaceSpoeAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SpoeAgent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceSpoeAgentBadRequest creates a ReplaceSpoeAgentBadRequest with default headers values
func NewReplaceSpoeAgentBadRequest() *ReplaceSpoeAgentBadRequest {
	return &ReplaceSpoeAgentBadRequest{}
}

/*
ReplaceSpoeAgentBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceSpoeAgentBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this replace spoe agent bad request response has a 2xx status code
func (o *ReplaceSpoeAgentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this replace spoe agent bad request response has a 3xx status code
func (o *ReplaceSpoeAgentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace spoe agent bad request response has a 4xx status code
func (o *ReplaceSpoeAgentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this replace spoe agent bad request response has a 5xx status code
func (o *ReplaceSpoeAgentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this replace spoe agent bad request response a status code equal to that given
func (o *ReplaceSpoeAgentBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the replace spoe agent bad request response
func (o *ReplaceSpoeAgentBadRequest) Code() int {
	return 400
}

func (o *ReplaceSpoeAgentBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe/spoe_agents/{name}][%d] replaceSpoeAgentBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceSpoeAgentBadRequest) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe/spoe_agents/{name}][%d] replaceSpoeAgentBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceSpoeAgentBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceSpoeAgentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceSpoeAgentNotFound creates a ReplaceSpoeAgentNotFound with default headers values
func NewReplaceSpoeAgentNotFound() *ReplaceSpoeAgentNotFound {
	return &ReplaceSpoeAgentNotFound{}
}

/*
ReplaceSpoeAgentNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type ReplaceSpoeAgentNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this replace spoe agent not found response has a 2xx status code
func (o *ReplaceSpoeAgentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this replace spoe agent not found response has a 3xx status code
func (o *ReplaceSpoeAgentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace spoe agent not found response has a 4xx status code
func (o *ReplaceSpoeAgentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this replace spoe agent not found response has a 5xx status code
func (o *ReplaceSpoeAgentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this replace spoe agent not found response a status code equal to that given
func (o *ReplaceSpoeAgentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the replace spoe agent not found response
func (o *ReplaceSpoeAgentNotFound) Code() int {
	return 404
}

func (o *ReplaceSpoeAgentNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe/spoe_agents/{name}][%d] replaceSpoeAgentNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceSpoeAgentNotFound) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe/spoe_agents/{name}][%d] replaceSpoeAgentNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceSpoeAgentNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceSpoeAgentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceSpoeAgentDefault creates a ReplaceSpoeAgentDefault with default headers values
func NewReplaceSpoeAgentDefault(code int) *ReplaceSpoeAgentDefault {
	return &ReplaceSpoeAgentDefault{
		_statusCode: code,
	}
}

/*
ReplaceSpoeAgentDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceSpoeAgentDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this replace spoe agent default response has a 2xx status code
func (o *ReplaceSpoeAgentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this replace spoe agent default response has a 3xx status code
func (o *ReplaceSpoeAgentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this replace spoe agent default response has a 4xx status code
func (o *ReplaceSpoeAgentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this replace spoe agent default response has a 5xx status code
func (o *ReplaceSpoeAgentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this replace spoe agent default response a status code equal to that given
func (o *ReplaceSpoeAgentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the replace spoe agent default response
func (o *ReplaceSpoeAgentDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceSpoeAgentDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe/spoe_agents/{name}][%d] replaceSpoeAgent default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceSpoeAgentDefault) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe/spoe_agents/{name}][%d] replaceSpoeAgent default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceSpoeAgentDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceSpoeAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
