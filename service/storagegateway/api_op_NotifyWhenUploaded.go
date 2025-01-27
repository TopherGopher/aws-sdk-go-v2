// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/NotifyWhenUploadedInput
type NotifyWhenUploadedInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the file share.
	//
	// FileShareARN is a required field
	FileShareARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s NotifyWhenUploadedInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *NotifyWhenUploadedInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "NotifyWhenUploadedInput"}

	if s.FileShareARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("FileShareARN"))
	}
	if s.FileShareARN != nil && len(*s.FileShareARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("FileShareARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/NotifyWhenUploadedOutput
type NotifyWhenUploadedOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the file share.
	FileShareARN *string `min:"50" type:"string"`

	// The randomly generated ID of the notification that was sent. This ID is in
	// UUID format.
	NotificationId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s NotifyWhenUploadedOutput) String() string {
	return awsutil.Prettify(s)
}

const opNotifyWhenUploaded = "NotifyWhenUploaded"

// NotifyWhenUploadedRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Sends you notification through CloudWatch Events when all files written to
// your file share have been uploaded to Amazon S3.
//
// AWS Storage Gateway can send a notification through Amazon CloudWatch Events
// when all files written to your file share up to that point in time have been
// uploaded to Amazon S3. These files include files written to the file share
// up to the time that you make a request for notification. When the upload
// is done, Storage Gateway sends you notification through an Amazon CloudWatch
// Event. You can configure CloudWatch Events to send the notification through
// event targets such as Amazon SNS or AWS Lambda function. This operation is
// only supported for file gateways.
//
// For more information, see Getting File Upload Notification in the Storage
// Gateway User Guide (https://docs.aws.amazon.com/storagegateway/latest/userguide/monitoring-file-gateway.html#get-upload-notification).
//
//    // Example sending a request using NotifyWhenUploadedRequest.
//    req := client.NotifyWhenUploadedRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/NotifyWhenUploaded
func (c *Client) NotifyWhenUploadedRequest(input *NotifyWhenUploadedInput) NotifyWhenUploadedRequest {
	op := &aws.Operation{
		Name:       opNotifyWhenUploaded,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &NotifyWhenUploadedInput{}
	}

	req := c.newRequest(op, input, &NotifyWhenUploadedOutput{})
	return NotifyWhenUploadedRequest{Request: req, Input: input, Copy: c.NotifyWhenUploadedRequest}
}

// NotifyWhenUploadedRequest is the request type for the
// NotifyWhenUploaded API operation.
type NotifyWhenUploadedRequest struct {
	*aws.Request
	Input *NotifyWhenUploadedInput
	Copy  func(*NotifyWhenUploadedInput) NotifyWhenUploadedRequest
}

// Send marshals and sends the NotifyWhenUploaded API request.
func (r NotifyWhenUploadedRequest) Send(ctx context.Context) (*NotifyWhenUploadedResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NotifyWhenUploadedResponse{
		NotifyWhenUploadedOutput: r.Request.Data.(*NotifyWhenUploadedOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NotifyWhenUploadedResponse is the response type for the
// NotifyWhenUploaded API operation.
type NotifyWhenUploadedResponse struct {
	*NotifyWhenUploadedOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NotifyWhenUploaded request.
func (r *NotifyWhenUploadedResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
