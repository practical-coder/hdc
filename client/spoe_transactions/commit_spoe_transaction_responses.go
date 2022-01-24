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

// CommitSpoeTransactionReader is a Reader for the CommitSpoeTransaction structure.
type CommitSpoeTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CommitSpoeTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCommitSpoeTransactionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCommitSpoeTransactionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCommitSpoeTransactionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCommitSpoeTransactionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCommitSpoeTransactionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCommitSpoeTransactionOK creates a CommitSpoeTransactionOK with default headers values
func NewCommitSpoeTransactionOK() *CommitSpoeTransactionOK {
	return &CommitSpoeTransactionOK{}
}

/*CommitSpoeTransactionOK handles this case with default header values.

Transaction successfully committed
*/
type CommitSpoeTransactionOK struct {
	Payload *models.SpoeTransaction
}

func (o *CommitSpoeTransactionOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe_transactions/{id}][%d] commitSpoeTransactionOK  %+v", 200, o.Payload)
}

func (o *CommitSpoeTransactionOK) GetPayload() *models.SpoeTransaction {
	return o.Payload
}

func (o *CommitSpoeTransactionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SpoeTransaction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCommitSpoeTransactionAccepted creates a CommitSpoeTransactionAccepted with default headers values
func NewCommitSpoeTransactionAccepted() *CommitSpoeTransactionAccepted {
	return &CommitSpoeTransactionAccepted{}
}

/*CommitSpoeTransactionAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type CommitSpoeTransactionAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string

	Payload *models.SpoeTransaction
}

func (o *CommitSpoeTransactionAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe_transactions/{id}][%d] commitSpoeTransactionAccepted  %+v", 202, o.Payload)
}

func (o *CommitSpoeTransactionAccepted) GetPayload() *models.SpoeTransaction {
	return o.Payload
}

func (o *CommitSpoeTransactionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	o.Payload = new(models.SpoeTransaction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCommitSpoeTransactionBadRequest creates a CommitSpoeTransactionBadRequest with default headers values
func NewCommitSpoeTransactionBadRequest() *CommitSpoeTransactionBadRequest {
	return &CommitSpoeTransactionBadRequest{}
}

/*CommitSpoeTransactionBadRequest handles this case with default header values.

Bad request
*/
type CommitSpoeTransactionBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CommitSpoeTransactionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe_transactions/{id}][%d] commitSpoeTransactionBadRequest  %+v", 400, o.Payload)
}

func (o *CommitSpoeTransactionBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CommitSpoeTransactionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCommitSpoeTransactionNotFound creates a CommitSpoeTransactionNotFound with default headers values
func NewCommitSpoeTransactionNotFound() *CommitSpoeTransactionNotFound {
	return &CommitSpoeTransactionNotFound{}
}

/*CommitSpoeTransactionNotFound handles this case with default header values.

The specified resource was not found
*/
type CommitSpoeTransactionNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CommitSpoeTransactionNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe_transactions/{id}][%d] commitSpoeTransactionNotFound  %+v", 404, o.Payload)
}

func (o *CommitSpoeTransactionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CommitSpoeTransactionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCommitSpoeTransactionDefault creates a CommitSpoeTransactionDefault with default headers values
func NewCommitSpoeTransactionDefault(code int) *CommitSpoeTransactionDefault {
	return &CommitSpoeTransactionDefault{
		_statusCode: code,
	}
}

/*CommitSpoeTransactionDefault handles this case with default header values.

General Error
*/
type CommitSpoeTransactionDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the commit spoe transaction default response
func (o *CommitSpoeTransactionDefault) Code() int {
	return o._statusCode
}

func (o *CommitSpoeTransactionDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe_transactions/{id}][%d] commitSpoeTransaction default  %+v", o._statusCode, o.Payload)
}

func (o *CommitSpoeTransactionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CommitSpoeTransactionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
