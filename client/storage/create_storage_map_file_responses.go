// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// CreateStorageMapFileReader is a Reader for the CreateStorageMapFile structure.
type CreateStorageMapFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateStorageMapFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateStorageMapFileCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateStorageMapFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateStorageMapFileConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateStorageMapFileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateStorageMapFileCreated creates a CreateStorageMapFileCreated with default headers values
func NewCreateStorageMapFileCreated() *CreateStorageMapFileCreated {
	return &CreateStorageMapFileCreated{}
}

/*
CreateStorageMapFileCreated describes a response with status code 201, with default header values.

Map file created with its entries
*/
type CreateStorageMapFileCreated struct {
	Payload *models.Map
}

// IsSuccess returns true when this create storage map file created response has a 2xx status code
func (o *CreateStorageMapFileCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create storage map file created response has a 3xx status code
func (o *CreateStorageMapFileCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create storage map file created response has a 4xx status code
func (o *CreateStorageMapFileCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create storage map file created response has a 5xx status code
func (o *CreateStorageMapFileCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create storage map file created response a status code equal to that given
func (o *CreateStorageMapFileCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateStorageMapFileCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/storage/maps][%d] createStorageMapFileCreated  %+v", 201, o.Payload)
}

func (o *CreateStorageMapFileCreated) String() string {
	return fmt.Sprintf("[POST /services/haproxy/storage/maps][%d] createStorageMapFileCreated  %+v", 201, o.Payload)
}

func (o *CreateStorageMapFileCreated) GetPayload() *models.Map {
	return o.Payload
}

func (o *CreateStorageMapFileCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Map)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStorageMapFileBadRequest creates a CreateStorageMapFileBadRequest with default headers values
func NewCreateStorageMapFileBadRequest() *CreateStorageMapFileBadRequest {
	return &CreateStorageMapFileBadRequest{}
}

/*
CreateStorageMapFileBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateStorageMapFileBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create storage map file bad request response has a 2xx status code
func (o *CreateStorageMapFileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create storage map file bad request response has a 3xx status code
func (o *CreateStorageMapFileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create storage map file bad request response has a 4xx status code
func (o *CreateStorageMapFileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create storage map file bad request response has a 5xx status code
func (o *CreateStorageMapFileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create storage map file bad request response a status code equal to that given
func (o *CreateStorageMapFileBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateStorageMapFileBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/storage/maps][%d] createStorageMapFileBadRequest  %+v", 400, o.Payload)
}

func (o *CreateStorageMapFileBadRequest) String() string {
	return fmt.Sprintf("[POST /services/haproxy/storage/maps][%d] createStorageMapFileBadRequest  %+v", 400, o.Payload)
}

func (o *CreateStorageMapFileBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateStorageMapFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateStorageMapFileConflict creates a CreateStorageMapFileConflict with default headers values
func NewCreateStorageMapFileConflict() *CreateStorageMapFileConflict {
	return &CreateStorageMapFileConflict{}
}

/*
CreateStorageMapFileConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateStorageMapFileConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create storage map file conflict response has a 2xx status code
func (o *CreateStorageMapFileConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create storage map file conflict response has a 3xx status code
func (o *CreateStorageMapFileConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create storage map file conflict response has a 4xx status code
func (o *CreateStorageMapFileConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create storage map file conflict response has a 5xx status code
func (o *CreateStorageMapFileConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create storage map file conflict response a status code equal to that given
func (o *CreateStorageMapFileConflict) IsCode(code int) bool {
	return code == 409
}

func (o *CreateStorageMapFileConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/storage/maps][%d] createStorageMapFileConflict  %+v", 409, o.Payload)
}

func (o *CreateStorageMapFileConflict) String() string {
	return fmt.Sprintf("[POST /services/haproxy/storage/maps][%d] createStorageMapFileConflict  %+v", 409, o.Payload)
}

func (o *CreateStorageMapFileConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateStorageMapFileConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateStorageMapFileDefault creates a CreateStorageMapFileDefault with default headers values
func NewCreateStorageMapFileDefault(code int) *CreateStorageMapFileDefault {
	return &CreateStorageMapFileDefault{
		_statusCode: code,
	}
}

/*
CreateStorageMapFileDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateStorageMapFileDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create storage map file default response
func (o *CreateStorageMapFileDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create storage map file default response has a 2xx status code
func (o *CreateStorageMapFileDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create storage map file default response has a 3xx status code
func (o *CreateStorageMapFileDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create storage map file default response has a 4xx status code
func (o *CreateStorageMapFileDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create storage map file default response has a 5xx status code
func (o *CreateStorageMapFileDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create storage map file default response a status code equal to that given
func (o *CreateStorageMapFileDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateStorageMapFileDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/storage/maps][%d] createStorageMapFile default  %+v", o._statusCode, o.Payload)
}

func (o *CreateStorageMapFileDefault) String() string {
	return fmt.Sprintf("[POST /services/haproxy/storage/maps][%d] createStorageMapFile default  %+v", o._statusCode, o.Payload)
}

func (o *CreateStorageMapFileDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateStorageMapFileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
