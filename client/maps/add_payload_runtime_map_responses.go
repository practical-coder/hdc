// Code generated by go-swagger; DO NOT EDIT.

package maps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// AddPayloadRuntimeMapReader is a Reader for the AddPayloadRuntimeMap structure.
type AddPayloadRuntimeMapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddPayloadRuntimeMapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddPayloadRuntimeMapCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddPayloadRuntimeMapBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddPayloadRuntimeMapDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddPayloadRuntimeMapCreated creates a AddPayloadRuntimeMapCreated with default headers values
func NewAddPayloadRuntimeMapCreated() *AddPayloadRuntimeMapCreated {
	return &AddPayloadRuntimeMapCreated{}
}

/* AddPayloadRuntimeMapCreated describes a response with status code 201, with default header values.

Map payload added
*/
type AddPayloadRuntimeMapCreated struct {
	Payload models.MapEntries
}

func (o *AddPayloadRuntimeMapCreated) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/maps/{name}][%d] addPayloadRuntimeMapCreated  %+v", 201, o.Payload)
}
func (o *AddPayloadRuntimeMapCreated) GetPayload() models.MapEntries {
	return o.Payload
}

func (o *AddPayloadRuntimeMapCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPayloadRuntimeMapBadRequest creates a AddPayloadRuntimeMapBadRequest with default headers values
func NewAddPayloadRuntimeMapBadRequest() *AddPayloadRuntimeMapBadRequest {
	return &AddPayloadRuntimeMapBadRequest{}
}

/* AddPayloadRuntimeMapBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type AddPayloadRuntimeMapBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *AddPayloadRuntimeMapBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/maps/{name}][%d] addPayloadRuntimeMapBadRequest  %+v", 400, o.Payload)
}
func (o *AddPayloadRuntimeMapBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddPayloadRuntimeMapBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddPayloadRuntimeMapDefault creates a AddPayloadRuntimeMapDefault with default headers values
func NewAddPayloadRuntimeMapDefault(code int) *AddPayloadRuntimeMapDefault {
	return &AddPayloadRuntimeMapDefault{
		_statusCode: code,
	}
}

/* AddPayloadRuntimeMapDefault describes a response with status code -1, with default header values.

General Error
*/
type AddPayloadRuntimeMapDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the add payload runtime map default response
func (o *AddPayloadRuntimeMapDefault) Code() int {
	return o._statusCode
}

func (o *AddPayloadRuntimeMapDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/maps/{name}][%d] addPayloadRuntimeMap default  %+v", o._statusCode, o.Payload)
}
func (o *AddPayloadRuntimeMapDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddPayloadRuntimeMapDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
