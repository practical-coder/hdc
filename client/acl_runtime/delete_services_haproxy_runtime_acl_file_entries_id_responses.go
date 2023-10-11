// Code generated by go-swagger; DO NOT EDIT.

package acl_runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v5/models"
)

// DeleteServicesHaproxyRuntimeACLFileEntriesIDReader is a Reader for the DeleteServicesHaproxyRuntimeACLFileEntriesID structure.
type DeleteServicesHaproxyRuntimeACLFileEntriesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteServicesHaproxyRuntimeACLFileEntriesIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteServicesHaproxyRuntimeACLFileEntriesIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteServicesHaproxyRuntimeACLFileEntriesIDNoContent creates a DeleteServicesHaproxyRuntimeACLFileEntriesIDNoContent with default headers values
func NewDeleteServicesHaproxyRuntimeACLFileEntriesIDNoContent() *DeleteServicesHaproxyRuntimeACLFileEntriesIDNoContent {
	return &DeleteServicesHaproxyRuntimeACLFileEntriesIDNoContent{}
}

/*
DeleteServicesHaproxyRuntimeACLFileEntriesIDNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteServicesHaproxyRuntimeACLFileEntriesIDNoContent struct {
}

// IsSuccess returns true when this delete services haproxy runtime Acl file entries Id no content response has a 2xx status code
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete services haproxy runtime Acl file entries Id no content response has a 3xx status code
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete services haproxy runtime Acl file entries Id no content response has a 4xx status code
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete services haproxy runtime Acl file entries Id no content response has a 5xx status code
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete services haproxy runtime Acl file entries Id no content response a status code equal to that given
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete services haproxy runtime Acl file entries Id no content response
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDNoContent) Code() int {
	return 204
}

func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/runtime/acl_file_entries/{id}][%d] deleteServicesHaproxyRuntimeAclFileEntriesIdNoContent ", 204)
}

func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/runtime/acl_file_entries/{id}][%d] deleteServicesHaproxyRuntimeAclFileEntriesIdNoContent ", 204)
}

func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest creates a DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest with default headers values
func NewDeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest() *DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest {
	return &DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest{}
}

/*
DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete services haproxy runtime Acl file entries Id bad request response has a 2xx status code
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete services haproxy runtime Acl file entries Id bad request response has a 3xx status code
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete services haproxy runtime Acl file entries Id bad request response has a 4xx status code
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete services haproxy runtime Acl file entries Id bad request response has a 5xx status code
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete services haproxy runtime Acl file entries Id bad request response a status code equal to that given
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete services haproxy runtime Acl file entries Id bad request response
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest) Code() int {
	return 400
}

func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/runtime/acl_file_entries/{id}][%d] deleteServicesHaproxyRuntimeAclFileEntriesIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/runtime/acl_file_entries/{id}][%d] deleteServicesHaproxyRuntimeAclFileEntriesIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound creates a DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound with default headers values
func NewDeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound() *DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound {
	return &DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound{}
}

/*
DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete services haproxy runtime Acl file entries Id not found response has a 2xx status code
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete services haproxy runtime Acl file entries Id not found response has a 3xx status code
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete services haproxy runtime Acl file entries Id not found response has a 4xx status code
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete services haproxy runtime Acl file entries Id not found response has a 5xx status code
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete services haproxy runtime Acl file entries Id not found response a status code equal to that given
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete services haproxy runtime Acl file entries Id not found response
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound) Code() int {
	return 404
}

func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/runtime/acl_file_entries/{id}][%d] deleteServicesHaproxyRuntimeAclFileEntriesIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/runtime/acl_file_entries/{id}][%d] deleteServicesHaproxyRuntimeAclFileEntriesIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteServicesHaproxyRuntimeACLFileEntriesIDDefault creates a DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault with default headers values
func NewDeleteServicesHaproxyRuntimeACLFileEntriesIDDefault(code int) *DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault {
	return &DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault{
		_statusCode: code,
	}
}

/*
DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete services haproxy runtime ACL file entries ID default response has a 2xx status code
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete services haproxy runtime ACL file entries ID default response has a 3xx status code
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete services haproxy runtime ACL file entries ID default response has a 4xx status code
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete services haproxy runtime ACL file entries ID default response has a 5xx status code
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete services haproxy runtime ACL file entries ID default response a status code equal to that given
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete services haproxy runtime ACL file entries ID default response
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/runtime/acl_file_entries/{id}][%d] DeleteServicesHaproxyRuntimeACLFileEntriesID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/runtime/acl_file_entries/{id}][%d] DeleteServicesHaproxyRuntimeACLFileEntriesID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
