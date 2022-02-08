// Code generated by go-swagger; DO NOT EDIT.

package http_request_rule

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

	"github.com/haproxytech/client-native/v3/models"
)

// GetHTTPRequestRulesReader is a Reader for the GetHTTPRequestRules structure.
type GetHTTPRequestRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHTTPRequestRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHTTPRequestRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetHTTPRequestRulesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHTTPRequestRulesOK creates a GetHTTPRequestRulesOK with default headers values
func NewGetHTTPRequestRulesOK() *GetHTTPRequestRulesOK {
	return &GetHTTPRequestRulesOK{}
}

/*GetHTTPRequestRulesOK handles this case with default header values.

Successful operation
*/
type GetHTTPRequestRulesOK struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetHTTPRequestRulesOKBody
}

func (o *GetHTTPRequestRulesOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_request_rules][%d] getHttpRequestRulesOK  %+v", 200, o.Payload)
}

func (o *GetHTTPRequestRulesOK) GetPayload() *GetHTTPRequestRulesOKBody {
	return o.Payload
}

func (o *GetHTTPRequestRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(GetHTTPRequestRulesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHTTPRequestRulesDefault creates a GetHTTPRequestRulesDefault with default headers values
func NewGetHTTPRequestRulesDefault(code int) *GetHTTPRequestRulesDefault {
	return &GetHTTPRequestRulesDefault{
		_statusCode: code,
	}
}

/*GetHTTPRequestRulesDefault handles this case with default header values.

General Error
*/
type GetHTTPRequestRulesDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get HTTP request rules default response
func (o *GetHTTPRequestRulesDefault) Code() int {
	return o._statusCode
}

func (o *GetHTTPRequestRulesDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_request_rules][%d] getHTTPRequestRules default  %+v", o._statusCode, o.Payload)
}

func (o *GetHTTPRequestRulesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHTTPRequestRulesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetHTTPRequestRulesOKBody get HTTP request rules o k body
swagger:model GetHTTPRequestRulesOKBody
*/
type GetHTTPRequestRulesOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.HTTPRequestRules `json:"data"`
}

// Validate validates this get HTTP request rules o k body
func (o *GetHTTPRequestRulesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHTTPRequestRulesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getHttpRequestRulesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getHttpRequestRulesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetHTTPRequestRulesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHTTPRequestRulesOKBody) UnmarshalBinary(b []byte) error {
	var res GetHTTPRequestRulesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
