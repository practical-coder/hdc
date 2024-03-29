// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v5/models"
)

// DeleteGroupReader is a Reader for the DeleteGroup structure.
type DeleteGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteGroupAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteGroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteGroupAccepted creates a DeleteGroupAccepted with default headers values
func NewDeleteGroupAccepted() *DeleteGroupAccepted {
	return &DeleteGroupAccepted{}
}

/*
DeleteGroupAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeleteGroupAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

// IsSuccess returns true when this delete group accepted response has a 2xx status code
func (o *DeleteGroupAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete group accepted response has a 3xx status code
func (o *DeleteGroupAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group accepted response has a 4xx status code
func (o *DeleteGroupAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete group accepted response has a 5xx status code
func (o *DeleteGroupAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this delete group accepted response a status code equal to that given
func (o *DeleteGroupAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the delete group accepted response
func (o *DeleteGroupAccepted) Code() int {
	return 202
}

func (o *DeleteGroupAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/groups/{name}][%d] deleteGroupAccepted ", 202)
}

func (o *DeleteGroupAccepted) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/groups/{name}][%d] deleteGroupAccepted ", 202)
}

func (o *DeleteGroupAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeleteGroupNoContent creates a DeleteGroupNoContent with default headers values
func NewDeleteGroupNoContent() *DeleteGroupNoContent {
	return &DeleteGroupNoContent{}
}

/*
DeleteGroupNoContent describes a response with status code 204, with default header values.

Group deleted
*/
type DeleteGroupNoContent struct {
}

// IsSuccess returns true when this delete group no content response has a 2xx status code
func (o *DeleteGroupNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete group no content response has a 3xx status code
func (o *DeleteGroupNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group no content response has a 4xx status code
func (o *DeleteGroupNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete group no content response has a 5xx status code
func (o *DeleteGroupNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete group no content response a status code equal to that given
func (o *DeleteGroupNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete group no content response
func (o *DeleteGroupNoContent) Code() int {
	return 204
}

func (o *DeleteGroupNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/groups/{name}][%d] deleteGroupNoContent ", 204)
}

func (o *DeleteGroupNoContent) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/groups/{name}][%d] deleteGroupNoContent ", 204)
}

func (o *DeleteGroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteGroupNotFound creates a DeleteGroupNotFound with default headers values
func NewDeleteGroupNotFound() *DeleteGroupNotFound {
	return &DeleteGroupNotFound{}
}

/*
DeleteGroupNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteGroupNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete group not found response has a 2xx status code
func (o *DeleteGroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete group not found response has a 3xx status code
func (o *DeleteGroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group not found response has a 4xx status code
func (o *DeleteGroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete group not found response has a 5xx status code
func (o *DeleteGroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete group not found response a status code equal to that given
func (o *DeleteGroupNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete group not found response
func (o *DeleteGroupNotFound) Code() int {
	return 404
}

func (o *DeleteGroupNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/groups/{name}][%d] deleteGroupNotFound  %+v", 404, o.Payload)
}

func (o *DeleteGroupNotFound) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/groups/{name}][%d] deleteGroupNotFound  %+v", 404, o.Payload)
}

func (o *DeleteGroupNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteGroupDefault creates a DeleteGroupDefault with default headers values
func NewDeleteGroupDefault(code int) *DeleteGroupDefault {
	return &DeleteGroupDefault{
		_statusCode: code,
	}
}

/*
DeleteGroupDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteGroupDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete group default response has a 2xx status code
func (o *DeleteGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete group default response has a 3xx status code
func (o *DeleteGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete group default response has a 4xx status code
func (o *DeleteGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete group default response has a 5xx status code
func (o *DeleteGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete group default response a status code equal to that given
func (o *DeleteGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete group default response
func (o *DeleteGroupDefault) Code() int {
	return o._statusCode
}

func (o *DeleteGroupDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/groups/{name}][%d] deleteGroup default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteGroupDefault) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/groups/{name}][%d] deleteGroup default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteGroupDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
