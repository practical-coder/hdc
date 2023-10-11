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

	"github.com/haproxytech/client-native/v5/models"
)

// GetSpoeGroupsReader is a Reader for the GetSpoeGroups structure.
type GetSpoeGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpoeGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSpoeGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSpoeGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSpoeGroupsOK creates a GetSpoeGroupsOK with default headers values
func NewGetSpoeGroupsOK() *GetSpoeGroupsOK {
	return &GetSpoeGroupsOK{}
}

/*
GetSpoeGroupsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetSpoeGroupsOK struct {

	/* Spoe configuration file version
	 */
	ConfigurationVersion string

	Payload *GetSpoeGroupsOKBody
}

// IsSuccess returns true when this get spoe groups o k response has a 2xx status code
func (o *GetSpoeGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get spoe groups o k response has a 3xx status code
func (o *GetSpoeGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get spoe groups o k response has a 4xx status code
func (o *GetSpoeGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get spoe groups o k response has a 5xx status code
func (o *GetSpoeGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get spoe groups o k response a status code equal to that given
func (o *GetSpoeGroupsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get spoe groups o k response
func (o *GetSpoeGroupsOK) Code() int {
	return 200
}

func (o *GetSpoeGroupsOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_groups][%d] getSpoeGroupsOK  %+v", 200, o.Payload)
}

func (o *GetSpoeGroupsOK) String() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_groups][%d] getSpoeGroupsOK  %+v", 200, o.Payload)
}

func (o *GetSpoeGroupsOK) GetPayload() *GetSpoeGroupsOKBody {
	return o.Payload
}

func (o *GetSpoeGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetSpoeGroupsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpoeGroupsDefault creates a GetSpoeGroupsDefault with default headers values
func NewGetSpoeGroupsDefault(code int) *GetSpoeGroupsDefault {
	return &GetSpoeGroupsDefault{
		_statusCode: code,
	}
}

/*
GetSpoeGroupsDefault describes a response with status code -1, with default header values.

General Error
*/
type GetSpoeGroupsDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get spoe groups default response has a 2xx status code
func (o *GetSpoeGroupsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get spoe groups default response has a 3xx status code
func (o *GetSpoeGroupsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get spoe groups default response has a 4xx status code
func (o *GetSpoeGroupsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get spoe groups default response has a 5xx status code
func (o *GetSpoeGroupsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get spoe groups default response a status code equal to that given
func (o *GetSpoeGroupsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get spoe groups default response
func (o *GetSpoeGroupsDefault) Code() int {
	return o._statusCode
}

func (o *GetSpoeGroupsDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_groups][%d] getSpoeGroups default  %+v", o._statusCode, o.Payload)
}

func (o *GetSpoeGroupsDefault) String() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_groups][%d] getSpoeGroups default  %+v", o._statusCode, o.Payload)
}

func (o *GetSpoeGroupsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSpoeGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetSpoeGroupsOKBody get spoe groups o k body
swagger:model GetSpoeGroupsOKBody
*/
type GetSpoeGroupsOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.SpoeGroups `json:"data"`
}

// Validate validates this get spoe groups o k body
func (o *GetSpoeGroupsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSpoeGroupsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getSpoeGroupsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getSpoeGroupsOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getSpoeGroupsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get spoe groups o k body based on the context it is used
func (o *GetSpoeGroupsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSpoeGroupsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getSpoeGroupsOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getSpoeGroupsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSpoeGroupsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSpoeGroupsOKBody) UnmarshalBinary(b []byte) error {
	var res GetSpoeGroupsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
