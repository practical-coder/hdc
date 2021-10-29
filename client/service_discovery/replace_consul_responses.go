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

// ReplaceConsulReader is a Reader for the ReplaceConsul structure.
type ReplaceConsulReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceConsulReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceConsulOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceConsulBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceConsulNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceConsulDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceConsulOK creates a ReplaceConsulOK with default headers values
func NewReplaceConsulOK() *ReplaceConsulOK {
	return &ReplaceConsulOK{}
}

/* ReplaceConsulOK describes a response with status code 200, with default header values.

Consul server replaced
*/
type ReplaceConsulOK struct {
	Payload *models.Consul
}

func (o *ReplaceConsulOK) Error() string {
	return fmt.Sprintf("[PUT /service_discovery/consul/{id}][%d] replaceConsulOK  %+v", 200, o.Payload)
}
func (o *ReplaceConsulOK) GetPayload() *models.Consul {
	return o.Payload
}

func (o *ReplaceConsulOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Consul)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceConsulBadRequest creates a ReplaceConsulBadRequest with default headers values
func NewReplaceConsulBadRequest() *ReplaceConsulBadRequest {
	return &ReplaceConsulBadRequest{}
}

/* ReplaceConsulBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceConsulBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceConsulBadRequest) Error() string {
	return fmt.Sprintf("[PUT /service_discovery/consul/{id}][%d] replaceConsulBadRequest  %+v", 400, o.Payload)
}
func (o *ReplaceConsulBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceConsulBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceConsulNotFound creates a ReplaceConsulNotFound with default headers values
func NewReplaceConsulNotFound() *ReplaceConsulNotFound {
	return &ReplaceConsulNotFound{}
}

/* ReplaceConsulNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type ReplaceConsulNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceConsulNotFound) Error() string {
	return fmt.Sprintf("[PUT /service_discovery/consul/{id}][%d] replaceConsulNotFound  %+v", 404, o.Payload)
}
func (o *ReplaceConsulNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceConsulNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceConsulDefault creates a ReplaceConsulDefault with default headers values
func NewReplaceConsulDefault(code int) *ReplaceConsulDefault {
	return &ReplaceConsulDefault{
		_statusCode: code,
	}
}

/* ReplaceConsulDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceConsulDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the replace consul default response
func (o *ReplaceConsulDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceConsulDefault) Error() string {
	return fmt.Sprintf("[PUT /service_discovery/consul/{id}][%d] replaceConsul default  %+v", o._statusCode, o.Payload)
}
func (o *ReplaceConsulDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceConsulDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
