// Code generated by go-swagger; DO NOT EDIT.

package http_response_rule

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

// GetHTTPResponseRulesReader is a Reader for the GetHTTPResponseRules structure.
type GetHTTPResponseRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHTTPResponseRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHTTPResponseRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetHTTPResponseRulesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHTTPResponseRulesOK creates a GetHTTPResponseRulesOK with default headers values
func NewGetHTTPResponseRulesOK() *GetHTTPResponseRulesOK {
	return &GetHTTPResponseRulesOK{}
}

/*GetHTTPResponseRulesOK handles this case with default header values.

Successful operation
*/
type GetHTTPResponseRulesOK struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetHTTPResponseRulesOKBody
}

func (o *GetHTTPResponseRulesOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_response_rules][%d] getHttpResponseRulesOK  %+v", 200, o.Payload)
}

func (o *GetHTTPResponseRulesOK) GetPayload() *GetHTTPResponseRulesOKBody {
	return o.Payload
}

func (o *GetHTTPResponseRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(GetHTTPResponseRulesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHTTPResponseRulesDefault creates a GetHTTPResponseRulesDefault with default headers values
func NewGetHTTPResponseRulesDefault(code int) *GetHTTPResponseRulesDefault {
	return &GetHTTPResponseRulesDefault{
		_statusCode: code,
	}
}

/*GetHTTPResponseRulesDefault handles this case with default header values.

General Error
*/
type GetHTTPResponseRulesDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get HTTP response rules default response
func (o *GetHTTPResponseRulesDefault) Code() int {
	return o._statusCode
}

func (o *GetHTTPResponseRulesDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_response_rules][%d] getHTTPResponseRules default  %+v", o._statusCode, o.Payload)
}

func (o *GetHTTPResponseRulesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHTTPResponseRulesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetHTTPResponseRulesOKBody get HTTP response rules o k body
swagger:model GetHTTPResponseRulesOKBody
*/
type GetHTTPResponseRulesOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.HTTPResponseRules `json:"data"`
}

// Validate validates this get HTTP response rules o k body
func (o *GetHTTPResponseRulesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHTTPResponseRulesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getHttpResponseRulesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getHttpResponseRulesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetHTTPResponseRulesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHTTPResponseRulesOKBody) UnmarshalBinary(b []byte) error {
	var res GetHTTPResponseRulesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
