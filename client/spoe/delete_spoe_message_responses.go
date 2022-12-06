// Code generated by go-swagger; DO NOT EDIT.

package spoe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// DeleteSpoeMessageReader is a Reader for the DeleteSpoeMessage structure.
type DeleteSpoeMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSpoeMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteSpoeMessageNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteSpoeMessageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteSpoeMessageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSpoeMessageNoContent creates a DeleteSpoeMessageNoContent with default headers values
func NewDeleteSpoeMessageNoContent() *DeleteSpoeMessageNoContent {
	return &DeleteSpoeMessageNoContent{}
}

/*
DeleteSpoeMessageNoContent describes a response with status code 204, with default header values.

Spoe message deleted
*/
type DeleteSpoeMessageNoContent struct {
}

// IsSuccess returns true when this delete spoe message no content response has a 2xx status code
func (o *DeleteSpoeMessageNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete spoe message no content response has a 3xx status code
func (o *DeleteSpoeMessageNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete spoe message no content response has a 4xx status code
func (o *DeleteSpoeMessageNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete spoe message no content response has a 5xx status code
func (o *DeleteSpoeMessageNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete spoe message no content response a status code equal to that given
func (o *DeleteSpoeMessageNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteSpoeMessageNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_messages/{name}][%d] deleteSpoeMessageNoContent ", 204)
}

func (o *DeleteSpoeMessageNoContent) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_messages/{name}][%d] deleteSpoeMessageNoContent ", 204)
}

func (o *DeleteSpoeMessageNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSpoeMessageNotFound creates a DeleteSpoeMessageNotFound with default headers values
func NewDeleteSpoeMessageNotFound() *DeleteSpoeMessageNotFound {
	return &DeleteSpoeMessageNotFound{}
}

/*
DeleteSpoeMessageNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteSpoeMessageNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete spoe message not found response has a 2xx status code
func (o *DeleteSpoeMessageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete spoe message not found response has a 3xx status code
func (o *DeleteSpoeMessageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete spoe message not found response has a 4xx status code
func (o *DeleteSpoeMessageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete spoe message not found response has a 5xx status code
func (o *DeleteSpoeMessageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete spoe message not found response a status code equal to that given
func (o *DeleteSpoeMessageNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteSpoeMessageNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_messages/{name}][%d] deleteSpoeMessageNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSpoeMessageNotFound) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_messages/{name}][%d] deleteSpoeMessageNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSpoeMessageNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteSpoeMessageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteSpoeMessageDefault creates a DeleteSpoeMessageDefault with default headers values
func NewDeleteSpoeMessageDefault(code int) *DeleteSpoeMessageDefault {
	return &DeleteSpoeMessageDefault{
		_statusCode: code,
	}
}

/*
DeleteSpoeMessageDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteSpoeMessageDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete spoe message default response
func (o *DeleteSpoeMessageDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete spoe message default response has a 2xx status code
func (o *DeleteSpoeMessageDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete spoe message default response has a 3xx status code
func (o *DeleteSpoeMessageDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete spoe message default response has a 4xx status code
func (o *DeleteSpoeMessageDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete spoe message default response has a 5xx status code
func (o *DeleteSpoeMessageDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete spoe message default response a status code equal to that given
func (o *DeleteSpoeMessageDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteSpoeMessageDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_messages/{name}][%d] deleteSpoeMessage default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSpoeMessageDefault) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_messages/{name}][%d] deleteSpoeMessage default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSpoeMessageDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteSpoeMessageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
