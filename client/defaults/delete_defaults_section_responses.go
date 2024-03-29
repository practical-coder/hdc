// Code generated by go-swagger; DO NOT EDIT.

package defaults

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v5/models"
)

// DeleteDefaultsSectionReader is a Reader for the DeleteDefaultsSection structure.
type DeleteDefaultsSectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDefaultsSectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteDefaultsSectionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteDefaultsSectionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteDefaultsSectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteDefaultsSectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteDefaultsSectionAccepted creates a DeleteDefaultsSectionAccepted with default headers values
func NewDeleteDefaultsSectionAccepted() *DeleteDefaultsSectionAccepted {
	return &DeleteDefaultsSectionAccepted{}
}

/*
DeleteDefaultsSectionAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeleteDefaultsSectionAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

// IsSuccess returns true when this delete defaults section accepted response has a 2xx status code
func (o *DeleteDefaultsSectionAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete defaults section accepted response has a 3xx status code
func (o *DeleteDefaultsSectionAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete defaults section accepted response has a 4xx status code
func (o *DeleteDefaultsSectionAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete defaults section accepted response has a 5xx status code
func (o *DeleteDefaultsSectionAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this delete defaults section accepted response a status code equal to that given
func (o *DeleteDefaultsSectionAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the delete defaults section accepted response
func (o *DeleteDefaultsSectionAccepted) Code() int {
	return 202
}

func (o *DeleteDefaultsSectionAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/named_defaults/{name}][%d] deleteDefaultsSectionAccepted ", 202)
}

func (o *DeleteDefaultsSectionAccepted) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/named_defaults/{name}][%d] deleteDefaultsSectionAccepted ", 202)
}

func (o *DeleteDefaultsSectionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeleteDefaultsSectionNoContent creates a DeleteDefaultsSectionNoContent with default headers values
func NewDeleteDefaultsSectionNoContent() *DeleteDefaultsSectionNoContent {
	return &DeleteDefaultsSectionNoContent{}
}

/*
DeleteDefaultsSectionNoContent describes a response with status code 204, with default header values.

Defaults section deleted
*/
type DeleteDefaultsSectionNoContent struct {
}

// IsSuccess returns true when this delete defaults section no content response has a 2xx status code
func (o *DeleteDefaultsSectionNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete defaults section no content response has a 3xx status code
func (o *DeleteDefaultsSectionNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete defaults section no content response has a 4xx status code
func (o *DeleteDefaultsSectionNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete defaults section no content response has a 5xx status code
func (o *DeleteDefaultsSectionNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete defaults section no content response a status code equal to that given
func (o *DeleteDefaultsSectionNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete defaults section no content response
func (o *DeleteDefaultsSectionNoContent) Code() int {
	return 204
}

func (o *DeleteDefaultsSectionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/named_defaults/{name}][%d] deleteDefaultsSectionNoContent ", 204)
}

func (o *DeleteDefaultsSectionNoContent) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/named_defaults/{name}][%d] deleteDefaultsSectionNoContent ", 204)
}

func (o *DeleteDefaultsSectionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDefaultsSectionNotFound creates a DeleteDefaultsSectionNotFound with default headers values
func NewDeleteDefaultsSectionNotFound() *DeleteDefaultsSectionNotFound {
	return &DeleteDefaultsSectionNotFound{}
}

/*
DeleteDefaultsSectionNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteDefaultsSectionNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete defaults section not found response has a 2xx status code
func (o *DeleteDefaultsSectionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete defaults section not found response has a 3xx status code
func (o *DeleteDefaultsSectionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete defaults section not found response has a 4xx status code
func (o *DeleteDefaultsSectionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete defaults section not found response has a 5xx status code
func (o *DeleteDefaultsSectionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete defaults section not found response a status code equal to that given
func (o *DeleteDefaultsSectionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete defaults section not found response
func (o *DeleteDefaultsSectionNotFound) Code() int {
	return 404
}

func (o *DeleteDefaultsSectionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/named_defaults/{name}][%d] deleteDefaultsSectionNotFound  %+v", 404, o.Payload)
}

func (o *DeleteDefaultsSectionNotFound) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/named_defaults/{name}][%d] deleteDefaultsSectionNotFound  %+v", 404, o.Payload)
}

func (o *DeleteDefaultsSectionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDefaultsSectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteDefaultsSectionDefault creates a DeleteDefaultsSectionDefault with default headers values
func NewDeleteDefaultsSectionDefault(code int) *DeleteDefaultsSectionDefault {
	return &DeleteDefaultsSectionDefault{
		_statusCode: code,
	}
}

/*
DeleteDefaultsSectionDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteDefaultsSectionDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete defaults section default response has a 2xx status code
func (o *DeleteDefaultsSectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete defaults section default response has a 3xx status code
func (o *DeleteDefaultsSectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete defaults section default response has a 4xx status code
func (o *DeleteDefaultsSectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete defaults section default response has a 5xx status code
func (o *DeleteDefaultsSectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete defaults section default response a status code equal to that given
func (o *DeleteDefaultsSectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete defaults section default response
func (o *DeleteDefaultsSectionDefault) Code() int {
	return o._statusCode
}

func (o *DeleteDefaultsSectionDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/named_defaults/{name}][%d] deleteDefaultsSection default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDefaultsSectionDefault) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/named_defaults/{name}][%d] deleteDefaultsSection default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDefaultsSectionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDefaultsSectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
