// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v5/models"
)

// ReplaceUserReader is a Reader for the ReplaceUser structure.
type ReplaceUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceUserAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceUserOK creates a ReplaceUserOK with default headers values
func NewReplaceUserOK() *ReplaceUserOK {
	return &ReplaceUserOK{}
}

/*
ReplaceUserOK describes a response with status code 200, with default header values.

User replaced
*/
type ReplaceUserOK struct {
	Payload *models.User
}

// IsSuccess returns true when this replace user o k response has a 2xx status code
func (o *ReplaceUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this replace user o k response has a 3xx status code
func (o *ReplaceUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace user o k response has a 4xx status code
func (o *ReplaceUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this replace user o k response has a 5xx status code
func (o *ReplaceUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this replace user o k response a status code equal to that given
func (o *ReplaceUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the replace user o k response
func (o *ReplaceUserOK) Code() int {
	return 200
}

func (o *ReplaceUserOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/users/{username}][%d] replaceUserOK  %+v", 200, o.Payload)
}

func (o *ReplaceUserOK) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/users/{username}][%d] replaceUserOK  %+v", 200, o.Payload)
}

func (o *ReplaceUserOK) GetPayload() *models.User {
	return o.Payload
}

func (o *ReplaceUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceUserAccepted creates a ReplaceUserAccepted with default headers values
func NewReplaceUserAccepted() *ReplaceUserAccepted {
	return &ReplaceUserAccepted{}
}

/*
ReplaceUserAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type ReplaceUserAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.User
}

// IsSuccess returns true when this replace user accepted response has a 2xx status code
func (o *ReplaceUserAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this replace user accepted response has a 3xx status code
func (o *ReplaceUserAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace user accepted response has a 4xx status code
func (o *ReplaceUserAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this replace user accepted response has a 5xx status code
func (o *ReplaceUserAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this replace user accepted response a status code equal to that given
func (o *ReplaceUserAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the replace user accepted response
func (o *ReplaceUserAccepted) Code() int {
	return 202
}

func (o *ReplaceUserAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/users/{username}][%d] replaceUserAccepted  %+v", 202, o.Payload)
}

func (o *ReplaceUserAccepted) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/users/{username}][%d] replaceUserAccepted  %+v", 202, o.Payload)
}

func (o *ReplaceUserAccepted) GetPayload() *models.User {
	return o.Payload
}

func (o *ReplaceUserAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceUserBadRequest creates a ReplaceUserBadRequest with default headers values
func NewReplaceUserBadRequest() *ReplaceUserBadRequest {
	return &ReplaceUserBadRequest{}
}

/*
ReplaceUserBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceUserBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this replace user bad request response has a 2xx status code
func (o *ReplaceUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this replace user bad request response has a 3xx status code
func (o *ReplaceUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace user bad request response has a 4xx status code
func (o *ReplaceUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this replace user bad request response has a 5xx status code
func (o *ReplaceUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this replace user bad request response a status code equal to that given
func (o *ReplaceUserBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the replace user bad request response
func (o *ReplaceUserBadRequest) Code() int {
	return 400
}

func (o *ReplaceUserBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/users/{username}][%d] replaceUserBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceUserBadRequest) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/users/{username}][%d] replaceUserBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceUserBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceUserNotFound creates a ReplaceUserNotFound with default headers values
func NewReplaceUserNotFound() *ReplaceUserNotFound {
	return &ReplaceUserNotFound{}
}

/*
ReplaceUserNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type ReplaceUserNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this replace user not found response has a 2xx status code
func (o *ReplaceUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this replace user not found response has a 3xx status code
func (o *ReplaceUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace user not found response has a 4xx status code
func (o *ReplaceUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this replace user not found response has a 5xx status code
func (o *ReplaceUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this replace user not found response a status code equal to that given
func (o *ReplaceUserNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the replace user not found response
func (o *ReplaceUserNotFound) Code() int {
	return 404
}

func (o *ReplaceUserNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/users/{username}][%d] replaceUserNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceUserNotFound) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/users/{username}][%d] replaceUserNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceUserNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceUserDefault creates a ReplaceUserDefault with default headers values
func NewReplaceUserDefault(code int) *ReplaceUserDefault {
	return &ReplaceUserDefault{
		_statusCode: code,
	}
}

/*
ReplaceUserDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceUserDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this replace user default response has a 2xx status code
func (o *ReplaceUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this replace user default response has a 3xx status code
func (o *ReplaceUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this replace user default response has a 4xx status code
func (o *ReplaceUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this replace user default response has a 5xx status code
func (o *ReplaceUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this replace user default response a status code equal to that given
func (o *ReplaceUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the replace user default response
func (o *ReplaceUserDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceUserDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/users/{username}][%d] replaceUser default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceUserDefault) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/users/{username}][%d] replaceUser default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceUserDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
