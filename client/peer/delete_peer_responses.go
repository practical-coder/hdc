// Code generated by go-swagger; DO NOT EDIT.

package peer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v3/models"
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

/*DeletePeerAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type DeletePeerAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string
}

func (o *DeletePeerAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/peer_section/{name}][%d] deletePeerAccepted ", 202)
}

func (o *DeletePeerAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	return nil
}

// NewDeletePeerNoContent creates a DeletePeerNoContent with default headers values
func NewDeletePeerNoContent() *DeletePeerNoContent {
	return &DeletePeerNoContent{}
}

/*DeletePeerNoContent handles this case with default header values.

Peer deleted
*/
type DeletePeerNoContent struct {
}

func (o *DeletePeerNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/peer_section/{name}][%d] deletePeerNoContent ", 204)
}

func (o *DeletePeerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePeerNotFound creates a DeletePeerNotFound with default headers values
func NewDeletePeerNotFound() *DeletePeerNotFound {
	return &DeletePeerNotFound{}
}

/*DeletePeerNotFound handles this case with default header values.

The specified resource was not found
*/
type DeletePeerNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeletePeerNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/peer_section/{name}][%d] deletePeerNotFound  %+v", 404, o.Payload)
}

func (o *DeletePeerNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeletePeerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

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

/*DeletePeerDefault handles this case with default header values.

General Error
*/
type DeletePeerDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete peer default response
func (o *DeletePeerDefault) Code() int {
	return o._statusCode
}

func (o *DeletePeerDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/peer_section/{name}][%d] deletePeer default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePeerDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeletePeerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
