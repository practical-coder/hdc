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

	"github.com/practical-coder/hdc/models"
)

// GetServerTemplateReader is a Reader for the GetServerTemplate structure.
type GetServerTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServerTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServerTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetServerTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetServerTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServerTemplateOK creates a GetServerTemplateOK with default headers values
func NewGetServerTemplateOK() *GetServerTemplateOK {
	return &GetServerTemplateOK{}
}

/* GetServerTemplateOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetServerTemplateOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion int64

	Payload *GetServerTemplateOKBody
}

func (o *GetServerTemplateOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/server_templates/{prefix}][%d] getServerTemplateOK  %+v", 200, o.Payload)
}
func (o *GetServerTemplateOK) GetPayload() *GetServerTemplateOKBody {
	return o.Payload
}

func (o *GetServerTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		valconfigurationVersion, err := swag.ConvertInt64(hdrConfigurationVersion)
		if err != nil {
			return errors.InvalidType("Configuration-Version", "header", "int64", hdrConfigurationVersion)
		}
		o.ConfigurationVersion = valconfigurationVersion
	}

	o.Payload = new(GetServerTemplateOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServerTemplateNotFound creates a GetServerTemplateNotFound with default headers values
func NewGetServerTemplateNotFound() *GetServerTemplateNotFound {
	return &GetServerTemplateNotFound{}
}

/* GetServerTemplateNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetServerTemplateNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetServerTemplateNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/server_templates/{prefix}][%d] getServerTemplateNotFound  %+v", 404, o.Payload)
}
func (o *GetServerTemplateNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServerTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetServerTemplateDefault creates a GetServerTemplateDefault with default headers values
func NewGetServerTemplateDefault(code int) *GetServerTemplateDefault {
	return &GetServerTemplateDefault{
		_statusCode: code,
	}
}

/* GetServerTemplateDefault describes a response with status code -1, with default header values.

General Error
*/
type GetServerTemplateDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get server template default response
func (o *GetServerTemplateDefault) Code() int {
	return o._statusCode
}

func (o *GetServerTemplateDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/server_templates/{prefix}][%d] getServerTemplate default  %+v", o._statusCode, o.Payload)
}
func (o *GetServerTemplateDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServerTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetServerTemplateOKBody get server template o k body
swagger:model GetServerTemplateOKBody
*/
type GetServerTemplateOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.ServerTemplate `json:"data,omitempty"`
}

// Validate validates this get server template o k body
func (o *GetServerTemplateOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServerTemplateOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getServerTemplateOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getServerTemplateOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get server template o k body based on the context it is used
func (o *GetServerTemplateOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServerTemplateOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getServerTemplateOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getServerTemplateOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetServerTemplateOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServerTemplateOKBody) UnmarshalBinary(b []byte) error {
	var res GetServerTemplateOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
