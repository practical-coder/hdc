// Code generated by go-swagger; DO NOT EDIT.

package table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v5/models"
)

// CreateTableReader is a Reader for the CreateTable structure.
type CreateTableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateTableCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateTableAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTableBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateTableConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateTableDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateTableCreated creates a CreateTableCreated with default headers values
func NewCreateTableCreated() *CreateTableCreated {
	return &CreateTableCreated{}
}

/*
CreateTableCreated describes a response with status code 201, with default header values.

Table created
*/
type CreateTableCreated struct {
	Payload *models.Table
}

// IsSuccess returns true when this create table created response has a 2xx status code
func (o *CreateTableCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create table created response has a 3xx status code
func (o *CreateTableCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create table created response has a 4xx status code
func (o *CreateTableCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create table created response has a 5xx status code
func (o *CreateTableCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create table created response a status code equal to that given
func (o *CreateTableCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create table created response
func (o *CreateTableCreated) Code() int {
	return 201
}

func (o *CreateTableCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/tables][%d] createTableCreated  %+v", 201, o.Payload)
}

func (o *CreateTableCreated) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/tables][%d] createTableCreated  %+v", 201, o.Payload)
}

func (o *CreateTableCreated) GetPayload() *models.Table {
	return o.Payload
}

func (o *CreateTableCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Table)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTableAccepted creates a CreateTableAccepted with default headers values
func NewCreateTableAccepted() *CreateTableAccepted {
	return &CreateTableAccepted{}
}

/*
CreateTableAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateTableAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.Table
}

// IsSuccess returns true when this create table accepted response has a 2xx status code
func (o *CreateTableAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create table accepted response has a 3xx status code
func (o *CreateTableAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create table accepted response has a 4xx status code
func (o *CreateTableAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this create table accepted response has a 5xx status code
func (o *CreateTableAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this create table accepted response a status code equal to that given
func (o *CreateTableAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the create table accepted response
func (o *CreateTableAccepted) Code() int {
	return 202
}

func (o *CreateTableAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/tables][%d] createTableAccepted  %+v", 202, o.Payload)
}

func (o *CreateTableAccepted) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/tables][%d] createTableAccepted  %+v", 202, o.Payload)
}

func (o *CreateTableAccepted) GetPayload() *models.Table {
	return o.Payload
}

func (o *CreateTableAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.Table)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTableBadRequest creates a CreateTableBadRequest with default headers values
func NewCreateTableBadRequest() *CreateTableBadRequest {
	return &CreateTableBadRequest{}
}

/*
CreateTableBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateTableBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create table bad request response has a 2xx status code
func (o *CreateTableBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create table bad request response has a 3xx status code
func (o *CreateTableBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create table bad request response has a 4xx status code
func (o *CreateTableBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create table bad request response has a 5xx status code
func (o *CreateTableBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create table bad request response a status code equal to that given
func (o *CreateTableBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create table bad request response
func (o *CreateTableBadRequest) Code() int {
	return 400
}

func (o *CreateTableBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/tables][%d] createTableBadRequest  %+v", 400, o.Payload)
}

func (o *CreateTableBadRequest) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/tables][%d] createTableBadRequest  %+v", 400, o.Payload)
}

func (o *CreateTableBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateTableBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateTableConflict creates a CreateTableConflict with default headers values
func NewCreateTableConflict() *CreateTableConflict {
	return &CreateTableConflict{}
}

/*
CreateTableConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateTableConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create table conflict response has a 2xx status code
func (o *CreateTableConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create table conflict response has a 3xx status code
func (o *CreateTableConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create table conflict response has a 4xx status code
func (o *CreateTableConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create table conflict response has a 5xx status code
func (o *CreateTableConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create table conflict response a status code equal to that given
func (o *CreateTableConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create table conflict response
func (o *CreateTableConflict) Code() int {
	return 409
}

func (o *CreateTableConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/tables][%d] createTableConflict  %+v", 409, o.Payload)
}

func (o *CreateTableConflict) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/tables][%d] createTableConflict  %+v", 409, o.Payload)
}

func (o *CreateTableConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateTableConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateTableDefault creates a CreateTableDefault with default headers values
func NewCreateTableDefault(code int) *CreateTableDefault {
	return &CreateTableDefault{
		_statusCode: code,
	}
}

/*
CreateTableDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateTableDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create table default response has a 2xx status code
func (o *CreateTableDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create table default response has a 3xx status code
func (o *CreateTableDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create table default response has a 4xx status code
func (o *CreateTableDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create table default response has a 5xx status code
func (o *CreateTableDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create table default response a status code equal to that given
func (o *CreateTableDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create table default response
func (o *CreateTableDefault) Code() int {
	return o._statusCode
}

func (o *CreateTableDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/tables][%d] createTable default  %+v", o._statusCode, o.Payload)
}

func (o *CreateTableDefault) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/tables][%d] createTable default  %+v", o._statusCode, o.Payload)
}

func (o *CreateTableDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateTableDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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