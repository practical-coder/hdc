// Code generated by go-swagger; DO NOT EDIT.

package frontend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v5/models"
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

/*
ReplaceFrontendOK describes a response with status code 200, with default header values.

Frontend replaced
*/
type ReplaceFrontendOK struct {
	Payload *models.Frontend
}

// IsSuccess returns true when this replace frontend o k response has a 2xx status code
func (o *ReplaceFrontendOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this replace frontend o k response has a 3xx status code
func (o *ReplaceFrontendOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace frontend o k response has a 4xx status code
func (o *ReplaceFrontendOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this replace frontend o k response has a 5xx status code
func (o *ReplaceFrontendOK) IsServerError() bool {
	return false
}

// IsCode returns true when this replace frontend o k response a status code equal to that given
func (o *ReplaceFrontendOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the replace frontend o k response
func (o *ReplaceFrontendOK) Code() int {
	return 200
}

func (o *ReplaceFrontendOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/frontends/{name}][%d] replaceFrontendOK  %+v", 200, o.Payload)
}

func (o *ReplaceFrontendOK) String() string {
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

/*
ReplaceFrontendAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type ReplaceFrontendAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.Frontend
}

// IsSuccess returns true when this replace frontend accepted response has a 2xx status code
func (o *ReplaceFrontendAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this replace frontend accepted response has a 3xx status code
func (o *ReplaceFrontendAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace frontend accepted response has a 4xx status code
func (o *ReplaceFrontendAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this replace frontend accepted response has a 5xx status code
func (o *ReplaceFrontendAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this replace frontend accepted response a status code equal to that given
func (o *ReplaceFrontendAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the replace frontend accepted response
func (o *ReplaceFrontendAccepted) Code() int {
	return 202
}

func (o *ReplaceFrontendAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/frontends/{name}][%d] replaceFrontendAccepted  %+v", 202, o.Payload)
}

func (o *ReplaceFrontendAccepted) String() string {
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
	return &ReplaceFrontendBadRequest{}
}

/*
ReplaceFrontendBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceFrontendBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this replace frontend bad request response has a 2xx status code
func (o *ReplaceFrontendBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this replace frontend bad request response has a 3xx status code
func (o *ReplaceFrontendBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace frontend bad request response has a 4xx status code
func (o *ReplaceFrontendBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this replace frontend bad request response has a 5xx status code
func (o *ReplaceFrontendBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this replace frontend bad request response a status code equal to that given
func (o *ReplaceFrontendBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the replace frontend bad request response
func (o *ReplaceFrontendBadRequest) Code() int {
	return 400
}

func (o *ReplaceFrontendBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/frontends/{name}][%d] replaceFrontendBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceFrontendBadRequest) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/frontends/{name}][%d] replaceFrontendBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceFrontendBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceFrontendBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceFrontendNotFound creates a ReplaceFrontendNotFound with default headers values
func NewReplaceFrontendNotFound() *ReplaceFrontendNotFound {
	return &ReplaceFrontendNotFound{}
}

/*
ReplaceFrontendNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type ReplaceFrontendNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this replace frontend not found response has a 2xx status code
func (o *ReplaceFrontendNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this replace frontend not found response has a 3xx status code
func (o *ReplaceFrontendNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace frontend not found response has a 4xx status code
func (o *ReplaceFrontendNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this replace frontend not found response has a 5xx status code
func (o *ReplaceFrontendNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this replace frontend not found response a status code equal to that given
func (o *ReplaceFrontendNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the replace frontend not found response
func (o *ReplaceFrontendNotFound) Code() int {
	return 404
}

func (o *ReplaceFrontendNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/frontends/{name}][%d] replaceFrontendNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceFrontendNotFound) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/frontends/{name}][%d] replaceFrontendNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceFrontendNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceFrontendNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceFrontendDefault creates a ReplaceFrontendDefault with default headers values
func NewReplaceFrontendDefault(code int) *ReplaceFrontendDefault {
	return &ReplaceFrontendDefault{
		_statusCode: code,
	}
}

/*
ReplaceFrontendDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceFrontendDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this replace frontend default response has a 2xx status code
func (o *ReplaceFrontendDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this replace frontend default response has a 3xx status code
func (o *ReplaceFrontendDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this replace frontend default response has a 4xx status code
func (o *ReplaceFrontendDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this replace frontend default response has a 5xx status code
func (o *ReplaceFrontendDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this replace frontend default response a status code equal to that given
func (o *ReplaceFrontendDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the replace frontend default response
func (o *ReplaceFrontendDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceFrontendDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/frontends/{name}][%d] replaceFrontend default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceFrontendDefault) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/frontends/{name}][%d] replaceFrontend default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceFrontendDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceFrontendDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
