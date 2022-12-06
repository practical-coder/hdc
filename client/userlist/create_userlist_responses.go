// Code generated by go-swagger; DO NOT EDIT.

package userlist

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// CreateUserlistReader is a Reader for the CreateUserlist structure.
type CreateUserlistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserlistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateUserlistCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateUserlistAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateUserlistBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateUserlistConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateUserlistDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateUserlistCreated creates a CreateUserlistCreated with default headers values
func NewCreateUserlistCreated() *CreateUserlistCreated {
	return &CreateUserlistCreated{}
}

/*
CreateUserlistCreated describes a response with status code 201, with default header values.

Userlist created
*/
type CreateUserlistCreated struct {
	Payload *models.Userlist
}

// IsSuccess returns true when this create userlist created response has a 2xx status code
func (o *CreateUserlistCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create userlist created response has a 3xx status code
func (o *CreateUserlistCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create userlist created response has a 4xx status code
func (o *CreateUserlistCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create userlist created response has a 5xx status code
func (o *CreateUserlistCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create userlist created response a status code equal to that given
func (o *CreateUserlistCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateUserlistCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/userlists][%d] createUserlistCreated  %+v", 201, o.Payload)
}

func (o *CreateUserlistCreated) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/userlists][%d] createUserlistCreated  %+v", 201, o.Payload)
}

func (o *CreateUserlistCreated) GetPayload() *models.Userlist {
	return o.Payload
}

func (o *CreateUserlistCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Userlist)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserlistAccepted creates a CreateUserlistAccepted with default headers values
func NewCreateUserlistAccepted() *CreateUserlistAccepted {
	return &CreateUserlistAccepted{}
}

/*
CreateUserlistAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateUserlistAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.Userlist
}

// IsSuccess returns true when this create userlist accepted response has a 2xx status code
func (o *CreateUserlistAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create userlist accepted response has a 3xx status code
func (o *CreateUserlistAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create userlist accepted response has a 4xx status code
func (o *CreateUserlistAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this create userlist accepted response has a 5xx status code
func (o *CreateUserlistAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this create userlist accepted response a status code equal to that given
func (o *CreateUserlistAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *CreateUserlistAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/userlists][%d] createUserlistAccepted  %+v", 202, o.Payload)
}

func (o *CreateUserlistAccepted) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/userlists][%d] createUserlistAccepted  %+v", 202, o.Payload)
}

func (o *CreateUserlistAccepted) GetPayload() *models.Userlist {
	return o.Payload
}

func (o *CreateUserlistAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.Userlist)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserlistBadRequest creates a CreateUserlistBadRequest with default headers values
func NewCreateUserlistBadRequest() *CreateUserlistBadRequest {
	return &CreateUserlistBadRequest{}
}

/*
CreateUserlistBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateUserlistBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create userlist bad request response has a 2xx status code
func (o *CreateUserlistBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create userlist bad request response has a 3xx status code
func (o *CreateUserlistBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create userlist bad request response has a 4xx status code
func (o *CreateUserlistBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create userlist bad request response has a 5xx status code
func (o *CreateUserlistBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create userlist bad request response a status code equal to that given
func (o *CreateUserlistBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateUserlistBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/userlists][%d] createUserlistBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserlistBadRequest) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/userlists][%d] createUserlistBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserlistBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUserlistBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateUserlistConflict creates a CreateUserlistConflict with default headers values
func NewCreateUserlistConflict() *CreateUserlistConflict {
	return &CreateUserlistConflict{}
}

/*
CreateUserlistConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateUserlistConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create userlist conflict response has a 2xx status code
func (o *CreateUserlistConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create userlist conflict response has a 3xx status code
func (o *CreateUserlistConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create userlist conflict response has a 4xx status code
func (o *CreateUserlistConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create userlist conflict response has a 5xx status code
func (o *CreateUserlistConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create userlist conflict response a status code equal to that given
func (o *CreateUserlistConflict) IsCode(code int) bool {
	return code == 409
}

func (o *CreateUserlistConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/userlists][%d] createUserlistConflict  %+v", 409, o.Payload)
}

func (o *CreateUserlistConflict) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/userlists][%d] createUserlistConflict  %+v", 409, o.Payload)
}

func (o *CreateUserlistConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUserlistConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateUserlistDefault creates a CreateUserlistDefault with default headers values
func NewCreateUserlistDefault(code int) *CreateUserlistDefault {
	return &CreateUserlistDefault{
		_statusCode: code,
	}
}

/*
CreateUserlistDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateUserlistDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create userlist default response
func (o *CreateUserlistDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create userlist default response has a 2xx status code
func (o *CreateUserlistDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create userlist default response has a 3xx status code
func (o *CreateUserlistDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create userlist default response has a 4xx status code
func (o *CreateUserlistDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create userlist default response has a 5xx status code
func (o *CreateUserlistDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create userlist default response a status code equal to that given
func (o *CreateUserlistDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateUserlistDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/userlists][%d] createUserlist default  %+v", o._statusCode, o.Payload)
}

func (o *CreateUserlistDefault) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/userlists][%d] createUserlist default  %+v", o._statusCode, o.Payload)
}

func (o *CreateUserlistDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUserlistDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
