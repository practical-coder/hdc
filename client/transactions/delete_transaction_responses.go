// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// DeleteTransactionReader is a Reader for the DeleteTransaction structure.
type DeleteTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteTransactionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteTransactionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteTransactionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteTransactionNoContent creates a DeleteTransactionNoContent with default headers values
func NewDeleteTransactionNoContent() *DeleteTransactionNoContent {
	return &DeleteTransactionNoContent{}
}

/*
DeleteTransactionNoContent describes a response with status code 204, with default header values.

Transaction deleted
*/
type DeleteTransactionNoContent struct {
}

// IsSuccess returns true when this delete transaction no content response has a 2xx status code
func (o *DeleteTransactionNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete transaction no content response has a 3xx status code
func (o *DeleteTransactionNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete transaction no content response has a 4xx status code
func (o *DeleteTransactionNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete transaction no content response has a 5xx status code
func (o *DeleteTransactionNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete transaction no content response a status code equal to that given
func (o *DeleteTransactionNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteTransactionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/transactions/{id}][%d] deleteTransactionNoContent ", 204)
}

func (o *DeleteTransactionNoContent) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/transactions/{id}][%d] deleteTransactionNoContent ", 204)
}

func (o *DeleteTransactionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTransactionNotFound creates a DeleteTransactionNotFound with default headers values
func NewDeleteTransactionNotFound() *DeleteTransactionNotFound {
	return &DeleteTransactionNotFound{}
}

/*
DeleteTransactionNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteTransactionNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete transaction not found response has a 2xx status code
func (o *DeleteTransactionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete transaction not found response has a 3xx status code
func (o *DeleteTransactionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete transaction not found response has a 4xx status code
func (o *DeleteTransactionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete transaction not found response has a 5xx status code
func (o *DeleteTransactionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete transaction not found response a status code equal to that given
func (o *DeleteTransactionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteTransactionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/transactions/{id}][%d] deleteTransactionNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTransactionNotFound) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/transactions/{id}][%d] deleteTransactionNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTransactionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteTransactionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTransactionDefault creates a DeleteTransactionDefault with default headers values
func NewDeleteTransactionDefault(code int) *DeleteTransactionDefault {
	return &DeleteTransactionDefault{
		_statusCode: code,
	}
}

/*
DeleteTransactionDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteTransactionDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete transaction default response
func (o *DeleteTransactionDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete transaction default response has a 2xx status code
func (o *DeleteTransactionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete transaction default response has a 3xx status code
func (o *DeleteTransactionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete transaction default response has a 4xx status code
func (o *DeleteTransactionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete transaction default response has a 5xx status code
func (o *DeleteTransactionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete transaction default response a status code equal to that given
func (o *DeleteTransactionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteTransactionDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/transactions/{id}][%d] deleteTransaction default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTransactionDefault) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/transactions/{id}][%d] deleteTransaction default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTransactionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteTransactionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
