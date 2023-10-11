// Code generated by go-swagger; DO NOT EDIT.

package tcp_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v5/models"
)

// DeleteTCPCheckReader is a Reader for the DeleteTCPCheck structure.
type DeleteTCPCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTCPCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteTCPCheckAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteTCPCheckNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteTCPCheckNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteTCPCheckDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteTCPCheckAccepted creates a DeleteTCPCheckAccepted with default headers values
func NewDeleteTCPCheckAccepted() *DeleteTCPCheckAccepted {
	return &DeleteTCPCheckAccepted{}
}

/*
DeleteTCPCheckAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeleteTCPCheckAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

// IsSuccess returns true when this delete Tcp check accepted response has a 2xx status code
func (o *DeleteTCPCheckAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete Tcp check accepted response has a 3xx status code
func (o *DeleteTCPCheckAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Tcp check accepted response has a 4xx status code
func (o *DeleteTCPCheckAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Tcp check accepted response has a 5xx status code
func (o *DeleteTCPCheckAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Tcp check accepted response a status code equal to that given
func (o *DeleteTCPCheckAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the delete Tcp check accepted response
func (o *DeleteTCPCheckAccepted) Code() int {
	return 202
}

func (o *DeleteTCPCheckAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/tcp_checks/{index}][%d] deleteTcpCheckAccepted ", 202)
}

func (o *DeleteTCPCheckAccepted) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/tcp_checks/{index}][%d] deleteTcpCheckAccepted ", 202)
}

func (o *DeleteTCPCheckAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeleteTCPCheckNoContent creates a DeleteTCPCheckNoContent with default headers values
func NewDeleteTCPCheckNoContent() *DeleteTCPCheckNoContent {
	return &DeleteTCPCheckNoContent{}
}

/*
DeleteTCPCheckNoContent describes a response with status code 204, with default header values.

TCP check deleted
*/
type DeleteTCPCheckNoContent struct {
}

// IsSuccess returns true when this delete Tcp check no content response has a 2xx status code
func (o *DeleteTCPCheckNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete Tcp check no content response has a 3xx status code
func (o *DeleteTCPCheckNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Tcp check no content response has a 4xx status code
func (o *DeleteTCPCheckNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Tcp check no content response has a 5xx status code
func (o *DeleteTCPCheckNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Tcp check no content response a status code equal to that given
func (o *DeleteTCPCheckNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete Tcp check no content response
func (o *DeleteTCPCheckNoContent) Code() int {
	return 204
}

func (o *DeleteTCPCheckNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/tcp_checks/{index}][%d] deleteTcpCheckNoContent ", 204)
}

func (o *DeleteTCPCheckNoContent) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/tcp_checks/{index}][%d] deleteTcpCheckNoContent ", 204)
}

func (o *DeleteTCPCheckNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTCPCheckNotFound creates a DeleteTCPCheckNotFound with default headers values
func NewDeleteTCPCheckNotFound() *DeleteTCPCheckNotFound {
	return &DeleteTCPCheckNotFound{}
}

/*
DeleteTCPCheckNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteTCPCheckNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete Tcp check not found response has a 2xx status code
func (o *DeleteTCPCheckNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Tcp check not found response has a 3xx status code
func (o *DeleteTCPCheckNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Tcp check not found response has a 4xx status code
func (o *DeleteTCPCheckNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Tcp check not found response has a 5xx status code
func (o *DeleteTCPCheckNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Tcp check not found response a status code equal to that given
func (o *DeleteTCPCheckNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete Tcp check not found response
func (o *DeleteTCPCheckNotFound) Code() int {
	return 404
}

func (o *DeleteTCPCheckNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/tcp_checks/{index}][%d] deleteTcpCheckNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTCPCheckNotFound) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/tcp_checks/{index}][%d] deleteTcpCheckNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTCPCheckNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteTCPCheckNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteTCPCheckDefault creates a DeleteTCPCheckDefault with default headers values
func NewDeleteTCPCheckDefault(code int) *DeleteTCPCheckDefault {
	return &DeleteTCPCheckDefault{
		_statusCode: code,
	}
}

/*
DeleteTCPCheckDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteTCPCheckDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete TCP check default response has a 2xx status code
func (o *DeleteTCPCheckDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete TCP check default response has a 3xx status code
func (o *DeleteTCPCheckDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete TCP check default response has a 4xx status code
func (o *DeleteTCPCheckDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete TCP check default response has a 5xx status code
func (o *DeleteTCPCheckDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete TCP check default response a status code equal to that given
func (o *DeleteTCPCheckDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete TCP check default response
func (o *DeleteTCPCheckDefault) Code() int {
	return o._statusCode
}

func (o *DeleteTCPCheckDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/tcp_checks/{index}][%d] deleteTCPCheck default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTCPCheckDefault) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/tcp_checks/{index}][%d] deleteTCPCheck default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTCPCheckDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteTCPCheckDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
