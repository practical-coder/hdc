// Code generated by go-swagger; DO NOT EDIT.

package http_request_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v5/models"
)

// DeleteHTTPRequestRuleReader is a Reader for the DeleteHTTPRequestRule structure.
type DeleteHTTPRequestRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteHTTPRequestRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteHTTPRequestRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteHTTPRequestRuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteHTTPRequestRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteHTTPRequestRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteHTTPRequestRuleAccepted creates a DeleteHTTPRequestRuleAccepted with default headers values
func NewDeleteHTTPRequestRuleAccepted() *DeleteHTTPRequestRuleAccepted {
	return &DeleteHTTPRequestRuleAccepted{}
}

/*
DeleteHTTPRequestRuleAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeleteHTTPRequestRuleAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

// IsSuccess returns true when this delete Http request rule accepted response has a 2xx status code
func (o *DeleteHTTPRequestRuleAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete Http request rule accepted response has a 3xx status code
func (o *DeleteHTTPRequestRuleAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Http request rule accepted response has a 4xx status code
func (o *DeleteHTTPRequestRuleAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Http request rule accepted response has a 5xx status code
func (o *DeleteHTTPRequestRuleAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Http request rule accepted response a status code equal to that given
func (o *DeleteHTTPRequestRuleAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the delete Http request rule accepted response
func (o *DeleteHTTPRequestRuleAccepted) Code() int {
	return 202
}

func (o *DeleteHTTPRequestRuleAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/http_request_rules/{index}][%d] deleteHttpRequestRuleAccepted ", 202)
}

func (o *DeleteHTTPRequestRuleAccepted) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/http_request_rules/{index}][%d] deleteHttpRequestRuleAccepted ", 202)
}

func (o *DeleteHTTPRequestRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeleteHTTPRequestRuleNoContent creates a DeleteHTTPRequestRuleNoContent with default headers values
func NewDeleteHTTPRequestRuleNoContent() *DeleteHTTPRequestRuleNoContent {
	return &DeleteHTTPRequestRuleNoContent{}
}

/*
DeleteHTTPRequestRuleNoContent describes a response with status code 204, with default header values.

HTTP Request Rule deleted
*/
type DeleteHTTPRequestRuleNoContent struct {
}

// IsSuccess returns true when this delete Http request rule no content response has a 2xx status code
func (o *DeleteHTTPRequestRuleNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete Http request rule no content response has a 3xx status code
func (o *DeleteHTTPRequestRuleNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Http request rule no content response has a 4xx status code
func (o *DeleteHTTPRequestRuleNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Http request rule no content response has a 5xx status code
func (o *DeleteHTTPRequestRuleNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Http request rule no content response a status code equal to that given
func (o *DeleteHTTPRequestRuleNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete Http request rule no content response
func (o *DeleteHTTPRequestRuleNoContent) Code() int {
	return 204
}

func (o *DeleteHTTPRequestRuleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/http_request_rules/{index}][%d] deleteHttpRequestRuleNoContent ", 204)
}

func (o *DeleteHTTPRequestRuleNoContent) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/http_request_rules/{index}][%d] deleteHttpRequestRuleNoContent ", 204)
}

func (o *DeleteHTTPRequestRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHTTPRequestRuleNotFound creates a DeleteHTTPRequestRuleNotFound with default headers values
func NewDeleteHTTPRequestRuleNotFound() *DeleteHTTPRequestRuleNotFound {
	return &DeleteHTTPRequestRuleNotFound{}
}

/*
DeleteHTTPRequestRuleNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteHTTPRequestRuleNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete Http request rule not found response has a 2xx status code
func (o *DeleteHTTPRequestRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Http request rule not found response has a 3xx status code
func (o *DeleteHTTPRequestRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Http request rule not found response has a 4xx status code
func (o *DeleteHTTPRequestRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Http request rule not found response has a 5xx status code
func (o *DeleteHTTPRequestRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Http request rule not found response a status code equal to that given
func (o *DeleteHTTPRequestRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete Http request rule not found response
func (o *DeleteHTTPRequestRuleNotFound) Code() int {
	return 404
}

func (o *DeleteHTTPRequestRuleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/http_request_rules/{index}][%d] deleteHttpRequestRuleNotFound  %+v", 404, o.Payload)
}

func (o *DeleteHTTPRequestRuleNotFound) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/http_request_rules/{index}][%d] deleteHttpRequestRuleNotFound  %+v", 404, o.Payload)
}

func (o *DeleteHTTPRequestRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteHTTPRequestRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteHTTPRequestRuleDefault creates a DeleteHTTPRequestRuleDefault with default headers values
func NewDeleteHTTPRequestRuleDefault(code int) *DeleteHTTPRequestRuleDefault {
	return &DeleteHTTPRequestRuleDefault{
		_statusCode: code,
	}
}

/*
DeleteHTTPRequestRuleDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteHTTPRequestRuleDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete HTTP request rule default response has a 2xx status code
func (o *DeleteHTTPRequestRuleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete HTTP request rule default response has a 3xx status code
func (o *DeleteHTTPRequestRuleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete HTTP request rule default response has a 4xx status code
func (o *DeleteHTTPRequestRuleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete HTTP request rule default response has a 5xx status code
func (o *DeleteHTTPRequestRuleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete HTTP request rule default response a status code equal to that given
func (o *DeleteHTTPRequestRuleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete HTTP request rule default response
func (o *DeleteHTTPRequestRuleDefault) Code() int {
	return o._statusCode
}

func (o *DeleteHTTPRequestRuleDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/http_request_rules/{index}][%d] deleteHTTPRequestRule default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteHTTPRequestRuleDefault) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/http_request_rules/{index}][%d] deleteHTTPRequestRule default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteHTTPRequestRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteHTTPRequestRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
