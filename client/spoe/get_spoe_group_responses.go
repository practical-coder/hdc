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

// GetSpoeGroupReader is a Reader for the GetSpoeGroup structure.
type GetSpoeGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpoeGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSpoeGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSpoeGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetSpoeGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSpoeGroupOK creates a GetSpoeGroupOK with default headers values
func NewGetSpoeGroupOK() *GetSpoeGroupOK {
	return &GetSpoeGroupOK{}
}

/*
GetSpoeGroupOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetSpoeGroupOK struct {

	/* Spoe configuration file version
	 */
	ConfigurationVersion string

	Payload *GetSpoeGroupOKBody
}

// IsSuccess returns true when this get spoe group o k response has a 2xx status code
func (o *GetSpoeGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get spoe group o k response has a 3xx status code
func (o *GetSpoeGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get spoe group o k response has a 4xx status code
func (o *GetSpoeGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get spoe group o k response has a 5xx status code
func (o *GetSpoeGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get spoe group o k response a status code equal to that given
func (o *GetSpoeGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get spoe group o k response
func (o *GetSpoeGroupOK) Code() int {
	return 200
}

func (o *GetSpoeGroupOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_groups/{name}][%d] getSpoeGroupOK  %+v", 200, o.Payload)
}

func (o *GetSpoeGroupOK) String() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_groups/{name}][%d] getSpoeGroupOK  %+v", 200, o.Payload)
}

func (o *GetSpoeGroupOK) GetPayload() *GetSpoeGroupOKBody {
	return o.Payload
}

func (o *GetSpoeGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetSpoeGroupOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpoeGroupNotFound creates a GetSpoeGroupNotFound with default headers values
func NewGetSpoeGroupNotFound() *GetSpoeGroupNotFound {
	return &GetSpoeGroupNotFound{}
}

/*
GetSpoeGroupNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetSpoeGroupNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get spoe group not found response has a 2xx status code
func (o *GetSpoeGroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get spoe group not found response has a 3xx status code
func (o *GetSpoeGroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get spoe group not found response has a 4xx status code
func (o *GetSpoeGroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get spoe group not found response has a 5xx status code
func (o *GetSpoeGroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get spoe group not found response a status code equal to that given
func (o *GetSpoeGroupNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get spoe group not found response
func (o *GetSpoeGroupNotFound) Code() int {
	return 404
}

func (o *GetSpoeGroupNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_groups/{name}][%d] getSpoeGroupNotFound  %+v", 404, o.Payload)
}

func (o *GetSpoeGroupNotFound) String() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_groups/{name}][%d] getSpoeGroupNotFound  %+v", 404, o.Payload)
}

func (o *GetSpoeGroupNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSpoeGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSpoeGroupDefault creates a GetSpoeGroupDefault with default headers values
func NewGetSpoeGroupDefault(code int) *GetSpoeGroupDefault {
	return &GetSpoeGroupDefault{
		_statusCode: code,
	}
}

/*
GetSpoeGroupDefault describes a response with status code -1, with default header values.

General Error
*/
type GetSpoeGroupDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get spoe group default response has a 2xx status code
func (o *GetSpoeGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get spoe group default response has a 3xx status code
func (o *GetSpoeGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get spoe group default response has a 4xx status code
func (o *GetSpoeGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get spoe group default response has a 5xx status code
func (o *GetSpoeGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get spoe group default response a status code equal to that given
func (o *GetSpoeGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get spoe group default response
func (o *GetSpoeGroupDefault) Code() int {
	return o._statusCode
}

func (o *GetSpoeGroupDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_groups/{name}][%d] getSpoeGroup default  %+v", o._statusCode, o.Payload)
}

func (o *GetSpoeGroupDefault) String() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_groups/{name}][%d] getSpoeGroup default  %+v", o._statusCode, o.Payload)
}

func (o *GetSpoeGroupDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSpoeGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetSpoeGroupOKBody get spoe group o k body
swagger:model GetSpoeGroupOKBody
*/
type GetSpoeGroupOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data *models.SpoeGroup `json:"data"`
}

// Validate validates this get spoe group o k body
func (o *GetSpoeGroupOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSpoeGroupOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getSpoeGroupOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSpoeGroupOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getSpoeGroupOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get spoe group o k body based on the context it is used
func (o *GetSpoeGroupOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSpoeGroupOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSpoeGroupOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getSpoeGroupOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSpoeGroupOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSpoeGroupOKBody) UnmarshalBinary(b []byte) error {
	var res GetSpoeGroupOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
