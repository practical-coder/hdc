// Code generated by go-swagger; DO NOT EDIT.

package acl

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v5/models"
)

// DeleteACLReader is a Reader for the DeleteACL structure.
type DeleteACLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteACLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteACLAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteACLNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteACLNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteACLDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteACLAccepted creates a DeleteACLAccepted with default headers values
func NewDeleteACLAccepted() *DeleteACLAccepted {
	return &DeleteACLAccepted{}
}

/*
DeleteACLAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeleteACLAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

// IsSuccess returns true when this delete Acl accepted response has a 2xx status code
func (o *DeleteACLAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete Acl accepted response has a 3xx status code
func (o *DeleteACLAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Acl accepted response has a 4xx status code
func (o *DeleteACLAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Acl accepted response has a 5xx status code
func (o *DeleteACLAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Acl accepted response a status code equal to that given
func (o *DeleteACLAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the delete Acl accepted response
func (o *DeleteACLAccepted) Code() int {
	return 202
}

func (o *DeleteACLAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/acls/{index}][%d] deleteAclAccepted ", 202)
}

func (o *DeleteACLAccepted) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/acls/{index}][%d] deleteAclAccepted ", 202)
}

func (o *DeleteACLAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeleteACLNoContent creates a DeleteACLNoContent with default headers values
func NewDeleteACLNoContent() *DeleteACLNoContent {
	return &DeleteACLNoContent{}
}

/*
DeleteACLNoContent describes a response with status code 204, with default header values.

ACL line deleted
*/
type DeleteACLNoContent struct {
}

// IsSuccess returns true when this delete Acl no content response has a 2xx status code
func (o *DeleteACLNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete Acl no content response has a 3xx status code
func (o *DeleteACLNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Acl no content response has a 4xx status code
func (o *DeleteACLNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Acl no content response has a 5xx status code
func (o *DeleteACLNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Acl no content response a status code equal to that given
func (o *DeleteACLNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete Acl no content response
func (o *DeleteACLNoContent) Code() int {
	return 204
}

func (o *DeleteACLNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/acls/{index}][%d] deleteAclNoContent ", 204)
}

func (o *DeleteACLNoContent) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/acls/{index}][%d] deleteAclNoContent ", 204)
}

func (o *DeleteACLNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteACLNotFound creates a DeleteACLNotFound with default headers values
func NewDeleteACLNotFound() *DeleteACLNotFound {
	return &DeleteACLNotFound{}
}

/*
DeleteACLNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteACLNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete Acl not found response has a 2xx status code
func (o *DeleteACLNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Acl not found response has a 3xx status code
func (o *DeleteACLNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Acl not found response has a 4xx status code
func (o *DeleteACLNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Acl not found response has a 5xx status code
func (o *DeleteACLNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Acl not found response a status code equal to that given
func (o *DeleteACLNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete Acl not found response
func (o *DeleteACLNotFound) Code() int {
	return 404
}

func (o *DeleteACLNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/acls/{index}][%d] deleteAclNotFound  %+v", 404, o.Payload)
}

func (o *DeleteACLNotFound) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/acls/{index}][%d] deleteAclNotFound  %+v", 404, o.Payload)
}

func (o *DeleteACLNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteACLNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteACLDefault creates a DeleteACLDefault with default headers values
func NewDeleteACLDefault(code int) *DeleteACLDefault {
	return &DeleteACLDefault{
		_statusCode: code,
	}
}

/*
DeleteACLDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteACLDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// IsSuccess returns true when this delete Acl default response has a 2xx status code
func (o *DeleteACLDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete Acl default response has a 3xx status code
func (o *DeleteACLDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete Acl default response has a 4xx status code
func (o *DeleteACLDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete Acl default response has a 5xx status code
func (o *DeleteACLDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete Acl default response a status code equal to that given
func (o *DeleteACLDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete Acl default response
func (o *DeleteACLDefault) Code() int {
	return o._statusCode
}

func (o *DeleteACLDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/acls/{index}][%d] deleteAcl default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteACLDefault) String() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/acls/{index}][%d] deleteAcl default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteACLDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteACLDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
