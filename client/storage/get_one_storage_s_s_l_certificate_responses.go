// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/models"
)

// GetOneStorageSSLCertificateReader is a Reader for the GetOneStorageSSLCertificate structure.
type GetOneStorageSSLCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOneStorageSSLCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOneStorageSSLCertificateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetOneStorageSSLCertificateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetOneStorageSSLCertificateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOneStorageSSLCertificateOK creates a GetOneStorageSSLCertificateOK with default headers values
func NewGetOneStorageSSLCertificateOK() *GetOneStorageSSLCertificateOK {
	return &GetOneStorageSSLCertificateOK{}
}

/* GetOneStorageSSLCertificateOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOneStorageSSLCertificateOK struct {
	Payload *models.SslCertificate
}

func (o *GetOneStorageSSLCertificateOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/storage/ssl_certificates/{name}][%d] getOneStorageSSLCertificateOK  %+v", 200, o.Payload)
}
func (o *GetOneStorageSSLCertificateOK) GetPayload() *models.SslCertificate {
	return o.Payload
}

func (o *GetOneStorageSSLCertificateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SslCertificate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOneStorageSSLCertificateNotFound creates a GetOneStorageSSLCertificateNotFound with default headers values
func NewGetOneStorageSSLCertificateNotFound() *GetOneStorageSSLCertificateNotFound {
	return &GetOneStorageSSLCertificateNotFound{}
}

/* GetOneStorageSSLCertificateNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetOneStorageSSLCertificateNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetOneStorageSSLCertificateNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/storage/ssl_certificates/{name}][%d] getOneStorageSSLCertificateNotFound  %+v", 404, o.Payload)
}
func (o *GetOneStorageSSLCertificateNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOneStorageSSLCertificateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetOneStorageSSLCertificateDefault creates a GetOneStorageSSLCertificateDefault with default headers values
func NewGetOneStorageSSLCertificateDefault(code int) *GetOneStorageSSLCertificateDefault {
	return &GetOneStorageSSLCertificateDefault{
		_statusCode: code,
	}
}

/* GetOneStorageSSLCertificateDefault describes a response with status code -1, with default header values.

General Error
*/
type GetOneStorageSSLCertificateDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get one storage s s l certificate default response
func (o *GetOneStorageSSLCertificateDefault) Code() int {
	return o._statusCode
}

func (o *GetOneStorageSSLCertificateDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/storage/ssl_certificates/{name}][%d] getOneStorageSSLCertificate default  %+v", o._statusCode, o.Payload)
}
func (o *GetOneStorageSSLCertificateDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOneStorageSSLCertificateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
