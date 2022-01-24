// Code generated by go-swagger; DO NOT EDIT.

package maps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v2/models"
)

// ClearRuntimeMapReader is a Reader for the ClearRuntimeMap structure.
type ClearRuntimeMapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClearRuntimeMapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewClearRuntimeMapNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewClearRuntimeMapNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewClearRuntimeMapDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClearRuntimeMapNoContent creates a ClearRuntimeMapNoContent with default headers values
func NewClearRuntimeMapNoContent() *ClearRuntimeMapNoContent {
	return &ClearRuntimeMapNoContent{}
}

/* ClearRuntimeMapNoContent describes a response with status code 204, with default header values.

All map entries deleted
*/
type ClearRuntimeMapNoContent struct {
}

func (o *ClearRuntimeMapNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/runtime/maps/{name}][%d] clearRuntimeMapNoContent ", 204)
}

func (o *ClearRuntimeMapNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClearRuntimeMapNotFound creates a ClearRuntimeMapNotFound with default headers values
func NewClearRuntimeMapNotFound() *ClearRuntimeMapNotFound {
	return &ClearRuntimeMapNotFound{}
}

/* ClearRuntimeMapNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type ClearRuntimeMapNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ClearRuntimeMapNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/runtime/maps/{name}][%d] clearRuntimeMapNotFound  %+v", 404, o.Payload)
}
func (o *ClearRuntimeMapNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ClearRuntimeMapNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewClearRuntimeMapDefault creates a ClearRuntimeMapDefault with default headers values
func NewClearRuntimeMapDefault(code int) *ClearRuntimeMapDefault {
	return &ClearRuntimeMapDefault{
		_statusCode: code,
	}
}

/* ClearRuntimeMapDefault describes a response with status code -1, with default header values.

General Error
*/
type ClearRuntimeMapDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the clear runtime map default response
func (o *ClearRuntimeMapDefault) Code() int {
	return o._statusCode
}

func (o *ClearRuntimeMapDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/runtime/maps/{name}][%d] clearRuntimeMap default  %+v", o._statusCode, o.Payload)
}
func (o *ClearRuntimeMapDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ClearRuntimeMapDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
