// Code generated by go-swagger; DO NOT EDIT.

package tcp_request_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v5/models"
)

// DeleteTCPRequestRuleReader is a Reader for the DeleteTCPRequestRule structure.
type DeleteTCPRequestRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTCPRequestRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteTCPRequestRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteTCPRequestRuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteTCPRequestRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteTCPRequestRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteTCPRequestRuleAccepted creates a DeleteTCPRequestRuleAccepted with default headers values
func NewDeleteTCPRequestRuleAccepted() *DeleteTCPRequestRuleAccepted {
	return &DeleteTCPRequestRuleAccepted{}
}

/*
DeleteTCPRequestRuleAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeleteTCPRequestRuleAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

// IsSuccess returns true when this delete Tcp request rule accepted response has a 2xx status code
func (o *DeleteTCPRequestRuleAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete Tcp request rule accepted response has a 3xx status code
func (o *DeleteTCPRequestRuleAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Tcp request rule accepted response has a 4xx status code
func (o *DeleteTCPRequestRuleAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Tcp request rule accepted response has a 5xx status code
func (o *DeleteTCPRequestRuleAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Tcp request rule accepted response a status code equal to that given
func (o *DeleteTCPRequestRuleAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the delete Tcp request rule accepted response
func (o *DeleteTCPRequestRuleAccepted) Code() int {
	return 202
}

func (o *DeleteTCPRequestRuleAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/tcp_request_rules/{index}][%d] deleteTcpRequestRuleAccepted ", 202)
}

func (o *DeleteTCPRequestRuleAccepted) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/tcp_request_rules/{index}][%d] deleteTcpRequestRuleAccepted ", 202)
}

func (o *DeleteTCPRequestRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeleteTCPRequestRuleNoContent creates a DeleteTCPRequestRuleNoContent with default headers values
func NewDeleteTCPRequestRuleNoContent() *DeleteTCPRequestRuleNoContent {
	return &DeleteTCPRequestRuleNoContent{}
}

/*
DeleteTCPRequestRuleNoContent describes a response with status code 204, with default header values.

TCP Request Rule deleted
*/
type DeleteTCPRequestRuleNoContent struct {
}

// IsSuccess returns true when this delete Tcp request rule no content response has a 2xx status code
func (o *DeleteTCPRequestRuleNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete Tcp request rule no content response has a 3xx status code
func (o *DeleteTCPRequestRuleNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Tcp request rule no content response has a 4xx status code
func (o *DeleteTCPRequestRuleNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Tcp request rule no content response has a 5xx status code
func (o *DeleteTCPRequestRuleNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Tcp request rule no content response a status code equal to that given
func (o *DeleteTCPRequestRuleNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete Tcp request rule no content response
func (o *DeleteTCPRequestRuleNoContent) Code() int {
	return 204
}

func (o *DeleteTCPRequestRuleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/tcp_request_rules/{index}][%d] deleteTcpRequestRuleNoContent ", 204)
}

func (o *DeleteTCPRequestRuleNoContent) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/tcp_request_rules/{index}][%d] deleteTcpRequestRuleNoContent ", 204)
}

func (o *DeleteTCPRequestRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTCPRequestRuleNotFound creates a DeleteTCPRequestRuleNotFound with default headers values
func NewDeleteTCPRequestRuleNotFound() *DeleteTCPRequestRuleNotFound {
	return &DeleteTCPRequestRuleNotFound{}
}

/*
DeleteTCPRequestRuleNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteTCPRequestRuleNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete Tcp request rule not found response has a 2xx status code
func (o *DeleteTCPRequestRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Tcp request rule not found response has a 3xx status code
func (o *DeleteTCPRequestRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Tcp request rule not found response has a 4xx status code
func (o *DeleteTCPRequestRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Tcp request rule not found response has a 5xx status code
func (o *DeleteTCPRequestRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Tcp request rule not found response a status code equal to that given
func (o *DeleteTCPRequestRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete Tcp request rule not found response
func (o *DeleteTCPRequestRuleNotFound) Code() int {
	return 404
}

func (o *DeleteTCPRequestRuleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/tcp_request_rules/{index}][%d] deleteTcpRequestRuleNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTCPRequestRuleNotFound) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/tcp_request_rules/{index}][%d] deleteTcpRequestRuleNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTCPRequestRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteTCPRequestRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteTCPRequestRuleDefault creates a DeleteTCPRequestRuleDefault with default headers values
func NewDeleteTCPRequestRuleDefault(code int) *DeleteTCPRequestRuleDefault {
	return &DeleteTCPRequestRuleDefault{
		_statusCode: code,
	}
}

/*
DeleteTCPRequestRuleDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteTCPRequestRuleDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete TCP request rule default response has a 2xx status code
func (o *DeleteTCPRequestRuleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete TCP request rule default response has a 3xx status code
func (o *DeleteTCPRequestRuleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete TCP request rule default response has a 4xx status code
func (o *DeleteTCPRequestRuleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete TCP request rule default response has a 5xx status code
func (o *DeleteTCPRequestRuleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete TCP request rule default response a status code equal to that given
func (o *DeleteTCPRequestRuleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete TCP request rule default response
func (o *DeleteTCPRequestRuleDefault) Code() int {
	return o._statusCode
}

func (o *DeleteTCPRequestRuleDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/tcp_request_rules/{index}][%d] deleteTCPRequestRule default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTCPRequestRuleDefault) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/tcp_request_rules/{index}][%d] deleteTCPRequestRule default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTCPRequestRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteTCPRequestRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
