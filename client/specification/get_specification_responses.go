// Code generated by go-swagger; DO NOT EDIT.

package specification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v2/models"
)

// GetSpecificationReader is a Reader for the GetSpecification structure.
type GetSpecificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpecificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSpecificationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSpecificationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSpecificationOK creates a GetSpecificationOK with default headers values
func NewGetSpecificationOK() *GetSpecificationOK {
	return &GetSpecificationOK{}
}

/* GetSpecificationOK describes a response with status code 200, with default header values.

Success
*/
type GetSpecificationOK struct {
	Payload interface{}
}

func (o *GetSpecificationOK) Error() string {
	return fmt.Sprintf("[GET /specification][%d] getSpecificationOK  %+v", 200, o.Payload)
}
func (o *GetSpecificationOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetSpecificationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpecificationDefault creates a GetSpecificationDefault with default headers values
func NewGetSpecificationDefault(code int) *GetSpecificationDefault {
	return &GetSpecificationDefault{
		_statusCode: code,
	}
}

/* GetSpecificationDefault describes a response with status code -1, with default header values.

General Error
*/
type GetSpecificationDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get specification default response
func (o *GetSpecificationDefault) Code() int {
	return o._statusCode
}

func (o *GetSpecificationDefault) Error() string {
	return fmt.Sprintf("[GET /specification][%d] getSpecification default  %+v", o._statusCode, o.Payload)
}
func (o *GetSpecificationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSpecificationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
