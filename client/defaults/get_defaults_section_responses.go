// Code generated by go-swagger; DO NOT EDIT.

package defaults

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

	"github.com/haproxytech/client-native/v4/models"
)

// GetDefaultsSectionReader is a Reader for the GetDefaultsSection structure.
type GetDefaultsSectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDefaultsSectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDefaultsSectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDefaultsSectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetDefaultsSectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDefaultsSectionOK creates a GetDefaultsSectionOK with default headers values
func NewGetDefaultsSectionOK() *GetDefaultsSectionOK {
	return &GetDefaultsSectionOK{}
}

/*
GetDefaultsSectionOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetDefaultsSectionOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetDefaultsSectionOKBody
}

// IsSuccess returns true when this get defaults section o k response has a 2xx status code
func (o *GetDefaultsSectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get defaults section o k response has a 3xx status code
func (o *GetDefaultsSectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get defaults section o k response has a 4xx status code
func (o *GetDefaultsSectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get defaults section o k response has a 5xx status code
func (o *GetDefaultsSectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get defaults section o k response a status code equal to that given
func (o *GetDefaultsSectionOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDefaultsSectionOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/named_defaults/{name}][%d] getDefaultsSectionOK  %+v", 200, o.Payload)
}

func (o *GetDefaultsSectionOK) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/named_defaults/{name}][%d] getDefaultsSectionOK  %+v", 200, o.Payload)
}

func (o *GetDefaultsSectionOK) GetPayload() *GetDefaultsSectionOKBody {
	return o.Payload
}

func (o *GetDefaultsSectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetDefaultsSectionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDefaultsSectionNotFound creates a GetDefaultsSectionNotFound with default headers values
func NewGetDefaultsSectionNotFound() *GetDefaultsSectionNotFound {
	return &GetDefaultsSectionNotFound{}
}

/*
GetDefaultsSectionNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetDefaultsSectionNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get defaults section not found response has a 2xx status code
func (o *GetDefaultsSectionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get defaults section not found response has a 3xx status code
func (o *GetDefaultsSectionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get defaults section not found response has a 4xx status code
func (o *GetDefaultsSectionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get defaults section not found response has a 5xx status code
func (o *GetDefaultsSectionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get defaults section not found response a status code equal to that given
func (o *GetDefaultsSectionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetDefaultsSectionNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/named_defaults/{name}][%d] getDefaultsSectionNotFound  %+v", 404, o.Payload)
}

func (o *GetDefaultsSectionNotFound) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/named_defaults/{name}][%d] getDefaultsSectionNotFound  %+v", 404, o.Payload)
}

func (o *GetDefaultsSectionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDefaultsSectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDefaultsSectionDefault creates a GetDefaultsSectionDefault with default headers values
func NewGetDefaultsSectionDefault(code int) *GetDefaultsSectionDefault {
	return &GetDefaultsSectionDefault{
		_statusCode: code,
	}
}

/*
GetDefaultsSectionDefault describes a response with status code -1, with default header values.

General Error
*/
type GetDefaultsSectionDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get defaults section default response
func (o *GetDefaultsSectionDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get defaults section default response has a 2xx status code
func (o *GetDefaultsSectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get defaults section default response has a 3xx status code
func (o *GetDefaultsSectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get defaults section default response has a 4xx status code
func (o *GetDefaultsSectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get defaults section default response has a 5xx status code
func (o *GetDefaultsSectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get defaults section default response a status code equal to that given
func (o *GetDefaultsSectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetDefaultsSectionDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/named_defaults/{name}][%d] getDefaultsSection default  %+v", o._statusCode, o.Payload)
}

func (o *GetDefaultsSectionDefault) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/named_defaults/{name}][%d] getDefaultsSection default  %+v", o._statusCode, o.Payload)
}

func (o *GetDefaultsSectionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDefaultsSectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetDefaultsSectionOKBody get defaults section o k body
swagger:model GetDefaultsSectionOKBody
*/
type GetDefaultsSectionOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.Defaults `json:"data,omitempty"`
}

// Validate validates this get defaults section o k body
func (o *GetDefaultsSectionOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDefaultsSectionOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getDefaultsSectionOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getDefaultsSectionOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get defaults section o k body based on the context it is used
func (o *GetDefaultsSectionOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDefaultsSectionOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getDefaultsSectionOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getDefaultsSectionOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDefaultsSectionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDefaultsSectionOKBody) UnmarshalBinary(b []byte) error {
	var res GetDefaultsSectionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}