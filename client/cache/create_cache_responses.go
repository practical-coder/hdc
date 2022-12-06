// Code generated by go-swagger; DO NOT EDIT.

package cache

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// CreateCacheReader is a Reader for the CreateCache structure.
type CreateCacheReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCacheReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateCacheCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateCacheAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateCacheBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateCacheConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateCacheDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateCacheCreated creates a CreateCacheCreated with default headers values
func NewCreateCacheCreated() *CreateCacheCreated {
	return &CreateCacheCreated{}
}

/*
CreateCacheCreated describes a response with status code 201, with default header values.

Cache created
*/
type CreateCacheCreated struct {
	Payload *models.Cache
}

// IsSuccess returns true when this create cache created response has a 2xx status code
func (o *CreateCacheCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create cache created response has a 3xx status code
func (o *CreateCacheCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cache created response has a 4xx status code
func (o *CreateCacheCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create cache created response has a 5xx status code
func (o *CreateCacheCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create cache created response a status code equal to that given
func (o *CreateCacheCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateCacheCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/caches][%d] createCacheCreated  %+v", 201, o.Payload)
}

func (o *CreateCacheCreated) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/caches][%d] createCacheCreated  %+v", 201, o.Payload)
}

func (o *CreateCacheCreated) GetPayload() *models.Cache {
	return o.Payload
}

func (o *CreateCacheCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cache)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCacheAccepted creates a CreateCacheAccepted with default headers values
func NewCreateCacheAccepted() *CreateCacheAccepted {
	return &CreateCacheAccepted{}
}

/*
CreateCacheAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateCacheAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.Cache
}

// IsSuccess returns true when this create cache accepted response has a 2xx status code
func (o *CreateCacheAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create cache accepted response has a 3xx status code
func (o *CreateCacheAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cache accepted response has a 4xx status code
func (o *CreateCacheAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this create cache accepted response has a 5xx status code
func (o *CreateCacheAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this create cache accepted response a status code equal to that given
func (o *CreateCacheAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *CreateCacheAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/caches][%d] createCacheAccepted  %+v", 202, o.Payload)
}

func (o *CreateCacheAccepted) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/caches][%d] createCacheAccepted  %+v", 202, o.Payload)
}

func (o *CreateCacheAccepted) GetPayload() *models.Cache {
	return o.Payload
}

func (o *CreateCacheAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.Cache)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCacheBadRequest creates a CreateCacheBadRequest with default headers values
func NewCreateCacheBadRequest() *CreateCacheBadRequest {
	return &CreateCacheBadRequest{}
}

/*
CreateCacheBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateCacheBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create cache bad request response has a 2xx status code
func (o *CreateCacheBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create cache bad request response has a 3xx status code
func (o *CreateCacheBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cache bad request response has a 4xx status code
func (o *CreateCacheBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create cache bad request response has a 5xx status code
func (o *CreateCacheBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create cache bad request response a status code equal to that given
func (o *CreateCacheBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateCacheBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/caches][%d] createCacheBadRequest  %+v", 400, o.Payload)
}

func (o *CreateCacheBadRequest) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/caches][%d] createCacheBadRequest  %+v", 400, o.Payload)
}

func (o *CreateCacheBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateCacheBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateCacheConflict creates a CreateCacheConflict with default headers values
func NewCreateCacheConflict() *CreateCacheConflict {
	return &CreateCacheConflict{}
}

/*
CreateCacheConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateCacheConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create cache conflict response has a 2xx status code
func (o *CreateCacheConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create cache conflict response has a 3xx status code
func (o *CreateCacheConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cache conflict response has a 4xx status code
func (o *CreateCacheConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create cache conflict response has a 5xx status code
func (o *CreateCacheConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create cache conflict response a status code equal to that given
func (o *CreateCacheConflict) IsCode(code int) bool {
	return code == 409
}

func (o *CreateCacheConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/caches][%d] createCacheConflict  %+v", 409, o.Payload)
}

func (o *CreateCacheConflict) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/caches][%d] createCacheConflict  %+v", 409, o.Payload)
}

func (o *CreateCacheConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateCacheConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateCacheDefault creates a CreateCacheDefault with default headers values
func NewCreateCacheDefault(code int) *CreateCacheDefault {
	return &CreateCacheDefault{
		_statusCode: code,
	}
}

/*
CreateCacheDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateCacheDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create cache default response
func (o *CreateCacheDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create cache default response has a 2xx status code
func (o *CreateCacheDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create cache default response has a 3xx status code
func (o *CreateCacheDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create cache default response has a 4xx status code
func (o *CreateCacheDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create cache default response has a 5xx status code
func (o *CreateCacheDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create cache default response a status code equal to that given
func (o *CreateCacheDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateCacheDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/caches][%d] createCache default  %+v", o._statusCode, o.Payload)
}

func (o *CreateCacheDefault) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/caches][%d] createCache default  %+v", o._statusCode, o.Payload)
}

func (o *CreateCacheDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateCacheDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
