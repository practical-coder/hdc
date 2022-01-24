// Code generated by go-swagger; DO NOT EDIT.

package spoe_transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v2/models"
)

// DeleteSpoeTransactionReader is a Reader for the DeleteSpoeTransaction structure.
type DeleteSpoeTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSpoeTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteSpoeTransactionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteSpoeTransactionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteSpoeTransactionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSpoeTransactionNoContent creates a DeleteSpoeTransactionNoContent with default headers values
func NewDeleteSpoeTransactionNoContent() *DeleteSpoeTransactionNoContent {
	return &DeleteSpoeTransactionNoContent{}
}

/*DeleteSpoeTransactionNoContent handles this case with default header values.

Transaction deleted
*/
type DeleteSpoeTransactionNoContent struct {
}

func (o *DeleteSpoeTransactionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe_transactions/{id}][%d] deleteSpoeTransactionNoContent ", 204)
}

func (o *DeleteSpoeTransactionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSpoeTransactionNotFound creates a DeleteSpoeTransactionNotFound with default headers values
func NewDeleteSpoeTransactionNotFound() *DeleteSpoeTransactionNotFound {
	return &DeleteSpoeTransactionNotFound{}
}

/*DeleteSpoeTransactionNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteSpoeTransactionNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteSpoeTransactionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe_transactions/{id}][%d] deleteSpoeTransactionNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSpoeTransactionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteSpoeTransactionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSpoeTransactionDefault creates a DeleteSpoeTransactionDefault with default headers values
func NewDeleteSpoeTransactionDefault(code int) *DeleteSpoeTransactionDefault {
	return &DeleteSpoeTransactionDefault{
		_statusCode: code,
	}
}

/*DeleteSpoeTransactionDefault handles this case with default header values.

General Error
*/
type DeleteSpoeTransactionDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete spoe transaction default response
func (o *DeleteSpoeTransactionDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSpoeTransactionDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe_transactions/{id}][%d] deleteSpoeTransaction default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSpoeTransactionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteSpoeTransactionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
