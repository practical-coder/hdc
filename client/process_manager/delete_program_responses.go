// Code generated by go-swagger; DO NOT EDIT.

package process_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// DeleteProgramReader is a Reader for the DeleteProgram structure.
type DeleteProgramReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProgramReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteProgramAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteProgramNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteProgramNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteProgramDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteProgramAccepted creates a DeleteProgramAccepted with default headers values
func NewDeleteProgramAccepted() *DeleteProgramAccepted {
	return &DeleteProgramAccepted{}
}

/*
DeleteProgramAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeleteProgramAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

// IsSuccess returns true when this delete program accepted response has a 2xx status code
func (o *DeleteProgramAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete program accepted response has a 3xx status code
func (o *DeleteProgramAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete program accepted response has a 4xx status code
func (o *DeleteProgramAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete program accepted response has a 5xx status code
func (o *DeleteProgramAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this delete program accepted response a status code equal to that given
func (o *DeleteProgramAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *DeleteProgramAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/programs/{name}][%d] deleteProgramAccepted ", 202)
}

func (o *DeleteProgramAccepted) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/programs/{name}][%d] deleteProgramAccepted ", 202)
}

func (o *DeleteProgramAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeleteProgramNoContent creates a DeleteProgramNoContent with default headers values
func NewDeleteProgramNoContent() *DeleteProgramNoContent {
	return &DeleteProgramNoContent{}
}

/*
DeleteProgramNoContent describes a response with status code 204, with default header values.

Program deleted
*/
type DeleteProgramNoContent struct {
}

// IsSuccess returns true when this delete program no content response has a 2xx status code
func (o *DeleteProgramNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete program no content response has a 3xx status code
func (o *DeleteProgramNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete program no content response has a 4xx status code
func (o *DeleteProgramNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete program no content response has a 5xx status code
func (o *DeleteProgramNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete program no content response a status code equal to that given
func (o *DeleteProgramNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteProgramNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/programs/{name}][%d] deleteProgramNoContent ", 204)
}

func (o *DeleteProgramNoContent) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/programs/{name}][%d] deleteProgramNoContent ", 204)
}

func (o *DeleteProgramNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProgramNotFound creates a DeleteProgramNotFound with default headers values
func NewDeleteProgramNotFound() *DeleteProgramNotFound {
	return &DeleteProgramNotFound{}
}

/*
DeleteProgramNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteProgramNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete program not found response has a 2xx status code
func (o *DeleteProgramNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete program not found response has a 3xx status code
func (o *DeleteProgramNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete program not found response has a 4xx status code
func (o *DeleteProgramNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete program not found response has a 5xx status code
func (o *DeleteProgramNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete program not found response a status code equal to that given
func (o *DeleteProgramNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteProgramNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/programs/{name}][%d] deleteProgramNotFound  %+v", 404, o.Payload)
}

func (o *DeleteProgramNotFound) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/programs/{name}][%d] deleteProgramNotFound  %+v", 404, o.Payload)
}

func (o *DeleteProgramNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteProgramNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteProgramDefault creates a DeleteProgramDefault with default headers values
func NewDeleteProgramDefault(code int) *DeleteProgramDefault {
	return &DeleteProgramDefault{
		_statusCode: code,
	}
}

/*
DeleteProgramDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteProgramDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete program default response
func (o *DeleteProgramDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete program default response has a 2xx status code
func (o *DeleteProgramDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete program default response has a 3xx status code
func (o *DeleteProgramDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete program default response has a 4xx status code
func (o *DeleteProgramDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete program default response has a 5xx status code
func (o *DeleteProgramDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete program default response a status code equal to that given
func (o *DeleteProgramDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteProgramDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/programs/{name}][%d] deleteProgram default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteProgramDefault) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/programs/{name}][%d] deleteProgram default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteProgramDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteProgramDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
