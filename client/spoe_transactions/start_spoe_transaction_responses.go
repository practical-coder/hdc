// Code generated by go-swagger; DO NOT EDIT.

package spoe_transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/client-native/v3/models"
)

// StartSpoeTransactionReader is a Reader for the StartSpoeTransaction structure.
type StartSpoeTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartSpoeTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewStartSpoeTransactionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewStartSpoeTransactionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewStartSpoeTransactionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartSpoeTransactionCreated creates a StartSpoeTransactionCreated with default headers values
func NewStartSpoeTransactionCreated() *StartSpoeTransactionCreated {
	return &StartSpoeTransactionCreated{}
}

/*StartSpoeTransactionCreated handles this case with default header values.

Transaction started
*/
type StartSpoeTransactionCreated struct {
	Payload *models.SpoeTransaction
}

func (o *StartSpoeTransactionCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe_transactions][%d] startSpoeTransactionCreated  %+v", 201, o.Payload)
}

func (o *StartSpoeTransactionCreated) GetPayload() *models.SpoeTransaction {
	return o.Payload
}

func (o *StartSpoeTransactionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SpoeTransaction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartSpoeTransactionTooManyRequests creates a StartSpoeTransactionTooManyRequests with default headers values
func NewStartSpoeTransactionTooManyRequests() *StartSpoeTransactionTooManyRequests {
	return &StartSpoeTransactionTooManyRequests{}
}

/*StartSpoeTransactionTooManyRequests handles this case with default header values.

Too many open transactions
*/
type StartSpoeTransactionTooManyRequests struct {
	Payload *StartSpoeTransactionTooManyRequestsBody
}

func (o *StartSpoeTransactionTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe_transactions][%d] startSpoeTransactionTooManyRequests  %+v", 429, o.Payload)
}

func (o *StartSpoeTransactionTooManyRequests) GetPayload() *StartSpoeTransactionTooManyRequestsBody {
	return o.Payload
}

func (o *StartSpoeTransactionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartSpoeTransactionTooManyRequestsBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartSpoeTransactionDefault creates a StartSpoeTransactionDefault with default headers values
func NewStartSpoeTransactionDefault(code int) *StartSpoeTransactionDefault {
	return &StartSpoeTransactionDefault{
		_statusCode: code,
	}
}

/*StartSpoeTransactionDefault handles this case with default header values.

General Error
*/
type StartSpoeTransactionDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the start spoe transaction default response
func (o *StartSpoeTransactionDefault) Code() int {
	return o._statusCode
}

func (o *StartSpoeTransactionDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe_transactions][%d] startSpoeTransaction default  %+v", o._statusCode, o.Payload)
}

func (o *StartSpoeTransactionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *StartSpoeTransactionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StartSpoeTransactionTooManyRequestsBody start spoe transaction too many requests body
swagger:model StartSpoeTransactionTooManyRequestsBody
*/
type StartSpoeTransactionTooManyRequestsBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this start spoe transaction too many requests body
func (o *StartSpoeTransactionTooManyRequestsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartSpoeTransactionTooManyRequestsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartSpoeTransactionTooManyRequestsBody) UnmarshalBinary(b []byte) error {
	var res StartSpoeTransactionTooManyRequestsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
