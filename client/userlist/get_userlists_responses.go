// Code generated by go-swagger; DO NOT EDIT.

package userlist

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

	"github.com/haproxytech/client-native/v4/models"
)

// GetUserlistsReader is a Reader for the GetUserlists structure.
type GetUserlistsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserlistsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserlistsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetUserlistsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUserlistsOK creates a GetUserlistsOK with default headers values
func NewGetUserlistsOK() *GetUserlistsOK {
	return &GetUserlistsOK{}
}

/* GetUserlistsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetUserlistsOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetUserlistsOKBody
}

func (o *GetUserlistsOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/userlists][%d] getUserlistsOK  %+v", 200, o.Payload)
}
func (o *GetUserlistsOK) GetPayload() *GetUserlistsOKBody {
	return o.Payload
}

func (o *GetUserlistsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetUserlistsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserlistsDefault creates a GetUserlistsDefault with default headers values
func NewGetUserlistsDefault(code int) *GetUserlistsDefault {
	return &GetUserlistsDefault{
		_statusCode: code,
	}
}

/* GetUserlistsDefault describes a response with status code -1, with default header values.

General Error
*/
type GetUserlistsDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get userlists default response
func (o *GetUserlistsDefault) Code() int {
	return o._statusCode
}

func (o *GetUserlistsDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/userlists][%d] getUserlists default  %+v", o._statusCode, o.Payload)
}
func (o *GetUserlistsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUserlistsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetUserlistsOKBody get userlists o k body
swagger:model GetUserlistsOKBody
*/
type GetUserlistsOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.Userlists `json:"data"`
}

// Validate validates this get userlists o k body
func (o *GetUserlistsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUserlistsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getUserlistsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getUserlistsOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getUserlistsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get userlists o k body based on the context it is used
func (o *GetUserlistsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUserlistsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getUserlistsOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getUserlistsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUserlistsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUserlistsOKBody) UnmarshalBinary(b []byte) error {
	var res GetUserlistsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}