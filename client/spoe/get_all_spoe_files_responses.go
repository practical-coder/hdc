// Code generated by go-swagger; DO NOT EDIT.

package spoe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/practical-coder/hdc/models"
)

// GetAllSpoeFilesReader is a Reader for the GetAllSpoeFiles structure.
type GetAllSpoeFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllSpoeFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllSpoeFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAllSpoeFilesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAllSpoeFilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAllSpoeFilesOK creates a GetAllSpoeFilesOK with default headers values
func NewGetAllSpoeFilesOK() *GetAllSpoeFilesOK {
	return &GetAllSpoeFilesOK{}
}

/* GetAllSpoeFilesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetAllSpoeFilesOK struct {
	Payload models.SpoeFiles
}

func (o *GetAllSpoeFilesOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_files][%d] getAllSpoeFilesOK  %+v", 200, o.Payload)
}
func (o *GetAllSpoeFilesOK) GetPayload() models.SpoeFiles {
	return o.Payload
}

func (o *GetAllSpoeFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllSpoeFilesNotFound creates a GetAllSpoeFilesNotFound with default headers values
func NewGetAllSpoeFilesNotFound() *GetAllSpoeFilesNotFound {
	return &GetAllSpoeFilesNotFound{}
}

/* GetAllSpoeFilesNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetAllSpoeFilesNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetAllSpoeFilesNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_files][%d] getAllSpoeFilesNotFound  %+v", 404, o.Payload)
}
func (o *GetAllSpoeFilesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllSpoeFilesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAllSpoeFilesDefault creates a GetAllSpoeFilesDefault with default headers values
func NewGetAllSpoeFilesDefault(code int) *GetAllSpoeFilesDefault {
	return &GetAllSpoeFilesDefault{
		_statusCode: code,
	}
}

/* GetAllSpoeFilesDefault describes a response with status code -1, with default header values.

General Error
*/
type GetAllSpoeFilesDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get all spoe files default response
func (o *GetAllSpoeFilesDefault) Code() int {
	return o._statusCode
}

func (o *GetAllSpoeFilesDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_files][%d] getAllSpoeFiles default  %+v", o._statusCode, o.Payload)
}
func (o *GetAllSpoeFilesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllSpoeFilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
