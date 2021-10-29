// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/models"
)

// GetOneStorageMapReader is a Reader for the GetOneStorageMap structure.
type GetOneStorageMapReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetOneStorageMapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOneStorageMapOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetOneStorageMapNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetOneStorageMapDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOneStorageMapOK creates a GetOneStorageMapOK with default headers values
func NewGetOneStorageMapOK(writer io.Writer) *GetOneStorageMapOK {
	return &GetOneStorageMapOK{

		Payload: writer,
	}
}

/* GetOneStorageMapOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOneStorageMapOK struct {
	Payload io.Writer
}

func (o *GetOneStorageMapOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/storage/maps/{name}][%d] getOneStorageMapOK  %+v", 200, o.Payload)
}
func (o *GetOneStorageMapOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *GetOneStorageMapOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOneStorageMapNotFound creates a GetOneStorageMapNotFound with default headers values
func NewGetOneStorageMapNotFound() *GetOneStorageMapNotFound {
	return &GetOneStorageMapNotFound{}
}

/* GetOneStorageMapNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetOneStorageMapNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetOneStorageMapNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/storage/maps/{name}][%d] getOneStorageMapNotFound  %+v", 404, o.Payload)
}
func (o *GetOneStorageMapNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOneStorageMapNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetOneStorageMapDefault creates a GetOneStorageMapDefault with default headers values
func NewGetOneStorageMapDefault(code int) *GetOneStorageMapDefault {
	return &GetOneStorageMapDefault{
		_statusCode: code,
	}
}

/* GetOneStorageMapDefault describes a response with status code -1, with default header values.

General Error
*/
type GetOneStorageMapDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get one storage map default response
func (o *GetOneStorageMapDefault) Code() int {
	return o._statusCode
}

func (o *GetOneStorageMapDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/storage/maps/{name}][%d] getOneStorageMap default  %+v", o._statusCode, o.Payload)
}
func (o *GetOneStorageMapDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOneStorageMapDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
