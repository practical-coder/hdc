// Code generated by go-swagger; DO NOT EDIT.

package http_error_rule

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

// GetHTTPErrorRuleReader is a Reader for the GetHTTPErrorRule structure.
type GetHTTPErrorRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHTTPErrorRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHTTPErrorRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetHTTPErrorRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetHTTPErrorRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHTTPErrorRuleOK creates a GetHTTPErrorRuleOK with default headers values
func NewGetHTTPErrorRuleOK() *GetHTTPErrorRuleOK {
	return &GetHTTPErrorRuleOK{}
}

/*
GetHTTPErrorRuleOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetHTTPErrorRuleOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetHTTPErrorRuleOKBody
}

// IsSuccess returns true when this get Http error rule o k response has a 2xx status code
func (o *GetHTTPErrorRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Http error rule o k response has a 3xx status code
func (o *GetHTTPErrorRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Http error rule o k response has a 4xx status code
func (o *GetHTTPErrorRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Http error rule o k response has a 5xx status code
func (o *GetHTTPErrorRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Http error rule o k response a status code equal to that given
func (o *GetHTTPErrorRuleOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetHTTPErrorRuleOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_error_rules/{index}][%d] getHttpErrorRuleOK  %+v", 200, o.Payload)
}

func (o *GetHTTPErrorRuleOK) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_error_rules/{index}][%d] getHttpErrorRuleOK  %+v", 200, o.Payload)
}

func (o *GetHTTPErrorRuleOK) GetPayload() *GetHTTPErrorRuleOKBody {
	return o.Payload
}

func (o *GetHTTPErrorRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetHTTPErrorRuleOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHTTPErrorRuleNotFound creates a GetHTTPErrorRuleNotFound with default headers values
func NewGetHTTPErrorRuleNotFound() *GetHTTPErrorRuleNotFound {
	return &GetHTTPErrorRuleNotFound{}
}

/*
GetHTTPErrorRuleNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetHTTPErrorRuleNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get Http error rule not found response has a 2xx status code
func (o *GetHTTPErrorRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Http error rule not found response has a 3xx status code
func (o *GetHTTPErrorRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Http error rule not found response has a 4xx status code
func (o *GetHTTPErrorRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Http error rule not found response has a 5xx status code
func (o *GetHTTPErrorRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get Http error rule not found response a status code equal to that given
func (o *GetHTTPErrorRuleNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetHTTPErrorRuleNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_error_rules/{index}][%d] getHttpErrorRuleNotFound  %+v", 404, o.Payload)
}

func (o *GetHTTPErrorRuleNotFound) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_error_rules/{index}][%d] getHttpErrorRuleNotFound  %+v", 404, o.Payload)
}

func (o *GetHTTPErrorRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHTTPErrorRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetHTTPErrorRuleDefault creates a GetHTTPErrorRuleDefault with default headers values
func NewGetHTTPErrorRuleDefault(code int) *GetHTTPErrorRuleDefault {
	return &GetHTTPErrorRuleDefault{
		_statusCode: code,
	}
}

/*
GetHTTPErrorRuleDefault describes a response with status code -1, with default header values.

General Error
*/
type GetHTTPErrorRuleDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get HTTP error rule default response
func (o *GetHTTPErrorRuleDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get HTTP error rule default response has a 2xx status code
func (o *GetHTTPErrorRuleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get HTTP error rule default response has a 3xx status code
func (o *GetHTTPErrorRuleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get HTTP error rule default response has a 4xx status code
func (o *GetHTTPErrorRuleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get HTTP error rule default response has a 5xx status code
func (o *GetHTTPErrorRuleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get HTTP error rule default response a status code equal to that given
func (o *GetHTTPErrorRuleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetHTTPErrorRuleDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_error_rules/{index}][%d] getHTTPErrorRule default  %+v", o._statusCode, o.Payload)
}

func (o *GetHTTPErrorRuleDefault) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_error_rules/{index}][%d] getHTTPErrorRule default  %+v", o._statusCode, o.Payload)
}

func (o *GetHTTPErrorRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHTTPErrorRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetHTTPErrorRuleOKBody get HTTP error rule o k body
swagger:model GetHTTPErrorRuleOKBody
*/
type GetHTTPErrorRuleOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.HTTPErrorRule `json:"data,omitempty"`
}

// Validate validates this get HTTP error rule o k body
func (o *GetHTTPErrorRuleOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHTTPErrorRuleOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getHttpErrorRuleOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getHttpErrorRuleOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get HTTP error rule o k body based on the context it is used
func (o *GetHTTPErrorRuleOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHTTPErrorRuleOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getHttpErrorRuleOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getHttpErrorRuleOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetHTTPErrorRuleOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHTTPErrorRuleOKBody) UnmarshalBinary(b []byte) error {
	var res GetHTTPErrorRuleOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}