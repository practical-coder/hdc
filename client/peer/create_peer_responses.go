// Code generated by go-swagger; DO NOT EDIT.

package peer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/cloud-native/models"
)

// CreatePeerReader is a Reader for the CreatePeer structure.
type CreatePeerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePeerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreatePeerCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreatePeerAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreatePeerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreatePeerConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreatePeerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreatePeerCreated creates a CreatePeerCreated with default headers values
func NewCreatePeerCreated() *CreatePeerCreated {
	return &CreatePeerCreated{}
}

/* CreatePeerCreated describes a response with status code 201, with default header values.

Peer created
*/
type CreatePeerCreated struct {
	Payload *models.PeerSection
}

func (o *CreatePeerCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/peer_section][%d] createPeerCreated  %+v", 201, o.Payload)
}
func (o *CreatePeerCreated) GetPayload() *models.PeerSection {
	return o.Payload
}

func (o *CreatePeerCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PeerSection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePeerAccepted creates a CreatePeerAccepted with default headers values
func NewCreatePeerAccepted() *CreatePeerAccepted {
	return &CreatePeerAccepted{}
}

/* CreatePeerAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreatePeerAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.PeerSection
}

func (o *CreatePeerAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/peer_section][%d] createPeerAccepted  %+v", 202, o.Payload)
}
func (o *CreatePeerAccepted) GetPayload() *models.PeerSection {
	return o.Payload
}

func (o *CreatePeerAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.PeerSection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePeerBadRequest creates a CreatePeerBadRequest with default headers values
func NewCreatePeerBadRequest() *CreatePeerBadRequest {
	return &CreatePeerBadRequest{}
}

/* CreatePeerBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreatePeerBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreatePeerBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/peer_section][%d] createPeerBadRequest  %+v", 400, o.Payload)
}
func (o *CreatePeerBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreatePeerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreatePeerConflict creates a CreatePeerConflict with default headers values
func NewCreatePeerConflict() *CreatePeerConflict {
	return &CreatePeerConflict{}
}

/* CreatePeerConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreatePeerConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreatePeerConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/peer_section][%d] createPeerConflict  %+v", 409, o.Payload)
}
func (o *CreatePeerConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreatePeerConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreatePeerDefault creates a CreatePeerDefault with default headers values
func NewCreatePeerDefault(code int) *CreatePeerDefault {
	return &CreatePeerDefault{
		_statusCode: code,
	}
}

/* CreatePeerDefault describes a response with status code -1, with default header values.

General Error
*/
type CreatePeerDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create peer default response
func (o *CreatePeerDefault) Code() int {
	return o._statusCode
}

func (o *CreatePeerDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/peer_section][%d] createPeer default  %+v", o._statusCode, o.Payload)
}
func (o *CreatePeerDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreatePeerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
