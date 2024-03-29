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
	"github.com/go-openapi/validate"

	"github.com/haproxytech/client-native/v5/models"
)

// GetServerSwitchingRulesReader is a Reader for the GetServerSwitchingRules structure.
type GetServerSwitchingRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServerSwitchingRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServerSwitchingRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetServerSwitchingRulesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServerSwitchingRulesOK creates a GetServerSwitchingRulesOK with default headers values
func NewGetServerSwitchingRulesOK() *GetServerSwitchingRulesOK {
	return &GetServerSwitchingRulesOK{}
}

/*
GetServerSwitchingRulesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetServerSwitchingRulesOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetServerSwitchingRulesOKBody
}

// IsSuccess returns true when this get server switching rules o k response has a 2xx status code
func (o *GetServerSwitchingRulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get server switching rules o k response has a 3xx status code
func (o *GetServerSwitchingRulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get server switching rules o k response has a 4xx status code
func (o *GetServerSwitchingRulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get server switching rules o k response has a 5xx status code
func (o *GetServerSwitchingRulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get server switching rules o k response a status code equal to that given
func (o *GetServerSwitchingRulesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get server switching rules o k response
func (o *GetServerSwitchingRulesOK) Code() int {
	return 200
}

func (o *GetServerSwitchingRulesOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/server_switching_rules][%d] getServerSwitchingRulesOK  %+v", 200, o.Payload)
}

func (o *GetServerSwitchingRulesOK) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/server_switching_rules][%d] getServerSwitchingRulesOK  %+v", 200, o.Payload)
}

func (o *GetServerSwitchingRulesOK) GetPayload() *GetServerSwitchingRulesOKBody {
	return o.Payload
}

func (o *GetServerSwitchingRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetServerSwitchingRulesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServerSwitchingRulesDefault creates a GetServerSwitchingRulesDefault with default headers values
func NewGetServerSwitchingRulesDefault(code int) *GetServerSwitchingRulesDefault {
	return &GetServerSwitchingRulesDefault{
		_statusCode: code,
	}
}

/*
GetServerSwitchingRulesDefault describes a response with status code -1, with default header values.

General Error
*/
type GetServerSwitchingRulesDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get server switching rules default response has a 2xx status code
func (o *GetServerSwitchingRulesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get server switching rules default response has a 3xx status code
func (o *GetServerSwitchingRulesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get server switching rules default response has a 4xx status code
func (o *GetServerSwitchingRulesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get server switching rules default response has a 5xx status code
func (o *GetServerSwitchingRulesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get server switching rules default response a status code equal to that given
func (o *GetServerSwitchingRulesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get server switching rules default response
func (o *GetServerSwitchingRulesDefault) Code() int {
	return o._statusCode
}

func (o *GetServerSwitchingRulesDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/server_switching_rules][%d] getServerSwitchingRules default  %+v", o._statusCode, o.Payload)
}

func (o *GetServerSwitchingRulesDefault) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/server_switching_rules][%d] getServerSwitchingRules default  %+v", o._statusCode, o.Payload)
}

func (o *GetServerSwitchingRulesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServerSwitchingRulesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetServerSwitchingRulesOKBody get server switching rules o k body
swagger:model GetServerSwitchingRulesOKBody
*/
type GetServerSwitchingRulesOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.ServerSwitchingRules `json:"data"`
}

// Validate validates this get server switching rules o k body
func (o *GetServerSwitchingRulesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServerSwitchingRulesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getServerSwitchingRulesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getServerSwitchingRulesOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getServerSwitchingRulesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get server switching rules o k body based on the context it is used
func (o *GetServerSwitchingRulesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServerSwitchingRulesOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getServerSwitchingRulesOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getServerSwitchingRulesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetServerSwitchingRulesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServerSwitchingRulesOKBody) UnmarshalBinary(b []byte) error {
	var res GetServerSwitchingRulesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
