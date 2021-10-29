// Code generated by go-swagger; DO NOT EDIT.

package maps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/models"
)

// GetAllRuntimeMapFilesReader is a Reader for the GetAllRuntimeMapFiles structure.
type GetAllRuntimeMapFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllRuntimeMapFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllRuntimeMapFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAllRuntimeMapFilesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAllRuntimeMapFilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAllRuntimeMapFilesOK creates a GetAllRuntimeMapFilesOK with default headers values
func NewGetAllRuntimeMapFilesOK() *GetAllRuntimeMapFilesOK {
	return &GetAllRuntimeMapFilesOK{}
}

/* GetAllRuntimeMapFilesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetAllRuntimeMapFilesOK struct {
	Payload models.Maps
}

func (o *GetAllRuntimeMapFilesOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/maps][%d] getAllRuntimeMapFilesOK  %+v", 200, o.Payload)
}
func (o *GetAllRuntimeMapFilesOK) GetPayload() models.Maps {
	return o.Payload
}

func (o *GetAllRuntimeMapFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllRuntimeMapFilesNotFound creates a GetAllRuntimeMapFilesNotFound with default headers values
func NewGetAllRuntimeMapFilesNotFound() *GetAllRuntimeMapFilesNotFound {
	return &GetAllRuntimeMapFilesNotFound{}
}

/* GetAllRuntimeMapFilesNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetAllRuntimeMapFilesNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetAllRuntimeMapFilesNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/maps][%d] getAllRuntimeMapFilesNotFound  %+v", 404, o.Payload)
}
func (o *GetAllRuntimeMapFilesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllRuntimeMapFilesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAllRuntimeMapFilesDefault creates a GetAllRuntimeMapFilesDefault with default headers values
func NewGetAllRuntimeMapFilesDefault(code int) *GetAllRuntimeMapFilesDefault {
	return &GetAllRuntimeMapFilesDefault{
		_statusCode: code,
	}
}

/* GetAllRuntimeMapFilesDefault describes a response with status code -1, with default header values.

General Error
*/
type GetAllRuntimeMapFilesDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get all runtime map files default response
func (o *GetAllRuntimeMapFilesDefault) Code() int {
	return o._statusCode
}

func (o *GetAllRuntimeMapFilesDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/maps][%d] getAllRuntimeMapFiles default  %+v", o._statusCode, o.Payload)
}
func (o *GetAllRuntimeMapFilesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllRuntimeMapFilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
