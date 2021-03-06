// Code generated by go-swagger; DO NOT EDIT.

package http_request_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
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

/* DeleteHTTPRequestRuleAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeleteHTTPRequestRuleAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteHTTPRequestRuleAccepted) Error() string {
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

/* DeleteHTTPRequestRuleNoContent describes a response with status code 204, with default header values.

HTTP Request Rule deleted
*/
type DeleteHTTPRequestRuleNoContent struct {
}

func (o *DeleteHTTPRequestRuleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/http_request_rules/{index}][%d] deleteHttpRequestRuleNoContent ", 204)
}

func (o *DeleteHTTPRequestRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHTTPRequestRuleNotFound creates a DeleteHTTPRequestRuleNotFound with default headers values
func NewDeleteHTTPRequestRuleNotFound() *DeleteHTTPRequestRuleNotFound {
	return &DeleteHTTPRequestRuleNotFound{}
}

/* DeleteHTTPRequestRuleNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteHTTPRequestRuleNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteHTTPRequestRuleNotFound) Error() string {
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

/* DeleteHTTPRequestRuleDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteHTTPRequestRuleDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete HTTP request rule default response
func (o *DeleteHTTPRequestRuleDefault) Code() int {
	return o._statusCode
}

func (o *DeleteHTTPRequestRuleDefault) Error() string {
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
