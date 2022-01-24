// Code generated by go-swagger; DO NOT EDIT.

package server_switching_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v2/models"
)

// ReplaceServerSwitchingRuleReader is a Reader for the ReplaceServerSwitchingRule structure.
type ReplaceServerSwitchingRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceServerSwitchingRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceServerSwitchingRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceServerSwitchingRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceServerSwitchingRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceServerSwitchingRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceServerSwitchingRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceServerSwitchingRuleOK creates a ReplaceServerSwitchingRuleOK with default headers values
func NewReplaceServerSwitchingRuleOK() *ReplaceServerSwitchingRuleOK {
	return &ReplaceServerSwitchingRuleOK{}
}

/*ReplaceServerSwitchingRuleOK handles this case with default header values.

Server Switching Rule replaced
*/
type ReplaceServerSwitchingRuleOK struct {
	Payload *models.ServerSwitchingRule
}

func (o *ReplaceServerSwitchingRuleOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/server_switching_rules/{index}][%d] replaceServerSwitchingRuleOK  %+v", 200, o.Payload)
}

func (o *ReplaceServerSwitchingRuleOK) GetPayload() *models.ServerSwitchingRule {
	return o.Payload
}

func (o *ReplaceServerSwitchingRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerSwitchingRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceServerSwitchingRuleAccepted creates a ReplaceServerSwitchingRuleAccepted with default headers values
func NewReplaceServerSwitchingRuleAccepted() *ReplaceServerSwitchingRuleAccepted {
	return &ReplaceServerSwitchingRuleAccepted{}
}

/*ReplaceServerSwitchingRuleAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type ReplaceServerSwitchingRuleAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string

	Payload *models.ServerSwitchingRule
}

func (o *ReplaceServerSwitchingRuleAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/server_switching_rules/{index}][%d] replaceServerSwitchingRuleAccepted  %+v", 202, o.Payload)
}

func (o *ReplaceServerSwitchingRuleAccepted) GetPayload() *models.ServerSwitchingRule {
	return o.Payload
}

func (o *ReplaceServerSwitchingRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	o.Payload = new(models.ServerSwitchingRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceServerSwitchingRuleBadRequest creates a ReplaceServerSwitchingRuleBadRequest with default headers values
func NewReplaceServerSwitchingRuleBadRequest() *ReplaceServerSwitchingRuleBadRequest {
	return &ReplaceServerSwitchingRuleBadRequest{}
}

/*ReplaceServerSwitchingRuleBadRequest handles this case with default header values.

Bad request
*/
type ReplaceServerSwitchingRuleBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceServerSwitchingRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/server_switching_rules/{index}][%d] replaceServerSwitchingRuleBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceServerSwitchingRuleBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceServerSwitchingRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceServerSwitchingRuleNotFound creates a ReplaceServerSwitchingRuleNotFound with default headers values
func NewReplaceServerSwitchingRuleNotFound() *ReplaceServerSwitchingRuleNotFound {
	return &ReplaceServerSwitchingRuleNotFound{}
}

/*ReplaceServerSwitchingRuleNotFound handles this case with default header values.

The specified resource was not found
*/
type ReplaceServerSwitchingRuleNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceServerSwitchingRuleNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/server_switching_rules/{index}][%d] replaceServerSwitchingRuleNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceServerSwitchingRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceServerSwitchingRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceServerSwitchingRuleDefault creates a ReplaceServerSwitchingRuleDefault with default headers values
func NewReplaceServerSwitchingRuleDefault(code int) *ReplaceServerSwitchingRuleDefault {
	return &ReplaceServerSwitchingRuleDefault{
		_statusCode: code,
	}
}

/*ReplaceServerSwitchingRuleDefault handles this case with default header values.

General Error
*/
type ReplaceServerSwitchingRuleDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the replace server switching rule default response
func (o *ReplaceServerSwitchingRuleDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceServerSwitchingRuleDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/server_switching_rules/{index}][%d] replaceServerSwitchingRule default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceServerSwitchingRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceServerSwitchingRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
