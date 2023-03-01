// Code generated by go-swagger; DO NOT EDIT.

package http_errors

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

// GetHTTPErrorsSectionsReader is a Reader for the GetHTTPErrorsSections structure.
type GetHTTPErrorsSectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHTTPErrorsSectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHTTPErrorsSectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetHTTPErrorsSectionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHTTPErrorsSectionsOK creates a GetHTTPErrorsSectionsOK with default headers values
func NewGetHTTPErrorsSectionsOK() *GetHTTPErrorsSectionsOK {
	return &GetHTTPErrorsSectionsOK{}
}

/*
GetHTTPErrorsSectionsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetHTTPErrorsSectionsOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetHTTPErrorsSectionsOKBody
}

// IsSuccess returns true when this get Http errors sections o k response has a 2xx status code
func (o *GetHTTPErrorsSectionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Http errors sections o k response has a 3xx status code
func (o *GetHTTPErrorsSectionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Http errors sections o k response has a 4xx status code
func (o *GetHTTPErrorsSectionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Http errors sections o k response has a 5xx status code
func (o *GetHTTPErrorsSectionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Http errors sections o k response a status code equal to that given
func (o *GetHTTPErrorsSectionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetHTTPErrorsSectionsOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_errors_sections][%d] getHttpErrorsSectionsOK  %+v", 200, o.Payload)
}

func (o *GetHTTPErrorsSectionsOK) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_errors_sections][%d] getHttpErrorsSectionsOK  %+v", 200, o.Payload)
}

func (o *GetHTTPErrorsSectionsOK) GetPayload() *GetHTTPErrorsSectionsOKBody {
	return o.Payload
}

func (o *GetHTTPErrorsSectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetHTTPErrorsSectionsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHTTPErrorsSectionsDefault creates a GetHTTPErrorsSectionsDefault with default headers values
func NewGetHTTPErrorsSectionsDefault(code int) *GetHTTPErrorsSectionsDefault {
	return &GetHTTPErrorsSectionsDefault{
		_statusCode: code,
	}
}

/*
GetHTTPErrorsSectionsDefault describes a response with status code -1, with default header values.

General Error
*/
type GetHTTPErrorsSectionsDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get HTTP errors sections default response
func (o *GetHTTPErrorsSectionsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get HTTP errors sections default response has a 2xx status code
func (o *GetHTTPErrorsSectionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get HTTP errors sections default response has a 3xx status code
func (o *GetHTTPErrorsSectionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get HTTP errors sections default response has a 4xx status code
func (o *GetHTTPErrorsSectionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get HTTP errors sections default response has a 5xx status code
func (o *GetHTTPErrorsSectionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get HTTP errors sections default response a status code equal to that given
func (o *GetHTTPErrorsSectionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetHTTPErrorsSectionsDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_errors_sections][%d] getHTTPErrorsSections default  %+v", o._statusCode, o.Payload)
}

func (o *GetHTTPErrorsSectionsDefault) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_errors_sections][%d] getHTTPErrorsSections default  %+v", o._statusCode, o.Payload)
}

func (o *GetHTTPErrorsSectionsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHTTPErrorsSectionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetHTTPErrorsSectionsOKBody get HTTP errors sections o k body
swagger:model GetHTTPErrorsSectionsOKBody
*/
type GetHTTPErrorsSectionsOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.HTTPErrorsSections `json:"data"`
}

// Validate validates this get HTTP errors sections o k body
func (o *GetHTTPErrorsSectionsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHTTPErrorsSectionsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getHttpErrorsSectionsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getHttpErrorsSectionsOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getHttpErrorsSectionsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get HTTP errors sections o k body based on the context it is used
func (o *GetHTTPErrorsSectionsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHTTPErrorsSectionsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getHttpErrorsSectionsOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getHttpErrorsSectionsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetHTTPErrorsSectionsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHTTPErrorsSectionsOKBody) UnmarshalBinary(b []byte) error {
	var res GetHTTPErrorsSectionsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}