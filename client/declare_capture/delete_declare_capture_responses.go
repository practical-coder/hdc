// Code generated by go-swagger; DO NOT EDIT.

package declare_capture

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
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

/*
DeleteDeclareCaptureAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeleteDeclareCaptureAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

// IsSuccess returns true when this delete declare capture accepted response has a 2xx status code
func (o *DeleteDeclareCaptureAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete declare capture accepted response has a 3xx status code
func (o *DeleteDeclareCaptureAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete declare capture accepted response has a 4xx status code
func (o *DeleteDeclareCaptureAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete declare capture accepted response has a 5xx status code
func (o *DeleteDeclareCaptureAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this delete declare capture accepted response a status code equal to that given
func (o *DeleteDeclareCaptureAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *DeleteDeclareCaptureAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/captures/{index}][%d] deleteDeclareCaptureAccepted ", 202)
}

func (o *DeleteDeclareCaptureAccepted) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/captures/{index}][%d] deleteDeclareCaptureAccepted ", 202)
}

func (o *DeleteDeclareCaptureAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeleteDeclareCaptureNoContent creates a DeleteDeclareCaptureNoContent with default headers values
func NewDeleteDeclareCaptureNoContent() *DeleteDeclareCaptureNoContent {
	return &DeleteDeclareCaptureNoContent{}
}

/*
DeleteDeclareCaptureNoContent describes a response with status code 204, with default header values.

Declare Capture deleted
*/
type DeleteDeclareCaptureNoContent struct {
}

// IsSuccess returns true when this delete declare capture no content response has a 2xx status code
func (o *DeleteDeclareCaptureNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete declare capture no content response has a 3xx status code
func (o *DeleteDeclareCaptureNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete declare capture no content response has a 4xx status code
func (o *DeleteDeclareCaptureNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete declare capture no content response has a 5xx status code
func (o *DeleteDeclareCaptureNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete declare capture no content response a status code equal to that given
func (o *DeleteDeclareCaptureNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteDeclareCaptureNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/captures/{index}][%d] deleteDeclareCaptureNoContent ", 204)
}

func (o *DeleteDeclareCaptureNoContent) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/captures/{index}][%d] deleteDeclareCaptureNoContent ", 204)
}

func (o *DeleteDeclareCaptureNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeclareCaptureNotFound creates a DeleteDeclareCaptureNotFound with default headers values
func NewDeleteDeclareCaptureNotFound() *DeleteDeclareCaptureNotFound {
	return &DeleteDeclareCaptureNotFound{}
}

/*
DeleteDeclareCaptureNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteDeclareCaptureNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete declare capture not found response has a 2xx status code
func (o *DeleteDeclareCaptureNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete declare capture not found response has a 3xx status code
func (o *DeleteDeclareCaptureNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete declare capture not found response has a 4xx status code
func (o *DeleteDeclareCaptureNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete declare capture not found response has a 5xx status code
func (o *DeleteDeclareCaptureNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete declare capture not found response a status code equal to that given
func (o *DeleteDeclareCaptureNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteDeclareCaptureNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/captures/{index}][%d] deleteDeclareCaptureNotFound  %+v", 404, o.Payload)
}

func (o *DeleteDeclareCaptureNotFound) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/captures/{index}][%d] deleteDeclareCaptureNotFound  %+v", 404, o.Payload)
}

func (o *DeleteDeclareCaptureNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDeclareCaptureNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteDeclareCaptureDefault creates a DeleteDeclareCaptureDefault with default headers values
func NewDeleteDeclareCaptureDefault(code int) *DeleteDeclareCaptureDefault {
	return &DeleteDeclareCaptureDefault{
		_statusCode: code,
	}
}

/*
DeleteDeclareCaptureDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteDeclareCaptureDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete declare capture default response
func (o *DeleteDeclareCaptureDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete declare capture default response has a 2xx status code
func (o *DeleteDeclareCaptureDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete declare capture default response has a 3xx status code
func (o *DeleteDeclareCaptureDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete declare capture default response has a 4xx status code
func (o *DeleteDeclareCaptureDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete declare capture default response has a 5xx status code
func (o *DeleteDeclareCaptureDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete declare capture default response a status code equal to that given
func (o *DeleteDeclareCaptureDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteDeclareCaptureDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/captures/{index}][%d] deleteDeclareCapture default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDeclareCaptureDefault) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/captures/{index}][%d] deleteDeclareCapture default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDeclareCaptureDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDeclareCaptureDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
