// Code generated by go-swagger; DO NOT EDIT.

package mailer_entry

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

// GetMailerEntriesReader is a Reader for the GetMailerEntries structure.
type GetMailerEntriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMailerEntriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMailerEntriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMailerEntriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMailerEntriesOK creates a GetMailerEntriesOK with default headers values
func NewGetMailerEntriesOK() *GetMailerEntriesOK {
	return &GetMailerEntriesOK{}
}

/*
GetMailerEntriesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetMailerEntriesOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetMailerEntriesOKBody
}

// IsSuccess returns true when this get mailer entries o k response has a 2xx status code
func (o *GetMailerEntriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get mailer entries o k response has a 3xx status code
func (o *GetMailerEntriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mailer entries o k response has a 4xx status code
func (o *GetMailerEntriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get mailer entries o k response has a 5xx status code
func (o *GetMailerEntriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get mailer entries o k response a status code equal to that given
func (o *GetMailerEntriesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get mailer entries o k response
func (o *GetMailerEntriesOK) Code() int {
	return 200
}

func (o *GetMailerEntriesOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/mailer_entries][%d] getMailerEntriesOK  %+v", 200, o.Payload)
}

func (o *GetMailerEntriesOK) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/mailer_entries][%d] getMailerEntriesOK  %+v", 200, o.Payload)
}

func (o *GetMailerEntriesOK) GetPayload() *GetMailerEntriesOKBody {
	return o.Payload
}

func (o *GetMailerEntriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetMailerEntriesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMailerEntriesDefault creates a GetMailerEntriesDefault with default headers values
func NewGetMailerEntriesDefault(code int) *GetMailerEntriesDefault {
	return &GetMailerEntriesDefault{
		_statusCode: code,
	}
}

/*
GetMailerEntriesDefault describes a response with status code -1, with default header values.

General Error
*/
type GetMailerEntriesDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get mailer entries default response has a 2xx status code
func (o *GetMailerEntriesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get mailer entries default response has a 3xx status code
func (o *GetMailerEntriesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get mailer entries default response has a 4xx status code
func (o *GetMailerEntriesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get mailer entries default response has a 5xx status code
func (o *GetMailerEntriesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get mailer entries default response a status code equal to that given
func (o *GetMailerEntriesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get mailer entries default response
func (o *GetMailerEntriesDefault) Code() int {
	return o._statusCode
}

func (o *GetMailerEntriesDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/mailer_entries][%d] getMailerEntries default  %+v", o._statusCode, o.Payload)
}

func (o *GetMailerEntriesDefault) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/mailer_entries][%d] getMailerEntries default  %+v", o._statusCode, o.Payload)
}

func (o *GetMailerEntriesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetMailerEntriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetMailerEntriesOKBody get mailer entries o k body
swagger:model GetMailerEntriesOKBody
*/
type GetMailerEntriesOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.MailerEntries `json:"data"`
}

// Validate validates this get mailer entries o k body
func (o *GetMailerEntriesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetMailerEntriesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getMailerEntriesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getMailerEntriesOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getMailerEntriesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get mailer entries o k body based on the context it is used
func (o *GetMailerEntriesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetMailerEntriesOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getMailerEntriesOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getMailerEntriesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetMailerEntriesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetMailerEntriesOKBody) UnmarshalBinary(b []byte) error {
	var res GetMailerEntriesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
