// Code generated by go-swagger; DO NOT EDIT.

package acl_runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/cloud-native/models"
)

// GetServicesHaproxyRuntimeACLFileEntriesIDReader is a Reader for the GetServicesHaproxyRuntimeACLFileEntriesID structure.
type GetServicesHaproxyRuntimeACLFileEntriesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServicesHaproxyRuntimeACLFileEntriesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetServicesHaproxyRuntimeACLFileEntriesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetServicesHaproxyRuntimeACLFileEntriesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetServicesHaproxyRuntimeACLFileEntriesIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServicesHaproxyRuntimeACLFileEntriesIDOK creates a GetServicesHaproxyRuntimeACLFileEntriesIDOK with default headers values
func NewGetServicesHaproxyRuntimeACLFileEntriesIDOK() *GetServicesHaproxyRuntimeACLFileEntriesIDOK {
	return &GetServicesHaproxyRuntimeACLFileEntriesIDOK{}
}

/* GetServicesHaproxyRuntimeACLFileEntriesIDOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetServicesHaproxyRuntimeACLFileEntriesIDOK struct {
	Payload *models.ACLFileEntry
}

func (o *GetServicesHaproxyRuntimeACLFileEntriesIDOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/acl_file_entries/{id}][%d] getServicesHaproxyRuntimeAclFileEntriesIdOK  %+v", 200, o.Payload)
}
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDOK) GetPayload() *models.ACLFileEntry {
	return o.Payload
}

func (o *GetServicesHaproxyRuntimeACLFileEntriesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ACLFileEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServicesHaproxyRuntimeACLFileEntriesIDBadRequest creates a GetServicesHaproxyRuntimeACLFileEntriesIDBadRequest with default headers values
func NewGetServicesHaproxyRuntimeACLFileEntriesIDBadRequest() *GetServicesHaproxyRuntimeACLFileEntriesIDBadRequest {
	return &GetServicesHaproxyRuntimeACLFileEntriesIDBadRequest{}
}

/* GetServicesHaproxyRuntimeACLFileEntriesIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetServicesHaproxyRuntimeACLFileEntriesIDBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetServicesHaproxyRuntimeACLFileEntriesIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/acl_file_entries/{id}][%d] getServicesHaproxyRuntimeAclFileEntriesIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServicesHaproxyRuntimeACLFileEntriesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetServicesHaproxyRuntimeACLFileEntriesIDNotFound creates a GetServicesHaproxyRuntimeACLFileEntriesIDNotFound with default headers values
func NewGetServicesHaproxyRuntimeACLFileEntriesIDNotFound() *GetServicesHaproxyRuntimeACLFileEntriesIDNotFound {
	return &GetServicesHaproxyRuntimeACLFileEntriesIDNotFound{}
}

/* GetServicesHaproxyRuntimeACLFileEntriesIDNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetServicesHaproxyRuntimeACLFileEntriesIDNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetServicesHaproxyRuntimeACLFileEntriesIDNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/acl_file_entries/{id}][%d] getServicesHaproxyRuntimeAclFileEntriesIdNotFound  %+v", 404, o.Payload)
}
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServicesHaproxyRuntimeACLFileEntriesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetServicesHaproxyRuntimeACLFileEntriesIDDefault creates a GetServicesHaproxyRuntimeACLFileEntriesIDDefault with default headers values
func NewGetServicesHaproxyRuntimeACLFileEntriesIDDefault(code int) *GetServicesHaproxyRuntimeACLFileEntriesIDDefault {
	return &GetServicesHaproxyRuntimeACLFileEntriesIDDefault{
		_statusCode: code,
	}
}

/* GetServicesHaproxyRuntimeACLFileEntriesIDDefault describes a response with status code -1, with default header values.

General Error
*/
type GetServicesHaproxyRuntimeACLFileEntriesIDDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get services haproxy runtime ACL file entries ID default response
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDDefault) Code() int {
	return o._statusCode
}

func (o *GetServicesHaproxyRuntimeACLFileEntriesIDDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/acl_file_entries/{id}][%d] GetServicesHaproxyRuntimeACLFileEntriesID default  %+v", o._statusCode, o.Payload)
}
func (o *GetServicesHaproxyRuntimeACLFileEntriesIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServicesHaproxyRuntimeACLFileEntriesIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
