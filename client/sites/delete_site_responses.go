// Code generated by go-swagger; DO NOT EDIT.

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"client-native/models"
)

// DeleteSiteReader is a Reader for the DeleteSite structure.
type DeleteSiteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSiteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteSiteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteSiteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteSiteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteSiteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSiteAccepted creates a DeleteSiteAccepted with default headers values
func NewDeleteSiteAccepted() *DeleteSiteAccepted {
	return &DeleteSiteAccepted{}
}

/* DeleteSiteAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeleteSiteAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteSiteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/sites/{name}][%d] deleteSiteAccepted ", 202)
}

func (o *DeleteSiteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeleteSiteNoContent creates a DeleteSiteNoContent with default headers values
func NewDeleteSiteNoContent() *DeleteSiteNoContent {
	return &DeleteSiteNoContent{}
}

/* DeleteSiteNoContent describes a response with status code 204, with default header values.

Site deleted
*/
type DeleteSiteNoContent struct {
}

func (o *DeleteSiteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/sites/{name}][%d] deleteSiteNoContent ", 204)
}

func (o *DeleteSiteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSiteNotFound creates a DeleteSiteNotFound with default headers values
func NewDeleteSiteNotFound() *DeleteSiteNotFound {
	return &DeleteSiteNotFound{}
}

/* DeleteSiteNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteSiteNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteSiteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/sites/{name}][%d] deleteSiteNotFound  %+v", 404, o.Payload)
}
func (o *DeleteSiteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteSiteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteSiteDefault creates a DeleteSiteDefault with default headers values
func NewDeleteSiteDefault(code int) *DeleteSiteDefault {
	return &DeleteSiteDefault{
		_statusCode: code,
	}
}

/* DeleteSiteDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteSiteDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete site default response
func (o *DeleteSiteDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSiteDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/sites/{name}][%d] deleteSite default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteSiteDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteSiteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
