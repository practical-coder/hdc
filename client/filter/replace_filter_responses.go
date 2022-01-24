// Code generated by go-swagger; DO NOT EDIT.

package filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v2/models"
)

// ReplaceFilterReader is a Reader for the ReplaceFilter structure.
type ReplaceFilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceFilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceFilterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceFilterAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceFilterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceFilterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceFilterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceFilterOK creates a ReplaceFilterOK with default headers values
func NewReplaceFilterOK() *ReplaceFilterOK {
	return &ReplaceFilterOK{}
}

/*ReplaceFilterOK handles this case with default header values.

Filter replaced
*/
type ReplaceFilterOK struct {
	Payload *models.Filter
}

func (o *ReplaceFilterOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/filters/{index}][%d] replaceFilterOK  %+v", 200, o.Payload)
}

func (o *ReplaceFilterOK) GetPayload() *models.Filter {
	return o.Payload
}

func (o *ReplaceFilterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Filter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceFilterAccepted creates a ReplaceFilterAccepted with default headers values
func NewReplaceFilterAccepted() *ReplaceFilterAccepted {
	return &ReplaceFilterAccepted{}
}

/*ReplaceFilterAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type ReplaceFilterAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string

	Payload *models.Filter
}

func (o *ReplaceFilterAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/filters/{index}][%d] replaceFilterAccepted  %+v", 202, o.Payload)
}

func (o *ReplaceFilterAccepted) GetPayload() *models.Filter {
	return o.Payload
}

func (o *ReplaceFilterAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	o.Payload = new(models.Filter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceFilterBadRequest creates a ReplaceFilterBadRequest with default headers values
func NewReplaceFilterBadRequest() *ReplaceFilterBadRequest {
	return &ReplaceFilterBadRequest{}
}

/*ReplaceFilterBadRequest handles this case with default header values.

Bad request
*/
type ReplaceFilterBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceFilterBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/filters/{index}][%d] replaceFilterBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceFilterBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceFilterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceFilterNotFound creates a ReplaceFilterNotFound with default headers values
func NewReplaceFilterNotFound() *ReplaceFilterNotFound {
	return &ReplaceFilterNotFound{}
}

/*ReplaceFilterNotFound handles this case with default header values.

The specified resource was not found
*/
type ReplaceFilterNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceFilterNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/filters/{index}][%d] replaceFilterNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceFilterNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceFilterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceFilterDefault creates a ReplaceFilterDefault with default headers values
func NewReplaceFilterDefault(code int) *ReplaceFilterDefault {
	return &ReplaceFilterDefault{
		_statusCode: code,
	}
}

/*ReplaceFilterDefault handles this case with default header values.

General Error
*/
type ReplaceFilterDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the replace filter default response
func (o *ReplaceFilterDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceFilterDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/filters/{index}][%d] replaceFilter default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceFilterDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceFilterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
