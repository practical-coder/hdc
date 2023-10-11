// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v5/models"
)

// CreateGroupReader is a Reader for the CreateGroup structure.
type CreateGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateGroupCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateGroupAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateGroupCreated creates a CreateGroupCreated with default headers values
func NewCreateGroupCreated() *CreateGroupCreated {
	return &CreateGroupCreated{}
}

/*
CreateGroupCreated describes a response with status code 201, with default header values.

Group created
*/
type CreateGroupCreated struct {
	Payload *models.Group
}

// IsSuccess returns true when this create group created response has a 2xx status code
func (o *CreateGroupCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create group created response has a 3xx status code
func (o *CreateGroupCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create group created response has a 4xx status code
func (o *CreateGroupCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create group created response has a 5xx status code
func (o *CreateGroupCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create group created response a status code equal to that given
func (o *CreateGroupCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create group created response
func (o *CreateGroupCreated) Code() int {
	return 201
}

func (o *CreateGroupCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/groups][%d] createGroupCreated  %+v", 201, o.Payload)
}

func (o *CreateGroupCreated) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/groups][%d] createGroupCreated  %+v", 201, o.Payload)
}

func (o *CreateGroupCreated) GetPayload() *models.Group {
	return o.Payload
}

func (o *CreateGroupCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGroupAccepted creates a CreateGroupAccepted with default headers values
func NewCreateGroupAccepted() *CreateGroupAccepted {
	return &CreateGroupAccepted{}
}

/*
CreateGroupAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateGroupAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.Group
}

// IsSuccess returns true when this create group accepted response has a 2xx status code
func (o *CreateGroupAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create group accepted response has a 3xx status code
func (o *CreateGroupAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create group accepted response has a 4xx status code
func (o *CreateGroupAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this create group accepted response has a 5xx status code
func (o *CreateGroupAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this create group accepted response a status code equal to that given
func (o *CreateGroupAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the create group accepted response
func (o *CreateGroupAccepted) Code() int {
	return 202
}

func (o *CreateGroupAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/groups][%d] createGroupAccepted  %+v", 202, o.Payload)
}

func (o *CreateGroupAccepted) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/groups][%d] createGroupAccepted  %+v", 202, o.Payload)
}

func (o *CreateGroupAccepted) GetPayload() *models.Group {
	return o.Payload
}

func (o *CreateGroupAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGroupBadRequest creates a CreateGroupBadRequest with default headers values
func NewCreateGroupBadRequest() *CreateGroupBadRequest {
	return &CreateGroupBadRequest{}
}

/*
CreateGroupBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateGroupBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create group bad request response has a 2xx status code
func (o *CreateGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create group bad request response has a 3xx status code
func (o *CreateGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create group bad request response has a 4xx status code
func (o *CreateGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create group bad request response has a 5xx status code
func (o *CreateGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create group bad request response a status code equal to that given
func (o *CreateGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create group bad request response
func (o *CreateGroupBadRequest) Code() int {
	return 400
}

func (o *CreateGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/groups][%d] createGroupBadRequest  %+v", 400, o.Payload)
}

func (o *CreateGroupBadRequest) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/groups][%d] createGroupBadRequest  %+v", 400, o.Payload)
}

func (o *CreateGroupBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateGroupConflict creates a CreateGroupConflict with default headers values
func NewCreateGroupConflict() *CreateGroupConflict {
	return &CreateGroupConflict{}
}

/*
CreateGroupConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateGroupConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create group conflict response has a 2xx status code
func (o *CreateGroupConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create group conflict response has a 3xx status code
func (o *CreateGroupConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create group conflict response has a 4xx status code
func (o *CreateGroupConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create group conflict response has a 5xx status code
func (o *CreateGroupConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create group conflict response a status code equal to that given
func (o *CreateGroupConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create group conflict response
func (o *CreateGroupConflict) Code() int {
	return 409
}

func (o *CreateGroupConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/groups][%d] createGroupConflict  %+v", 409, o.Payload)
}

func (o *CreateGroupConflict) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/groups][%d] createGroupConflict  %+v", 409, o.Payload)
}

func (o *CreateGroupConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateGroupDefault creates a CreateGroupDefault with default headers values
func NewCreateGroupDefault(code int) *CreateGroupDefault {
	return &CreateGroupDefault{
		_statusCode: code,
	}
}

/*
CreateGroupDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateGroupDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create group default response has a 2xx status code
func (o *CreateGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create group default response has a 3xx status code
func (o *CreateGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create group default response has a 4xx status code
func (o *CreateGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create group default response has a 5xx status code
func (o *CreateGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create group default response a status code equal to that given
func (o *CreateGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create group default response
func (o *CreateGroupDefault) Code() int {
	return o._statusCode
}

func (o *CreateGroupDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/groups][%d] createGroup default  %+v", o._statusCode, o.Payload)
}

func (o *CreateGroupDefault) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/groups][%d] createGroup default  %+v", o._statusCode, o.Payload)
}

func (o *CreateGroupDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
