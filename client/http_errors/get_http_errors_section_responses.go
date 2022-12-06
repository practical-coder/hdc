// Code generated by go-swagger; DO NOT EDIT.

package http_errors

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

// GetHTTPErrorsSectionReader is a Reader for the GetHTTPErrorsSection structure.
type GetHTTPErrorsSectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHTTPErrorsSectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHTTPErrorsSectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetHTTPErrorsSectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetHTTPErrorsSectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHTTPErrorsSectionOK creates a GetHTTPErrorsSectionOK with default headers values
func NewGetHTTPErrorsSectionOK() *GetHTTPErrorsSectionOK {
	return &GetHTTPErrorsSectionOK{}
}

/*
GetHTTPErrorsSectionOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetHTTPErrorsSectionOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetHTTPErrorsSectionOKBody
}

// IsSuccess returns true when this get Http errors section o k response has a 2xx status code
func (o *GetHTTPErrorsSectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Http errors section o k response has a 3xx status code
func (o *GetHTTPErrorsSectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Http errors section o k response has a 4xx status code
func (o *GetHTTPErrorsSectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Http errors section o k response has a 5xx status code
func (o *GetHTTPErrorsSectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Http errors section o k response a status code equal to that given
func (o *GetHTTPErrorsSectionOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetHTTPErrorsSectionOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_errors_sections/{name}][%d] getHttpErrorsSectionOK  %+v", 200, o.Payload)
}

func (o *GetHTTPErrorsSectionOK) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_errors_sections/{name}][%d] getHttpErrorsSectionOK  %+v", 200, o.Payload)
}

func (o *GetHTTPErrorsSectionOK) GetPayload() *GetHTTPErrorsSectionOKBody {
	return o.Payload
}

func (o *GetHTTPErrorsSectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetHTTPErrorsSectionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHTTPErrorsSectionNotFound creates a GetHTTPErrorsSectionNotFound with default headers values
func NewGetHTTPErrorsSectionNotFound() *GetHTTPErrorsSectionNotFound {
	return &GetHTTPErrorsSectionNotFound{}
}

/*
GetHTTPErrorsSectionNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetHTTPErrorsSectionNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get Http errors section not found response has a 2xx status code
func (o *GetHTTPErrorsSectionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Http errors section not found response has a 3xx status code
func (o *GetHTTPErrorsSectionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Http errors section not found response has a 4xx status code
func (o *GetHTTPErrorsSectionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Http errors section not found response has a 5xx status code
func (o *GetHTTPErrorsSectionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get Http errors section not found response a status code equal to that given
func (o *GetHTTPErrorsSectionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetHTTPErrorsSectionNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_errors_sections/{name}][%d] getHttpErrorsSectionNotFound  %+v", 404, o.Payload)
}

func (o *GetHTTPErrorsSectionNotFound) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_errors_sections/{name}][%d] getHttpErrorsSectionNotFound  %+v", 404, o.Payload)
}

func (o *GetHTTPErrorsSectionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHTTPErrorsSectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetHTTPErrorsSectionDefault creates a GetHTTPErrorsSectionDefault with default headers values
func NewGetHTTPErrorsSectionDefault(code int) *GetHTTPErrorsSectionDefault {
	return &GetHTTPErrorsSectionDefault{
		_statusCode: code,
	}
}

/*
GetHTTPErrorsSectionDefault describes a response with status code -1, with default header values.

General Error
*/
type GetHTTPErrorsSectionDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get HTTP errors section default response
func (o *GetHTTPErrorsSectionDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get HTTP errors section default response has a 2xx status code
func (o *GetHTTPErrorsSectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get HTTP errors section default response has a 3xx status code
func (o *GetHTTPErrorsSectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get HTTP errors section default response has a 4xx status code
func (o *GetHTTPErrorsSectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get HTTP errors section default response has a 5xx status code
func (o *GetHTTPErrorsSectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get HTTP errors section default response a status code equal to that given
func (o *GetHTTPErrorsSectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetHTTPErrorsSectionDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_errors_sections/{name}][%d] getHTTPErrorsSection default  %+v", o._statusCode, o.Payload)
}

func (o *GetHTTPErrorsSectionDefault) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_errors_sections/{name}][%d] getHTTPErrorsSection default  %+v", o._statusCode, o.Payload)
}

func (o *GetHTTPErrorsSectionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHTTPErrorsSectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetHTTPErrorsSectionOKBody get HTTP errors section o k body
swagger:model GetHTTPErrorsSectionOKBody
*/
type GetHTTPErrorsSectionOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.HTTPErrorsSection `json:"data,omitempty"`
}

// Validate validates this get HTTP errors section o k body
func (o *GetHTTPErrorsSectionOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHTTPErrorsSectionOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getHttpErrorsSectionOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getHttpErrorsSectionOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get HTTP errors section o k body based on the context it is used
func (o *GetHTTPErrorsSectionOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHTTPErrorsSectionOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getHttpErrorsSectionOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getHttpErrorsSectionOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetHTTPErrorsSectionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHTTPErrorsSectionOKBody) UnmarshalBinary(b []byte) error {
	var res GetHTTPErrorsSectionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
