// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package efs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticfilesystem-2015-02-01/DeleteFileSystemRequest
type DeleteFileSystemInput struct {
	_ struct{} `type:"structure"`

	// The ID of the file system you want to delete.
	//
	// FileSystemId is a required field
	FileSystemId *string `location:"uri" locationName:"FileSystemId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteFileSystemInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteFileSystemInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteFileSystemInput"}

	if s.FileSystemId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FileSystemId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteFileSystemInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FileSystemId != nil {
		v := *s.FileSystemId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FileSystemId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticfilesystem-2015-02-01/DeleteFileSystemOutput
type DeleteFileSystemOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteFileSystemOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteFileSystemOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteFileSystem = "DeleteFileSystem"

// DeleteFileSystemRequest returns a request value for making API operation for
// Amazon Elastic File System.
//
// Deletes a file system, permanently severing access to its contents. Upon
// return, the file system no longer exists and you can't access any contents
// of the deleted file system.
//
// You can't delete a file system that is in use. That is, if the file system
// has any mount targets, you must first delete them. For more information,
// see DescribeMountTargets and DeleteMountTarget.
//
// The DeleteFileSystem call returns while the file system state is still deleting.
// You can check the file system deletion status by calling the DescribeFileSystems
// operation, which returns a list of file systems in your account. If you pass
// file system ID or creation token for the deleted file system, the DescribeFileSystems
// returns a 404 FileSystemNotFound error.
//
// This operation requires permissions for the elasticfilesystem:DeleteFileSystem
// action.
//
//    // Example sending a request using DeleteFileSystemRequest.
//    req := client.DeleteFileSystemRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticfilesystem-2015-02-01/DeleteFileSystem
func (c *Client) DeleteFileSystemRequest(input *DeleteFileSystemInput) DeleteFileSystemRequest {
	op := &aws.Operation{
		Name:       opDeleteFileSystem,
		HTTPMethod: "DELETE",
		HTTPPath:   "/2015-02-01/file-systems/{FileSystemId}",
	}

	if input == nil {
		input = &DeleteFileSystemInput{}
	}

	req := c.newRequest(op, input, &DeleteFileSystemOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteFileSystemRequest{Request: req, Input: input, Copy: c.DeleteFileSystemRequest}
}

// DeleteFileSystemRequest is the request type for the
// DeleteFileSystem API operation.
type DeleteFileSystemRequest struct {
	*aws.Request
	Input *DeleteFileSystemInput
	Copy  func(*DeleteFileSystemInput) DeleteFileSystemRequest
}

// Send marshals and sends the DeleteFileSystem API request.
func (r DeleteFileSystemRequest) Send(ctx context.Context) (*DeleteFileSystemResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteFileSystemResponse{
		DeleteFileSystemOutput: r.Request.Data.(*DeleteFileSystemOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteFileSystemResponse is the response type for the
// DeleteFileSystem API operation.
type DeleteFileSystemResponse struct {
	*DeleteFileSystemOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteFileSystem request.
func (r *DeleteFileSystemResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
