// Code generated by go-swagger; DO NOT EDIT.

package service_discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/models"
)

// DeleteConsulReader is a Reader for the DeleteConsul structure.
type DeleteConsulReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteConsulReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteConsulNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteConsulNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteConsulDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteConsulNoContent creates a DeleteConsulNoContent with default headers values
func NewDeleteConsulNoContent() *DeleteConsulNoContent {
	return &DeleteConsulNoContent{}
}

/* DeleteConsulNoContent describes a response with status code 204, with default header values.

Consul server deleted
*/
type DeleteConsulNoContent struct {
}

func (o *DeleteConsulNoContent) Error() string {
	return fmt.Sprintf("[DELETE /service_discovery/consul/{id}][%d] deleteConsulNoContent ", 204)
}

func (o *DeleteConsulNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteConsulNotFound creates a DeleteConsulNotFound with default headers values
func NewDeleteConsulNotFound() *DeleteConsulNotFound {
	return &DeleteConsulNotFound{}
}

/* DeleteConsulNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteConsulNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteConsulNotFound) Error() string {
	return fmt.Sprintf("[DELETE /service_discovery/consul/{id}][%d] deleteConsulNotFound  %+v", 404, o.Payload)
}
func (o *DeleteConsulNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteConsulNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteConsulDefault creates a DeleteConsulDefault with default headers values
func NewDeleteConsulDefault(code int) *DeleteConsulDefault {
	return &DeleteConsulDefault{
		_statusCode: code,
	}
}

/* DeleteConsulDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteConsulDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete consul default response
func (o *DeleteConsulDefault) Code() int {
	return o._statusCode
}

func (o *DeleteConsulDefault) Error() string {
	return fmt.Sprintf("[DELETE /service_discovery/consul/{id}][%d] deleteConsul default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteConsulDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteConsulDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
