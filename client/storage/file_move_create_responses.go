// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sapcc/ontap-restapi/models"
)

// FileMoveCreateReader is a Reader for the FileMoveCreate structure.
type FileMoveCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FileMoveCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewFileMoveCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFileMoveCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFileMoveCreateCreated creates a FileMoveCreateCreated with default headers values
func NewFileMoveCreateCreated() *FileMoveCreateCreated {
	return &FileMoveCreateCreated{}
}

/*
FileMoveCreateCreated describes a response with status code 201, with default header values.

Created
*/
type FileMoveCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string
}

// IsSuccess returns true when this file move create created response has a 2xx status code
func (o *FileMoveCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this file move create created response has a 3xx status code
func (o *FileMoveCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this file move create created response has a 4xx status code
func (o *FileMoveCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this file move create created response has a 5xx status code
func (o *FileMoveCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this file move create created response a status code equal to that given
func (o *FileMoveCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the file move create created response
func (o *FileMoveCreateCreated) Code() int {
	return 201
}

func (o *FileMoveCreateCreated) Error() string {
	return fmt.Sprintf("[POST /storage/file/moves][%d] fileMoveCreateCreated", 201)
}

func (o *FileMoveCreateCreated) String() string {
	return fmt.Sprintf("[POST /storage/file/moves][%d] fileMoveCreateCreated", 201)
}

func (o *FileMoveCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	return nil
}

// NewFileMoveCreateDefault creates a FileMoveCreateDefault with default headers values
func NewFileMoveCreateDefault(code int) *FileMoveCreateDefault {
	return &FileMoveCreateDefault{
		_statusCode: code,
	}
}

/*
	FileMoveCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 918236 | The specified \"volume.uuid\" and \"volume.name\" do not refer to the same volume. |
| 2621462 | SVM `svm.name` does not exist. |
| 2621706 | The specified \"svm.uuid\" and \"svm.name\" do not refer to the same SVM. |
| 7012352 | File locations are inconsistent. All files must be on the same volume. |
| 7012353 | Exceeded the file operations supported number of files. |
| 7012354 | Unable to pair the number of source files to destination files. |
| 7012357 | Cannot start a file operation until all cluster nodes support the file operations capability. |
| 7012358 | The specified source path is invalid. |
| 7012359 | The specified destination path is invalid. |
| 7012360 | The SVMs are not in an intracluster peering relationship. |
| 7012361 | The SVMs peering relationship does not include application \"file-move\". |
| 7012362 | The SVMs are not yet in a peered state. |
| 7012363 | Cannot move files. All file operations must be managed by the destination SVM's administrator. |
| 7012364 | Cannot move file from SVM \"svm.name\" to SVM \"svm.name\". Moving a file between SVMs is not supported. |
| 7012365 | Copying a file between clusters is not supported. |
| 7012367 | A reference path may only be specified if multiple source paths are specified. |
| 7012368 | The reference path must have a matching source path. |
| 7012371 | The reference cutover time exceeds the maximum allowable time. |
| 7012374 | Source volume and destination volume have different home clusters. |
| 7012376 | Operation not allowed on a volume that is part of a SnapMirror Synchronous relationship. |
| 7012377 | Cannot start a file move operation on the volume because an active volume conversion is in progress. |
| 7013352 | One or more source files must be specified in \"files_to_move.sources\". |
| 7013353 | One or more destination files or directories must be specified in \"files_to_move.destinations\". |
| 7013354 | \"path\" must be specified. |
| 7013355 | Moving files between FlexVol volumes and FlexGroup volumes or constituents is not supported. |
| 7013356 | Moving files between FlexVol volumes and FlexGroup constituents is not supported. |
| 7013357 | The specified volume could not be found. |
| 7013358 | The specified SVM or volume UUID could not be found. |
| 7013359 | The SVM and volume must both be provided. |
| 7018861 | Cannot start the file operation. The volume is offline. |
| 7018877 | Maximum combined total (50) of file and LUN copy and move operations reached. When one or more of the operations has completed, try the command again. |
| 7018937 | The file is already on the destination constituent. |
| 13107222 | Internal error. |
| 13107415 | Failed to lookup a volume property. |
| 13107431 | Failed to lookup an SVM property. |
| 13109260 | Failed to enable granular data on the volume. |
| 144179201 | The file move start operation failed. |
| 144179206 | Source file does not exist. |
| 144179207 | Volume capacity balancing requires an effective cluster version of 9.10.1 or later. |
| 144180200 | Destination constituent not a member of FlexGroup volume. |
| 144180201 | Destination constituent not properly configured. |
| 144180203 | Volume capacity rebalancing is not supported on FlexCache volumes. |
| 144180204 | Volume capacity rebalancing is not supported on object store volumes. |
| 144180205 | The system is busy. |
| 144180207 | Volume capacity rebalancing is not supported on inactive MetroCluster configurations. |
| 144180208 | Disruptive file movement is not supported when granular data is enabled on the volume. Try the operation again using \"disruptive=false\". |
| 144181200 | Too many source or destination files are specified for a file move within a FlexGroup volume. There must be one source file identifying a file on a FlexGroup volume and either zero or one destination files identifying the destination constituent. |
| 144181201 | For a file move within a FlexGroup volume, the source volume must be a FlexGroup volume, and the destination volume must be a constituent. |
| 144181202 | The specified source volume UUID is for a constituent. For a file move within a FlexGroup volume, the source volume must be a FlexGroup volume, and the destination volume must be a constituent. |
| 144181203 | A destination constituent must be provided in \"files_to_move.destinations\" if it is not being selected automatically. Use the \"automatic\" query to enable automatic destination constituent selection. |
| 144181204 | A destination constituent is provided while automatic destination constituent selection is enabled with the \"automatic\" query. |
| 144181205 | The destination volume is not a constituent. For a file move within a FlexGroup volume, the destination volume must be a constituent of the source FlexGroup volume. |
| 144181207 | The destination constituent SVM is not the same as the source SVM. For a file move within a FlexGroup volume, the destination constituent must be a constituent of the source FlexGroup volume. |
| 144181208 | The destination file path is different from the source file path. For a file move within a FlexGroup volume, the path of the source file does not change. |
| 144181209 | The specified SVM for the destination constituent differs from the SVM of the source FlexGroup volume. For a file move within a FlexGroup volume, the destination constituent must be a constituent of the source FlexGroup volume. |
| 144181210 | The \"force\" parameter is not supported unless the \"disruptive\" parameter is specified as \"true\". |
| 144182201 | Volume capacity rebalancing using non-disruptive file move operations and granular data requires an effective cluster version of 9.11.1 or later. |
| 144182228 | Operation is not supported on FlexGroup volumes with only one constituent. |
| 196608143 | Cannot start the operation. The volume is undergoing a secure purge operation. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type FileMoveCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this file move create default response has a 2xx status code
func (o *FileMoveCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this file move create default response has a 3xx status code
func (o *FileMoveCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this file move create default response has a 4xx status code
func (o *FileMoveCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this file move create default response has a 5xx status code
func (o *FileMoveCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this file move create default response a status code equal to that given
func (o *FileMoveCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the file move create default response
func (o *FileMoveCreateDefault) Code() int {
	return o._statusCode
}

func (o *FileMoveCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/file/moves][%d] file_move_create default %s", o._statusCode, payload)
}

func (o *FileMoveCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/file/moves][%d] file_move_create default %s", o._statusCode, payload)
}

func (o *FileMoveCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FileMoveCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
