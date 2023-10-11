// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v5/models"
)

// DeleteStorageMapReader is a Reader for the DeleteStorageMap structure.
type DeleteStorageMapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteStorageMapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteStorageMapNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteStorageMapNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteStorageMapDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteStorageMapNoContent creates a DeleteStorageMapNoContent with default headers values
func NewDeleteStorageMapNoContent() *DeleteStorageMapNoContent {
	return &DeleteStorageMapNoContent{}
}

/*
DeleteStorageMapNoContent describes a response with status code 204, with default header values.

Map file deleted
*/
type DeleteStorageMapNoContent struct {
}

// IsSuccess returns true when this delete storage map no content response has a 2xx status code
func (o *DeleteStorageMapNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete storage map no content response has a 3xx status code
func (o *DeleteStorageMapNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete storage map no content response has a 4xx status code
func (o *DeleteStorageMapNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete storage map no content response has a 5xx status code
func (o *DeleteStorageMapNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete storage map no content response a status code equal to that given
func (o *DeleteStorageMapNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete storage map no content response
func (o *DeleteStorageMapNoContent) Code() int {
	return 204
}

func (o *DeleteStorageMapNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/storage/maps/{name}][%d] deleteStorageMapNoContent ", 204)
}

func (o *DeleteStorageMapNoContent) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/storage/maps/{name}][%d] deleteStorageMapNoContent ", 204)
}

func (o *DeleteStorageMapNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteStorageMapNotFound creates a DeleteStorageMapNotFound with default headers values
func NewDeleteStorageMapNotFound() *DeleteStorageMapNotFound {
	return &DeleteStorageMapNotFound{}
}

/*
DeleteStorageMapNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteStorageMapNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete storage map not found response has a 2xx status code
func (o *DeleteStorageMapNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete storage map not found response has a 3xx status code
func (o *DeleteStorageMapNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete storage map not found response has a 4xx status code
func (o *DeleteStorageMapNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete storage map not found response has a 5xx status code
func (o *DeleteStorageMapNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete storage map not found response a status code equal to that given
func (o *DeleteStorageMapNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete storage map not found response
func (o *DeleteStorageMapNotFound) Code() int {
	return 404
}

func (o *DeleteStorageMapNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/storage/maps/{name}][%d] deleteStorageMapNotFound  %+v", 404, o.Payload)
}

func (o *DeleteStorageMapNotFound) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/storage/maps/{name}][%d] deleteStorageMapNotFound  %+v", 404, o.Payload)
}

func (o *DeleteStorageMapNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteStorageMapNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteStorageMapDefault creates a DeleteStorageMapDefault with default headers values
func NewDeleteStorageMapDefault(code int) *DeleteStorageMapDefault {
	return &DeleteStorageMapDefault{
		_statusCode: code,
	}
}

/*
DeleteStorageMapDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteStorageMapDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete storage map default response has a 2xx status code
func (o *DeleteStorageMapDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete storage map default response has a 3xx status code
func (o *DeleteStorageMapDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete storage map default response has a 4xx status code
func (o *DeleteStorageMapDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete storage map default response has a 5xx status code
func (o *DeleteStorageMapDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete storage map default response a status code equal to that given
func (o *DeleteStorageMapDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete storage map default response
func (o *DeleteStorageMapDefault) Code() int {
	return o._statusCode
}

func (o *DeleteStorageMapDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/storage/maps/{name}][%d] deleteStorageMap default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteStorageMapDefault) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/storage/maps/{name}][%d] deleteStorageMap default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteStorageMapDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteStorageMapDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
