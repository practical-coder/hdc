// Code generated by go-swagger; DO NOT EDIT.

package server_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/models"
)

// DeleteServerTemplateReader is a Reader for the DeleteServerTemplate structure.
type DeleteServerTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteServerTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteServerTemplateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteServerTemplateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteServerTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteServerTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteServerTemplateAccepted creates a DeleteServerTemplateAccepted with default headers values
func NewDeleteServerTemplateAccepted() *DeleteServerTemplateAccepted {
	return &DeleteServerTemplateAccepted{}
}

/* DeleteServerTemplateAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeleteServerTemplateAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteServerTemplateAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/server_templates/{prefix}][%d] deleteServerTemplateAccepted ", 202)
}

func (o *DeleteServerTemplateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeleteServerTemplateNoContent creates a DeleteServerTemplateNoContent with default headers values
func NewDeleteServerTemplateNoContent() *DeleteServerTemplateNoContent {
	return &DeleteServerTemplateNoContent{}
}

/* DeleteServerTemplateNoContent describes a response with status code 204, with default header values.

Server template deleted
*/
type DeleteServerTemplateNoContent struct {
}

func (o *DeleteServerTemplateNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/server_templates/{prefix}][%d] deleteServerTemplateNoContent ", 204)
}

func (o *DeleteServerTemplateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServerTemplateNotFound creates a DeleteServerTemplateNotFound with default headers values
func NewDeleteServerTemplateNotFound() *DeleteServerTemplateNotFound {
	return &DeleteServerTemplateNotFound{}
}

/* DeleteServerTemplateNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteServerTemplateNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteServerTemplateNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/server_templates/{prefix}][%d] deleteServerTemplateNotFound  %+v", 404, o.Payload)
}
func (o *DeleteServerTemplateNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteServerTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteServerTemplateDefault creates a DeleteServerTemplateDefault with default headers values
func NewDeleteServerTemplateDefault(code int) *DeleteServerTemplateDefault {
	return &DeleteServerTemplateDefault{
		_statusCode: code,
	}
}

/* DeleteServerTemplateDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteServerTemplateDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete server template default response
func (o *DeleteServerTemplateDefault) Code() int {
	return o._statusCode
}

func (o *DeleteServerTemplateDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/server_templates/{prefix}][%d] deleteServerTemplate default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteServerTemplateDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteServerTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
