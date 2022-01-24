// Code generated by go-swagger; DO NOT EDIT.

package spoe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/haproxytech/client-native/v2/models"
)

// GetSpoeScopesReader is a Reader for the GetSpoeScopes structure.
type GetSpoeScopesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpoeScopesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSpoeScopesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSpoeScopesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSpoeScopesOK creates a GetSpoeScopesOK with default headers values
func NewGetSpoeScopesOK() *GetSpoeScopesOK {
	return &GetSpoeScopesOK{}
}

/*GetSpoeScopesOK handles this case with default header values.

Successful operation
*/
type GetSpoeScopesOK struct {
	/*Spoe configuration file version
	 */
	ConfigurationVersion string

	Payload *GetSpoeScopesOKBody
}

func (o *GetSpoeScopesOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_scopes][%d] getSpoeScopesOK  %+v", 200, o.Payload)
}

func (o *GetSpoeScopesOK) GetPayload() *GetSpoeScopesOKBody {
	return o.Payload
}

func (o *GetSpoeScopesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(GetSpoeScopesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpoeScopesDefault creates a GetSpoeScopesDefault with default headers values
func NewGetSpoeScopesDefault(code int) *GetSpoeScopesDefault {
	return &GetSpoeScopesDefault{
		_statusCode: code,
	}
}

/*GetSpoeScopesDefault handles this case with default header values.

General Error
*/
type GetSpoeScopesDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get spoe scopes default response
func (o *GetSpoeScopesDefault) Code() int {
	return o._statusCode
}

func (o *GetSpoeScopesDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_scopes][%d] getSpoeScopes default  %+v", o._statusCode, o.Payload)
}

func (o *GetSpoeScopesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSpoeScopesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetSpoeScopesOKBody get spoe scopes o k body
swagger:model GetSpoeScopesOKBody
*/
type GetSpoeScopesOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.SpoeScopes `json:"data"`
}

// Validate validates this get spoe scopes o k body
func (o *GetSpoeScopesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSpoeScopesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getSpoeScopesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getSpoeScopesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSpoeScopesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSpoeScopesOKBody) UnmarshalBinary(b []byte) error {
	var res GetSpoeScopesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
