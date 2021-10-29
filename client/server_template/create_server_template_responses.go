// Code generated by go-swagger; DO NOT EDIT.

package server_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/practical-coder/hdc/models"
)

// CreateServerTemplateReader is a Reader for the CreateServerTemplate structure.
type CreateServerTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateServerTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateServerTemplateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateServerTemplateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateServerTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateServerTemplateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateServerTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateServerTemplateCreated creates a CreateServerTemplateCreated with default headers values
func NewCreateServerTemplateCreated() *CreateServerTemplateCreated {
	return &CreateServerTemplateCreated{}
}

/* CreateServerTemplateCreated describes a response with status code 201, with default header values.

Server template created
*/
type CreateServerTemplateCreated struct {
	Payload *models.ServerTemplate
}

func (o *CreateServerTemplateCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/server_templates][%d] createServerTemplateCreated  %+v", 201, o.Payload)
}
func (o *CreateServerTemplateCreated) GetPayload() *models.ServerTemplate {
	return o.Payload
}

func (o *CreateServerTemplateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServerTemplateAccepted creates a CreateServerTemplateAccepted with default headers values
func NewCreateServerTemplateAccepted() *CreateServerTemplateAccepted {
	return &CreateServerTemplateAccepted{}
}

/* CreateServerTemplateAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateServerTemplateAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.ServerTemplate
}

func (o *CreateServerTemplateAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/server_templates][%d] createServerTemplateAccepted  %+v", 202, o.Payload)
}
func (o *CreateServerTemplateAccepted) GetPayload() *models.ServerTemplate {
	return o.Payload
}

func (o *CreateServerTemplateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.ServerTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServerTemplateBadRequest creates a CreateServerTemplateBadRequest with default headers values
func NewCreateServerTemplateBadRequest() *CreateServerTemplateBadRequest {
	return &CreateServerTemplateBadRequest{}
}

/* CreateServerTemplateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateServerTemplateBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateServerTemplateBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/server_templates][%d] createServerTemplateBadRequest  %+v", 400, o.Payload)
}
func (o *CreateServerTemplateBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateServerTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateServerTemplateConflict creates a CreateServerTemplateConflict with default headers values
func NewCreateServerTemplateConflict() *CreateServerTemplateConflict {
	return &CreateServerTemplateConflict{}
}

/* CreateServerTemplateConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateServerTemplateConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateServerTemplateConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/server_templates][%d] createServerTemplateConflict  %+v", 409, o.Payload)
}
func (o *CreateServerTemplateConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateServerTemplateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateServerTemplateDefault creates a CreateServerTemplateDefault with default headers values
func NewCreateServerTemplateDefault(code int) *CreateServerTemplateDefault {
	return &CreateServerTemplateDefault{
		_statusCode: code,
	}
}

/* CreateServerTemplateDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateServerTemplateDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create server template default response
func (o *CreateServerTemplateDefault) Code() int {
	return o._statusCode
}

func (o *CreateServerTemplateDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/server_templates][%d] createServerTemplate default  %+v", o._statusCode, o.Payload)
}
func (o *CreateServerTemplateDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateServerTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
