// Code generated by go-swagger; DO NOT EDIT.

package filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// DeleteFilterReader is a Reader for the DeleteFilter structure.
type DeleteFilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteFilterAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteFilterNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteFilterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteFilterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteFilterAccepted creates a DeleteFilterAccepted with default headers values
func NewDeleteFilterAccepted() *DeleteFilterAccepted {
	return &DeleteFilterAccepted{}
}

/*
DeleteFilterAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeleteFilterAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

// IsSuccess returns true when this delete filter accepted response has a 2xx status code
func (o *DeleteFilterAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete filter accepted response has a 3xx status code
func (o *DeleteFilterAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete filter accepted response has a 4xx status code
func (o *DeleteFilterAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete filter accepted response has a 5xx status code
func (o *DeleteFilterAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this delete filter accepted response a status code equal to that given
func (o *DeleteFilterAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *DeleteFilterAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/filters/{index}][%d] deleteFilterAccepted ", 202)
}

func (o *DeleteFilterAccepted) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/filters/{index}][%d] deleteFilterAccepted ", 202)
}

func (o *DeleteFilterAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeleteFilterNoContent creates a DeleteFilterNoContent with default headers values
func NewDeleteFilterNoContent() *DeleteFilterNoContent {
	return &DeleteFilterNoContent{}
}

/*
DeleteFilterNoContent describes a response with status code 204, with default header values.

Filter deleted
*/
type DeleteFilterNoContent struct {
}

// IsSuccess returns true when this delete filter no content response has a 2xx status code
func (o *DeleteFilterNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete filter no content response has a 3xx status code
func (o *DeleteFilterNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete filter no content response has a 4xx status code
func (o *DeleteFilterNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete filter no content response has a 5xx status code
func (o *DeleteFilterNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete filter no content response a status code equal to that given
func (o *DeleteFilterNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteFilterNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/filters/{index}][%d] deleteFilterNoContent ", 204)
}

func (o *DeleteFilterNoContent) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/filters/{index}][%d] deleteFilterNoContent ", 204)
}

func (o *DeleteFilterNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFilterNotFound creates a DeleteFilterNotFound with default headers values
func NewDeleteFilterNotFound() *DeleteFilterNotFound {
	return &DeleteFilterNotFound{}
}

/*
DeleteFilterNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteFilterNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete filter not found response has a 2xx status code
func (o *DeleteFilterNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete filter not found response has a 3xx status code
func (o *DeleteFilterNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete filter not found response has a 4xx status code
func (o *DeleteFilterNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete filter not found response has a 5xx status code
func (o *DeleteFilterNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete filter not found response a status code equal to that given
func (o *DeleteFilterNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteFilterNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/filters/{index}][%d] deleteFilterNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFilterNotFound) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/filters/{index}][%d] deleteFilterNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFilterNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteFilterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteFilterDefault creates a DeleteFilterDefault with default headers values
func NewDeleteFilterDefault(code int) *DeleteFilterDefault {
	return &DeleteFilterDefault{
		_statusCode: code,
	}
}

/*
DeleteFilterDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteFilterDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete filter default response
func (o *DeleteFilterDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete filter default response has a 2xx status code
func (o *DeleteFilterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete filter default response has a 3xx status code
func (o *DeleteFilterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete filter default response has a 4xx status code
func (o *DeleteFilterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete filter default response has a 5xx status code
func (o *DeleteFilterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete filter default response a status code equal to that given
func (o *DeleteFilterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteFilterDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/filters/{index}][%d] deleteFilter default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteFilterDefault) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/filters/{index}][%d] deleteFilter default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteFilterDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteFilterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
