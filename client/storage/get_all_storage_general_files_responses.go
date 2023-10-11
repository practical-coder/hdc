// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v5/models"
)

// GetAllStorageGeneralFilesReader is a Reader for the GetAllStorageGeneralFiles structure.
type GetAllStorageGeneralFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllStorageGeneralFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllStorageGeneralFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAllStorageGeneralFilesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAllStorageGeneralFilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAllStorageGeneralFilesOK creates a GetAllStorageGeneralFilesOK with default headers values
func NewGetAllStorageGeneralFilesOK() *GetAllStorageGeneralFilesOK {
	return &GetAllStorageGeneralFilesOK{}
}

/*
GetAllStorageGeneralFilesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetAllStorageGeneralFilesOK struct {
	Payload models.GeneralFiles
}

// IsSuccess returns true when this get all storage general files o k response has a 2xx status code
func (o *GetAllStorageGeneralFilesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all storage general files o k response has a 3xx status code
func (o *GetAllStorageGeneralFilesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all storage general files o k response has a 4xx status code
func (o *GetAllStorageGeneralFilesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all storage general files o k response has a 5xx status code
func (o *GetAllStorageGeneralFilesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all storage general files o k response a status code equal to that given
func (o *GetAllStorageGeneralFilesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get all storage general files o k response
func (o *GetAllStorageGeneralFilesOK) Code() int {
	return 200
}

func (o *GetAllStorageGeneralFilesOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/storage/general][%d] getAllStorageGeneralFilesOK  %+v", 200, o.Payload)
}

func (o *GetAllStorageGeneralFilesOK) String() string {
	return fmt.Sprintf("[GET /services/haproxy/storage/general][%d] getAllStorageGeneralFilesOK  %+v", 200, o.Payload)
}

func (o *GetAllStorageGeneralFilesOK) GetPayload() models.GeneralFiles {
	return o.Payload
}

func (o *GetAllStorageGeneralFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllStorageGeneralFilesNotFound creates a GetAllStorageGeneralFilesNotFound with default headers values
func NewGetAllStorageGeneralFilesNotFound() *GetAllStorageGeneralFilesNotFound {
	return &GetAllStorageGeneralFilesNotFound{}
}

/*
GetAllStorageGeneralFilesNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetAllStorageGeneralFilesNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get all storage general files not found response has a 2xx status code
func (o *GetAllStorageGeneralFilesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all storage general files not found response has a 3xx status code
func (o *GetAllStorageGeneralFilesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all storage general files not found response has a 4xx status code
func (o *GetAllStorageGeneralFilesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all storage general files not found response has a 5xx status code
func (o *GetAllStorageGeneralFilesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get all storage general files not found response a status code equal to that given
func (o *GetAllStorageGeneralFilesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get all storage general files not found response
func (o *GetAllStorageGeneralFilesNotFound) Code() int {
	return 404
}

func (o *GetAllStorageGeneralFilesNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/storage/general][%d] getAllStorageGeneralFilesNotFound  %+v", 404, o.Payload)
}

func (o *GetAllStorageGeneralFilesNotFound) String() string {
	return fmt.Sprintf("[GET /services/haproxy/storage/general][%d] getAllStorageGeneralFilesNotFound  %+v", 404, o.Payload)
}

func (o *GetAllStorageGeneralFilesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllStorageGeneralFilesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAllStorageGeneralFilesDefault creates a GetAllStorageGeneralFilesDefault with default headers values
func NewGetAllStorageGeneralFilesDefault(code int) *GetAllStorageGeneralFilesDefault {
	return &GetAllStorageGeneralFilesDefault{
		_statusCode: code,
	}
}

/*
GetAllStorageGeneralFilesDefault describes a response with status code -1, with default header values.

General Error
*/
type GetAllStorageGeneralFilesDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get all storage general files default response has a 2xx status code
func (o *GetAllStorageGeneralFilesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get all storage general files default response has a 3xx status code
func (o *GetAllStorageGeneralFilesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get all storage general files default response has a 4xx status code
func (o *GetAllStorageGeneralFilesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get all storage general files default response has a 5xx status code
func (o *GetAllStorageGeneralFilesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get all storage general files default response a status code equal to that given
func (o *GetAllStorageGeneralFilesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get all storage general files default response
func (o *GetAllStorageGeneralFilesDefault) Code() int {
	return o._statusCode
}

func (o *GetAllStorageGeneralFilesDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/storage/general][%d] getAllStorageGeneralFiles default  %+v", o._statusCode, o.Payload)
}

func (o *GetAllStorageGeneralFilesDefault) String() string {
	return fmt.Sprintf("[GET /services/haproxy/storage/general][%d] getAllStorageGeneralFiles default  %+v", o._statusCode, o.Payload)
}

func (o *GetAllStorageGeneralFilesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllStorageGeneralFilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
