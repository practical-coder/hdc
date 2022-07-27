// Code generated by go-swagger; DO NOT EDIT.

package tcp_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// ReplaceTCPCheckReader is a Reader for the ReplaceTCPCheck structure.
type ReplaceTCPCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceTCPCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceTCPCheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceTCPCheckAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceTCPCheckBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceTCPCheckNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceTCPCheckDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceTCPCheckOK creates a ReplaceTCPCheckOK with default headers values
func NewReplaceTCPCheckOK() *ReplaceTCPCheckOK {
	return &ReplaceTCPCheckOK{}
}

/* ReplaceTCPCheckOK describes a response with status code 200, with default header values.

TCP check replaced
*/
type ReplaceTCPCheckOK struct {
	Payload *models.TCPCheck
}

func (o *ReplaceTCPCheckOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/tcp_checks/{index}][%d] replaceTcpCheckOK  %+v", 200, o.Payload)
}
func (o *ReplaceTCPCheckOK) GetPayload() *models.TCPCheck {
	return o.Payload
}

func (o *ReplaceTCPCheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TCPCheck)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceTCPCheckAccepted creates a ReplaceTCPCheckAccepted with default headers values
func NewReplaceTCPCheckAccepted() *ReplaceTCPCheckAccepted {
	return &ReplaceTCPCheckAccepted{}
}

/* ReplaceTCPCheckAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type ReplaceTCPCheckAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.TCPCheck
}

func (o *ReplaceTCPCheckAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/tcp_checks/{index}][%d] replaceTcpCheckAccepted  %+v", 202, o.Payload)
}
func (o *ReplaceTCPCheckAccepted) GetPayload() *models.TCPCheck {
	return o.Payload
}

func (o *ReplaceTCPCheckAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.TCPCheck)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceTCPCheckBadRequest creates a ReplaceTCPCheckBadRequest with default headers values
func NewReplaceTCPCheckBadRequest() *ReplaceTCPCheckBadRequest {
	return &ReplaceTCPCheckBadRequest{}
}

/* ReplaceTCPCheckBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceTCPCheckBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceTCPCheckBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/tcp_checks/{index}][%d] replaceTcpCheckBadRequest  %+v", 400, o.Payload)
}
func (o *ReplaceTCPCheckBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceTCPCheckBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceTCPCheckNotFound creates a ReplaceTCPCheckNotFound with default headers values
func NewReplaceTCPCheckNotFound() *ReplaceTCPCheckNotFound {
	return &ReplaceTCPCheckNotFound{}
}

/* ReplaceTCPCheckNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type ReplaceTCPCheckNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceTCPCheckNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/tcp_checks/{index}][%d] replaceTcpCheckNotFound  %+v", 404, o.Payload)
}
func (o *ReplaceTCPCheckNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceTCPCheckNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceTCPCheckDefault creates a ReplaceTCPCheckDefault with default headers values
func NewReplaceTCPCheckDefault(code int) *ReplaceTCPCheckDefault {
	return &ReplaceTCPCheckDefault{
		_statusCode: code,
	}
}

/* ReplaceTCPCheckDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceTCPCheckDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the replace TCP check default response
func (o *ReplaceTCPCheckDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceTCPCheckDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/tcp_checks/{index}][%d] replaceTCPCheck default  %+v", o._statusCode, o.Payload)
}
func (o *ReplaceTCPCheckDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceTCPCheckDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
