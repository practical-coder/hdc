// Code generated by go-swagger; DO NOT EDIT.

package tcp_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v2/models"
)

// ReplaceTCPResponseRuleReader is a Reader for the ReplaceTCPResponseRule structure.
type ReplaceTCPResponseRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceTCPResponseRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceTCPResponseRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceTCPResponseRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceTCPResponseRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceTCPResponseRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceTCPResponseRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceTCPResponseRuleOK creates a ReplaceTCPResponseRuleOK with default headers values
func NewReplaceTCPResponseRuleOK() *ReplaceTCPResponseRuleOK {
	return &ReplaceTCPResponseRuleOK{}
}

/* ReplaceTCPResponseRuleOK describes a response with status code 200, with default header values.

TCP Response Rule replaced
*/
type ReplaceTCPResponseRuleOK struct {
	Payload *models.TCPResponseRule
}

func (o *ReplaceTCPResponseRuleOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/tcp_response_rules/{index}][%d] replaceTcpResponseRuleOK  %+v", 200, o.Payload)
}
func (o *ReplaceTCPResponseRuleOK) GetPayload() *models.TCPResponseRule {
	return o.Payload
}

func (o *ReplaceTCPResponseRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TCPResponseRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceTCPResponseRuleAccepted creates a ReplaceTCPResponseRuleAccepted with default headers values
func NewReplaceTCPResponseRuleAccepted() *ReplaceTCPResponseRuleAccepted {
	return &ReplaceTCPResponseRuleAccepted{}
}

/* ReplaceTCPResponseRuleAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type ReplaceTCPResponseRuleAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.TCPResponseRule
}

func (o *ReplaceTCPResponseRuleAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/tcp_response_rules/{index}][%d] replaceTcpResponseRuleAccepted  %+v", 202, o.Payload)
}
func (o *ReplaceTCPResponseRuleAccepted) GetPayload() *models.TCPResponseRule {
	return o.Payload
}

func (o *ReplaceTCPResponseRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.TCPResponseRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceTCPResponseRuleBadRequest creates a ReplaceTCPResponseRuleBadRequest with default headers values
func NewReplaceTCPResponseRuleBadRequest() *ReplaceTCPResponseRuleBadRequest {
	return &ReplaceTCPResponseRuleBadRequest{}
}

/* ReplaceTCPResponseRuleBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceTCPResponseRuleBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceTCPResponseRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/tcp_response_rules/{index}][%d] replaceTcpResponseRuleBadRequest  %+v", 400, o.Payload)
}
func (o *ReplaceTCPResponseRuleBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceTCPResponseRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceTCPResponseRuleNotFound creates a ReplaceTCPResponseRuleNotFound with default headers values
func NewReplaceTCPResponseRuleNotFound() *ReplaceTCPResponseRuleNotFound {
	return &ReplaceTCPResponseRuleNotFound{}
}

/* ReplaceTCPResponseRuleNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type ReplaceTCPResponseRuleNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceTCPResponseRuleNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/tcp_response_rules/{index}][%d] replaceTcpResponseRuleNotFound  %+v", 404, o.Payload)
}
func (o *ReplaceTCPResponseRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceTCPResponseRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceTCPResponseRuleDefault creates a ReplaceTCPResponseRuleDefault with default headers values
func NewReplaceTCPResponseRuleDefault(code int) *ReplaceTCPResponseRuleDefault {
	return &ReplaceTCPResponseRuleDefault{
		_statusCode: code,
	}
}

/* ReplaceTCPResponseRuleDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceTCPResponseRuleDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the replace TCP response rule default response
func (o *ReplaceTCPResponseRuleDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceTCPResponseRuleDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/tcp_response_rules/{index}][%d] replaceTCPResponseRule default  %+v", o._statusCode, o.Payload)
}
func (o *ReplaceTCPResponseRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceTCPResponseRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
