// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// GetRuntimeServerReader is a Reader for the GetRuntimeServer structure.
type GetRuntimeServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRuntimeServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRuntimeServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRuntimeServerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetRuntimeServerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRuntimeServerOK creates a GetRuntimeServerOK with default headers values
func NewGetRuntimeServerOK() *GetRuntimeServerOK {
	return &GetRuntimeServerOK{}
}

/*
GetRuntimeServerOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetRuntimeServerOK struct {
	Payload *models.RuntimeServer
}

// IsSuccess returns true when this get runtime server o k response has a 2xx status code
func (o *GetRuntimeServerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get runtime server o k response has a 3xx status code
func (o *GetRuntimeServerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runtime server o k response has a 4xx status code
func (o *GetRuntimeServerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get runtime server o k response has a 5xx status code
func (o *GetRuntimeServerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get runtime server o k response a status code equal to that given
func (o *GetRuntimeServerOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRuntimeServerOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/servers/{name}][%d] getRuntimeServerOK  %+v", 200, o.Payload)
}

func (o *GetRuntimeServerOK) String() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/servers/{name}][%d] getRuntimeServerOK  %+v", 200, o.Payload)
}

func (o *GetRuntimeServerOK) GetPayload() *models.RuntimeServer {
	return o.Payload
}

func (o *GetRuntimeServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeServer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRuntimeServerNotFound creates a GetRuntimeServerNotFound with default headers values
func NewGetRuntimeServerNotFound() *GetRuntimeServerNotFound {
	return &GetRuntimeServerNotFound{}
}

/*
GetRuntimeServerNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetRuntimeServerNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get runtime server not found response has a 2xx status code
func (o *GetRuntimeServerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runtime server not found response has a 3xx status code
func (o *GetRuntimeServerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runtime server not found response has a 4xx status code
func (o *GetRuntimeServerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get runtime server not found response has a 5xx status code
func (o *GetRuntimeServerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get runtime server not found response a status code equal to that given
func (o *GetRuntimeServerNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRuntimeServerNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/servers/{name}][%d] getRuntimeServerNotFound  %+v", 404, o.Payload)
}

func (o *GetRuntimeServerNotFound) String() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/servers/{name}][%d] getRuntimeServerNotFound  %+v", 404, o.Payload)
}

func (o *GetRuntimeServerNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRuntimeServerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRuntimeServerDefault creates a GetRuntimeServerDefault with default headers values
func NewGetRuntimeServerDefault(code int) *GetRuntimeServerDefault {
	return &GetRuntimeServerDefault{
		_statusCode: code,
	}
}

/*
GetRuntimeServerDefault describes a response with status code -1, with default header values.

General Error
*/
type GetRuntimeServerDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get runtime server default response
func (o *GetRuntimeServerDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get runtime server default response has a 2xx status code
func (o *GetRuntimeServerDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get runtime server default response has a 3xx status code
func (o *GetRuntimeServerDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get runtime server default response has a 4xx status code
func (o *GetRuntimeServerDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get runtime server default response has a 5xx status code
func (o *GetRuntimeServerDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get runtime server default response a status code equal to that given
func (o *GetRuntimeServerDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetRuntimeServerDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/servers/{name}][%d] getRuntimeServer default  %+v", o._statusCode, o.Payload)
}

func (o *GetRuntimeServerDefault) String() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/servers/{name}][%d] getRuntimeServer default  %+v", o._statusCode, o.Payload)
}

func (o *GetRuntimeServerDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRuntimeServerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
