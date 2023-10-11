// Code generated by go-swagger; DO NOT EDIT.

package filter

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

	"github.com/haproxytech/client-native/v5/models"
)

// GetFiltersReader is a Reader for the GetFilters structure.
type GetFiltersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFiltersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFiltersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetFiltersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFiltersOK creates a GetFiltersOK with default headers values
func NewGetFiltersOK() *GetFiltersOK {
	return &GetFiltersOK{}
}

/*
GetFiltersOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetFiltersOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetFiltersOKBody
}

// IsSuccess returns true when this get filters o k response has a 2xx status code
func (o *GetFiltersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get filters o k response has a 3xx status code
func (o *GetFiltersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get filters o k response has a 4xx status code
func (o *GetFiltersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get filters o k response has a 5xx status code
func (o *GetFiltersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get filters o k response a status code equal to that given
func (o *GetFiltersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get filters o k response
func (o *GetFiltersOK) Code() int {
	return 200
}

func (o *GetFiltersOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/filters][%d] getFiltersOK  %+v", 200, o.Payload)
}

func (o *GetFiltersOK) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/filters][%d] getFiltersOK  %+v", 200, o.Payload)
}

func (o *GetFiltersOK) GetPayload() *GetFiltersOKBody {
	return o.Payload
}

func (o *GetFiltersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetFiltersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFiltersDefault creates a GetFiltersDefault with default headers values
func NewGetFiltersDefault(code int) *GetFiltersDefault {
	return &GetFiltersDefault{
		_statusCode: code,
	}
}

/*
GetFiltersDefault describes a response with status code -1, with default header values.

General Error
*/
type GetFiltersDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get filters default response has a 2xx status code
func (o *GetFiltersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get filters default response has a 3xx status code
func (o *GetFiltersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get filters default response has a 4xx status code
func (o *GetFiltersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get filters default response has a 5xx status code
func (o *GetFiltersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get filters default response a status code equal to that given
func (o *GetFiltersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get filters default response
func (o *GetFiltersDefault) Code() int {
	return o._statusCode
}

func (o *GetFiltersDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/filters][%d] getFilters default  %+v", o._statusCode, o.Payload)
}

func (o *GetFiltersDefault) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/filters][%d] getFilters default  %+v", o._statusCode, o.Payload)
}

func (o *GetFiltersDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFiltersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetFiltersOKBody get filters o k body
swagger:model GetFiltersOKBody
*/
type GetFiltersOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.Filters `json:"data"`
}

// Validate validates this get filters o k body
func (o *GetFiltersOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFiltersOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getFiltersOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getFiltersOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getFiltersOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get filters o k body based on the context it is used
func (o *GetFiltersOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFiltersOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getFiltersOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getFiltersOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetFiltersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFiltersOKBody) UnmarshalBinary(b []byte) error {
	var res GetFiltersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
