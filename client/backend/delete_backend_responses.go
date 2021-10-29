// Code generated by go-swagger; DO NOT EDIT.

package backend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/models"
)

// DeleteBackendReader is a Reader for the DeleteBackend structure.
type DeleteBackendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBackendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteBackendAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteBackendNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteBackendNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteBackendDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteBackendAccepted creates a DeleteBackendAccepted with default headers values
func NewDeleteBackendAccepted() *DeleteBackendAccepted {
	return &DeleteBackendAccepted{}
}

/* DeleteBackendAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeleteBackendAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteBackendAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/backends/{name}][%d] deleteBackendAccepted ", 202)
}

func (o *DeleteBackendAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeleteBackendNoContent creates a DeleteBackendNoContent with default headers values
func NewDeleteBackendNoContent() *DeleteBackendNoContent {
	return &DeleteBackendNoContent{}
}

/* DeleteBackendNoContent describes a response with status code 204, with default header values.

Backend deleted
*/
type DeleteBackendNoContent struct {
}

func (o *DeleteBackendNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/backends/{name}][%d] deleteBackendNoContent ", 204)
}

func (o *DeleteBackendNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteBackendNotFound creates a DeleteBackendNotFound with default headers values
func NewDeleteBackendNotFound() *DeleteBackendNotFound {
	return &DeleteBackendNotFound{}
}

/* DeleteBackendNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteBackendNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteBackendNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/backends/{name}][%d] deleteBackendNotFound  %+v", 404, o.Payload)
}
func (o *DeleteBackendNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteBackendNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteBackendDefault creates a DeleteBackendDefault with default headers values
func NewDeleteBackendDefault(code int) *DeleteBackendDefault {
	return &DeleteBackendDefault{
		_statusCode: code,
	}
}

/* DeleteBackendDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteBackendDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete backend default response
func (o *DeleteBackendDefault) Code() int {
	return o._statusCode
}

func (o *DeleteBackendDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/backends/{name}][%d] deleteBackend default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteBackendDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteBackendDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
