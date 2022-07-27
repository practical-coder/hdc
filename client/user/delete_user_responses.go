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

// DeleteUserReader is a Reader for the DeleteUser structure.
type DeleteUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteUserAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteUserAccepted creates a DeleteUserAccepted with default headers values
func NewDeleteUserAccepted() *DeleteUserAccepted {
	return &DeleteUserAccepted{}
}

/* DeleteUserAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeleteUserAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteUserAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/users/{username}][%d] deleteUserAccepted ", 202)
}

func (o *DeleteUserAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeleteUserNoContent creates a DeleteUserNoContent with default headers values
func NewDeleteUserNoContent() *DeleteUserNoContent {
	return &DeleteUserNoContent{}
}

/* DeleteUserNoContent describes a response with status code 204, with default header values.

User deleted
*/
type DeleteUserNoContent struct {
}

func (o *DeleteUserNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/users/{username}][%d] deleteUserNoContent ", 204)
}

func (o *DeleteUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserNotFound creates a DeleteUserNotFound with default headers values
func NewDeleteUserNotFound() *DeleteUserNotFound {
	return &DeleteUserNotFound{}
}

/* DeleteUserNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteUserNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteUserNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/users/{username}][%d] deleteUserNotFound  %+v", 404, o.Payload)
}
func (o *DeleteUserNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteUserDefault creates a DeleteUserDefault with default headers values
func NewDeleteUserDefault(code int) *DeleteUserDefault {
	return &DeleteUserDefault{
		_statusCode: code,
	}
}

/* DeleteUserDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteUserDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete user default response
func (o *DeleteUserDefault) Code() int {
	return o._statusCode
}

func (o *DeleteUserDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/users/{username}][%d] deleteUser default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteUserDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
