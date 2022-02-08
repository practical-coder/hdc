// Code generated by go-swagger; DO NOT EDIT.

package backend_switching_rule

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

// GetBackendSwitchingRulesReader is a Reader for the GetBackendSwitchingRules structure.
type GetBackendSwitchingRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackendSwitchingRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBackendSwitchingRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetBackendSwitchingRulesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBackendSwitchingRulesOK creates a GetBackendSwitchingRulesOK with default headers values
func NewGetBackendSwitchingRulesOK() *GetBackendSwitchingRulesOK {
	return &GetBackendSwitchingRulesOK{}
}

/*GetBackendSwitchingRulesOK handles this case with default header values.

Successful operation
*/
type GetBackendSwitchingRulesOK struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetBackendSwitchingRulesOKBody
}

func (o *GetBackendSwitchingRulesOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/backend_switching_rules][%d] getBackendSwitchingRulesOK  %+v", 200, o.Payload)
}

func (o *GetBackendSwitchingRulesOK) GetPayload() *GetBackendSwitchingRulesOKBody {
	return o.Payload
}

func (o *GetBackendSwitchingRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(GetBackendSwitchingRulesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackendSwitchingRulesDefault creates a GetBackendSwitchingRulesDefault with default headers values
func NewGetBackendSwitchingRulesDefault(code int) *GetBackendSwitchingRulesDefault {
	return &GetBackendSwitchingRulesDefault{
		_statusCode: code,
	}
}

/*GetBackendSwitchingRulesDefault handles this case with default header values.

General Error
*/
type GetBackendSwitchingRulesDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get backend switching rules default response
func (o *GetBackendSwitchingRulesDefault) Code() int {
	return o._statusCode
}

func (o *GetBackendSwitchingRulesDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/backend_switching_rules][%d] getBackendSwitchingRules default  %+v", o._statusCode, o.Payload)
}

func (o *GetBackendSwitchingRulesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBackendSwitchingRulesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetBackendSwitchingRulesOKBody get backend switching rules o k body
swagger:model GetBackendSwitchingRulesOKBody
*/
type GetBackendSwitchingRulesOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.BackendSwitchingRules `json:"data"`
}

// Validate validates this get backend switching rules o k body
func (o *GetBackendSwitchingRulesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBackendSwitchingRulesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getBackendSwitchingRulesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getBackendSwitchingRulesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetBackendSwitchingRulesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBackendSwitchingRulesOKBody) UnmarshalBinary(b []byte) error {
	var res GetBackendSwitchingRulesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
