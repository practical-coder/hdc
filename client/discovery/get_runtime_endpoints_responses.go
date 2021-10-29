// Code generated by go-swagger; DO NOT EDIT.

package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/cloud-native/models"
)

// GetRuntimeEndpointsReader is a Reader for the GetRuntimeEndpoints structure.
type GetRuntimeEndpointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRuntimeEndpointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRuntimeEndpointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetRuntimeEndpointsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRuntimeEndpointsOK creates a GetRuntimeEndpointsOK with default headers values
func NewGetRuntimeEndpointsOK() *GetRuntimeEndpointsOK {
	return &GetRuntimeEndpointsOK{}
}

/* GetRuntimeEndpointsOK describes a response with status code 200, with default header values.

Success
*/
type GetRuntimeEndpointsOK struct {
	Payload models.Endpoints
}

func (o *GetRuntimeEndpointsOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime][%d] getRuntimeEndpointsOK  %+v", 200, o.Payload)
}
func (o *GetRuntimeEndpointsOK) GetPayload() models.Endpoints {
	return o.Payload
}

func (o *GetRuntimeEndpointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRuntimeEndpointsDefault creates a GetRuntimeEndpointsDefault with default headers values
func NewGetRuntimeEndpointsDefault(code int) *GetRuntimeEndpointsDefault {
	return &GetRuntimeEndpointsDefault{
		_statusCode: code,
	}
}

/* GetRuntimeEndpointsDefault describes a response with status code -1, with default header values.

General Error
*/
type GetRuntimeEndpointsDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get runtime endpoints default response
func (o *GetRuntimeEndpointsDefault) Code() int {
	return o._statusCode
}

func (o *GetRuntimeEndpointsDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime][%d] getRuntimeEndpoints default  %+v", o._statusCode, o.Payload)
}
func (o *GetRuntimeEndpointsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRuntimeEndpointsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
