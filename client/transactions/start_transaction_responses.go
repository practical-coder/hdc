// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/client-native/v2/models"
)

// StartTransactionReader is a Reader for the StartTransaction structure.
type StartTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewStartTransactionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewStartTransactionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewStartTransactionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartTransactionCreated creates a StartTransactionCreated with default headers values
func NewStartTransactionCreated() *StartTransactionCreated {
	return &StartTransactionCreated{}
}

/*StartTransactionCreated handles this case with default header values.

Transaction started
*/
type StartTransactionCreated struct {
	Payload *models.Transaction
}

func (o *StartTransactionCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/transactions][%d] startTransactionCreated  %+v", 201, o.Payload)
}

func (o *StartTransactionCreated) GetPayload() *models.Transaction {
	return o.Payload
}

func (o *StartTransactionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Transaction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartTransactionTooManyRequests creates a StartTransactionTooManyRequests with default headers values
func NewStartTransactionTooManyRequests() *StartTransactionTooManyRequests {
	return &StartTransactionTooManyRequests{}
}

/*StartTransactionTooManyRequests handles this case with default header values.

Too many open transactions
*/
type StartTransactionTooManyRequests struct {
	Payload *StartTransactionTooManyRequestsBody
}

func (o *StartTransactionTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/transactions][%d] startTransactionTooManyRequests  %+v", 429, o.Payload)
}

func (o *StartTransactionTooManyRequests) GetPayload() *StartTransactionTooManyRequestsBody {
	return o.Payload
}

func (o *StartTransactionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartTransactionTooManyRequestsBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartTransactionDefault creates a StartTransactionDefault with default headers values
func NewStartTransactionDefault(code int) *StartTransactionDefault {
	return &StartTransactionDefault{
		_statusCode: code,
	}
}

/*StartTransactionDefault handles this case with default header values.

General Error
*/
type StartTransactionDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the start transaction default response
func (o *StartTransactionDefault) Code() int {
	return o._statusCode
}

func (o *StartTransactionDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/transactions][%d] startTransaction default  %+v", o._statusCode, o.Payload)
}

func (o *StartTransactionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *StartTransactionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StartTransactionTooManyRequestsBody start transaction too many requests body
swagger:model StartTransactionTooManyRequestsBody
*/
type StartTransactionTooManyRequestsBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this start transaction too many requests body
func (o *StartTransactionTooManyRequestsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartTransactionTooManyRequestsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartTransactionTooManyRequestsBody) UnmarshalBinary(b []byte) error {
	var res StartTransactionTooManyRequestsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
