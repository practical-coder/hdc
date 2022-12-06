// Code generated by go-swagger; DO NOT EDIT.

package spoe_transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
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

/*
CommitSpoeTransactionOK describes a response with status code 200, with default header values.

Transaction successfully committed
*/
type CommitSpoeTransactionOK struct {
	Payload *models.SpoeTransaction
}

// IsSuccess returns true when this commit spoe transaction o k response has a 2xx status code
func (o *CommitSpoeTransactionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this commit spoe transaction o k response has a 3xx status code
func (o *CommitSpoeTransactionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this commit spoe transaction o k response has a 4xx status code
func (o *CommitSpoeTransactionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this commit spoe transaction o k response has a 5xx status code
func (o *CommitSpoeTransactionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this commit spoe transaction o k response a status code equal to that given
func (o *CommitSpoeTransactionOK) IsCode(code int) bool {
	return code == 200
}

func (o *CommitSpoeTransactionOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe_transactions/{id}][%d] commitSpoeTransactionOK  %+v", 200, o.Payload)
}

func (o *CommitSpoeTransactionOK) String() string {
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

/*
CommitSpoeTransactionAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CommitSpoeTransactionAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.SpoeTransaction
}

// IsSuccess returns true when this commit spoe transaction accepted response has a 2xx status code
func (o *CommitSpoeTransactionAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this commit spoe transaction accepted response has a 3xx status code
func (o *CommitSpoeTransactionAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this commit spoe transaction accepted response has a 4xx status code
func (o *CommitSpoeTransactionAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this commit spoe transaction accepted response has a 5xx status code
func (o *CommitSpoeTransactionAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this commit spoe transaction accepted response a status code equal to that given
func (o *CommitSpoeTransactionAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *CommitSpoeTransactionAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe_transactions/{id}][%d] commitSpoeTransactionAccepted  %+v", 202, o.Payload)
}

func (o *CommitSpoeTransactionAccepted) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe_transactions/{id}][%d] commitSpoeTransactionAccepted  %+v", 202, o.Payload)
}

func (o *CommitSpoeTransactionAccepted) GetPayload() *models.SpoeTransaction {
	return o.Payload
}

func (o *CommitSpoeTransactionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

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

/*
CommitSpoeTransactionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CommitSpoeTransactionBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this commit spoe transaction bad request response has a 2xx status code
func (o *CommitSpoeTransactionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this commit spoe transaction bad request response has a 3xx status code
func (o *CommitSpoeTransactionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this commit spoe transaction bad request response has a 4xx status code
func (o *CommitSpoeTransactionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this commit spoe transaction bad request response has a 5xx status code
func (o *CommitSpoeTransactionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this commit spoe transaction bad request response a status code equal to that given
func (o *CommitSpoeTransactionBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CommitSpoeTransactionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe_transactions/{id}][%d] commitSpoeTransactionBadRequest  %+v", 400, o.Payload)
}

func (o *CommitSpoeTransactionBadRequest) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe_transactions/{id}][%d] commitSpoeTransactionBadRequest  %+v", 400, o.Payload)
}

func (o *CommitSpoeTransactionBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CommitSpoeTransactionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCommitSpoeTransactionNotFound creates a CommitSpoeTransactionNotFound with default headers values
func NewCommitSpoeTransactionNotFound() *CommitSpoeTransactionNotFound {
	return &CommitSpoeTransactionNotFound{}
}

/*
CommitSpoeTransactionNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type CommitSpoeTransactionNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this commit spoe transaction not found response has a 2xx status code
func (o *CommitSpoeTransactionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this commit spoe transaction not found response has a 3xx status code
func (o *CommitSpoeTransactionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this commit spoe transaction not found response has a 4xx status code
func (o *CommitSpoeTransactionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this commit spoe transaction not found response has a 5xx status code
func (o *CommitSpoeTransactionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this commit spoe transaction not found response a status code equal to that given
func (o *CommitSpoeTransactionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CommitSpoeTransactionNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe_transactions/{id}][%d] commitSpoeTransactionNotFound  %+v", 404, o.Payload)
}

func (o *CommitSpoeTransactionNotFound) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe_transactions/{id}][%d] commitSpoeTransactionNotFound  %+v", 404, o.Payload)
}

func (o *CommitSpoeTransactionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CommitSpoeTransactionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCommitSpoeTransactionDefault creates a CommitSpoeTransactionDefault with default headers values
func NewCommitSpoeTransactionDefault(code int) *CommitSpoeTransactionDefault {
	return &CommitSpoeTransactionDefault{
		_statusCode: code,
	}
}

/*
CommitSpoeTransactionDefault describes a response with status code -1, with default header values.

General Error
*/
type CommitSpoeTransactionDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the commit spoe transaction default response
func (o *CommitSpoeTransactionDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this commit spoe transaction default response has a 2xx status code
func (o *CommitSpoeTransactionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this commit spoe transaction default response has a 3xx status code
func (o *CommitSpoeTransactionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this commit spoe transaction default response has a 4xx status code
func (o *CommitSpoeTransactionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this commit spoe transaction default response has a 5xx status code
func (o *CommitSpoeTransactionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this commit spoe transaction default response a status code equal to that given
func (o *CommitSpoeTransactionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CommitSpoeTransactionDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe_transactions/{id}][%d] commitSpoeTransaction default  %+v", o._statusCode, o.Payload)
}

func (o *CommitSpoeTransactionDefault) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe_transactions/{id}][%d] commitSpoeTransaction default  %+v", o._statusCode, o.Payload)
}

func (o *CommitSpoeTransactionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CommitSpoeTransactionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
