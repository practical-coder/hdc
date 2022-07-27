// Code generated by go-swagger; DO NOT EDIT.

package reloads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// GetReloadReader is a Reader for the GetReload structure.
type GetReloadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReloadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReloadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetReloadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetReloadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetReloadOK creates a GetReloadOK with default headers values
func NewGetReloadOK() *GetReloadOK {
	return &GetReloadOK{}
}

/* GetReloadOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetReloadOK struct {
	Payload *models.Reload
}

func (o *GetReloadOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/reloads/{id}][%d] getReloadOK  %+v", 200, o.Payload)
}
func (o *GetReloadOK) GetPayload() *models.Reload {
	return o.Payload
}

func (o *GetReloadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Reload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReloadNotFound creates a GetReloadNotFound with default headers values
func NewGetReloadNotFound() *GetReloadNotFound {
	return &GetReloadNotFound{}
}

/* GetReloadNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetReloadNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetReloadNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/reloads/{id}][%d] getReloadNotFound  %+v", 404, o.Payload)
}
func (o *GetReloadNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetReloadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetReloadDefault creates a GetReloadDefault with default headers values
func NewGetReloadDefault(code int) *GetReloadDefault {
	return &GetReloadDefault{
		_statusCode: code,
	}
}

/* GetReloadDefault describes a response with status code -1, with default header values.

General Error
*/
type GetReloadDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get reload default response
func (o *GetReloadDefault) Code() int {
	return o._statusCode
}

func (o *GetReloadDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/reloads/{id}][%d] getReload default  %+v", o._statusCode, o.Payload)
}
func (o *GetReloadDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetReloadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
