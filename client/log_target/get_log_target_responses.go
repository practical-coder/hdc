// Code generated by go-swagger; DO NOT EDIT.

package log_target

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

// GetLogTargetReader is a Reader for the GetLogTarget structure.
type GetLogTargetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogTargetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogTargetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetLogTargetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetLogTargetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLogTargetOK creates a GetLogTargetOK with default headers values
func NewGetLogTargetOK() *GetLogTargetOK {
	return &GetLogTargetOK{}
}

/*
GetLogTargetOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetLogTargetOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetLogTargetOKBody
}

// IsSuccess returns true when this get log target o k response has a 2xx status code
func (o *GetLogTargetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get log target o k response has a 3xx status code
func (o *GetLogTargetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get log target o k response has a 4xx status code
func (o *GetLogTargetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get log target o k response has a 5xx status code
func (o *GetLogTargetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get log target o k response a status code equal to that given
func (o *GetLogTargetOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetLogTargetOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/log_targets/{index}][%d] getLogTargetOK  %+v", 200, o.Payload)
}

func (o *GetLogTargetOK) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/log_targets/{index}][%d] getLogTargetOK  %+v", 200, o.Payload)
}

func (o *GetLogTargetOK) GetPayload() *GetLogTargetOKBody {
	return o.Payload
}

func (o *GetLogTargetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetLogTargetOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogTargetNotFound creates a GetLogTargetNotFound with default headers values
func NewGetLogTargetNotFound() *GetLogTargetNotFound {
	return &GetLogTargetNotFound{}
}

/*
GetLogTargetNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetLogTargetNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get log target not found response has a 2xx status code
func (o *GetLogTargetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get log target not found response has a 3xx status code
func (o *GetLogTargetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get log target not found response has a 4xx status code
func (o *GetLogTargetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get log target not found response has a 5xx status code
func (o *GetLogTargetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get log target not found response a status code equal to that given
func (o *GetLogTargetNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetLogTargetNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/log_targets/{index}][%d] getLogTargetNotFound  %+v", 404, o.Payload)
}

func (o *GetLogTargetNotFound) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/log_targets/{index}][%d] getLogTargetNotFound  %+v", 404, o.Payload)
}

func (o *GetLogTargetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLogTargetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetLogTargetDefault creates a GetLogTargetDefault with default headers values
func NewGetLogTargetDefault(code int) *GetLogTargetDefault {
	return &GetLogTargetDefault{
		_statusCode: code,
	}
}

/*
GetLogTargetDefault describes a response with status code -1, with default header values.

General Error
*/
type GetLogTargetDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get log target default response
func (o *GetLogTargetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get log target default response has a 2xx status code
func (o *GetLogTargetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get log target default response has a 3xx status code
func (o *GetLogTargetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get log target default response has a 4xx status code
func (o *GetLogTargetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get log target default response has a 5xx status code
func (o *GetLogTargetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get log target default response a status code equal to that given
func (o *GetLogTargetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetLogTargetDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/log_targets/{index}][%d] getLogTarget default  %+v", o._statusCode, o.Payload)
}

func (o *GetLogTargetDefault) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/log_targets/{index}][%d] getLogTarget default  %+v", o._statusCode, o.Payload)
}

func (o *GetLogTargetDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLogTargetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetLogTargetOKBody get log target o k body
swagger:model GetLogTargetOKBody
*/
type GetLogTargetOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.LogTarget `json:"data,omitempty"`
}

// Validate validates this get log target o k body
func (o *GetLogTargetOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLogTargetOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLogTargetOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLogTargetOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get log target o k body based on the context it is used
func (o *GetLogTargetOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLogTargetOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLogTargetOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLogTargetOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLogTargetOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLogTargetOKBody) UnmarshalBinary(b []byte) error {
	var res GetLogTargetOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
