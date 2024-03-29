// Code generated by go-swagger; DO NOT EDIT.

package http_request_rule

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

// GetHTTPRequestRuleReader is a Reader for the GetHTTPRequestRule structure.
type GetHTTPRequestRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHTTPRequestRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHTTPRequestRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetHTTPRequestRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetHTTPRequestRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHTTPRequestRuleOK creates a GetHTTPRequestRuleOK with default headers values
func NewGetHTTPRequestRuleOK() *GetHTTPRequestRuleOK {
	return &GetHTTPRequestRuleOK{}
}

/*
GetHTTPRequestRuleOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetHTTPRequestRuleOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetHTTPRequestRuleOKBody
}

// IsSuccess returns true when this get Http request rule o k response has a 2xx status code
func (o *GetHTTPRequestRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Http request rule o k response has a 3xx status code
func (o *GetHTTPRequestRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Http request rule o k response has a 4xx status code
func (o *GetHTTPRequestRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Http request rule o k response has a 5xx status code
func (o *GetHTTPRequestRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Http request rule o k response a status code equal to that given
func (o *GetHTTPRequestRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Http request rule o k response
func (o *GetHTTPRequestRuleOK) Code() int {
	return 200
}

func (o *GetHTTPRequestRuleOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_request_rules/{index}][%d] getHttpRequestRuleOK  %+v", 200, o.Payload)
}

func (o *GetHTTPRequestRuleOK) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_request_rules/{index}][%d] getHttpRequestRuleOK  %+v", 200, o.Payload)
}

func (o *GetHTTPRequestRuleOK) GetPayload() *GetHTTPRequestRuleOKBody {
	return o.Payload
}

func (o *GetHTTPRequestRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetHTTPRequestRuleOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHTTPRequestRuleNotFound creates a GetHTTPRequestRuleNotFound with default headers values
func NewGetHTTPRequestRuleNotFound() *GetHTTPRequestRuleNotFound {
	return &GetHTTPRequestRuleNotFound{}
}

/*
GetHTTPRequestRuleNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetHTTPRequestRuleNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get Http request rule not found response has a 2xx status code
func (o *GetHTTPRequestRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Http request rule not found response has a 3xx status code
func (o *GetHTTPRequestRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Http request rule not found response has a 4xx status code
func (o *GetHTTPRequestRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Http request rule not found response has a 5xx status code
func (o *GetHTTPRequestRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get Http request rule not found response a status code equal to that given
func (o *GetHTTPRequestRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get Http request rule not found response
func (o *GetHTTPRequestRuleNotFound) Code() int {
	return 404
}

func (o *GetHTTPRequestRuleNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_request_rules/{index}][%d] getHttpRequestRuleNotFound  %+v", 404, o.Payload)
}

func (o *GetHTTPRequestRuleNotFound) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_request_rules/{index}][%d] getHttpRequestRuleNotFound  %+v", 404, o.Payload)
}

func (o *GetHTTPRequestRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHTTPRequestRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetHTTPRequestRuleDefault creates a GetHTTPRequestRuleDefault with default headers values
func NewGetHTTPRequestRuleDefault(code int) *GetHTTPRequestRuleDefault {
	return &GetHTTPRequestRuleDefault{
		_statusCode: code,
	}
}

/*
GetHTTPRequestRuleDefault describes a response with status code -1, with default header values.

General Error
*/
type GetHTTPRequestRuleDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get HTTP request rule default response has a 2xx status code
func (o *GetHTTPRequestRuleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get HTTP request rule default response has a 3xx status code
func (o *GetHTTPRequestRuleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get HTTP request rule default response has a 4xx status code
func (o *GetHTTPRequestRuleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get HTTP request rule default response has a 5xx status code
func (o *GetHTTPRequestRuleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get HTTP request rule default response a status code equal to that given
func (o *GetHTTPRequestRuleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get HTTP request rule default response
func (o *GetHTTPRequestRuleDefault) Code() int {
	return o._statusCode
}

func (o *GetHTTPRequestRuleDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_request_rules/{index}][%d] getHTTPRequestRule default  %+v", o._statusCode, o.Payload)
}

func (o *GetHTTPRequestRuleDefault) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_request_rules/{index}][%d] getHTTPRequestRule default  %+v", o._statusCode, o.Payload)
}

func (o *GetHTTPRequestRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHTTPRequestRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetHTTPRequestRuleOKBody get HTTP request rule o k body
swagger:model GetHTTPRequestRuleOKBody
*/
type GetHTTPRequestRuleOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.HTTPRequestRule `json:"data,omitempty"`
}

// Validate validates this get HTTP request rule o k body
func (o *GetHTTPRequestRuleOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHTTPRequestRuleOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getHttpRequestRuleOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getHttpRequestRuleOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get HTTP request rule o k body based on the context it is used
func (o *GetHTTPRequestRuleOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHTTPRequestRuleOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getHttpRequestRuleOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getHttpRequestRuleOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetHTTPRequestRuleOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHTTPRequestRuleOKBody) UnmarshalBinary(b []byte) error {
	var res GetHTTPRequestRuleOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
