// Code generated by go-swagger; DO NOT EDIT.

package object_store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sapcc/ontap-restapi-client/swagger/models"
)

// S3BucketSvmCreateReader is a Reader for the S3BucketSvmCreate structure.
type S3BucketSvmCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3BucketSvmCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewS3BucketSvmCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewS3BucketSvmCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3BucketSvmCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3BucketSvmCreateCreated creates a S3BucketSvmCreateCreated with default headers values
func NewS3BucketSvmCreateCreated() *S3BucketSvmCreateCreated {
	return &S3BucketSvmCreateCreated{}
}

/*
S3BucketSvmCreateCreated describes a response with status code 201, with default header values.

Created
*/
type S3BucketSvmCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this s3 bucket svm create created response has a 2xx status code
func (o *S3BucketSvmCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 bucket svm create created response has a 3xx status code
func (o *S3BucketSvmCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 bucket svm create created response has a 4xx status code
func (o *S3BucketSvmCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 bucket svm create created response has a 5xx status code
func (o *S3BucketSvmCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 bucket svm create created response a status code equal to that given
func (o *S3BucketSvmCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the s3 bucket svm create created response
func (o *S3BucketSvmCreateCreated) Code() int {
	return 201
}

func (o *S3BucketSvmCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/s3/services/{svm.uuid}/buckets][%d] s3BucketSvmCreateCreated %s", 201, payload)
}

func (o *S3BucketSvmCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/s3/services/{svm.uuid}/buckets][%d] s3BucketSvmCreateCreated %s", 201, payload)
}

func (o *S3BucketSvmCreateCreated) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *S3BucketSvmCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3BucketSvmCreateAccepted creates a S3BucketSvmCreateAccepted with default headers values
func NewS3BucketSvmCreateAccepted() *S3BucketSvmCreateAccepted {
	return &S3BucketSvmCreateAccepted{}
}

/*
S3BucketSvmCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type S3BucketSvmCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this s3 bucket svm create accepted response has a 2xx status code
func (o *S3BucketSvmCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 bucket svm create accepted response has a 3xx status code
func (o *S3BucketSvmCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 bucket svm create accepted response has a 4xx status code
func (o *S3BucketSvmCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 bucket svm create accepted response has a 5xx status code
func (o *S3BucketSvmCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 bucket svm create accepted response a status code equal to that given
func (o *S3BucketSvmCreateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the s3 bucket svm create accepted response
func (o *S3BucketSvmCreateAccepted) Code() int {
	return 202
}

func (o *S3BucketSvmCreateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/s3/services/{svm.uuid}/buckets][%d] s3BucketSvmCreateAccepted %s", 202, payload)
}

func (o *S3BucketSvmCreateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/s3/services/{svm.uuid}/buckets][%d] s3BucketSvmCreateAccepted %s", 202, payload)
}

func (o *S3BucketSvmCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *S3BucketSvmCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3BucketSvmCreateDefault creates a S3BucketSvmCreateDefault with default headers values
func NewS3BucketSvmCreateDefault(code int) *S3BucketSvmCreateDefault {
	return &S3BucketSvmCreateDefault{
		_statusCode: code,
	}
}

/*
	S3BucketSvmCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error code | Message |
| ---------- | ------- |
| 92405777   | "Failed to create bucket \\\"{bucket name}\\\" for SVM \\\"{svm.name}\\\". Reason: {Reason of failure}. ";
| 92405785   | "Bucket name \\\"{bucket name}\\\" contains invalid characters or invalid character combinations. Valid characters for a bucket name are 0-9, a-z, \\\".\\\", and \\\"-\\\". Invalid character combinations are \\\".-\\\", \\\"-.\\\", and \\\"..\\\". ";
| 92405786   | "Bucket name \\\"{bucket name}\\\" is not valid. Bucket names must have between 3 and 63 characters. ";
| 92405811   | "Failed to create bucket \\\"{bucket name}\\\" for SVM \\\"{svm.name}\\\". Wait a few minutes and try the operation again.";
| 92405812   | "Failed to create the object store volume. Reason: {Reason for failure}.";
| 92405819   | "Cannot provision an object store server volume for bucket \\\"{bucket name}\\\" in SVM \\\"{svm.name}\\\" on the following aggregates because they are SnapLock aggregates: {List of aggregates.name}.";
| 92405820   | "Failed to check whether the aggregate \\\"{aggregates.name}\\\" is a FabricPool. Reason: {Reason for failure}.";
| 92405821   | "Cannot provision an object store server volume for bucket \\\"{bucket name}\\\" in SVM \\\"{svm.name}\\\" on the following aggregates because they are FabricPool: {List of aggregates.name}.";
| 92405827   | "Internal Error. Unable to generate object store volume name.";
| 92405857   | "One or more aggregates must be specified if \\\"constituents_per_aggregate\\\" is specified.";
| 92405858   | "Failed to \\\"create\\\" the \\\"bucket\\\" because the operation is only supported on data SVMs.";
| 92405859   | "The specified \\\"aggregates.uuid\\\" \\\"{aggregates.uuid}\\\" does not exist.";
| 92405860   | "The specified "aggregates.name" "{aggregates.name}" and "aggregates.uuid" "{aggregates.uuid}" do not refer to same aggregate.";
| 92405861   | "The specified SVM UUID or bucket UUID does not exist.";
| 92405863   | "Failed to create access policies for bucket \\\"{bucket name}\\\". Reason: {Reason of Failure}. Resolve all the issues and retry the operation.";
| 92405863   | "Failed to create event selector for bucket \\\"{bucket name}\\\".  Reason: {Reason of failure}. Reason: object-store-server audit configuration not created for SVM \\\"{svm.name}\\\". Resolve all the issues and retry the operation.";
| 92405891   | "The resources specified in the access policy are not valid. Valid ways to specify a resource are \\\"*\\\", \\\"<bucket-name>\\\", \\\"<bucket-name>/.../...\\\". Valid characters for a resource are 0-9, A-Z, a-z, \\\"_\\\", \\\"+\\\", \\\",\\\", \\\";\\\", \\\":\\\", \\\";\\\", \\\"=\\\", \\\".\\\", \\\"&\\\", \\\"@\\\", \\\"?\\\", \\\"(\\\", \\\")\\\", \\\"'\\\", \\\"*\\\", \\\"!\\\", \\\"-\\\" and \\\"\\$\\\".";
| 92405894   | "Statements, principals and resources list can have a maximum of 10 entries.";
| 92405897   | "The principals specified in the access policy are not in the correct format. User name must be in between 1 and 64 characters. Valid characters for a user name are 0-9, A-Z, a-z, \\\"_\\\", \\\"+\\\", \\\"=\\\", \\\",\\\", \\\".\\\", \\\"@\\\", and \\\"-\\\". ";
| 92405898   | "The SID specified in the access policy is not valid. Valid characters for a SID are 0-9, A-Z and a-z.";
| 92733688   | "Failed to create bucket \\\"{bucket name}\\\" for SVM \\\"{svm.name}\\\". Reason: Unable to find a suitable aggregate for volume \\\"{volume name}\\\" in SVM \\\"{svm.name}\\\". Reason: No candidate aggregates are available for storage services: performance. ";
| 460770     | "[Job job_number] Job failed: Failed to create bucket \\\"{bucket name}\\\" for SVM \\\"{svm name}\\\". Reason: {Reason of failure}. ";
| 655551     | "The specified path \\\"{path name}\\\" does not exist in the namespace belonging to SVM \\\"{svm.name}\\\". ";
| 92406161   | "Failed to enable locking on bucket \\\"{bucket name}\\\" in Vserver \\\"{vserver name}\\\". Enabling locking on a bucket requires an effective cluster version of 9.14.1 or later.";
| 92406164   | "Initializing system clock for the purpose of locking an S3 bucket.";
| 92406165   | "Setting \\\"-default-retention-period\\\" on an S3 bucket requires \\\"-retention-mode\\\" to be set to \\\"compliance\\\" or \\\"governance\\\".";
| 92406166   | "Cannot enable locking on a NAS bucket.";
| 92406170   | "Cannot set \\\"-default-retention-period\\\" on object store bucket \\\"{0}\\\" in Vserver \\\"{1}\\\". Setting the default retention period on an object store bucket requires an effective cluster version of 9.14.1 or later.";
| 92406171   | "Cannot set \\\"{retention_mode}\\\" to \\\"compliance\\\" in a MetroCluster configuration";
| 92406174   | "Internal error. Failed to complete bucket create workflow with \\\"-retention-mode\\\" set to \\\"compliance\\\" or \\\"governance\\\". Reason: {0}";
| 92406175   | "The SnapLock compliance clock is not running. Use the \\\"snaplock compliance-clock initialize\\\" command to initialize the compliance clock, and then try the operation again.";
| 92406176   | "The SnapLock compliance clock is not running on the MetroCluster partner cluster. Use the \\\"snaplock compliance-clock initialize\\\" command to initialize the compliance clock on the MetroCluster partner cluster, and then try the operation again.";
| 92406230   | "The value for \\\"retention.default_period\\\" parameter for object store bucket \\\"<bucket>\\\" cannot be greater than the maximum lock retention period set in the object store server for SVM \\\"<SVM>\\\". Check the maximum allowed lock retention period present in the object store server for SVM \\\"<SVM>\\\" and try the operation again.";
| 92406236   | "The value for \\\"retention.default_period\\\" parameter for object store bucket \\\"<bucket>\\\" cannot be less than the minimum lock retention period set in the object store server for SVM \\\"<SVM>\\\". Check the minimum allowed lock retention period present in the object store server for SVM \\\"<SVM>\\\" and try the operation again.";
| 92406217   | "The specified \"allowed_headers\" is not valid because it contains more than one wild card (\"*\") character.";
| 92406224   | "A Cross-Origin Resource Sharing (CORS) rule must have an origin and HTTP method specified.";
| 92406222   | "Cannot specify Cross-Origin Resource Sharing (CORS) configuration for object store bucket \\\"<bucket>\\\" on SVM \\\"<SVM>\\\". Specifying such configuration is supported on object store volumes created in ONTAP 9.8 or later releases only.";
| 92406211   | "The specified method \"DONE\" is not valid. Valid methods are GET, PUT, DELETE, HEAD, and POST.";
| 92405863   | "Failed to create CORS rules for bucket \"bb1\". Reason: \"Field \"index\" cannot be specified for this operation.\". Resolve all the issues and retry the operation.";
| 92406228   | "Cannot exceed the maximum limit of 100 Cross-Origin Resource Sharing (CORS) rules per S3 bucket \\\"<bucket>\\\" in SVM \\\"<SVM>\\\".";;
*/
type S3BucketSvmCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this s3 bucket svm create default response has a 2xx status code
func (o *S3BucketSvmCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this s3 bucket svm create default response has a 3xx status code
func (o *S3BucketSvmCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this s3 bucket svm create default response has a 4xx status code
func (o *S3BucketSvmCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this s3 bucket svm create default response has a 5xx status code
func (o *S3BucketSvmCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this s3 bucket svm create default response a status code equal to that given
func (o *S3BucketSvmCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the s3 bucket svm create default response
func (o *S3BucketSvmCreateDefault) Code() int {
	return o._statusCode
}

func (o *S3BucketSvmCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/s3/services/{svm.uuid}/buckets][%d] s3_bucket_svm_create default %s", o._statusCode, payload)
}

func (o *S3BucketSvmCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /protocols/s3/services/{svm.uuid}/buckets][%d] s3_bucket_svm_create default %s", o._statusCode, payload)
}

func (o *S3BucketSvmCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3BucketSvmCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
