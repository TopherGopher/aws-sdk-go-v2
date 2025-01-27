// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/DeleteCommentRequest
type DeleteCommentInput struct {
	_ struct{} `type:"structure"`

	// Amazon WorkDocs authentication token. Do not set this field when using administrative
	// API actions, as in accessing the API using AWS credentials.
	AuthenticationToken *string `location:"header" locationName:"Authentication" min:"1" type:"string"`

	// The ID of the comment.
	//
	// CommentId is a required field
	CommentId *string `location:"uri" locationName:"CommentId" min:"1" type:"string" required:"true"`

	// The ID of the document.
	//
	// DocumentId is a required field
	DocumentId *string `location:"uri" locationName:"DocumentId" min:"1" type:"string" required:"true"`

	// The ID of the document version.
	//
	// VersionId is a required field
	VersionId *string `location:"uri" locationName:"VersionId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteCommentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCommentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteCommentInput"}
	if s.AuthenticationToken != nil && len(*s.AuthenticationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationToken", 1))
	}

	if s.CommentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CommentId"))
	}
	if s.CommentId != nil && len(*s.CommentId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CommentId", 1))
	}

	if s.DocumentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DocumentId"))
	}
	if s.DocumentId != nil && len(*s.DocumentId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DocumentId", 1))
	}

	if s.VersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VersionId"))
	}
	if s.VersionId != nil && len(*s.VersionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VersionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteCommentInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AuthenticationToken != nil {
		v := *s.AuthenticationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Authentication", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CommentId != nil {
		v := *s.CommentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "CommentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DocumentId != nil {
		v := *s.DocumentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DocumentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "VersionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/DeleteCommentOutput
type DeleteCommentOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteCommentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteCommentOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteComment = "DeleteComment"

// DeleteCommentRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Deletes the specified comment from the document version.
//
//    // Example sending a request using DeleteCommentRequest.
//    req := client.DeleteCommentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/DeleteComment
func (c *Client) DeleteCommentRequest(input *DeleteCommentInput) DeleteCommentRequest {
	op := &aws.Operation{
		Name:       opDeleteComment,
		HTTPMethod: "DELETE",
		HTTPPath:   "/api/v1/documents/{DocumentId}/versions/{VersionId}/comment/{CommentId}",
	}

	if input == nil {
		input = &DeleteCommentInput{}
	}

	req := c.newRequest(op, input, &DeleteCommentOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteCommentRequest{Request: req, Input: input, Copy: c.DeleteCommentRequest}
}

// DeleteCommentRequest is the request type for the
// DeleteComment API operation.
type DeleteCommentRequest struct {
	*aws.Request
	Input *DeleteCommentInput
	Copy  func(*DeleteCommentInput) DeleteCommentRequest
}

// Send marshals and sends the DeleteComment API request.
func (r DeleteCommentRequest) Send(ctx context.Context) (*DeleteCommentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteCommentResponse{
		DeleteCommentOutput: r.Request.Data.(*DeleteCommentOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteCommentResponse is the response type for the
// DeleteComment API operation.
type DeleteCommentResponse struct {
	*DeleteCommentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteComment request.
func (r *DeleteCommentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
