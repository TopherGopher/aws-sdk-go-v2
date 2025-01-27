// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworkscm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/DeleteBackupRequest
type DeleteBackupInput struct {
	_ struct{} `type:"structure"`

	// The ID of the backup to delete. Run the DescribeBackups command to get a
	// list of backup IDs. Backup IDs are in the format ServerName-yyyyMMddHHmmssSSS.
	//
	// BackupId is a required field
	BackupId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBackupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBackupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBackupInput"}

	if s.BackupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/DeleteBackupResponse
type DeleteBackupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteBackupOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteBackup = "DeleteBackup"

// DeleteBackupRequest returns a request value for making API operation for
// AWS OpsWorks CM.
//
// Deletes a backup. You can delete both manual and automated backups. This
// operation is asynchronous.
//
// An InvalidStateException is thrown when a backup deletion is already in progress.
// A ResourceNotFoundException is thrown when the backup does not exist. A ValidationException
// is thrown when parameters of the request are not valid.
//
//    // Example sending a request using DeleteBackupRequest.
//    req := client.DeleteBackupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/DeleteBackup
func (c *Client) DeleteBackupRequest(input *DeleteBackupInput) DeleteBackupRequest {
	op := &aws.Operation{
		Name:       opDeleteBackup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteBackupInput{}
	}

	req := c.newRequest(op, input, &DeleteBackupOutput{})
	return DeleteBackupRequest{Request: req, Input: input, Copy: c.DeleteBackupRequest}
}

// DeleteBackupRequest is the request type for the
// DeleteBackup API operation.
type DeleteBackupRequest struct {
	*aws.Request
	Input *DeleteBackupInput
	Copy  func(*DeleteBackupInput) DeleteBackupRequest
}

// Send marshals and sends the DeleteBackup API request.
func (r DeleteBackupRequest) Send(ctx context.Context) (*DeleteBackupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBackupResponse{
		DeleteBackupOutput: r.Request.Data.(*DeleteBackupOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBackupResponse is the response type for the
// DeleteBackup API operation.
type DeleteBackupResponse struct {
	*DeleteBackupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBackup request.
func (r *DeleteBackupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
