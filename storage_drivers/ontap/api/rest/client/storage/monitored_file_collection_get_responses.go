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

// MonitoredFileCollectionGetReader is a Reader for the MonitoredFileCollectionGet structure.
type MonitoredFileCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MonitoredFileCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMonitoredFileCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMonitoredFileCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMonitoredFileCollectionGetOK creates a MonitoredFileCollectionGetOK with default headers values
func NewMonitoredFileCollectionGetOK() *MonitoredFileCollectionGetOK {
	return &MonitoredFileCollectionGetOK{}
}

/* MonitoredFileCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type MonitoredFileCollectionGetOK struct {
	Payload *models.MonitoredFileResponse
}

func (o *MonitoredFileCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/monitored-files][%d] monitoredFileCollectionGetOK  %+v", 200, o.Payload)
}
func (o *MonitoredFileCollectionGetOK) GetPayload() *models.MonitoredFileResponse {
	return o.Payload
}

func (o *MonitoredFileCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MonitoredFileResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMonitoredFileCollectionGetDefault creates a MonitoredFileCollectionGetDefault with default headers values
func NewMonitoredFileCollectionGetDefault(code int) *MonitoredFileCollectionGetDefault {
	return &MonitoredFileCollectionGetDefault{
		_statusCode: code,
	}
}

/* MonitoredFileCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type MonitoredFileCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the monitored file collection get default response
func (o *MonitoredFileCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *MonitoredFileCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/monitored-files][%d] monitored_file_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *MonitoredFileCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MonitoredFileCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
