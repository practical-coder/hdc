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

	"github.com/haproxytech/client-native/v5/models"
)

// GetFilterReader is a Reader for the GetFilter structure.
type GetFilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFilterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetFilterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetFilterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFilterOK creates a GetFilterOK with default headers values
func NewGetFilterOK() *GetFilterOK {
	return &GetFilterOK{}
}

/*
GetFilterOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetFilterOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetFilterOKBody
}

// IsSuccess returns true when this get filter o k response has a 2xx status code
func (o *GetFilterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get filter o k response has a 3xx status code
func (o *GetFilterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get filter o k response has a 4xx status code
func (o *GetFilterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get filter o k response has a 5xx status code
func (o *GetFilterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get filter o k response a status code equal to that given
func (o *GetFilterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get filter o k response
func (o *GetFilterOK) Code() int {
	return 200
}

func (o *GetFilterOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/filters/{index}][%d] getFilterOK  %+v", 200, o.Payload)
}

func (o *GetFilterOK) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/filters/{index}][%d] getFilterOK  %+v", 200, o.Payload)
}

func (o *GetFilterOK) GetPayload() *GetFilterOKBody {
	return o.Payload
}

func (o *GetFilterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetFilterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilterNotFound creates a GetFilterNotFound with default headers values
func NewGetFilterNotFound() *GetFilterNotFound {
	return &GetFilterNotFound{}
}

/*
GetFilterNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetFilterNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get filter not found response has a 2xx status code
func (o *GetFilterNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get filter not found response has a 3xx status code
func (o *GetFilterNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get filter not found response has a 4xx status code
func (o *GetFilterNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get filter not found response has a 5xx status code
func (o *GetFilterNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get filter not found response a status code equal to that given
func (o *GetFilterNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get filter not found response
func (o *GetFilterNotFound) Code() int {
	return 404
}

func (o *GetFilterNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/filters/{index}][%d] getFilterNotFound  %+v", 404, o.Payload)
}

func (o *GetFilterNotFound) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/filters/{index}][%d] getFilterNotFound  %+v", 404, o.Payload)
}

func (o *GetFilterNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFilterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFilterDefault creates a GetFilterDefault with default headers values
func NewGetFilterDefault(code int) *GetFilterDefault {
	return &GetFilterDefault{
		_statusCode: code,
	}
}

/*
GetFilterDefault describes a response with status code -1, with default header values.

General Error
*/
type GetFilterDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get filter default response has a 2xx status code
func (o *GetFilterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get filter default response has a 3xx status code
func (o *GetFilterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get filter default response has a 4xx status code
func (o *GetFilterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get filter default response has a 5xx status code
func (o *GetFilterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get filter default response a status code equal to that given
func (o *GetFilterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get filter default response
func (o *GetFilterDefault) Code() int {
	return o._statusCode
}

func (o *GetFilterDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/filters/{index}][%d] getFilter default  %+v", o._statusCode, o.Payload)
}

func (o *GetFilterDefault) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/filters/{index}][%d] getFilter default  %+v", o._statusCode, o.Payload)
}

func (o *GetFilterDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFilterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetFilterOKBody get filter o k body
swagger:model GetFilterOKBody
*/
type GetFilterOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.Filter `json:"data,omitempty"`
}

// Validate validates this get filter o k body
func (o *GetFilterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFilterOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getFilterOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getFilterOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get filter o k body based on the context it is used
func (o *GetFilterOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFilterOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getFilterOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getFilterOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetFilterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFilterOKBody) UnmarshalBinary(b []byte) error {
	var res GetFilterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
