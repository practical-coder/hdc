// Code generated by go-swagger; DO NOT EDIT.

package http_error_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v5/models"
)

// ReplaceHTTPErrorRuleReader is a Reader for the ReplaceHTTPErrorRule structure.
type ReplaceHTTPErrorRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceHTTPErrorRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceHTTPErrorRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceHTTPErrorRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceHTTPErrorRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceHTTPErrorRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceHTTPErrorRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceHTTPErrorRuleOK creates a ReplaceHTTPErrorRuleOK with default headers values
func NewReplaceHTTPErrorRuleOK() *ReplaceHTTPErrorRuleOK {
	return &ReplaceHTTPErrorRuleOK{}
}

/*
ReplaceHTTPErrorRuleOK describes a response with status code 200, with default header values.

HTTP Error Rule replaced
*/
type ReplaceHTTPErrorRuleOK struct {
	Payload *models.HTTPErrorRule
}

// IsSuccess returns true when this replace Http error rule o k response has a 2xx status code
func (o *ReplaceHTTPErrorRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this replace Http error rule o k response has a 3xx status code
func (o *ReplaceHTTPErrorRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace Http error rule o k response has a 4xx status code
func (o *ReplaceHTTPErrorRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this replace Http error rule o k response has a 5xx status code
func (o *ReplaceHTTPErrorRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this replace Http error rule o k response a status code equal to that given
func (o *ReplaceHTTPErrorRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the replace Http error rule o k response
func (o *ReplaceHTTPErrorRuleOK) Code() int {
	return 200
}

func (o *ReplaceHTTPErrorRuleOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/http_error_rules/{index}][%d] replaceHttpErrorRuleOK  %+v", 200, o.Payload)
}

func (o *ReplaceHTTPErrorRuleOK) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/http_error_rules/{index}][%d] replaceHttpErrorRuleOK  %+v", 200, o.Payload)
}

func (o *ReplaceHTTPErrorRuleOK) GetPayload() *models.HTTPErrorRule {
	return o.Payload
}

func (o *ReplaceHTTPErrorRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPErrorRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceHTTPErrorRuleAccepted creates a ReplaceHTTPErrorRuleAccepted with default headers values
func NewReplaceHTTPErrorRuleAccepted() *ReplaceHTTPErrorRuleAccepted {
	return &ReplaceHTTPErrorRuleAccepted{}
}

/*
ReplaceHTTPErrorRuleAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type ReplaceHTTPErrorRuleAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.HTTPErrorRule
}

// IsSuccess returns true when this replace Http error rule accepted response has a 2xx status code
func (o *ReplaceHTTPErrorRuleAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this replace Http error rule accepted response has a 3xx status code
func (o *ReplaceHTTPErrorRuleAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace Http error rule accepted response has a 4xx status code
func (o *ReplaceHTTPErrorRuleAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this replace Http error rule accepted response has a 5xx status code
func (o *ReplaceHTTPErrorRuleAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this replace Http error rule accepted response a status code equal to that given
func (o *ReplaceHTTPErrorRuleAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the replace Http error rule accepted response
func (o *ReplaceHTTPErrorRuleAccepted) Code() int {
	return 202
}

func (o *ReplaceHTTPErrorRuleAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/http_error_rules/{index}][%d] replaceHttpErrorRuleAccepted  %+v", 202, o.Payload)
}

func (o *ReplaceHTTPErrorRuleAccepted) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/http_error_rules/{index}][%d] replaceHttpErrorRuleAccepted  %+v", 202, o.Payload)
}

func (o *ReplaceHTTPErrorRuleAccepted) GetPayload() *models.HTTPErrorRule {
	return o.Payload
}

func (o *ReplaceHTTPErrorRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.HTTPErrorRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceHTTPErrorRuleBadRequest creates a ReplaceHTTPErrorRuleBadRequest with default headers values
func NewReplaceHTTPErrorRuleBadRequest() *ReplaceHTTPErrorRuleBadRequest {
	return &ReplaceHTTPErrorRuleBadRequest{}
}

/*
ReplaceHTTPErrorRuleBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceHTTPErrorRuleBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this replace Http error rule bad request response has a 2xx status code
func (o *ReplaceHTTPErrorRuleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this replace Http error rule bad request response has a 3xx status code
func (o *ReplaceHTTPErrorRuleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace Http error rule bad request response has a 4xx status code
func (o *ReplaceHTTPErrorRuleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this replace Http error rule bad request response has a 5xx status code
func (o *ReplaceHTTPErrorRuleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this replace Http error rule bad request response a status code equal to that given
func (o *ReplaceHTTPErrorRuleBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the replace Http error rule bad request response
func (o *ReplaceHTTPErrorRuleBadRequest) Code() int {
	return 400
}

func (o *ReplaceHTTPErrorRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/http_error_rules/{index}][%d] replaceHttpErrorRuleBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceHTTPErrorRuleBadRequest) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/http_error_rules/{index}][%d] replaceHttpErrorRuleBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceHTTPErrorRuleBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceHTTPErrorRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceHTTPErrorRuleNotFound creates a ReplaceHTTPErrorRuleNotFound with default headers values
func NewReplaceHTTPErrorRuleNotFound() *ReplaceHTTPErrorRuleNotFound {
	return &ReplaceHTTPErrorRuleNotFound{}
}

/*
ReplaceHTTPErrorRuleNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type ReplaceHTTPErrorRuleNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this replace Http error rule not found response has a 2xx status code
func (o *ReplaceHTTPErrorRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this replace Http error rule not found response has a 3xx status code
func (o *ReplaceHTTPErrorRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace Http error rule not found response has a 4xx status code
func (o *ReplaceHTTPErrorRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this replace Http error rule not found response has a 5xx status code
func (o *ReplaceHTTPErrorRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this replace Http error rule not found response a status code equal to that given
func (o *ReplaceHTTPErrorRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the replace Http error rule not found response
func (o *ReplaceHTTPErrorRuleNotFound) Code() int {
	return 404
}

func (o *ReplaceHTTPErrorRuleNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/http_error_rules/{index}][%d] replaceHttpErrorRuleNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceHTTPErrorRuleNotFound) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/http_error_rules/{index}][%d] replaceHttpErrorRuleNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceHTTPErrorRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceHTTPErrorRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceHTTPErrorRuleDefault creates a ReplaceHTTPErrorRuleDefault with default headers values
func NewReplaceHTTPErrorRuleDefault(code int) *ReplaceHTTPErrorRuleDefault {
	return &ReplaceHTTPErrorRuleDefault{
		_statusCode: code,
	}
}

/*
ReplaceHTTPErrorRuleDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceHTTPErrorRuleDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this replace HTTP error rule default response has a 2xx status code
func (o *ReplaceHTTPErrorRuleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this replace HTTP error rule default response has a 3xx status code
func (o *ReplaceHTTPErrorRuleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this replace HTTP error rule default response has a 4xx status code
func (o *ReplaceHTTPErrorRuleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this replace HTTP error rule default response has a 5xx status code
func (o *ReplaceHTTPErrorRuleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this replace HTTP error rule default response a status code equal to that given
func (o *ReplaceHTTPErrorRuleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the replace HTTP error rule default response
func (o *ReplaceHTTPErrorRuleDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceHTTPErrorRuleDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/http_error_rules/{index}][%d] replaceHTTPErrorRule default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceHTTPErrorRuleDefault) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/http_error_rules/{index}][%d] replaceHTTPErrorRule default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceHTTPErrorRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceHTTPErrorRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
