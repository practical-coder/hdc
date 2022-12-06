// Code generated by go-swagger; DO NOT EDIT.

package global

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

// GetGlobalReader is a Reader for the GetGlobal structure.
type GetGlobalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGlobalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGlobalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetGlobalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetGlobalOK creates a GetGlobalOK with default headers values
func NewGetGlobalOK() *GetGlobalOK {
	return &GetGlobalOK{}
}

/*
GetGlobalOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetGlobalOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetGlobalOKBody
}

// IsSuccess returns true when this get global o k response has a 2xx status code
func (o *GetGlobalOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get global o k response has a 3xx status code
func (o *GetGlobalOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get global o k response has a 4xx status code
func (o *GetGlobalOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get global o k response has a 5xx status code
func (o *GetGlobalOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get global o k response a status code equal to that given
func (o *GetGlobalOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetGlobalOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/global][%d] getGlobalOK  %+v", 200, o.Payload)
}

func (o *GetGlobalOK) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/global][%d] getGlobalOK  %+v", 200, o.Payload)
}

func (o *GetGlobalOK) GetPayload() *GetGlobalOKBody {
	return o.Payload
}

func (o *GetGlobalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetGlobalOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGlobalDefault creates a GetGlobalDefault with default headers values
func NewGetGlobalDefault(code int) *GetGlobalDefault {
	return &GetGlobalDefault{
		_statusCode: code,
	}
}

/*
GetGlobalDefault describes a response with status code -1, with default header values.

General Error
*/
type GetGlobalDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get global default response
func (o *GetGlobalDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get global default response has a 2xx status code
func (o *GetGlobalDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get global default response has a 3xx status code
func (o *GetGlobalDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get global default response has a 4xx status code
func (o *GetGlobalDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get global default response has a 5xx status code
func (o *GetGlobalDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get global default response a status code equal to that given
func (o *GetGlobalDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetGlobalDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/global][%d] getGlobal default  %+v", o._statusCode, o.Payload)
}

func (o *GetGlobalDefault) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/global][%d] getGlobal default  %+v", o._statusCode, o.Payload)
}

func (o *GetGlobalDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGlobalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetGlobalOKBody get global o k body
swagger:model GetGlobalOKBody
*/
type GetGlobalOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.Global `json:"data,omitempty"`
}

// Validate validates this get global o k body
func (o *GetGlobalOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetGlobalOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getGlobalOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getGlobalOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get global o k body based on the context it is used
func (o *GetGlobalOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetGlobalOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getGlobalOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getGlobalOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetGlobalOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetGlobalOKBody) UnmarshalBinary(b []byte) error {
	var res GetGlobalOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
