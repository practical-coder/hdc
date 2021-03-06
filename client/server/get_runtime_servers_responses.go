// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// GetRuntimeServersReader is a Reader for the GetRuntimeServers structure.
type GetRuntimeServersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRuntimeServersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRuntimeServersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetRuntimeServersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRuntimeServersOK creates a GetRuntimeServersOK with default headers values
func NewGetRuntimeServersOK() *GetRuntimeServersOK {
	return &GetRuntimeServersOK{}
}

/* GetRuntimeServersOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetRuntimeServersOK struct {
	Payload models.RuntimeServers
}

func (o *GetRuntimeServersOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/servers][%d] getRuntimeServersOK  %+v", 200, o.Payload)
}
func (o *GetRuntimeServersOK) GetPayload() models.RuntimeServers {
	return o.Payload
}

func (o *GetRuntimeServersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRuntimeServersDefault creates a GetRuntimeServersDefault with default headers values
func NewGetRuntimeServersDefault(code int) *GetRuntimeServersDefault {
	return &GetRuntimeServersDefault{
		_statusCode: code,
	}
}

/* GetRuntimeServersDefault describes a response with status code -1, with default header values.

General Error
*/
type GetRuntimeServersDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get runtime servers default response
func (o *GetRuntimeServersDefault) Code() int {
	return o._statusCode
}

func (o *GetRuntimeServersDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/servers][%d] getRuntimeServers default  %+v", o._statusCode, o.Payload)
}
func (o *GetRuntimeServersDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRuntimeServersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
