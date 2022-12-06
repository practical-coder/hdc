// Code generated by go-swagger; DO NOT EDIT.

package peer

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

	"github.com/haproxytech/client-native/v4/models"
)

// GetPeerSectionReader is a Reader for the GetPeerSection structure.
type GetPeerSectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPeerSectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPeerSectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetPeerSectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetPeerSectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPeerSectionOK creates a GetPeerSectionOK with default headers values
func NewGetPeerSectionOK() *GetPeerSectionOK {
	return &GetPeerSectionOK{}
}

/*
GetPeerSectionOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetPeerSectionOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetPeerSectionOKBody
}

// IsSuccess returns true when this get peer section o k response has a 2xx status code
func (o *GetPeerSectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get peer section o k response has a 3xx status code
func (o *GetPeerSectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get peer section o k response has a 4xx status code
func (o *GetPeerSectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get peer section o k response has a 5xx status code
func (o *GetPeerSectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get peer section o k response a status code equal to that given
func (o *GetPeerSectionOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPeerSectionOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/peer_section/{name}][%d] getPeerSectionOK  %+v", 200, o.Payload)
}

func (o *GetPeerSectionOK) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/peer_section/{name}][%d] getPeerSectionOK  %+v", 200, o.Payload)
}

func (o *GetPeerSectionOK) GetPayload() *GetPeerSectionOKBody {
	return o.Payload
}

func (o *GetPeerSectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetPeerSectionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPeerSectionNotFound creates a GetPeerSectionNotFound with default headers values
func NewGetPeerSectionNotFound() *GetPeerSectionNotFound {
	return &GetPeerSectionNotFound{}
}

/*
GetPeerSectionNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetPeerSectionNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this get peer section not found response has a 2xx status code
func (o *GetPeerSectionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get peer section not found response has a 3xx status code
func (o *GetPeerSectionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get peer section not found response has a 4xx status code
func (o *GetPeerSectionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get peer section not found response has a 5xx status code
func (o *GetPeerSectionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get peer section not found response a status code equal to that given
func (o *GetPeerSectionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetPeerSectionNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/peer_section/{name}][%d] getPeerSectionNotFound  %+v", 404, o.Payload)
}

func (o *GetPeerSectionNotFound) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/peer_section/{name}][%d] getPeerSectionNotFound  %+v", 404, o.Payload)
}

func (o *GetPeerSectionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPeerSectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPeerSectionDefault creates a GetPeerSectionDefault with default headers values
func NewGetPeerSectionDefault(code int) *GetPeerSectionDefault {
	return &GetPeerSectionDefault{
		_statusCode: code,
	}
}

/*
GetPeerSectionDefault describes a response with status code -1, with default header values.

General Error
*/
type GetPeerSectionDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get peer section default response
func (o *GetPeerSectionDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get peer section default response has a 2xx status code
func (o *GetPeerSectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get peer section default response has a 3xx status code
func (o *GetPeerSectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get peer section default response has a 4xx status code
func (o *GetPeerSectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get peer section default response has a 5xx status code
func (o *GetPeerSectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get peer section default response a status code equal to that given
func (o *GetPeerSectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetPeerSectionDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/peer_section/{name}][%d] getPeerSection default  %+v", o._statusCode, o.Payload)
}

func (o *GetPeerSectionDefault) String() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/peer_section/{name}][%d] getPeerSection default  %+v", o._statusCode, o.Payload)
}

func (o *GetPeerSectionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPeerSectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetPeerSectionOKBody get peer section o k body
swagger:model GetPeerSectionOKBody
*/
type GetPeerSectionOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.PeerSection `json:"data,omitempty"`
}

// Validate validates this get peer section o k body
func (o *GetPeerSectionOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPeerSectionOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getPeerSectionOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getPeerSectionOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get peer section o k body based on the context it is used
func (o *GetPeerSectionOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPeerSectionOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getPeerSectionOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getPeerSectionOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPeerSectionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPeerSectionOKBody) UnmarshalBinary(b []byte) error {
	var res GetPeerSectionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
