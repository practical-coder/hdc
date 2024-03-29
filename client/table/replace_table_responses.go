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

// ReplaceTableReader is a Reader for the ReplaceTable structure.
type ReplaceTableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceTableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceTableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceTableAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceTableBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceTableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceTableDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceTableOK creates a ReplaceTableOK with default headers values
func NewReplaceTableOK() *ReplaceTableOK {
	return &ReplaceTableOK{}
}

/*
ReplaceTableOK describes a response with status code 200, with default header values.

Table replaced
*/
type ReplaceTableOK struct {
	Payload *models.Table
}

// IsSuccess returns true when this replace table o k response has a 2xx status code
func (o *ReplaceTableOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this replace table o k response has a 3xx status code
func (o *ReplaceTableOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace table o k response has a 4xx status code
func (o *ReplaceTableOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this replace table o k response has a 5xx status code
func (o *ReplaceTableOK) IsServerError() bool {
	return false
}

// IsCode returns true when this replace table o k response a status code equal to that given
func (o *ReplaceTableOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the replace table o k response
func (o *ReplaceTableOK) Code() int {
	return 200
}

func (o *ReplaceTableOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/tables/{name}][%d] replaceTableOK  %+v", 200, o.Payload)
}

func (o *ReplaceTableOK) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/tables/{name}][%d] replaceTableOK  %+v", 200, o.Payload)
}

func (o *ReplaceTableOK) GetPayload() *models.Table {
	return o.Payload
}

func (o *ReplaceTableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Table)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceTableAccepted creates a ReplaceTableAccepted with default headers values
func NewReplaceTableAccepted() *ReplaceTableAccepted {
	return &ReplaceTableAccepted{}
}

/*
ReplaceTableAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type ReplaceTableAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.Table
}

// IsSuccess returns true when this replace table accepted response has a 2xx status code
func (o *ReplaceTableAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this replace table accepted response has a 3xx status code
func (o *ReplaceTableAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace table accepted response has a 4xx status code
func (o *ReplaceTableAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this replace table accepted response has a 5xx status code
func (o *ReplaceTableAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this replace table accepted response a status code equal to that given
func (o *ReplaceTableAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the replace table accepted response
func (o *ReplaceTableAccepted) Code() int {
	return 202
}

func (o *ReplaceTableAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/tables/{name}][%d] replaceTableAccepted  %+v", 202, o.Payload)
}

func (o *ReplaceTableAccepted) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/tables/{name}][%d] replaceTableAccepted  %+v", 202, o.Payload)
}

func (o *ReplaceTableAccepted) GetPayload() *models.Table {
	return o.Payload
}

func (o *ReplaceTableAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceTableBadRequest creates a ReplaceTableBadRequest with default headers values
func NewReplaceTableBadRequest() *ReplaceTableBadRequest {
	return &ReplaceTableBadRequest{}
}

/*
ReplaceTableBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceTableBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this replace table bad request response has a 2xx status code
func (o *ReplaceTableBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this replace table bad request response has a 3xx status code
func (o *ReplaceTableBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace table bad request response has a 4xx status code
func (o *ReplaceTableBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this replace table bad request response has a 5xx status code
func (o *ReplaceTableBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this replace table bad request response a status code equal to that given
func (o *ReplaceTableBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the replace table bad request response
func (o *ReplaceTableBadRequest) Code() int {
	return 400
}

func (o *ReplaceTableBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/tables/{name}][%d] replaceTableBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceTableBadRequest) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/tables/{name}][%d] replaceTableBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceTableBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceTableBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceTableNotFound creates a ReplaceTableNotFound with default headers values
func NewReplaceTableNotFound() *ReplaceTableNotFound {
	return &ReplaceTableNotFound{}
}

/*
ReplaceTableNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type ReplaceTableNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this replace table not found response has a 2xx status code
func (o *ReplaceTableNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this replace table not found response has a 3xx status code
func (o *ReplaceTableNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace table not found response has a 4xx status code
func (o *ReplaceTableNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this replace table not found response has a 5xx status code
func (o *ReplaceTableNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this replace table not found response a status code equal to that given
func (o *ReplaceTableNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the replace table not found response
func (o *ReplaceTableNotFound) Code() int {
	return 404
}

func (o *ReplaceTableNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/tables/{name}][%d] replaceTableNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceTableNotFound) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/tables/{name}][%d] replaceTableNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceTableNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceTableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceTableDefault creates a ReplaceTableDefault with default headers values
func NewReplaceTableDefault(code int) *ReplaceTableDefault {
	return &ReplaceTableDefault{
		_statusCode: code,
	}
}

/*
ReplaceTableDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceTableDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this replace table default response has a 2xx status code
func (o *ReplaceTableDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this replace table default response has a 3xx status code
func (o *ReplaceTableDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this replace table default response has a 4xx status code
func (o *ReplaceTableDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this replace table default response has a 5xx status code
func (o *ReplaceTableDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this replace table default response a status code equal to that given
func (o *ReplaceTableDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the replace table default response
func (o *ReplaceTableDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceTableDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/tables/{name}][%d] replaceTable default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceTableDefault) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/tables/{name}][%d] replaceTable default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceTableDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceTableDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
