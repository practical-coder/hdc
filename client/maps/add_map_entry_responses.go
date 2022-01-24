// Code generated by go-swagger; DO NOT EDIT.

package maps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v2/models"
)

// AddMapEntryReader is a Reader for the AddMapEntry structure.
type AddMapEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddMapEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddMapEntryCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddMapEntryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddMapEntryConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddMapEntryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddMapEntryCreated creates a AddMapEntryCreated with default headers values
func NewAddMapEntryCreated() *AddMapEntryCreated {
	return &AddMapEntryCreated{}
}

/* AddMapEntryCreated describes a response with status code 201, with default header values.

Map entry created
*/
type AddMapEntryCreated struct {
	Payload *models.MapEntry
}

func (o *AddMapEntryCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/runtime/maps_entries][%d] addMapEntryCreated  %+v", 201, o.Payload)
}
func (o *AddMapEntryCreated) GetPayload() *models.MapEntry {
	return o.Payload
}

func (o *AddMapEntryCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MapEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddMapEntryBadRequest creates a AddMapEntryBadRequest with default headers values
func NewAddMapEntryBadRequest() *AddMapEntryBadRequest {
	return &AddMapEntryBadRequest{}
}

/* AddMapEntryBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type AddMapEntryBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *AddMapEntryBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/runtime/maps_entries][%d] addMapEntryBadRequest  %+v", 400, o.Payload)
}
func (o *AddMapEntryBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddMapEntryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddMapEntryConflict creates a AddMapEntryConflict with default headers values
func NewAddMapEntryConflict() *AddMapEntryConflict {
	return &AddMapEntryConflict{}
}

/* AddMapEntryConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type AddMapEntryConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *AddMapEntryConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/runtime/maps_entries][%d] addMapEntryConflict  %+v", 409, o.Payload)
}
func (o *AddMapEntryConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddMapEntryConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddMapEntryDefault creates a AddMapEntryDefault with default headers values
func NewAddMapEntryDefault(code int) *AddMapEntryDefault {
	return &AddMapEntryDefault{
		_statusCode: code,
	}
}

/* AddMapEntryDefault describes a response with status code -1, with default header values.

General Error
*/
type AddMapEntryDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the add map entry default response
func (o *AddMapEntryDefault) Code() int {
	return o._statusCode
}

func (o *AddMapEntryDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/runtime/maps_entries][%d] addMapEntry default  %+v", o._statusCode, o.Payload)
}
func (o *AddMapEntryDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddMapEntryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
