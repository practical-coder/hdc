// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"client-native/models"
)

// GetAllStorageSSLCertificatesReader is a Reader for the GetAllStorageSSLCertificates structure.
type GetAllStorageSSLCertificatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllStorageSSLCertificatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllStorageSSLCertificatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAllStorageSSLCertificatesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAllStorageSSLCertificatesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAllStorageSSLCertificatesOK creates a GetAllStorageSSLCertificatesOK with default headers values
func NewGetAllStorageSSLCertificatesOK() *GetAllStorageSSLCertificatesOK {
	return &GetAllStorageSSLCertificatesOK{}
}

/* GetAllStorageSSLCertificatesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetAllStorageSSLCertificatesOK struct {
	Payload models.SslCertificates
}

func (o *GetAllStorageSSLCertificatesOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/storage/ssl_certificates][%d] getAllStorageSSLCertificatesOK  %+v", 200, o.Payload)
}
func (o *GetAllStorageSSLCertificatesOK) GetPayload() models.SslCertificates {
	return o.Payload
}

func (o *GetAllStorageSSLCertificatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllStorageSSLCertificatesNotFound creates a GetAllStorageSSLCertificatesNotFound with default headers values
func NewGetAllStorageSSLCertificatesNotFound() *GetAllStorageSSLCertificatesNotFound {
	return &GetAllStorageSSLCertificatesNotFound{}
}

/* GetAllStorageSSLCertificatesNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetAllStorageSSLCertificatesNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetAllStorageSSLCertificatesNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/storage/ssl_certificates][%d] getAllStorageSSLCertificatesNotFound  %+v", 404, o.Payload)
}
func (o *GetAllStorageSSLCertificatesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllStorageSSLCertificatesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAllStorageSSLCertificatesDefault creates a GetAllStorageSSLCertificatesDefault with default headers values
func NewGetAllStorageSSLCertificatesDefault(code int) *GetAllStorageSSLCertificatesDefault {
	return &GetAllStorageSSLCertificatesDefault{
		_statusCode: code,
	}
}

/* GetAllStorageSSLCertificatesDefault describes a response with status code -1, with default header values.

General Error
*/
type GetAllStorageSSLCertificatesDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get all storage s s l certificates default response
func (o *GetAllStorageSSLCertificatesDefault) Code() int {
	return o._statusCode
}

func (o *GetAllStorageSSLCertificatesDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/storage/ssl_certificates][%d] getAllStorageSSLCertificates default  %+v", o._statusCode, o.Payload)
}
func (o *GetAllStorageSSLCertificatesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllStorageSSLCertificatesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
