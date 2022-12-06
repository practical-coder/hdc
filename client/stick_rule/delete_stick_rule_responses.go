// Code generated by go-swagger; DO NOT EDIT.

package stick_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// DeleteStickRuleReader is a Reader for the DeleteStickRule structure.
type DeleteStickRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteStickRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteStickRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteStickRuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteStickRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteStickRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteStickRuleAccepted creates a DeleteStickRuleAccepted with default headers values
func NewDeleteStickRuleAccepted() *DeleteStickRuleAccepted {
	return &DeleteStickRuleAccepted{}
}

/*
DeleteStickRuleAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeleteStickRuleAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

// IsSuccess returns true when this delete stick rule accepted response has a 2xx status code
func (o *DeleteStickRuleAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete stick rule accepted response has a 3xx status code
func (o *DeleteStickRuleAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete stick rule accepted response has a 4xx status code
func (o *DeleteStickRuleAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete stick rule accepted response has a 5xx status code
func (o *DeleteStickRuleAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this delete stick rule accepted response a status code equal to that given
func (o *DeleteStickRuleAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *DeleteStickRuleAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/stick_rules/{index}][%d] deleteStickRuleAccepted ", 202)
}

func (o *DeleteStickRuleAccepted) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/stick_rules/{index}][%d] deleteStickRuleAccepted ", 202)
}

func (o *DeleteStickRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeleteStickRuleNoContent creates a DeleteStickRuleNoContent with default headers values
func NewDeleteStickRuleNoContent() *DeleteStickRuleNoContent {
	return &DeleteStickRuleNoContent{}
}

/*
DeleteStickRuleNoContent describes a response with status code 204, with default header values.

Stick Rule deleted
*/
type DeleteStickRuleNoContent struct {
}

// IsSuccess returns true when this delete stick rule no content response has a 2xx status code
func (o *DeleteStickRuleNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete stick rule no content response has a 3xx status code
func (o *DeleteStickRuleNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete stick rule no content response has a 4xx status code
func (o *DeleteStickRuleNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete stick rule no content response has a 5xx status code
func (o *DeleteStickRuleNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete stick rule no content response a status code equal to that given
func (o *DeleteStickRuleNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteStickRuleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/stick_rules/{index}][%d] deleteStickRuleNoContent ", 204)
}

func (o *DeleteStickRuleNoContent) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/stick_rules/{index}][%d] deleteStickRuleNoContent ", 204)
}

func (o *DeleteStickRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteStickRuleNotFound creates a DeleteStickRuleNotFound with default headers values
func NewDeleteStickRuleNotFound() *DeleteStickRuleNotFound {
	return &DeleteStickRuleNotFound{}
}

/*
DeleteStickRuleNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteStickRuleNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete stick rule not found response has a 2xx status code
func (o *DeleteStickRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete stick rule not found response has a 3xx status code
func (o *DeleteStickRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete stick rule not found response has a 4xx status code
func (o *DeleteStickRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete stick rule not found response has a 5xx status code
func (o *DeleteStickRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete stick rule not found response a status code equal to that given
func (o *DeleteStickRuleNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteStickRuleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/stick_rules/{index}][%d] deleteStickRuleNotFound  %+v", 404, o.Payload)
}

func (o *DeleteStickRuleNotFound) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/stick_rules/{index}][%d] deleteStickRuleNotFound  %+v", 404, o.Payload)
}

func (o *DeleteStickRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteStickRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteStickRuleDefault creates a DeleteStickRuleDefault with default headers values
func NewDeleteStickRuleDefault(code int) *DeleteStickRuleDefault {
	return &DeleteStickRuleDefault{
		_statusCode: code,
	}
}

/*
DeleteStickRuleDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteStickRuleDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete stick rule default response
func (o *DeleteStickRuleDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete stick rule default response has a 2xx status code
func (o *DeleteStickRuleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete stick rule default response has a 3xx status code
func (o *DeleteStickRuleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete stick rule default response has a 4xx status code
func (o *DeleteStickRuleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete stick rule default response has a 5xx status code
func (o *DeleteStickRuleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete stick rule default response a status code equal to that given
func (o *DeleteStickRuleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteStickRuleDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/stick_rules/{index}][%d] deleteStickRule default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteStickRuleDefault) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/stick_rules/{index}][%d] deleteStickRule default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteStickRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteStickRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
