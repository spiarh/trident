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

// FileInfoCreateReader is a Reader for the FileInfoCreate structure.
type FileInfoCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FileInfoCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewFileInfoCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFileInfoCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFileInfoCreateCreated creates a FileInfoCreateCreated with default headers values
func NewFileInfoCreateCreated() *FileInfoCreateCreated {
	return &FileInfoCreateCreated{}
}

/* FileInfoCreateCreated describes a response with status code 201, with default header values.

Created
*/
type FileInfoCreateCreated struct {
}

func (o *FileInfoCreateCreated) Error() string {
	return fmt.Sprintf("[POST /storage/volumes/{volume.uuid}/files/{path}][%d] fileInfoCreateCreated ", 201)
}

func (o *FileInfoCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFileInfoCreateDefault creates a FileInfoCreateDefault with default headers values
func NewFileInfoCreateDefault(code int) *FileInfoCreateDefault {
	return &FileInfoCreateDefault{
		_statusCode: code,
	}
}

/* FileInfoCreateDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 917505 | The SVM does not exist. |
| 917525 | The volume in the symlink path does not exist in the SVM. |
| 917698 | The volume in the symlink path is not mounted in the namespace. |
| 6488064 | This command is not supported. |
| 6488065 | The volume in the symlink path is invalid. |
| 6488066 | Mounting the unjunctioned volume in the symlink path failed. |
| 6488069 | Internal file error. |
| 6488084 | Failed to create {path} because the "unix_permissions" field was not specified. |
| 6488085 | Failed to create {path} because the "type" field was not specified. |
| 8257536 | This operation is not supported for the system volume specified in the symlink path. |
| 8257541 | Failed to compute the SVM identification from this content. |
| 8257542 | This operation is not supported for the administrative SVM. |
| 9437549 | This operation is not allowed on SVMs with Infinite Volume. |
| 13172837 | This operation is not permitted because the SVM is locked for a migrate operation. |

*/
type FileInfoCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the file info create default response
func (o *FileInfoCreateDefault) Code() int {
	return o._statusCode
}

func (o *FileInfoCreateDefault) Error() string {
	return fmt.Sprintf("[POST /storage/volumes/{volume.uuid}/files/{path}][%d] file_info_create default  %+v", o._statusCode, o.Payload)
}
func (o *FileInfoCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FileInfoCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
