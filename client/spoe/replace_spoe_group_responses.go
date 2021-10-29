// Code generated by go-swagger; DO NOT EDIT.

package spoe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/models"
)

// ReplaceSpoeGroupReader is a Reader for the ReplaceSpoeGroup structure.
type ReplaceSpoeGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceSpoeGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceSpoeGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceSpoeGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceSpoeGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceSpoeGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceSpoeGroupOK creates a ReplaceSpoeGroupOK with default headers values
func NewReplaceSpoeGroupOK() *ReplaceSpoeGroupOK {
	return &ReplaceSpoeGroupOK{}
}

/* ReplaceSpoeGroupOK describes a response with status code 200, with default header values.

Spoe groups replaced
*/
type ReplaceSpoeGroupOK struct {
	Payload *models.SpoeGroup
}

func (o *ReplaceSpoeGroupOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe/spoe_groups/{name}][%d] replaceSpoeGroupOK  %+v", 200, o.Payload)
}
func (o *ReplaceSpoeGroupOK) GetPayload() *models.SpoeGroup {
	return o.Payload
}

func (o *ReplaceSpoeGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SpoeGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceSpoeGroupBadRequest creates a ReplaceSpoeGroupBadRequest with default headers values
func NewReplaceSpoeGroupBadRequest() *ReplaceSpoeGroupBadRequest {
	return &ReplaceSpoeGroupBadRequest{}
}

/* ReplaceSpoeGroupBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceSpoeGroupBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceSpoeGroupBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe/spoe_groups/{name}][%d] replaceSpoeGroupBadRequest  %+v", 400, o.Payload)
}
func (o *ReplaceSpoeGroupBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceSpoeGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceSpoeGroupNotFound creates a ReplaceSpoeGroupNotFound with default headers values
func NewReplaceSpoeGroupNotFound() *ReplaceSpoeGroupNotFound {
	return &ReplaceSpoeGroupNotFound{}
}

/* ReplaceSpoeGroupNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type ReplaceSpoeGroupNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceSpoeGroupNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe/spoe_groups/{name}][%d] replaceSpoeGroupNotFound  %+v", 404, o.Payload)
}
func (o *ReplaceSpoeGroupNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceSpoeGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceSpoeGroupDefault creates a ReplaceSpoeGroupDefault with default headers values
func NewReplaceSpoeGroupDefault(code int) *ReplaceSpoeGroupDefault {
	return &ReplaceSpoeGroupDefault{
		_statusCode: code,
	}
}

/* ReplaceSpoeGroupDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceSpoeGroupDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the replace spoe group default response
func (o *ReplaceSpoeGroupDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceSpoeGroupDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe/spoe_groups/{name}][%d] replaceSpoeGroup default  %+v", o._statusCode, o.Payload)
}
func (o *ReplaceSpoeGroupDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceSpoeGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
