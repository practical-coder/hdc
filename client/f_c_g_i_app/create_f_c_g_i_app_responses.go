// Code generated by go-swagger; DO NOT EDIT.

package f_c_g_i_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// CreateFCGIAppReader is a Reader for the CreateFCGIApp structure.
type CreateFCGIAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateFCGIAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateFCGIAppCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateFCGIAppAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateFCGIAppBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateFCGIAppConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateFCGIAppDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateFCGIAppCreated creates a CreateFCGIAppCreated with default headers values
func NewCreateFCGIAppCreated() *CreateFCGIAppCreated {
	return &CreateFCGIAppCreated{}
}

/*
CreateFCGIAppCreated describes a response with status code 201, with default header values.

Application created
*/
type CreateFCGIAppCreated struct {
	Payload *models.FcgiApp
}

// IsSuccess returns true when this create f c g i app created response has a 2xx status code
func (o *CreateFCGIAppCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create f c g i app created response has a 3xx status code
func (o *CreateFCGIAppCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create f c g i app created response has a 4xx status code
func (o *CreateFCGIAppCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create f c g i app created response has a 5xx status code
func (o *CreateFCGIAppCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create f c g i app created response a status code equal to that given
func (o *CreateFCGIAppCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateFCGIAppCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/fcgi_apps][%d] createFCGIAppCreated  %+v", 201, o.Payload)
}

func (o *CreateFCGIAppCreated) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/fcgi_apps][%d] createFCGIAppCreated  %+v", 201, o.Payload)
}

func (o *CreateFCGIAppCreated) GetPayload() *models.FcgiApp {
	return o.Payload
}

func (o *CreateFCGIAppCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FcgiApp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFCGIAppAccepted creates a CreateFCGIAppAccepted with default headers values
func NewCreateFCGIAppAccepted() *CreateFCGIAppAccepted {
	return &CreateFCGIAppAccepted{}
}

/*
CreateFCGIAppAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateFCGIAppAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.FcgiApp
}

// IsSuccess returns true when this create f c g i app accepted response has a 2xx status code
func (o *CreateFCGIAppAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create f c g i app accepted response has a 3xx status code
func (o *CreateFCGIAppAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create f c g i app accepted response has a 4xx status code
func (o *CreateFCGIAppAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this create f c g i app accepted response has a 5xx status code
func (o *CreateFCGIAppAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this create f c g i app accepted response a status code equal to that given
func (o *CreateFCGIAppAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *CreateFCGIAppAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/fcgi_apps][%d] createFCGIAppAccepted  %+v", 202, o.Payload)
}

func (o *CreateFCGIAppAccepted) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/fcgi_apps][%d] createFCGIAppAccepted  %+v", 202, o.Payload)
}

func (o *CreateFCGIAppAccepted) GetPayload() *models.FcgiApp {
	return o.Payload
}

func (o *CreateFCGIAppAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.FcgiApp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFCGIAppBadRequest creates a CreateFCGIAppBadRequest with default headers values
func NewCreateFCGIAppBadRequest() *CreateFCGIAppBadRequest {
	return &CreateFCGIAppBadRequest{}
}

/*
CreateFCGIAppBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateFCGIAppBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create f c g i app bad request response has a 2xx status code
func (o *CreateFCGIAppBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create f c g i app bad request response has a 3xx status code
func (o *CreateFCGIAppBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create f c g i app bad request response has a 4xx status code
func (o *CreateFCGIAppBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create f c g i app bad request response has a 5xx status code
func (o *CreateFCGIAppBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create f c g i app bad request response a status code equal to that given
func (o *CreateFCGIAppBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateFCGIAppBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/fcgi_apps][%d] createFCGIAppBadRequest  %+v", 400, o.Payload)
}

func (o *CreateFCGIAppBadRequest) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/fcgi_apps][%d] createFCGIAppBadRequest  %+v", 400, o.Payload)
}

func (o *CreateFCGIAppBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateFCGIAppBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateFCGIAppConflict creates a CreateFCGIAppConflict with default headers values
func NewCreateFCGIAppConflict() *CreateFCGIAppConflict {
	return &CreateFCGIAppConflict{}
}

/*
CreateFCGIAppConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateFCGIAppConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create f c g i app conflict response has a 2xx status code
func (o *CreateFCGIAppConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create f c g i app conflict response has a 3xx status code
func (o *CreateFCGIAppConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create f c g i app conflict response has a 4xx status code
func (o *CreateFCGIAppConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create f c g i app conflict response has a 5xx status code
func (o *CreateFCGIAppConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create f c g i app conflict response a status code equal to that given
func (o *CreateFCGIAppConflict) IsCode(code int) bool {
	return code == 409
}

func (o *CreateFCGIAppConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/fcgi_apps][%d] createFCGIAppConflict  %+v", 409, o.Payload)
}

func (o *CreateFCGIAppConflict) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/fcgi_apps][%d] createFCGIAppConflict  %+v", 409, o.Payload)
}

func (o *CreateFCGIAppConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateFCGIAppConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateFCGIAppDefault creates a CreateFCGIAppDefault with default headers values
func NewCreateFCGIAppDefault(code int) *CreateFCGIAppDefault {
	return &CreateFCGIAppDefault{
		_statusCode: code,
	}
}

/*
CreateFCGIAppDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateFCGIAppDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create f c g i app default response
func (o *CreateFCGIAppDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create f c g i app default response has a 2xx status code
func (o *CreateFCGIAppDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create f c g i app default response has a 3xx status code
func (o *CreateFCGIAppDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create f c g i app default response has a 4xx status code
func (o *CreateFCGIAppDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create f c g i app default response has a 5xx status code
func (o *CreateFCGIAppDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create f c g i app default response a status code equal to that given
func (o *CreateFCGIAppDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateFCGIAppDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/fcgi_apps][%d] createFCGIApp default  %+v", o._statusCode, o.Payload)
}

func (o *CreateFCGIAppDefault) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/fcgi_apps][%d] createFCGIApp default  %+v", o._statusCode, o.Payload)
}

func (o *CreateFCGIAppDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateFCGIAppDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
