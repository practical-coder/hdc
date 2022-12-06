// Code generated by go-swagger; DO NOT EDIT.

package declare_capture

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// CreateDeclareCaptureReader is a Reader for the CreateDeclareCapture structure.
type CreateDeclareCaptureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDeclareCaptureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDeclareCaptureCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateDeclareCaptureAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDeclareCaptureBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateDeclareCaptureConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateDeclareCaptureDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateDeclareCaptureCreated creates a CreateDeclareCaptureCreated with default headers values
func NewCreateDeclareCaptureCreated() *CreateDeclareCaptureCreated {
	return &CreateDeclareCaptureCreated{}
}

/*
CreateDeclareCaptureCreated describes a response with status code 201, with default header values.

Declare capture created
*/
type CreateDeclareCaptureCreated struct {
	Payload *models.Capture
}

// IsSuccess returns true when this create declare capture created response has a 2xx status code
func (o *CreateDeclareCaptureCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create declare capture created response has a 3xx status code
func (o *CreateDeclareCaptureCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create declare capture created response has a 4xx status code
func (o *CreateDeclareCaptureCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create declare capture created response has a 5xx status code
func (o *CreateDeclareCaptureCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create declare capture created response a status code equal to that given
func (o *CreateDeclareCaptureCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateDeclareCaptureCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/captures][%d] createDeclareCaptureCreated  %+v", 201, o.Payload)
}

func (o *CreateDeclareCaptureCreated) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/captures][%d] createDeclareCaptureCreated  %+v", 201, o.Payload)
}

func (o *CreateDeclareCaptureCreated) GetPayload() *models.Capture {
	return o.Payload
}

func (o *CreateDeclareCaptureCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Capture)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDeclareCaptureAccepted creates a CreateDeclareCaptureAccepted with default headers values
func NewCreateDeclareCaptureAccepted() *CreateDeclareCaptureAccepted {
	return &CreateDeclareCaptureAccepted{}
}

/*
CreateDeclareCaptureAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateDeclareCaptureAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.Capture
}

// IsSuccess returns true when this create declare capture accepted response has a 2xx status code
func (o *CreateDeclareCaptureAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create declare capture accepted response has a 3xx status code
func (o *CreateDeclareCaptureAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create declare capture accepted response has a 4xx status code
func (o *CreateDeclareCaptureAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this create declare capture accepted response has a 5xx status code
func (o *CreateDeclareCaptureAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this create declare capture accepted response a status code equal to that given
func (o *CreateDeclareCaptureAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *CreateDeclareCaptureAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/captures][%d] createDeclareCaptureAccepted  %+v", 202, o.Payload)
}

func (o *CreateDeclareCaptureAccepted) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/captures][%d] createDeclareCaptureAccepted  %+v", 202, o.Payload)
}

func (o *CreateDeclareCaptureAccepted) GetPayload() *models.Capture {
	return o.Payload
}

func (o *CreateDeclareCaptureAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.Capture)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDeclareCaptureBadRequest creates a CreateDeclareCaptureBadRequest with default headers values
func NewCreateDeclareCaptureBadRequest() *CreateDeclareCaptureBadRequest {
	return &CreateDeclareCaptureBadRequest{}
}

/*
CreateDeclareCaptureBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateDeclareCaptureBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create declare capture bad request response has a 2xx status code
func (o *CreateDeclareCaptureBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create declare capture bad request response has a 3xx status code
func (o *CreateDeclareCaptureBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create declare capture bad request response has a 4xx status code
func (o *CreateDeclareCaptureBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create declare capture bad request response has a 5xx status code
func (o *CreateDeclareCaptureBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create declare capture bad request response a status code equal to that given
func (o *CreateDeclareCaptureBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateDeclareCaptureBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/captures][%d] createDeclareCaptureBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDeclareCaptureBadRequest) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/captures][%d] createDeclareCaptureBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDeclareCaptureBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDeclareCaptureBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateDeclareCaptureConflict creates a CreateDeclareCaptureConflict with default headers values
func NewCreateDeclareCaptureConflict() *CreateDeclareCaptureConflict {
	return &CreateDeclareCaptureConflict{}
}

/*
CreateDeclareCaptureConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateDeclareCaptureConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create declare capture conflict response has a 2xx status code
func (o *CreateDeclareCaptureConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create declare capture conflict response has a 3xx status code
func (o *CreateDeclareCaptureConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create declare capture conflict response has a 4xx status code
func (o *CreateDeclareCaptureConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create declare capture conflict response has a 5xx status code
func (o *CreateDeclareCaptureConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create declare capture conflict response a status code equal to that given
func (o *CreateDeclareCaptureConflict) IsCode(code int) bool {
	return code == 409
}

func (o *CreateDeclareCaptureConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/captures][%d] createDeclareCaptureConflict  %+v", 409, o.Payload)
}

func (o *CreateDeclareCaptureConflict) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/captures][%d] createDeclareCaptureConflict  %+v", 409, o.Payload)
}

func (o *CreateDeclareCaptureConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDeclareCaptureConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateDeclareCaptureDefault creates a CreateDeclareCaptureDefault with default headers values
func NewCreateDeclareCaptureDefault(code int) *CreateDeclareCaptureDefault {
	return &CreateDeclareCaptureDefault{
		_statusCode: code,
	}
}

/*
CreateDeclareCaptureDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateDeclareCaptureDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create declare capture default response
func (o *CreateDeclareCaptureDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create declare capture default response has a 2xx status code
func (o *CreateDeclareCaptureDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create declare capture default response has a 3xx status code
func (o *CreateDeclareCaptureDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create declare capture default response has a 4xx status code
func (o *CreateDeclareCaptureDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create declare capture default response has a 5xx status code
func (o *CreateDeclareCaptureDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create declare capture default response a status code equal to that given
func (o *CreateDeclareCaptureDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateDeclareCaptureDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/captures][%d] createDeclareCapture default  %+v", o._statusCode, o.Payload)
}

func (o *CreateDeclareCaptureDefault) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/captures][%d] createDeclareCapture default  %+v", o._statusCode, o.Payload)
}

func (o *CreateDeclareCaptureDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDeclareCaptureDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
