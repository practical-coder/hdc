// Code generated by go-swagger; DO NOT EDIT.

package stick_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models"
)

// ReplaceStickRuleReader is a Reader for the ReplaceStickRule structure.
type ReplaceStickRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceStickRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceStickRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceStickRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceStickRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceStickRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceStickRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceStickRuleOK creates a ReplaceStickRuleOK with default headers values
func NewReplaceStickRuleOK() *ReplaceStickRuleOK {
	return &ReplaceStickRuleOK{}
}

/* ReplaceStickRuleOK describes a response with status code 200, with default header values.

Stick Rule replaced
*/
type ReplaceStickRuleOK struct {
	Payload *models.StickRule
}

func (o *ReplaceStickRuleOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/stick_rules/{index}][%d] replaceStickRuleOK  %+v", 200, o.Payload)
}
func (o *ReplaceStickRuleOK) GetPayload() *models.StickRule {
	return o.Payload
}

func (o *ReplaceStickRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StickRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceStickRuleAccepted creates a ReplaceStickRuleAccepted with default headers values
func NewReplaceStickRuleAccepted() *ReplaceStickRuleAccepted {
	return &ReplaceStickRuleAccepted{}
}

/* ReplaceStickRuleAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type ReplaceStickRuleAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.StickRule
}

func (o *ReplaceStickRuleAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/stick_rules/{index}][%d] replaceStickRuleAccepted  %+v", 202, o.Payload)
}
func (o *ReplaceStickRuleAccepted) GetPayload() *models.StickRule {
	return o.Payload
}

func (o *ReplaceStickRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.StickRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceStickRuleBadRequest creates a ReplaceStickRuleBadRequest with default headers values
func NewReplaceStickRuleBadRequest() *ReplaceStickRuleBadRequest {
	var (
		// initialize headers with default values
		configurationVersionDefault = int64(0)
	)

	return &ReplaceStickRuleBadRequest{

		ConfigurationVersion: configurationVersionDefault,
	}
}

/* ReplaceStickRuleBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceStickRuleBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *ReplaceStickRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/stick_rules/{index}][%d] replaceStickRuleBadRequest  %+v", 400, o.Payload)
}
func (o *ReplaceStickRuleBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceStickRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		valconfigurationVersion, err := swag.ConvertInt64(hdrConfigurationVersion)
		if err != nil {
			return errors.InvalidType("Configuration-Version", "header", "int64", hdrConfigurationVersion)
		}
		o.ConfigurationVersion = valconfigurationVersion
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceStickRuleNotFound creates a ReplaceStickRuleNotFound with default headers values
func NewReplaceStickRuleNotFound() *ReplaceStickRuleNotFound {
	var (
		// initialize headers with default values
		configurationVersionDefault = int64(0)
	)

	return &ReplaceStickRuleNotFound{

		ConfigurationVersion: configurationVersionDefault,
	}
}

/* ReplaceStickRuleNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type ReplaceStickRuleNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *ReplaceStickRuleNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/stick_rules/{index}][%d] replaceStickRuleNotFound  %+v", 404, o.Payload)
}
func (o *ReplaceStickRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceStickRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		valconfigurationVersion, err := swag.ConvertInt64(hdrConfigurationVersion)
		if err != nil {
			return errors.InvalidType("Configuration-Version", "header", "int64", hdrConfigurationVersion)
		}
		o.ConfigurationVersion = valconfigurationVersion
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceStickRuleDefault creates a ReplaceStickRuleDefault with default headers values
func NewReplaceStickRuleDefault(code int) *ReplaceStickRuleDefault {
	var (
		// initialize headers with default values
		configurationVersionDefault = int64(0)
	)

	return &ReplaceStickRuleDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

/* ReplaceStickRuleDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceStickRuleDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the replace stick rule default response
func (o *ReplaceStickRuleDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceStickRuleDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/stick_rules/{index}][%d] replaceStickRule default  %+v", o._statusCode, o.Payload)
}
func (o *ReplaceStickRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceStickRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		valconfigurationVersion, err := swag.ConvertInt64(hdrConfigurationVersion)
		if err != nil {
			return errors.InvalidType("Configuration-Version", "header", "int64", hdrConfigurationVersion)
		}
		o.ConfigurationVersion = valconfigurationVersion
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
