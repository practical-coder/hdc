// Code generated by go-swagger; DO NOT EDIT.

package frontend

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

// ReplaceFrontendReader is a Reader for the ReplaceFrontend structure.
type ReplaceFrontendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceFrontendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceFrontendOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceFrontendAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceFrontendBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceFrontendNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceFrontendDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceFrontendOK creates a ReplaceFrontendOK with default headers values
func NewReplaceFrontendOK() *ReplaceFrontendOK {
	return &ReplaceFrontendOK{}
}

/* ReplaceFrontendOK describes a response with status code 200, with default header values.

Frontend replaced
*/
type ReplaceFrontendOK struct {
	Payload *models.Frontend
}

func (o *ReplaceFrontendOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/frontends/{name}][%d] replaceFrontendOK  %+v", 200, o.Payload)
}
func (o *ReplaceFrontendOK) GetPayload() *models.Frontend {
	return o.Payload
}

func (o *ReplaceFrontendOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Frontend)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceFrontendAccepted creates a ReplaceFrontendAccepted with default headers values
func NewReplaceFrontendAccepted() *ReplaceFrontendAccepted {
	return &ReplaceFrontendAccepted{}
}

/* ReplaceFrontendAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type ReplaceFrontendAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.Frontend
}

func (o *ReplaceFrontendAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/frontends/{name}][%d] replaceFrontendAccepted  %+v", 202, o.Payload)
}
func (o *ReplaceFrontendAccepted) GetPayload() *models.Frontend {
	return o.Payload
}

func (o *ReplaceFrontendAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.Frontend)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceFrontendBadRequest creates a ReplaceFrontendBadRequest with default headers values
func NewReplaceFrontendBadRequest() *ReplaceFrontendBadRequest {
	var (
		// initialize headers with default values
		configurationVersionDefault = int64(0)
	)

	return &ReplaceFrontendBadRequest{

		ConfigurationVersion: configurationVersionDefault,
	}
}

/* ReplaceFrontendBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceFrontendBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *ReplaceFrontendBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/frontends/{name}][%d] replaceFrontendBadRequest  %+v", 400, o.Payload)
}
func (o *ReplaceFrontendBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceFrontendBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceFrontendNotFound creates a ReplaceFrontendNotFound with default headers values
func NewReplaceFrontendNotFound() *ReplaceFrontendNotFound {
	var (
		// initialize headers with default values
		configurationVersionDefault = int64(0)
	)

	return &ReplaceFrontendNotFound{

		ConfigurationVersion: configurationVersionDefault,
	}
}

/* ReplaceFrontendNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type ReplaceFrontendNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *ReplaceFrontendNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/frontends/{name}][%d] replaceFrontendNotFound  %+v", 404, o.Payload)
}
func (o *ReplaceFrontendNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceFrontendNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceFrontendDefault creates a ReplaceFrontendDefault with default headers values
func NewReplaceFrontendDefault(code int) *ReplaceFrontendDefault {
	var (
		// initialize headers with default values
		configurationVersionDefault = int64(0)
	)

	return &ReplaceFrontendDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

/* ReplaceFrontendDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceFrontendDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the replace frontend default response
func (o *ReplaceFrontendDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceFrontendDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/frontends/{name}][%d] replaceFrontend default  %+v", o._statusCode, o.Payload)
}
func (o *ReplaceFrontendDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceFrontendDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
