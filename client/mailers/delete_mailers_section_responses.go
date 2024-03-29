// Code generated by go-swagger; DO NOT EDIT.

package mailers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v5/models"
)

// DeleteMailersSectionReader is a Reader for the DeleteMailersSection structure.
type DeleteMailersSectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMailersSectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteMailersSectionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteMailersSectionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteMailersSectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteMailersSectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMailersSectionAccepted creates a DeleteMailersSectionAccepted with default headers values
func NewDeleteMailersSectionAccepted() *DeleteMailersSectionAccepted {
	return &DeleteMailersSectionAccepted{}
}

/*
DeleteMailersSectionAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeleteMailersSectionAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

// IsSuccess returns true when this delete mailers section accepted response has a 2xx status code
func (o *DeleteMailersSectionAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete mailers section accepted response has a 3xx status code
func (o *DeleteMailersSectionAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mailers section accepted response has a 4xx status code
func (o *DeleteMailersSectionAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete mailers section accepted response has a 5xx status code
func (o *DeleteMailersSectionAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this delete mailers section accepted response a status code equal to that given
func (o *DeleteMailersSectionAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the delete mailers section accepted response
func (o *DeleteMailersSectionAccepted) Code() int {
	return 202
}

func (o *DeleteMailersSectionAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/mailers_section/{name}][%d] deleteMailersSectionAccepted ", 202)
}

func (o *DeleteMailersSectionAccepted) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/mailers_section/{name}][%d] deleteMailersSectionAccepted ", 202)
}

func (o *DeleteMailersSectionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeleteMailersSectionNoContent creates a DeleteMailersSectionNoContent with default headers values
func NewDeleteMailersSectionNoContent() *DeleteMailersSectionNoContent {
	return &DeleteMailersSectionNoContent{}
}

/*
DeleteMailersSectionNoContent describes a response with status code 204, with default header values.

Mailers deleted
*/
type DeleteMailersSectionNoContent struct {
}

// IsSuccess returns true when this delete mailers section no content response has a 2xx status code
func (o *DeleteMailersSectionNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete mailers section no content response has a 3xx status code
func (o *DeleteMailersSectionNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mailers section no content response has a 4xx status code
func (o *DeleteMailersSectionNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete mailers section no content response has a 5xx status code
func (o *DeleteMailersSectionNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete mailers section no content response a status code equal to that given
func (o *DeleteMailersSectionNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete mailers section no content response
func (o *DeleteMailersSectionNoContent) Code() int {
	return 204
}

func (o *DeleteMailersSectionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/mailers_section/{name}][%d] deleteMailersSectionNoContent ", 204)
}

func (o *DeleteMailersSectionNoContent) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/mailers_section/{name}][%d] deleteMailersSectionNoContent ", 204)
}

func (o *DeleteMailersSectionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMailersSectionNotFound creates a DeleteMailersSectionNotFound with default headers values
func NewDeleteMailersSectionNotFound() *DeleteMailersSectionNotFound {
	return &DeleteMailersSectionNotFound{}
}

/*
DeleteMailersSectionNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteMailersSectionNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete mailers section not found response has a 2xx status code
func (o *DeleteMailersSectionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete mailers section not found response has a 3xx status code
func (o *DeleteMailersSectionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mailers section not found response has a 4xx status code
func (o *DeleteMailersSectionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete mailers section not found response has a 5xx status code
func (o *DeleteMailersSectionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete mailers section not found response a status code equal to that given
func (o *DeleteMailersSectionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete mailers section not found response
func (o *DeleteMailersSectionNotFound) Code() int {
	return 404
}

func (o *DeleteMailersSectionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/mailers_section/{name}][%d] deleteMailersSectionNotFound  %+v", 404, o.Payload)
}

func (o *DeleteMailersSectionNotFound) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/mailers_section/{name}][%d] deleteMailersSectionNotFound  %+v", 404, o.Payload)
}

func (o *DeleteMailersSectionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteMailersSectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteMailersSectionDefault creates a DeleteMailersSectionDefault with default headers values
func NewDeleteMailersSectionDefault(code int) *DeleteMailersSectionDefault {
	return &DeleteMailersSectionDefault{
		_statusCode: code,
	}
}

/*
DeleteMailersSectionDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteMailersSectionDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete mailers section default response has a 2xx status code
func (o *DeleteMailersSectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete mailers section default response has a 3xx status code
func (o *DeleteMailersSectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete mailers section default response has a 4xx status code
func (o *DeleteMailersSectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete mailers section default response has a 5xx status code
func (o *DeleteMailersSectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete mailers section default response a status code equal to that given
func (o *DeleteMailersSectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete mailers section default response
func (o *DeleteMailersSectionDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMailersSectionDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/mailers_section/{name}][%d] deleteMailersSection default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMailersSectionDefault) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/mailers_section/{name}][%d] deleteMailersSection default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMailersSectionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteMailersSectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
