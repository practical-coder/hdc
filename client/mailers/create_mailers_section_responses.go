// Code generated by go-swagger; DO NOT EDIT.

package mailers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v5/models"
)

// CreateMailersSectionReader is a Reader for the CreateMailersSection structure.
type CreateMailersSectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMailersSectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateMailersSectionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateMailersSectionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateMailersSectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateMailersSectionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateMailersSectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateMailersSectionCreated creates a CreateMailersSectionCreated with default headers values
func NewCreateMailersSectionCreated() *CreateMailersSectionCreated {
	return &CreateMailersSectionCreated{}
}

/*
CreateMailersSectionCreated describes a response with status code 201, with default header values.

Mailers created
*/
type CreateMailersSectionCreated struct {
	Payload *models.MailersSection
}

// IsSuccess returns true when this create mailers section created response has a 2xx status code
func (o *CreateMailersSectionCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create mailers section created response has a 3xx status code
func (o *CreateMailersSectionCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create mailers section created response has a 4xx status code
func (o *CreateMailersSectionCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create mailers section created response has a 5xx status code
func (o *CreateMailersSectionCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create mailers section created response a status code equal to that given
func (o *CreateMailersSectionCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create mailers section created response
func (o *CreateMailersSectionCreated) Code() int {
	return 201
}

func (o *CreateMailersSectionCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/mailers_section][%d] createMailersSectionCreated  %+v", 201, o.Payload)
}

func (o *CreateMailersSectionCreated) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/mailers_section][%d] createMailersSectionCreated  %+v", 201, o.Payload)
}

func (o *CreateMailersSectionCreated) GetPayload() *models.MailersSection {
	return o.Payload
}

func (o *CreateMailersSectionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MailersSection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMailersSectionAccepted creates a CreateMailersSectionAccepted with default headers values
func NewCreateMailersSectionAccepted() *CreateMailersSectionAccepted {
	return &CreateMailersSectionAccepted{}
}

/*
CreateMailersSectionAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateMailersSectionAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.MailersSection
}

// IsSuccess returns true when this create mailers section accepted response has a 2xx status code
func (o *CreateMailersSectionAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create mailers section accepted response has a 3xx status code
func (o *CreateMailersSectionAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create mailers section accepted response has a 4xx status code
func (o *CreateMailersSectionAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this create mailers section accepted response has a 5xx status code
func (o *CreateMailersSectionAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this create mailers section accepted response a status code equal to that given
func (o *CreateMailersSectionAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the create mailers section accepted response
func (o *CreateMailersSectionAccepted) Code() int {
	return 202
}

func (o *CreateMailersSectionAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/mailers_section][%d] createMailersSectionAccepted  %+v", 202, o.Payload)
}

func (o *CreateMailersSectionAccepted) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/mailers_section][%d] createMailersSectionAccepted  %+v", 202, o.Payload)
}

func (o *CreateMailersSectionAccepted) GetPayload() *models.MailersSection {
	return o.Payload
}

func (o *CreateMailersSectionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.MailersSection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMailersSectionBadRequest creates a CreateMailersSectionBadRequest with default headers values
func NewCreateMailersSectionBadRequest() *CreateMailersSectionBadRequest {
	return &CreateMailersSectionBadRequest{}
}

/*
CreateMailersSectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateMailersSectionBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create mailers section bad request response has a 2xx status code
func (o *CreateMailersSectionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create mailers section bad request response has a 3xx status code
func (o *CreateMailersSectionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create mailers section bad request response has a 4xx status code
func (o *CreateMailersSectionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create mailers section bad request response has a 5xx status code
func (o *CreateMailersSectionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create mailers section bad request response a status code equal to that given
func (o *CreateMailersSectionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create mailers section bad request response
func (o *CreateMailersSectionBadRequest) Code() int {
	return 400
}

func (o *CreateMailersSectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/mailers_section][%d] createMailersSectionBadRequest  %+v", 400, o.Payload)
}

func (o *CreateMailersSectionBadRequest) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/mailers_section][%d] createMailersSectionBadRequest  %+v", 400, o.Payload)
}

func (o *CreateMailersSectionBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateMailersSectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateMailersSectionConflict creates a CreateMailersSectionConflict with default headers values
func NewCreateMailersSectionConflict() *CreateMailersSectionConflict {
	return &CreateMailersSectionConflict{}
}

/*
CreateMailersSectionConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateMailersSectionConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create mailers section conflict response has a 2xx status code
func (o *CreateMailersSectionConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create mailers section conflict response has a 3xx status code
func (o *CreateMailersSectionConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create mailers section conflict response has a 4xx status code
func (o *CreateMailersSectionConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create mailers section conflict response has a 5xx status code
func (o *CreateMailersSectionConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create mailers section conflict response a status code equal to that given
func (o *CreateMailersSectionConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create mailers section conflict response
func (o *CreateMailersSectionConflict) Code() int {
	return 409
}

func (o *CreateMailersSectionConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/mailers_section][%d] createMailersSectionConflict  %+v", 409, o.Payload)
}

func (o *CreateMailersSectionConflict) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/mailers_section][%d] createMailersSectionConflict  %+v", 409, o.Payload)
}

func (o *CreateMailersSectionConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateMailersSectionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateMailersSectionDefault creates a CreateMailersSectionDefault with default headers values
func NewCreateMailersSectionDefault(code int) *CreateMailersSectionDefault {
	return &CreateMailersSectionDefault{
		_statusCode: code,
	}
}

/*
CreateMailersSectionDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateMailersSectionDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create mailers section default response has a 2xx status code
func (o *CreateMailersSectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create mailers section default response has a 3xx status code
func (o *CreateMailersSectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create mailers section default response has a 4xx status code
func (o *CreateMailersSectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create mailers section default response has a 5xx status code
func (o *CreateMailersSectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create mailers section default response a status code equal to that given
func (o *CreateMailersSectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create mailers section default response
func (o *CreateMailersSectionDefault) Code() int {
	return o._statusCode
}

func (o *CreateMailersSectionDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/mailers_section][%d] createMailersSection default  %+v", o._statusCode, o.Payload)
}

func (o *CreateMailersSectionDefault) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/mailers_section][%d] createMailersSection default  %+v", o._statusCode, o.Payload)
}

func (o *CreateMailersSectionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateMailersSectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
