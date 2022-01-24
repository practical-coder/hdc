// Code generated by go-swagger; DO NOT EDIT.

package server_template

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

	"github.com/haproxytech/client-native/v2/models"
)

// GetServerTemplatesReader is a Reader for the GetServerTemplates structure.
type GetServerTemplatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServerTemplatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServerTemplatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetServerTemplatesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServerTemplatesOK creates a GetServerTemplatesOK with default headers values
func NewGetServerTemplatesOK() *GetServerTemplatesOK {
	return &GetServerTemplatesOK{}
}

/* GetServerTemplatesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetServerTemplatesOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion int64

	Payload *GetServerTemplatesOKBody
}

func (o *GetServerTemplatesOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/server_templates][%d] getServerTemplatesOK  %+v", 200, o.Payload)
}
func (o *GetServerTemplatesOK) GetPayload() *GetServerTemplatesOKBody {
	return o.Payload
}

func (o *GetServerTemplatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		valconfigurationVersion, err := swag.ConvertInt64(hdrConfigurationVersion)
		if err != nil {
			return errors.InvalidType("Configuration-Version", "header", "int64", hdrConfigurationVersion)
		}
		o.ConfigurationVersion = valconfigurationVersion
	}

	o.Payload = new(GetServerTemplatesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServerTemplatesDefault creates a GetServerTemplatesDefault with default headers values
func NewGetServerTemplatesDefault(code int) *GetServerTemplatesDefault {
	return &GetServerTemplatesDefault{
		_statusCode: code,
	}
}

/* GetServerTemplatesDefault describes a response with status code -1, with default header values.

General Error
*/
type GetServerTemplatesDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get server templates default response
func (o *GetServerTemplatesDefault) Code() int {
	return o._statusCode
}

func (o *GetServerTemplatesDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/server_templates][%d] getServerTemplates default  %+v", o._statusCode, o.Payload)
}
func (o *GetServerTemplatesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServerTemplatesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetServerTemplatesOKBody get server templates o k body
swagger:model GetServerTemplatesOKBody
*/
type GetServerTemplatesOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.ServerTemplates `json:"data"`
}

// Validate validates this get server templates o k body
func (o *GetServerTemplatesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServerTemplatesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getServerTemplatesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getServerTemplatesOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getServerTemplatesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get server templates o k body based on the context it is used
func (o *GetServerTemplatesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServerTemplatesOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getServerTemplatesOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getServerTemplatesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetServerTemplatesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServerTemplatesOKBody) UnmarshalBinary(b []byte) error {
	var res GetServerTemplatesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
