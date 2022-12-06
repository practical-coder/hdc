// Code generated by go-swagger; DO NOT EDIT.

package backend_switching_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// CreateBackendSwitchingRuleReader is a Reader for the CreateBackendSwitchingRule structure.
type CreateBackendSwitchingRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBackendSwitchingRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateBackendSwitchingRuleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateBackendSwitchingRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateBackendSwitchingRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateBackendSwitchingRuleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateBackendSwitchingRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateBackendSwitchingRuleCreated creates a CreateBackendSwitchingRuleCreated with default headers values
func NewCreateBackendSwitchingRuleCreated() *CreateBackendSwitchingRuleCreated {
	return &CreateBackendSwitchingRuleCreated{}
}

/*
CreateBackendSwitchingRuleCreated describes a response with status code 201, with default header values.

Backend Switching Rule created
*/
type CreateBackendSwitchingRuleCreated struct {
	Payload *models.BackendSwitchingRule
}

// IsSuccess returns true when this create backend switching rule created response has a 2xx status code
func (o *CreateBackendSwitchingRuleCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create backend switching rule created response has a 3xx status code
func (o *CreateBackendSwitchingRuleCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create backend switching rule created response has a 4xx status code
func (o *CreateBackendSwitchingRuleCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create backend switching rule created response has a 5xx status code
func (o *CreateBackendSwitchingRuleCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create backend switching rule created response a status code equal to that given
func (o *CreateBackendSwitchingRuleCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateBackendSwitchingRuleCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/backend_switching_rules][%d] createBackendSwitchingRuleCreated  %+v", 201, o.Payload)
}

func (o *CreateBackendSwitchingRuleCreated) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/backend_switching_rules][%d] createBackendSwitchingRuleCreated  %+v", 201, o.Payload)
}

func (o *CreateBackendSwitchingRuleCreated) GetPayload() *models.BackendSwitchingRule {
	return o.Payload
}

func (o *CreateBackendSwitchingRuleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackendSwitchingRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBackendSwitchingRuleAccepted creates a CreateBackendSwitchingRuleAccepted with default headers values
func NewCreateBackendSwitchingRuleAccepted() *CreateBackendSwitchingRuleAccepted {
	return &CreateBackendSwitchingRuleAccepted{}
}

/*
CreateBackendSwitchingRuleAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateBackendSwitchingRuleAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.BackendSwitchingRule
}

// IsSuccess returns true when this create backend switching rule accepted response has a 2xx status code
func (o *CreateBackendSwitchingRuleAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create backend switching rule accepted response has a 3xx status code
func (o *CreateBackendSwitchingRuleAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create backend switching rule accepted response has a 4xx status code
func (o *CreateBackendSwitchingRuleAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this create backend switching rule accepted response has a 5xx status code
func (o *CreateBackendSwitchingRuleAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this create backend switching rule accepted response a status code equal to that given
func (o *CreateBackendSwitchingRuleAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *CreateBackendSwitchingRuleAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/backend_switching_rules][%d] createBackendSwitchingRuleAccepted  %+v", 202, o.Payload)
}

func (o *CreateBackendSwitchingRuleAccepted) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/backend_switching_rules][%d] createBackendSwitchingRuleAccepted  %+v", 202, o.Payload)
}

func (o *CreateBackendSwitchingRuleAccepted) GetPayload() *models.BackendSwitchingRule {
	return o.Payload
}

func (o *CreateBackendSwitchingRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.BackendSwitchingRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBackendSwitchingRuleBadRequest creates a CreateBackendSwitchingRuleBadRequest with default headers values
func NewCreateBackendSwitchingRuleBadRequest() *CreateBackendSwitchingRuleBadRequest {
	return &CreateBackendSwitchingRuleBadRequest{}
}

/*
CreateBackendSwitchingRuleBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateBackendSwitchingRuleBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create backend switching rule bad request response has a 2xx status code
func (o *CreateBackendSwitchingRuleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create backend switching rule bad request response has a 3xx status code
func (o *CreateBackendSwitchingRuleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create backend switching rule bad request response has a 4xx status code
func (o *CreateBackendSwitchingRuleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create backend switching rule bad request response has a 5xx status code
func (o *CreateBackendSwitchingRuleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create backend switching rule bad request response a status code equal to that given
func (o *CreateBackendSwitchingRuleBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateBackendSwitchingRuleBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/backend_switching_rules][%d] createBackendSwitchingRuleBadRequest  %+v", 400, o.Payload)
}

func (o *CreateBackendSwitchingRuleBadRequest) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/backend_switching_rules][%d] createBackendSwitchingRuleBadRequest  %+v", 400, o.Payload)
}

func (o *CreateBackendSwitchingRuleBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateBackendSwitchingRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateBackendSwitchingRuleConflict creates a CreateBackendSwitchingRuleConflict with default headers values
func NewCreateBackendSwitchingRuleConflict() *CreateBackendSwitchingRuleConflict {
	return &CreateBackendSwitchingRuleConflict{}
}

/*
CreateBackendSwitchingRuleConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateBackendSwitchingRuleConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create backend switching rule conflict response has a 2xx status code
func (o *CreateBackendSwitchingRuleConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create backend switching rule conflict response has a 3xx status code
func (o *CreateBackendSwitchingRuleConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create backend switching rule conflict response has a 4xx status code
func (o *CreateBackendSwitchingRuleConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create backend switching rule conflict response has a 5xx status code
func (o *CreateBackendSwitchingRuleConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create backend switching rule conflict response a status code equal to that given
func (o *CreateBackendSwitchingRuleConflict) IsCode(code int) bool {
	return code == 409
}

func (o *CreateBackendSwitchingRuleConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/backend_switching_rules][%d] createBackendSwitchingRuleConflict  %+v", 409, o.Payload)
}

func (o *CreateBackendSwitchingRuleConflict) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/backend_switching_rules][%d] createBackendSwitchingRuleConflict  %+v", 409, o.Payload)
}

func (o *CreateBackendSwitchingRuleConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateBackendSwitchingRuleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateBackendSwitchingRuleDefault creates a CreateBackendSwitchingRuleDefault with default headers values
func NewCreateBackendSwitchingRuleDefault(code int) *CreateBackendSwitchingRuleDefault {
	return &CreateBackendSwitchingRuleDefault{
		_statusCode: code,
	}
}

/*
CreateBackendSwitchingRuleDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateBackendSwitchingRuleDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create backend switching rule default response
func (o *CreateBackendSwitchingRuleDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create backend switching rule default response has a 2xx status code
func (o *CreateBackendSwitchingRuleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create backend switching rule default response has a 3xx status code
func (o *CreateBackendSwitchingRuleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create backend switching rule default response has a 4xx status code
func (o *CreateBackendSwitchingRuleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create backend switching rule default response has a 5xx status code
func (o *CreateBackendSwitchingRuleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create backend switching rule default response a status code equal to that given
func (o *CreateBackendSwitchingRuleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateBackendSwitchingRuleDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/backend_switching_rules][%d] createBackendSwitchingRule default  %+v", o._statusCode, o.Payload)
}

func (o *CreateBackendSwitchingRuleDefault) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/backend_switching_rules][%d] createBackendSwitchingRule default  %+v", o._statusCode, o.Payload)
}

func (o *CreateBackendSwitchingRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateBackendSwitchingRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
