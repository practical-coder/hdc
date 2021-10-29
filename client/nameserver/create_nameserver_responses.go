// Code generated by go-swagger; DO NOT EDIT.

package nameserver

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

// CreateNameserverReader is a Reader for the CreateNameserver structure.
type CreateNameserverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNameserverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNameserverCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateNameserverAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateNameserverBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateNameserverConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateNameserverDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateNameserverCreated creates a CreateNameserverCreated with default headers values
func NewCreateNameserverCreated() *CreateNameserverCreated {
	return &CreateNameserverCreated{}
}

/* CreateNameserverCreated describes a response with status code 201, with default header values.

Nameserver created
*/
type CreateNameserverCreated struct {
	Payload *models.Nameserver
}

func (o *CreateNameserverCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/nameservers][%d] createNameserverCreated  %+v", 201, o.Payload)
}
func (o *CreateNameserverCreated) GetPayload() *models.Nameserver {
	return o.Payload
}

func (o *CreateNameserverCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Nameserver)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNameserverAccepted creates a CreateNameserverAccepted with default headers values
func NewCreateNameserverAccepted() *CreateNameserverAccepted {
	return &CreateNameserverAccepted{}
}

/* CreateNameserverAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateNameserverAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.Nameserver
}

func (o *CreateNameserverAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/nameservers][%d] createNameserverAccepted  %+v", 202, o.Payload)
}
func (o *CreateNameserverAccepted) GetPayload() *models.Nameserver {
	return o.Payload
}

func (o *CreateNameserverAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.Nameserver)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNameserverBadRequest creates a CreateNameserverBadRequest with default headers values
func NewCreateNameserverBadRequest() *CreateNameserverBadRequest {
	var (
		// initialize headers with default values
		configurationVersionDefault = int64(0)
	)

	return &CreateNameserverBadRequest{

		ConfigurationVersion: configurationVersionDefault,
	}
}

/* CreateNameserverBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateNameserverBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *CreateNameserverBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/nameservers][%d] createNameserverBadRequest  %+v", 400, o.Payload)
}
func (o *CreateNameserverBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateNameserverBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateNameserverConflict creates a CreateNameserverConflict with default headers values
func NewCreateNameserverConflict() *CreateNameserverConflict {
	var (
		// initialize headers with default values
		configurationVersionDefault = int64(0)
	)

	return &CreateNameserverConflict{

		ConfigurationVersion: configurationVersionDefault,
	}
}

/* CreateNameserverConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateNameserverConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *CreateNameserverConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/nameservers][%d] createNameserverConflict  %+v", 409, o.Payload)
}
func (o *CreateNameserverConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateNameserverConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateNameserverDefault creates a CreateNameserverDefault with default headers values
func NewCreateNameserverDefault(code int) *CreateNameserverDefault {
	var (
		// initialize headers with default values
		configurationVersionDefault = int64(0)
	)

	return &CreateNameserverDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

/* CreateNameserverDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateNameserverDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the create nameserver default response
func (o *CreateNameserverDefault) Code() int {
	return o._statusCode
}

func (o *CreateNameserverDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/nameservers][%d] createNameserver default  %+v", o._statusCode, o.Payload)
}
func (o *CreateNameserverDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateNameserverDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
