// Code generated by go-swagger; DO NOT EDIT.

package spoe

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

	"github.com/haproxytech/client-native/v4/models"
)

// GetSpoeScopeReader is a Reader for the GetSpoeScope structure.
type GetSpoeScopeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpoeScopeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSpoeScopeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSpoeScopeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetSpoeScopeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSpoeScopeOK creates a GetSpoeScopeOK with default headers values
func NewGetSpoeScopeOK() *GetSpoeScopeOK {
	return &GetSpoeScopeOK{}
}

/*
GetSpoeScopeOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetSpoeScopeOK struct {

	/* Spoe configuration file version
	 */
	ConfigurationVersion string

	Payload *GetSpoeScopeOKBody
}

// IsSuccess returns true when this get spoe scope o k response has a 2xx status code
func (o *GetSpoeScopeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get spoe scope o k response has a 3xx status code
func (o *GetSpoeScopeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get spoe scope o k response has a 4xx status code
func (o *GetSpoeScopeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get spoe scope o k response has a 5xx status code
func (o *GetSpoeScopeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get spoe scope o k response a status code equal to that given
func (o *GetSpoeScopeOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetSpoeScopeOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_scopes/{name}][%d] getSpoeScopeOK  %+v", 200, o.Payload)
}

func (o *GetSpoeScopeOK) String() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_scopes/{name}][%d] getSpoeScopeOK  %+v", 200, o.Payload)
}

func (o *GetSpoeScopeOK) GetPayload() *GetSpoeScopeOKBody {
	return o.Payload
}

func (o *GetSpoeScopeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetSpoeScopeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpoeScopeNotFound creates a GetSpoeScopeNotFound with default headers values
func NewGetSpoeScopeNotFound() *GetSpoeScopeNotFound {
	return &GetSpoeScopeNotFound{}
}

/*
GetSpoeScopeNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetSpoeScopeNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get spoe scope not found response has a 2xx status code
func (o *GetSpoeScopeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get spoe scope not found response has a 3xx status code
func (o *GetSpoeScopeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get spoe scope not found response has a 4xx status code
func (o *GetSpoeScopeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get spoe scope not found response has a 5xx status code
func (o *GetSpoeScopeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get spoe scope not found response a status code equal to that given
func (o *GetSpoeScopeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetSpoeScopeNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_scopes/{name}][%d] getSpoeScopeNotFound  %+v", 404, o.Payload)
}

func (o *GetSpoeScopeNotFound) String() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_scopes/{name}][%d] getSpoeScopeNotFound  %+v", 404, o.Payload)
}

func (o *GetSpoeScopeNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSpoeScopeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSpoeScopeDefault creates a GetSpoeScopeDefault with default headers values
func NewGetSpoeScopeDefault(code int) *GetSpoeScopeDefault {
	return &GetSpoeScopeDefault{
		_statusCode: code,
	}
}

/*
GetSpoeScopeDefault describes a response with status code -1, with default header values.

General Error
*/
type GetSpoeScopeDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get spoe scope default response
func (o *GetSpoeScopeDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get spoe scope default response has a 2xx status code
func (o *GetSpoeScopeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get spoe scope default response has a 3xx status code
func (o *GetSpoeScopeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get spoe scope default response has a 4xx status code
func (o *GetSpoeScopeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get spoe scope default response has a 5xx status code
func (o *GetSpoeScopeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get spoe scope default response a status code equal to that given
func (o *GetSpoeScopeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetSpoeScopeDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_scopes/{name}][%d] getSpoeScope default  %+v", o._statusCode, o.Payload)
}

func (o *GetSpoeScopeDefault) String() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_scopes/{name}][%d] getSpoeScope default  %+v", o._statusCode, o.Payload)
}

func (o *GetSpoeScopeDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSpoeScopeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
GetSpoeScopeOKBody get spoe scope o k body
swagger:model GetSpoeScopeOKBody
*/
type GetSpoeScopeOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data *models.SpoeScope `json:"data"`
}

// Validate validates this get spoe scope o k body
func (o *GetSpoeScopeOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSpoeScopeOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getSpoeScopeOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := validate.Required("getSpoeScopeOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSpoeScopeOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getSpoeScopeOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get spoe scope o k body based on the context it is used
func (o *GetSpoeScopeOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSpoeScopeOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSpoeScopeOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getSpoeScopeOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSpoeScopeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSpoeScopeOKBody) UnmarshalBinary(b []byte) error {
	var res GetSpoeScopeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
