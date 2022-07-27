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

// ShowRuntimeMapReader is a Reader for the ShowRuntimeMap structure.
type ShowRuntimeMapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowRuntimeMapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowRuntimeMapOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewShowRuntimeMapNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewShowRuntimeMapDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewShowRuntimeMapOK creates a ShowRuntimeMapOK with default headers values
func NewShowRuntimeMapOK() *ShowRuntimeMapOK {
	return &ShowRuntimeMapOK{}
}

/* ShowRuntimeMapOK describes a response with status code 200, with default header values.

Successful operation
*/
type ShowRuntimeMapOK struct {
	Payload models.MapEntries
}

func (o *ShowRuntimeMapOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/maps_entries][%d] showRuntimeMapOK  %+v", 200, o.Payload)
}
func (o *ShowRuntimeMapOK) GetPayload() models.MapEntries {
	return o.Payload
}

func (o *ShowRuntimeMapOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowRuntimeMapNotFound creates a ShowRuntimeMapNotFound with default headers values
func NewShowRuntimeMapNotFound() *ShowRuntimeMapNotFound {
	return &ShowRuntimeMapNotFound{}
}

/* ShowRuntimeMapNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type ShowRuntimeMapNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ShowRuntimeMapNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/maps_entries][%d] showRuntimeMapNotFound  %+v", 404, o.Payload)
}
func (o *ShowRuntimeMapNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ShowRuntimeMapNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewShowRuntimeMapDefault creates a ShowRuntimeMapDefault with default headers values
func NewShowRuntimeMapDefault(code int) *ShowRuntimeMapDefault {
	return &ShowRuntimeMapDefault{
		_statusCode: code,
	}
}

/* ShowRuntimeMapDefault describes a response with status code -1, with default header values.

General Error
*/
type ShowRuntimeMapDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the show runtime map default response
func (o *ShowRuntimeMapDefault) Code() int {
	return o._statusCode
}

func (o *ShowRuntimeMapDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/maps_entries][%d] showRuntimeMap default  %+v", o._statusCode, o.Payload)
}
func (o *ShowRuntimeMapDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ShowRuntimeMapDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
