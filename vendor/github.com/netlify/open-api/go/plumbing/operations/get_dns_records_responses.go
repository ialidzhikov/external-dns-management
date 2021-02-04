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

// GetDNSRecordsReader is a Reader for the GetDNSRecords structure.
type GetDNSRecordsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDNSRecordsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDNSRecordsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDNSRecordsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDNSRecordsOK creates a GetDNSRecordsOK with default headers values
func NewGetDNSRecordsOK() *GetDNSRecordsOK {
	return &GetDNSRecordsOK{}
}

/*GetDNSRecordsOK handles this case with default header values.

get all DNS records for a single DNS zone
*/
type GetDNSRecordsOK struct {
	Payload models.DNSRecords
}

func (o *GetDNSRecordsOK) Error() string {
	return fmt.Sprintf("[GET /dns_zones/{zone_id}/dns_records][%d] getDnsRecordsOK  %+v", 200, o.Payload)
}

func (o *GetDNSRecordsOK) GetPayload() models.DNSRecords {
	return o.Payload
}

func (o *GetDNSRecordsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDNSRecordsDefault creates a GetDNSRecordsDefault with default headers values
func NewGetDNSRecordsDefault(code int) *GetDNSRecordsDefault {
	return &GetDNSRecordsDefault{
		_statusCode: code,
	}
}

/*GetDNSRecordsDefault handles this case with default header values.

error
*/
type GetDNSRecordsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get Dns records default response
func (o *GetDNSRecordsDefault) Code() int {
	return o._statusCode
}

func (o *GetDNSRecordsDefault) Error() string {
	return fmt.Sprintf("[GET /dns_zones/{zone_id}/dns_records][%d] getDnsRecords default  %+v", o._statusCode, o.Payload)
}

func (o *GetDNSRecordsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDNSRecordsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
