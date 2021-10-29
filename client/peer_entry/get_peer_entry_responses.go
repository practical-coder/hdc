// Code generated by go-swagger; DO NOT EDIT.

package peer_entry

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

	"github.com/haproxytech/client-native/models"
)

// GetPeerEntryReader is a Reader for the GetPeerEntry structure.
type GetPeerEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPeerEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPeerEntryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetPeerEntryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetPeerEntryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPeerEntryOK creates a GetPeerEntryOK with default headers values
func NewGetPeerEntryOK() *GetPeerEntryOK {
	return &GetPeerEntryOK{}
}

/* GetPeerEntryOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetPeerEntryOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetPeerEntryOKBody
}

func (o *GetPeerEntryOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/peer_entries/{name}][%d] getPeerEntryOK  %+v", 200, o.Payload)
}
func (o *GetPeerEntryOK) GetPayload() *GetPeerEntryOKBody {
	return o.Payload
}

func (o *GetPeerEntryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetPeerEntryOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPeerEntryNotFound creates a GetPeerEntryNotFound with default headers values
func NewGetPeerEntryNotFound() *GetPeerEntryNotFound {
	return &GetPeerEntryNotFound{}
}

/* GetPeerEntryNotFound describes a response with status code 404, with default header values.

The specified resource already exists
*/
type GetPeerEntryNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetPeerEntryNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/peer_entries/{name}][%d] getPeerEntryNotFound  %+v", 404, o.Payload)
}
func (o *GetPeerEntryNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPeerEntryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPeerEntryDefault creates a GetPeerEntryDefault with default headers values
func NewGetPeerEntryDefault(code int) *GetPeerEntryDefault {
	return &GetPeerEntryDefault{
		_statusCode: code,
	}
}

/* GetPeerEntryDefault describes a response with status code -1, with default header values.

General Error
*/
type GetPeerEntryDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get peer entry default response
func (o *GetPeerEntryDefault) Code() int {
	return o._statusCode
}

func (o *GetPeerEntryDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/peer_entries/{name}][%d] getPeerEntry default  %+v", o._statusCode, o.Payload)
}
func (o *GetPeerEntryDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPeerEntryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetPeerEntryOKBody get peer entry o k body
swagger:model GetPeerEntryOKBody
*/
type GetPeerEntryOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.PeerEntry `json:"data,omitempty"`
}

// Validate validates this get peer entry o k body
func (o *GetPeerEntryOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPeerEntryOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getPeerEntryOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getPeerEntryOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get peer entry o k body based on the context it is used
func (o *GetPeerEntryOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPeerEntryOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getPeerEntryOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getPeerEntryOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPeerEntryOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPeerEntryOKBody) UnmarshalBinary(b []byte) error {
	var res GetPeerEntryOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
