// Code generated by go-swagger; DO NOT EDIT.

package userlist

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/client-native/v3/models"
)

// GetUserlistReader is a Reader for the GetUserlist structure.
type GetUserlistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserlistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserlistOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetUserlistNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetUserlistDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUserlistOK creates a GetUserlistOK with default headers values
func NewGetUserlistOK() *GetUserlistOK {
	return &GetUserlistOK{}
}

/*GetUserlistOK handles this case with default header values.

Successful operation
*/
type GetUserlistOK struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetUserlistOKBody
}

func (o *GetUserlistOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/userlists/{name}][%d] getUserlistOK  %+v", 200, o.Payload)
}

func (o *GetUserlistOK) GetPayload() *GetUserlistOKBody {
	return o.Payload
}

func (o *GetUserlistOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(GetUserlistOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserlistNotFound creates a GetUserlistNotFound with default headers values
func NewGetUserlistNotFound() *GetUserlistNotFound {
	return &GetUserlistNotFound{}
}

/*GetUserlistNotFound handles this case with default header values.

The specified resource already exists
*/
type GetUserlistNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetUserlistNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/userlists/{name}][%d] getUserlistNotFound  %+v", 404, o.Payload)
}

func (o *GetUserlistNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUserlistNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserlistDefault creates a GetUserlistDefault with default headers values
func NewGetUserlistDefault(code int) *GetUserlistDefault {
	return &GetUserlistDefault{
		_statusCode: code,
	}
}

/*GetUserlistDefault handles this case with default header values.

General Error
*/
type GetUserlistDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get userlist default response
func (o *GetUserlistDefault) Code() int {
	return o._statusCode
}

func (o *GetUserlistDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/userlists/{name}][%d] getUserlist default  %+v", o._statusCode, o.Payload)
}

func (o *GetUserlistDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUserlistDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetUserlistOKBody get userlist o k body
swagger:model GetUserlistOKBody
*/
type GetUserlistOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.Userlist `json:"data,omitempty"`
}

// Validate validates this get userlist o k body
func (o *GetUserlistOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUserlistOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getUserlistOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUserlistOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUserlistOKBody) UnmarshalBinary(b []byte) error {
	var res GetUserlistOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
