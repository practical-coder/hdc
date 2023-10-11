// Code generated by go-swagger; DO NOT EDIT.

package information

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v5/models"
)

// GetHaproxyProcessInfoReader is a Reader for the GetHaproxyProcessInfo structure.
type GetHaproxyProcessInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHaproxyProcessInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHaproxyProcessInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetHaproxyProcessInfoDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHaproxyProcessInfoOK creates a GetHaproxyProcessInfoOK with default headers values
func NewGetHaproxyProcessInfoOK() *GetHaproxyProcessInfoOK {
	return &GetHaproxyProcessInfoOK{}
}

/*
GetHaproxyProcessInfoOK describes a response with status code 200, with default header values.

Success
*/
type GetHaproxyProcessInfoOK struct {
	Payload models.ProcessInfos
}

// IsSuccess returns true when this get haproxy process info o k response has a 2xx status code
func (o *GetHaproxyProcessInfoOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get haproxy process info o k response has a 3xx status code
func (o *GetHaproxyProcessInfoOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get haproxy process info o k response has a 4xx status code
func (o *GetHaproxyProcessInfoOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get haproxy process info o k response has a 5xx status code
func (o *GetHaproxyProcessInfoOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get haproxy process info o k response a status code equal to that given
func (o *GetHaproxyProcessInfoOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get haproxy process info o k response
func (o *GetHaproxyProcessInfoOK) Code() int {
	return 200
}

func (o *GetHaproxyProcessInfoOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/info][%d] getHaproxyProcessInfoOK  %+v", 200, o.Payload)
}

func (o *GetHaproxyProcessInfoOK) String() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/info][%d] getHaproxyProcessInfoOK  %+v", 200, o.Payload)
}

func (o *GetHaproxyProcessInfoOK) GetPayload() models.ProcessInfos {
	return o.Payload
}

func (o *GetHaproxyProcessInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHaproxyProcessInfoDefault creates a GetHaproxyProcessInfoDefault with default headers values
func NewGetHaproxyProcessInfoDefault(code int) *GetHaproxyProcessInfoDefault {
	return &GetHaproxyProcessInfoDefault{
		_statusCode: code,
	}
}

/*
GetHaproxyProcessInfoDefault describes a response with status code -1, with default header values.

General Error
*/
type GetHaproxyProcessInfoDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get haproxy process info default response has a 2xx status code
func (o *GetHaproxyProcessInfoDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get haproxy process info default response has a 3xx status code
func (o *GetHaproxyProcessInfoDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get haproxy process info default response has a 4xx status code
func (o *GetHaproxyProcessInfoDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get haproxy process info default response has a 5xx status code
func (o *GetHaproxyProcessInfoDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get haproxy process info default response a status code equal to that given
func (o *GetHaproxyProcessInfoDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get haproxy process info default response
func (o *GetHaproxyProcessInfoDefault) Code() int {
	return o._statusCode
}

func (o *GetHaproxyProcessInfoDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/info][%d] getHaproxyProcessInfo default  %+v", o._statusCode, o.Payload)
}

func (o *GetHaproxyProcessInfoDefault) String() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/info][%d] getHaproxyProcessInfo default  %+v", o._statusCode, o.Payload)
}

func (o *GetHaproxyProcessInfoDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHaproxyProcessInfoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
