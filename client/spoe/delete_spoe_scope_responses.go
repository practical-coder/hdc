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

// DeleteSpoeScopeReader is a Reader for the DeleteSpoeScope structure.
type DeleteSpoeScopeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSpoeScopeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteSpoeScopeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteSpoeScopeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteSpoeScopeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSpoeScopeNoContent creates a DeleteSpoeScopeNoContent with default headers values
func NewDeleteSpoeScopeNoContent() *DeleteSpoeScopeNoContent {
	return &DeleteSpoeScopeNoContent{}
}

/*
DeleteSpoeScopeNoContent describes a response with status code 204, with default header values.

Spoe scope deleted
*/
type DeleteSpoeScopeNoContent struct {
}

// IsSuccess returns true when this delete spoe scope no content response has a 2xx status code
func (o *DeleteSpoeScopeNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete spoe scope no content response has a 3xx status code
func (o *DeleteSpoeScopeNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete spoe scope no content response has a 4xx status code
func (o *DeleteSpoeScopeNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete spoe scope no content response has a 5xx status code
func (o *DeleteSpoeScopeNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete spoe scope no content response a status code equal to that given
func (o *DeleteSpoeScopeNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete spoe scope no content response
func (o *DeleteSpoeScopeNoContent) Code() int {
	return 204
}

func (o *DeleteSpoeScopeNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_scopes/{name}][%d] deleteSpoeScopeNoContent ", 204)
}

func (o *DeleteSpoeScopeNoContent) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_scopes/{name}][%d] deleteSpoeScopeNoContent ", 204)
}

func (o *DeleteSpoeScopeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSpoeScopeNotFound creates a DeleteSpoeScopeNotFound with default headers values
func NewDeleteSpoeScopeNotFound() *DeleteSpoeScopeNotFound {
	return &DeleteSpoeScopeNotFound{}
}

/*
DeleteSpoeScopeNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteSpoeScopeNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete spoe scope not found response has a 2xx status code
func (o *DeleteSpoeScopeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete spoe scope not found response has a 3xx status code
func (o *DeleteSpoeScopeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete spoe scope not found response has a 4xx status code
func (o *DeleteSpoeScopeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete spoe scope not found response has a 5xx status code
func (o *DeleteSpoeScopeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete spoe scope not found response a status code equal to that given
func (o *DeleteSpoeScopeNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete spoe scope not found response
func (o *DeleteSpoeScopeNotFound) Code() int {
	return 404
}

func (o *DeleteSpoeScopeNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_scopes/{name}][%d] deleteSpoeScopeNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSpoeScopeNotFound) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_scopes/{name}][%d] deleteSpoeScopeNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSpoeScopeNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteSpoeScopeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteSpoeScopeDefault creates a DeleteSpoeScopeDefault with default headers values
func NewDeleteSpoeScopeDefault(code int) *DeleteSpoeScopeDefault {
	return &DeleteSpoeScopeDefault{
		_statusCode: code,
	}
}

/*
DeleteSpoeScopeDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteSpoeScopeDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete spoe scope default response has a 2xx status code
func (o *DeleteSpoeScopeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete spoe scope default response has a 3xx status code
func (o *DeleteSpoeScopeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete spoe scope default response has a 4xx status code
func (o *DeleteSpoeScopeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete spoe scope default response has a 5xx status code
func (o *DeleteSpoeScopeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete spoe scope default response a status code equal to that given
func (o *DeleteSpoeScopeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete spoe scope default response
func (o *DeleteSpoeScopeDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSpoeScopeDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_scopes/{name}][%d] deleteSpoeScope default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSpoeScopeDefault) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_scopes/{name}][%d] deleteSpoeScope default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSpoeScopeDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteSpoeScopeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
