// Code generated by go-swagger; DO NOT EDIT.

package acl_runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v2/models"
)

// GetServicesHaproxyRuntimeAclsReader is a Reader for the GetServicesHaproxyRuntimeAcls structure.
type GetServicesHaproxyRuntimeAclsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServicesHaproxyRuntimeAclsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServicesHaproxyRuntimeAclsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetServicesHaproxyRuntimeAclsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServicesHaproxyRuntimeAclsOK creates a GetServicesHaproxyRuntimeAclsOK with default headers values
func NewGetServicesHaproxyRuntimeAclsOK() *GetServicesHaproxyRuntimeAclsOK {
	return &GetServicesHaproxyRuntimeAclsOK{}
}

/* GetServicesHaproxyRuntimeAclsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetServicesHaproxyRuntimeAclsOK struct {
	Payload models.ACLFiles
}

func (o *GetServicesHaproxyRuntimeAclsOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/acls][%d] getServicesHaproxyRuntimeAclsOK  %+v", 200, o.Payload)
}
func (o *GetServicesHaproxyRuntimeAclsOK) GetPayload() models.ACLFiles {
	return o.Payload
}

func (o *GetServicesHaproxyRuntimeAclsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServicesHaproxyRuntimeAclsDefault creates a GetServicesHaproxyRuntimeAclsDefault with default headers values
func NewGetServicesHaproxyRuntimeAclsDefault(code int) *GetServicesHaproxyRuntimeAclsDefault {
	return &GetServicesHaproxyRuntimeAclsDefault{
		_statusCode: code,
	}
}

/* GetServicesHaproxyRuntimeAclsDefault describes a response with status code -1, with default header values.

General Error
*/
type GetServicesHaproxyRuntimeAclsDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get services haproxy runtime acls default response
func (o *GetServicesHaproxyRuntimeAclsDefault) Code() int {
	return o._statusCode
}

func (o *GetServicesHaproxyRuntimeAclsDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/acls][%d] GetServicesHaproxyRuntimeAcls default  %+v", o._statusCode, o.Payload)
}
func (o *GetServicesHaproxyRuntimeAclsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServicesHaproxyRuntimeAclsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
