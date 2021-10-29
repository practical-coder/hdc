// Code generated by go-swagger; DO NOT EDIT.

package filter

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

/* CreateFilterCreated describes a response with status code 201, with default header values.

Filter created
*/
type CreateFilterCreated struct {
	Payload *models.Filter
}

func (o *CreateFilterCreated) Error() string {
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

/* CreateFilterAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateFilterAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.Filter
}

func (o *CreateFilterAccepted) Error() string {
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
	var (
		// initialize headers with default values
		configurationVersionDefault = int64(0)
	)

	return &CreateFilterBadRequest{

		ConfigurationVersion: configurationVersionDefault,
	}
}

/* CreateFilterBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateFilterBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *CreateFilterBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/filters][%d] createFilterBadRequest  %+v", 400, o.Payload)
}
func (o *CreateFilterBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateFilterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateFilterConflict creates a CreateFilterConflict with default headers values
func NewCreateFilterConflict() *CreateFilterConflict {
	var (
		// initialize headers with default values
		configurationVersionDefault = int64(0)
	)

	return &CreateFilterConflict{

		ConfigurationVersion: configurationVersionDefault,
	}
}

/* CreateFilterConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateFilterConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *CreateFilterConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/filters][%d] createFilterConflict  %+v", 409, o.Payload)
}
func (o *CreateFilterConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateFilterConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateFilterDefault creates a CreateFilterDefault with default headers values
func NewCreateFilterDefault(code int) *CreateFilterDefault {
	var (
		// initialize headers with default values
		configurationVersionDefault = int64(0)
	)

	return &CreateFilterDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

/* CreateFilterDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateFilterDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the create filter default response
func (o *CreateFilterDefault) Code() int {
	return o._statusCode
}

func (o *CreateFilterDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/filters][%d] createFilter default  %+v", o._statusCode, o.Payload)
}
func (o *CreateFilterDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateFilterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
