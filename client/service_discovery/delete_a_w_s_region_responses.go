// Code generated by go-swagger; DO NOT EDIT.

package service_discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v3/models"
)

// DeleteAWSRegionReader is a Reader for the DeleteAWSRegion structure.
type DeleteAWSRegionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAWSRegionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteAWSRegionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteAWSRegionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteAWSRegionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAWSRegionNoContent creates a DeleteAWSRegionNoContent with default headers values
func NewDeleteAWSRegionNoContent() *DeleteAWSRegionNoContent {
	return &DeleteAWSRegionNoContent{}
}

/*DeleteAWSRegionNoContent handles this case with default header values.

Resource deleted
*/
type DeleteAWSRegionNoContent struct {
}

func (o *DeleteAWSRegionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /service_discovery/aws/{id}][%d] deleteAWSRegionNoContent ", 204)
}

func (o *DeleteAWSRegionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAWSRegionNotFound creates a DeleteAWSRegionNotFound with default headers values
func NewDeleteAWSRegionNotFound() *DeleteAWSRegionNotFound {
	return &DeleteAWSRegionNotFound{}
}

/*DeleteAWSRegionNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteAWSRegionNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteAWSRegionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /service_discovery/aws/{id}][%d] deleteAWSRegionNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAWSRegionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteAWSRegionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAWSRegionDefault creates a DeleteAWSRegionDefault with default headers values
func NewDeleteAWSRegionDefault(code int) *DeleteAWSRegionDefault {
	return &DeleteAWSRegionDefault{
		_statusCode: code,
	}
}

/*DeleteAWSRegionDefault handles this case with default header values.

General Error
*/
type DeleteAWSRegionDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete a w s region default response
func (o *DeleteAWSRegionDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAWSRegionDefault) Error() string {
	return fmt.Sprintf("[DELETE /service_discovery/aws/{id}][%d] deleteAWSRegion default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAWSRegionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteAWSRegionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
