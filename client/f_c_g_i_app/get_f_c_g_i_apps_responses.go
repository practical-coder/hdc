// Code generated by go-swagger; DO NOT EDIT.

package f_c_g_i_app

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

// GetFCGIAppsReader is a Reader for the GetFCGIApps structure.
type GetFCGIAppsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFCGIAppsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFCGIAppsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetFCGIAppsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFCGIAppsOK creates a GetFCGIAppsOK with default headers values
func NewGetFCGIAppsOK() *GetFCGIAppsOK {
	return &GetFCGIAppsOK{}
}

/*
GetFCGIAppsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetFCGIAppsOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetFCGIAppsOKBody
}

// IsSuccess returns true when this get f c g i apps o k response has a 2xx status code
func (o *GetFCGIAppsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get f c g i apps o k response has a 3xx status code
func (o *GetFCGIAppsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get f c g i apps o k response has a 4xx status code
func (o *GetFCGIAppsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get f c g i apps o k response has a 5xx status code
func (o *GetFCGIAppsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get f c g i apps o k response a status code equal to that given
func (o *GetFCGIAppsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetFCGIAppsOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/fcgi_apps][%d] getFCGIAppsOK  %+v", 200, o.Payload)
}

func (o *GetFCGIAppsOK) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/fcgi_apps][%d] getFCGIAppsOK  %+v", 200, o.Payload)
}

func (o *GetFCGIAppsOK) GetPayload() *GetFCGIAppsOKBody {
	return o.Payload
}

func (o *GetFCGIAppsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetFCGIAppsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFCGIAppsDefault creates a GetFCGIAppsDefault with default headers values
func NewGetFCGIAppsDefault(code int) *GetFCGIAppsDefault {
	return &GetFCGIAppsDefault{
		_statusCode: code,
	}
}

/*
GetFCGIAppsDefault describes a response with status code -1, with default header values.

General Error
*/
type GetFCGIAppsDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get f c g i apps default response
func (o *GetFCGIAppsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get f c g i apps default response has a 2xx status code
func (o *GetFCGIAppsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get f c g i apps default response has a 3xx status code
func (o *GetFCGIAppsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get f c g i apps default response has a 4xx status code
func (o *GetFCGIAppsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get f c g i apps default response has a 5xx status code
func (o *GetFCGIAppsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get f c g i apps default response a status code equal to that given
func (o *GetFCGIAppsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetFCGIAppsDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/fcgi_apps][%d] getFCGIApps default  %+v", o._statusCode, o.Payload)
}

func (o *GetFCGIAppsDefault) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/fcgi_apps][%d] getFCGIApps default  %+v", o._statusCode, o.Payload)
}

func (o *GetFCGIAppsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFCGIAppsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetFCGIAppsOKBody get f c g i apps o k body
swagger:model GetFCGIAppsOKBody
*/
type GetFCGIAppsOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.FcgiApps `json:"data"`
}

// Validate validates this get f c g i apps o k body
func (o *GetFCGIAppsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFCGIAppsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getFCGIAppsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getFCGIAppsOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getFCGIAppsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get f c g i apps o k body based on the context it is used
func (o *GetFCGIAppsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFCGIAppsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getFCGIAppsOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getFCGIAppsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetFCGIAppsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFCGIAppsOKBody) UnmarshalBinary(b []byte) error {
	var res GetFCGIAppsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}