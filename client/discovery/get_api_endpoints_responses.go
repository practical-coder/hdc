// Code generated by go-swagger; DO NOT EDIT.

package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v3/models"
)

// GetAPIEndpointsReader is a Reader for the GetAPIEndpoints structure.
type GetAPIEndpointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIEndpointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIEndpointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAPIEndpointsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAPIEndpointsOK creates a GetAPIEndpointsOK with default headers values
func NewGetAPIEndpointsOK() *GetAPIEndpointsOK {
	return &GetAPIEndpointsOK{}
}

/*GetAPIEndpointsOK handles this case with default header values.

Success
*/
type GetAPIEndpointsOK struct {
	Payload models.Endpoints
}

func (o *GetAPIEndpointsOK) Error() string {
	return fmt.Sprintf("[GET /][%d] getApiEndpointsOK  %+v", 200, o.Payload)
}

func (o *GetAPIEndpointsOK) GetPayload() models.Endpoints {
	return o.Payload
}

func (o *GetAPIEndpointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIEndpointsDefault creates a GetAPIEndpointsDefault with default headers values
func NewGetAPIEndpointsDefault(code int) *GetAPIEndpointsDefault {
	return &GetAPIEndpointsDefault{
		_statusCode: code,
	}
}

/*GetAPIEndpointsDefault handles this case with default header values.

General Error
*/
type GetAPIEndpointsDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get API endpoints default response
func (o *GetAPIEndpointsDefault) Code() int {
	return o._statusCode
}

func (o *GetAPIEndpointsDefault) Error() string {
	return fmt.Sprintf("[GET /][%d] getAPIEndpoints default  %+v", o._statusCode, o.Payload)
}

func (o *GetAPIEndpointsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAPIEndpointsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
