// Code generated by go-swagger; DO NOT EDIT.

package defaults

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/client-native/v2/models"
)

// GetDefaultsReader is a Reader for the GetDefaults structure.
type GetDefaultsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDefaultsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDefaultsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDefaultsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDefaultsOK creates a GetDefaultsOK with default headers values
func NewGetDefaultsOK() *GetDefaultsOK {
	return &GetDefaultsOK{}
}

/*GetDefaultsOK handles this case with default header values.

Successful operation
*/
type GetDefaultsOK struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetDefaultsOKBody
}

func (o *GetDefaultsOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/defaults][%d] getDefaultsOK  %+v", 200, o.Payload)
}

func (o *GetDefaultsOK) GetPayload() *GetDefaultsOKBody {
	return o.Payload
}

func (o *GetDefaultsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(GetDefaultsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDefaultsDefault creates a GetDefaultsDefault with default headers values
func NewGetDefaultsDefault(code int) *GetDefaultsDefault {
	return &GetDefaultsDefault{
		_statusCode: code,
	}
}

/*GetDefaultsDefault handles this case with default header values.

General Error
*/
type GetDefaultsDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get defaults default response
func (o *GetDefaultsDefault) Code() int {
	return o._statusCode
}

func (o *GetDefaultsDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/defaults][%d] getDefaults default  %+v", o._statusCode, o.Payload)
}

func (o *GetDefaultsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDefaultsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetDefaultsOKBody get defaults o k body
swagger:model GetDefaultsOKBody
*/
type GetDefaultsOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.Defaults `json:"data,omitempty"`
}

// Validate validates this get defaults o k body
func (o *GetDefaultsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDefaultsOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getDefaultsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDefaultsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDefaultsOKBody) UnmarshalBinary(b []byte) error {
	var res GetDefaultsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
