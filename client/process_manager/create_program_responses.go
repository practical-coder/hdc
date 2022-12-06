// Code generated by go-swagger; DO NOT EDIT.

package process_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// CreateProgramReader is a Reader for the CreateProgram structure.
type CreateProgramReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProgramReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateProgramCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateProgramAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateProgramBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateProgramConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateProgramDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateProgramCreated creates a CreateProgramCreated with default headers values
func NewCreateProgramCreated() *CreateProgramCreated {
	return &CreateProgramCreated{}
}

/*
CreateProgramCreated describes a response with status code 201, with default header values.

Program created
*/
type CreateProgramCreated struct {
	Payload *models.Program
}

// IsSuccess returns true when this create program created response has a 2xx status code
func (o *CreateProgramCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create program created response has a 3xx status code
func (o *CreateProgramCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create program created response has a 4xx status code
func (o *CreateProgramCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create program created response has a 5xx status code
func (o *CreateProgramCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create program created response a status code equal to that given
func (o *CreateProgramCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateProgramCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/programs][%d] createProgramCreated  %+v", 201, o.Payload)
}

func (o *CreateProgramCreated) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/programs][%d] createProgramCreated  %+v", 201, o.Payload)
}

func (o *CreateProgramCreated) GetPayload() *models.Program {
	return o.Payload
}

func (o *CreateProgramCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Program)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProgramAccepted creates a CreateProgramAccepted with default headers values
func NewCreateProgramAccepted() *CreateProgramAccepted {
	return &CreateProgramAccepted{}
}

/*
CreateProgramAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateProgramAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.Program
}

// IsSuccess returns true when this create program accepted response has a 2xx status code
func (o *CreateProgramAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create program accepted response has a 3xx status code
func (o *CreateProgramAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create program accepted response has a 4xx status code
func (o *CreateProgramAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this create program accepted response has a 5xx status code
func (o *CreateProgramAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this create program accepted response a status code equal to that given
func (o *CreateProgramAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *CreateProgramAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/programs][%d] createProgramAccepted  %+v", 202, o.Payload)
}

func (o *CreateProgramAccepted) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/programs][%d] createProgramAccepted  %+v", 202, o.Payload)
}

func (o *CreateProgramAccepted) GetPayload() *models.Program {
	return o.Payload
}

func (o *CreateProgramAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.Program)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProgramBadRequest creates a CreateProgramBadRequest with default headers values
func NewCreateProgramBadRequest() *CreateProgramBadRequest {
	return &CreateProgramBadRequest{}
}

/*
CreateProgramBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateProgramBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create program bad request response has a 2xx status code
func (o *CreateProgramBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create program bad request response has a 3xx status code
func (o *CreateProgramBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create program bad request response has a 4xx status code
func (o *CreateProgramBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create program bad request response has a 5xx status code
func (o *CreateProgramBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create program bad request response a status code equal to that given
func (o *CreateProgramBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateProgramBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/programs][%d] createProgramBadRequest  %+v", 400, o.Payload)
}

func (o *CreateProgramBadRequest) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/programs][%d] createProgramBadRequest  %+v", 400, o.Payload)
}

func (o *CreateProgramBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateProgramBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateProgramConflict creates a CreateProgramConflict with default headers values
func NewCreateProgramConflict() *CreateProgramConflict {
	return &CreateProgramConflict{}
}

/*
CreateProgramConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateProgramConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create program conflict response has a 2xx status code
func (o *CreateProgramConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create program conflict response has a 3xx status code
func (o *CreateProgramConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create program conflict response has a 4xx status code
func (o *CreateProgramConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create program conflict response has a 5xx status code
func (o *CreateProgramConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create program conflict response a status code equal to that given
func (o *CreateProgramConflict) IsCode(code int) bool {
	return code == 409
}

func (o *CreateProgramConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/programs][%d] createProgramConflict  %+v", 409, o.Payload)
}

func (o *CreateProgramConflict) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/programs][%d] createProgramConflict  %+v", 409, o.Payload)
}

func (o *CreateProgramConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateProgramConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateProgramDefault creates a CreateProgramDefault with default headers values
func NewCreateProgramDefault(code int) *CreateProgramDefault {
	return &CreateProgramDefault{
		_statusCode: code,
	}
}

/*
CreateProgramDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateProgramDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create program default response
func (o *CreateProgramDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create program default response has a 2xx status code
func (o *CreateProgramDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create program default response has a 3xx status code
func (o *CreateProgramDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create program default response has a 4xx status code
func (o *CreateProgramDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create program default response has a 5xx status code
func (o *CreateProgramDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create program default response a status code equal to that given
func (o *CreateProgramDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateProgramDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/programs][%d] createProgram default  %+v", o._statusCode, o.Payload)
}

func (o *CreateProgramDefault) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/programs][%d] createProgram default  %+v", o._statusCode, o.Payload)
}

func (o *CreateProgramDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateProgramDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
