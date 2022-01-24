// Code generated by go-swagger; DO NOT EDIT.

package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v2/models"
)

// GetConfigurationEndpointsReader is a Reader for the GetConfigurationEndpoints structure.
type GetConfigurationEndpointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfigurationEndpointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConfigurationEndpointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetConfigurationEndpointsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetConfigurationEndpointsOK creates a GetConfigurationEndpointsOK with default headers values
func NewGetConfigurationEndpointsOK() *GetConfigurationEndpointsOK {
	return &GetConfigurationEndpointsOK{}
}

/* GetConfigurationEndpointsOK describes a response with status code 200, with default header values.

Success
*/
type GetConfigurationEndpointsOK struct {
	Payload models.Endpoints
}

func (o *GetConfigurationEndpointsOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration][%d] getConfigurationEndpointsOK  %+v", 200, o.Payload)
}
func (o *GetConfigurationEndpointsOK) GetPayload() models.Endpoints {
	return o.Payload
}

func (o *GetConfigurationEndpointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationEndpointsDefault creates a GetConfigurationEndpointsDefault with default headers values
func NewGetConfigurationEndpointsDefault(code int) *GetConfigurationEndpointsDefault {
	return &GetConfigurationEndpointsDefault{
		_statusCode: code,
	}
}

/* GetConfigurationEndpointsDefault describes a response with status code -1, with default header values.

General Error
*/
type GetConfigurationEndpointsDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get configuration endpoints default response
func (o *GetConfigurationEndpointsDefault) Code() int {
	return o._statusCode
}

func (o *GetConfigurationEndpointsDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration][%d] getConfigurationEndpoints default  %+v", o._statusCode, o.Payload)
}
func (o *GetConfigurationEndpointsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetConfigurationEndpointsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
