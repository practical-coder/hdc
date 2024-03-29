// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v5/models"
)

// DeleteRuntimeServerReader is a Reader for the DeleteRuntimeServer structure.
type DeleteRuntimeServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRuntimeServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteRuntimeServerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRuntimeServerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRuntimeServerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteRuntimeServerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteRuntimeServerNoContent creates a DeleteRuntimeServerNoContent with default headers values
func NewDeleteRuntimeServerNoContent() *DeleteRuntimeServerNoContent {
	return &DeleteRuntimeServerNoContent{}
}

/*
DeleteRuntimeServerNoContent describes a response with status code 204, with default header values.

Server deleted
*/
type DeleteRuntimeServerNoContent struct {
}

// IsSuccess returns true when this delete runtime server no content response has a 2xx status code
func (o *DeleteRuntimeServerNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete runtime server no content response has a 3xx status code
func (o *DeleteRuntimeServerNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete runtime server no content response has a 4xx status code
func (o *DeleteRuntimeServerNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete runtime server no content response has a 5xx status code
func (o *DeleteRuntimeServerNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete runtime server no content response a status code equal to that given
func (o *DeleteRuntimeServerNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete runtime server no content response
func (o *DeleteRuntimeServerNoContent) Code() int {
	return 204
}

func (o *DeleteRuntimeServerNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/runtime/servers/{name}][%d] deleteRuntimeServerNoContent ", 204)
}

func (o *DeleteRuntimeServerNoContent) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/runtime/servers/{name}][%d] deleteRuntimeServerNoContent ", 204)
}

func (o *DeleteRuntimeServerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRuntimeServerBadRequest creates a DeleteRuntimeServerBadRequest with default headers values
func NewDeleteRuntimeServerBadRequest() *DeleteRuntimeServerBadRequest {
	return &DeleteRuntimeServerBadRequest{}
}

/*
DeleteRuntimeServerBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteRuntimeServerBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete runtime server bad request response has a 2xx status code
func (o *DeleteRuntimeServerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete runtime server bad request response has a 3xx status code
func (o *DeleteRuntimeServerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete runtime server bad request response has a 4xx status code
func (o *DeleteRuntimeServerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete runtime server bad request response has a 5xx status code
func (o *DeleteRuntimeServerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete runtime server bad request response a status code equal to that given
func (o *DeleteRuntimeServerBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete runtime server bad request response
func (o *DeleteRuntimeServerBadRequest) Code() int {
	return 400
}

func (o *DeleteRuntimeServerBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/runtime/servers/{name}][%d] deleteRuntimeServerBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRuntimeServerBadRequest) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/runtime/servers/{name}][%d] deleteRuntimeServerBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRuntimeServerBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteRuntimeServerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRuntimeServerNotFound creates a DeleteRuntimeServerNotFound with default headers values
func NewDeleteRuntimeServerNotFound() *DeleteRuntimeServerNotFound {
	return &DeleteRuntimeServerNotFound{}
}

/*
DeleteRuntimeServerNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteRuntimeServerNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete runtime server not found response has a 2xx status code
func (o *DeleteRuntimeServerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete runtime server not found response has a 3xx status code
func (o *DeleteRuntimeServerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete runtime server not found response has a 4xx status code
func (o *DeleteRuntimeServerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete runtime server not found response has a 5xx status code
func (o *DeleteRuntimeServerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete runtime server not found response a status code equal to that given
func (o *DeleteRuntimeServerNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete runtime server not found response
func (o *DeleteRuntimeServerNotFound) Code() int {
	return 404
}

func (o *DeleteRuntimeServerNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/runtime/servers/{name}][%d] deleteRuntimeServerNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRuntimeServerNotFound) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/runtime/servers/{name}][%d] deleteRuntimeServerNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRuntimeServerNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteRuntimeServerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRuntimeServerDefault creates a DeleteRuntimeServerDefault with default headers values
func NewDeleteRuntimeServerDefault(code int) *DeleteRuntimeServerDefault {
	return &DeleteRuntimeServerDefault{
		_statusCode: code,
	}
}

/*
DeleteRuntimeServerDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteRuntimeServerDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete runtime server default response has a 2xx status code
func (o *DeleteRuntimeServerDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete runtime server default response has a 3xx status code
func (o *DeleteRuntimeServerDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete runtime server default response has a 4xx status code
func (o *DeleteRuntimeServerDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete runtime server default response has a 5xx status code
func (o *DeleteRuntimeServerDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete runtime server default response a status code equal to that given
func (o *DeleteRuntimeServerDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete runtime server default response
func (o *DeleteRuntimeServerDefault) Code() int {
	return o._statusCode
}

func (o *DeleteRuntimeServerDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/runtime/servers/{name}][%d] deleteRuntimeServer default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteRuntimeServerDefault) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/runtime/servers/{name}][%d] deleteRuntimeServer default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteRuntimeServerDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteRuntimeServerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
