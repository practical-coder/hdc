// Code generated by go-swagger; DO NOT EDIT.

package userlist

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v3/models"
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

/*CreateUserlistCreated handles this case with default header values.

Userlist created
*/
type CreateUserlistCreated struct {
	Payload *models.Userlist
}

func (o *CreateUserlistCreated) Error() string {
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

/*CreateUserlistAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type CreateUserlistAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string

	Payload *models.Userlist
}

func (o *CreateUserlistAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/userlists][%d] createUserlistAccepted  %+v", 202, o.Payload)
}

func (o *CreateUserlistAccepted) GetPayload() *models.Userlist {
	return o.Payload
}

func (o *CreateUserlistAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

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

/*CreateUserlistBadRequest handles this case with default header values.

Bad request
*/
type CreateUserlistBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateUserlistBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/userlists][%d] createUserlistBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserlistBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUserlistBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

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

/*CreateUserlistConflict handles this case with default header values.

The specified resource already exists
*/
type CreateUserlistConflict struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateUserlistConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/userlists][%d] createUserlistConflict  %+v", 409, o.Payload)
}

func (o *CreateUserlistConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUserlistConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

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

/*CreateUserlistDefault handles this case with default header values.

General Error
*/
type CreateUserlistDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create userlist default response
func (o *CreateUserlistDefault) Code() int {
	return o._statusCode
}

func (o *CreateUserlistDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/userlists][%d] createUserlist default  %+v", o._statusCode, o.Payload)
}

func (o *CreateUserlistDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUserlistDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}