// Code generated by go-swagger; DO NOT EDIT.

package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/models"
)

// GetSpoeEndpointsReader is a Reader for the GetSpoeEndpoints structure.
type GetSpoeEndpointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpoeEndpointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSpoeEndpointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSpoeEndpointsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSpoeEndpointsOK creates a GetSpoeEndpointsOK with default headers values
func NewGetSpoeEndpointsOK() *GetSpoeEndpointsOK {
	return &GetSpoeEndpointsOK{}
}

/* GetSpoeEndpointsOK describes a response with status code 200, with default header values.

Success
*/
type GetSpoeEndpointsOK struct {
	Payload models.Endpoints
}

func (o *GetSpoeEndpointsOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe][%d] getSpoeEndpointsOK  %+v", 200, o.Payload)
}
func (o *GetSpoeEndpointsOK) GetPayload() models.Endpoints {
	return o.Payload
}

func (o *GetSpoeEndpointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpoeEndpointsDefault creates a GetSpoeEndpointsDefault with default headers values
func NewGetSpoeEndpointsDefault(code int) *GetSpoeEndpointsDefault {
	return &GetSpoeEndpointsDefault{
		_statusCode: code,
	}
}

/* GetSpoeEndpointsDefault describes a response with status code -1, with default header values.

General Error
*/
type GetSpoeEndpointsDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get spoe endpoints default response
func (o *GetSpoeEndpointsDefault) Code() int {
	return o._statusCode
}

func (o *GetSpoeEndpointsDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe][%d] getSpoeEndpoints default  %+v", o._statusCode, o.Payload)
}
func (o *GetSpoeEndpointsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSpoeEndpointsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
