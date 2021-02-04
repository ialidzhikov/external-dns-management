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

// EnableHookReader is a Reader for the EnableHook structure.
type EnableHookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableHookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnableHookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEnableHookDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEnableHookOK creates a EnableHookOK with default headers values
func NewEnableHookOK() *EnableHookOK {
	return &EnableHookOK{}
}

/*EnableHookOK handles this case with default header values.

OK
*/
type EnableHookOK struct {
	Payload *models.Hook
}

func (o *EnableHookOK) Error() string {
	return fmt.Sprintf("[POST /hooks/{hook_id}/enable][%d] enableHookOK  %+v", 200, o.Payload)
}

func (o *EnableHookOK) GetPayload() *models.Hook {
	return o.Payload
}

func (o *EnableHookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Hook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableHookDefault creates a EnableHookDefault with default headers values
func NewEnableHookDefault(code int) *EnableHookDefault {
	return &EnableHookDefault{
		_statusCode: code,
	}
}

/*EnableHookDefault handles this case with default header values.

error
*/
type EnableHookDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the enable hook default response
func (o *EnableHookDefault) Code() int {
	return o._statusCode
}

func (o *EnableHookDefault) Error() string {
	return fmt.Sprintf("[POST /hooks/{hook_id}/enable][%d] enableHook default  %+v", o._statusCode, o.Payload)
}

func (o *EnableHookDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *EnableHookDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
