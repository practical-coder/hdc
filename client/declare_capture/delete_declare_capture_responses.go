// Code generated by go-swagger; DO NOT EDIT.

package declare_capture

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v3/models"
)

// DeleteDeclareCaptureReader is a Reader for the DeleteDeclareCapture structure.
type DeleteDeclareCaptureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDeclareCaptureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteDeclareCaptureAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteDeclareCaptureNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteDeclareCaptureNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteDeclareCaptureDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteDeclareCaptureAccepted creates a DeleteDeclareCaptureAccepted with default headers values
func NewDeleteDeclareCaptureAccepted() *DeleteDeclareCaptureAccepted {
	return &DeleteDeclareCaptureAccepted{}
}

/*DeleteDeclareCaptureAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type DeleteDeclareCaptureAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteDeclareCaptureAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/captures/{index}][%d] deleteDeclareCaptureAccepted ", 202)
}

func (o *DeleteDeclareCaptureAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	return nil
}

// NewDeleteDeclareCaptureNoContent creates a DeleteDeclareCaptureNoContent with default headers values
func NewDeleteDeclareCaptureNoContent() *DeleteDeclareCaptureNoContent {
	return &DeleteDeclareCaptureNoContent{}
}

/*DeleteDeclareCaptureNoContent handles this case with default header values.

Declare Capture deleted
*/
type DeleteDeclareCaptureNoContent struct {
}

func (o *DeleteDeclareCaptureNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/captures/{index}][%d] deleteDeclareCaptureNoContent ", 204)
}

func (o *DeleteDeclareCaptureNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeclareCaptureNotFound creates a DeleteDeclareCaptureNotFound with default headers values
func NewDeleteDeclareCaptureNotFound() *DeleteDeclareCaptureNotFound {
	return &DeleteDeclareCaptureNotFound{}
}

/*DeleteDeclareCaptureNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteDeclareCaptureNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteDeclareCaptureNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/captures/{index}][%d] deleteDeclareCaptureNotFound  %+v", 404, o.Payload)
}

func (o *DeleteDeclareCaptureNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDeclareCaptureNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeclareCaptureDefault creates a DeleteDeclareCaptureDefault with default headers values
func NewDeleteDeclareCaptureDefault(code int) *DeleteDeclareCaptureDefault {
	return &DeleteDeclareCaptureDefault{
		_statusCode: code,
	}
}

/*DeleteDeclareCaptureDefault handles this case with default header values.

General Error
*/
type DeleteDeclareCaptureDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete declare capture default response
func (o *DeleteDeclareCaptureDefault) Code() int {
	return o._statusCode
}

func (o *DeleteDeclareCaptureDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/captures/{index}][%d] deleteDeclareCapture default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDeclareCaptureDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDeclareCaptureDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}