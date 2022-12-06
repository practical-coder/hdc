// Code generated by go-swagger; DO NOT EDIT.

package filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// CreateFilterReader is a Reader for the CreateFilter structure.
type CreateFilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateFilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateFilterCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateFilterAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateFilterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateFilterConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateFilterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateFilterCreated creates a CreateFilterCreated with default headers values
func NewCreateFilterCreated() *CreateFilterCreated {
	return &CreateFilterCreated{}
}

/*
CreateFilterCreated describes a response with status code 201, with default header values.

Filter created
*/
type CreateFilterCreated struct {
	Payload *models.Filter
}

// IsSuccess returns true when this create filter created response has a 2xx status code
func (o *CreateFilterCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create filter created response has a 3xx status code
func (o *CreateFilterCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create filter created response has a 4xx status code
func (o *CreateFilterCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create filter created response has a 5xx status code
func (o *CreateFilterCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create filter created response a status code equal to that given
func (o *CreateFilterCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateFilterCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/filters][%d] createFilterCreated  %+v", 201, o.Payload)
}

func (o *CreateFilterCreated) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/filters][%d] createFilterCreated  %+v", 201, o.Payload)
}

func (o *CreateFilterCreated) GetPayload() *models.Filter {
	return o.Payload
}

func (o *CreateFilterCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Filter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFilterAccepted creates a CreateFilterAccepted with default headers values
func NewCreateFilterAccepted() *CreateFilterAccepted {
	return &CreateFilterAccepted{}
}

/*
CreateFilterAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateFilterAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.Filter
}

// IsSuccess returns true when this create filter accepted response has a 2xx status code
func (o *CreateFilterAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create filter accepted response has a 3xx status code
func (o *CreateFilterAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create filter accepted response has a 4xx status code
func (o *CreateFilterAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this create filter accepted response has a 5xx status code
func (o *CreateFilterAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this create filter accepted response a status code equal to that given
func (o *CreateFilterAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *CreateFilterAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/filters][%d] createFilterAccepted  %+v", 202, o.Payload)
}

func (o *CreateFilterAccepted) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/filters][%d] createFilterAccepted  %+v", 202, o.Payload)
}

func (o *CreateFilterAccepted) GetPayload() *models.Filter {
	return o.Payload
}

func (o *CreateFilterAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.Filter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFilterBadRequest creates a CreateFilterBadRequest with default headers values
func NewCreateFilterBadRequest() *CreateFilterBadRequest {
	return &CreateFilterBadRequest{}
}

/*
CreateFilterBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateFilterBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create filter bad request response has a 2xx status code
func (o *CreateFilterBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create filter bad request response has a 3xx status code
func (o *CreateFilterBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create filter bad request response has a 4xx status code
func (o *CreateFilterBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create filter bad request response has a 5xx status code
func (o *CreateFilterBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create filter bad request response a status code equal to that given
func (o *CreateFilterBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateFilterBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/filters][%d] createFilterBadRequest  %+v", 400, o.Payload)
}

func (o *CreateFilterBadRequest) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/filters][%d] createFilterBadRequest  %+v", 400, o.Payload)
}

func (o *CreateFilterBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateFilterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateFilterConflict creates a CreateFilterConflict with default headers values
func NewCreateFilterConflict() *CreateFilterConflict {
	return &CreateFilterConflict{}
}

/*
CreateFilterConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateFilterConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this create filter conflict response has a 2xx status code
func (o *CreateFilterConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create filter conflict response has a 3xx status code
func (o *CreateFilterConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create filter conflict response has a 4xx status code
func (o *CreateFilterConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create filter conflict response has a 5xx status code
func (o *CreateFilterConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create filter conflict response a status code equal to that given
func (o *CreateFilterConflict) IsCode(code int) bool {
	return code == 409
}

func (o *CreateFilterConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/filters][%d] createFilterConflict  %+v", 409, o.Payload)
}

func (o *CreateFilterConflict) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/filters][%d] createFilterConflict  %+v", 409, o.Payload)
}

func (o *CreateFilterConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateFilterConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateFilterDefault creates a CreateFilterDefault with default headers values
func NewCreateFilterDefault(code int) *CreateFilterDefault {
	return &CreateFilterDefault{
		_statusCode: code,
	}
}

/*
CreateFilterDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateFilterDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create filter default response
func (o *CreateFilterDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create filter default response has a 2xx status code
func (o *CreateFilterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create filter default response has a 3xx status code
func (o *CreateFilterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create filter default response has a 4xx status code
func (o *CreateFilterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create filter default response has a 5xx status code
func (o *CreateFilterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create filter default response a status code equal to that given
func (o *CreateFilterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateFilterDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/filters][%d] createFilter default  %+v", o._statusCode, o.Payload)
}

func (o *CreateFilterDefault) String() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/filters][%d] createFilter default  %+v", o._statusCode, o.Payload)
}

func (o *CreateFilterDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateFilterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
