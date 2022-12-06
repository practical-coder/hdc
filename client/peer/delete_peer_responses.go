// Code generated by go-swagger; DO NOT EDIT.

package peer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// DeletePeerReader is a Reader for the DeletePeer structure.
type DeletePeerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePeerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeletePeerAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeletePeerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeletePeerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeletePeerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePeerAccepted creates a DeletePeerAccepted with default headers values
func NewDeletePeerAccepted() *DeletePeerAccepted {
	return &DeletePeerAccepted{}
}

/*
DeletePeerAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeletePeerAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

// IsSuccess returns true when this delete peer accepted response has a 2xx status code
func (o *DeletePeerAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete peer accepted response has a 3xx status code
func (o *DeletePeerAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete peer accepted response has a 4xx status code
func (o *DeletePeerAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete peer accepted response has a 5xx status code
func (o *DeletePeerAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this delete peer accepted response a status code equal to that given
func (o *DeletePeerAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *DeletePeerAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/peer_section/{name}][%d] deletePeerAccepted ", 202)
}

func (o *DeletePeerAccepted) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/peer_section/{name}][%d] deletePeerAccepted ", 202)
}

func (o *DeletePeerAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeletePeerNoContent creates a DeletePeerNoContent with default headers values
func NewDeletePeerNoContent() *DeletePeerNoContent {
	return &DeletePeerNoContent{}
}

/*
DeletePeerNoContent describes a response with status code 204, with default header values.

Peer deleted
*/
type DeletePeerNoContent struct {
}

// IsSuccess returns true when this delete peer no content response has a 2xx status code
func (o *DeletePeerNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete peer no content response has a 3xx status code
func (o *DeletePeerNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete peer no content response has a 4xx status code
func (o *DeletePeerNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete peer no content response has a 5xx status code
func (o *DeletePeerNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete peer no content response a status code equal to that given
func (o *DeletePeerNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeletePeerNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/peer_section/{name}][%d] deletePeerNoContent ", 204)
}

func (o *DeletePeerNoContent) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/peer_section/{name}][%d] deletePeerNoContent ", 204)
}

func (o *DeletePeerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePeerNotFound creates a DeletePeerNotFound with default headers values
func NewDeletePeerNotFound() *DeletePeerNotFound {
	return &DeletePeerNotFound{}
}

/*
DeletePeerNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeletePeerNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete peer not found response has a 2xx status code
func (o *DeletePeerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete peer not found response has a 3xx status code
func (o *DeletePeerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete peer not found response has a 4xx status code
func (o *DeletePeerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete peer not found response has a 5xx status code
func (o *DeletePeerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete peer not found response a status code equal to that given
func (o *DeletePeerNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeletePeerNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/peer_section/{name}][%d] deletePeerNotFound  %+v", 404, o.Payload)
}

func (o *DeletePeerNotFound) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/peer_section/{name}][%d] deletePeerNotFound  %+v", 404, o.Payload)
}

func (o *DeletePeerNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeletePeerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeletePeerDefault creates a DeletePeerDefault with default headers values
func NewDeletePeerDefault(code int) *DeletePeerDefault {
	return &DeletePeerDefault{
		_statusCode: code,
	}
}

/*
DeletePeerDefault describes a response with status code -1, with default header values.

General Error
*/
type DeletePeerDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete peer default response
func (o *DeletePeerDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete peer default response has a 2xx status code
func (o *DeletePeerDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete peer default response has a 3xx status code
func (o *DeletePeerDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete peer default response has a 4xx status code
func (o *DeletePeerDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete peer default response has a 5xx status code
func (o *DeletePeerDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete peer default response a status code equal to that given
func (o *DeletePeerDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeletePeerDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/peer_section/{name}][%d] deletePeer default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePeerDefault) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/peer_section/{name}][%d] deletePeer default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePeerDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeletePeerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
