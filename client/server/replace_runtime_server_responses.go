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

// ReplaceRuntimeServerReader is a Reader for the ReplaceRuntimeServer structure.
type ReplaceRuntimeServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceRuntimeServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceRuntimeServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceRuntimeServerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceRuntimeServerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceRuntimeServerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceRuntimeServerOK creates a ReplaceRuntimeServerOK with default headers values
func NewReplaceRuntimeServerOK() *ReplaceRuntimeServerOK {
	return &ReplaceRuntimeServerOK{}
}

/*
ReplaceRuntimeServerOK describes a response with status code 200, with default header values.

Server transient settings replaced
*/
type ReplaceRuntimeServerOK struct {
	Payload *models.RuntimeServer
}

// IsSuccess returns true when this replace runtime server o k response has a 2xx status code
func (o *ReplaceRuntimeServerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this replace runtime server o k response has a 3xx status code
func (o *ReplaceRuntimeServerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace runtime server o k response has a 4xx status code
func (o *ReplaceRuntimeServerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this replace runtime server o k response has a 5xx status code
func (o *ReplaceRuntimeServerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this replace runtime server o k response a status code equal to that given
func (o *ReplaceRuntimeServerOK) IsCode(code int) bool {
	return code == 200
}

func (o *ReplaceRuntimeServerOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/servers/{name}][%d] replaceRuntimeServerOK  %+v", 200, o.Payload)
}

func (o *ReplaceRuntimeServerOK) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/servers/{name}][%d] replaceRuntimeServerOK  %+v", 200, o.Payload)
}

func (o *ReplaceRuntimeServerOK) GetPayload() *models.RuntimeServer {
	return o.Payload
}

func (o *ReplaceRuntimeServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeServer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceRuntimeServerBadRequest creates a ReplaceRuntimeServerBadRequest with default headers values
func NewReplaceRuntimeServerBadRequest() *ReplaceRuntimeServerBadRequest {
	return &ReplaceRuntimeServerBadRequest{}
}

/*
ReplaceRuntimeServerBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceRuntimeServerBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this replace runtime server bad request response has a 2xx status code
func (o *ReplaceRuntimeServerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this replace runtime server bad request response has a 3xx status code
func (o *ReplaceRuntimeServerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace runtime server bad request response has a 4xx status code
func (o *ReplaceRuntimeServerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this replace runtime server bad request response has a 5xx status code
func (o *ReplaceRuntimeServerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this replace runtime server bad request response a status code equal to that given
func (o *ReplaceRuntimeServerBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ReplaceRuntimeServerBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/servers/{name}][%d] replaceRuntimeServerBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceRuntimeServerBadRequest) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/servers/{name}][%d] replaceRuntimeServerBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceRuntimeServerBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceRuntimeServerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceRuntimeServerNotFound creates a ReplaceRuntimeServerNotFound with default headers values
func NewReplaceRuntimeServerNotFound() *ReplaceRuntimeServerNotFound {
	return &ReplaceRuntimeServerNotFound{}
}

/*
ReplaceRuntimeServerNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type ReplaceRuntimeServerNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this replace runtime server not found response has a 2xx status code
func (o *ReplaceRuntimeServerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this replace runtime server not found response has a 3xx status code
func (o *ReplaceRuntimeServerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace runtime server not found response has a 4xx status code
func (o *ReplaceRuntimeServerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this replace runtime server not found response has a 5xx status code
func (o *ReplaceRuntimeServerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this replace runtime server not found response a status code equal to that given
func (o *ReplaceRuntimeServerNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ReplaceRuntimeServerNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/servers/{name}][%d] replaceRuntimeServerNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceRuntimeServerNotFound) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/servers/{name}][%d] replaceRuntimeServerNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceRuntimeServerNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceRuntimeServerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceRuntimeServerDefault creates a ReplaceRuntimeServerDefault with default headers values
func NewReplaceRuntimeServerDefault(code int) *ReplaceRuntimeServerDefault {
	return &ReplaceRuntimeServerDefault{
		_statusCode: code,
	}
}

/*
ReplaceRuntimeServerDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceRuntimeServerDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the replace runtime server default response
func (o *ReplaceRuntimeServerDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this replace runtime server default response has a 2xx status code
func (o *ReplaceRuntimeServerDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this replace runtime server default response has a 3xx status code
func (o *ReplaceRuntimeServerDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this replace runtime server default response has a 4xx status code
func (o *ReplaceRuntimeServerDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this replace runtime server default response has a 5xx status code
func (o *ReplaceRuntimeServerDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this replace runtime server default response a status code equal to that given
func (o *ReplaceRuntimeServerDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ReplaceRuntimeServerDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/servers/{name}][%d] replaceRuntimeServer default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceRuntimeServerDefault) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/servers/{name}][%d] replaceRuntimeServer default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceRuntimeServerDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceRuntimeServerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
