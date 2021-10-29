// Code generated by go-swagger; DO NOT EDIT.

package service_discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/practical-coder/hdc/models"
)

// CreateAWSRegionReader is a Reader for the CreateAWSRegion structure.
type CreateAWSRegionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAWSRegionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAWSRegionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAWSRegionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateAWSRegionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateAWSRegionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateAWSRegionCreated creates a CreateAWSRegionCreated with default headers values
func NewCreateAWSRegionCreated() *CreateAWSRegionCreated {
	return &CreateAWSRegionCreated{}
}

/* CreateAWSRegionCreated describes a response with status code 201, with default header values.

Resource created
*/
type CreateAWSRegionCreated struct {
	Payload *models.AwsRegion
}

func (o *CreateAWSRegionCreated) Error() string {
	return fmt.Sprintf("[POST /service_discovery/aws][%d] createAWSRegionCreated  %+v", 201, o.Payload)
}
func (o *CreateAWSRegionCreated) GetPayload() *models.AwsRegion {
	return o.Payload
}

func (o *CreateAWSRegionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AwsRegion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAWSRegionBadRequest creates a CreateAWSRegionBadRequest with default headers values
func NewCreateAWSRegionBadRequest() *CreateAWSRegionBadRequest {
	return &CreateAWSRegionBadRequest{}
}

/* CreateAWSRegionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateAWSRegionBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateAWSRegionBadRequest) Error() string {
	return fmt.Sprintf("[POST /service_discovery/aws][%d] createAWSRegionBadRequest  %+v", 400, o.Payload)
}
func (o *CreateAWSRegionBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateAWSRegionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateAWSRegionConflict creates a CreateAWSRegionConflict with default headers values
func NewCreateAWSRegionConflict() *CreateAWSRegionConflict {
	return &CreateAWSRegionConflict{}
}

/* CreateAWSRegionConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateAWSRegionConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateAWSRegionConflict) Error() string {
	return fmt.Sprintf("[POST /service_discovery/aws][%d] createAWSRegionConflict  %+v", 409, o.Payload)
}
func (o *CreateAWSRegionConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateAWSRegionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateAWSRegionDefault creates a CreateAWSRegionDefault with default headers values
func NewCreateAWSRegionDefault(code int) *CreateAWSRegionDefault {
	return &CreateAWSRegionDefault{
		_statusCode: code,
	}
}

/* CreateAWSRegionDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateAWSRegionDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create a w s region default response
func (o *CreateAWSRegionDefault) Code() int {
	return o._statusCode
}

func (o *CreateAWSRegionDefault) Error() string {
	return fmt.Sprintf("[POST /service_discovery/aws][%d] createAWSRegion default  %+v", o._statusCode, o.Payload)
}
func (o *CreateAWSRegionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateAWSRegionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
