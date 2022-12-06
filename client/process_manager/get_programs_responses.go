// Code generated by go-swagger; DO NOT EDIT.

package process_manager

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

// GetProgramsReader is a Reader for the GetPrograms structure.
type GetProgramsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProgramsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProgramsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetProgramsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProgramsOK creates a GetProgramsOK with default headers values
func NewGetProgramsOK() *GetProgramsOK {
	return &GetProgramsOK{}
}

/*
GetProgramsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetProgramsOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetProgramsOKBody
}

// IsSuccess returns true when this get programs o k response has a 2xx status code
func (o *GetProgramsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get programs o k response has a 3xx status code
func (o *GetProgramsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get programs o k response has a 4xx status code
func (o *GetProgramsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get programs o k response has a 5xx status code
func (o *GetProgramsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get programs o k response a status code equal to that given
func (o *GetProgramsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetProgramsOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/programs][%d] getProgramsOK  %+v", 200, o.Payload)
}

func (o *GetProgramsOK) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/programs][%d] getProgramsOK  %+v", 200, o.Payload)
}

func (o *GetProgramsOK) GetPayload() *GetProgramsOKBody {
	return o.Payload
}

func (o *GetProgramsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetProgramsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProgramsDefault creates a GetProgramsDefault with default headers values
func NewGetProgramsDefault(code int) *GetProgramsDefault {
	return &GetProgramsDefault{
		_statusCode: code,
	}
}

/*
GetProgramsDefault describes a response with status code -1, with default header values.

General Error
*/
type GetProgramsDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get programs default response
func (o *GetProgramsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get programs default response has a 2xx status code
func (o *GetProgramsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get programs default response has a 3xx status code
func (o *GetProgramsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get programs default response has a 4xx status code
func (o *GetProgramsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get programs default response has a 5xx status code
func (o *GetProgramsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get programs default response a status code equal to that given
func (o *GetProgramsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetProgramsDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/programs][%d] getPrograms default  %+v", o._statusCode, o.Payload)
}

func (o *GetProgramsDefault) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/programs][%d] getPrograms default  %+v", o._statusCode, o.Payload)
}

func (o *GetProgramsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetProgramsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetProgramsOKBody get programs o k body
swagger:model GetProgramsOKBody
*/
type GetProgramsOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.Programs `json:"data"`
}

// Validate validates this get programs o k body
func (o *GetProgramsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetProgramsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getProgramsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getProgramsOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getProgramsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get programs o k body based on the context it is used
func (o *GetProgramsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetProgramsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getProgramsOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getProgramsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetProgramsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetProgramsOKBody) UnmarshalBinary(b []byte) error {
	var res GetProgramsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
