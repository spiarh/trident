// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SplitLoadGetReader is a Reader for the SplitLoadGet structure.
type SplitLoadGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SplitLoadGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSplitLoadGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSplitLoadGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSplitLoadGetOK creates a SplitLoadGetOK with default headers values
func NewSplitLoadGetOK() *SplitLoadGetOK {
	return &SplitLoadGetOK{}
}

/* SplitLoadGetOK describes a response with status code 200, with default header values.

OK
*/
type SplitLoadGetOK struct {
	Payload *models.SplitLoad
}

func (o *SplitLoadGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/file/clone/split-loads/{node.uuid}][%d] splitLoadGetOK  %+v", 200, o.Payload)
}
func (o *SplitLoadGetOK) GetPayload() *models.SplitLoad {
	return o.Payload
}

func (o *SplitLoadGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SplitLoad)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSplitLoadGetDefault creates a SplitLoadGetDefault with default headers values
func NewSplitLoadGetDefault(code int) *SplitLoadGetDefault {
	return &SplitLoadGetDefault{
		_statusCode: code,
	}
}

/* SplitLoadGetDefault describes a response with status code -1, with default header values.

Error
*/
type SplitLoadGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the split load get default response
func (o *SplitLoadGetDefault) Code() int {
	return o._statusCode
}

func (o *SplitLoadGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/file/clone/split-loads/{node.uuid}][%d] split_load_get default  %+v", o._statusCode, o.Payload)
}
func (o *SplitLoadGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SplitLoadGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
