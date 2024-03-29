// Code generated by go-swagger; DO NOT EDIT.

package tcp_request_rule

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

// GetTCPRequestRuleReader is a Reader for the GetTCPRequestRule structure.
type GetTCPRequestRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTCPRequestRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTCPRequestRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTCPRequestRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetTCPRequestRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTCPRequestRuleOK creates a GetTCPRequestRuleOK with default headers values
func NewGetTCPRequestRuleOK() *GetTCPRequestRuleOK {
	return &GetTCPRequestRuleOK{}
}

/*
GetTCPRequestRuleOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetTCPRequestRuleOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetTCPRequestRuleOKBody
}

// IsSuccess returns true when this get Tcp request rule o k response has a 2xx status code
func (o *GetTCPRequestRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Tcp request rule o k response has a 3xx status code
func (o *GetTCPRequestRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Tcp request rule o k response has a 4xx status code
func (o *GetTCPRequestRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Tcp request rule o k response has a 5xx status code
func (o *GetTCPRequestRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Tcp request rule o k response a status code equal to that given
func (o *GetTCPRequestRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Tcp request rule o k response
func (o *GetTCPRequestRuleOK) Code() int {
	return 200
}

func (o *GetTCPRequestRuleOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/tcp_request_rules/{index}][%d] getTcpRequestRuleOK  %+v", 200, o.Payload)
}

func (o *GetTCPRequestRuleOK) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/tcp_request_rules/{index}][%d] getTcpRequestRuleOK  %+v", 200, o.Payload)
}

func (o *GetTCPRequestRuleOK) GetPayload() *GetTCPRequestRuleOKBody {
	return o.Payload
}

func (o *GetTCPRequestRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetTCPRequestRuleOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTCPRequestRuleNotFound creates a GetTCPRequestRuleNotFound with default headers values
func NewGetTCPRequestRuleNotFound() *GetTCPRequestRuleNotFound {
	return &GetTCPRequestRuleNotFound{}
}

/*
GetTCPRequestRuleNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetTCPRequestRuleNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get Tcp request rule not found response has a 2xx status code
func (o *GetTCPRequestRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Tcp request rule not found response has a 3xx status code
func (o *GetTCPRequestRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Tcp request rule not found response has a 4xx status code
func (o *GetTCPRequestRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Tcp request rule not found response has a 5xx status code
func (o *GetTCPRequestRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get Tcp request rule not found response a status code equal to that given
func (o *GetTCPRequestRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get Tcp request rule not found response
func (o *GetTCPRequestRuleNotFound) Code() int {
	return 404
}

func (o *GetTCPRequestRuleNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/tcp_request_rules/{index}][%d] getTcpRequestRuleNotFound  %+v", 404, o.Payload)
}

func (o *GetTCPRequestRuleNotFound) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/tcp_request_rules/{index}][%d] getTcpRequestRuleNotFound  %+v", 404, o.Payload)
}

func (o *GetTCPRequestRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTCPRequestRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetTCPRequestRuleDefault creates a GetTCPRequestRuleDefault with default headers values
func NewGetTCPRequestRuleDefault(code int) *GetTCPRequestRuleDefault {
	return &GetTCPRequestRuleDefault{
		_statusCode: code,
	}
}

/*
GetTCPRequestRuleDefault describes a response with status code -1, with default header values.

General Error
*/
type GetTCPRequestRuleDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get TCP request rule default response has a 2xx status code
func (o *GetTCPRequestRuleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get TCP request rule default response has a 3xx status code
func (o *GetTCPRequestRuleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get TCP request rule default response has a 4xx status code
func (o *GetTCPRequestRuleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get TCP request rule default response has a 5xx status code
func (o *GetTCPRequestRuleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get TCP request rule default response a status code equal to that given
func (o *GetTCPRequestRuleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get TCP request rule default response
func (o *GetTCPRequestRuleDefault) Code() int {
	return o._statusCode
}

func (o *GetTCPRequestRuleDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/tcp_request_rules/{index}][%d] getTCPRequestRule default  %+v", o._statusCode, o.Payload)
}

func (o *GetTCPRequestRuleDefault) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/tcp_request_rules/{index}][%d] getTCPRequestRule default  %+v", o._statusCode, o.Payload)
}

func (o *GetTCPRequestRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTCPRequestRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetTCPRequestRuleOKBody get TCP request rule o k body
swagger:model GetTCPRequestRuleOKBody
*/
type GetTCPRequestRuleOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.TCPRequestRule `json:"data,omitempty"`
}

// Validate validates this get TCP request rule o k body
func (o *GetTCPRequestRuleOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTCPRequestRuleOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getTcpRequestRuleOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getTcpRequestRuleOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get TCP request rule o k body based on the context it is used
func (o *GetTCPRequestRuleOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTCPRequestRuleOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getTcpRequestRuleOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getTcpRequestRuleOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTCPRequestRuleOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTCPRequestRuleOKBody) UnmarshalBinary(b []byte) error {
	var res GetTCPRequestRuleOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
