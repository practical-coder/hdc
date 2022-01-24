// Code generated by go-swagger; DO NOT EDIT.

package bind

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

// GetBindReader is a Reader for the GetBind structure.
type GetBindReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBindReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBindOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetBindNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetBindDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBindOK creates a GetBindOK with default headers values
func NewGetBindOK() *GetBindOK {
	return &GetBindOK{}
}

/*GetBindOK handles this case with default header values.

Successful operation
*/
type GetBindOK struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetBindOKBody
}

func (o *GetBindOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/binds/{name}][%d] getBindOK  %+v", 200, o.Payload)
}

func (o *GetBindOK) GetPayload() *GetBindOKBody {
	return o.Payload
}

func (o *GetBindOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(GetBindOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBindNotFound creates a GetBindNotFound with default headers values
func NewGetBindNotFound() *GetBindNotFound {
	return &GetBindNotFound{}
}

/*GetBindNotFound handles this case with default header values.

The specified resource already exists
*/
type GetBindNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetBindNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/binds/{name}][%d] getBindNotFound  %+v", 404, o.Payload)
}

func (o *GetBindNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBindNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBindDefault creates a GetBindDefault with default headers values
func NewGetBindDefault(code int) *GetBindDefault {
	return &GetBindDefault{
		_statusCode: code,
	}
}

/*GetBindDefault handles this case with default header values.

General Error
*/
type GetBindDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get bind default response
func (o *GetBindDefault) Code() int {
	return o._statusCode
}

func (o *GetBindDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/binds/{name}][%d] getBind default  %+v", o._statusCode, o.Payload)
}

func (o *GetBindDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBindDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetBindOKBody get bind o k body
swagger:model GetBindOKBody
*/
type GetBindOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.Bind `json:"data,omitempty"`
}

// Validate validates this get bind o k body
func (o *GetBindOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBindOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getBindOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetBindOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBindOKBody) UnmarshalBinary(b []byte) error {
	var res GetBindOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
