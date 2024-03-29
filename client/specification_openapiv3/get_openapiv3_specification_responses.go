// Code generated by go-swagger; DO NOT EDIT.

package specification_openapiv3

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v5/models"
)

// GetOpenapiv3SpecificationReader is a Reader for the GetOpenapiv3Specification structure.
type GetOpenapiv3SpecificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOpenapiv3SpecificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOpenapiv3SpecificationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetOpenapiv3SpecificationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOpenapiv3SpecificationOK creates a GetOpenapiv3SpecificationOK with default headers values
func NewGetOpenapiv3SpecificationOK() *GetOpenapiv3SpecificationOK {
	return &GetOpenapiv3SpecificationOK{}
}

/*
GetOpenapiv3SpecificationOK describes a response with status code 200, with default header values.

Success
*/
type GetOpenapiv3SpecificationOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get openapiv3 specification o k response has a 2xx status code
func (o *GetOpenapiv3SpecificationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get openapiv3 specification o k response has a 3xx status code
func (o *GetOpenapiv3SpecificationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get openapiv3 specification o k response has a 4xx status code
func (o *GetOpenapiv3SpecificationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get openapiv3 specification o k response has a 5xx status code
func (o *GetOpenapiv3SpecificationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get openapiv3 specification o k response a status code equal to that given
func (o *GetOpenapiv3SpecificationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get openapiv3 specification o k response
func (o *GetOpenapiv3SpecificationOK) Code() int {
	return 200
}

func (o *GetOpenapiv3SpecificationOK) Error() string {
	return fmt.Sprintf("[GET /specification_openapiv3][%d] getOpenapiv3SpecificationOK  %+v", 200, o.Payload)
}

func (o *GetOpenapiv3SpecificationOK) String() string {
	return fmt.Sprintf("[GET /specification_openapiv3][%d] getOpenapiv3SpecificationOK  %+v", 200, o.Payload)
}

func (o *GetOpenapiv3SpecificationOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOpenapiv3SpecificationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOpenapiv3SpecificationDefault creates a GetOpenapiv3SpecificationDefault with default headers values
func NewGetOpenapiv3SpecificationDefault(code int) *GetOpenapiv3SpecificationDefault {
	return &GetOpenapiv3SpecificationDefault{
		_statusCode: code,
	}
}

/*
GetOpenapiv3SpecificationDefault describes a response with status code -1, with default header values.

General Error
*/
type GetOpenapiv3SpecificationDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get openapiv3 specification default response has a 2xx status code
func (o *GetOpenapiv3SpecificationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get openapiv3 specification default response has a 3xx status code
func (o *GetOpenapiv3SpecificationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get openapiv3 specification default response has a 4xx status code
func (o *GetOpenapiv3SpecificationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get openapiv3 specification default response has a 5xx status code
func (o *GetOpenapiv3SpecificationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get openapiv3 specification default response a status code equal to that given
func (o *GetOpenapiv3SpecificationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get openapiv3 specification default response
func (o *GetOpenapiv3SpecificationDefault) Code() int {
	return o._statusCode
}

func (o *GetOpenapiv3SpecificationDefault) Error() string {
	return fmt.Sprintf("[GET /specification_openapiv3][%d] getOpenapiv3Specification default  %+v", o._statusCode, o.Payload)
}

func (o *GetOpenapiv3SpecificationDefault) String() string {
	return fmt.Sprintf("[GET /specification_openapiv3][%d] getOpenapiv3Specification default  %+v", o._statusCode, o.Payload)
}

func (o *GetOpenapiv3SpecificationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOpenapiv3SpecificationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
