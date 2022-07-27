// Code generated by go-swagger; DO NOT EDIT.

package resolver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// DeleteResolverReader is a Reader for the DeleteResolver structure.
type DeleteResolverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteResolverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteResolverAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteResolverNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteResolverNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteResolverDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteResolverAccepted creates a DeleteResolverAccepted with default headers values
func NewDeleteResolverAccepted() *DeleteResolverAccepted {
	return &DeleteResolverAccepted{}
}

/* DeleteResolverAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeleteResolverAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteResolverAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/resolvers/{name}][%d] deleteResolverAccepted ", 202)
}

func (o *DeleteResolverAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeleteResolverNoContent creates a DeleteResolverNoContent with default headers values
func NewDeleteResolverNoContent() *DeleteResolverNoContent {
	return &DeleteResolverNoContent{}
}

/* DeleteResolverNoContent describes a response with status code 204, with default header values.

Resolver deleted
*/
type DeleteResolverNoContent struct {
}

func (o *DeleteResolverNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/resolvers/{name}][%d] deleteResolverNoContent ", 204)
}

func (o *DeleteResolverNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteResolverNotFound creates a DeleteResolverNotFound with default headers values
func NewDeleteResolverNotFound() *DeleteResolverNotFound {
	return &DeleteResolverNotFound{}
}

/* DeleteResolverNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteResolverNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteResolverNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/resolvers/{name}][%d] deleteResolverNotFound  %+v", 404, o.Payload)
}
func (o *DeleteResolverNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteResolverNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteResolverDefault creates a DeleteResolverDefault with default headers values
func NewDeleteResolverDefault(code int) *DeleteResolverDefault {
	return &DeleteResolverDefault{
		_statusCode: code,
	}
}

/* DeleteResolverDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteResolverDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete resolver default response
func (o *DeleteResolverDefault) Code() int {
	return o._statusCode
}

func (o *DeleteResolverDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/resolvers/{name}][%d] deleteResolver default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteResolverDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteResolverDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
