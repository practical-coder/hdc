// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// DeleteStorageMapReader is a Reader for the DeleteStorageMap structure.
type DeleteStorageMapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteStorageMapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteStorageMapNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteStorageMapNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteStorageMapDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteStorageMapNoContent creates a DeleteStorageMapNoContent with default headers values
func NewDeleteStorageMapNoContent() *DeleteStorageMapNoContent {
	return &DeleteStorageMapNoContent{}
}

/* DeleteStorageMapNoContent describes a response with status code 204, with default header values.

Map file deleted
*/
type DeleteStorageMapNoContent struct {
}

func (o *DeleteStorageMapNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/storage/maps/{name}][%d] deleteStorageMapNoContent ", 204)
}

func (o *DeleteStorageMapNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteStorageMapNotFound creates a DeleteStorageMapNotFound with default headers values
func NewDeleteStorageMapNotFound() *DeleteStorageMapNotFound {
	return &DeleteStorageMapNotFound{}
}

/* DeleteStorageMapNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteStorageMapNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteStorageMapNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/storage/maps/{name}][%d] deleteStorageMapNotFound  %+v", 404, o.Payload)
}
func (o *DeleteStorageMapNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteStorageMapNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteStorageMapDefault creates a DeleteStorageMapDefault with default headers values
func NewDeleteStorageMapDefault(code int) *DeleteStorageMapDefault {
	return &DeleteStorageMapDefault{
		_statusCode: code,
	}
}

/* DeleteStorageMapDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteStorageMapDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete storage map default response
func (o *DeleteStorageMapDefault) Code() int {
	return o._statusCode
}

func (o *DeleteStorageMapDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/storage/maps/{name}][%d] deleteStorageMap default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteStorageMapDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteStorageMapDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
