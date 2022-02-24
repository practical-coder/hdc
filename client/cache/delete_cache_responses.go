// Code generated by go-swagger; DO NOT EDIT.

package cache

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v3/models"
)

// DeleteCacheReader is a Reader for the DeleteCache structure.
type DeleteCacheReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCacheReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteCacheAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteCacheNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteCacheNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteCacheDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCacheAccepted creates a DeleteCacheAccepted with default headers values
func NewDeleteCacheAccepted() *DeleteCacheAccepted {
	return &DeleteCacheAccepted{}
}

/*DeleteCacheAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type DeleteCacheAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteCacheAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/caches/{name}][%d] deleteCacheAccepted ", 202)
}

func (o *DeleteCacheAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	return nil
}

// NewDeleteCacheNoContent creates a DeleteCacheNoContent with default headers values
func NewDeleteCacheNoContent() *DeleteCacheNoContent {
	return &DeleteCacheNoContent{}
}

/*DeleteCacheNoContent handles this case with default header values.

Cache deleted
*/
type DeleteCacheNoContent struct {
}

func (o *DeleteCacheNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/caches/{name}][%d] deleteCacheNoContent ", 204)
}

func (o *DeleteCacheNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCacheNotFound creates a DeleteCacheNotFound with default headers values
func NewDeleteCacheNotFound() *DeleteCacheNotFound {
	return &DeleteCacheNotFound{}
}

/*DeleteCacheNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteCacheNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteCacheNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/caches/{name}][%d] deleteCacheNotFound  %+v", 404, o.Payload)
}

func (o *DeleteCacheNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteCacheNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCacheDefault creates a DeleteCacheDefault with default headers values
func NewDeleteCacheDefault(code int) *DeleteCacheDefault {
	return &DeleteCacheDefault{
		_statusCode: code,
	}
}

/*DeleteCacheDefault handles this case with default header values.

General Error
*/
type DeleteCacheDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete cache default response
func (o *DeleteCacheDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCacheDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/caches/{name}][%d] deleteCache default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCacheDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteCacheDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
