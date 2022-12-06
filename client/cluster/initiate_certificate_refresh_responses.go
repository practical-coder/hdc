// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// InitiateCertificateRefreshReader is a Reader for the InitiateCertificateRefresh structure.
type InitiateCertificateRefreshReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InitiateCertificateRefreshReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInitiateCertificateRefreshOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewInitiateCertificateRefreshForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewInitiateCertificateRefreshDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewInitiateCertificateRefreshOK creates a InitiateCertificateRefreshOK with default headers values
func NewInitiateCertificateRefreshOK() *InitiateCertificateRefreshOK {
	return &InitiateCertificateRefreshOK{}
}

/*
InitiateCertificateRefreshOK describes a response with status code 200, with default header values.

refresh activated
*/
type InitiateCertificateRefreshOK struct {
}

// IsSuccess returns true when this initiate certificate refresh o k response has a 2xx status code
func (o *InitiateCertificateRefreshOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this initiate certificate refresh o k response has a 3xx status code
func (o *InitiateCertificateRefreshOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this initiate certificate refresh o k response has a 4xx status code
func (o *InitiateCertificateRefreshOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this initiate certificate refresh o k response has a 5xx status code
func (o *InitiateCertificateRefreshOK) IsServerError() bool {
	return false
}

// IsCode returns true when this initiate certificate refresh o k response a status code equal to that given
func (o *InitiateCertificateRefreshOK) IsCode(code int) bool {
	return code == 200
}

func (o *InitiateCertificateRefreshOK) Error() string {
	return fmt.Sprintf("[POST /cluster/certificate][%d] initiateCertificateRefreshOK ", 200)
}

func (o *InitiateCertificateRefreshOK) String() string {
	return fmt.Sprintf("[POST /cluster/certificate][%d] initiateCertificateRefreshOK ", 200)
}

func (o *InitiateCertificateRefreshOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInitiateCertificateRefreshForbidden creates a InitiateCertificateRefreshForbidden with default headers values
func NewInitiateCertificateRefreshForbidden() *InitiateCertificateRefreshForbidden {
	return &InitiateCertificateRefreshForbidden{}
}

/*
InitiateCertificateRefreshForbidden describes a response with status code 403, with default header values.

refresh not possible
*/
type InitiateCertificateRefreshForbidden struct {
}

// IsSuccess returns true when this initiate certificate refresh forbidden response has a 2xx status code
func (o *InitiateCertificateRefreshForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this initiate certificate refresh forbidden response has a 3xx status code
func (o *InitiateCertificateRefreshForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this initiate certificate refresh forbidden response has a 4xx status code
func (o *InitiateCertificateRefreshForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this initiate certificate refresh forbidden response has a 5xx status code
func (o *InitiateCertificateRefreshForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this initiate certificate refresh forbidden response a status code equal to that given
func (o *InitiateCertificateRefreshForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *InitiateCertificateRefreshForbidden) Error() string {
	return fmt.Sprintf("[POST /cluster/certificate][%d] initiateCertificateRefreshForbidden ", 403)
}

func (o *InitiateCertificateRefreshForbidden) String() string {
	return fmt.Sprintf("[POST /cluster/certificate][%d] initiateCertificateRefreshForbidden ", 403)
}

func (o *InitiateCertificateRefreshForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInitiateCertificateRefreshDefault creates a InitiateCertificateRefreshDefault with default headers values
func NewInitiateCertificateRefreshDefault(code int) *InitiateCertificateRefreshDefault {
	return &InitiateCertificateRefreshDefault{
		_statusCode: code,
	}
}

/*
InitiateCertificateRefreshDefault describes a response with status code -1, with default header values.

General Error
*/
type InitiateCertificateRefreshDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the initiate certificate refresh default response
func (o *InitiateCertificateRefreshDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this initiate certificate refresh default response has a 2xx status code
func (o *InitiateCertificateRefreshDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this initiate certificate refresh default response has a 3xx status code
func (o *InitiateCertificateRefreshDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this initiate certificate refresh default response has a 4xx status code
func (o *InitiateCertificateRefreshDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this initiate certificate refresh default response has a 5xx status code
func (o *InitiateCertificateRefreshDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this initiate certificate refresh default response a status code equal to that given
func (o *InitiateCertificateRefreshDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *InitiateCertificateRefreshDefault) Error() string {
	return fmt.Sprintf("[POST /cluster/certificate][%d] initiateCertificateRefresh default  %+v", o._statusCode, o.Payload)
}

func (o *InitiateCertificateRefreshDefault) String() string {
	return fmt.Sprintf("[POST /cluster/certificate][%d] initiateCertificateRefresh default  %+v", o._statusCode, o.Payload)
}

func (o *InitiateCertificateRefreshDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *InitiateCertificateRefreshDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
