// Code generated by go-swagger; DO NOT EDIT.

package configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v2/models"
)

// PostHAProxyConfigurationReader is a Reader for the PostHAProxyConfiguration structure.
type PostHAProxyConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostHAProxyConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostHAProxyConfigurationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPostHAProxyConfigurationAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostHAProxyConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPostHAProxyConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostHAProxyConfigurationCreated creates a PostHAProxyConfigurationCreated with default headers values
func NewPostHAProxyConfigurationCreated() *PostHAProxyConfigurationCreated {
	return &PostHAProxyConfigurationCreated{}
}

/* PostHAProxyConfigurationCreated describes a response with status code 201, with default header values.

New HAProxy configuration pushed
*/
type PostHAProxyConfigurationCreated struct {
	Payload string
}

func (o *PostHAProxyConfigurationCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/raw][%d] postHAProxyConfigurationCreated  %+v", 201, o.Payload)
}
func (o *PostHAProxyConfigurationCreated) GetPayload() string {
	return o.Payload
}

func (o *PostHAProxyConfigurationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostHAProxyConfigurationAccepted creates a PostHAProxyConfigurationAccepted with default headers values
func NewPostHAProxyConfigurationAccepted() *PostHAProxyConfigurationAccepted {
	return &PostHAProxyConfigurationAccepted{}
}

/* PostHAProxyConfigurationAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type PostHAProxyConfigurationAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload string
}

func (o *PostHAProxyConfigurationAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/raw][%d] postHAProxyConfigurationAccepted  %+v", 202, o.Payload)
}
func (o *PostHAProxyConfigurationAccepted) GetPayload() string {
	return o.Payload
}

func (o *PostHAProxyConfigurationAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostHAProxyConfigurationBadRequest creates a PostHAProxyConfigurationBadRequest with default headers values
func NewPostHAProxyConfigurationBadRequest() *PostHAProxyConfigurationBadRequest {
	return &PostHAProxyConfigurationBadRequest{}
}

/* PostHAProxyConfigurationBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostHAProxyConfigurationBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *PostHAProxyConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/raw][%d] postHAProxyConfigurationBadRequest  %+v", 400, o.Payload)
}
func (o *PostHAProxyConfigurationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostHAProxyConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPostHAProxyConfigurationDefault creates a PostHAProxyConfigurationDefault with default headers values
func NewPostHAProxyConfigurationDefault(code int) *PostHAProxyConfigurationDefault {
	return &PostHAProxyConfigurationDefault{
		_statusCode: code,
	}
}

/* PostHAProxyConfigurationDefault describes a response with status code -1, with default header values.

General Error
*/
type PostHAProxyConfigurationDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the post h a proxy configuration default response
func (o *PostHAProxyConfigurationDefault) Code() int {
	return o._statusCode
}

func (o *PostHAProxyConfigurationDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/raw][%d] postHAProxyConfiguration default  %+v", o._statusCode, o.Payload)
}
func (o *PostHAProxyConfigurationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostHAProxyConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
