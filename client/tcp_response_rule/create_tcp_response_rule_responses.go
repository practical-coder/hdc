// Code generated by go-swagger; DO NOT EDIT.

package tcp_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v2/models"
)

// CreateTCPResponseRuleReader is a Reader for the CreateTCPResponseRule structure.
type CreateTCPResponseRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTCPResponseRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateTCPResponseRuleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateTCPResponseRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTCPResponseRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateTCPResponseRuleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateTCPResponseRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateTCPResponseRuleCreated creates a CreateTCPResponseRuleCreated with default headers values
func NewCreateTCPResponseRuleCreated() *CreateTCPResponseRuleCreated {
	return &CreateTCPResponseRuleCreated{}
}

/* CreateTCPResponseRuleCreated describes a response with status code 201, with default header values.

TCP Response Rule created
*/
type CreateTCPResponseRuleCreated struct {
	Payload *models.TCPResponseRule
}

func (o *CreateTCPResponseRuleCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/tcp_response_rules][%d] createTcpResponseRuleCreated  %+v", 201, o.Payload)
}
func (o *CreateTCPResponseRuleCreated) GetPayload() *models.TCPResponseRule {
	return o.Payload
}

func (o *CreateTCPResponseRuleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TCPResponseRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTCPResponseRuleAccepted creates a CreateTCPResponseRuleAccepted with default headers values
func NewCreateTCPResponseRuleAccepted() *CreateTCPResponseRuleAccepted {
	return &CreateTCPResponseRuleAccepted{}
}

/* CreateTCPResponseRuleAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateTCPResponseRuleAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.TCPResponseRule
}

func (o *CreateTCPResponseRuleAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/tcp_response_rules][%d] createTcpResponseRuleAccepted  %+v", 202, o.Payload)
}
func (o *CreateTCPResponseRuleAccepted) GetPayload() *models.TCPResponseRule {
	return o.Payload
}

func (o *CreateTCPResponseRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.TCPResponseRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTCPResponseRuleBadRequest creates a CreateTCPResponseRuleBadRequest with default headers values
func NewCreateTCPResponseRuleBadRequest() *CreateTCPResponseRuleBadRequest {
	return &CreateTCPResponseRuleBadRequest{}
}

/* CreateTCPResponseRuleBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateTCPResponseRuleBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateTCPResponseRuleBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/tcp_response_rules][%d] createTcpResponseRuleBadRequest  %+v", 400, o.Payload)
}
func (o *CreateTCPResponseRuleBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateTCPResponseRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateTCPResponseRuleConflict creates a CreateTCPResponseRuleConflict with default headers values
func NewCreateTCPResponseRuleConflict() *CreateTCPResponseRuleConflict {
	return &CreateTCPResponseRuleConflict{}
}

/* CreateTCPResponseRuleConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateTCPResponseRuleConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateTCPResponseRuleConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/tcp_response_rules][%d] createTcpResponseRuleConflict  %+v", 409, o.Payload)
}
func (o *CreateTCPResponseRuleConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateTCPResponseRuleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateTCPResponseRuleDefault creates a CreateTCPResponseRuleDefault with default headers values
func NewCreateTCPResponseRuleDefault(code int) *CreateTCPResponseRuleDefault {
	return &CreateTCPResponseRuleDefault{
		_statusCode: code,
	}
}

/* CreateTCPResponseRuleDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateTCPResponseRuleDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create TCP response rule default response
func (o *CreateTCPResponseRuleDefault) Code() int {
	return o._statusCode
}

func (o *CreateTCPResponseRuleDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/tcp_response_rules][%d] createTCPResponseRule default  %+v", o._statusCode, o.Payload)
}
func (o *CreateTCPResponseRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateTCPResponseRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
