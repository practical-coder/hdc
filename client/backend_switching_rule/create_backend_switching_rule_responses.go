// Code generated by go-swagger; DO NOT EDIT.

package backend_switching_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v2/models"
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

/*CreateBackendSwitchingRuleCreated handles this case with default header values.

Backend Switching Rule created
*/
type CreateBackendSwitchingRuleCreated struct {
	Payload *models.BackendSwitchingRule
}

func (o *CreateBackendSwitchingRuleCreated) Error() string {
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

/*CreateBackendSwitchingRuleAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type CreateBackendSwitchingRuleAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string

	Payload *models.BackendSwitchingRule
}

func (o *CreateBackendSwitchingRuleAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/backend_switching_rules][%d] createBackendSwitchingRuleAccepted  %+v", 202, o.Payload)
}

func (o *CreateBackendSwitchingRuleAccepted) GetPayload() *models.BackendSwitchingRule {
	return o.Payload
}

func (o *CreateBackendSwitchingRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

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

/*CreateBackendSwitchingRuleBadRequest handles this case with default header values.

Bad request
*/
type CreateBackendSwitchingRuleBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateBackendSwitchingRuleBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/backend_switching_rules][%d] createBackendSwitchingRuleBadRequest  %+v", 400, o.Payload)
}

func (o *CreateBackendSwitchingRuleBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateBackendSwitchingRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

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

/*CreateBackendSwitchingRuleConflict handles this case with default header values.

The specified resource already exists
*/
type CreateBackendSwitchingRuleConflict struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateBackendSwitchingRuleConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/backend_switching_rules][%d] createBackendSwitchingRuleConflict  %+v", 409, o.Payload)
}

func (o *CreateBackendSwitchingRuleConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateBackendSwitchingRuleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

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

/*CreateBackendSwitchingRuleDefault handles this case with default header values.

General Error
*/
type CreateBackendSwitchingRuleDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create backend switching rule default response
func (o *CreateBackendSwitchingRuleDefault) Code() int {
	return o._statusCode
}

func (o *CreateBackendSwitchingRuleDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/backend_switching_rules][%d] createBackendSwitchingRule default  %+v", o._statusCode, o.Payload)
}

func (o *CreateBackendSwitchingRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateBackendSwitchingRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
