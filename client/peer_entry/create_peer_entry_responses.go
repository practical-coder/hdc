// Code generated by go-swagger; DO NOT EDIT.

package peer_entry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v2/models"
)

// CreatePeerEntryReader is a Reader for the CreatePeerEntry structure.
type CreatePeerEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePeerEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreatePeerEntryCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreatePeerEntryAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreatePeerEntryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreatePeerEntryConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreatePeerEntryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreatePeerEntryCreated creates a CreatePeerEntryCreated with default headers values
func NewCreatePeerEntryCreated() *CreatePeerEntryCreated {
	return &CreatePeerEntryCreated{}
}

/*CreatePeerEntryCreated handles this case with default header values.

PeerEntry created
*/
type CreatePeerEntryCreated struct {
	Payload *models.PeerEntry
}

func (o *CreatePeerEntryCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/peer_entries][%d] createPeerEntryCreated  %+v", 201, o.Payload)
}

func (o *CreatePeerEntryCreated) GetPayload() *models.PeerEntry {
	return o.Payload
}

func (o *CreatePeerEntryCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PeerEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePeerEntryAccepted creates a CreatePeerEntryAccepted with default headers values
func NewCreatePeerEntryAccepted() *CreatePeerEntryAccepted {
	return &CreatePeerEntryAccepted{}
}

/*CreatePeerEntryAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type CreatePeerEntryAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string

	Payload *models.PeerEntry
}

func (o *CreatePeerEntryAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/peer_entries][%d] createPeerEntryAccepted  %+v", 202, o.Payload)
}

func (o *CreatePeerEntryAccepted) GetPayload() *models.PeerEntry {
	return o.Payload
}

func (o *CreatePeerEntryAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	o.Payload = new(models.PeerEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePeerEntryBadRequest creates a CreatePeerEntryBadRequest with default headers values
func NewCreatePeerEntryBadRequest() *CreatePeerEntryBadRequest {
	return &CreatePeerEntryBadRequest{}
}

/*CreatePeerEntryBadRequest handles this case with default header values.

Bad request
*/
type CreatePeerEntryBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreatePeerEntryBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/peer_entries][%d] createPeerEntryBadRequest  %+v", 400, o.Payload)
}

func (o *CreatePeerEntryBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreatePeerEntryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePeerEntryConflict creates a CreatePeerEntryConflict with default headers values
func NewCreatePeerEntryConflict() *CreatePeerEntryConflict {
	return &CreatePeerEntryConflict{}
}

/*CreatePeerEntryConflict handles this case with default header values.

The specified resource already exists
*/
type CreatePeerEntryConflict struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreatePeerEntryConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/peer_entries][%d] createPeerEntryConflict  %+v", 409, o.Payload)
}

func (o *CreatePeerEntryConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreatePeerEntryConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePeerEntryDefault creates a CreatePeerEntryDefault with default headers values
func NewCreatePeerEntryDefault(code int) *CreatePeerEntryDefault {
	return &CreatePeerEntryDefault{
		_statusCode: code,
	}
}

/*CreatePeerEntryDefault handles this case with default header values.

General Error
*/
type CreatePeerEntryDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create peer entry default response
func (o *CreatePeerEntryDefault) Code() int {
	return o._statusCode
}

func (o *CreatePeerEntryDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/peer_entries][%d] createPeerEntry default  %+v", o._statusCode, o.Payload)
}

func (o *CreatePeerEntryDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreatePeerEntryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
