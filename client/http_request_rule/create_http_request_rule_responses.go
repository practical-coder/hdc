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

// CreateHTTPRequestRuleReader is a Reader for the CreateHTTPRequestRule structure.
type CreateHTTPRequestRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateHTTPRequestRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateHTTPRequestRuleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateHTTPRequestRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateHTTPRequestRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateHTTPRequestRuleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateHTTPRequestRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateHTTPRequestRuleCreated creates a CreateHTTPRequestRuleCreated with default headers values
func NewCreateHTTPRequestRuleCreated() *CreateHTTPRequestRuleCreated {
	return &CreateHTTPRequestRuleCreated{}
}

/*
CreateHTTPRequestRuleCreated describes a response with status code 201, with default header values.

HTTP Request Rule created
*/
type CreateHTTPRequestRuleCreated struct {
	Payload *models.HTTPRequestRule
}

// IsSuccess returns true when this create Http request rule created response has a 2xx status code
func (o *CreateHTTPRequestRuleCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create Http request rule created response has a 3xx status code
func (o *CreateHTTPRequestRuleCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Http request rule created response has a 4xx status code
func (o *CreateHTTPRequestRuleCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create Http request rule created response has a 5xx status code
func (o *CreateHTTPRequestRuleCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create Http request rule created response a status code equal to that given
func (o *CreateHTTPRequestRuleCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create Http request rule created response
func (o *CreateHTTPRequestRuleCreated) Code() int {
	return 201
}

func (o *CreateHTTPRequestRuleCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_request_rules][%d] createHttpRequestRuleCreated  %+v", 201, o.Payload)
}

func (o *CreateHTTPRequestRuleCreated) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_request_rules][%d] createHttpRequestRuleCreated  %+v", 201, o.Payload)
}

func (o *CreateHTTPRequestRuleCreated) GetPayload() *models.HTTPRequestRule {
	return o.Payload
}

func (o *CreateHTTPRequestRuleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPRequestRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHTTPRequestRuleAccepted creates a CreateHTTPRequestRuleAccepted with default headers values
func NewCreateHTTPRequestRuleAccepted() *CreateHTTPRequestRuleAccepted {
	return &CreateHTTPRequestRuleAccepted{}
}

/*
CreateHTTPRequestRuleAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateHTTPRequestRuleAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.HTTPRequestRule
}

// IsSuccess returns true when this create Http request rule accepted response has a 2xx status code
func (o *CreateHTTPRequestRuleAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create Http request rule accepted response has a 3xx status code
func (o *CreateHTTPRequestRuleAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Http request rule accepted response has a 4xx status code
func (o *CreateHTTPRequestRuleAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this create Http request rule accepted response has a 5xx status code
func (o *CreateHTTPRequestRuleAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this create Http request rule accepted response a status code equal to that given
func (o *CreateHTTPRequestRuleAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the create Http request rule accepted response
func (o *CreateHTTPRequestRuleAccepted) Code() int {
	return 202
}

func (o *CreateHTTPRequestRuleAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_request_rules][%d] createHttpRequestRuleAccepted  %+v", 202, o.Payload)
}

func (o *CreateHTTPRequestRuleAccepted) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_request_rules][%d] createHttpRequestRuleAccepted  %+v", 202, o.Payload)
}

func (o *CreateHTTPRequestRuleAccepted) GetPayload() *models.HTTPRequestRule {
	return o.Payload
}

func (o *CreateHTTPRequestRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.HTTPRequestRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHTTPRequestRuleBadRequest creates a CreateHTTPRequestRuleBadRequest with default headers values
func NewCreateHTTPRequestRuleBadRequest() *CreateHTTPRequestRuleBadRequest {
	return &CreateHTTPRequestRuleBadRequest{}
}

/*
CreateHTTPRequestRuleBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateHTTPRequestRuleBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create Http request rule bad request response has a 2xx status code
func (o *CreateHTTPRequestRuleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create Http request rule bad request response has a 3xx status code
func (o *CreateHTTPRequestRuleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Http request rule bad request response has a 4xx status code
func (o *CreateHTTPRequestRuleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create Http request rule bad request response has a 5xx status code
func (o *CreateHTTPRequestRuleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create Http request rule bad request response a status code equal to that given
func (o *CreateHTTPRequestRuleBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create Http request rule bad request response
func (o *CreateHTTPRequestRuleBadRequest) Code() int {
	return 400
}

func (o *CreateHTTPRequestRuleBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_request_rules][%d] createHttpRequestRuleBadRequest  %+v", 400, o.Payload)
}

func (o *CreateHTTPRequestRuleBadRequest) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_request_rules][%d] createHttpRequestRuleBadRequest  %+v", 400, o.Payload)
}

func (o *CreateHTTPRequestRuleBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateHTTPRequestRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateHTTPRequestRuleConflict creates a CreateHTTPRequestRuleConflict with default headers values
func NewCreateHTTPRequestRuleConflict() *CreateHTTPRequestRuleConflict {
	return &CreateHTTPRequestRuleConflict{}
}

/*
CreateHTTPRequestRuleConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateHTTPRequestRuleConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create Http request rule conflict response has a 2xx status code
func (o *CreateHTTPRequestRuleConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create Http request rule conflict response has a 3xx status code
func (o *CreateHTTPRequestRuleConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Http request rule conflict response has a 4xx status code
func (o *CreateHTTPRequestRuleConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create Http request rule conflict response has a 5xx status code
func (o *CreateHTTPRequestRuleConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create Http request rule conflict response a status code equal to that given
func (o *CreateHTTPRequestRuleConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create Http request rule conflict response
func (o *CreateHTTPRequestRuleConflict) Code() int {
	return 409
}

func (o *CreateHTTPRequestRuleConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_request_rules][%d] createHttpRequestRuleConflict  %+v", 409, o.Payload)
}

func (o *CreateHTTPRequestRuleConflict) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_request_rules][%d] createHttpRequestRuleConflict  %+v", 409, o.Payload)
}

func (o *CreateHTTPRequestRuleConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateHTTPRequestRuleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateHTTPRequestRuleDefault creates a CreateHTTPRequestRuleDefault with default headers values
func NewCreateHTTPRequestRuleDefault(code int) *CreateHTTPRequestRuleDefault {
	return &CreateHTTPRequestRuleDefault{
		_statusCode: code,
	}
}

/*
CreateHTTPRequestRuleDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateHTTPRequestRuleDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create HTTP request rule default response has a 2xx status code
func (o *CreateHTTPRequestRuleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create HTTP request rule default response has a 3xx status code
func (o *CreateHTTPRequestRuleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create HTTP request rule default response has a 4xx status code
func (o *CreateHTTPRequestRuleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create HTTP request rule default response has a 5xx status code
func (o *CreateHTTPRequestRuleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create HTTP request rule default response a status code equal to that given
func (o *CreateHTTPRequestRuleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create HTTP request rule default response
func (o *CreateHTTPRequestRuleDefault) Code() int {
	return o._statusCode
}

func (o *CreateHTTPRequestRuleDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_request_rules][%d] createHTTPRequestRule default  %+v", o._statusCode, o.Payload)
}

func (o *CreateHTTPRequestRuleDefault) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_request_rules][%d] createHTTPRequestRule default  %+v", o._statusCode, o.Payload)
}

func (o *CreateHTTPRequestRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateHTTPRequestRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
