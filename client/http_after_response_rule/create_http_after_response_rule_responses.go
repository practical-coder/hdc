// Code generated by go-swagger; DO NOT EDIT.

package http_after_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// CreateHTTPAfterResponseRuleReader is a Reader for the CreateHTTPAfterResponseRule structure.
type CreateHTTPAfterResponseRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateHTTPAfterResponseRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateHTTPAfterResponseRuleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateHTTPAfterResponseRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateHTTPAfterResponseRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateHTTPAfterResponseRuleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateHTTPAfterResponseRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateHTTPAfterResponseRuleCreated creates a CreateHTTPAfterResponseRuleCreated with default headers values
func NewCreateHTTPAfterResponseRuleCreated() *CreateHTTPAfterResponseRuleCreated {
	return &CreateHTTPAfterResponseRuleCreated{}
}

/*
CreateHTTPAfterResponseRuleCreated describes a response with status code 201, with default header values.

HTTP Response Rule created
*/
type CreateHTTPAfterResponseRuleCreated struct {
	Payload *models.HTTPAfterResponseRule
}

// IsSuccess returns true when this create Http after response rule created response has a 2xx status code
func (o *CreateHTTPAfterResponseRuleCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create Http after response rule created response has a 3xx status code
func (o *CreateHTTPAfterResponseRuleCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Http after response rule created response has a 4xx status code
func (o *CreateHTTPAfterResponseRuleCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create Http after response rule created response has a 5xx status code
func (o *CreateHTTPAfterResponseRuleCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create Http after response rule created response a status code equal to that given
func (o *CreateHTTPAfterResponseRuleCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateHTTPAfterResponseRuleCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_after_response_rules][%d] createHttpAfterResponseRuleCreated  %+v", 201, o.Payload)
}

func (o *CreateHTTPAfterResponseRuleCreated) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_after_response_rules][%d] createHttpAfterResponseRuleCreated  %+v", 201, o.Payload)
}

func (o *CreateHTTPAfterResponseRuleCreated) GetPayload() *models.HTTPAfterResponseRule {
	return o.Payload
}

func (o *CreateHTTPAfterResponseRuleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPAfterResponseRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHTTPAfterResponseRuleAccepted creates a CreateHTTPAfterResponseRuleAccepted with default headers values
func NewCreateHTTPAfterResponseRuleAccepted() *CreateHTTPAfterResponseRuleAccepted {
	return &CreateHTTPAfterResponseRuleAccepted{}
}

/*
CreateHTTPAfterResponseRuleAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateHTTPAfterResponseRuleAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.HTTPAfterResponseRule
}

// IsSuccess returns true when this create Http after response rule accepted response has a 2xx status code
func (o *CreateHTTPAfterResponseRuleAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create Http after response rule accepted response has a 3xx status code
func (o *CreateHTTPAfterResponseRuleAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Http after response rule accepted response has a 4xx status code
func (o *CreateHTTPAfterResponseRuleAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this create Http after response rule accepted response has a 5xx status code
func (o *CreateHTTPAfterResponseRuleAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this create Http after response rule accepted response a status code equal to that given
func (o *CreateHTTPAfterResponseRuleAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *CreateHTTPAfterResponseRuleAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_after_response_rules][%d] createHttpAfterResponseRuleAccepted  %+v", 202, o.Payload)
}

func (o *CreateHTTPAfterResponseRuleAccepted) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_after_response_rules][%d] createHttpAfterResponseRuleAccepted  %+v", 202, o.Payload)
}

func (o *CreateHTTPAfterResponseRuleAccepted) GetPayload() *models.HTTPAfterResponseRule {
	return o.Payload
}

func (o *CreateHTTPAfterResponseRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.HTTPAfterResponseRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHTTPAfterResponseRuleBadRequest creates a CreateHTTPAfterResponseRuleBadRequest with default headers values
func NewCreateHTTPAfterResponseRuleBadRequest() *CreateHTTPAfterResponseRuleBadRequest {
	return &CreateHTTPAfterResponseRuleBadRequest{}
}

/*
CreateHTTPAfterResponseRuleBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateHTTPAfterResponseRuleBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create Http after response rule bad request response has a 2xx status code
func (o *CreateHTTPAfterResponseRuleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create Http after response rule bad request response has a 3xx status code
func (o *CreateHTTPAfterResponseRuleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Http after response rule bad request response has a 4xx status code
func (o *CreateHTTPAfterResponseRuleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create Http after response rule bad request response has a 5xx status code
func (o *CreateHTTPAfterResponseRuleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create Http after response rule bad request response a status code equal to that given
func (o *CreateHTTPAfterResponseRuleBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateHTTPAfterResponseRuleBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_after_response_rules][%d] createHttpAfterResponseRuleBadRequest  %+v", 400, o.Payload)
}

func (o *CreateHTTPAfterResponseRuleBadRequest) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_after_response_rules][%d] createHttpAfterResponseRuleBadRequest  %+v", 400, o.Payload)
}

func (o *CreateHTTPAfterResponseRuleBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateHTTPAfterResponseRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateHTTPAfterResponseRuleConflict creates a CreateHTTPAfterResponseRuleConflict with default headers values
func NewCreateHTTPAfterResponseRuleConflict() *CreateHTTPAfterResponseRuleConflict {
	return &CreateHTTPAfterResponseRuleConflict{}
}

/*
CreateHTTPAfterResponseRuleConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateHTTPAfterResponseRuleConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create Http after response rule conflict response has a 2xx status code
func (o *CreateHTTPAfterResponseRuleConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create Http after response rule conflict response has a 3xx status code
func (o *CreateHTTPAfterResponseRuleConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Http after response rule conflict response has a 4xx status code
func (o *CreateHTTPAfterResponseRuleConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create Http after response rule conflict response has a 5xx status code
func (o *CreateHTTPAfterResponseRuleConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create Http after response rule conflict response a status code equal to that given
func (o *CreateHTTPAfterResponseRuleConflict) IsCode(code int) bool {
	return code == 409
}

func (o *CreateHTTPAfterResponseRuleConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_after_response_rules][%d] createHttpAfterResponseRuleConflict  %+v", 409, o.Payload)
}

func (o *CreateHTTPAfterResponseRuleConflict) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_after_response_rules][%d] createHttpAfterResponseRuleConflict  %+v", 409, o.Payload)
}

func (o *CreateHTTPAfterResponseRuleConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateHTTPAfterResponseRuleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateHTTPAfterResponseRuleDefault creates a CreateHTTPAfterResponseRuleDefault with default headers values
func NewCreateHTTPAfterResponseRuleDefault(code int) *CreateHTTPAfterResponseRuleDefault {
	return &CreateHTTPAfterResponseRuleDefault{
		_statusCode: code,
	}
}

/*
CreateHTTPAfterResponseRuleDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateHTTPAfterResponseRuleDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create HTTP after response rule default response
func (o *CreateHTTPAfterResponseRuleDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create HTTP after response rule default response has a 2xx status code
func (o *CreateHTTPAfterResponseRuleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create HTTP after response rule default response has a 3xx status code
func (o *CreateHTTPAfterResponseRuleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create HTTP after response rule default response has a 4xx status code
func (o *CreateHTTPAfterResponseRuleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create HTTP after response rule default response has a 5xx status code
func (o *CreateHTTPAfterResponseRuleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create HTTP after response rule default response a status code equal to that given
func (o *CreateHTTPAfterResponseRuleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateHTTPAfterResponseRuleDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_after_response_rules][%d] createHTTPAfterResponseRule default  %+v", o._statusCode, o.Payload)
}

func (o *CreateHTTPAfterResponseRuleDefault) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/http_after_response_rules][%d] createHTTPAfterResponseRule default  %+v", o._statusCode, o.Payload)
}

func (o *CreateHTTPAfterResponseRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateHTTPAfterResponseRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
