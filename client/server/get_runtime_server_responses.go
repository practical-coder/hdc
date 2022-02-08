// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v3/models"
)

// GetRuntimeServerReader is a Reader for the GetRuntimeServer structure.
type GetRuntimeServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRuntimeServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRuntimeServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRuntimeServerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetRuntimeServerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRuntimeServerOK creates a GetRuntimeServerOK with default headers values
func NewGetRuntimeServerOK() *GetRuntimeServerOK {
	return &GetRuntimeServerOK{}
}

/*GetRuntimeServerOK handles this case with default header values.

Successful operation
*/
type GetRuntimeServerOK struct {
	Payload *models.RuntimeServer
}

func (o *GetRuntimeServerOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/servers/{name}][%d] getRuntimeServerOK  %+v", 200, o.Payload)
}

func (o *GetRuntimeServerOK) GetPayload() *models.RuntimeServer {
	return o.Payload
}

func (o *GetRuntimeServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeServer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRuntimeServerNotFound creates a GetRuntimeServerNotFound with default headers values
func NewGetRuntimeServerNotFound() *GetRuntimeServerNotFound {
	return &GetRuntimeServerNotFound{}
}

/*GetRuntimeServerNotFound handles this case with default header values.

The specified resource was not found
*/
type GetRuntimeServerNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetRuntimeServerNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/servers/{name}][%d] getRuntimeServerNotFound  %+v", 404, o.Payload)
}

func (o *GetRuntimeServerNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRuntimeServerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRuntimeServerDefault creates a GetRuntimeServerDefault with default headers values
func NewGetRuntimeServerDefault(code int) *GetRuntimeServerDefault {
	return &GetRuntimeServerDefault{
		_statusCode: code,
	}
}

/*GetRuntimeServerDefault handles this case with default header values.

General Error
*/
type GetRuntimeServerDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get runtime server default response
func (o *GetRuntimeServerDefault) Code() int {
	return o._statusCode
}

func (o *GetRuntimeServerDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/servers/{name}][%d] getRuntimeServer default  %+v", o._statusCode, o.Payload)
}

func (o *GetRuntimeServerDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRuntimeServerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
