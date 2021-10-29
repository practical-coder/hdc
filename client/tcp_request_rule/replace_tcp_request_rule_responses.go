// Code generated by go-swagger; DO NOT EDIT.

package tcp_request_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/models"
)

// ReplaceTCPRequestRuleReader is a Reader for the ReplaceTCPRequestRule structure.
type ReplaceTCPRequestRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceTCPRequestRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceTCPRequestRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceTCPRequestRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceTCPRequestRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceTCPRequestRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceTCPRequestRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceTCPRequestRuleOK creates a ReplaceTCPRequestRuleOK with default headers values
func NewReplaceTCPRequestRuleOK() *ReplaceTCPRequestRuleOK {
	return &ReplaceTCPRequestRuleOK{}
}

/* ReplaceTCPRequestRuleOK describes a response with status code 200, with default header values.

TCP Request Rule replaced
*/
type ReplaceTCPRequestRuleOK struct {
	Payload *models.TCPRequestRule
}

func (o *ReplaceTCPRequestRuleOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/tcp_request_rules/{index}][%d] replaceTcpRequestRuleOK  %+v", 200, o.Payload)
}
func (o *ReplaceTCPRequestRuleOK) GetPayload() *models.TCPRequestRule {
	return o.Payload
}

func (o *ReplaceTCPRequestRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TCPRequestRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceTCPRequestRuleAccepted creates a ReplaceTCPRequestRuleAccepted with default headers values
func NewReplaceTCPRequestRuleAccepted() *ReplaceTCPRequestRuleAccepted {
	return &ReplaceTCPRequestRuleAccepted{}
}

/* ReplaceTCPRequestRuleAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type ReplaceTCPRequestRuleAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.TCPRequestRule
}

func (o *ReplaceTCPRequestRuleAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/tcp_request_rules/{index}][%d] replaceTcpRequestRuleAccepted  %+v", 202, o.Payload)
}
func (o *ReplaceTCPRequestRuleAccepted) GetPayload() *models.TCPRequestRule {
	return o.Payload
}

func (o *ReplaceTCPRequestRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.TCPRequestRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceTCPRequestRuleBadRequest creates a ReplaceTCPRequestRuleBadRequest with default headers values
func NewReplaceTCPRequestRuleBadRequest() *ReplaceTCPRequestRuleBadRequest {
	return &ReplaceTCPRequestRuleBadRequest{}
}

/* ReplaceTCPRequestRuleBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceTCPRequestRuleBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceTCPRequestRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/tcp_request_rules/{index}][%d] replaceTcpRequestRuleBadRequest  %+v", 400, o.Payload)
}
func (o *ReplaceTCPRequestRuleBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceTCPRequestRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceTCPRequestRuleNotFound creates a ReplaceTCPRequestRuleNotFound with default headers values
func NewReplaceTCPRequestRuleNotFound() *ReplaceTCPRequestRuleNotFound {
	return &ReplaceTCPRequestRuleNotFound{}
}

/* ReplaceTCPRequestRuleNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type ReplaceTCPRequestRuleNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceTCPRequestRuleNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/tcp_request_rules/{index}][%d] replaceTcpRequestRuleNotFound  %+v", 404, o.Payload)
}
func (o *ReplaceTCPRequestRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceTCPRequestRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceTCPRequestRuleDefault creates a ReplaceTCPRequestRuleDefault with default headers values
func NewReplaceTCPRequestRuleDefault(code int) *ReplaceTCPRequestRuleDefault {
	return &ReplaceTCPRequestRuleDefault{
		_statusCode: code,
	}
}

/* ReplaceTCPRequestRuleDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceTCPRequestRuleDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the replace TCP request rule default response
func (o *ReplaceTCPRequestRuleDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceTCPRequestRuleDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/tcp_request_rules/{index}][%d] replaceTCPRequestRule default  %+v", o._statusCode, o.Payload)
}
func (o *ReplaceTCPRequestRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceTCPRequestRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
