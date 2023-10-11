// Code generated by go-swagger; DO NOT EDIT.

package global

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v5/models"
)

// ReplaceGlobalReader is a Reader for the ReplaceGlobal structure.
type ReplaceGlobalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceGlobalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceGlobalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceGlobalAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceGlobalBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceGlobalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceGlobalOK creates a ReplaceGlobalOK with default headers values
func NewReplaceGlobalOK() *ReplaceGlobalOK {
	return &ReplaceGlobalOK{}
}

/*
ReplaceGlobalOK describes a response with status code 200, with default header values.

Global replaced
*/
type ReplaceGlobalOK struct {
	Payload *models.Global
}

// IsSuccess returns true when this replace global o k response has a 2xx status code
func (o *ReplaceGlobalOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this replace global o k response has a 3xx status code
func (o *ReplaceGlobalOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace global o k response has a 4xx status code
func (o *ReplaceGlobalOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this replace global o k response has a 5xx status code
func (o *ReplaceGlobalOK) IsServerError() bool {
	return false
}

// IsCode returns true when this replace global o k response a status code equal to that given
func (o *ReplaceGlobalOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the replace global o k response
func (o *ReplaceGlobalOK) Code() int {
	return 200
}

func (o *ReplaceGlobalOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/global][%d] replaceGlobalOK  %+v", 200, o.Payload)
}

func (o *ReplaceGlobalOK) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/global][%d] replaceGlobalOK  %+v", 200, o.Payload)
}

func (o *ReplaceGlobalOK) GetPayload() *models.Global {
	return o.Payload
}

func (o *ReplaceGlobalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Global)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceGlobalAccepted creates a ReplaceGlobalAccepted with default headers values
func NewReplaceGlobalAccepted() *ReplaceGlobalAccepted {
	return &ReplaceGlobalAccepted{}
}

/*
ReplaceGlobalAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type ReplaceGlobalAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.Global
}

// IsSuccess returns true when this replace global accepted response has a 2xx status code
func (o *ReplaceGlobalAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this replace global accepted response has a 3xx status code
func (o *ReplaceGlobalAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace global accepted response has a 4xx status code
func (o *ReplaceGlobalAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this replace global accepted response has a 5xx status code
func (o *ReplaceGlobalAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this replace global accepted response a status code equal to that given
func (o *ReplaceGlobalAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the replace global accepted response
func (o *ReplaceGlobalAccepted) Code() int {
	return 202
}

func (o *ReplaceGlobalAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/global][%d] replaceGlobalAccepted  %+v", 202, o.Payload)
}

func (o *ReplaceGlobalAccepted) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/global][%d] replaceGlobalAccepted  %+v", 202, o.Payload)
}

func (o *ReplaceGlobalAccepted) GetPayload() *models.Global {
	return o.Payload
}

func (o *ReplaceGlobalAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.Global)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceGlobalBadRequest creates a ReplaceGlobalBadRequest with default headers values
func NewReplaceGlobalBadRequest() *ReplaceGlobalBadRequest {
	return &ReplaceGlobalBadRequest{}
}

/*
ReplaceGlobalBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceGlobalBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this replace global bad request response has a 2xx status code
func (o *ReplaceGlobalBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this replace global bad request response has a 3xx status code
func (o *ReplaceGlobalBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace global bad request response has a 4xx status code
func (o *ReplaceGlobalBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this replace global bad request response has a 5xx status code
func (o *ReplaceGlobalBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this replace global bad request response a status code equal to that given
func (o *ReplaceGlobalBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the replace global bad request response
func (o *ReplaceGlobalBadRequest) Code() int {
	return 400
}

func (o *ReplaceGlobalBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/global][%d] replaceGlobalBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceGlobalBadRequest) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/global][%d] replaceGlobalBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceGlobalBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceGlobalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceGlobalDefault creates a ReplaceGlobalDefault with default headers values
func NewReplaceGlobalDefault(code int) *ReplaceGlobalDefault {
	return &ReplaceGlobalDefault{
		_statusCode: code,
	}
}

/*
ReplaceGlobalDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceGlobalDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this replace global default response has a 2xx status code
func (o *ReplaceGlobalDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this replace global default response has a 3xx status code
func (o *ReplaceGlobalDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this replace global default response has a 4xx status code
func (o *ReplaceGlobalDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this replace global default response has a 5xx status code
func (o *ReplaceGlobalDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this replace global default response a status code equal to that given
func (o *ReplaceGlobalDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the replace global default response
func (o *ReplaceGlobalDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceGlobalDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/global][%d] replaceGlobal default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceGlobalDefault) String() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/global][%d] replaceGlobal default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceGlobalDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceGlobalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
