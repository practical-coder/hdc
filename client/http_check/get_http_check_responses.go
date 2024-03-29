// Code generated by go-swagger; DO NOT EDIT.

package http_check

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

// GetHTTPCheckReader is a Reader for the GetHTTPCheck structure.
type GetHTTPCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHTTPCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHTTPCheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetHTTPCheckNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetHTTPCheckDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHTTPCheckOK creates a GetHTTPCheckOK with default headers values
func NewGetHTTPCheckOK() *GetHTTPCheckOK {
	return &GetHTTPCheckOK{}
}

/*
GetHTTPCheckOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetHTTPCheckOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetHTTPCheckOKBody
}

// IsSuccess returns true when this get Http check o k response has a 2xx status code
func (o *GetHTTPCheckOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Http check o k response has a 3xx status code
func (o *GetHTTPCheckOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Http check o k response has a 4xx status code
func (o *GetHTTPCheckOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Http check o k response has a 5xx status code
func (o *GetHTTPCheckOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Http check o k response a status code equal to that given
func (o *GetHTTPCheckOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Http check o k response
func (o *GetHTTPCheckOK) Code() int {
	return 200
}

func (o *GetHTTPCheckOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_checks/{index}][%d] getHttpCheckOK  %+v", 200, o.Payload)
}

func (o *GetHTTPCheckOK) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_checks/{index}][%d] getHttpCheckOK  %+v", 200, o.Payload)
}

func (o *GetHTTPCheckOK) GetPayload() *GetHTTPCheckOKBody {
	return o.Payload
}

func (o *GetHTTPCheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetHTTPCheckOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHTTPCheckNotFound creates a GetHTTPCheckNotFound with default headers values
func NewGetHTTPCheckNotFound() *GetHTTPCheckNotFound {
	return &GetHTTPCheckNotFound{}
}

/*
GetHTTPCheckNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetHTTPCheckNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get Http check not found response has a 2xx status code
func (o *GetHTTPCheckNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Http check not found response has a 3xx status code
func (o *GetHTTPCheckNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Http check not found response has a 4xx status code
func (o *GetHTTPCheckNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Http check not found response has a 5xx status code
func (o *GetHTTPCheckNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get Http check not found response a status code equal to that given
func (o *GetHTTPCheckNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get Http check not found response
func (o *GetHTTPCheckNotFound) Code() int {
	return 404
}

func (o *GetHTTPCheckNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_checks/{index}][%d] getHttpCheckNotFound  %+v", 404, o.Payload)
}

func (o *GetHTTPCheckNotFound) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_checks/{index}][%d] getHttpCheckNotFound  %+v", 404, o.Payload)
}

func (o *GetHTTPCheckNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHTTPCheckNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetHTTPCheckDefault creates a GetHTTPCheckDefault with default headers values
func NewGetHTTPCheckDefault(code int) *GetHTTPCheckDefault {
	return &GetHTTPCheckDefault{
		_statusCode: code,
	}
}

/*
GetHTTPCheckDefault describes a response with status code -1, with default header values.

General Error
*/
type GetHTTPCheckDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get HTTP check default response has a 2xx status code
func (o *GetHTTPCheckDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get HTTP check default response has a 3xx status code
func (o *GetHTTPCheckDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get HTTP check default response has a 4xx status code
func (o *GetHTTPCheckDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get HTTP check default response has a 5xx status code
func (o *GetHTTPCheckDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get HTTP check default response a status code equal to that given
func (o *GetHTTPCheckDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get HTTP check default response
func (o *GetHTTPCheckDefault) Code() int {
	return o._statusCode
}

func (o *GetHTTPCheckDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_checks/{index}][%d] getHTTPCheck default  %+v", o._statusCode, o.Payload)
}

func (o *GetHTTPCheckDefault) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_checks/{index}][%d] getHTTPCheck default  %+v", o._statusCode, o.Payload)
}

func (o *GetHTTPCheckDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHTTPCheckDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetHTTPCheckOKBody get HTTP check o k body
swagger:model GetHTTPCheckOKBody
*/
type GetHTTPCheckOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.HTTPCheck `json:"data,omitempty"`
}

// Validate validates this get HTTP check o k body
func (o *GetHTTPCheckOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHTTPCheckOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getHttpCheckOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getHttpCheckOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get HTTP check o k body based on the context it is used
func (o *GetHTTPCheckOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHTTPCheckOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getHttpCheckOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getHttpCheckOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetHTTPCheckOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHTTPCheckOKBody) UnmarshalBinary(b []byte) error {
	var res GetHTTPCheckOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
