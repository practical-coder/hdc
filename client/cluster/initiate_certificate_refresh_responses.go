// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/cloud-native/models"
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

/* InitiateCertificateRefreshOK describes a response with status code 200, with default header values.

refresh activated
*/
type InitiateCertificateRefreshOK struct {
}

func (o *InitiateCertificateRefreshOK) Error() string {
	return fmt.Sprintf("[POST /cluster/certificate][%d] initiateCertificateRefreshOK ", 200)
}

func (o *InitiateCertificateRefreshOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInitiateCertificateRefreshForbidden creates a InitiateCertificateRefreshForbidden with default headers values
func NewInitiateCertificateRefreshForbidden() *InitiateCertificateRefreshForbidden {
	return &InitiateCertificateRefreshForbidden{}
}

/* InitiateCertificateRefreshForbidden describes a response with status code 403, with default header values.

refresh not possible
*/
type InitiateCertificateRefreshForbidden struct {
}

func (o *InitiateCertificateRefreshForbidden) Error() string {
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

/* InitiateCertificateRefreshDefault describes a response with status code -1, with default header values.

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

func (o *InitiateCertificateRefreshDefault) Error() string {
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
