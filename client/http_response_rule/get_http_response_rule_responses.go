// Code generated by go-swagger; DO NOT EDIT.

package http_response_rule

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

// GetHTTPResponseRuleReader is a Reader for the GetHTTPResponseRule structure.
type GetHTTPResponseRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHTTPResponseRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHTTPResponseRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetHTTPResponseRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetHTTPResponseRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHTTPResponseRuleOK creates a GetHTTPResponseRuleOK with default headers values
func NewGetHTTPResponseRuleOK() *GetHTTPResponseRuleOK {
	return &GetHTTPResponseRuleOK{}
}

/*
GetHTTPResponseRuleOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetHTTPResponseRuleOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetHTTPResponseRuleOKBody
}

// IsSuccess returns true when this get Http response rule o k response has a 2xx status code
func (o *GetHTTPResponseRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Http response rule o k response has a 3xx status code
func (o *GetHTTPResponseRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Http response rule o k response has a 4xx status code
func (o *GetHTTPResponseRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Http response rule o k response has a 5xx status code
func (o *GetHTTPResponseRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Http response rule o k response a status code equal to that given
func (o *GetHTTPResponseRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Http response rule o k response
func (o *GetHTTPResponseRuleOK) Code() int {
	return 200
}

func (o *GetHTTPResponseRuleOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_response_rules/{index}][%d] getHttpResponseRuleOK  %+v", 200, o.Payload)
}

func (o *GetHTTPResponseRuleOK) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_response_rules/{index}][%d] getHttpResponseRuleOK  %+v", 200, o.Payload)
}

func (o *GetHTTPResponseRuleOK) GetPayload() *GetHTTPResponseRuleOKBody {
	return o.Payload
}

func (o *GetHTTPResponseRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetHTTPResponseRuleOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHTTPResponseRuleNotFound creates a GetHTTPResponseRuleNotFound with default headers values
func NewGetHTTPResponseRuleNotFound() *GetHTTPResponseRuleNotFound {
	return &GetHTTPResponseRuleNotFound{}
}

/*
GetHTTPResponseRuleNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetHTTPResponseRuleNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get Http response rule not found response has a 2xx status code
func (o *GetHTTPResponseRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Http response rule not found response has a 3xx status code
func (o *GetHTTPResponseRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Http response rule not found response has a 4xx status code
func (o *GetHTTPResponseRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Http response rule not found response has a 5xx status code
func (o *GetHTTPResponseRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get Http response rule not found response a status code equal to that given
func (o *GetHTTPResponseRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get Http response rule not found response
func (o *GetHTTPResponseRuleNotFound) Code() int {
	return 404
}

func (o *GetHTTPResponseRuleNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_response_rules/{index}][%d] getHttpResponseRuleNotFound  %+v", 404, o.Payload)
}

func (o *GetHTTPResponseRuleNotFound) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_response_rules/{index}][%d] getHttpResponseRuleNotFound  %+v", 404, o.Payload)
}

func (o *GetHTTPResponseRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHTTPResponseRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetHTTPResponseRuleDefault creates a GetHTTPResponseRuleDefault with default headers values
func NewGetHTTPResponseRuleDefault(code int) *GetHTTPResponseRuleDefault {
	return &GetHTTPResponseRuleDefault{
		_statusCode: code,
	}
}

/*
GetHTTPResponseRuleDefault describes a response with status code -1, with default header values.

General Error
*/
type GetHTTPResponseRuleDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get HTTP response rule default response has a 2xx status code
func (o *GetHTTPResponseRuleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get HTTP response rule default response has a 3xx status code
func (o *GetHTTPResponseRuleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get HTTP response rule default response has a 4xx status code
func (o *GetHTTPResponseRuleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get HTTP response rule default response has a 5xx status code
func (o *GetHTTPResponseRuleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get HTTP response rule default response a status code equal to that given
func (o *GetHTTPResponseRuleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get HTTP response rule default response
func (o *GetHTTPResponseRuleDefault) Code() int {
	return o._statusCode
}

func (o *GetHTTPResponseRuleDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_response_rules/{index}][%d] getHTTPResponseRule default  %+v", o._statusCode, o.Payload)
}

func (o *GetHTTPResponseRuleDefault) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_response_rules/{index}][%d] getHTTPResponseRule default  %+v", o._statusCode, o.Payload)
}

func (o *GetHTTPResponseRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHTTPResponseRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetHTTPResponseRuleOKBody get HTTP response rule o k body
swagger:model GetHTTPResponseRuleOKBody
*/
type GetHTTPResponseRuleOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.HTTPResponseRule `json:"data,omitempty"`
}

// Validate validates this get HTTP response rule o k body
func (o *GetHTTPResponseRuleOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHTTPResponseRuleOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getHttpResponseRuleOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getHttpResponseRuleOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get HTTP response rule o k body based on the context it is used
func (o *GetHTTPResponseRuleOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHTTPResponseRuleOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getHttpResponseRuleOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getHttpResponseRuleOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetHTTPResponseRuleOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHTTPResponseRuleOKBody) UnmarshalBinary(b []byte) error {
	var res GetHTTPResponseRuleOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
