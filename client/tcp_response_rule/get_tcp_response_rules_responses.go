// Code generated by go-swagger; DO NOT EDIT.

package tcp_response_rule

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

	"client-native/models"
)

// GetTCPResponseRulesReader is a Reader for the GetTCPResponseRules structure.
type GetTCPResponseRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTCPResponseRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTCPResponseRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetTCPResponseRulesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTCPResponseRulesOK creates a GetTCPResponseRulesOK with default headers values
func NewGetTCPResponseRulesOK() *GetTCPResponseRulesOK {
	return &GetTCPResponseRulesOK{}
}

/* GetTCPResponseRulesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetTCPResponseRulesOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetTCPResponseRulesOKBody
}

func (o *GetTCPResponseRulesOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/tcp_response_rules][%d] getTcpResponseRulesOK  %+v", 200, o.Payload)
}
func (o *GetTCPResponseRulesOK) GetPayload() *GetTCPResponseRulesOKBody {
	return o.Payload
}

func (o *GetTCPResponseRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetTCPResponseRulesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTCPResponseRulesDefault creates a GetTCPResponseRulesDefault with default headers values
func NewGetTCPResponseRulesDefault(code int) *GetTCPResponseRulesDefault {
	return &GetTCPResponseRulesDefault{
		_statusCode: code,
	}
}

/* GetTCPResponseRulesDefault describes a response with status code -1, with default header values.

General Error
*/
type GetTCPResponseRulesDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get TCP response rules default response
func (o *GetTCPResponseRulesDefault) Code() int {
	return o._statusCode
}

func (o *GetTCPResponseRulesDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/tcp_response_rules][%d] getTCPResponseRules default  %+v", o._statusCode, o.Payload)
}
func (o *GetTCPResponseRulesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTCPResponseRulesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetTCPResponseRulesOKBody get TCP response rules o k body
swagger:model GetTCPResponseRulesOKBody
*/
type GetTCPResponseRulesOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.TCPResponseRules `json:"data"`
}

// Validate validates this get TCP response rules o k body
func (o *GetTCPResponseRulesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTCPResponseRulesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getTcpResponseRulesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getTcpResponseRulesOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getTcpResponseRulesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get TCP response rules o k body based on the context it is used
func (o *GetTCPResponseRulesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTCPResponseRulesOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getTcpResponseRulesOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getTcpResponseRulesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTCPResponseRulesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTCPResponseRulesOKBody) UnmarshalBinary(b []byte) error {
	var res GetTCPResponseRulesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
