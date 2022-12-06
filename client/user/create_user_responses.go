// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// CreateUserReader is a Reader for the CreateUser structure.
type CreateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateUserCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateUserAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateUserConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateUserCreated creates a CreateUserCreated with default headers values
func NewCreateUserCreated() *CreateUserCreated {
	return &CreateUserCreated{}
}

/*
CreateUserCreated describes a response with status code 201, with default header values.

User created
*/
type CreateUserCreated struct {
	Payload *models.User
}

// IsSuccess returns true when this create user created response has a 2xx status code
func (o *CreateUserCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create user created response has a 3xx status code
func (o *CreateUserCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user created response has a 4xx status code
func (o *CreateUserCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create user created response has a 5xx status code
func (o *CreateUserCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create user created response a status code equal to that given
func (o *CreateUserCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateUserCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/users][%d] createUserCreated  %+v", 201, o.Payload)
}

func (o *CreateUserCreated) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/users][%d] createUserCreated  %+v", 201, o.Payload)
}

func (o *CreateUserCreated) GetPayload() *models.User {
	return o.Payload
}

func (o *CreateUserCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserAccepted creates a CreateUserAccepted with default headers values
func NewCreateUserAccepted() *CreateUserAccepted {
	return &CreateUserAccepted{}
}

/*
CreateUserAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateUserAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.User
}

// IsSuccess returns true when this create user accepted response has a 2xx status code
func (o *CreateUserAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create user accepted response has a 3xx status code
func (o *CreateUserAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user accepted response has a 4xx status code
func (o *CreateUserAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this create user accepted response has a 5xx status code
func (o *CreateUserAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this create user accepted response a status code equal to that given
func (o *CreateUserAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *CreateUserAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/users][%d] createUserAccepted  %+v", 202, o.Payload)
}

func (o *CreateUserAccepted) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/users][%d] createUserAccepted  %+v", 202, o.Payload)
}

func (o *CreateUserAccepted) GetPayload() *models.User {
	return o.Payload
}

func (o *CreateUserAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserBadRequest creates a CreateUserBadRequest with default headers values
func NewCreateUserBadRequest() *CreateUserBadRequest {
	return &CreateUserBadRequest{}
}

/*
CreateUserBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateUserBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create user bad request response has a 2xx status code
func (o *CreateUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user bad request response has a 3xx status code
func (o *CreateUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user bad request response has a 4xx status code
func (o *CreateUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user bad request response has a 5xx status code
func (o *CreateUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create user bad request response a status code equal to that given
func (o *CreateUserBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateUserBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/users][%d] createUserBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserBadRequest) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/users][%d] createUserBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateUserConflict creates a CreateUserConflict with default headers values
func NewCreateUserConflict() *CreateUserConflict {
	return &CreateUserConflict{}
}

/*
CreateUserConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateUserConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create user conflict response has a 2xx status code
func (o *CreateUserConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user conflict response has a 3xx status code
func (o *CreateUserConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user conflict response has a 4xx status code
func (o *CreateUserConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user conflict response has a 5xx status code
func (o *CreateUserConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create user conflict response a status code equal to that given
func (o *CreateUserConflict) IsCode(code int) bool {
	return code == 409
}

func (o *CreateUserConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/users][%d] createUserConflict  %+v", 409, o.Payload)
}

func (o *CreateUserConflict) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/users][%d] createUserConflict  %+v", 409, o.Payload)
}

func (o *CreateUserConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUserConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateUserDefault creates a CreateUserDefault with default headers values
func NewCreateUserDefault(code int) *CreateUserDefault {
	return &CreateUserDefault{
		_statusCode: code,
	}
}

/*
CreateUserDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateUserDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create user default response
func (o *CreateUserDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create user default response has a 2xx status code
func (o *CreateUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create user default response has a 3xx status code
func (o *CreateUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create user default response has a 4xx status code
func (o *CreateUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create user default response has a 5xx status code
func (o *CreateUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create user default response a status code equal to that given
func (o *CreateUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateUserDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/users][%d] createUser default  %+v", o._statusCode, o.Payload)
}

func (o *CreateUserDefault) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/users][%d] createUser default  %+v", o._statusCode, o.Payload)
}

func (o *CreateUserDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
