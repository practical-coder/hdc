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

// DeleteSpoeAgentReader is a Reader for the DeleteSpoeAgent structure.
type DeleteSpoeAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSpoeAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteSpoeAgentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteSpoeAgentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteSpoeAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSpoeAgentNoContent creates a DeleteSpoeAgentNoContent with default headers values
func NewDeleteSpoeAgentNoContent() *DeleteSpoeAgentNoContent {
	return &DeleteSpoeAgentNoContent{}
}

/*
DeleteSpoeAgentNoContent describes a response with status code 204, with default header values.

Spoe agent deleted
*/
type DeleteSpoeAgentNoContent struct {
}

// IsSuccess returns true when this delete spoe agent no content response has a 2xx status code
func (o *DeleteSpoeAgentNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete spoe agent no content response has a 3xx status code
func (o *DeleteSpoeAgentNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete spoe agent no content response has a 4xx status code
func (o *DeleteSpoeAgentNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete spoe agent no content response has a 5xx status code
func (o *DeleteSpoeAgentNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete spoe agent no content response a status code equal to that given
func (o *DeleteSpoeAgentNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete spoe agent no content response
func (o *DeleteSpoeAgentNoContent) Code() int {
	return 204
}

func (o *DeleteSpoeAgentNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_agents/{name}][%d] deleteSpoeAgentNoContent ", 204)
}

func (o *DeleteSpoeAgentNoContent) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_agents/{name}][%d] deleteSpoeAgentNoContent ", 204)
}

func (o *DeleteSpoeAgentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSpoeAgentNotFound creates a DeleteSpoeAgentNotFound with default headers values
func NewDeleteSpoeAgentNotFound() *DeleteSpoeAgentNotFound {
	return &DeleteSpoeAgentNotFound{}
}

/*
DeleteSpoeAgentNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteSpoeAgentNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete spoe agent not found response has a 2xx status code
func (o *DeleteSpoeAgentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete spoe agent not found response has a 3xx status code
func (o *DeleteSpoeAgentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete spoe agent not found response has a 4xx status code
func (o *DeleteSpoeAgentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete spoe agent not found response has a 5xx status code
func (o *DeleteSpoeAgentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete spoe agent not found response a status code equal to that given
func (o *DeleteSpoeAgentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete spoe agent not found response
func (o *DeleteSpoeAgentNotFound) Code() int {
	return 404
}

func (o *DeleteSpoeAgentNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_agents/{name}][%d] deleteSpoeAgentNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSpoeAgentNotFound) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_agents/{name}][%d] deleteSpoeAgentNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSpoeAgentNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteSpoeAgentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteSpoeAgentDefault creates a DeleteSpoeAgentDefault with default headers values
func NewDeleteSpoeAgentDefault(code int) *DeleteSpoeAgentDefault {
	return &DeleteSpoeAgentDefault{
		_statusCode: code,
	}
}

/*
DeleteSpoeAgentDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteSpoeAgentDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete spoe agent default response has a 2xx status code
func (o *DeleteSpoeAgentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete spoe agent default response has a 3xx status code
func (o *DeleteSpoeAgentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete spoe agent default response has a 4xx status code
func (o *DeleteSpoeAgentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete spoe agent default response has a 5xx status code
func (o *DeleteSpoeAgentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete spoe agent default response a status code equal to that given
func (o *DeleteSpoeAgentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete spoe agent default response
func (o *DeleteSpoeAgentDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSpoeAgentDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_agents/{name}][%d] deleteSpoeAgent default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSpoeAgentDefault) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_agents/{name}][%d] deleteSpoeAgent default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSpoeAgentDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteSpoeAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
