// Code generated by go-swagger; DO NOT EDIT.

package spoe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models"
)

// CreateSpoeAgentReader is a Reader for the CreateSpoeAgent structure.
type CreateSpoeAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSpoeAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateSpoeAgentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateSpoeAgentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateSpoeAgentConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateSpoeAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSpoeAgentCreated creates a CreateSpoeAgentCreated with default headers values
func NewCreateSpoeAgentCreated() *CreateSpoeAgentCreated {
	return &CreateSpoeAgentCreated{}
}

/* CreateSpoeAgentCreated describes a response with status code 201, with default header values.

Spoe agent created
*/
type CreateSpoeAgentCreated struct {
	Payload *models.SpoeAgent
}

func (o *CreateSpoeAgentCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_agents][%d] createSpoeAgentCreated  %+v", 201, o.Payload)
}
func (o *CreateSpoeAgentCreated) GetPayload() *models.SpoeAgent {
	return o.Payload
}

func (o *CreateSpoeAgentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SpoeAgent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSpoeAgentBadRequest creates a CreateSpoeAgentBadRequest with default headers values
func NewCreateSpoeAgentBadRequest() *CreateSpoeAgentBadRequest {
	var (
		// initialize headers with default values
		configurationVersionDefault = int64(0)
	)

	return &CreateSpoeAgentBadRequest{

		ConfigurationVersion: configurationVersionDefault,
	}
}

/* CreateSpoeAgentBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateSpoeAgentBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *CreateSpoeAgentBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_agents][%d] createSpoeAgentBadRequest  %+v", 400, o.Payload)
}
func (o *CreateSpoeAgentBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSpoeAgentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		valconfigurationVersion, err := swag.ConvertInt64(hdrConfigurationVersion)
		if err != nil {
			return errors.InvalidType("Configuration-Version", "header", "int64", hdrConfigurationVersion)
		}
		o.ConfigurationVersion = valconfigurationVersion
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSpoeAgentConflict creates a CreateSpoeAgentConflict with default headers values
func NewCreateSpoeAgentConflict() *CreateSpoeAgentConflict {
	var (
		// initialize headers with default values
		configurationVersionDefault = int64(0)
	)

	return &CreateSpoeAgentConflict{

		ConfigurationVersion: configurationVersionDefault,
	}
}

/* CreateSpoeAgentConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateSpoeAgentConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *CreateSpoeAgentConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_agents][%d] createSpoeAgentConflict  %+v", 409, o.Payload)
}
func (o *CreateSpoeAgentConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSpoeAgentConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		valconfigurationVersion, err := swag.ConvertInt64(hdrConfigurationVersion)
		if err != nil {
			return errors.InvalidType("Configuration-Version", "header", "int64", hdrConfigurationVersion)
		}
		o.ConfigurationVersion = valconfigurationVersion
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSpoeAgentDefault creates a CreateSpoeAgentDefault with default headers values
func NewCreateSpoeAgentDefault(code int) *CreateSpoeAgentDefault {
	var (
		// initialize headers with default values
		configurationVersionDefault = int64(0)
	)

	return &CreateSpoeAgentDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

/* CreateSpoeAgentDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateSpoeAgentDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the create spoe agent default response
func (o *CreateSpoeAgentDefault) Code() int {
	return o._statusCode
}

func (o *CreateSpoeAgentDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_agents][%d] createSpoeAgent default  %+v", o._statusCode, o.Payload)
}
func (o *CreateSpoeAgentDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSpoeAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		valconfigurationVersion, err := swag.ConvertInt64(hdrConfigurationVersion)
		if err != nil {
			return errors.InvalidType("Configuration-Version", "header", "int64", hdrConfigurationVersion)
		}
		o.ConfigurationVersion = valconfigurationVersion
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
