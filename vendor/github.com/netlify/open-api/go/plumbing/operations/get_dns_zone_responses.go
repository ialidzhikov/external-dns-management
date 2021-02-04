// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/go/models"
)

// GetDNSZoneReader is a Reader for the GetDNSZone structure.
type GetDNSZoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDNSZoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDNSZoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDNSZoneDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDNSZoneOK creates a GetDNSZoneOK with default headers values
func NewGetDNSZoneOK() *GetDNSZoneOK {
	return &GetDNSZoneOK{}
}

/*GetDNSZoneOK handles this case with default header values.

get a single DNS zone
*/
type GetDNSZoneOK struct {
	Payload *models.DNSZone
}

func (o *GetDNSZoneOK) Error() string {
	return fmt.Sprintf("[GET /dns_zones/{zone_id}][%d] getDnsZoneOK  %+v", 200, o.Payload)
}

func (o *GetDNSZoneOK) GetPayload() *models.DNSZone {
	return o.Payload
}

func (o *GetDNSZoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DNSZone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDNSZoneDefault creates a GetDNSZoneDefault with default headers values
func NewGetDNSZoneDefault(code int) *GetDNSZoneDefault {
	return &GetDNSZoneDefault{
		_statusCode: code,
	}
}

/*GetDNSZoneDefault handles this case with default header values.

error
*/
type GetDNSZoneDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get Dns zone default response
func (o *GetDNSZoneDefault) Code() int {
	return o._statusCode
}

func (o *GetDNSZoneDefault) Error() string {
	return fmt.Sprintf("[GET /dns_zones/{zone_id}][%d] getDnsZone default  %+v", o._statusCode, o.Payload)
}

func (o *GetDNSZoneDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDNSZoneDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
