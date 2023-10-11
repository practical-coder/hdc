// Code generated by go-swagger; DO NOT EDIT.

package server_switching_rule

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

// GetServerSwitchingRuleReader is a Reader for the GetServerSwitchingRule structure.
type GetServerSwitchingRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServerSwitchingRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServerSwitchingRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetServerSwitchingRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetServerSwitchingRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServerSwitchingRuleOK creates a GetServerSwitchingRuleOK with default headers values
func NewGetServerSwitchingRuleOK() *GetServerSwitchingRuleOK {
	return &GetServerSwitchingRuleOK{}
}

/*
GetServerSwitchingRuleOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetServerSwitchingRuleOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetServerSwitchingRuleOKBody
}

// IsSuccess returns true when this get server switching rule o k response has a 2xx status code
func (o *GetServerSwitchingRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get server switching rule o k response has a 3xx status code
func (o *GetServerSwitchingRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get server switching rule o k response has a 4xx status code
func (o *GetServerSwitchingRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get server switching rule o k response has a 5xx status code
func (o *GetServerSwitchingRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get server switching rule o k response a status code equal to that given
func (o *GetServerSwitchingRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get server switching rule o k response
func (o *GetServerSwitchingRuleOK) Code() int {
	return 200
}

func (o *GetServerSwitchingRuleOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/server_switching_rules/{index}][%d] getServerSwitchingRuleOK  %+v", 200, o.Payload)
}

func (o *GetServerSwitchingRuleOK) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/server_switching_rules/{index}][%d] getServerSwitchingRuleOK  %+v", 200, o.Payload)
}

func (o *GetServerSwitchingRuleOK) GetPayload() *GetServerSwitchingRuleOKBody {
	return o.Payload
}

func (o *GetServerSwitchingRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetServerSwitchingRuleOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServerSwitchingRuleNotFound creates a GetServerSwitchingRuleNotFound with default headers values
func NewGetServerSwitchingRuleNotFound() *GetServerSwitchingRuleNotFound {
	return &GetServerSwitchingRuleNotFound{}
}

/*
GetServerSwitchingRuleNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetServerSwitchingRuleNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get server switching rule not found response has a 2xx status code
func (o *GetServerSwitchingRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get server switching rule not found response has a 3xx status code
func (o *GetServerSwitchingRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get server switching rule not found response has a 4xx status code
func (o *GetServerSwitchingRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get server switching rule not found response has a 5xx status code
func (o *GetServerSwitchingRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get server switching rule not found response a status code equal to that given
func (o *GetServerSwitchingRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get server switching rule not found response
func (o *GetServerSwitchingRuleNotFound) Code() int {
	return 404
}

func (o *GetServerSwitchingRuleNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/server_switching_rules/{index}][%d] getServerSwitchingRuleNotFound  %+v", 404, o.Payload)
}

func (o *GetServerSwitchingRuleNotFound) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/server_switching_rules/{index}][%d] getServerSwitchingRuleNotFound  %+v", 404, o.Payload)
}

func (o *GetServerSwitchingRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServerSwitchingRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetServerSwitchingRuleDefault creates a GetServerSwitchingRuleDefault with default headers values
func NewGetServerSwitchingRuleDefault(code int) *GetServerSwitchingRuleDefault {
	return &GetServerSwitchingRuleDefault{
		_statusCode: code,
	}
}

/*
GetServerSwitchingRuleDefault describes a response with status code -1, with default header values.

General Error
*/
type GetServerSwitchingRuleDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get server switching rule default response has a 2xx status code
func (o *GetServerSwitchingRuleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get server switching rule default response has a 3xx status code
func (o *GetServerSwitchingRuleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get server switching rule default response has a 4xx status code
func (o *GetServerSwitchingRuleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get server switching rule default response has a 5xx status code
func (o *GetServerSwitchingRuleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get server switching rule default response a status code equal to that given
func (o *GetServerSwitchingRuleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get server switching rule default response
func (o *GetServerSwitchingRuleDefault) Code() int {
	return o._statusCode
}

func (o *GetServerSwitchingRuleDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/server_switching_rules/{index}][%d] getServerSwitchingRule default  %+v", o._statusCode, o.Payload)
}

func (o *GetServerSwitchingRuleDefault) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/server_switching_rules/{index}][%d] getServerSwitchingRule default  %+v", o._statusCode, o.Payload)
}

func (o *GetServerSwitchingRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServerSwitchingRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetServerSwitchingRuleOKBody get server switching rule o k body
swagger:model GetServerSwitchingRuleOKBody
*/
type GetServerSwitchingRuleOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.ServerSwitchingRule `json:"data,omitempty"`
}

// Validate validates this get server switching rule o k body
func (o *GetServerSwitchingRuleOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServerSwitchingRuleOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getServerSwitchingRuleOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getServerSwitchingRuleOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get server switching rule o k body based on the context it is used
func (o *GetServerSwitchingRuleOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServerSwitchingRuleOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getServerSwitchingRuleOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getServerSwitchingRuleOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetServerSwitchingRuleOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServerSwitchingRuleOKBody) UnmarshalBinary(b []byte) error {
	var res GetServerSwitchingRuleOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
