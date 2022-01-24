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

// ReplacePeerEntryReader is a Reader for the ReplacePeerEntry structure.
type ReplacePeerEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplacePeerEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplacePeerEntryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplacePeerEntryAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplacePeerEntryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplacePeerEntryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplacePeerEntryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplacePeerEntryOK creates a ReplacePeerEntryOK with default headers values
func NewReplacePeerEntryOK() *ReplacePeerEntryOK {
	return &ReplacePeerEntryOK{}
}

/*ReplacePeerEntryOK handles this case with default header values.

PeerEntry replaced
*/
type ReplacePeerEntryOK struct {
	Payload *models.PeerEntry
}

func (o *ReplacePeerEntryOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/peer_entries/{name}][%d] replacePeerEntryOK  %+v", 200, o.Payload)
}

func (o *ReplacePeerEntryOK) GetPayload() *models.PeerEntry {
	return o.Payload
}

func (o *ReplacePeerEntryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PeerEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplacePeerEntryAccepted creates a ReplacePeerEntryAccepted with default headers values
func NewReplacePeerEntryAccepted() *ReplacePeerEntryAccepted {
	return &ReplacePeerEntryAccepted{}
}

/*ReplacePeerEntryAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type ReplacePeerEntryAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string

	Payload *models.PeerEntry
}

func (o *ReplacePeerEntryAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/peer_entries/{name}][%d] replacePeerEntryAccepted  %+v", 202, o.Payload)
}

func (o *ReplacePeerEntryAccepted) GetPayload() *models.PeerEntry {
	return o.Payload
}

func (o *ReplacePeerEntryAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	o.Payload = new(models.PeerEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplacePeerEntryBadRequest creates a ReplacePeerEntryBadRequest with default headers values
func NewReplacePeerEntryBadRequest() *ReplacePeerEntryBadRequest {
	return &ReplacePeerEntryBadRequest{}
}

/*ReplacePeerEntryBadRequest handles this case with default header values.

Bad request
*/
type ReplacePeerEntryBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplacePeerEntryBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/peer_entries/{name}][%d] replacePeerEntryBadRequest  %+v", 400, o.Payload)
}

func (o *ReplacePeerEntryBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplacePeerEntryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplacePeerEntryNotFound creates a ReplacePeerEntryNotFound with default headers values
func NewReplacePeerEntryNotFound() *ReplacePeerEntryNotFound {
	return &ReplacePeerEntryNotFound{}
}

/*ReplacePeerEntryNotFound handles this case with default header values.

The specified resource was not found
*/
type ReplacePeerEntryNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplacePeerEntryNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/peer_entries/{name}][%d] replacePeerEntryNotFound  %+v", 404, o.Payload)
}

func (o *ReplacePeerEntryNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplacePeerEntryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplacePeerEntryDefault creates a ReplacePeerEntryDefault with default headers values
func NewReplacePeerEntryDefault(code int) *ReplacePeerEntryDefault {
	return &ReplacePeerEntryDefault{
		_statusCode: code,
	}
}

/*ReplacePeerEntryDefault handles this case with default header values.

General Error
*/
type ReplacePeerEntryDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the replace peer entry default response
func (o *ReplacePeerEntryDefault) Code() int {
	return o._statusCode
}

func (o *ReplacePeerEntryDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/peer_entries/{name}][%d] replacePeerEntry default  %+v", o._statusCode, o.Payload)
}

func (o *ReplacePeerEntryDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplacePeerEntryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
