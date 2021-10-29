// Code generated by go-swagger; DO NOT EDIT.

package maps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/practical-coder/hdc/models"
)

// ReplaceRuntimeMapEntryReader is a Reader for the ReplaceRuntimeMapEntry structure.
type ReplaceRuntimeMapEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceRuntimeMapEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceRuntimeMapEntryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceRuntimeMapEntryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceRuntimeMapEntryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceRuntimeMapEntryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceRuntimeMapEntryOK creates a ReplaceRuntimeMapEntryOK with default headers values
func NewReplaceRuntimeMapEntryOK() *ReplaceRuntimeMapEntryOK {
	return &ReplaceRuntimeMapEntryOK{}
}

/* ReplaceRuntimeMapEntryOK describes a response with status code 200, with default header values.

Map value replaced
*/
type ReplaceRuntimeMapEntryOK struct {
	Payload *models.MapEntry
}

func (o *ReplaceRuntimeMapEntryOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/maps_entries/{id}][%d] replaceRuntimeMapEntryOK  %+v", 200, o.Payload)
}
func (o *ReplaceRuntimeMapEntryOK) GetPayload() *models.MapEntry {
	return o.Payload
}

func (o *ReplaceRuntimeMapEntryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MapEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceRuntimeMapEntryBadRequest creates a ReplaceRuntimeMapEntryBadRequest with default headers values
func NewReplaceRuntimeMapEntryBadRequest() *ReplaceRuntimeMapEntryBadRequest {
	return &ReplaceRuntimeMapEntryBadRequest{}
}

/* ReplaceRuntimeMapEntryBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceRuntimeMapEntryBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceRuntimeMapEntryBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/maps_entries/{id}][%d] replaceRuntimeMapEntryBadRequest  %+v", 400, o.Payload)
}
func (o *ReplaceRuntimeMapEntryBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceRuntimeMapEntryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceRuntimeMapEntryNotFound creates a ReplaceRuntimeMapEntryNotFound with default headers values
func NewReplaceRuntimeMapEntryNotFound() *ReplaceRuntimeMapEntryNotFound {
	return &ReplaceRuntimeMapEntryNotFound{}
}

/* ReplaceRuntimeMapEntryNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type ReplaceRuntimeMapEntryNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceRuntimeMapEntryNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/maps_entries/{id}][%d] replaceRuntimeMapEntryNotFound  %+v", 404, o.Payload)
}
func (o *ReplaceRuntimeMapEntryNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceRuntimeMapEntryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceRuntimeMapEntryDefault creates a ReplaceRuntimeMapEntryDefault with default headers values
func NewReplaceRuntimeMapEntryDefault(code int) *ReplaceRuntimeMapEntryDefault {
	return &ReplaceRuntimeMapEntryDefault{
		_statusCode: code,
	}
}

/* ReplaceRuntimeMapEntryDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceRuntimeMapEntryDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the replace runtime map entry default response
func (o *ReplaceRuntimeMapEntryDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceRuntimeMapEntryDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/maps_entries/{id}][%d] replaceRuntimeMapEntry default  %+v", o._statusCode, o.Payload)
}
func (o *ReplaceRuntimeMapEntryDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceRuntimeMapEntryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*ReplaceRuntimeMapEntryBody replace runtime map entry body
swagger:model ReplaceRuntimeMapEntryBody
*/
type ReplaceRuntimeMapEntryBody struct {

	// Map value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this replace runtime map entry body
func (o *ReplaceRuntimeMapEntryBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ReplaceRuntimeMapEntryBody) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this replace runtime map entry body based on context it is used
func (o *ReplaceRuntimeMapEntryBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ReplaceRuntimeMapEntryBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ReplaceRuntimeMapEntryBody) UnmarshalBinary(b []byte) error {
	var res ReplaceRuntimeMapEntryBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
