// Code generated by go-swagger; DO NOT EDIT.

package peer_entry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
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

/*
CreatePeerEntryCreated describes a response with status code 201, with default header values.

PeerEntry created
*/
type CreatePeerEntryCreated struct {
	Payload *models.PeerEntry
}

// IsSuccess returns true when this create peer entry created response has a 2xx status code
func (o *CreatePeerEntryCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create peer entry created response has a 3xx status code
func (o *CreatePeerEntryCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create peer entry created response has a 4xx status code
func (o *CreatePeerEntryCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create peer entry created response has a 5xx status code
func (o *CreatePeerEntryCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create peer entry created response a status code equal to that given
func (o *CreatePeerEntryCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreatePeerEntryCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/peer_entries][%d] createPeerEntryCreated  %+v", 201, o.Payload)
}

func (o *CreatePeerEntryCreated) String() string {
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

/*
CreatePeerEntryAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreatePeerEntryAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.PeerEntry
}

// IsSuccess returns true when this create peer entry accepted response has a 2xx status code
func (o *CreatePeerEntryAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create peer entry accepted response has a 3xx status code
func (o *CreatePeerEntryAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create peer entry accepted response has a 4xx status code
func (o *CreatePeerEntryAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this create peer entry accepted response has a 5xx status code
func (o *CreatePeerEntryAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this create peer entry accepted response a status code equal to that given
func (o *CreatePeerEntryAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *CreatePeerEntryAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/peer_entries][%d] createPeerEntryAccepted  %+v", 202, o.Payload)
}

func (o *CreatePeerEntryAccepted) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/peer_entries][%d] createPeerEntryAccepted  %+v", 202, o.Payload)
}

func (o *CreatePeerEntryAccepted) GetPayload() *models.PeerEntry {
	return o.Payload
}

func (o *CreatePeerEntryAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

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

/*
CreatePeerEntryBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreatePeerEntryBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create peer entry bad request response has a 2xx status code
func (o *CreatePeerEntryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create peer entry bad request response has a 3xx status code
func (o *CreatePeerEntryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create peer entry bad request response has a 4xx status code
func (o *CreatePeerEntryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create peer entry bad request response has a 5xx status code
func (o *CreatePeerEntryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create peer entry bad request response a status code equal to that given
func (o *CreatePeerEntryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreatePeerEntryBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/peer_entries][%d] createPeerEntryBadRequest  %+v", 400, o.Payload)
}

func (o *CreatePeerEntryBadRequest) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/peer_entries][%d] createPeerEntryBadRequest  %+v", 400, o.Payload)
}

func (o *CreatePeerEntryBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreatePeerEntryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreatePeerEntryConflict creates a CreatePeerEntryConflict with default headers values
func NewCreatePeerEntryConflict() *CreatePeerEntryConflict {
	return &CreatePeerEntryConflict{}
}

/*
CreatePeerEntryConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreatePeerEntryConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create peer entry conflict response has a 2xx status code
func (o *CreatePeerEntryConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create peer entry conflict response has a 3xx status code
func (o *CreatePeerEntryConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create peer entry conflict response has a 4xx status code
func (o *CreatePeerEntryConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create peer entry conflict response has a 5xx status code
func (o *CreatePeerEntryConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create peer entry conflict response a status code equal to that given
func (o *CreatePeerEntryConflict) IsCode(code int) bool {
	return code == 409
}

func (o *CreatePeerEntryConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/peer_entries][%d] createPeerEntryConflict  %+v", 409, o.Payload)
}

func (o *CreatePeerEntryConflict) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/peer_entries][%d] createPeerEntryConflict  %+v", 409, o.Payload)
}

func (o *CreatePeerEntryConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreatePeerEntryConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreatePeerEntryDefault creates a CreatePeerEntryDefault with default headers values
func NewCreatePeerEntryDefault(code int) *CreatePeerEntryDefault {
	return &CreatePeerEntryDefault{
		_statusCode: code,
	}
}

/*
CreatePeerEntryDefault describes a response with status code -1, with default header values.

General Error
*/
type CreatePeerEntryDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create peer entry default response
func (o *CreatePeerEntryDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create peer entry default response has a 2xx status code
func (o *CreatePeerEntryDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create peer entry default response has a 3xx status code
func (o *CreatePeerEntryDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create peer entry default response has a 4xx status code
func (o *CreatePeerEntryDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create peer entry default response has a 5xx status code
func (o *CreatePeerEntryDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create peer entry default response a status code equal to that given
func (o *CreatePeerEntryDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreatePeerEntryDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/peer_entries][%d] createPeerEntry default  %+v", o._statusCode, o.Payload)
}

func (o *CreatePeerEntryDefault) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/peer_entries][%d] createPeerEntry default  %+v", o._statusCode, o.Payload)
}

func (o *CreatePeerEntryDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreatePeerEntryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
